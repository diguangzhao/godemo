package main

import (
	"fmt"
	"reflect"
)

/*
* @Author: guangzhao.di
* @Date:   2019-12-19 21:11:46
* @Last Modified by:   guangzhao.di
* @Last Modified time: 2019-12-19 22:06:43
*/

type User struct {
	Id int
	Name string
	Age int
}

func (user User) Hello()  {
	fmt.Println("hello user", user)
}

func main() {
	user := User{1, "张三", 10}
	Info(user)
}
func Info(o interface{})  {
	t := reflect.TypeOf(o)
	fmt.Println("t", t)

	v := reflect.ValueOf(o)
	fmt.Println("value", v)

	fmt.Printf()
	fmt.Println(t.String())



}
