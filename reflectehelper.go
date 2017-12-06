package gohelper

import (
    "reflect"
    "fmt"
    "errors"
    "strings"
    "github.com/liu-junyong/go-logger/logger"
)

func SetField(obj interface{}, name string, value interface{}) error {
    structValue := reflect.ValueOf(obj).Elem()
    structFieldValue := structValue.FieldByName(name)

    if !structFieldValue.IsValid() {
        return fmt.Errorf("No such field: %s in obj", name)
    }

    if !structFieldValue.CanSet() {
        return fmt.Errorf("Cannot set %s field value", name)
    }

    structFieldType := structFieldValue.Type()
    val := reflect.ValueOf(value)
    if structFieldType != val.Type() {
        return errors.New("Provided value type didn't match obj field type")
    }

    structFieldValue.Set(val)
    return nil
}

type MyStruct struct {
    Name string
    Age  int64
}

func (s *MyStruct) FillStruct(m map[string]interface{}) error {
    for k, v := range m {
        err := SetField(s, k, v)
        if err != nil {
            return err
        }
    }
    return nil
}



func Struct2Map_lower(obj interface{}) map[string]interface{} {
    defer func() {
        if r := recover(); r != nil {
            logger.Error(r)
        }
    }()

    t := reflect.TypeOf(obj)
    v := reflect.ValueOf(obj)

    var data = make(map[string]interface{})
    for i := 0; i < t.NumField(); i++ {
        data[ strings.ToLower(t.Field(i).Name)] = v.Field(i).Interface()
        //logger.Info(t.Field(i))
    }
    return data
}



func Struct2Map(obj interface{}) map[string]interface{} {
    defer func() {
        if r := recover(); r != nil {
            logger.Error(r)
        }
    }()

    t := reflect.TypeOf(obj)
    v := reflect.ValueOf(obj)

    var data = make(map[string]interface{})
    for i := 0; i < t.NumField(); i++ {
        data[t.Field(i).Name] = v.Field(i).Interface()
    }
    return data
}


//类型转换
func TypeConversion(value string, ntype string) (reflect.Value, error) {
    if ntype == "string" {
        return reflect.ValueOf(value), nil
    } else if ntype == "time.Time" {
        t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
        return reflect.ValueOf(t), err
    } else if ntype == "Time" {
        t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
        return reflect.ValueOf(t), err
    } else if ntype == "int" {
        i, err := strconv.Atoi(value)
        return reflect.ValueOf(i), err
    } else if ntype == "int8" {
        i, err := strconv.ParseInt(value, 10, 64)
        return reflect.ValueOf(int8(i)), err
    } else if ntype == "int32" {
        i, err := strconv.ParseInt(value, 10, 64)
        return reflect.ValueOf(int64(i)), err
    } else if ntype == "int64" {
        i, err := strconv.ParseInt(value, 10, 64)
        return reflect.ValueOf(i), err
    } else if ntype == "float32" {
        i, err := strconv.ParseFloat(value, 64)
        return reflect.ValueOf(float32(i)), err
    } else if ntype == "float64" {
        i, err := strconv.ParseFloat(value, 64)
        return reflect.ValueOf(i), err
    }

    //else if .......增加其他一些类型的转换

    return reflect.ValueOf(value), errors.New("未知的类型：" + ntype)
}

