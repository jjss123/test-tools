package types

import "time"

// BlockStore Phase
const (
	SnapshotPending  = CommonPending
	SnapshotRunning  = CommonRunning
	SnapshotToDelete = CommonToDelete
)

type SnapshotSpec struct {
	VolumeID string `json:"volume_id"`
	TenantId    string `json:"tenantid"`
	Size        int    `yaml:"size" json:"size"`
}

type SnapshotDef struct {
	Meta `json:",omitempty" yaml:",inline"`
	Spec SnapshotSpec `yaml:"spec" json:"spec,omitempty"`
}

type SnapshotStatus struct {
	BasicStatus `yaml:",inline"`
	ID          string `yaml:"id" json:"id"`
}

type SnapShot struct {
	SnapshotDef `json:",omitempty" yaml:",inline"`
	Status        SnapshotStatus `yaml:"status" json:"status,omitempty"`
}

func (s *SnapShot) SetPhase(phase string) {
	s.Status.Phase = phase
}

func (s *SnapShot) GetPhase() string {
	return s.Status.Phase
}

func (s *SnapShot) SetRequestId(requestId string) {
	s.Status.RequestId = requestId
}

func (s *SnapShot) RefreshUpdatetime() {
	s.Status.Updatetime = time.Now().Round(0)
}

func (s *SnapShot) GetDef() interface{} {
	return &(s.SnapshotDef)
}