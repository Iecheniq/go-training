package type_assertion

import (
	"fmt"
)

func AssertTypeIsPrimitive(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("The value %v is int", t)
	case string:
		fmt.Printf("The value %v is string", t)
	case bool:
		fmt.Printf("The value %v is bool", t)
	case float32:
		fmt.Printf("The value %v is float32", t)
	case float64:
		fmt.Printf("The value %v is float64", t)
	default:
		fmt.Printf("Type %T is not a primitve", t)
	}
}
