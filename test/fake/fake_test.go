package fake_test

import (
	"testing"

	"github.com/integration-system/isp-kit/test/fake"
	"github.com/stretchr/testify/require"
)

type SomeStruct struct {
	A string
	B bool
}

func Test(t *testing.T) {
	require := require.New(t)

	intValue := fake.It[int]()
	require.NotEmpty(intValue)

	stringSlice := fake.It[[]string]()
	require.NotEmpty(stringSlice)

	structSlice := fake.It[[]SomeStruct]()
	t.Log(structSlice)
	require.NotEmpty(structSlice)
}
