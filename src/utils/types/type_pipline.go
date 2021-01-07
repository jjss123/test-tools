package types

import "time"

// Pipeline
const (
	// Initing
	PipelineIniting = CommonIniting
	// Running
	PipelineRunning = CommonRunning
	// Completed
	PipelineCompleted = "Completed"
	// PipelineToDelete
	PipelineToDelete = CommonToDelete
	// Error
	PipelineError = "Error"
)

type Pipeline struct {
	PipelineDef     `yaml:",inline"`
	PipelineRuntime `yaml:",inline"`
}

type PipelineDef struct {
	Meta `yaml:",inline"`
	Spec PipelineSpec `yaml:"spec" json:"spec"`
}

type PipelineSpec struct {
	// yaml name slice
	Steps       []string              `yaml:"steps" json:"steps"`
	ReplicaSets map[string]ReplicaSet `yaml:"replicasets" json:"replicasets"`
	Jobs        map[string]Job        `yaml:"jobs" json:"jobs"`
	Scripts     map[string]Script     `yaml:"scripts" json:"scripts"`
	Updates     map[string]Update     `yaml:"updates" json:"updates"`
}

type PipelineRuntime struct {
	Status PipelineStatus `yaml:"status" json:"status"`
}

type PipelineStatus struct {
	BasicStatus `yaml:",inline"`
	// all task execute history
	Steps       []PipelineStep `yaml:"steps" json:"steps"`
	CurrentStep int            `yaml:"current_step" json:"current_step"`
}

const (
	PipelineStepInitingPhase   = "init"
	PipelineStepStartedPhase   = "started"
	PipelineStepCompletedPhase = "completed"
)

type PipelineStep struct {
	Phase        string `yaml:"phase" json:"phase"`
	YamlFileName string `yaml:"yaml_file_name" json:"yaml_file_name"`
	Result       string `yaml:"result" json:"result"`
	ExitCode     int    `yaml:"exit_code" json:"exit_code"`
}

func (t *Pipeline) GetDef() interface{} {
	return &(t.PipelineDef)
}

func (t *Pipeline) SetPhase(phase string) {
	t.Status.Phase = phase
}

func (t *Pipeline) GetPhase() string {
	return t.Status.Phase
}

func (t *Pipeline) SetRequestId(requestId string) {
	t.Status.RequestId = requestId
}

func (t *Pipeline) RefreshUpdatetime() {
	t.Status.Updatetime = time.Now().Round(0)
}
