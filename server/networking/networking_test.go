package networking

import (
	"testing"
	"webserver/server/router"
)

func Test_initListener(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		rcs []router.ControlledRoutes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initListener("3000", tt.rcs)
		})
	}
}
