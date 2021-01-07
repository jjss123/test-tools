package types

import "time"

// SecurityGroup Phase
const (
	// Initing
	SecurityGroupIniting = CommonIniting
	// Pending
	SecurityGroupPending = CommonPending
	// Running
	SecurityGroupRunning = CommonRunning
	// ToDelete
	SecurityGroupToDelete = CommonToDelete
)

const (
	ExternalSGDefaultIP = "0.0.0.0/0"
)

type SecurityGroupRuleIPData struct {
	Expect map[string]bool
	Actual map[string]bool
}

type SecurityGroupStatus struct {
	BasicStatus     `yaml:",inline"`
	SecurityGroupID string            `yaml:"securitygroupid" json:"securitygroupid"`
	RuleIPs         DiffPair          `yaml:"ruleips" json:"ruleips"`
	IPs             []string          `yaml:"ips" json:"ips,omitempty"`
	IPv6s           []string          `yaml:"ipv6s" json:"ipv6s,omitempty"`
	Ports           []int             `yaml:"ports" json:"ports,omitempty"`
	PortsRange      []ExposePortsSpec `yaml:"ports_range" json:"ports_range,omitempty"`
	Rules           *DiffPair         `yaml:"rules" json:"rules,omitempty"`
}

type SecurityGroupRuntime struct {
	Status SecurityGroupStatus `yaml:"status" json:"status"`
}

type SecurityGroupSpec struct {
	// SG_JVESSEL/SG_INTERNAL/SG_EXTERNAL
	Type       string            `yaml:"type" json:"type"`
	VpcId      string            `yaml:"vpcid" json:"vpcid"`
	TenantId   string            `yaml:"tenantid" json:"tenantid"`
	Ports      []int             `yaml:"ports" json:"ports,omitempty"`
	PortsRange []ExposePortsSpec `yaml:"expose_range_ports" json:"expose_range_ports,omitempty"`
}

type SecurityGroupDef struct {
	Meta `yaml:",inline"`
	Spec SecurityGroupSpec `yaml:"spec" json:"spec"`
}

type SecurityGroup struct {
	SecurityGroupDef     `yaml:",inline"`
	SecurityGroupRuntime `yaml:",inline"`
}

func (s *SecurityGroup) SetPhase(phase string) {
	s.Status.Phase = phase
}

func (s *SecurityGroup) GetPhase() string {
	return s.Status.Phase
}

func (s *SecurityGroup) SetRequestId(requestId string) {
	s.Status.RequestId = requestId
}

func (s *SecurityGroup) RefreshUpdatetime() {
	s.Status.Updatetime = time.Now().Round(0)
}

func (s *SecurityGroup) GetDef() interface{} {
	return &(s.SecurityGroupDef)
}
