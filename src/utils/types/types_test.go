package types

import (
	"testing"
)

// check if implements Resource interface
func TestResource(t *testing.T) {
	test := func(re interface{}) Resource {
		return re.(Resource)
	}
	test(new(Container))
	test(new(Job))
	test(new(Nlb))
	test(new(Space))
	test(new(ReplicaSet))
	test(new(FloatIP))
	test(new(SecurityGroup))
	test(new(Script))
	test(new(BlockStore))
	test(new(Schedule))
	test(new(ImportService))
}
