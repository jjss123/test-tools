package types

import "sort"

type DynamicResourceAZ struct {
	CPU               int `json:"cpu"` // total
	Memory            int `json:"memory"`
	LocalDataDisk     int `json:"local_data_disk"`
	CloudDataDisk     int `json:"cloud_data_disk"`
	CPUOfJD           int `json:"cpu_of_jd"` // jd means jd pool
	MemoryOfJD        int `json:"memory_of_jd"`
	LocalDataDiskOfJD int `json:"local_data_disk_of_jd"`
	CloudDataDiskOfJD int `json:"cloud_data_disk_of_jd"`
}

type DynamicResource struct {
	Service  int `json:"service"`
	UserPin  int `json:"user_pin"`
	VPC      int `json:"vpc"`
	Instance int `json:"instance"`

	VM              int `json:"vm"`
	Container       int `json:"container"`
	NativeContainer int `json:"native_container"`
	FloatingIP      int `json:"floating_ip"`
	Nlb             int `json:"nlb"`

	DynamicResourceAZ
	Azs map[string]*DynamicResourceAZ `json:"azs"`
}

func (p *DynamicResource) Init() {
	p.Azs = map[string]*DynamicResourceAZ{}
}

type ResourceStatisticsUsrePin struct {
	Pin string `json:"pin"`
	*DynamicResource
}

type ResourceFlavor struct {
	Flavor string `json:"flavor"`
	Total  int    `json:"total"`
}

type ResourceStatisticsFlavor struct {
	Flavors []*ResourceFlavor `json:"flavors"`
}

type ResourceUserPin struct {
	UserPin string `json:"userPin"`
	Total   int    `json:"total"`
}

type ResourceUserPinArr []ResourceUserPin
type ResourceStatisticsTopUserPin struct {
	CPU    ResourceUserPinArr `json:"cpu"`
	Memory ResourceUserPinArr `json:"memory"`
	Disk   ResourceUserPinArr `json:"disk"`
}

type DynamicResourceRegion map[string]*DynamicResourceAZ

func InitDynamicResourceRegion() DynamicResourceRegion {
	return map[string]*DynamicResourceAZ{}
}
func (p ResourceUserPinArr) Len() int {
	return len(p)
}

func (p ResourceUserPinArr) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ResourceUserPinArr) Less(i, j int) bool {
	return p[i].Total < p[j].Total
}

func (p *ResourceStatisticsTopUserPin) SortByMap(cpu, memory, disk map[string]int, topNum int) {
	topcpu := make(ResourceUserPinArr, len(cpu))
	i := 0
	for k, v := range cpu {
		topcpu[i] = ResourceUserPin{k, v}
		i += 1
	}

	topmemeroy := make(ResourceUserPinArr, len(memory))
	i = 0
	for k, v := range memory {
		topmemeroy[i] = ResourceUserPin{k, v}
		i += 1
	}

	topdisk := make(ResourceUserPinArr, len(disk))
	i = 0
	for k, v := range disk {
		topdisk[i] = ResourceUserPin{k, v}
		i += 1
	}

	p.Sort(topcpu, topmemeroy, topdisk, topNum)
}

func (p *ResourceStatisticsTopUserPin) SortBySlice(topUserPins []*ResourceStatisticsTopUserPin, topNum int) {
	topcpu := ResourceUserPinArr{}
	topmemeroy := ResourceUserPinArr{}
	topdisk := ResourceUserPinArr{}

	for _, val := range topUserPins {
		topcpu = append(topcpu, val.CPU...)
		topmemeroy = append(topmemeroy, val.Memory...)
		topdisk = append(topdisk, val.Disk...)
	}

	p.Sort(topcpu, topmemeroy, topdisk, topNum)
}

func (p *ResourceStatisticsTopUserPin) Sort(topcpu, topmemeroy, topdisk ResourceUserPinArr, topNum int) {
	sort.Sort(sort.Reverse(topcpu))
	sort.Sort(sort.Reverse(topmemeroy))
	sort.Sort(sort.Reverse(topdisk))

	for i := 0; i < topNum && i < len(topcpu); i++ {
		p.CPU = append(p.CPU, topcpu[i])
	}
	for i := 0; i < topNum && i < len(topmemeroy); i++ {
		p.Memory = append(p.Memory, topmemeroy[i])
	}
	for i := 0; i < topNum && i < len(topdisk); i++ {
		p.Disk = append(p.Disk, topdisk[i])
	}
}

const (
	ResourceInstanceStatus_Normal            = "Normal"
	ResourceInstanceStatus_Updating          = "Updating"
	ResourceInstanceStatus_UpdatingCompleted = "UpdatingCompleted"
	ResourceInstanceStatus_Abnormal          = "Abnormal"
	ResourceInstanceOpStatus_Blocked         = "Blocked"
	ResourceInstanceOpStatus_Suspend         = "Suspend"
	ResourceInstanceOpStatus_UnGreenHealth   = "UnGreenHealth"
)

type ResourceDistributionItem struct {
	Name  string `json:"name"`
	Total int    `json:"total"`
}
