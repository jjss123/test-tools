package types

import (
	"time"
)

const (
	// use -1 to distinguish not-set-int-field from zero-int-field during unmarshal process
	// type data struct{ V int }
	// a := data{};    json.Unmarshal("{}", &a) // a.V => 0
	// a = data{V:-1}; json.Unmarshal("{}", &a) // a.V => -1
	// a = data{V:0};  json.Unmarshal("{}", &a) // a.V => 0
	UnsignedIntNotSetFlag = -1
)

// factories produce resources with default values
func initStatusCommon(status *BasicStatus) {
	status.Context = make(map[string]interface{})
	status.Createtime = time.Now().Round(0)
	status.Updatetime = time.Now().Round(0)
}

func NewContainer() *Container {
	o := &Container{}
	o.Kind = ContainerResource
	o.Status.Phase = ContainerIniting
	initStatusCommon(&(o.Status.BasicStatus))
	o.Metadata.Labels = make(map[string]string)
	o.Status.Volumes = make(map[string]interface{})
	o.Spec.Volumes = make(map[string]Volume)
	return o
}

func NewVM() *VM {
	o := &VM{}
	o.Kind = VMResource
	o.Status.Phase = VMIniting
	initStatusCommon(&(o.Status.BasicStatus))
	o.Metadata.Labels = make(map[string]string)
	o.Status.Volumes = make(map[string]interface{})
	return o
}

func NewNC() *NC {
	o := &NC{}
	o.Kind = NCResource
	o.Status.Phase = NCIniting
	initStatusCommon(&(o.Status.BasicStatus))
	o.Metadata.Labels = make(map[string]string)
	o.Status.Volumes = make(map[string]interface{})
	return o
}

func NewReplicaSet() *ReplicaSet {
	o := &ReplicaSet{}
	o.Kind = ReplicaSetResource
	o.Status.Phase = ReplicaSetPending
	initStatusCommon(&(o.Status.BasicStatus))
	o.Metadata.Labels = make(map[string]string)
	o.Spec.Replica = UnsignedIntNotSetFlag
	o.Spec.RollingUpdate.Strategy = RollingUpdateStepStrategy
	o.Spec.RollingUpdate.MaxUnavailable = 1
	return o
}

func NewJob() *Job {
	o := &Job{}
	o.Kind = JobResource
	o.Status.Phase = JobPending
	initStatusCommon(&(o.Status.BasicStatus))
	o.Metadata.Labels = make(map[string]string)
	o.Status.ExitCode = JobDefaultResultValue
	return o
}

func NewSpace() *Space {
	o := &Space{}
	o.Kind = SpaceResource
	o.Status.Phase = SpacePending
	o.Status.SGSetting = make(map[string]interface{})
	initStatusCommon(&(o.Status.BasicStatus))
	o.Metadata.Labels = make(map[string]string)
	return o
}

func NewNlb() *Nlb {
	o := &Nlb{}
	o.Kind = NlbResource
	o.Status.Phase = NlbIniting
	initStatusCommon(&(o.Status.BasicStatus))
	o.Metadata.Labels = make(map[string]string)
	o.Status.NlbContext = NewDiffPair()
	o.Status.NlbRuleID = make(map[int]string)
	return o
}

func NewFloatIP() *FloatIP {
	o := &FloatIP{}
	o.Kind = FloatIPResource
	o.Status.Phase = FloatIPIniting
	initStatusCommon(&(o.Status.BasicStatus))
	o.Metadata.Labels = make(map[string]string)
	o.Status.BindPort = NewDiffPair()
	return o
}

func NewSecurityGroup() *SecurityGroup {
	o := &SecurityGroup{}
	o.Kind = SecurityGroupResource
	o.Status.Phase = SecurityGroupIniting
	initStatusCommon(&(o.Status.BasicStatus))
	o.Metadata.Labels = make(map[string]string)
	return o
}

func NewScript() *Script {
	o := &Script{}
	o.Kind = ScriptResource
	o.Metadata.Labels = make(map[string]string)
	return o
}

func NewBlockStore() *BlockStore {
	o := &BlockStore{}
	o.Kind = BlockStoreResource
	o.Status.Phase = BlockStorePending
	initStatusCommon(&(o.Status.BasicStatus))
	o.Metadata.Labels = make(map[string]string)
	return o
}

func NewSnapshot() *SnapShot {
	s := &SnapShot{}
	s.Kind = SnapshotResource
	s.Status.Phase = SnapshotPending
	initStatusCommon(&(s.Status.BasicStatus))
	s.Metadata.Labels = make(map[string]string)
	return s
}

func NewSchedule() *Schedule {
	o := &Schedule{}
	o.Kind = ScheduleResource
	o.Status.Createtime = time.Now().Round(0)
	o.Status.Updatetime = time.Now().Round(0)
	o.Metadata.Labels = make(map[string]string)
	return o
}

func NewImportService() *ImportService {
	o := &ImportService{}
	o.Kind = ImportServiceResource
	o.Status.Createtime = time.Now().Round(0)
	o.Status.Updatetime = time.Now().Round(0)
	o.Metadata.Labels = make(map[string]string)
	return o
}

func NewUnderlayEntry() *UnderlayEntry {
	o := &UnderlayEntry{}
	o.Kind = UnderlayEntryResource
	o.Status.Phase = UnderlayEntryPending
	initStatusCommon(&(o.Status.BasicStatus))
	o.Metadata.Labels = make(map[string]string)
	return o
}

func NewPipeline() *Pipeline {
	o := &Pipeline{}
	o.Kind = PipelineResource
	o.Spec.ReplicaSets = make(map[string]ReplicaSet)
	o.Spec.Jobs = make(map[string]Job)
	o.Spec.Scripts = make(map[string]Script)
	o.Spec.Updates = make(map[string]Update)
	o.Spec.Steps = make([]string, 0)
	o.Status.Phase = PipelineIniting
	o.Status.Steps = make([]PipelineStep, 0)
	o.Status.CurrentStep = 0
	initStatusCommon(&(o.Status.BasicStatus))
	o.Metadata.Labels = make(map[string]string)
	return o
}

func NewUpdate() *Update {
	o := &Update{}
	o.Kind = UpdateResource
	o.Status.Phase = UpdateIniting
	o.Spec.Template = make(map[string]interface{})
	initStatusCommon(&o.Status.BasicStatus)
	o.Metadata.Labels = make(map[string]string)
	return o
}

func NewContainerAlone() *Container {
	o := &Container{}
	o.Kind = ContainerAloneResource
	o.Status.Phase = ContainerIniting
	initStatusCommon(&(o.Status.BasicStatus))
	o.Metadata.Labels = make(map[string]string)
	o.Status.Volumes = make(map[string]interface{})
	o.Spec.Volumes = make(map[string]Volume)
	return o
}

func NewBlockStoreAlone() *BlockStore {
	o := &BlockStore{}
	o.Kind = BlockStoreAloneResource
	o.Status.Phase = BlockStorePending
	initStatusCommon(&(o.Status.BasicStatus))
	o.Metadata.Labels = make(map[string]string)
	return o
}
