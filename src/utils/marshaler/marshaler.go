// wrap json/yaml marshaler into unified interface
package marshaler

import (
	json "github.com/json-iterator/go"
	"gopkg.in/yaml.v2"
)

type Marshaler interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

type JsonMarshaler struct {
}

func NewJsonMarshaler() Marshaler {
	return new(JsonMarshaler)
}

func (t *JsonMarshaler) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (t *JsonMarshaler) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

type YamlMarshaler struct {
}

func NewYamlMarshaler() Marshaler {
	return new(YamlMarshaler)
}

func (t *YamlMarshaler) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (t *YamlMarshaler) Unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}
