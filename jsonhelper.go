package gohelper

import (
	"github.com/liu-junyong/go-logger/logger"
	"encoding/json"
	"strconv"
	"reflect"
)


func Json2int(hh interface{}) int32 {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
		}
	}()

	if hh == nil {
		return 0
	}

	heifan := 0
	switch hh.(type) {
	case float64:
		heifan = int(hh.(float64))
		return int32(heifan)
	case int32:
		heifan = int(hh.(int32))
		return int32(heifan)
	case int64:
		heifan = int(hh.(int64))
		return int32(heifan)
	case string:
		heifan, _ = strconv.Atoi(hh.(string))
		return int32(heifan)
	}
	return int32(hh.(int))
}



func Json2float32(hh interface{}) float32 {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
		}
	}()

	if hh == nil {
		return 0
	}

	heifan := 0
	switch hh.(type) {
	case float64:
		heifan = int(hh.(float64))
		return float32(heifan)
	case float32:
		heifan = int(hh.(float64))
		return float32(heifan)
	case int32:
		heifan = int(hh.(int32))
		return float32(heifan)
	case int64:
		heifan = int(hh.(int64))
		return float32(heifan)
	}
	return float32(hh.(int))
}

func Json2Bool(hh interface{}) bool {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
		}
	}()

	if hh == nil {
		return false
	}

	heifan := false
	switch hh.(type) {
	case bool:
		heifan = hh.(bool)
	}
	return heifan
}

func Json2String(hh interface{}) string {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
		}
	}()

	if hh == nil {
		return ""
	}

	heifan := ""
	switch hh.(type) {
	case string:
		heifan = hh.(string)
	}
	return heifan
}

func ParseValue(opt string, key1 string,key2 string) string {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
		}
	}()

	v := make(map[string]interface{})
	json.Unmarshal([]byte(opt), &v)
	ret1 := Json2String(v[key1])
	if len(ret1) == 0 {
		return  Json2String(v[key2])
	}
	return ret1
}

func ParseValueInt(opt string, key1 string, key2 string) int32 {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
		}
	}()

	v := make(map[string]interface{})
	json.Unmarshal([]byte(opt), &v)
	ret1 := Json2int(v[key1])
	if ret1 == 0 {
		return Json2int(v[key2])
	}
	return ret1
}



func ParseValueBool(opt string, key1 string, key2 string) bool {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
		}
	}()

	v := make(map[string]interface{})
	json.Unmarshal([]byte(opt), &v)
	ret1 := Json2Bool(v[key1])
	if ret1 == false {
		return Json2Bool(v[key2])
	}
	return ret1
}




func ToSlice(arr interface{}) []interface{} {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
		}
	}()

	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("toslice arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}



func Json2sliceObj(hh interface{}) []interface{} {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
		}
	}()

	ret := make([]interface{},0)
	if hh == nil {
		return ret
	}

	v := reflect.ValueOf(hh)
	if v.Kind() != reflect.Slice {
		return ret
	}

	ret1 := ToSlice(hh)
	for _,v := range ret1  {
		k := Json2String( v )
		ret = append( ret ,k )
	}

	return ret

}




func Json2sliceStr(hh interface{}) []string {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
		}
	}()

	ret := make([]string,0)
	if hh == nil {
		return ret
	}

	ret1 := ToSlice(hh)
	for _,v := range ret1  {
		k := Json2String( v )
		ret = append( ret ,k )
	}

	return ret

}


func Json2slice(hh interface{}) []int32 {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
		}
	}()

	ret := make([]int32,0)
	if hh == nil {
		return ret
	}

	ret1 := ToSlice(hh)
	for _,v := range ret1  {
		k := Json2int( v )
		if k > 0 {
			ret = append( ret ,k )
		}
	}

	return ret

}