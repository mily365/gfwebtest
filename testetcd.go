package main

import "xpass/app/workflow/regdisc"

func main() {
	ser, _ := regdisc.NewServiceReg([]string{"192.168.4.118:2379"}, 5)
	err := ser.PutService("/node/001", "192.168.4.118")
	if err != nil {
		panic(err.Error())
	}
	//select {}

}
