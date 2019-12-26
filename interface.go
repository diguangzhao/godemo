package main

import "fmt"

/*
* @Author: guangzhao.di
* @Date:   2019-12-19 21:11:46
* @Last Modified by:   guangzhao.di
* @Last Modified time: 2019-12-19 22:06:43
*/

type USB interface {
	Name() string
	Connecter
}

type Connecter interface {
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (phoneConnecter *PhoneConnecter) Name() string  {
	fmt.Println(phoneConnecter.name)
	return phoneConnecter.name
}

func (phoneConnecter *PhoneConnecter) Connect()  {
	fmt.Println("phoneConnecter connecteds")
}

func (phoneConnecter *PhoneConnecter) Connect2()  {
	fmt.Println("phoneConnecter connecteds2")
}

func main() {
	connecter := &PhoneConnecter{}
	Disconnect(connecter)
}

func Disconnect(usb interface{})  {
	if phoneConnecter, ok := usb.(*PhoneConnecter); ok {
		fmt.Println(phoneConnecter)
		fmt.Println(ok)
	}

}