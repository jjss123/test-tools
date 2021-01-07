package types

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

type dummy struct {
	Value string
}

func TestJsonMarshal_happy(t *testing.T) {
	//setup
	v := dummy{Value: "abc"}

	//exercise
	data, err := JsonMarshal(v)

	//verify
	assert.Nil(t, err)
	assert.Equal(t, data, "{\"Value\":\"abc\"}")
}

func TestJsonUnmarshal_happy(t *testing.T) {
	//setup
	data := "{\"Value\":\"abc\"}"
	v := new(dummy)

	//exercise
	err := JsonUnmarshal(data, v)

	//verify
	assert.Nil(t, err)
	assert.Equal(t, v.Value, "abc")
}

func TestJsonUnmarshal_bad(t *testing.T) {
	//setup
	data := "{\"Value\":\"abc}"
	v := new(dummy)

	//exercise
	err := JsonUnmarshal(data, v)

	//verify
	assert.NotNil(t, err)
}

func isTypeOf(ele interface{}, kind string) bool {
	return strings.HasSuffix(reflect.TypeOf(ele).String(), "."+kind)
}

func TestMakeResourceByKind(t *testing.T) {
	test := func(kind string) {
		ele := makeResourceByKind(kind)
		assert.True(t, isTypeOf(ele, kind))
	}

	test("Container")
	test("Job")
	test("Space")
	test("ReplicaSet")
}

func TestSmartJsonUnmarshal_happy(t *testing.T) {
	//setup
	data := "{\"kind\":\"Container\"}"

	//exercise
	v, err := SmartJsonUnmarshal(data)

	//verify
	assert.Nil(t, err)
	assert.True(t, isTypeOf(v, "Container"))
}

func TestSmartJsonUnmarshal_unknownResource(t *testing.T) {
	//setup
	data := "{\"kind\":\"Dummy\"}"

	//exercise
	_, err := SmartJsonUnmarshal(data)

	//verify
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), UnknownResource)
}
