package main

import "xpass/app/workflow/regdisc"

func main() {
	cli, _ := regdisc.NewClientDis([]string{"192.168.4.118:2379"})
	cli.GetService("/node")
	select {}
}
