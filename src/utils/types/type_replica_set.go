package types

import "time"

// ReplicaSet Phase
const (
	// Pending
	ReplicaSetPending = CommonPending
	// Launching
	ReplicaSetLaunching = "Launching"
	// Running
	ReplicaSetRunning = CommonRunning
	// Updating
	ReplicaSetUpdating = "Updating"
	// ToDelete
	ReplicaSetToDelete = CommonToDelete
)

const (
	YamlPathImage            = "spec.template.image"
	YamlPathImageID          = "spec.template.image_id"
	YamlPathReplica          = "spec.replica"
	YamlPathFlavor           = "spec.template.flavor"
	YamlPathRateFlavor       = "spec.template.rate_flavor"
	YamlPathRSFloatIP        = "spec.template.floatip"
	YamlPathVolume           = "spec.template.volumes"
	YamlPathLocalMount       = "spec.template.localmount"
	YamlPathExposes          = "spec.template.exposes"
	YamlPathExposeRangePorts = "spec.template.expose_range_ports"
	YamlPathRSNlb            = "spec.template.nlb"
)

type ReplicaSetStatus struct {
	BasicStatus `yaml:",inline"`
	AZ          string `yaml:"az" json:"az"`
}

type ReplicaSetSpec struct {
	Replica     int         `yaml:"replica" json:"replica"`
	ComputeType string      `yaml:"compute_type" json:"compute_type,omitempty"`
	Template    ComputeSpec `yaml:"template" json:"template"`
	// RollingUpdate
	RollingUpdate RollingUpdate `yaml:"rollingUpdate" json:"rollingUpdate"`
}

type RollingUpdate struct {
	Strategy       string `yaml:"strategy" json:"strategy"`
	MaxUnavailable int    `yaml:"maxUnavailable" json:"maxUnavailable"`
}

type ReplicaSetDef struct {
	Meta `yaml:",inline"`
	Spec ReplicaSetSpec `yaml:"spec" json:"spec"`
}

type ReplicaSetRuntime struct {
	Status ReplicaSetStatus `yaml:"status" json:"status"`
	//definition of this container before last modification
	LastDef *ReplicaSet `yaml:"lastdef" json:"lastdef"`
}

func (t *ReplicaSetRuntime) SetPhase(phase string) {
	t.Status.Phase = phase
}

func (t *ReplicaSetRuntime) GetPhase() string {
	return t.Status.Phase
}

func (t *ReplicaSetRuntime) SetRequestId(requestId string) {
	t.Status.RequestId = requestId
}

func (t *ReplicaSetRuntime) RefreshUpdatetime() {
	t.Status.Updatetime = time.Now().Round(0)
}

type ReplicaSet struct {
	ReplicaSetDef     `yaml:",inline"`
	ReplicaSetRuntime `yaml:",inline"`
}

func (t *ReplicaSet) GetDef() interface{} {
	return &(t.ReplicaSetDef)
}
