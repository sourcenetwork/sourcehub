package testutil

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

func ProtoEq[Msg proto.Message](t *testing.T, left, right Msg) {
	opt := protocmp.Transform()

	if diff := cmp.Diff(left, right, opt); diff != "" {
		t.Errorf("mismatch (-left +right):\n%s", diff)
	}
}
