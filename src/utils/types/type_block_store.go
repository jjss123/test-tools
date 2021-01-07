package types

import "time"

// BlockStore Phase
const (
	BlockStorePending  = CommonPending
	BlockStoreRunning  = CommonRunning
	BlockStoreToDelete = CommonToDelete
)

const (
	YamlPathVolumeSize              = "size"
	YamlPathVolumeKind              = "type"
	YamlPathVolumeType              = "volumetype"
	YamlPathVolumePath              = "path"
	BlockstoreCreateSnapshotContext = "snapshot_name"

	CtxBlockStoreRecreateTimes = "RECREATE_TIMES"
)

type BlockStoreSpec struct {
	Type         string `yaml:"type" json:"type"`
	Size         int    `yaml:"size" json:"size"`
	TenantId     string `yaml:"tenantid" json:"tenantid"`
	SnapshotName string `yaml:"snapshot_name" json:"snapshot_name,omitempty"`
	ServiceCode  string `yaml:"service_code" json:"service_code"`
}

type BlockStoreDef struct {
	Meta `json:",omitempty" yaml:",inline"`
	Spec BlockStoreSpec `yaml:"spec" json:"spec,omitempty"`
}

type BlockStoreStatus struct {
	BasicStatus `yaml:",inline"`
	ID          string `yaml:"id" json:"id"`
	AZ          string `yaml:"az" json:"az"`
	Format      bool   `json:"format,omitempty"`
}

type BlockStore struct {
	BlockStoreDef `json:",omitempty" yaml:",inline"`
	Status        BlockStoreStatus `yaml:"status" json:"status,omitempty"`
}

func (s *BlockStore) SetPhase(phase string) {
	s.Status.Phase = phase
}

func (s *BlockStore) GetPhase() string {
	return s.Status.Phase
}

func (s *BlockStore) SetRequestId(requestId string) {
	s.Status.RequestId = requestId
}

func (s *BlockStore) RefreshUpdatetime() {
	s.Status.Updatetime = time.Now().Round(0)
}

func (s *BlockStore) GetDef() interface{} {
	return &(s.BlockStoreDef)
}
