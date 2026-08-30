// Package lifecyclesdk defines the deployment-neutral Lifecycle boundary.
package lifecyclesdk

import (
	"context"
	"fmt"
	"strings"

	"github.com/domainry/domainry-lifecycle-sdk/contract"
	"github.com/domainry/domainry-lifecycle-sdk/modulehost"
	"github.com/domainry/domainry-lifecycle-sdk/repository"
)

type DeploymentMode string

const (
	DeploymentModeModule DeploymentMode = "module"
	ProtocolVersionV1                   = "domainry-lifecycle-protocol-v1"
)

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

type Binding interface {
	Descriptor() Descriptor
	Repository() repository.LifecycleRepository
	UploadArtifacts(UploadArtifactOptions) (contract.UploadArtifactStore, error)
	SubjectArtifacts(string) (contract.SubjectArtifactStore, error)
	ArchiveStore() contract.ArchiveStore
	WithinTransaction(context.Context, func(context.Context) error) error
	Close(context.Context) error
}
