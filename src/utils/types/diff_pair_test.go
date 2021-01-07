package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiff(t *testing.T) {
	var first, second, add, del map[string]interface{}

	first = map[string]interface{}{"a": true, "b": true}
	second = map[string]interface{}{"c": true, "b": true}
	add, del = Diff(first, second)

	assert.Equal(t, 1, len(add))
	assert.NotNil(t, add["a"])
	assert.Equal(t, 1, len(del))
	assert.NotNil(t, del["c"])

	first = map[string]interface{}{"a": true, "b": true}
	second = map[string]interface{}{"a": true, "b": true}
	add, del = Diff(first, second)

	assert.Equal(t, 0, len(add))
	assert.Equal(t, 0, len(del))
}

func TestHasDiff(t *testing.T) {
	var first, second map[string]interface{}

	first = map[string]interface{}{"a": true, "b": true}
	second = map[string]interface{}{"c": true, "b": true}
	has := HasDiff(first, second)
	assert.True(t, has)

	first = map[string]interface{}{"a": true, "b": true}
	second = map[string]interface{}{"a": true, "b": true}
	has = HasDiff(first, second)
	assert.False(t, has)
}
