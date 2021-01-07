package types

import "time"

const (
	UpdateIniting  = CommonIniting
	UpdateRunning  = CommonRunning
	UpdateComplete = "Completed"
	UpdateError    = "Error"
	UpdateToDelete = CommonToDelete

	UpdateKindNoChange   = "no_change"
	UpdateKindFlavor     = "replicaset_flavor"
	UpdateKindRateFlavor = "replicaset_rate_flavor"
	UpdateKindReplica    = "replicaset_replica"
	UpdateKindImage      = "replicaset_image"
	UpdateKindFloatIP    = "replicaset_fi"
	UpdateKindVolumn     = "blockstore_flavor"
	UpdateKindLocalMount = "replicaset_localmount"
	UpdateKindExposes    = "replicaset_exposes"
	UpdateKindNLBFloatIP = "nlb_fi"
	UpdateKindMulti      = "_muti"
	UpdateKindForce      = "force"
	UpdateKindRSNlb      = "replicaset_nlb"
	UpdateKindNlbPort    = "nlb_port"
)

type Update struct {
	UpdateDef     `yaml:",inline"`
	UpdateRuntime `yaml:",inline"`
}

type UpdateDef struct {
	Meta `yaml:",inline"`
	Spec UpdateSpec `yaml:"spec" json:"spec"`
}

type UpdateRuntime struct {
	Status UpdateStatus `yaml:"status" json:"status"`
}

type UpdateSpec struct {
	UpdateKind  string                 `yaml:"update_kind" json:"update_kind"`
	Template    map[string]interface{} `yaml:"template" json:"template"`
	IgnorePhase bool                   `yaml:"template" json:"ignore_phase,omitempty"`
}

type UpdateStatus struct {
	BasicStatus `yaml:",inline"`
}

func (t *Update) GetDef() interface{} {
	return &(t.UpdateDef)
}

func (t *Update) SetPhase(phase string) {
	t.Status.Phase = phase
}

func (t *Update) GetPhase() string {
	return t.Status.Phase
}

func (t *Update) SetRequestId(requestId string) {
	t.Status.RequestId = requestId
}

func (t *Update) RefreshUpdatetime() {
	t.Status.Updatetime = time.Now().Round(0)
}

type UpdateSerializer struct {
	Meta `yaml:",inline"`
	Spec struct {
		UpdateKind string `yaml:"update_kind" json:"update_kind"`
		Template   struct {
			Meta `yaml:",inline"`
		} `yaml:"template" json:"template"`
	} `yaml:"spec" json:"spec"`
}

//
//type UpdateTemplate struct {
//	Spec struct {
//		Template string `yaml:"template"`
//	} `yaml:"spec"`
//}
