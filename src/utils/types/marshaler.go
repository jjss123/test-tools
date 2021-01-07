package types

import (
	"encoding/json"
	"errors"
	yaml2 "github.com/ghodss/yaml"
	"gopkg.in/yaml.v2"
	. "testTools/src/utils/marshaler"
)

var JsonMarshal = func(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	return string(b), err
}

var JsonUnmarshal = func(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}

var YamlMarshal = func(v interface{}) (string, error) {
	b, err := yaml.Marshal(v)
	return string(b), err
}

// make resource by kind
// return nil for unknown kind
// TODO: use reflection?
func makeResourceByKind(kind string) Resource {
	if kind == ContainerResource {
		return NewContainer()
	} else if kind == VMResource {
		return NewVM()
	} else if kind == NCResource {
		return NewNC()
	} else if kind == ReplicaSetResource {
		return NewReplicaSet()
	} else if kind == JobResource {
		return NewJob()
	} else if kind == NlbResource {
		return NewNlb()
	} else if kind == SpaceResource {
		return NewSpace()
	} else if kind == FloatIPResource {
		return NewFloatIP()
	} else if kind == SecurityGroupResource {
		return NewSecurityGroup()
	} else if kind == ScriptResource {
		return NewScript()
	} else if kind == BlockStoreResource {
		return NewBlockStore()
	} else if kind == SnapshotResource {
		return NewSnapshot()
	} else if kind == ScheduleResource {
		return NewSchedule()
	} else if kind == ImportServiceResource {
		return NewImportService()
	} else if kind == UnderlayEntryResource {
		return NewUnderlayEntry()
	} else if kind == PipelineResource {
		return NewPipeline()
	} else if kind == UpdateResource {
		return NewUpdate()
	} else if kind == SnapshotResource {
		return NewSnapshot()
	} else if kind == ContainerAloneResource {
		return NewContainerAlone()
	} else if kind == BlockStoreAloneResource {
		return NewBlockStoreAlone()
	}
	return nil
}

const UnknownResource = "Unknown resource"

// unmarshal resource in json format, return the pointer
var SmartUnmarshal = func(data string, marshaler Marshaler) (Resource, error) {

	meta := new(Meta)
	err := marshaler.Unmarshal([]byte(data), meta)
	if err != nil {
		return nil, err
	}

	if meta.Kind == UpdateResource {
		v := makeResourceByKind(UpdateResource)
		switch marshaler.(type) {

		case *YamlMarshaler:
			jsonData, err := yaml2.YAMLToJSON([]byte(data))
			if err != nil {
				return nil, err
			}
			jsonMarshaler := NewJsonMarshaler()
			err = jsonMarshaler.Unmarshal([]byte(jsonData), v)
			if err != nil {
				return nil, err
			}
			return v, nil
		default:
			err = marshaler.Unmarshal([]byte(data), v)
			if err != nil {
				return nil, err
			}
			return v, nil
		}

	} else {
		v := makeResourceByKind(meta.Kind)
		if v == nil {
			return nil, errors.New(UnknownResource)
		}

		err = marshaler.Unmarshal([]byte(data), v)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

var SmartJsonUnmarshal = func(data string) (Resource, error) {
	return SmartUnmarshal(data, NewJsonMarshaler())
}

var SmartYamlUnmarshal = func(data string) (Resource, error) {
	//jsonData, err := yaml2.YAMLToJSON([]byte(data))
	//if err != nil {
	//	return nil, err
	//}
	//return SmartUnmarshal(string(jsonData), NewJsonMarshaler())
	return SmartUnmarshal(data, NewYamlMarshaler())
}
