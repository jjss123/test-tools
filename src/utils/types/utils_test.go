package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasContainerSpec(t *testing.T) {
	assert.True(t, HasContainerSpec(ContainerResource))
	assert.True(t, HasContainerSpec(ReplicaSetResource))
	assert.True(t, HasContainerSpec(JobResource))
	assert.False(t, HasContainerSpec(SecurityGroupResource))
}

func TestGetImage(t *testing.T) {
	con := NewContainer()
	con.Spec.Image = "a"
	assert.Equal(t, "a", GetImage(con))

	rep := NewReplicaSet()
	rep.Spec.Template.Image = "b"
	assert.Equal(t, "b", GetImage(rep))

	job := NewJob()
	job.Spec.Template.Image = "c"
	assert.Equal(t, "c", GetImage(job))

	assert.Equal(t, "", GetImage(NewScript()))
	assert.Equal(t, "", GetImage(NewSpace()))
}

func TestSetImageID(t *testing.T) {
	con := NewContainer()
	SetImageId(con, "a")
	assert.Equal(t, "a", con.Spec.ImageID)

	rep := NewReplicaSet()
	SetImageId(rep, "b")
	assert.Equal(t, "b", rep.Spec.Template.ImageID)

	job := NewJob()
	SetImageId(job, "c")
	assert.Equal(t, "c", job.Spec.Template.ImageID)
}

func TestGetValueByPath(t *testing.T) {
	data := make(map[string]interface{})
	b := make(map[string]interface{})
	data["A"] = b
	b["b"] = "asdf"
	result, err := GetValueByPath(data, "A.b")
	assert.Equal(t, nil, err)
	assert.Equal(t, "asdf", result)
}
