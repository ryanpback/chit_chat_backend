package helpers

import (
	"fmt"
	"reflect"
)

// ConvertReceiverIDs converts an interface type containing an array of int into a slice of int64
func ConvertReceiverIDs(r interface{}) []int64 {
	var ids = make([]int64, 0)
	rIds := reflect.ValueOf(r)

	for i := 0; i < rIds.Len(); i++ {
		id := int64(rIds.Index(i).Int())
		ids = append(ids, id)
	}

	return ids
}

// ConvertInterfaceToString converts an interface (likely from a map[string]interface{} type) to a string
func ConvertInterfaceToString(i interface{}) string {
	return fmt.Sprintf("%v", i)
}
