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
	o := cloudsigma.ServerRequest{Name: name, Cpu: cpu, Memory: memory, VncPassword: vncPassword}
	l := []cloudsigma.ServerRequest{o}
	cs := cloudsigma.NewServers()
	args := cs.NewCreate(l)
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}

func (s *Server) Delete(uuid string) ([]byte, error) {
	cs := cloudsigma.NewServers()
	args := cs.NewDelete(uuid)
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}

func (s *Server) Start(uuid string) ([]byte, error) {
	cs := cloudsigma.NewServers()
	args := cs.NewStart(uuid)
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}

func (s *Server) Stop(uuid string) ([]byte, error) {
	cs := cloudsigma.NewServers()
	args := cs.NewStop(uuid)
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}

func (s *Server) Shutdown(uuid string) ([]byte, error) {
	cs := cloudsigma.NewServers()
	args := cs.NewShutdown(uuid)
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}

type Drives struct{}
type Drive struct{}

func (c *Cloud) NewDrives() *Drives {
	return &Drives{}
}

func (c *Cloud) NewDrive() *Drive {
	return &Drive{}
}

func (d *Drives) List() ([]byte, error) {
	cs := cloudsigma.NewDrives()
	args := cs.NewList()
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}

func (s *Drive) Create(name string, size int, media string) ([]byte, error) {
	o := cloudsigma.DriveRequest{Media: media, Name: name, Size: size}
	l := []cloudsigma.DriveRequest{o}
	cs := cloudsigma.NewDrives()
	args := cs.NewCreate(l)
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}

func (s *Drive) Delete(uuid string) ([]byte, error) {
	cs := cloudsigma.NewDrives()
	args := cs.NewDelete(uuid)
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}

type CurrentUsage struct{}

func (c *Cloud) NewCurrentUsage() *CurrentUsage {
	return &CurrentUsage{}
}

func (u *CurrentUsage) List() ([]byte, error) {
	cs := cloudsigma.NewCurrentUsage()
	args := cs.NewList()
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}

type BurstUsage struct{}

func (c *Cloud) NewBurstUsage() *BurstUsage {
	return &BurstUsage{}
}

func (u *BurstUsage) List() ([]byte, error) {
	cs := cloudsigma.NewBurstUsage()
	args := cs.NewList()
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}

type VLANs struct{}

func (c *Cloud) NewVLANs() *VLANs {
	return &VLANs{}
}

func (v *VLANs) List() ([]byte, error) {
	cs := cloudsigma.NewVlans()
	args := cs.NewList()
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}

type IPs struct{}

func (c *Cloud) NewIPs() *IPs {
	return &IPs{}
}

func (i *IPs) List() ([]byte, error) {
	cs := cloudsigma.NewIps()
	args := cs.NewList()
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}
