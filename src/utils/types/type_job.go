package types

import "time"

// Job Phase
const (
	JobPending   = CommonPending
	JobLaunching = "Launching"
	JobRunning   = CommonRunning
	JobCompleted = "Completed"
	JobToDelete  = CommonToDelete
)

const (
	JobDefaultResultValue = -1
	JobCreatedErrorValue  = -2
)

type JobSpec struct {
	Template ComputeSpec `yaml:"template" json:"template"`
}

type JobStatus struct {
	BasicStatus `yaml:",inline"`
	ExitCode    int    `yaml:"exitcode" json:"exitcode"`
	AZ          string `yaml:"az" json:"az"`
}

type JobDef struct {
	Meta `yaml:",inline"`
	Spec JobSpec `yaml:"spec" json:"spec"`
}

type JobRuntime struct {
	Status JobStatus `yaml:"status" json:"status"`
}

func (t *JobRuntime) SetPhase(phase string) {
	t.Status.Phase = phase
}

func (t *JobRuntime) GetPhase() string {
	return t.Status.Phase
}

func (t *JobRuntime) SetRequestId(requestId string) {
	t.Status.RequestId = requestId
}

func (t *JobRuntime) RefreshUpdatetime() {
	t.Status.Updatetime = time.Now().Round(0)
}

type Job struct {
	JobDef     `yaml:",inline"`
	JobRuntime `yaml:",inline"`
}

func (t *Job) GetDef() interface{} {
	return &(t.JobDef)
}
