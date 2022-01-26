package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"sort"
)

const defaultKey = "ID"

// StructsToMap 默认通过 key ID 分组成 map
func StructsToMap(in interface{}) (map[string]string, error) {
	return StructsToOfKey(defaultKey,in)
}

// StructsToOfKey 自定义 key 分组成 map
func StructsToOfKey(key string, in interface{}) (map[string]string, error) {
	out := make(map[string]string)

	value := reflect.ValueOf(in)

	if in == nil{
		return out,nil
	}

	if value.Kind() == reflect.Ptr{
		return out,nil
	}

	if value.Kind() == reflect.Struct{
		t := value.Type()
		_, b := t.FieldByName(key)
		if b {
			jsonKey, _ := json.Marshal(value.FieldByName(key).Interface())
			jsonValue, _ := json.Marshal(value.Interface())
			out[string(jsonKey)] = string(jsonValue)
		}else{
			log.Println("struct doesn't has the field ",key)
			return out,fmt.Errorf("struct doesn't has the field : "+key)
		}
		return out,nil
	}

	if value.Kind() == reflect.Slice{
		ln := value.Len()
		for i := 0; i < ln; i++ {
			index := value.Index(i)
			t := index.Type()
			_, b := t.FieldByName(key)
			if b {
				jsonKey, _ := json.Marshal(index.FieldByName(key).Interface())
				jsonValue, _ := json.Marshal(index.Interface())
				out[string(jsonKey)] = string(jsonValue)
			}
		}
	}

	return out, nil
}

func MapToStructs(dataMap map[string]string) interface{} {
	var result []interface{}
	if len(dataMap) != 0 {
		keys := sortKey(dataMap)
		for _,v:= range keys{
			var data interface{}
			_ = json.Unmarshal([]byte(dataMap[v]), &data)
			result = append(result,data)
		}
	}
	return result
}

func sortKey(dataMap map[string]string) []string {
	var keys []string
	for k,_:= range dataMap{
		keys = append(keys,k)
	}
	sort.Strings(keys)
	return keys
}

func changeRv(rv reflect.Value) {
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() == reflect.Struct {
		changeStruct(rv)
	}
	if rv.Kind() == reflect.Slice {
		changeSlice(rv)
	}
}

func changeSlice(rv reflect.Value) {
	ln := rv.Len()
	if ln == 0 && rv.CanAddr() {
		var elem reflect.Value

		typ := rv.Type().Elem()
		if typ.Kind() == reflect.Ptr {
			elem = reflect.New(typ.Elem())
		}
		if typ.Kind() == reflect.Struct {
			elem = reflect.New(typ).Elem()
		}

		rv.Set(reflect.Append(rv, elem))
	}

	ln = rv.Len()
	for i := 0; i < ln; i++ {
		changeRv(rv.Index(i))
	}
}

// assumes rv is a struct
func changeStruct(rv reflect.Value) {
	if !rv.CanAddr() {
		return
	}
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)

		switch field.Kind() {
		case reflect.String:
			field.SetString("fred")
		case reflect.Int:
			field.SetInt(54)
		default:
			fmt.Println("unknown field")
		}
	}
}

