/**
 * @Author: Tomonori
 * @Date: 2019/11/14 12:29
 * @Desc:
 */
package main

import "reflect"

type PropertyBase struct {
	Type  reflect.Type
	Value reflect.Value
}

type StructPropertyInfo struct {
	//PrimaryKey
}
