// Code generated DO NOT EDIT

package mesh



import (
   "sync"
   "github.com/muka/go-bluetooth/bluez"
   "github.com/muka/go-bluetooth/util"
   "github.com/muka/go-bluetooth/props"
   "github.com/godbus/dbus"
)

var ProvisionAgent1Interface = "org.bluez.mesh.ProvisionAgent1"


// NewProvisionAgent1 create a new instance of ProvisionAgent1
//
// Args:
// - servicePath: unique name
// - objectPath: freely definable
func NewProvisionAgent1(servicePath string, objectPath dbus.ObjectPath) (*ProvisionAgent1, error) {
	a := new(ProvisionAgent1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  servicePath,
			Iface: ProvisionAgent1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(ProvisionAgent1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}


/*
ProvisionAgent1 Provisioning Agent Hierarchy

*/
type ProvisionAgent1 struct {
	client     				*bluez.Client
	propertiesSignal 	chan *dbus.Signal
	objectManagerSignal chan *dbus.Signal
	objectManager       *bluez.ObjectManager
	Properties 				*ProvisionAgent1Properties
	watchPropertiesChannel chan *dbus.Signal
}

// ProvisionAgent1Properties contains the exposed properties of an interface
type ProvisionAgent1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

	/*
	Capabilities An array of strings with the following allowed values:
			"blink"
			"beep"
			"vibrate"
			"out-numeric"
			"out-alpha"
			"push"
			"twist"
			"in-numeric"
			"in-alpha"
			"static-oob"
			"public-oob"
	*/
	Capabilities []string

	/*
	OutOfBandInfo Indicates availability of OOB data. An array of strings with the
		following allowed values:
			"other"
			"uri"
			"machine-code-2d"
			"bar-code"
			"nfc"
			"number"
			"string"
			"on-box"
			"in-box"
			"on-paper",
			"in-manual"
			"on-device"
	*/
	OutOfBandInfo []string

	/*
	URI Uniform Resource Identifier points to out-of-band (OOB)
		information (e.g., a public key)
	*/
	URI string

}

//Lock access to properties
func (p *ProvisionAgent1Properties) Lock() {
	p.lock.Lock()
}

//Unlock access to properties
func (p *ProvisionAgent1Properties) Unlock() {
	p.lock.Unlock()
}




// SetCapabilities set Capabilities value
func (a *ProvisionAgent1) SetCapabilities(v []string) error {
	return a.SetProperty("Capabilities", v)
}



// GetCapabilities get Capabilities value
func (a *ProvisionAgent1) GetCapabilities() ([]string, error) {
	v, err := a.GetProperty("Capabilities")
	if err != nil {
		return []string{}, err
	}
	return v.Value().([]string), nil
}




// SetOutOfBandInfo set OutOfBandInfo value
func (a *ProvisionAgent1) SetOutOfBandInfo(v []string) error {
	return a.SetProperty("OutOfBandInfo", v)
}



// GetOutOfBandInfo get OutOfBandInfo value
func (a *ProvisionAgent1) GetOutOfBandInfo() ([]string, error) {
	v, err := a.GetProperty("OutOfBandInfo")
	if err != nil {
		return []string{}, err
	}
	return v.Value().([]string), nil
}




// SetURI set URI value
func (a *ProvisionAgent1) SetURI(v string) error {
	return a.SetProperty("URI", v)
}



// GetURI get URI value
func (a *ProvisionAgent1) GetURI() (string, error) {
	v, err := a.GetProperty("URI")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}



// Close the connection
func (a *ProvisionAgent1) Close() {
	
	a.unregisterPropertiesSignal()
	
	a.client.Disconnect()
}

// Path return ProvisionAgent1 object path
func (a *ProvisionAgent1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Client return ProvisionAgent1 dbus client
func (a *ProvisionAgent1) Client() *bluez.Client {
	return a.client
}

// Interface return ProvisionAgent1 interface
func (a *ProvisionAgent1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *ProvisionAgent1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

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


// ToMap convert a ProvisionAgent1Properties to map
func (a *ProvisionAgent1Properties) ToMap() (map[string]interface{}, error) {
	return props.ToMap(a), nil
}

// FromMap convert a map to an ProvisionAgent1Properties
func (a *ProvisionAgent1Properties) FromMap(props map[string]interface{}) (*ProvisionAgent1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an ProvisionAgent1Properties
func (a *ProvisionAgent1Properties) FromDBusMap(props map[string]dbus.Variant) (*ProvisionAgent1Properties, error) {
	s := new(ProvisionAgent1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// ToProps return the properties interface
func (a *ProvisionAgent1) ToProps() bluez.Properties {
	return a.Properties
}

// GetWatchPropertiesChannel return the dbus channel to receive properties interface
func (a *ProvisionAgent1) GetWatchPropertiesChannel() chan *dbus.Signal {
	return a.watchPropertiesChannel
}

// SetWatchPropertiesChannel set the dbus channel to receive properties interface
func (a *ProvisionAgent1) SetWatchPropertiesChannel(c chan *dbus.Signal) {
	a.watchPropertiesChannel = c
}

// GetProperties load all available properties
func (a *ProvisionAgent1) GetProperties() (*ProvisionAgent1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *ProvisionAgent1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *ProvisionAgent1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *ProvisionAgent1) GetPropertiesSignal() (chan *dbus.Signal, error) {

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
func (a *ProvisionAgent1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *ProvisionAgent1) WatchProperties() (chan *bluez.PropertyChanged, error) {
	return bluez.WatchProperties(a)
}

func (a *ProvisionAgent1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	return bluez.UnwatchProperties(a, ch)
}




/*
PrivateKey 
		This method is called during provisioning if the Provisioner
		has requested Out-Of-Band ECC key exchange. The Private key is
		returned to the Daemon, and the Public Key is delivered to the
		remote Provisioner using a method that does not involve the
		Bluetooth Mesh system. The Private Key returned must be 32
		octets in size, or the Provisioning procedure will fail and be
		canceled.

		This function will only be called if the Provisioner has
		requested pre-determined keys to be exchanged Out-of-Band, and
		the local role is Unprovisioned device.


*/
func (a *ProvisionAgent1) PrivateKey() ([]byte, error) {
	
	var val0 []byte
	err := a.client.Call("PrivateKey", 0, ).Store(&val0)
	return val0, err	
}

/*
PublicKey 
		This method is called during provisioning if the local device is
		the Provisioner, and is requestng Out-Of-Band ECC key exchange.
		The Public key is returned to the Daemon that is the matched
		pair of the Private key of the remote device. The Public Key
		returned must be 64 octets in size, or the Provisioning
		procedure will fail and be canceled.

		This function will only be called if the Provisioner has
		requested pre-determined keys to be exchanged Out-of-Band, and
		the local role is Provisioner.


*/
func (a *ProvisionAgent1) PublicKey() ([]byte, error) {
	
	var val0 []byte
	err := a.client.Call("PublicKey", 0, ).Store(&val0)
	return val0, err	
}

/*
DisplayString 
		This method is called when the Daemon has something important
		for the Agent to Display, but does not require any additional
		input locally. For instance: "Enter "ABCDE" on remote device".


*/
func (a *ProvisionAgent1) DisplayString(value string) error {
	
	return a.client.Call("DisplayString", 0, value).Store()
	
}

/*
DisplayNumeric 
		This method is called when the Daemon has something important
		for the Agent to Display, but does not require any additional
		input locally. For instance: "Enter 14939264 on remote device".

		The type parameter indicates the display method. Allowed values
		are:
			"blink" - Locally blink LED
			"beep" - Locally make a noise
			"vibrate" - Locally vibrate
			"out-numeric" - Display value to enter remotely
			"push" - Request pushes on remote button
			"twist" - Request twists on remote knob

		The number parameter is the specific value represented by the
		Prompt.


*/
func (a *ProvisionAgent1) DisplayNumeric(type string, number uint32) error {
	
	return a.client.Call("DisplayNumeric", 0, type, number).Store()
	
}

/*
PromptNumeric 
		This method is called when the Daemon requests the user to
		enter a decimal value between 1-99999999.

		The type parameter indicates the input method. Allowed values
		are:
			"blink" - Enter times remote LED blinked
			"beep" - Enter times remote device beeped
			"vibrate" - Enter times remote device vibrated
			"in-numeric" - Enter remotely displayed value
			"push" - Push local button remotely requested times
			"twist" - Twist local knob remotely requested times


		This agent should prompt the user for specific input. For
		instance: "Enter value being displayed by remote device".


*/
func (a *ProvisionAgent1) PromptNumeric(type string) (uint32, error) {
	
	var val0 uint32
	err := a.client.Call("PromptNumeric", 0, type).Store(&val0)
	return val0, err	
}

/*
PromptStatic 
		This method is called when the Daemon requires a 16 octet byte
		array, as an Out-of-Band authentication.

		The type parameter indicates the input method. Allowed values
		are:
			"static-oob" - return 16 octet array
			"in-alpha" - return 16 octet alpha array

		The Static data returned must be 16 octets in size, or the
		Provisioning procedure will fail and be canceled. If input type
		is "in-alpha", the printable characters should be
		left-justified, with trailing 0x00 octets filling the remaining
		bytes.


*/
func (a *ProvisionAgent1) PromptStatic(type string) ([]byte, error) {
	
	var val0 []byte
	err := a.client.Call("PromptStatic", 0, type).Store(&val0)
	return val0, err	
}

/*
Cancel 
		This method gets called by the daemon to cancel any existing
		Agent Requests. When called, any pending user input should be
		canceled, and any display requests removed.


Properties:
	array{string} Capabilities [read-only]

		An array of strings with the following allowed values:
			"blink"
			"beep"
			"vibrate"
			"out-numeric"
			"out-alpha"
			"push"
			"twist"
			"in-numeric"
			"in-alpha"
			"static-oob"
			"public-oob"

	array{string} OutOfBandInfo [read-only, optional]

		Indicates availability of OOB data. An array of strings with the
		following allowed values:
			"other"
			"uri"
			"machine-code-2d"
			"bar-code"
			"nfc"
			"number"
			"string"
			"on-box"
			"in-box"
			"on-paper",
			"in-manual"
			"on-device"

	string URI [read-only, optional]

		Uniform Resource Identifier points to out-of-band (OOB)
		information (e.g., a public key)

*/
func (a *ProvisionAgent1) Cancel() error {
	
	return a.client.Call("Cancel", 0, ).Store()
	
}

