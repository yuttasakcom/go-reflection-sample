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
			val := v.Field(i)

			if i > 0 {
				sb.WriteString(", ")
			}
			displayName := v.Type().Field(i).Name
			tagName, ok := v.Type().Field(i).Tag.Lookup("report")
			if ok {
				tags := strings.Split(tagName, ",")
				if len(tags) == 2 {
					displayName = tags[0]
					if displayName == "" {
						displayName = v.Type().Field(i).Name
					}
					if tags[1] == "uppercase" {
						upper := strings.ToUpper(v.Field(i).String())
						val = reflect.ValueOf(upper)
					}
				}
			}
			sb.WriteString(displayName + " : ")

			text(sb, &val)
		}
		sb.WriteString(")")
	case reflect.String:
		sb.WriteString(v.String())
	case reflect.Int:
		sb.WriteString(fmt.Sprintf("%d", v.Int()))
	}
}
