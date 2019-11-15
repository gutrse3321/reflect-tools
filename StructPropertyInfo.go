/**
 * @Author: Tomonori
 * @Date: 2019/11/14 12:29
 * @Desc:
 */
package main

import "reflect"

type PropertyBase struct {
	Name      string
	Type      string
	OtherInfo reflect.StructField
	ValueOf   reflect.Value
}

type StructPropertyInfo struct {
	//PrimaryKey
}
