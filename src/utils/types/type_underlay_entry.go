package types

import "time"

// UnderlayEntry Phase
const (
	// Pending
	UnderlayEntryPending = CommonPending
	// Running
	UnderlayEntryRunning = CommonRunning
	// ToDelete
	UnderlayEntryToDelete = CommonToDelete
)

const (
	UnderlayEntryNetworkContainerKind = ContainerResource
	UnderlayEntryNetworkVMKind        = VMResource
	UnderlayEntryNetworkNCKind        = NCResource
	UnderlayEntryNetworkNlbKind       = NlbResource
)

const (
	UnderlayEntryOpenApiType = "openapi"
	UnderlayEntryWebAppType  = "webapp"
)

type UnderlayEntry struct {
	UnderlayEntryDef     `yaml:",inline"`
	UnderlayEntryRuntime `yaml:",inline"`
}

type UnderlayEntryDef struct {
	Meta `yaml:",inline"`
	Spec UnderlayEntrySpec `yaml:"spec" json:"spec"`
}

type UnderlayEntrySpec struct {
	Type         string `yaml:"type" json:"type"`
	Tag          string `yaml:"tag" json:"tag"`
	Protocol     string `yaml:"protocol" json:"protocol"`
	Port         int    `yaml:"port" json:"port"`
	ResourceKind string `yaml:"resource_kind" json:"resource_kind"`
	ResourceName string `yaml:"resource_name" json:"resource_name"`

	IAMResourceType    string `yaml:"iam_resource_type" json:"iam_resource_type,omitempty"`
	IAMResourceID      string `yaml:"iam_resource_id" json:"iam_resource_id,omitempty"`
	IAMOperationID     string `yaml:"iam_operation_id" json:"iam_operation_id,omitempty"`
	IAMSubResourceType string `yaml:"iam_sub_resource_type" json:"iam_sub_resource_type,omitempty"`
	IAMSubResourceID   string `yaml:"iam_sub_resource_id" json:"iam_sub_resource_id,omitempty"`
	IAMServiceName     string `yaml:"iam_service_name" json:"iam_service_name,omitempty"`
}

type UnderlayEntryRuntime struct {
	Status UnderlayEntryStatus `yaml:"status" json:"status"`
}

type UnderlayEntryStatus struct {
	BasicStatus `yaml:",inline"`
	// now access type: floatip
	AccessType   string `yaml:"access_type" json:"access_type"`
	AccessDomain string `yaml:"access_domain" json:"access_domain"`
}

func (t *UnderlayEntry) GetDef() interface{} {
	return &(t.UnderlayEntryDef)
}

func (t *UnderlayEntry) SetPhase(phase string) {
	t.Status.Phase = phase
}

func (t *UnderlayEntry) GetPhase() string {
	return t.Status.Phase
}

func (t *UnderlayEntry) SetRequestId(requestId string) {
	t.Status.RequestId = requestId
}

func (t *UnderlayEntry) RefreshUpdatetime() {
	t.Status.Updatetime = time.Now().Round(0)
}
