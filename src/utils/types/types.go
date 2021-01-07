package types

import "time"

//All top resources are separated into two parts:
// 1.xxDef, definition fields of resource;
// 2.xxRuntime, runtime fields of resource;

// supported resources Kind
const (
	ContainerResource       = "Container"
	ComputeResource         = "Compute"
	VMResource              = "VM"
	NCResource              = "NC"
	ReplicaSetResource      = "ReplicaSet"
	JobResource             = "Job"
	NlbResource             = "Nlb"
	SpaceResource           = "Space"
	FloatIPResource         = "FloatIP"
	SecurityGroupResource   = "SecurityGroup"
	ScriptResource          = "Script"
	BlockStoreResource      = "BlockStore"
	SnapshotResource        = "Snapshot"
	ScheduleResource        = "Schedule"
	ImportServiceResource   = "ImportService"
	UnderlayEntryResource   = "UnderlayEntry"
	PipelineResource        = "Pipeline"
	UpdateResource          = "Update"
	ContainerAloneResource  = "ContainerAlone"
	BlockStoreAloneResource = "BlockStoreAlone"
)

const (
	VMComputeType        = "vm"
	NCComputeType        = "nc"
	ContainerComputeType = "container"
)

const (
	IpVersionV4   = "v4"
	IpVersionDual = "v4&v6"
)

// all 'type' is a 'Resource'
type Resource interface {
	GetMeta() *Meta
	GetDef() interface{}
	SetPhase(phase string)
	GetPhase() string
	SetRequestId(requestId string)
	RefreshUpdatetime()
}

// common resource phase
const (
	CommonIniting  = "Initing"
	CommonPending  = "Pending"
	CommonRunning  = "Running"
	CommonToDelete = "ToDelete"
)

// common yaml path
const (
	YamlPathKind      = "kind"
	YamlPathLabels    = "metadata.labels"
	YamlPathName      = "metadata.name"
	YamlPathNamespace = "metadata.namespace"
)

/* some common struct */

// meta of resource
type Meta struct {
	Kind     string `yaml:"kind" json:"kind"`
	Metadata struct {
		Name      string            `yaml:"name" json:"name"`
		Namespace string            `yaml:"namespace" json:"namespace"`
		Labels    map[string]string `yaml:"labels" json:"labels"`
	} `yaml:"metadata" json:"metadata"`
}

func (t *Meta) GetMeta() *Meta {
	return t
}

// status of resources
type BasicStatus struct {
	Phase     string `yaml:"phase" json:"phase"`
	RequestId string `yaml:"requestid" json:"requestid"`
	// extension field for controllers or worker to store self-defined key in resources
	// to avoid conflicts, better provide special prefix for key
	Context    map[string]interface{} `yaml:"context" json:"context"`
	Createtime time.Time              `yaml:"createtime" json:"createtime"`
	Updatetime time.Time              `yaml:"updatetime" json:"updatetime"`
	Deletetime time.Time              `yaml:"deletetime" json:"deletetime"`
}

// nlb/float_ip/container/vm
type ExposeSpec struct {
	Tag  string `yaml:"tag" json:"tag"`
	Port int    `yaml:"port" json:"port"`
}

//[Start,End]
type ExposePortsSpec struct {
	Start int `yaml:"start" json:"start"`
	End   int `yaml:"end" json:"end"`
}

// nlb/container/vm
type FloatIPExpose struct {
	Type        string            `yaml:"type" json:"type"`
	BandWidth   int               `yaml:"bandwidth" json:"bandwidth,omitempty"`
	Exposes     []ExposeSpec      `yaml:"exposes" json:"exposes"`
	ExposePorts []ExposePortsSpec `yaml:"expose_range_ports" json:"expose_range_ports,omitempty"`
}

type ExposePortsSpecSlice []ExposePortsSpec

func (c ExposePortsSpecSlice) Len() int {
	return len(c)
}

func (c ExposePortsSpecSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c ExposePortsSpecSlice) Less(i, j int) bool {
	if c[i].Start < c[j].Start {
		return true
	} else if c[i].Start > c[j].Start {
		return false
	}

	if c[i].End < c[j].End {
		return true
	}
	return false
}
