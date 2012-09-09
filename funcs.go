package main

import (
	"fmt"
)

var defaultFuncs = map[string]interface{}{
	"str": String,
	"dec": Decimal,
}

func String(v interface{}) string {
	switch x := v.(type) {
	case string:
		// FIXME x should be escaped
		return fmt.Sprintf("\"%s\"", x)
	case float64:
		return fmt.Sprintf("\"%.2f\"", x)
	}
	return ""
}

func Decimal(v interface{}) string {
	if f, ok := v.(float64); ok {
		return fmt.Sprintf("%.2f", f)
	}
	return ""
}
