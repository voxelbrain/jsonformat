package main

import (
	"fmt"
	"strings"
)

var defaultFuncs = map[string]interface{}{
	"str": String,
	"dec": Decimal,
}

func String(v interface{}) string {
	switch x := v.(type) {
	case string:
		x = strings.Replace(x, `\`, `\\`, -1)
		x = strings.Replace(x, `"`, `\"`, -1)
		return fmt.Sprintf("\"%s\"", x)
	case float64:
		return fmt.Sprintf("\"%f\"", x)
	}
	return ""
}

func Decimal(dec int, v interface{}) string {
	if f, ok := v.(float64); ok {
		fmtstr := fmt.Sprintf("%%.%df", dec)
		return fmt.Sprintf(fmtstr, f)
	}
	return ""
}
