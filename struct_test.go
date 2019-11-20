/**
 * @Author: Tomonori
 * @Date: 2019/11/20 11:50
 * @Desc:
 */
package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"testing"
)

func TestCheckTag(test *testing.T) {
	utility := NewStructPropertyUtility()
	test1 := &TestStrcut{Name: "tomo", Age: "18", Gender: 1}
	keyExist, valExist, err := utility.CheckTagKey(test1, "Name", "fuck")
	if err != nil {
		panic(err)
	}
	log.Println("key:", keyExist, " val:", valExist)
}

func TestCopyNotNull(test *testing.T) {
	utility := NewStructPropertyUtility()
	test1 := &TestStrcut{Name: "tomo", Age: "18", Gender: 1}
	test2 := &Test2Struct{}
	err := utility.CopyNotNull(test1, test2)
	if err != nil {
		panic(err)
	}
	log.Println("main target:", test2)
}

func TestStructToMap(test *testing.T) {
	utility := NewStructPropertyUtility()
	test1 := &TestStrcut{Name: "tomo", Age: "18", Gender: 1}

	resultMap, _ := utility.StructToMap(test1)
	resultMap2, _ := utility.StructToMap(test1)
	log.Println(resultMap)
	log.Println(resultMap2)
}

func Test0(test *testing.T) {
	test1 := &TestStrcut{Name: "tomo", Age: "18", Gender: 1}
	valueOf := reflect.ValueOf(*test1)
	typeOf := reflect.TypeOf(*test1)
	fieldSize := valueOf.NumField()
	fmt.Println("typeOf:", typeOf.Kind())

	for i := 0; i < fieldSize; i++ {
		fmt.Println("字段名称:", typeOf.Field(i).Name)
		fmt.Println("字段标签:", typeOf.Field(i).Tag)
		fmt.Println("字段标签信息:", strings.HasPrefix(string(typeOf.Field(i).Tag), "fuck:"))
		fmt.Println("字段的值:", valueOf.Field(i))
		fmt.Println("字段在结构体中的字节偏移量:", typeOf.Field(i).Offset)
		fmt.Println("索引", typeOf.Field(i).Index)
		fmt.Println("字段是否持有值:", valueOf.Field(i).IsValid()) //均有值，应使用IsZero判断
		//fmt.Println("值是否是nil:", valueOf.Field(i).IsNil()) //持有的值是否为nil，值的分类必须是通道、函数、接口、映射、指针、切片之一
		fmt.Println("值是否是零值:", valueOf.Field(i).IsZero())
		fmt.Printf("是否是隐藏字段： %v\n\n", typeOf.Field(i).Anonymous)
	}
}
