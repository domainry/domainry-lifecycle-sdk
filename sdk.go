// Package lifecyclesdk defines the deployment-neutral Lifecycle boundary.
package lifecyclesdk

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/domainry/domainry-foundation/modulecapability"
	"github.com/domainry/domainry-lifecycle-sdk/access"
	"github.com/domainry/domainry-lifecycle-sdk/contract"
	model "github.com/domainry/domainry-lifecycle-sdk/model"
	"github.com/domainry/domainry-lifecycle-sdk/modulehost"
)

type DeploymentMode string

const (
	DeploymentModeModule    DeploymentMode = "module"
	ProtocolVersionV1                      = "domainry-lifecycle-protocol-v1"
	PermissionPolicyManage                 = "lifecycle.policy.manage"
	PermissionCleanupRun                   = "lifecycle.cleanup.run"
	PermissionSubjectManage                = "lifecycle.subject.manage"
)

type Error struct {
	StatusCode int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message,omitempty"`
	Retryable  bool   `json:"retryable,omitempty"`
	Cause      error  `json:"-"`
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	if strings.TrimSpace(e.Message) != "" {
		return e.Code + ": " + e.Message
	}
	return e.Code
}

func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Cause
}

type ApplicationRef struct {
	RuntimeID string `json:"runtime_id"`
}

func (r ApplicationRef) Validate() error {
	if strings.TrimSpace(r.RuntimeID) == "" {
		return fmt.Errorf("Lifecycle runtime identity is required")
	}
	return nil
}

type Capabilities struct {
	Governance      bool
	SubjectRequests bool
	RetentionWorker bool
	UploadArtifacts bool
	ArchiveEvidence bool
}

type Descriptor struct {
	ProtocolVersion string
	Mode            DeploymentMode
	Capabilities    Capabilities
}

func (d Descriptor) Validate() error {
	if d.ProtocolVersion != ProtocolVersionV1 || d.Mode != DeploymentModeModule || !d.Capabilities.Governance || !d.Capabilities.SubjectRequests {
		return fmt.Errorf("invalid Lifecycle module descriptor")
	}
	return nil
}

type UploadArtifactOptions struct {
	Root              string
	Fields            contract.UploadFieldCatalog
	References        contract.UploadArtifactReferenceResolver
	ExpiredReferences contract.ExpiredUploadReferenceCleaner
}

type Factory interface {
	OpenModule(context.Context, ApplicationRef, modulehost.Host) (Binding, error)
}

type OwnerExtensions struct {
	Executors       []contract.OwnerLifecycleExecutor
	SubjectResolver contract.SubjectIdentityResolver
	SubjectHandlers []contract.SubjectExecutionHandler
	ExternalErasure contract.ExternalErasureHandler
	Artifacts       contract.SubjectArtifactStore
	UploadArtifacts contract.UploadArtifactStore
}

type Governance interface {
	PublishPolicy(context.Context, model.PolicyVersion, access.Principal) (model.PolicyVersion, error)
	ListPolicies(context.Context, access.Principal) ([]model.PolicyVersion, error)
	CreateLegalHold(context.Context, model.LegalHold, access.Principal) (model.LegalHold, error)
	EndLegalHold(context.Context, string, string, string, string, time.Time, access.Principal) (model.LegalHold, error)
	PreviewCleanup(context.Context, string, string, access.Principal, time.Time) (contract.CleanupPreview, error)
	CreateCleanupJob(context.Context, model.CleanupJob, access.Principal) (model.CleanupJob, error)
	ProcessCleanupJob(context.Context, string, string, string, time.Duration, int, time.Time, access.Principal) (model.CleanupJob, error)
	Metrics(context.Context, access.Principal, time.Time) (model.Metrics, error)
	ListArchiveEntries(context.Context, string, int, access.Principal) ([]model.ArchiveEntry, error)
	CreateSubjectRequest(context.Context, model.SubjectRequest, access.Principal) (model.SubjectRequest, error)
	VerifySubjectRequest(context.Context, string, string, string, access.Principal) (model.SubjectRequest, error)
	PreviewSubjectRequest(context.Context, string, string, access.Principal) (model.SubjectRequest, error)
	ApproveSubjectRequest(context.Context, string, string, access.Principal) (model.SubjectRequest, error)
	ExecuteSubjectRequest(context.Context, string, string, access.Principal) (model.SubjectRequest, error)
	DownloadSubjectExport(context.Context, string, string, access.Principal, time.Time) (json.RawMessage, error)
	ListExternalErasures(context.Context, string, access.Principal) ([]model.ExternalErasure, error)
	ReconcileExternalErasure(context.Context, string, string, access.Principal, time.Time) (model.ExternalErasure, error)
	ReplayRegisteredDeletions(context.Context, string, int, access.Principal) (int, error)
}

type System interface {
	InstallDefaultPolicies(context.Context, string, access.Principal, time.Time) error
	Health(context.Context, access.SystemScope, time.Time) (map[string]any, error)
}

type WorkerTick struct {
	LeaseOwner string
	BatchSize  int
	JobLimit   int
	Now        time.Time
	Scope      access.SystemScope
}

type WorkerTickResult struct {
	ProcessedJobs    int
	DeletedArtifacts int
}

type LocalWorkers interface {
	Tick(context.Context, WorkerTick) (WorkerTickResult, error)
}

type Binding interface {
	modulecapability.Binding
	Descriptor() Descriptor
	BindOwners(context.Context, OwnerExtensions) error
	Governance() Governance
	System() System
	LocalWorkers() (LocalWorkers, bool)
	UploadArtifacts(UploadArtifactOptions) (contract.UploadFileArtifactStore, error)
	SubjectArtifacts(string) (contract.SubjectArtifactStore, error)
	ArchiveStore() contract.ArchiveStore
	Close(context.Context) error
}
