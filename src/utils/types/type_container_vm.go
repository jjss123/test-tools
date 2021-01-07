package types

import (
	"fmt"
	"time"
)

//Container Phase
const (
	ContainerIniting          = CommonIniting      //API Server,Controller set Pending
	ContainerPending          = CommonPending      //Controller set Pending; Worker listen Pending
	ContainerLaunching        = "Launching"        //Worker set Launching;Prober listen Launching
	ContainerUpgrade          = "Upgrade"          //Controller set Upgrade, Worker listen Upgrade
	ContainerRunning          = CommonRunning      //Prober set Running, Prober,Worker listen Running
	ContainerStop             = "Stop"             //Prober set Stop; Controller listen Stop
	ContainerRestart          = "Restart"          //Controller set; Worker listen Restart
	ContainerToDelete         = CommonToDelete     //Controller set; Worker listen ToDelete
	ContainerPreDelete        = "PreDelete"        //Controller set if need to get exit code ;Worker listenand get exit code
	ContainerScale            = "Scale"            //Controller set Scale; Worker listen Scale
	ContainerScaleUnSatisfied = "ScaleUnSatisfied" //Worker set ScaleUnSatisfied; Controller listen ScaleUnSatisfied
	ContainerFrozen           = "Frozen"           //Replicaset set Frozen; For container not cycle restart
	ContainerForceDelete      = "ForceDelete"      //Space delete then make container to ContainerForceDelete
	ContainerWaitAZRecover    = "WaitAZRecover"    //Wait AZ Recover
	ContainerDeleteAZRefuge   = "DeleteAZRefuge"   //Delete resource in azRefuge
	ContainerVolumeExtend     = "VolumeExtend"     //Controller set VolumeExtend,Worker listen VolumeExtend
	ContainerRollback         = "Rollback"         // rollback vm to an specific version of snapshot
)

//VM Phase
const (
	VMIniting          = ContainerIniting
	VMPending          = ContainerPending
	VMLaunching        = ContainerLaunching
	VMUpgrade          = ContainerUpgrade
	VMRunning          = ContainerRunning
	VMStop             = ContainerStop
	VMRestart          = ContainerRestart
	VMToDelete         = ContainerToDelete
	VMPreDelete        = ContainerPreDelete
	VMScale            = ContainerScale
	VMScaleUnSatisfied = ContainerScaleUnSatisfied
	VMFrozen           = ContainerFrozen
	VMForceDelete      = ContainerForceDelete
	VMWaitAZRecover    = ContainerWaitAZRecover
	VMVolumeExtend     = ContainerVolumeExtend
)

//native container Phase
const (
	NCIniting          = ContainerIniting
	NCPending          = ContainerPending
	NCLaunching        = ContainerLaunching
	NCUpgrade          = ContainerUpgrade
	NCRunning          = ContainerRunning
	NCStop             = ContainerStop
	NCRestart          = ContainerRestart
	NCToDelete         = ContainerToDelete
	NCPreDelete        = ContainerPreDelete
	NCScale            = ContainerScale
	NCScaleUnSatisfied = ContainerScaleUnSatisfied
	NCFrozen           = ContainerFrozen
	NCForceDelete      = ContainerForceDelete
	NCWaitAZRecover    = ContainerWaitAZRecover
	NCVolumeExtend     = ContainerVolumeExtend
)

// container CTX
const (
	CtxContainerCopy                            = "COPY"
	CtxVMRollbackToSnapshot                     = "ROLLBACK_SNAPSHOT"
	CtxContainerRestart                         = "RESTART"
	CtxVolumeName                               = "VOLUME_NAME"
	CtxContainerCreateFailedAndToFroceDeleteTag = "CREATE-AND-TO-FORCEDELETE"
	CtxContainerUpgradeFailed                   = "UpgradeFailed"
	CtxLastPhaseBeforeGoToStop                  = "last_phase_before_goto_stop"
	CtxRunningTime                              = "RUNNING_TIME"
)
const (
	SystemDiskTypeLocal = "local"
	SystemDiskTypeZBS   = "zbs"
)
const (
	SystemDiskParamSSD = "ssd"
	SystemDiskParamHDD = "hdd"
)

const (
	RateFlavorPrefix = "rate_"
)

type Container struct {
	Compute `yaml:",inline"`
}
type VM struct {
	Compute `yaml:",inline"`
}
type NC struct {
	Compute `yaml:",inline"`
}

type Compute struct {
	ComputeDef     `yaml:",inline"`
	ComputeRuntime `yaml:",inline"`
}

type ComputeDef struct {
	Meta `yaml:",inline"`
	Spec ComputeSpec `yaml:"spec" json:"spec"`
}

type ComputeRuntime struct {
	Status  ComputeStatus `yaml:"status" json:"status"`
	LastDef *Compute      `yaml:"lastdef" json:"lastdef"` //definition of this container before last modification
}

type ComputeSpec struct {
	Image       string              `yaml:"image" json:"image"`
	ImageID     string              `yaml:"image_id" json:"image_id,omitempty"`
	Flavor      string              `yaml:"flavor" json:"flavor"`
	RateFlavor  *RateFlavor         `yaml:"rate_flavor" json:"rate_flavor,omitempty"`
	LocalMount  string              `yaml:"localmount" json:"localmount"`
	Network     VPCNetwork          `yaml:"network" json:"network"`
	Nlb         NlbLink             `yaml:"nlb" json:"nlb"`
	Env         []EnvItem           `yaml:"env" json:"env"`
	ConfigMap   []ConfigMapItem     `yaml:"configmap" json:"configmap,omitempty"`
	Exposes     []ExposeSpec        `yaml:"exposes" json:"exposes"`
	ExposePorts []ExposePortsSpec   `yaml:"expose_range_ports" json:"expose_range_ports,omitempty"`
	FloatIP     FloatIPExpose       `yaml:"floatip" json:"floatip"`
	SourceType  string              `yaml:"source_type" json:"source_type"`
	Volumes     ComputeVolumeSpec   `yaml:"volumes" json:"volumes"`
	IOLimit     *IOLimit            `yaml:"iolimit" json:"iolimit,omitempty"`
	Ulimit      []Ulimit            `yaml:"ulimit" json:"ulimit,omitempty"`
	Snapshots   map[string][]string `yaml:"snapshots" json:"snapshots,omitempty"`
	Probe       Probe               `yaml:"probe" json:"probe"`
	AZ          string              `yaml:"az" json:"az"`
	DiskType    string              `yaml:"disk_type" json:"disk_type,omitempty"`
	DiskSize    int                 `yaml:"disk_size" json:"disk_size,omitempty"`
}

type ComputeStatus struct {
	BasicStatus  `yaml:",inline"`
	ResourceId   string `yaml:"resource_id" json:"resource_id"`
	Ip           string `yaml:"ip" json:"ip"`
	Ipv6         string `yaml:"ipv6" json:"ipv6,omitempty"`
	ExposeDomain string `yaml:"expose_domain" json:"expose_domain"`
	Machine      string `yaml:"machine" json:"machine,omitempty"`
	HostIp       string `yaml:"hostip" json:"hostip"`
	PortID       string `yaml:"portid" json:"portid"`
	//TODO change json tag from containerid to id
	Id        string   `yaml:"containerid" json:"containerid"`
	SGContext DiffPair `yaml:"sgcontext" json:"sgcontext"`
	//SnapshotContext    *DiffPair              `yaml:"snapshotcontext" json:"snapshotcontext,omitempty"`
	Exited             bool                   `yaml:"exited" json:"exited"`
	ExitCode           int                    `yaml:"exitcode" json:"exitcode"`
	Volumes            map[string]interface{} `yaml:"volumes" json:"volumes"`
	AntiAffinityConfig string                 `yaml:"antiAffinityConfig" json:"antiAffinityConfig"`
	AZ                 string                 `yaml:"az" json:"az"`
	SystemDiskID       string                 `yaml:"systemDiskId" json:"systemDiskId,omitempty"`
	FromSnapshotName   string                 `yaml:"fromSnapshotName" json:"fromSnapshotName,omitempty"`
}

type RateFlavor struct {
	Rate       int              `yaml:"rate" json:"rate"`
	BaseFlavor string           `yaml:"base_flavor" json:"base_flavor"` // muste rate type flavor
	Disk       *int             `yaml:"disk" json:"disk,omitempty"`     // GB
	Machines   []string         `yaml:"machines" json:"machines"`
	IOLimits   []MachineIOLimit `yaml:"iolimits" json:"iolimits,omitempty"`
}

type VPCNetwork struct {
	VpcId    string `yaml:"vpcid" json:"vpcid"`
	SubnetId string `yaml:"subnetid" json:"subnetid"`
	TenantId string `yaml:"tenantid" json:"tenantid"`
}

type EnvItem struct {
	Name  string `yaml:"name" json:"name"`
	Value string `yaml:"value" json:"value"`
}

type ConfigMapItem struct {
	Key   string `yaml:"key" json:"key"`
	Value string `yaml:"value" json:"value"`
}

// block store volume type
const (
	BlockStoreVolumeSSDType = "ssd"
	BlockStoreVolumeHDDType = "hdd"
)

// volume type
const (
	BlockStoreVolume = "blockstore"
	LocalStoreVolume = "local"
)

type ComputeVolumeSpec map[string]Volume

func (v ComputeVolumeSpec) GetAllBlockStore() map[string]Volume {
	bs := make(map[string]Volume)

	for k, v := range v {
		if v.Type == BlockStoreVolume {
			bs[k] = v
		}
	}

	return bs
}

func (v ComputeVolumeSpec) GetAllLocalStore() map[string]Volume {
	bs := make(map[string]Volume)

	for k, v := range v {
		if v.Type == LocalStoreVolume {
			bs[k] = v
		}
	}

	return bs
}

type Volume struct {
	// volume big type
	Type string `yaml:"type" json:"type"`
	// zbs: size的存储单位是G
	Size int    `yaml:"size" json:"size"`
	Path string `yaml:"path" json:"path"`
	// volume detail store type
	VolumeType string `yaml:"volumetype" json:"volumetype"`
}

type IOLimit struct {
	ReadBytes  int `yaml:"read_bytes" json:"read_bytes,omitempty"`
	WriteBytes int `yaml:"write_bytes" json:"write_bytes,omitempty"`
	ReadIOPS   int `yaml:"read_iops" json:"read_iops,omitempty"`
	WriteIOPS  int `yaml:"write_iops" json:"write_iops,omitempty"`
}

type MachineIOLimit struct {
	Machine    string `yaml:"machine" json:"machine,omitempty"`
	ReadBytes  int    `yaml:"read_bytes" json:"read_bytes,omitempty"`
	WriteBytes int    `yaml:"write_bytes" json:"write_bytes,omitempty"`
	ReadIOPS   int    `yaml:"read_iops" json:"read_iops,omitempty"`
	WriteIOPS  int    `yaml:"write_iops" json:"write_iops,omitempty"`
}

type Ulimit struct {
	Name string `yaml:"name" json:"name"`
	Soft int    `yaml:"soft" json:"soft"`
	Hard int    `yaml:"hard" json:"hard"`
}

type Probe struct {
	// Number of seconds after the container has started before probe initiated.
	InitDelay int `yaml:"initdelay" json:"initdelay"`
	// How often (in seconds) to perform the probe
	Period int `yaml:"period" json:"period"`
	// Number of seconds after which the probe times out
	Timeout int `yaml:"timeout" json:"timeout"`
}

func (t *Compute) GetDef() interface{} {
	return &t.ComputeDef
}

func (t *Compute) SetPhase(phase string) {
	t.Status.Phase = phase
}

func (t *Compute) GetPhase() string {
	return t.Status.Phase
}

func (t *Compute) SetRequestId(requestId string) {
	t.Status.RequestId = requestId
}

func (t *Compute) RefreshUpdatetime() {
	t.Status.Updatetime = time.Now().Round(0)
}

func (t *Compute) RefreshRunningtime() {
	t.Status.Context[CtxRunningTime] = time.Now().Round(0).String()
}

func GetCompute(res Resource) *Compute {
	if c, ok := res.(*Compute); ok {
		return c
	}

	kind := res.GetMeta().Kind
	switch kind {
	case ContainerResource:
		return &res.(*Container).Compute
	case VMResource:
		return &res.(*VM).Compute
	case NCResource:
		return &res.(*NC).Compute
	case ContainerAloneResource:
		return &res.(*Container).Compute
	default:
		return nil
	}
}

// note: return a copy of com
func ComputeToResource(com *Compute) Resource {
	kind := com.GetMeta().Kind
	switch kind {
	case ContainerResource:
		return &Container{Compute: *com}
	case VMResource:
		return &VM{Compute: *com}
	default:
		return nil
	}
}

func IsValidRateFlavor(rateFlavor *RateFlavor) bool {
	return rateFlavor != nil && rateFlavor.Rate > 0 && rateFlavor.BaseFlavor != "" && len(rateFlavor.Machines) > 0
}

func GetFlavor(flavor string, rateFlavor *RateFlavor) string {
	if rateFlavor != nil && rateFlavor.Rate > 0 && rateFlavor.BaseFlavor != "" && len(rateFlavor.Machines) > 0 {
		return fmt.Sprintf("%s%d_%s", RateFlavorPrefix, rateFlavor.Rate, rateFlavor.BaseFlavor)
	}
	return flavor
}

/*
// in example: rate_8_paas_1c_2g
func GetBaseFlavor(flavor string) string {
	if !strings.HasPrefix(flavor, RateFlavorPrefix) {
		return ""
	}
	items := strings.Split(flavor, "_")
	if len(items) <= 2 {
		return ""
	}
	return strings.Join(items[2:], "_")
}
*/
