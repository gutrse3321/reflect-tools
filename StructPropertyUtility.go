/**
 * @Author: Tomonori
 * @Date: 2019/11/14 15:27
 * @Desc:
 */
package main

import (
	"errors"
	"reflect"
	"strings"
)

type StructPropertyUtility struct {
	cacheStructProperties map[string][]*PropertyBase
}

func NewStructPropertyUtility() *StructPropertyUtility {
	utility := &StructPropertyUtility{
		cacheStructProperties: make(map[string][]*PropertyBase),
	}
	return utility
}

//获取所有可导出属性信息
func (s *StructPropertyUtility) GetProperties(origin interface{}) ([]*PropertyBase, error) {
	if !s.isStructType(origin) {
		return nil, errors.New("must be has type struct")
	}

	name := s.getStructName(origin)
	result := s.cacheStructProperties[name]
	if result == nil {
		typeOf, valueOf := s.getElem(origin)
		fieldNum := typeOf.NumField()
		var1 := 0
		for var1 < fieldNum {
			fieldType, fieldValue := typeOf.Field(var1), valueOf.Field(var1)
			if fieldValue.CanSet() {
				base := &PropertyBase{
					Name:      fieldType.Name,
					Type:      fieldType.Type.String(),
					OtherInfo: fieldType,
					ValueOf:   fieldValue,
				}
				result = append(result, base)
			}
			var1++
		}
		s.cacheStructProperties[name] = result
		return result, nil
	}

	return result, nil
}

//转换结构体属性和值为映射
func (s *StructPropertyUtility) StructToMap(origin interface{}) (result map[string]interface{}, err error) {
	result = make(map[string]interface{})
	properties, err := s.GetProperties(origin)
	if err != nil {
		return nil, err
	}

	for _, item := range properties {
		result[item.Name] = s.getRealValue(item.ValueOf)
	}
	return
}

//拷贝非空字段值
func (s *StructPropertyUtility) CopyNotNull(origin, target interface{}) (err error) {
	if origin == nil || target == nil {
		return errors.New("struct not be null")
	}

	originProperties, err := s.GetProperties(origin)
	if err != nil {
		return err
	}
	targetProperties, err := s.GetProperties(target)
	if err != nil {
		return err
	}

	for _, originItem := range originProperties {
		for _, targetItem := range targetProperties {
			if targetItem.Name == originItem.Name && targetItem.ValueOf.CanSet() {
				targetItem.ValueOf.Set(originItem.ValueOf)
			}
		}
	}
	return nil
}

func (s *StructPropertyUtility) CheckTagKey(origin interface{}, field, tag string) (keyExist, valExist bool, err error) {
	if origin == nil || field == "" || tag == "" {
		return false, false, errors.New("Struct required or Field required or Tag required")
	}

	properties, err := s.GetProperties(origin)
	if err != nil {
		return false, false, err
	}

	var key, val bool
	for _, item := range properties {
		if item.Name == field && strings.HasPrefix(string(item.OtherInfo.Tag), tag+":") {
			key = true
			if item.OtherInfo.Tag.Get(tag) != "" {
				val = true
			}
			return key, val, nil
		}
	}
	return key, val, nil
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
	entityType, _ := s.getElem(entity)
	return entityType.Name()
}

func (s *StructPropertyUtility) isStructType(entity interface{}) bool {
	typeOf, _ := s.getElem(entity)
	if typeOf.Kind() != reflect.Struct {
		return false
	}
	return true
}

func (s *StructPropertyUtility) getElem(entity interface{}) (reflect.Type, reflect.Value) {
	//Elem() 如果取到值非Interface 或 pointer会panic错误，使用Elem()方法转换为源地址的reflect.Value或reflect.Type，才能进行后续操作
	//否则就是指针或接口的Value或Type了
	//而且用了这个必定要传指针或接口类型的参数
	typeOf := reflect.TypeOf(entity).Elem()
	valueOf := reflect.ValueOf(entity).Elem()
	return typeOf, valueOf
}
