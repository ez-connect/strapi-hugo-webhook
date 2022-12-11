package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteKeys(t *testing.T) {
	req := map[string]any{
		"a": 1,
		"b": "2",
		"c": map[string]any{"x": 1, "y": 2},
		"d": 4,
	}

	res := deleteKeys(req, "a")
	assert.Nil(t, res["a"])

	res = deleteKeys(req, "b", "c")
	assert.Nil(t, res["b"])
	assert.Nil(t, res["c"])
	assert.Equal(t, 4, res["d"])

	assert.NotNil(t, req["a"])
	assert.NotNil(t, req["b"])
	assert.NotNil(t, req["d"])
}
