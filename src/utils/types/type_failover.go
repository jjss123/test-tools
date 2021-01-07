package types

import "fmt"

const (
	BlockType_All       = "block_all"        //manager.BlockAll
	BlockType_AZ        = "block_az"         //manager.BlockAz
	BlockType_Rack      = "block_rack"       //manager.BlockRack
	BlockType_HostGroup = "block_host_group" //manager.BlockHostGroup
	BlockType_Host      = "block_host"       //manager.BlockHost
	BlockType_Service   = "block_service"    //manager.BlockService

	BlockType_ServiceAZ       = "block_service_az"
	BlockType_ServiceRelease  = "block_service_release"
	BlockType_ServiceInstance = "block_service_instance"
)

type Failover struct {
	Enable    bool     `json:"enable"`
	AZs       []string `json:"azs"`
	Releases  []string `json:"releases"`
	Instances []string `json:"instances"`
	Service   string   `json:"service"`
}

func (p *Failover) Init() {
	p.Enable = false
	p.AZs = []string{}
	p.Releases = []string{}
	p.Instances = []string{}
}

type ServiceFailover map[string]*Failover

func InitServiceFailover() ServiceFailover{
	return make(ServiceFailover)
}

func (f *Failover) String() string {
	return fmt.Sprintf("%+v", *f)
}

type FailoverBody struct {
	BlockType string `json:"blockType"`
	Block     bool   `json:"block"`
	BlockItem string `json:"blockItem"`
}

type GlobalFailover struct {
	Enable     bool     `json:"enable"`
	AZs        []string `json:"azs"`
	Racks      []string `json:"racks"`
	HostGroups []string `json:"hostGroups"`
	Hosts      []string `json:"hosts"`
}

func (p *GlobalFailover) Init() {
	p.Enable = false
	p.AZs = []string{}
	p.Racks = []string{}
	p.HostGroups = []string{}
	p.Hosts = []string{}
}

const (
	BlockResult_Exist = -1
)
