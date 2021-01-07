package types

import "time"

const (
	AntiAffinity_ReplicaSetsIndex   = "ReplicaSetsIndex"
	AntiAffinity_ReplicaSetInternal = "ReplicaSetInternal"
	AntiAffinity_ReplicaSets        = "ReplicaSets"
	ScheduleStrategy_Preferred      = "Preferred"
	ScheduleStrategy_Required_Host  = "HostRequired"
	ScheduleStrategy_Required_Rack  = "RackRequired"
	ScheduleStrategy_Required       = "Required"
)

/*                Host      Rack     Group
Preferred          N         N         N       All levels are preferred
HostRequired       Y         N         N       Just Host level required, others is preferred
RackRequired       Y         Y         N       Rack and Host levels are required, group level is preferred
Required           Y         Y         Y       All levels are required
*/

type Schedule struct {
	ScheduleDef     `yaml:",inline"`
	ScheduleRuntime `yaml:",inline"`
}

type ScheduleDef struct {
	Meta `yaml:",inline"`
	Spec ScheduleSpec `yaml:"spec" json:"spec"`
}

type ScheduleSpec struct {
	AntiAffinity []AntiAffinityRule `yaml:"antiaffinity" json:"antiaffinity"`
}

type AntiAffinityRule struct {
	Type      string   `yaml:"type" json:"type"`
	Strategy  string   `yaml:"strategy" json:"strategy"`
	Resources []string `yaml:"resources" json:"resources"`
}

type ScheduleRuntime struct {
	Status ScheduleStatus `yaml:"status" json:"status"`
}

type ScheduleStatus struct {
	Createtime   time.Time                  `yaml:"createtime" json:"createtime"`
	Updatetime   time.Time                  `yaml:"updatetime" json:"updatetime"`
	Deletetime   time.Time                  `yaml:"deletetime" json:"deletetime"`
	AntiAffinity []AntiAffinityRuleInstance `yaml:"antiaffinity" json:"antiaffinity"`
}

type AntiAffinityRuleInstance struct {
	Type     string `yaml:"type" json:"type"`
	Strategy string `yaml:"strategy" json:"strategy"`
	// expect containers' names
	Expect []string `yaml:"expect" json:"expect"`
}

func (s *Schedule) SetPhase(phase string) {
}

func (s *Schedule) GetPhase() string {
	return ""
}

func (s *Schedule) SetRequestId(requestId string) {
}

func (s *Schedule) RefreshUpdatetime() {
	s.Status.Updatetime = time.Now().Round(0)
}

func (s *Schedule) GetDef() interface{} {
	return &(s.ScheduleDef)
}
