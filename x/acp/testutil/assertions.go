package testutil

import (
	"reflect"
	"testing"

	proto "github.com/cosmos/gogoproto/proto"

	"github.com/cosmos/cosmos-sdk/types"
)

func AssertEventEmmited(t *testing.T, ctx types.Context, event any) {
	var ev types.Event

	switch cast := event.(type) {
	case types.Event:
		ev = cast
	default:
		var err error
		ev, err = types.TypedEventToEvent(cast.(proto.Message))
		if err != nil {
			panic(err)
		}
	}

	for _, e := range ctx.EventManager().Events() {
		if reflect.DeepEqual(e, ev) {
			return
		}
	}
	t.Fatalf("EventManager did not emit wanted event: want %v", event)
}
