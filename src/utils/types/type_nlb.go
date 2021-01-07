package types

import "time"

//Nlb Phase
const (
	//API Server set Pending
	NlbIniting = CommonIniting
	//Controller set Pending
	NlbPending = CommonPending
	//Worker Set Running
	NlbRunning  = CommonRunning
	NlbToDelete = CommonToDelete
)

const (
	NlbDiffPairKey_Id              = "id"
	NlbDiffPairKey_ContainerName   = "container_name"
	NlbDiffPairKey_ContainerIp     = "container_ip"
	NlbDiffPairKey_ContainerIpv6   = "container_ipv6"
	NlbDiffPairKey_ContainerPortId = "container_port_id"
	NlbDiffPairKey_TargetId        = "target_id"
	NlbDiffPairKey_TargetIdv6      = "target_idv6"
)

const (
	YamlPathNlbFloatIP          = "spec.floatip"
	YamlPathNlbExposes          = "spec.exposes"
	YamlPathNlbExposeRangePorts = "spec.expose_range_ports"
	YamlPathNlbPort             = "spec.port"
)

type NlbDiffPairValue = map[string]interface{}

type NlbStatus struct {
	BasicStatus  `yaml:",inline"`
	Ip           string `yaml:"ip" json:"ip"`
	Ipv6         string `yaml:"ipv6" json:"ipv6,omitempty"`
	ExposeDomain string `yaml:"expose_domain" json:"expose_domain"`
	PortID       string `yaml:"portid" json:"portid"`
	NlbID        string `yaml:"nlb_id" json:"nlb_id"`
	NlbPoolID    string `yaml:"nlbpool_id" json:"nlbpool_id"`
	//RuleID -> key: port; value : ruleID
	NlbRuleID map[int]string `yaml:"nlbrule_id" json:"nlbrule_id"`
	//TargetID -> key : containerName; value : targetID
	// todo delete this one
	NlbContext DiffPair `yaml:"nlb_context" json:"nlb_context"`
	SGContext  DiffPair `yaml:"sgcontext" json:"sgcontext"`
}

type NlbSpec struct {
	Port        []int             `yaml:"port" json:"port"`
	VpcId       string            `yaml:"vpcid" json:"vpcid"`
	SubnetId    string            `yaml:"subnetid" json:"subnetid"`
	TenantId    string            `yaml:"tenantid" json:"tenantid"`
	Exposes     []ExposeSpec      `yaml:"exposes" json:"exposes"`
	ExposePorts []ExposePortsSpec `yaml:"expose_range_ports" json:"expose_range_ports,omitempty"`
	FloatIP     FloatIPExpose     `yaml:"floatip" json:"floatip"`
}

type NlbLink struct {
	Name string `yaml:"name" json:"name"`
}

type NlbDef struct {
	Meta `yaml:",inline"`
	Spec NlbSpec `yaml:"spec" json:"spec"`
}

type NlbRuntime struct {
	Status NlbStatus `yaml:"status" json:"status"`
}

func (t *NlbRuntime) SetPhase(phase string) {
	t.Status.Phase = phase
}

func (t *NlbRuntime) GetPhase() string {
	return t.Status.Phase
}

func (t *NlbRuntime) SetRequestId(requestId string) {
	t.Status.RequestId = requestId
}

func (t *NlbRuntime) RefreshUpdatetime() {
	t.Status.Updatetime = time.Now().Round(0)
}

type Nlb struct {
	NlbDef     `yaml:",inline"`
	NlbRuntime `yaml:",inline"`
}

func (t *Nlb) GetDef() interface{} {
	return &(t.NlbDef)
}
