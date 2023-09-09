package k8s

import (
	"encoding/json"

	"github.com/raanfefu/basajaun/admission/pkg"
)

func AddPatches(arr []pkg.Patch, operation string, path string, value interface{}) []pkg.Patch {
	patch := &pkg.Patch{
		Op:    operation,
		Path:  path,
		Value: value,
	}
	if arr == nil {
		arr = make([]pkg.Patch, 1)
	}
	arr = append(arr, *patch)
	return arr
}

func EncodeBase64Patches(arr []pkg.Patch) []byte {
	ajson, _ := json.MarshalIndent(arr, "", "   ")
	spatches := string(ajson)
	return []byte(spatches)
}
