package lifecyclemodel

import "time"

type Operation string

const (
	OperationArchive Operation = "archive"
	OperationPurge   Operation = "purge"
	OperationErase   Operation = "erase"
)

type ResourceTarget struct {
	WorkspaceID  string `json:"workspace_id"`
	Owner        string `json:"owner,omitempty"`
	ResourceType string `json:"resource_type,omitempty"`
	ResourceID   string `json:"resource_id,omitempty"`
}

type LegalHold struct {
	ID            string     `json:"id"`
	WorkspaceID   string     `json:"workspace_id"`
	Owner         string     `json:"owner,omitempty"`
	ResourceType  string     `json:"resource_type,omitempty"`
	ResourceID    string     `json:"resource_id,omitempty"`
	Reason        string     `json:"reason"`
	Authority     string     `json:"authority"`
	StartsAt      time.Time  `json:"starts_at"`
	EndsAt        *time.Time `json:"ends_at,omitempty"`
	ReviewAt      time.Time  `json:"review_at"`
	AuditEvidence string     `json:"audit_evidence"`
}

type EligibilityInput struct {
	Operation           Operation
	Target              ResourceTarget
	OwnerEligible       bool
	ActiveProcess       bool
	PendingOutbox       bool
	Referenced          bool
	BackupPolicyBlocked bool
	LegalHolds          []LegalHold
	Now                 time.Time
}

type EligibilityDecision struct {
	Eligible bool
	Blockers []string
}
