package consts

import "fmt"

const (
	verMajor = 0
	verMinor = 1
	verPatch = 0
)

var Ver = fmt.Sprintf("%d.%d.%d", verMajor, verMinor, verPatch)
