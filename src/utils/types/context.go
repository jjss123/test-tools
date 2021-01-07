package types

import (
	"fmt"
	"strings"
)

var CtxComputeBlockStoreFormat = "VOLUME_FORMAT_%s"

func NewComputeBlockStoreFormatContext(bsName string) string {
	return fmt.Sprintf(CtxComputeBlockStoreFormat, strings.ToUpper(bsName))
}
