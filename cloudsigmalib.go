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

var (
	cloud *Cloud
)

func NewCloud(location string) (*Cloud, error) {
	if strings.TrimSpace(location) == "" {
		return nil, errors.New("Location cannot be empty.")
	}
	cloud = &Cloud{Location: location}
	return cloud, nil
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

func (c *Cloud) setArgs(args *cloudsigma.Args) (*cloudsigma.Args, error) {
	if args.RequiresAuth && c.BasicAuth == nil {
		return nil, errors.New("BasicAuth must not be nil.")
	}
	args = c.setAuth(args)
	args.Location = c.Location
	return args, nil
}

func (c *Cloud) setAuth(args *cloudsigma.Args) *cloudsigma.Args {
	if c.BasicAuth == nil {
		return args
	}
	args.Username = c.BasicAuth.Username
	args.Password = c.BasicAuth.Password
	return args
}

func (c *Cloud) GetLocations() ([]byte, error) {
	o := cloudsigma.NewLocations()
	args := o.NewList()
	args, err := c.setArgs(args)
	if err != nil {
		return nil, err
	}
	return c.sendRequest(args)
}

func (c *Cloud) GetBalance() ([]byte, error) {
	o := cloudsigma.NewBalance()
	args := o.NewList()
	args, err := c.setArgs(args)
	if err != nil {
		return nil, err
	}
	return c.sendRequest(args)
}

type Servers struct{}

type Server struct{}

func (c *Cloud) NewServers() *Servers {
	s := &Servers{}
	return s
}

func (c *Cloud) NewServer() *Server {
	s := &Server{}
	return s
}

func (s *Servers) List() ([]byte, error) {
	cs := cloudsigma.NewServers()
	args := cs.NewList()
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}

func (s *Server) Create(name string, cpu int, memory int, vncPassword string) ([]byte, error) {
	ns := cloudsigma.ServerRequest{Name: name, Cpu: cpu, Memory: memory, VncPassword: vncPassword}
	nsList := []cloudsigma.ServerRequest{ns}
	cs := cloudsigma.NewServers()
	args := cs.NewCreate(nsList)
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}
