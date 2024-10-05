package basic

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	var (
		input = 1
		expectOutput = 2
	)

	actualOuput := AddOne(input)
	require.Equal(t, expectOutput, actualOuput)
}



// require versus assert:
// - require failed, interupt function immediately
// - assert failed, continue run next statements