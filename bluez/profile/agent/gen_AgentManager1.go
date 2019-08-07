package agent



import (
  "sync"
  "github.com/muka/go-bluetooth/bluez"
  "github.com/godbus/dbus"
)

var AgentManager1Interface = "org.bluez.AgentManager1"


// NewAgentManager1 create a new instance of AgentManager1
//
// Args:

func NewAgentManager1() (*AgentManager1, error) {
	a := new(AgentManager1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez",
			Iface: AgentManager1Interface,
			Path:  dbus.ObjectPath("/org/bluez"),
			Bus:   bluez.SystemBus,
		},
	)
	
	return a, nil
}


/*
AgentManager1 Agent Manager hierarchy

*/
type AgentManager1 struct {
	client     				*bluez.Client
	propertiesSignal 	chan *dbus.Signal
	objectManagerSignal chan *dbus.Signal
	objectManager       *bluez.ObjectManager
	Properties 				*AgentManager1Properties
}

// AgentManager1Properties contains the exposed properties of an interface
type AgentManager1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

}

//Lock access to properties
func (p *AgentManager1Properties) Lock() {
	p.lock.Lock()
}

//Unlock access to properties
func (p *AgentManager1Properties) Unlock() {
	p.lock.Unlock()
}



// Close the connection
func (a *AgentManager1) Close() {
	
	a.client.Disconnect()
}

// Path return AgentManager1 object path
func (a *AgentManager1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Interface return AgentManager1 interface
func (a *AgentManager1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *AgentManager1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

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




/*
RegisterAgent 
			This registers an agent handler.

			The object path defines the path of the agent
			that will be called when user input is needed.

			Every application can register its own agent and
			for all actions triggered by that application its
			agent is used.

			It is not required by an application to register
			an agent. If an application does chooses to not
			register an agent, the default agent is used. This
			is on most cases a good idea. Only application
			like a pairing wizard should register their own
			agent.

			An application can only register one agent. Multiple
			agents per application is not supported.

			The capability parameter can have the values
			"DisplayOnly", "DisplayYesNo", "KeyboardOnly",
			"NoInputNoOutput" and "KeyboardDisplay" which
			reflects the input and output capabilities of the
			agent.

			If an empty string is used it will fallback to
			"KeyboardDisplay".

			Possible errors: org.bluez.Error.InvalidArguments
					 org.bluez.Error.AlreadyExists


*/
func (a *AgentManager1) RegisterAgent(agent dbus.ObjectPath, capability string) error {
	
	return a.client.Call("RegisterAgent", 0, agent, capability).Store()
	
}

/*
UnregisterAgent 
			This unregisters the agent that has been previously
			registered. The object path parameter must match the
			same value that has been used on registration.

			Possible errors: org.bluez.Error.DoesNotExist


*/
func (a *AgentManager1) UnregisterAgent(agent dbus.ObjectPath) error {
	
	return a.client.Call("UnregisterAgent", 0, agent).Store()
	
}

/*
RequestDefaultAgent 
			This requests is to make the application agent
			the default agent. The application is required
			to register an agent.

			Special permission might be required to become
			the default agent.

			Possible errors: org.bluez.Error.DoesNotExist



*/
func (a *AgentManager1) RequestDefaultAgent(agent dbus.ObjectPath) error {
	
	return a.client.Call("RequestDefaultAgent", 0, agent).Store()
	
}

