package main

import (
	"fmt"
	"github.com/russmack/cloudsigmalib"
)

func main() {
	cloud, err := cloudsigmalib.NewCloud("zrh")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Using cloud location:", cloud.Location)
	resp, err := cloud.GetLocations()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Response:", string(resp))

	cloud.BasicAuth = &cloudsigmalib.BasicAuth{"my@email", "mypass"}
	resp, err = cloud.GetBalance()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Response:", string(resp))

	servers := cloud.NewServers()
	resp, err = cloud.List(servers)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Response:", string(resp))

}
