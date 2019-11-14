/**
 * @Author: Tomonori
 * @Date: 2019/11/14 15:27
 * @Desc:
 */
package main

import (
	"errors"
	"reflect"
)

type StructPropertyUtility struct {
	cacheStructProperty map[string]*PropertyBase
}

func NewStructPropertyUtility() *StructPropertyUtility {
	utility := &StructPropertyUtility{cacheStructProperty: make(map[string]*PropertyBase)}
	return utility
}

//获取所有属性信息
//TODO 现在是直接获取反射，需要封装个property的专门的结构体来记录所有的属性信息
func (s *StructPropertyUtility) GetProperties(origin interface{}) (result *PropertyBase) {
	if !s.isStructType(origin) {
		return
	}

	name := s.getStructName(origin)
	result = s.cacheStructProperty[name]
	if result == nil {
		base := &PropertyBase{
			Type:  reflect.TypeOf(origin).Elem(),
			Value: reflect.ValueOf(origin).Elem(),
		}
		s.cacheStructProperty[name] = base
		return base
	}

	return
}

//转换结构体属性和值为映射
func (s *StructPropertyUtility) StructToMap(origin interface{}) (result map[string]interface{}) {
	if !s.isStructType(origin) {
		return
	}
	result = make(map[string]interface{})
	properties := s.GetProperties(origin)
	var1 := 0
	for var1 < properties.Value.NumField() {
		field := properties.Value.Field(var1)
		if field.CanSet() {
			result[properties.Type.Field(var1).Name] = s.getRealValue(field)
		}
		var1++
	}
	return
}

//拷贝非空字段值
func (s *StructPropertyUtility) CopyNotNull(origin, target interface{}) (err error) {
	if origin == nil || target == nil {
		return errors.New("struct not be null")
	}

	if !s.isStructType(origin) || !s.isStructType(target) {
		return errors.New("params has not be type struct")
	}

	//Elem() 如果取到值非Interface 或 pointer，使用Elem()方法转换为源地址的reflect.Value或reflect.Type，才能进行后续操作
	//否则就是指针或接口的Value或Type了
	//而且用了这个必定要传指针或接口类型的参数
	originType := reflect.TypeOf(origin).Elem()
	targetType := reflect.TypeOf(target).Elem()
	originValue := reflect.ValueOf(origin).Elem()
	targetValue := reflect.ValueOf(target).Elem()
	originFieldSize := originType.NumField()
	targetFieldSize := targetType.NumField()
	var1 := 0
	for var1 < originFieldSize {
		fieldName := originType.Field(var1).Name
		fieldValue := originValue.Field(var1)
		var2 := 0
		for var2 < targetFieldSize {
			if targetType.Field(var2).Name == fieldName && targetValue.Field(var2).CanSet() {
				targetValue.Field(var2).Set(fieldValue)
			}
			var2++
		}
		var1++
	}
	return nil
}

func (s *StructPropertyUtility) getRealValue(valueOf reflect.Value) (result interface{}) {
	switch valueOf.Kind() {
	case reflect.Bool:
		result = valueOf.Bool()
	case reflect.Int:
		result = valueOf.Int()
	case reflect.Int64:
		result = valueOf.Int()
	case reflect.Float32:
		result = valueOf.Float()
	case reflect.Float64:
		result = valueOf.Float()
	case reflect.String:
		result = valueOf.String()
	}
	return
}

func (s *StructPropertyUtility) getStructName(entity interface{}) string {
	entityType := reflect.TypeOf(entity).Elem()
	return entityType.Name()
}

func (s *StructPropertyUtility) isStructType(entity interface{}) bool {
	typeOf := reflect.TypeOf(entity).Elem()
	if typeOf.Kind() != reflect.Struct {
		return false
	}
	return true
}
