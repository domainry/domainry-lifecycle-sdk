package lifecyclesdk

import (
	"testing"

	"github.com/domainry/domainry-foundation/apperror"
)

func TestErrorClassificationContract(t *testing.T) {
	tests := []struct {
		class ErrorClass
		kind  apperror.ErrorKind
		code  string
	}{
		{ErrorClassBadRequest, apperror.KindBadRequest, "lifecycle.validation_failed"},
		{ErrorClassForbidden, apperror.KindForbidden, "lifecycle.permission_denied"},
		{ErrorClassConflict, apperror.KindConflict, "lifecycle.execution_fence_rejected"},
		{ErrorClassUnavailable, apperror.KindUnavailable, "lifecycle.database_unavailable"},
	}
	for _, test := range tests {
		err := &Error{Class: test.class, Code: test.code}
		if apperror.KindOf(err) != test.kind || apperror.CodeOf(err) != test.code {
			t.Fatalf("class=%q kind=%q code=%q", test.class, apperror.KindOf(err), apperror.CodeOf(err))
		}
	}
}
