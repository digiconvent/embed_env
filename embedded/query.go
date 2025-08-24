package embedded

import (
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func ToQuery(someStruct any) (string, error) {
	v := reflect.ValueOf(someStruct)
	t := reflect.TypeOf(someStruct)

	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("StructToQuery: expected struct but got %s", t.Kind())
	}

	var params []string
	for i := range t.NumField() {
		field := t.Field(i)
		value := v.Field(i)

		if field.PkgPath != "" {
			continue
		}

		key := field.Tag.Get("name")
		if key == "" {
			key = field.Name
		}

		if isZero(value) {
			continue
		}

		params = append(params, fmt.Sprintf("%s=%s", url.QueryEscape(key), url.QueryEscape(fmt.Sprint(value.Interface()))))
	}
	sort.Strings(params)
	result := strings.Join(params, "&")
	return result, nil
}

func isZero(v reflect.Value) bool {
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}

func FromQuery(someStruct any, query *string) error {
	values, err := url.ParseQuery(*query)
	if err != nil {
		return err
	}

	v := reflect.ValueOf(someStruct).Elem()
	t := v.Type()

	for i := range t.NumField() {
		field := t.Field(i)
		value := v.Field(i)

		if field.PkgPath != "" {
			continue
		}

		key := field.Tag.Get("name")
		if key == "" {
			key = field.Name
		}

		if paramValues, ok := values[key]; ok && len(paramValues) > 0 {
			raw := paramValues[0]
			switch value.Kind() {
			case reflect.String:
				value.SetString(raw)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if iv, err := strconv.ParseInt(raw, 10, 64); err == nil {
					value.SetInt(iv)
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				if uv, err := strconv.ParseUint(raw, 10, 64); err == nil {
					value.SetUint(uv)
				}
			case reflect.Float32, reflect.Float64:
				if fv, err := strconv.ParseFloat(raw, 64); err == nil {
					value.SetFloat(fv)
				}
			case reflect.Bool:
				if bv, err := strconv.ParseBool(raw); err == nil {
					value.SetBool(bv)
				}
			}
		}
	}

	return nil
}
