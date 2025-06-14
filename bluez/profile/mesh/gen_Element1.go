// Code generated by go-bluetooth generator DO NOT EDIT.

package mesh

import (
	"sync"

	"github.com/godbus/dbus/v5"
	"github.com/bfontana/go-bluetooth/bluez"
	"github.com/bfontana/go-bluetooth/props"
	"github.com/bfontana/go-bluetooth/util"
)

var Element1Interface = "org.bluez.mesh.Element1"

// NewElement1 create a new instance of Element1
//
// Args:
// - servicePath: unique name
// - objectPath: <app_defined_element_path>
func NewElement1(servicePath string, objectPath dbus.ObjectPath) (*Element1, error) {
	a := new(Element1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  servicePath,
			Iface: Element1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	a.Properties = new(Element1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	return a, nil
}

/*
Element1 Mesh Element Hierarchy
*/
type Element1 struct {
	client                 *bluez.Client
	propertiesSignal       chan *dbus.Signal
	objectManagerSignal    chan *dbus.Signal
	objectManager          *bluez.ObjectManager
	Properties             *Element1Properties
	watchPropertiesChannel chan *dbus.Signal
}

// Element1Properties contains the exposed properties of an interface
type Element1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

	/*
		Location Location descriptor as defined in the GATT Bluetooth Namespace
			Descriptors section of the Bluetooth SIG Assigned Numbers
	*/
	Location uint16

	/*
		Models An array of SIG Models:

				id - SIG Model Identifier

				options - a dictionary that may contain additional model
				info. The following keys are defined:
	*/
	Models []ConfigurationItem

	/*
		Publish supports publication mechanism
	*/
	Publish bool

	/*
		Subscribe supports subscription mechanism

			The array may be empty.
	*/
	Subscribe bool

	/*
		VendorModels An array of Vendor Models:

				vendor - a 16-bit Bluetooth-assigned Company ID as
				defined by Bluetooth SIG.

				id - a 16-bit vendor-assigned Model Identifier

				options - a dictionary that may contain additional model
				info. The following keys are defined:
	*/
	VendorModels []VendorOptionsItem
}

// Lock access to properties
func (p *Element1Properties) Lock() {
	p.lock.Lock()
}

// Unlock access to properties
func (p *Element1Properties) Unlock() {
	p.lock.Unlock()
}

// GetLocation get Location value
func (a *Element1) GetLocation() (uint16, error) {
	v, err := a.GetProperty("Location")
	if err != nil {
		return uint16(0), err
	}
	return v.Value().(uint16), nil
}

// GetModels get Models value
func (a *Element1) GetModels() ([]ConfigurationItem, error) {
	v, err := a.GetProperty("Models")
	if err != nil {
		return []ConfigurationItem{}, err
	}
	return v.Value().([]ConfigurationItem), nil
}

// SetPublish set Publish value
func (a *Element1) SetPublish(v bool) error {
	return a.SetProperty("Publish", v)
}

// GetPublish get Publish value
func (a *Element1) GetPublish() (bool, error) {
	v, err := a.GetProperty("Publish")
	if err != nil {
		return false, err
	}
	return v.Value().(bool), nil
}

// SetSubscribe set Subscribe value
func (a *Element1) SetSubscribe(v bool) error {
	return a.SetProperty("Subscribe", v)
}

// GetSubscribe get Subscribe value
func (a *Element1) GetSubscribe() (bool, error) {
	v, err := a.GetProperty("Subscribe")
	if err != nil {
		return false, err
	}
	return v.Value().(bool), nil
}

// GetVendorModels get VendorModels value
func (a *Element1) GetVendorModels() ([]VendorOptionsItem, error) {
	v, err := a.GetProperty("VendorModels")
	if err != nil {
		return []VendorOptionsItem{}, err
	}
	return v.Value().([]VendorOptionsItem), nil
}

// Close the connection
func (a *Element1) Close() {
	a.unregisterPropertiesSignal()
	a.client.Disconnect()
}

// Path return Element1 object path
func (a *Element1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Client return Element1 dbus client
func (a *Element1) Client() *bluez.Client {
	return a.client
}

// Interface return Element1 interface
func (a *Element1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *Element1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

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

// ToMap convert a Element1Properties to map
func (a *Element1Properties) ToMap() (map[string]interface{}, error) {
	return props.ToMap(a), nil
}

// FromMap convert a map to an Element1Properties
func (a *Element1Properties) FromMap(props map[string]interface{}) (*Element1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an Element1Properties
func (a *Element1Properties) FromDBusMap(props map[string]dbus.Variant) (*Element1Properties, error) {
	s := new(Element1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// ToProps return the properties interface
func (a *Element1) ToProps() bluez.Properties {
	return a.Properties
}

// GetWatchPropertiesChannel return the dbus channel to receive properties interface
func (a *Element1) GetWatchPropertiesChannel() chan *dbus.Signal {
	return a.watchPropertiesChannel
}

// SetWatchPropertiesChannel set the dbus channel to receive properties interface
func (a *Element1) SetWatchPropertiesChannel(c chan *dbus.Signal) {
	a.watchPropertiesChannel = c
}

// GetProperties load all available properties
func (a *Element1) GetProperties() (*Element1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *Element1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *Element1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *Element1) GetPropertiesSignal() (chan *dbus.Signal, error) {

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
func (a *Element1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *Element1) WatchProperties() (chan *bluez.PropertyChanged, error) {
	return bluez.WatchProperties(a)
}

func (a *Element1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	return bluez.UnwatchProperties(a, ch)
}

/*
MessageReceived

	This method is called by bluetooth-meshd daemon when a message
	arrives addressed to the application.
	The source parameter is unicast address of the remote
	node-element that sent the message.
	The key_index parameter indicates which application key has been
	used to decode the incoming message. The same key_index should
	be used by the application when sending a response to this
	message (in case a response is expected).
	The destination parameter contains the destination address of
	received message. Underlying variant types are:
	uint16
		Destination is an unicast address, or a well known
		group address
	array{byte}
		Destination is a virtual address label
	The data parameter is the incoming message.
*/
func (a *Element1) MessageReceived(source uint16, key_index uint16, destination dbus.Variant, data []byte) error {
	return a.client.Call("MessageReceived", 0, source, key_index, destination, data).Store()
}

/*
DevKeyMessageReceived

	This method is called by meshd daemon when a message arrives
	addressed to the application, which was sent with the remote
	node's device key.
	The source parameter is unicast address of the remote
	node-element that sent the message.
	The remote parameter if true indicates that the device key
	used to decrypt the message was from the sender. False
	indicates that the local nodes device key was used, and the
	message has permissions to modify local states.
	The net_index parameter indicates what subnet the message was
	received on, and if a response is required, the same subnet
	must be used to send the response.
	The data parameter is the incoming message.
*/
func (a *Element1) DevKeyMessageReceived(source uint16, remote bool, net_index uint16, data []byte) error {
	return a.client.Call("DevKeyMessageReceived", 0, source, remote, net_index, data).Store()
}

/*
UpdateModelConfiguration

	This method is called by bluetooth-meshd daemon when a model's
	configuration is updated.
	The model_id parameter contains BT SIG Model Identifier or, if
	Vendor key is present in config dictionary, a 16-bit
	vendor-assigned Model Identifier.
	The config parameter is a dictionary with the following keys
	defined:
	array{uint16} Bindings
		Indices of application keys bound to the model
	uint32 PublicationPeriod
		Model publication period in milliseconds
	uint16 Vendor
		A 16-bit Bluetooth-assigned Company Identifier of the
		vendor as defined by Bluetooth SIG
	array{variant} Subscriptions
		Addresses the model is subscribed to.
		Each address is provided either as uint16 for group
		addresses, or as array{byte} for virtual labels.
*/
func (a *Element1) UpdateModelConfiguration(model_id uint16, config map[string]interface{}) error {
	return a.client.Call("UpdateModelConfiguration", 0, model_id, config).Store()
}
