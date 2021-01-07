package types

import "fmt"

const (
	TypeService         = "Service"
	TypeServiceAZ       = "ServiceAZ"
	TypeServiceRelease  = "ServiceRelease"
	TypeServiceInstance = "ServiceInstance"
)

type ServiceBlock struct {
	Enable    bool                   `json:"enable"`
	Service   string                 `json:"service"`
	AZs       []string               `json:"azs"`
	Releases  []string               `json:"releases"`
	Instances map[string]interface{} `json:"instances"`
}

func NewServiceBlock(service string) *ServiceBlock {
	block := &ServiceBlock{
		Service:   service,
		AZs:       make([]string, 0),
		Releases:  make([]string, 0),
		Instances: make(map[string]interface{}),
	}
	return block
}

func (block *ServiceBlock) IsEmpty() bool {
	if block.Enable == true {
		return false
	}
	if len(block.AZs) != 0 {
		return false
	}
	if len(block.Releases) != 0 {
		return false
	}
	if len(block.Instances) != 0 {
		return false
	}
	return true
}

func (block *ServiceBlock) Block(blockType, blockItem string) bool {
	switch blockType {
	case TypeService:
		if block.Enable == true {
			return false
		}
		block.Enable = true
	case TypeServiceAZ:
		for _, az := range block.AZs {
			if az == blockItem {
				return false
			}
		}
		block.AZs = append(block.AZs, blockItem)
	case TypeServiceRelease:
		for _, release := range block.Releases {
			if release == blockItem {
				return false
			}
		}
		block.Releases = append(block.Releases, blockItem)
	case TypeServiceInstance:
		if _, ok := block.Instances[blockItem]; ok {
			return false
		}
		block.Instances[blockItem] = true
	}
	return true
}

func (block *ServiceBlock) UnBlock(blockType, blockItem string) bool {
	switch blockType {
	case TypeService:
		if block.Enable == false {
			return false
		}
		block.Enable = false
	case TypeServiceAZ:
		newAZSlice := make([]string, 0)
		for _, az := range block.AZs {
			if az == blockItem {
				continue
			}
			newAZSlice = append(newAZSlice, az)
		}
		if len(newAZSlice) == len(block.AZs) {
			return false
		}
		block.AZs = newAZSlice
	case TypeServiceRelease:
		newReleaseSlice := make([]string, 0)
		for _, release := range block.Releases {
			if release == blockItem {
				continue
			}
			newReleaseSlice = append(newReleaseSlice, release)
		}
		if len(newReleaseSlice) == len(block.Releases) {
			return false
		}
		block.Releases = newReleaseSlice
	case TypeServiceInstance:
		if _, ok := block.Instances[blockItem]; !ok {
			return false
		}
		delete(block.Instances, blockItem)
	}
	return true
}

func (block *ServiceBlock) String() string {
	return fmt.Sprintf("%+v", *block)
}
