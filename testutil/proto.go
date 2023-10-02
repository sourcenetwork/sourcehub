package testutil

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

func ProtoEq[Msg proto.Message](t *testing.T, want, got Msg) {
	opt := protocmp.Transform()

	if diff := cmp.Diff(want, got, opt); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
