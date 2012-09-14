package main

import (
	"fmt"
	"strings"
)

var defaultFuncs = map[string]interface{}{
	"str":        String,
	"dec":        Decimal,
	"eq":         Equal,
	"eq_igncase": EqualIgnoreCase,
}

func String(v ...interface{}) string {
	if v == nil || len(v) == 0 {
		return "\"\""
	}
	switch x := v[0].(type) {
	case string:
		x = strings.Replace(x, `\`, `\\`, -1)
		x = strings.Replace(x, `"`, `\"`, -1)
		return fmt.Sprintf("\"%s\"", x)
	case float64:
		return fmt.Sprintf("\"%f\"", x)
	}
	return "\"\""
}

func Decimal(dec int, v ...interface{}) string {
	if v == nil || len(v) == 0 {
		return ""
	}
	if f, ok := v[0].(float64); ok {
		fmtstr := fmt.Sprintf("%%.%df", dec)
		return fmt.Sprintf(fmtstr, f)
	}
	return ""
}

func Equal(v ...interface{}) interface{} {
	if v == nil || len(v) < 2 {
		return nil
	}
	if v[0] == v[1] {
		return v[0]
	}
	return nil
}

func EqualIgnoreCase(s ...string) string {
	if len(s) < 2 {
		return ""
	}
	if strings.ToLower(s[0]) == strings.ToLower(s[1]) {
		return s[0]
	}
	return ""
}
