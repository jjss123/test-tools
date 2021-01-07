package types

import "time"

// Space Phase
const (
	SpacePending           = CommonPending
	SpaceLaunching         = "Launching"
	SpaceRunning           = CommonRunning
	SpaceUpdating          = "Updating"
	SpaceUpgrading         = "Upgrading"
	SpaceUpdatingCompleted = "UpdatingCompleted"
	SpaceToDelete          = CommonToDelete
	SpaceRollbacking       = "Rollbacking"
)

const (
	DefaultNamespace = "default"
)
const (
	DefaultPortContext         = "exposedPortList"
	CtxSpaceGreenHealthSuspend = "GREEN_HEALTH_SUSPEND"
)

type SpaceUpdateEvent struct {
	Type string `yaml:"type" json:"type"`
	//json format of updating resource
	Content string `yaml:"content" json:"content"`
}

type SpaceStatus struct {
	BasicStatus           `yaml:",inline"`
	UpdateEvent           []SpaceUpdateEvent     `yaml:"updateevent" json:"updateevent"`
	UndoRequestId         string                 `yaml:"undorequestid" json:"undorequestid"`
	SGSetting             map[string]interface{} `yaml:"sgsetting" json:"sgsetting"`
	LastGreenHealthResult int                    `yaml:"last_green_health_result" json:"last_green_health_result,omitempty"`
	IsProtectDelete       bool                   `yaml:"is_protect_delete" json:"is_protect_delete"`
	DeleteProtectTime     int                    `yaml:"delete_protect_time" json:"delete_protect_time,omitempty"`
	IsSuspend             bool                   `yaml:"is_suspend" json:"is_suspend"`
}

const (
	SG_JVESSEL  = "J" //JVESSEL
	SG_INTERNAL = "I" //INTERNAL
	SG_EXTERNAL = "E" //EXTERNAL
)

func GetAllSGTypes() []string {
	return []string{SG_JVESSEL, SG_INTERNAL, SG_EXTERNAL}
}

type SpaceSpec struct {
	VpcId              string           `yaml:"vpcid" json:"vpcid"`
	SubnetId           string           `yaml:"subnetid" json:"subnetid"`
	TenantId           string           `yaml:"tenantid" json:"tenantid"`
	PreinstallWhiteIps []string         `yaml:"preinstall_whiteips" json:"preinstall_whiteips"`
	GreenHealth        *HealthCheckSpec `yaml:"green_health" json:"green_health,omitempty"`
}

type HealthCheckSpec struct {
	Cmd           string   `yaml:"cmd" json:"cmd,omitempty"`
	TargetList    []string `yaml:"target_list" json:"target_list,omitempty"`
	MaxRetryTimes int      `yaml:"max_retry_times" json:"max_retry_times,omitempty"`
}

type SpaceDef struct {
	Meta `yaml:",inline"`
	Spec SpaceSpec `yaml:"spec" json:"spec"`
}

type SpaceRuntime struct {
	Status SpaceStatus `yaml:"status" json:"status"`
}

func (t *SpaceRuntime) SetPhase(phase string) {
	t.Status.Phase = phase
}

func (t *SpaceRuntime) GetPhase() string {
	return t.Status.Phase
}

func (t *SpaceRuntime) SetRequestId(requestId string) {
	t.Status.RequestId = requestId
}

func (t *SpaceRuntime) RefreshUpdatetime() {
	t.Status.Updatetime = time.Now().Round(0)
}

type Space struct {
	SpaceDef     `yaml:",inline"`
	SpaceRuntime `yaml:",inline"`
}

func (t *Space) GetDef() interface{} {
	return &(t.SpaceDef)
}

var rollbackablePhasesOfSpace = []string{SpaceLaunching, SpaceRunning, SpaceUpdating}

func IsSpaceRollbackable(space *Space) bool {
	for _, v := range rollbackablePhasesOfSpace {
		if v == space.Status.Phase {
			return true
		}
	}
	return false
}

type ServiceStatus struct {
	Phases struct {
		Normal   int      `json:"normal"`
		Updating []string `json:"updating"`
		Abnormal []string `json:"abnormal"`
	} `json:"phases"`
	NoneFailovers []string `json:"noneFailovers"`
	Block         bool     `json:"block"`
}

type InstanceStatus struct {
	Instances map[string]string `json:"instances"`
}
