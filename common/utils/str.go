package utils

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/mangenotwork/extras/common/logger"
)

// Str2Int64 string -> int
func Str2Int(str string) int {
	if i, err := strconv.Atoi(str); err == nil {
		return i
	}
	return 0
}

// Str2Int64 string -> int32
func Str2Int32(str string) int32 {
	if i, err := strconv.ParseInt(str, 10, 32); err == nil {
		return int32(i)
	}
	return 0
}

// Str2Int64 string -> int64
func Str2Int64(str string) int64 {
	if i, err := strconv.ParseInt(str, 10, 64); err == nil {
		return i
	}
	return 0
}

// Str2Float string -> float64
func Str2Float64(str string) float64 {
	v1, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return v1
}

func Str2Bool(str string) bool {
	for _, v := range []string{"0","f","F","FALSE","false","False","no","否"}{
		if str == v {
			return false
		}
	}
	return true
}

func Int642Str(i int64) string {
	return strconv.FormatInt(i,10)
}


// Any2String interface{} -> string
func Any2String(i interface{}) string {
	if i == nil {
		return ""
	}
	if reflect.ValueOf(i).Kind() == reflect.String{
		return i.(string)
	}
	var buf bytes.Buffer
	stringValue(reflect.ValueOf(i), 0, &buf)
	return buf.String()
}

func stringValue(v reflect.Value, indent int, buf *bytes.Buffer) {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Struct:
		buf.WriteString("{\n")
		for i := 0; i < v.Type().NumField(); i++ {
			ft := v.Type().Field(i)
			fv := v.Field(i)
			if ft.Name[0:1] == strings.ToLower(ft.Name[0:1]) {
				// ignore unexported fields
				continue
			}
			if (fv.Kind() == reflect.Ptr || fv.Kind() == reflect.Slice) && fv.IsNil() {
				// ignore unset fields
				continue
			}
			buf.WriteString(strings.Repeat(" ", indent+2))
			buf.WriteString(ft.Name + ": ")
			if tag := ft.Tag.Get("sensitive"); tag == "true" {
				buf.WriteString("<sensitive>")
			} else {
				stringValue(fv, indent+2, buf)
			}
			buf.WriteString(",\n")
		}
		buf.WriteString("\n" + strings.Repeat(" ", indent) + "}")

	case reflect.Slice:
		nl, id, id2 := "", "", ""
		if v.Len() > 3 {
			nl, id, id2 = "\n", strings.Repeat(" ", indent), strings.Repeat(" ", indent+2)
		}
		buf.WriteString("[" + nl)
		for i := 0; i < v.Len(); i++ {
			buf.WriteString(id2)
			stringValue(v.Index(i), indent+2, buf)

			if i < v.Len()-1 {
				buf.WriteString("," + nl)
			}
		}
		buf.WriteString(nl + id + "]")

	case reflect.Map:
		buf.WriteString("{\n")
		for i, k := range v.MapKeys() {
			buf.WriteString(strings.Repeat(" ", indent+2))
			buf.WriteString(k.String() + ": ")
			stringValue(v.MapIndex(k), indent+2, buf)

			if i < v.Len()-1 {
				buf.WriteString(",\n")
			}
		}
		buf.WriteString("\n" + strings.Repeat(" ", indent) + "}")

	default:
		format := "%v"
		switch v.Interface().(type) {
		case string:
			format = "%q"
		}
		_,_= fmt.Fprintf(buf, format, v.Interface())
	}
}

// P2E panic -> error
func P2E() {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Panic error: %v", r)
		}
	}()
}

//切片去除空元素
func SliceDelNullString(sli []string) []string {
	rse := make([]string, 0)
	for _,v := range sli {
		if v != "" {
			rse = append(rse, v)
		}
	}
	return rse
}

