package lifecyclesdk

// Lifecycle Action keys are stable executable identities. HTTP paths and SDK
// method names are bindings and may change without changing these keys.
const (
	ActionLifecyclePoliciesList              = "lifecycle.policies.list"
	ActionLifecyclePoliciesPublish           = "lifecycle.policies.publish"
	ActionLifecycleLegalHoldsList            = "lifecycle.legal_holds.list"
	ActionLifecycleLegalHoldsCreate          = "lifecycle.legal_holds.create"
	ActionLifecycleLegalHoldsEnd             = "lifecycle.legal_holds.end"
	ActionLifecycleCleanupPreview            = "lifecycle.cleanup.preview"
	ActionLifecycleCleanupJobsCreate         = "lifecycle.cleanup_jobs.create"
	ActionLifecycleCleanupJobsProcess        = "lifecycle.cleanup_jobs.process"
	ActionLifecycleMetricsRead               = "lifecycle.metrics.read"
	ActionLifecycleArchiveList               = "lifecycle.archive.list"
	ActionLifecycleSubjectRequestsCreate     = "lifecycle.subject_requests.create"
	ActionLifecycleSubjectRequestsList       = "lifecycle.subject_requests.list"
	ActionLifecycleSubjectRequestsRead       = "lifecycle.subject_requests.read"
	ActionLifecycleSubjectRequestsVerify     = "lifecycle.subject_requests.verify"
	ActionLifecycleSubjectRequestsPreview    = "lifecycle.subject_requests.preview"
	ActionLifecycleSubjectRequestsApprove    = "lifecycle.subject_requests.approve"
	ActionLifecycleSubjectRequestsExecute    = "lifecycle.subject_requests.execute"
	ActionLifecycleSubjectExportsDownload    = "lifecycle.subject_exports.download"
	ActionLifecycleExternalErasuresList      = "lifecycle.external_erasures.list"
	ActionLifecycleExternalErasuresReconcile = "lifecycle.external_erasures.reconcile"
	ActionLifecycleDeletionsReplay           = "lifecycle.deletions.replay"
	ActionLifecycleDefaultPoliciesInstall    = "lifecycle.default_policies.install"
	ActionLifecycleHealthRead                = "lifecycle.health.read"
	ActionLifecycleWorkersTick               = "lifecycle.workers.tick"
)

const (
	CapabilityLifecycleGovernance = "lifecycle.governance"
	CapabilityLifecycleSubjects   = "lifecycle.subjects"
	CapabilityLifecycleOperations = "lifecycle.operations"
)
