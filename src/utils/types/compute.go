package types

type ComputeInfo struct {
	Host           string `json:"host"`
	Service        string `json:"service"`
	InstanceID     string `json:"instance_id"`
	ResourceName   string `json:"resource_name"`
	Kind           string `json:"kind"`
	Flavor         string `json:"flavor"`
	Pin            string `json:"pin"`
	IaaSInstanceID string `json:"iaas_instance_id"`
	Image          string `json:"image"`
	ImageID        string `json:"image_id"`
	CreateTime     string `json:"create_time"`
	UpdateTime     string `json:"update_time"`
}

type ServiceConsumeCpuOfHost struct {
	Region       string `json:"region"`
	AZ           string `json:"az"`
	PoolName     string `json:"pool_name"`
	Machine      string `json:"machine"`
	IP           string `json:"ip"`
	Service      string `json:"service"`
	ServiceCode  string `json:"service_code"`
	Cpus         int    `json:"cpus"`
	CpusReserved int    `json:"cpus_reserved"`
	TotalVcpus   int    `json:"total_vcpus"`
	CpusUsed     int    `json:"cpus_used"`
}
