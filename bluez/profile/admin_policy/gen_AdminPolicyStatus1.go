// Code generated by go-bluetooth generator DO NOT EDIT.

package admin_policy

import (
	"sync"

	"github.com/godbus/dbus/v5"
	"github.com/bfontana/go-bluetooth/bluez"
	"github.com/bfontana/go-bluetooth/props"
	"github.com/bfontana/go-bluetooth/util"
)

var AdminPolicyStatus1Interface = "org.bluez.AdminPolicyStatus1"

// NewAdminPolicyStatus1 create a new instance of AdminPolicyStatus1
//
// Args:
// - objectPath: [variable prefix]/{hci0,hci1,...}
func NewAdminPolicyStatus1(objectPath dbus.ObjectPath) (*AdminPolicyStatus1, error) {
	a := new(AdminPolicyStatus1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez",
			Iface: AdminPolicyStatus1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	a.Properties = new(AdminPolicyStatus1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	return a, nil
}

/*
AdminPolicyStatus1 Admin Policy Status hierarchy
*/
type AdminPolicyStatus1 struct {
	client                 *bluez.Client
	propertiesSignal       chan *dbus.Signal
	objectManagerSignal    chan *dbus.Signal
	objectManager          *bluez.ObjectManager
	Properties             *AdminPolicyStatus1Properties
	watchPropertiesChannel chan *dbus.Signal
}

// AdminPolicyStatus1Properties contains the exposed properties of an interface
type AdminPolicyStatus1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

	/*
		ServiceAllowList Current value of service allow list.
	*/
	ServiceAllowList []string
}

// Lock access to properties
func (p *AdminPolicyStatus1Properties) Lock() {
	p.lock.Lock()
}

// Unlock access to properties
func (p *AdminPolicyStatus1Properties) Unlock() {
	p.lock.Unlock()
}

// SetServiceAllowList set ServiceAllowList value
func (a *AdminPolicyStatus1) SetServiceAllowList(v []string) error {
	return a.SetProperty("ServiceAllowList", v)
}

// GetServiceAllowList get ServiceAllowList value
func (a *AdminPolicyStatus1) GetServiceAllowList() ([]string, error) {
	v, err := a.GetProperty("ServiceAllowList")
	if err != nil {
		return []string{}, err
	}
	return v.Value().([]string), nil
}

// Close the connection
func (a *AdminPolicyStatus1) Close() {
	a.unregisterPropertiesSignal()
	a.client.Disconnect()
}

// Path return AdminPolicyStatus1 object path
func (a *AdminPolicyStatus1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Client return AdminPolicyStatus1 dbus client
func (a *AdminPolicyStatus1) Client() *bluez.Client {
	return a.client
}

// Interface return AdminPolicyStatus1 interface
func (a *AdminPolicyStatus1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *AdminPolicyStatus1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

	if a.objectManagerSignal == nil {
		if a.objectManager == nil {
			om, err := bluez.GetObjectManager()
			if err != nil {
				return nil, nil, err
			}
			a.objectManager = om
		}

		s, err := a.objectManager.Register()
		if err != nil {
			return nil, nil, err
		}
		a.objectManagerSignal = s
	}

	cancel := func() {
		if a.objectManagerSignal == nil {
			return
		}
		a.objectManagerSignal <- nil
		a.objectManager.Unregister(a.objectManagerSignal)
		a.objectManagerSignal = nil
	}

	return a.objectManagerSignal, cancel, nil
}

// ToMap convert a AdminPolicyStatus1Properties to map
func (a *AdminPolicyStatus1Properties) ToMap() (map[string]interface{}, error) {
	return props.ToMap(a), nil
}

// FromMap convert a map to an AdminPolicyStatus1Properties
func (a *AdminPolicyStatus1Properties) FromMap(props map[string]interface{}) (*AdminPolicyStatus1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an AdminPolicyStatus1Properties
func (a *AdminPolicyStatus1Properties) FromDBusMap(props map[string]dbus.Variant) (*AdminPolicyStatus1Properties, error) {
	s := new(AdminPolicyStatus1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// ToProps return the properties interface
func (a *AdminPolicyStatus1) ToProps() bluez.Properties {
	return a.Properties
}

// GetWatchPropertiesChannel return the dbus channel to receive properties interface
func (a *AdminPolicyStatus1) GetWatchPropertiesChannel() chan *dbus.Signal {
	return a.watchPropertiesChannel
}

// SetWatchPropertiesChannel set the dbus channel to receive properties interface
func (a *AdminPolicyStatus1) SetWatchPropertiesChannel(c chan *dbus.Signal) {
	a.watchPropertiesChannel = c
}

// GetProperties load all available properties
func (a *AdminPolicyStatus1) GetProperties() (*AdminPolicyStatus1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *AdminPolicyStatus1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *AdminPolicyStatus1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *AdminPolicyStatus1) GetPropertiesSignal() (chan *dbus.Signal, error) {

	if a.propertiesSignal == nil {
		s, err := a.client.Register(a.client.Config.Path, bluez.PropertiesInterface)
		if err != nil {
			return nil, err
		}
		a.propertiesSignal = s
	}

	return a.propertiesSignal, nil
}

// Unregister for changes signalling
func (a *AdminPolicyStatus1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *AdminPolicyStatus1) WatchProperties() (chan *bluez.PropertyChanged, error) {
	return bluez.WatchProperties(a)
}

func (a *AdminPolicyStatus1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	return bluez.UnwatchProperties(a, ch)
}
