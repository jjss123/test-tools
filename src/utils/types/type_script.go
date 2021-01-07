package types

type ScriptSpec struct {
	Target  string `yaml:"target" json:"target,omitempty"`
	Cmd     string `yaml:"cmd" json:"cmd,omitempty"`
	Timeout int    `yaml:"timeout" json:"timeout,omitempty"` // second as unit
}

type ScriptDef struct {
	Meta `json:",omitempty" yaml:",inline"`
	Spec ScriptSpec `yaml:"spec" json:"spec,omitempty"`
}

type ScriptStatus struct {
	Result string `yaml:"result" json:"result"`
}

/**
 * Definition for a special resource - Script which provide abstraction of synchronous-call in jvessel
 * Execution record will stored in template system
 */
type Script struct {
	ScriptDef `json:",omitempty" yaml:",inline"`
	Status    ScriptStatus `yaml:"status" json:"status,omitempty"`
}

//useless method
func (s *Script) SetPhase(phase string)         {}
func (s *Script) GetPhase() string              { return "" }
func (s *Script) SetRequestId(requestId string) {}
func (s *Script) RefreshUpdatetime()            {}

func (s *Script) GetDef() interface{} {
	return &(s.ScriptDef)
}
