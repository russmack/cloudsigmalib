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

	resp, err = cloud.GetCloudStatus()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n\nResponse:", string(resp))

	cloud.BasicAuth = &cloudsigmalib.BasicAuth{login.Username, login.Password}

	resp, err = cloud.GetCapabilities()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n\nResponse:", string(resp))

	resp, err = cloud.GetBalance()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n\nResponse:", string(resp))

	img := cloud.NewImage()
	uuid := "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	resp, err = img.Download(uuid, "MyDownloadedImage.img")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n\nResponse:", string(resp))

	return

	resp, err = cloud.GetSubscriptions()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n\nResponse:", string(resp))

	resp, err = cloud.GetProfile()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n\nResponse:", string(resp))

	resp, err = cloud.GetTransactions()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n\nResponse:", string(resp))

	resp, err = cloud.GetLicenses()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n\nResponse:", string(resp))

	keypairs := cloud.NewKeypairs()
	resp, err = keypairs.List()
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

	acls := cloud.NewACLs()
	resp, err = acls.List()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n\nResponse:", string(resp))

	tags := cloud.NewTags()
	resp, err = tags.List()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n\nResponse:", string(resp))

	/*
		notificationContactsCreator := cloud.NewNotificationContacts()
		resp, err = notificationContactsCreator.Create("my@testemail.com", "MyName Not", "+11235554567")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("\n\nResponse:", string(resp))
	*/

	notificationContacts := cloud.NewNotificationContacts()
	resp, err = notificationContacts.List()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n\nResponse:", string(resp))

	notificationPrefs := cloud.NewNotificationPreferences()
	resp, err = notificationPrefs.List()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n\nResponse:", string(resp))

	dailyBurstUsage := cloud.NewDailyBurstUsage()
	resp, err = dailyBurstUsage.List()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n\nResponse:", string(resp))
}
