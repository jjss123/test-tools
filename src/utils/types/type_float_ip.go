package types

import "time"

// FloatIP Phase
const (
	FloatIPIniting  = CommonIniting
	FloatIPPending  = CommonPending
	FloatIPUpdating = "Updating"
	FloatIPScale    = "Scale"
	FloatIPRunning  = CommonRunning
	FloatIPToDelete = CommonToDelete
)

// floatip type
const (
	InternalFloatIP = "internal" //use for vpc and management network
	PublicFloatIP   = "public"   //use for internet
	PrivateFloatIP  = "private"  // use for vpc and management network
)

// default bandwidth
const (
	DefaultPublicBandWidth   = 1
	DefaultPrivateBandWidth  = 100
	DefaultInternalBandWidth = 1000
)

const (
	FloatIPCanVisibility    = 0
	FloatIPCanNotVisibility = 1
)

const (
	YamlFloatIPPathType        = "type"
	YamlFloatIPPathBandWidth   = "bandwidth"
	YamlFloatIPPathExposes     = "exposes"
	YamlFloatIPPathExposePorts = "expose_range_ports"
)

type FloatIPStatus struct {
	BasicStatus  `yaml:",inline"`
	IP           string   `yaml:"ip" json:"ip"`
	IPv6         string   `yaml:"ipv6" json:"ipv6,omitempty"`
	ExposeDomain string   `yaml:"expose_domain" json:"expose_domain"`
	ID           string   `yaml:"id" json:"id"`
	PortID       string   `yaml:"portid" json:"portid"`
	BindPort     DiffPair `yaml:"bindport" json:"bindport"`
	SGContext    DiffPair `yaml:"sgcontext" json:"sgcontext"`
}

//t.Status.Updatetime = time.Now()
type FloatIPRuntime struct {
	Status FloatIPStatus `yaml:"status" json:"status"`
}

func (f *FloatIP) SetPhase(phase string) {
	f.Status.Phase = phase
}

func (f *FloatIP) GetPhase() string {
	return f.Status.Phase
}

func (f *FloatIP) SetRequestId(requestId string) {
	f.Status.RequestId = requestId
}

func (f *FloatIP) RefreshUpdatetime() {
	f.Status.Updatetime = time.Now().Round(0)
}

type FloatIPSpec struct {
	Type        string            `yaml:"type" json:"type,omitempty"`
	BandWidth   int               `yaml:"bandwidth" json:"bandwidth,omitempty"`
	TenantId    string            `yaml:"tenantid" json:"tenantid"`
	Exposes     []ExposeSpec      `yaml:"exposes" json:"exposes"`
	ExposePorts []ExposePortsSpec `yaml:"expose_range_ports" json:"expose_range_ports,omitempty"`
}

type FloatIPDef struct {
	Meta `yaml:",inline"`
	Spec FloatIPSpec `yaml:"spec" json:"spec"`
}

type FloatIP struct {
	FloatIPDef     `yaml:",inline"`
	FloatIPRuntime `yaml:",inline"`
}

func (f *FloatIP) GetDef() interface{} {
	return &(f.FloatIPDef)
}
