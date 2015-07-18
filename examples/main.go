package main

import (
	"fmt"
	"github.com/russmack/cloudsigmalib"
	"os"
)

func main() {
	config := cloudsigmalib.NewConfig()
	_, err := config.Load()
	if err != nil {
		fmt.Println("Unable to load config.", err)
		os.Exit(1)
	}
	login := config.Login()

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
	fmt.Println("\n\nResponse:", string(resp))

	cloud.BasicAuth = &cloudsigmalib.BasicAuth{login.Username, login.Password}
	resp, err = cloud.GetBalance()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n\nResponse:", string(resp))

	servers := cloud.NewServers()
	resp, err = servers.List()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n\nResponse:", string(resp))

	/*
		server := cloud.NewServer()
		resp, err = server.Create("Test server 123", 1000, 536870912, "thisispass")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("\n\nResponse:", string(resp))
	*/

	drives := cloud.NewDrives()
	resp, err = drives.List()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n\nResponse:", string(resp))

	currUsage := cloud.NewCurrentUsage()
	resp, err = currUsage.List()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n\nResponse:", string(resp))

	burstUsage := cloud.NewBurstUsage()
	resp, err = burstUsage.List()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n\nResponse:", string(resp))

	vlans := cloud.NewVLANs()
	resp, err = vlans.List()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n\nResponse:", string(resp))

	ips := cloud.NewIPs()
	resp, err = ips.List()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n\nResponse:", string(resp))
}
