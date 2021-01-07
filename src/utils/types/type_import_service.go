package types

import "time"

type ImportService struct {
	ImportServiceDef     `yaml:",inline"`
	ImportServiceRuntime `yaml:",inline"`
}

type ImportServiceDef struct {
	Meta `yaml:",inline"`
	Spec ImportServiceSpec `yaml:"spec" json:"spec"`
}

type ImportServiceSpec struct {
	Services []ServiceEntity `yaml:"services" json:"services"`
}

type ServiceEntity struct {
	// used to refer imported service instance in others resources, for example {{.importService.ReferValue.endpointXXX}}
	Refer string `yaml:"refer" json:"refer"`
	// service name
	Service string `yaml:"service" json:"service"`
	// service version
	Version string `yaml:"version" json:"version"`
	// service instance class
	InstanceClass string `yaml:"instanceClass" json:"instanceClass"`
	// extra params to create imported service, which not included in flavor
	Params map[string]interface{} `yaml:"params" json:"params"`
	// service instance id
	InstanceId string `yaml:"instanceId" json:"instanceId"`
}

type ImportServiceRuntime struct {
	Status ImportServiceStatus `yaml:"status" json:"status"`
}

type ImportServiceStatus struct {
	Createtime time.Time `yaml:"createtime" json:"createtime"`
	Updatetime time.Time `yaml:"updatetime" json:"updatetime"`
	Deletetime time.Time `yaml:"deletetime" json:"deletetime"`
}

func (t *ImportService) GetDef() interface{} {
	return &(t.ImportServiceDef)
}

func (t *ImportService) SetPhase(phase string) {
}

func (t *ImportService) GetPhase() string {
	return ""
}

func (t *ImportService) SetRequestId(requestId string) {
}

func (t *ImportService) RefreshUpdatetime() {
	t.Status.Updatetime = time.Now().Round(0)
}
