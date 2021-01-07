package types

const (
	LabelChangeEvent = "LabelChange"

	ContainerScaleEvent        = "ContainerScale"
	ContainerUpgradeEvent      = "ContainerUpgrade"
	ContainerVolumeExtendEvent = "ContainerVolumeExtend"

	VMScaleEvent        = "VMScale"
	VMUpgradeEvent      = "VMUpgrade"
	VMVolumeExtendEvent = "VMVolumeExtend"

	ReplicaSetVerticalExpansionScaleEvent = "ReplicaSetVerticalExpansionScale"
	ReplicaSetVerticalReductionScaleEvent = "ReplicaSetVerticalReductionScale"
	ReplicaSetHorizontalScaleEvent        = "ReplicaSetHorizontalScale"
	ReplicaSetUpgradeEvent                = "ReplicaSetUpgrade"
	ReplicaSetVolumeExtendEvent           = "ReplicaSetVolumeExtend"
	ReplicaSetRecreateEvent               = "ReplicaSetRecreate"

	AddWhitelistEvent    = "AddWhitelist"
	RemoveWhitelistEvent = "RemoveWhitelist"

	AZMigrateEvent = "AZMigrateEvent"
	AZRecoverEvent = "AZRecoverEvent"

	AZMoveEvent = "AZMoveEvent"

	PipelineUpdateEvent = "PipelineUpdate"

	FloatipUpdateEvent = "FloatipUpdate"

	ReplicaSetSpecUpdateEvent = "SpecUpdate"

	ReplicaSetNlbUpdateEvent = "RSNlbUpdate"

	VMRecoverEvent = "VMRecover"
)

// update strategy type
const (
	RollingUpdateStepStrategy    = "step"
	RollingUpdatePercentStrategy = "percent"
)

type SpaceWhiteList struct {
	Ips []string
}

type AZMigrateInfo struct {
	RS     string
	FromAZ string
	ToAZ   string
}

type AZMoveInfo struct {
	RS     string
	FromAZ string
	ToAZ   string
}

type AZRecoverInfo struct {
	RS                string
	IsUsedOldInstance bool
}

type PipelineUpdate struct {
	Name string
	Step int
}

type VMRecoverInfo struct {
	VMName       string
	SnapshotName string
}

var UnmarshalSpaceUpdateEvent = func(e *SpaceUpdateEvent) (interface{}, error) {
	if e.Type == AddWhitelistEvent || e.Type == RemoveWhitelistEvent {
		whitelist := new(SpaceWhiteList)
		err := JsonUnmarshal(e.Content, whitelist)
		if err != nil {
			return nil, err
		}
		return whitelist, nil
	} else if e.Type == AZMigrateEvent {
		migrateInfo := new(AZMigrateInfo)
		err := JsonUnmarshal(e.Content, migrateInfo)
		if err != nil {
			return nil, err
		}
		return migrateInfo, nil
	} else if e.Type == AZMoveEvent {
		moveInfo := new(AZMoveInfo)
		err := JsonUnmarshal(e.Content, moveInfo)
		if err != nil {
			return nil, err
		}
		return moveInfo, err
	} else if e.Type == AZRecoverEvent {
		recoverInfo := new(AZRecoverInfo)
		err := JsonUnmarshal(e.Content, recoverInfo)
		if err != nil {
			return nil, err
		}
		return recoverInfo, nil
	} else if e.Type == PipelineUpdateEvent {
		pipelineUpdateInfo := new(PipelineUpdate)
		err := JsonUnmarshal(e.Content, pipelineUpdateInfo)
		if err != nil {
			return nil, err
		}
		return pipelineUpdateInfo, nil
	} else if e.Type == VMRecoverEvent {
		recoverInfo := new(VMRecoverInfo)
		err := JsonUnmarshal(e.Content, recoverInfo)
		if err != nil {
			return nil, err
		}
		return recoverInfo, nil
	}
	return SmartJsonUnmarshal(e.Content)
}
