package types

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

type containerSpecOps struct {
	GetImage   func(resource Resource) string
	SetImageId func(resource Resource, imageId string)
}

var (
	resource2ContainerSpecOps = map[string]containerSpecOps{
		ContainerResource: {
			GetImage: func(resource Resource) string {
				c := resource.(*Container)
				return c.Spec.Image
			},
			SetImageId: func(resource Resource, imageId string) {
				c := resource.(*Container)
				c.Spec.ImageID = imageId
			},
		},
		ReplicaSetResource: {
			GetImage: func(resource Resource) string {
				r := resource.(*ReplicaSet)
				return r.Spec.Template.Image
			},
			SetImageId: func(resource Resource, imageId string) {
				r := resource.(*ReplicaSet)
				r.Spec.Template.ImageID = imageId
			},
		},
		JobResource: {
			GetImage: func(resource Resource) string {
				j := resource.(*Job)
				return j.Spec.Template.Image
			},
			SetImageId: func(resource Resource, imageId string) {
				j := resource.(*Job)
				j.Spec.Template.ImageID = imageId
			},
		},
	}
)

var GetImage = func(resource Resource) string {
	ops, exist := resource2ContainerSpecOps[resource.GetMeta().Kind]
	if exist {
		return ops.GetImage(resource)
	}
	return ""
}

var HasImageDefinition = func(kind string) bool {
	return HasContainerSpec(kind)
}

var HasContainerSpec = func(kind string) bool {
	_, exist := resource2ContainerSpecOps[kind]
	return exist
}

var SetImageId = func(resource Resource, imageId string) {
	ops, exist := resource2ContainerSpecOps[resource.GetMeta().Kind]
	if exist {
		ops.SetImageId(resource, imageId)
	}
}

var MakeSecurityGroupName = func(spaceName string, t string) string {
	return makeSGName(spaceName, t)
}

var GetExternalSGName = func(spaceName string) string {
	return makeSGName(spaceName, SG_EXTERNAL)
}

var GetInternalSGName = func(spaceName string) string {
	return makeSGName(spaceName, SG_INTERNAL)
}

const SecurityGroupNameSuffix = "SG"

var makeSGName = func(spaceName string, t string) string {
	return spaceName + "-" + t + SecurityGroupNameSuffix
}

const BlockStoreNameSuffix = "-BS"

var MakeBlockStoreResourceName = func(containerName string, volName string) string {
	return containerName + "-" + volName + BlockStoreNameSuffix
}

const scheduleSuffix = ".SCHEDULE"

var MakeScheduleName = func(spaceName string) string {
	return spaceName + scheduleSuffix
}

const ImportServiceNameSuffix = "-IS"

var MakeImportServiceName = func(spaceName string) string {
	return spaceName + ImportServiceNameSuffix
}

const DefaultFloatIPSuffix = "-FI"

var MakeFloatIPName = func(nameOfFloatIPTarget string) string {
	return nameOfFloatIPTarget + DefaultFloatIPSuffix
}

var GetTargetNameOfFloatIP = func(floatIPName string) string {
	return floatIPName[:len(floatIPName)-len(DefaultFloatIPSuffix)]
}

var MakeWebAppDomain = func(metaName, domain string) string {
	return fmt.Sprintf("%s.%s", metaName, domain)
}

// get value in map by path
var GetValueByPath = func(data map[string]interface{}, path string) (value interface{}, err error) {
	pathList := strings.Split(path, ".")
	var res interface{}
	res = data
	for _, key := range pathList {
		_, ok := res.(map[string]interface{})
		if !ok {
			return nil, errors.New("path not exist")
		}
		res, ok = res.(map[string]interface{})[key]
		if !ok {
			return nil, errors.New("path not exist")
		}
	}
	return res, nil
}
