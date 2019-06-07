package report

import (
	"fmt"
	"reflect"
	"strings"
)

// Text report
func Text(v interface{}) string {
	val := reflect.ValueOf(v)
	var sb strings.Builder

	text(&sb, &val)

	return sb.String()
}

func text(sb *strings.Builder, v *reflect.Value) {
	switch v.Kind() {
	case reflect.Struct:
		structName := v.Type().Name()
		if len(structName) == 0 {
			structName = "Anonymous"
		}
		sb.WriteString(structName)
		sb.WriteString("(")
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(v.Type().Field(i).Name + " : ")
			val := v.Field(i)
			text(sb, &val)
		}
		sb.WriteString(")")
	case reflect.String:
		sb.WriteString(v.String())
	case reflect.Int:
		sb.WriteString(fmt.Sprintf("%d", v.Int()))
	}
}
