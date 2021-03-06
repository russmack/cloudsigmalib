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

func (c *Cloud) sendDownloadRequest(args *cloudsigma.Args) ([]byte, error) {
	client := &cloudsigma.Client{}
	resp, err := client.Download(nil, args)
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

func (c *Cloud) GetCloudStatus() ([]byte, error) {
	o := cloudsigma.NewCloudStatus()
	args := o.NewList()
	args, err := c.setArgs(args)
	if err != nil {
		return nil, err
	}
	return c.sendRequest(args)
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

func (c *Cloud) GetCapabilities() ([]byte, error) {
	o := cloudsigma.NewCapabilities()
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

func (c *Cloud) GetTransactions() ([]byte, error) {
	o := cloudsigma.NewTransactions()
	args := o.NewList()
	args, err := c.setArgs(args)
	if err != nil {
		return nil, err
	}
	return c.sendRequest(args)
}

func (c *Cloud) GetSubscriptions() ([]byte, error) {
	o := cloudsigma.NewSubscriptions()
	args := o.NewList()
	args, err := c.setArgs(args)
	if err != nil {
		return nil, err
	}
	return c.sendRequest(args)
}

func (c *Cloud) GetLicenses() ([]byte, error) {
	o := cloudsigma.NewLicenses()
	args := o.NewList()
	args, err := c.setArgs(args)
	if err != nil {
		return nil, err
	}
	return c.sendRequest(args)
}

func (c *Cloud) GetProfile() ([]byte, error) {
	o := cloudsigma.NewProfile()
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

type Image struct{}

func (c *Cloud) NewImage() *Image {
	return &Image{}
}

func (o *Image) Download(uuid string, filename string) ([]byte, error) {
	cs := cloudsigma.NewImages()
	args := cs.NewDownload(uuid, filename)
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendDownloadRequest(args)
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

type DailyBurstUsage struct{}

func (c *Cloud) NewDailyBurstUsage() *DailyBurstUsage {
	return &DailyBurstUsage{}
}

func (u *DailyBurstUsage) List() ([]byte, error) {
	cs := cloudsigma.NewDailyBurstUsage()
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

type ACLs struct{}

func (c *Cloud) NewACLs() *ACLs {
	return &ACLs{}
}

func (a *ACLs) List() ([]byte, error) {
	cs := cloudsigma.NewAcls()
	args := cs.NewList()
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}

type Tags struct{}

func (c *Cloud) NewTags() *Tags {
	return &Tags{}
}

func (a *Tags) List() ([]byte, error) {
	cs := cloudsigma.NewTags()
	args := cs.NewList()
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}

type NotificationContacts struct{}

func (c *Cloud) NewNotificationContacts() *NotificationContacts {
	return &NotificationContacts{}
}

func (a *NotificationContacts) List() ([]byte, error) {
	cs := cloudsigma.NewNotificationContacts()
	args := cs.NewList()
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}

func (a *NotificationContacts) Create(email string, name string, phone string) ([]byte, error) {
	o := []cloudsigma.ContactRequest{{Email: email, Name: name, Phone: phone}}
	cs := cloudsigma.NewNotificationContacts()
	args := cs.NewCreate(o)
	args, err := cloud.setArgs(args)
	if err != nil {
		e := errors.New(err.Error() + " Ensure phone number starts with a +.")
		return nil, e
	}
	return cloud.sendRequest(args)
}

type NotificationPreferences struct{}

func (c *Cloud) NewNotificationPreferences() *NotificationPreferences {
	return &NotificationPreferences{}
}

func (a *NotificationPreferences) List() ([]byte, error) {
	cs := cloudsigma.NewNotificationPreferences()
	args := cs.NewList()
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}

type Snapshots struct{}
type Snapshot struct{}

func (c *Cloud) NewSnapshots() *Snapshots {
	return &Snapshots{}
}

func (c *Cloud) NewSnapshot() *Snapshot {
	return &Snapshot{}
}

func (s *Snapshots) List() ([]byte, error) {
	cs := cloudsigma.NewSnapshots()
	args := cs.NewList()
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}

func (s *Snapshot) List(uuid string) ([]byte, error) {
	cs := cloudsigma.NewSnapshots()
	args := cs.NewGet(uuid)
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}

func (s *Snapshot) Create(driveUuid string, name string) ([]byte, error) {
	o := cloudsigma.SnapshotRequest{Drive: driveUuid, Name: name}
	cs := cloudsigma.NewSnapshots()
	args := cs.NewCreate(o)
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}

func (s *Snapshot) Delete(uuid string) ([]byte, error) {
	cs := cloudsigma.NewSnapshots()
	args := cs.NewDelete(uuid)
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}

type Keypairs struct{}

func (c *Cloud) NewKeypairs() *Keypairs {
	return &Keypairs{}
}

func (s *Keypairs) List() ([]byte, error) {
	cs := cloudsigma.NewKeypairs()
	args := cs.NewList()
	args, err := cloud.setArgs(args)
	if err != nil {
		return nil, err
	}
	return cloud.sendRequest(args)
}
