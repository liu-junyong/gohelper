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
