package errors

const (
	InvalidInput                        = -2
	InternalError                       = -1
	UpdateConflict                      = -9
	CreateResourceInAPIServerError      = -5
	BatchUpdateResourceInAPIServerError = -6
	InvalidChart                        = -7
	ScriptTimeout                       = -8
	ScriptError                         = -3
	ScheduleError                       = -4
	QuotaExceeded                       = -10
	InstanceNotExist                    = 4
	InstanceExistedInHistory            = 5
	InstanceCreateFailed                = 6
	InstanceCreating                    = 7
	InstanceDeleteFailed                = 8
	InstanceProtectDelete               = 9
	InstanceWaitingDeleteResult         = 10
)

var ErrorTypeMapping = map[int]string{
	-1: "InternalError",
	-2: "InvalidInput",
	-3: "ScriptError",
	-4: "ScheduleError",
	-5: "CreateResourceInAPIServerError",
	-6: "BatchUpdateResourceInAPIServerError",
	-7: "InvalidChart",
	-8: "ScriptTimeout",
	-9: "UpdateConflict",
}

//error messages
const (
	MsgServiceNotFound             = "service not found"
	MsgServiceRequired             = "service required"
	MsgServiceInstanceNotFound     = "instance not found"
	MsgServiceInstanceExisted      = "instance already existed"
	MsgChartExisted                = "chart already existed or existed before"
	MsgChartNotFound               = "chart not found"
	MsgFileNotFound                = "file not found"
	MsgSourceTypeNotFound          = "source type not found"
	MsgSourceTypeNameEmpty         = "source type should not empty"
	MsgRequestAlreadyExecuted      = "same requestId execute before, please pass a new requestId when you perform action on instance every time"
	MsgFlavorTranslateYamlNotFound = "flavor translation file not found"
	MsgFlavorTranslateYamlInvalid  = "flavor translation file is not a valid yaml"
	MsgFlavorTranslationNotFound   = "cannot find flavor translation in file"
	MsgVarsYamlInvalid             = "vars file is not a valid yaml"
	MsgNoFileInChart               = "no file found in chart"
	MsgReleaseNotFound             = "release with passed service and version not found"
	MsgVersionRequired             = "version required"
	MsgPresetWhiteListFormatWarn   = "presetWhiteList should be this format ip/mask, such as 192.168.0.1/16"
	MsgInstanceIdEmpty             = "instanceId should not empty"
	MsgIpPrefixRequired            = "should provide at least one ipPrefix"
	MsgIpPrefixIncorrectFormat     = "ipPrefix should be of type []string"
	MsgAddWhitelistToOpenedSpace   = "instance access is opened, add white list will not affect the access control"
	MsgAddWhitelistToClosedSpace   = "instance access is closed, remove white list will not affect the access control"
	MsgNameEmpty                   = "name should not empty"
	MsgTagEmpty                    = "tag should not empty"
	MsgImageNotFound               = "image not found"
	MsgChecksumEmpty               = "checksum should not empty"
	MsgImageExist                  = "image already exist, please use a unique name and tag"
	MsgImageChecksumNotMatched     = "image checksum not match, please use md5sum to calculate checksum"
	MsgCookieNotExist              = "cookie not exist"
	ScheduleResourceNumOverOne     = "schedule resource num over one"
	ServiceNameCanNotEmpty         = "service name can not empty"
	MsgMatrixValidateParamFailed   = "matrix validate param failed"
	MsgReleaseHasCreatedInstance   = "release can not be deleted because there are instances that have been created"
	MsgReleaseIsInuse              = "release cannot be deleted becasuse it is inuse version"
	MsgChartAlreadyBindAction      = "chart can not be deleted because it already bind action"
	MsgStatusEmpty                 = "status required"
	MsgWaitExposetagTimeout        = "wait expose tag timeout"
	MsgQuotaExceeded               = "quota exceeded, please clean up resources"

	// Flavor
	MsgFlavorNameRequired     = "flavor name can not empty"
	MsgCPUNumIllegal          = "cpu num illegal"
	MsgMemoryNumIllegal       = "memory num illegal"
	MsgDiskNumillegal         = "disk num illegal"
	MsgFlavorNotFound         = "flavor not found"
	MsgProtectTimeIllegal     = "protect_time must be int"
	MsgUnmarshalRequestFailed = "unmarshal request failed"

	// Mysql
	MsgInsertDataIntoDBFailed = "insert data into db failed"
	MsgQueryDataFromDBFailed  = "get records from db failed"
	MsgDeleteDataFromDBFailed = "delete data from db failed"
	MsgCompareAndSwapFailed   = "compare and save data failed"

	// Image
	MsgImageHasExisted = "current upload image already existed in glance"

	// Region
	MsgRegionRequired   = "region required"
	MsgRegionNotExisted = "region not existed"

	// Action
	MsgActionNotFound = "action not found in service"

	// Instance
	MsgInstanceCreating      = "instance exists but creating"
	MsgInstanceCreatedFailed = "instance exists but created failed"
	MsgInstanceExisted       = "instance exists and created successful"
	MsgInstanceDeleteFailed  = "instance delete failed, it may be creating"


	// Snapshot
	MsgSnapshotNameExist = "snapshot exists"

	MsgPinRequired = "pin required"


	//Proxy
	MsgProxyParseUrlFailed = "proxy parse url failed"
)
