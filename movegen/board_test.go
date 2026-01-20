package movegen_test

import (
	"testing"

	"github.com/amit152116/chess_server/movegen"
)

func TestBoard_GetFEN(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		fen  string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := movegen.LoadFen(tt.fen)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			got := b.GetFEN()
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GetFEN() = %v, want %v", got, tt.want)
			}
		})
	}
}
