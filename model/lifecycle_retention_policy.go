package lifecyclemodel

import "time"

type RetentionClass string

const (
	RetentionClassProduct    RetentionClass = "product_retention"
	RetentionClassLegalAudit RetentionClass = "legal_audit_retention"
	RetentionClassTechnical  RetentionClass = "technical_ttl"
	RetentionClassUserErase  RetentionClass = "user_requested_erase"
)

type Sensitivity string

const (
	SensitivityInternal  Sensitivity = "internal"
	SensitivityPII       Sensitivity = "pii"
	SensitivitySensitive Sensitivity = "sensitive"
	SensitivityFinancial Sensitivity = "financial"
	SensitivitySecurity  Sensitivity = "security"
	SensitivityAudit     Sensitivity = "audit"
)

type BackupBehavior string

const (
	BackupBehaviorStandard         BackupBehavior = "standard_restore_then_reconcile"
	BackupBehaviorDelayedErase     BackupBehavior = "delayed_erase_after_restore"
	BackupBehaviorComplianceLocked BackupBehavior = "compliance_locked"
)

type EraseBehavior string

const (
	EraseBehaviorDelete      EraseBehavior = "delete"
	EraseBehaviorAnonymize   EraseBehavior = "anonymize"
	EraseBehaviorNotEligible EraseBehavior = "not_eligible"
)

type RetentionPolicy struct {
	Key                     string                   `json:"key"`
	Version                 string                   `json:"version"`
	Owner                   string                   `json:"owner"`
	Class                   RetentionClass           `json:"class"`
	Sensitivity             []Sensitivity            `json:"sensitivity,omitempty"`
	DefaultRetention        time.Duration            `json:"default_retention"`
	MinimumRetention        time.Duration            `json:"minimum_retention"`
	StatusRetention         map[string]time.Duration `json:"status_retention,omitempty"`
	ReplayWindow            time.Duration            `json:"replay_window,omitempty"`
	WorkspaceMayExtend      bool                     `json:"workspace_may_extend"`
	WorkspaceMayReduce      bool                     `json:"workspace_may_reduce"`
	LegalHoldEligible       bool                     `json:"legal_hold_eligible"`
	BackupBehavior          BackupBehavior           `json:"backup_behavior"`
	EraseBehavior           EraseBehavior            `json:"erase_behavior"`
	RequiredReferenceChecks []string                 `json:"required_reference_checks,omitempty"`
}

type WorkspaceRetentionOverride struct {
	WorkspaceID string        `json:"workspace_id"`
	Retention   time.Duration `json:"retention"`
}
