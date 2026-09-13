package contract

import (
	"context"
	"time"

	"github.com/domainry/domainry-lifecycle-sdk/access"
	model "github.com/domainry/domainry-lifecycle-sdk/model"
)

// AccountErasureApproval is trusted Action provenance, resolved and frozen by
// Runtime. Project handlers never select a workspace or an Identity user here.
type AccountErasureApproval struct {
	WorkspaceID string `json:"workspace_id"`
	RequestID   string `json:"request_id"`
	SubjectID   string `json:"subject_id"`
	RequestedBy string `json:"requested_by"`
	ApprovedBy  string `json:"approved_by"`
	OwnerOrgID  string `json:"owner_org_id"`
	ActionKey   string `json:"action_key"`
	ApprovalID  string `json:"approval_id"`
	BindingKey  string `json:"binding_key"`
	ObjectKey   string `json:"object_key"`
	ProfileID   string `json:"profile_id"`
}

type AccountErasureReference struct {
	WorkspaceID string
	RequestID   string
	OwnerOrgID  string
	BindingKey  string
	ObjectKey   string
	ProfileID   string
}

// AccountErasures queues approved business requests inside the caller's Action
// transaction. The existing Lifecycle executor performs irreversible effects
// after commit; queued/failed requests are never reported as erased.
type AccountErasures interface {
	StageApprovedAccountErasure(context.Context, AccountErasureApproval, access.SystemScope) (model.SubjectRequest, error)
	GetAccountErasure(context.Context, AccountErasureReference, access.SystemScope) (model.SubjectRequest, error)
	ProcessApprovedAccountErasures(context.Context, string, int, time.Time, access.SystemScope) (int, error)
}
