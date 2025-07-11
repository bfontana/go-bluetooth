// Code generated by go-bluetooth generator DO NOT EDIT.

package advertisement_monitor

import (
	"sync"

	"github.com/godbus/dbus/v5"
	"github.com/bfontana/go-bluetooth/bluez"
	"github.com/bfontana/go-bluetooth/props"
	"github.com/bfontana/go-bluetooth/util"
)

var AdvertisementMonitor1Interface = "org.bluez.AdvertisementMonitor1"

// NewAdvertisementMonitor1 create a new instance of AdvertisementMonitor1
//
// Args:
// - objectPath: freely definable
func NewAdvertisementMonitor1(objectPath dbus.ObjectPath) (*AdvertisementMonitor1, error) {
	a := new(AdvertisementMonitor1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez",
			Iface: AdvertisementMonitor1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	a.Properties = new(AdvertisementMonitor1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	return a, nil
}

/*
AdvertisementMonitor1 Advertisement Monitor hierarchy
*/
type AdvertisementMonitor1 struct {
	client                 *bluez.Client
	propertiesSignal       chan *dbus.Signal
	objectManagerSignal    chan *dbus.Signal
	objectManager          *bluez.ObjectManager
	Properties             *AdvertisementMonitor1Properties
	watchPropertiesChannel chan *dbus.Signal
}

// AdvertisementMonitor1Properties contains the exposed properties of an interface
type AdvertisementMonitor1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

	/*
		Patterns If the Type property is set to "or_patterns", then this
				property must exist and have at least one entry in the
				array.

				The structure of a pattern contains the following:
				uint8 start_position
					The index in an AD data field where the search
					should start. The beginning of an AD data field
					is index 0.
				uint8 AD_data_type
					See https://www.bluetooth.com/specifications/
					assigned-numbers/generic-access-profile/ for
					the possible allowed value.
	*/
	Patterns []Pattern

	/*
		RSSIHighThreshold Used in conjunction with RSSIHighTimeout to determine
				whether a device becomes in-range. Valid range is
				-127 to 20 (dBm), while 127 indicates unset.
	*/
	RSSIHighThreshold int16

	/*
		RSSIHighTimeout The time it takes to consider a device as in-range.
				If this many seconds elapses while we continuously
				receive signals at least as strong as RSSIHighThreshold,
				a currently out-of-range device will be considered as
				in-range (found). Valid range is 1 to 300 (seconds),
				while 0 indicates unset.
	*/
	RSSIHighTimeout uint16

	/*
		RSSILowThreshold Used in conjunction with RSSILowTimeout to determine
				whether a device becomes out-of-range. Valid range is
				-127 to 20 (dBm), while 127 indicates unset.
	*/
	RSSILowThreshold int16

	/*
		RSSILowTimeout The time it takes to consider a device as out-of-range.
				If this many seconds elapses without receiving any
				signal at least as strong as RSSILowThreshold, a
				currently in-range device will be considered as
				out-of-range (lost). Valid range is 1 to 300 (seconds),
				while 0 indicates unset.
	*/
	RSSILowTimeout uint16

	/*
		RSSISamplingPeriod Grouping rules on how to propagate the received
				advertisement packets to the client. Valid range is 0 to
				255 while 256 indicates unset.

				The meaning of this property is as follows:
				0:
					All advertisement packets from in-range devices
					would be propagated.
				255:
					Only the first advertisement packet of in-range
					devices would be propagated. If the device
					becomes lost, then the first packet when it is
					found again will also be propagated.
				1 to 254:
					Advertisement packets would be grouped into
					100ms * N time period. Packets in the same group
					will only be reported once, with the RSSI value
					being averaged out.

				Currently this is unimplemented in user space, so the
				value is only used to be forwarded to the kernel.
	*/
	RSSISamplingPeriod uint16

	/*
		Type The type of the monitor. See SupportedMonitorTypes in
				org.bluez.AdvertisementMonitorManager1 for the available
				options.
	*/
	Type string
}

// Lock access to properties
func (p *AdvertisementMonitor1Properties) Lock() {
	p.lock.Lock()
}

// Unlock access to properties
func (p *AdvertisementMonitor1Properties) Unlock() {
	p.lock.Unlock()
}

// GetPatterns get Patterns value
func (a *AdvertisementMonitor1) GetPatterns() ([]Pattern, error) {
	v, err := a.GetProperty("Patterns")
	if err != nil {
		return []Pattern{}, err
	}
	return v.Value().([]Pattern), nil
}

// GetRSSIHighThreshold get RSSIHighThreshold value
func (a *AdvertisementMonitor1) GetRSSIHighThreshold() (int16, error) {
	v, err := a.GetProperty("RSSIHighThreshold")
	if err != nil {
		return int16(0), err
	}
	return v.Value().(int16), nil
}

// GetRSSIHighTimeout get RSSIHighTimeout value
func (a *AdvertisementMonitor1) GetRSSIHighTimeout() (uint16, error) {
	v, err := a.GetProperty("RSSIHighTimeout")
	if err != nil {
		return uint16(0), err
	}
	return v.Value().(uint16), nil
}

// GetRSSILowThreshold get RSSILowThreshold value
func (a *AdvertisementMonitor1) GetRSSILowThreshold() (int16, error) {
	v, err := a.GetProperty("RSSILowThreshold")
	if err != nil {
		return int16(0), err
	}
	return v.Value().(int16), nil
}

// GetRSSILowTimeout get RSSILowTimeout value
func (a *AdvertisementMonitor1) GetRSSILowTimeout() (uint16, error) {
	v, err := a.GetProperty("RSSILowTimeout")
	if err != nil {
		return uint16(0), err
	}
	return v.Value().(uint16), nil
}

// GetRSSISamplingPeriod get RSSISamplingPeriod value
func (a *AdvertisementMonitor1) GetRSSISamplingPeriod() (uint16, error) {
	v, err := a.GetProperty("RSSISamplingPeriod")
	if err != nil {
		return uint16(0), err
	}
	return v.Value().(uint16), nil
}

// GetType get Type value
func (a *AdvertisementMonitor1) GetType() (string, error) {
	v, err := a.GetProperty("Type")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}

// Close the connection
func (a *AdvertisementMonitor1) Close() {
	a.unregisterPropertiesSignal()
	a.client.Disconnect()
}

// Path return AdvertisementMonitor1 object path
func (a *AdvertisementMonitor1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Client return AdvertisementMonitor1 dbus client
func (a *AdvertisementMonitor1) Client() *bluez.Client {
	return a.client
}

// Interface return AdvertisementMonitor1 interface
func (a *AdvertisementMonitor1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *AdvertisementMonitor1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

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

// ToMap convert a AdvertisementMonitor1Properties to map
func (a *AdvertisementMonitor1Properties) ToMap() (map[string]interface{}, error) {
	return props.ToMap(a), nil
}

// FromMap convert a map to an AdvertisementMonitor1Properties
func (a *AdvertisementMonitor1Properties) FromMap(props map[string]interface{}) (*AdvertisementMonitor1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an AdvertisementMonitor1Properties
func (a *AdvertisementMonitor1Properties) FromDBusMap(props map[string]dbus.Variant) (*AdvertisementMonitor1Properties, error) {
	s := new(AdvertisementMonitor1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// ToProps return the properties interface
func (a *AdvertisementMonitor1) ToProps() bluez.Properties {
	return a.Properties
}

// GetWatchPropertiesChannel return the dbus channel to receive properties interface
func (a *AdvertisementMonitor1) GetWatchPropertiesChannel() chan *dbus.Signal {
	return a.watchPropertiesChannel
}

// SetWatchPropertiesChannel set the dbus channel to receive properties interface
func (a *AdvertisementMonitor1) SetWatchPropertiesChannel(c chan *dbus.Signal) {
	a.watchPropertiesChannel = c
}

// GetProperties load all available properties
func (a *AdvertisementMonitor1) GetProperties() (*AdvertisementMonitor1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *AdvertisementMonitor1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *AdvertisementMonitor1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *AdvertisementMonitor1) GetPropertiesSignal() (chan *dbus.Signal, error) {

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
func (a *AdvertisementMonitor1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *AdvertisementMonitor1) WatchProperties() (chan *bluez.PropertyChanged, error) {
	return bluez.WatchProperties(a)
}

func (a *AdvertisementMonitor1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	return bluez.UnwatchProperties(a, ch)
}

/*
Release

	This gets called as a signal for a client to perform
	clean-up when (1)a monitor cannot be activated after it
	was exposed or (2)a monitor has been deactivated.
*/
func (a *AdvertisementMonitor1) Release() error {
	return a.client.Call("Release", 0).Store()
}

/*
Activate

	After a monitor was exposed, this gets called as a
	signal for client to get acknowledged when a monitor
	has been activated, so the client can expect to receive
	calls on DeviceFound() or DeviceLost().
*/
func (a *AdvertisementMonitor1) Activate() error {
	return a.client.Call("Activate", 0).Store()
}

/*
DeviceFound

	This gets called to notify the client of finding the
	targeted device. Once receiving the call, the client
	should start to monitor the corresponding device to
	retrieve the changes on RSSI and advertisement content.
*/
func (a *AdvertisementMonitor1) DeviceFound(device dbus.ObjectPath) error {
	return a.client.Call("DeviceFound", 0, device).Store()
}

/*
DeviceLost

	This gets called to notify the client of losing the
	targeted device. Once receiving this call, the client
	should stop monitoring the corresponding device.
*/
func (a *AdvertisementMonitor1) DeviceLost(device dbus.ObjectPath) error {
	return a.client.Call("DeviceLost", 0, device).Store()
}
