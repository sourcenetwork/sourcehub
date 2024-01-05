package embedded

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLocalACPBuilds(t *testing.T) {
	acp, err := NewLocalACP()

	require.Nil(t, err)
	require.NotNil(t, acp.GetMsgService())
	require.NotNil(t, acp.GetQueryService())
}
