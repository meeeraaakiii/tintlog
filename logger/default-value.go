package tl

import (
	"reflect"
	"strings"
)

/*
Fills zero-valued fields in dst from def.

It logs each fill via logf(fieldName, defaultValue) if provided.

Respects struct tag `default:"skip"` to skip a field.

For now this function only handles top-level fields.
*/
func ApplyDefaults[T any](dst *T, def T, logf func(field string, defVal any)) {
	v := reflect.ValueOf(dst).Elem()
	dv := reflect.ValueOf(def)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)

		// Only exported, settable fields
		f := v.Field(i)
		if !f.CanSet() || sf.PkgPath != "" { // unexported
			continue
		}

		// Optional opt-out
		if sf.Tag.Get("default") == "skip" {
			continue
		}

		// Zero or replaceable?
		if IsZeroOrEmpty(f) {
			// Fill from default
			f.Set(dv.Field(i))

			if logf != nil {
				name := jsonOrFieldName(sf)
				logf(name, dv.Field(i).Interface())
			}
		}
	}
}

// Treats zero values as zero, and ALSO treats slices with len==0 as zero
func IsZeroOrEmpty(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		return v.Len() == 0
	case reflect.Pointer, reflect.Interface:
		if v.IsNil() {
			return true
		}
		return IsZeroOrEmpty(v.Elem())
	default:
		// For everything else, use the standard zero check.
		// reflect.Value.IsZero is available in modern Go and faster than DeepEqual-to-zero.
		return v.IsZero()
	}
}

func jsonOrFieldName(sf reflect.StructField) string {
	tag := sf.Tag.Get("json")
	if tag == "-" || tag == "" {
		return sf.Name
	}
	// strip omitempty, etc.
	if idx := strings.Index(tag, ","); idx >= 0 {
		tag = tag[:idx]
	}
	if tag == "" {
		return sf.Name
	}
	return tag
}
