package testutil

import (
	"time"

	gogotypes "github.com/cosmos/gogoproto/types"
)

// MustDateTimeToProto parses a time.DateTime (YYYY-MM-DD HH:MM:SS) timestamp
// and converts into a proto Timestamp.
// Panics if input is invalid
func MustDateTimeToProto(timestamp string) *gogotypes.Timestamp {
	t, err := time.Parse(time.DateTime, timestamp)
	if err != nil {
		panic(err)
	}

	ts, err := gogotypes.TimestampProto(t)
	if err != nil {
		panic(err)
	}

	return ts
}
