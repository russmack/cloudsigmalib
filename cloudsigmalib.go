package cloudsigmalib

import (
	"errors"
	"fmt"
	"github.com/russmack/cloudsigma"
	"strings"
)

type Cloud struct {
	Location  string
	BasicAuth *BasicAuth
}

type BasicAuth struct {
	Username string
	Password string
}

func NewCloud(location string) (*Cloud, error) {
	if strings.TrimSpace(location) == "" {
		return nil, errors.New("Location cannot be empty.")
	}
	c := &Cloud{Location: location}
	return c, nil
}

func (c *Cloud) sendRequest(args *cloudsigma.Args) ([]byte, error) {
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Printf("Error calling client. %s", err)
		return nil, err
	}
	return resp, nil
}

func (c *Cloud) setAuth(args *cloudsigma.Args) *cloudsigma.Args {
	args.Username = c.BasicAuth.Username
	args.Password = c.BasicAuth.Password
	return args
}

func (c *Cloud) GetLocations() ([]byte, error) {
	o := cloudsigma.NewLocations()
	args := o.NewList()
	args.Location = c.Location
	return c.sendRequest(args)
}

func (c *Cloud) GetBalance() ([]byte, error) {
	if c.BasicAuth == nil {
		return nil, errors.New("BasicAuth must not be nil.")
	}
	o := cloudsigma.NewBalance()
	args := o.NewList()
	args = c.setAuth(args)
	args.Location = c.Location
	return c.sendRequest(args)
}

func (c *Cloud) ListServers() ([]byte, error) {
	o := cloudsigma.NewServers()
	args := o.NewList()
	args = c.setAuth(args)
	args.Location = c.Location
	return c.sendRequest(args)
}

func (c *Cloud) CreateServer(name string, cpu int, memory int, vncPassword string) ([]byte, error) {
	o := cloudsigma.NewServers()
	newServers := []cloudsigma.ServerRequest{
		cloudsigma.ServerRequest{name, cpu, memory, vncPassword},
	}
	args := o.NewCreate(newServers)
	args = c.setAuth(args)
	args.Location = c.Location
	return c.sendRequest(args)
}
