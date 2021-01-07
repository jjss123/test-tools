package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewContainer(t *testing.T) {
	o := NewContainer()
	assert.Equal(t, ContainerResource, o.Kind)
	assert.Equal(t, ContainerIniting, o.Status.Phase)
	assert.NotNil(t, o.Status.Context)
	assert.NotNil(t, o.Metadata.Labels)
}

func TestNewReplicaSet(t *testing.T) {
	o := NewReplicaSet()
	assert.Equal(t, ReplicaSetResource, o.Kind)
	assert.Equal(t, ReplicaSetPending, o.Status.Phase)
	assert.NotNil(t, o.Status.Context)
	assert.NotNil(t, o.Metadata.Labels)
}

func TestNewJob(t *testing.T) {
	o := NewJob()
	assert.Equal(t, JobResource, o.Kind)
	assert.Equal(t, JobPending, o.Status.Phase)
	assert.NotNil(t, o.Status.Context)
	assert.NotNil(t, o.Metadata.Labels)
}

func TestNewSpace(t *testing.T) {
	o := NewSpace()
	assert.Equal(t, SpaceResource, o.Kind)
	assert.Equal(t, SpacePending, o.Status.Phase)
	assert.NotNil(t, o.Status.Context)
	assert.NotNil(t, o.Metadata.Labels)
}

func TestNewNlb(t *testing.T) {
	o := NewNlb()
	assert.Equal(t, NlbResource, o.Kind)
	assert.Equal(t, NlbIniting, o.Status.Phase)
	assert.NotNil(t, o.Status.Context)
	assert.NotNil(t, o.Metadata.Labels)
}

func TestNewFloatIP(t *testing.T) {
	o := NewFloatIP()
	assert.Equal(t, FloatIPResource, o.Kind)
	assert.Equal(t, FloatIPIniting, o.Status.Phase)
	assert.NotNil(t, o.Status.Context)
	assert.NotNil(t, o.Metadata.Labels)
}

func TestNewSecurityGroup(t *testing.T) {
	o := NewSecurityGroup()
	assert.Equal(t, SecurityGroupResource, o.Kind)
	assert.Equal(t, SecurityGroupIniting, o.Status.Phase)
	assert.NotNil(t, o.Status.Context)
	assert.NotNil(t, o.Metadata.Labels)
}

func TestNewScript(t *testing.T) {
	o := NewScript()
	assert.Equal(t, ScriptResource, o.Kind)
	assert.NotNil(t, o.Metadata.Labels)
}

func TestNewBlockStore(t *testing.T) {
	o := NewBlockStore()
	assert.Equal(t, BlockStoreResource, o.Kind)
	assert.Equal(t, BlockStorePending, o.Status.Phase)
	assert.NotNil(t, o.Status.Context)
	assert.NotNil(t, o.Metadata.Labels)
}
