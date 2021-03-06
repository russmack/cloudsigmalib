# cloudsigmalib

A simple, clean wrapper api for the cloudsigma core project.

[![Build Status](https://travis-ci.org/russmack/cloudsigmalib.svg?branch=master)](https://travis-ci.org/russmack/cloudsigmalib)

---
#### Status: Beta.
---

## Usage
```
	cloud, err := cloudsigmalib.NewCloud("zrh")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := cloud.GetLocations()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Response:", string(resp))

	cloud.BasicAuth = &cloudsigmalib.BasicAuth{"my@email", "mypass"}

	server := cloud.NewServer()
	resp, err = server.Create("Test server 123", 1000, 536870912, "thisispass")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Response:", string(resp))

	servers := cloud.NewServers()
	resp, err = servers.List()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Response:", string(resp))
```

## Example
```
cd examples
go run main.go
```

## Features, so far:

- [X] Servers [List, Create, Delete, Start, Stop, Shutdown]
- [X] Drives [List, Create, Delete]
- [X] VLANs
- [X] IPs
- [X] ACLs
- [X] Tags
- [X] Snapshots [List, list single, create, delete]
- [X] Images [Download]
- [X] Current Usage
- [X] Burst Usage
- [X] Daily Burst Usage
- [X] Locations
- [X] Balance
- [X] Transactions
- [X] Subscriptions
- [X] Profile
- [X] Cloud Status
- [X] Capabilities
- [X] Notification Contacts [list, create]
- [X] Notification Preferences
- [X] Keypairs

## License
BSD 3-Clause: [LICENSE.txt](LICENSE.txt)

[<img alt="LICENSE" src="http://img.shields.io/pypi/l/Django.svg?style=flat-square"/>](LICENSE.txt)
