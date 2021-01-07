package types

import (
	"errors"
)

// All System Injected Resource Label Keys should have unified prefix
const InjectedLabelPrefix = "_vessel_"
const (
	ServiceNameLabel             = InjectedLabelPrefix + "service_name"
	ServiceVersionLabel          = InjectedLabelPrefix + "service_version"
	ComputeNameLabel             = InjectedLabelPrefix + "compute_name"
	ActionLabel                  = InjectedLabelPrefix + "action"
	SourceFileLabel              = InjectedLabelPrefix + "source_file"
	SpaceLabel                   = InjectedLabelPrefix + "space"
	ReplicaSetLabel              = InjectedLabelPrefix + "replicaset"
	JobLabel                     = InjectedLabelPrefix + "job"
	AzLabel                      = InjectedLabelPrefix + "az"
	RackLabel                    = InjectedLabelPrefix + "rack"
	HostLabel                    = InjectedLabelPrefix + "host"
	HostGroupIDLabel             = InjectedLabelPrefix + "host_group_id"
	NlbLabel                     = InjectedLabelPrefix + "nlb"
	NlbUnbindLabel               = InjectedLabelPrefix + "unbind_nlb"
	FloatIPLabel                 = InjectedLabelPrefix + "floatip"
	OrphanContainerLabel         = InjectedLabelPrefix + "orphan_container"
	OrphanContainerLabelValue    = "yes"
	ContainerLabel               = InjectedLabelPrefix + "container"
	VMLabel                      = InjectedLabelPrefix + "vm"
	NCLabel                      = InjectedLabelPrefix + "nc"
	BlockStoreLabel              = InjectedLabelPrefix + "bs"
	ContainerFlavorCPULabel      = InjectedLabelPrefix + "flavor_cpu"
	ContainerFlavorMemoryLabel   = InjectedLabelPrefix + "flavor_memory"
	ContainerFlavorDiskSizeLabel = InjectedLabelPrefix + "flavor_disk_size"
	ImportServiceLabel           = InjectedLabelPrefix + "import_service"
	RegionLabel                  = InjectedLabelPrefix + "region"
	PinLabel                     = InjectedLabelPrefix + "pin"
	PipelineLabel                = InjectedLabelPrefix + "pipeline"
	UpdateLabel                  = InjectedLabelPrefix + "update"
	ExposedPortLabel             = InjectedLabelPrefix + "exposed_port"
	ExposeRangePortsLabel        = InjectedLabelPrefix + "expose_range_ports"
	IaaSInstanceIDLabel          = InjectedLabelPrefix + "iaas_instance_id"
	IpVersionLabel               = InjectedLabelPrefix + "ip_version"
)

var (
	ErrLabelKeyNotExist = errors.New("lable key not exist in resource")
)

// get label value by key
func GetLabelValue(key string, r Resource) (string, error) {
	value, ok := r.GetMeta().Metadata.Labels[key]
	if ok {
		return value, nil
	} else {
		return "", ErrLabelKeyNotExist
	}
}
