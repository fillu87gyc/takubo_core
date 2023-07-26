package config

type Network struct {
	backendIP     string
	dynamixelIP   string
	wizWebIP      string
	backendPort   string
	dynamixelPort string
	wizWebPort    string
}

var ApiVersion = "/v2/"

func NewNetwork() Network {
	return Network{
		backendIP:     BackendIP,
		dynamixelIP:   DynamixelIP,
		wizWebIP:      WizWebIP,
		backendPort:   BackendPort,
		dynamixelPort: DynamixelPort,
		wizWebPort:    WizWebPort,
	}
}

func (n *Network) BackendAddr() string {
	return "http://" + n.backendIP + ":" + (n.backendPort) + ApiVersion
}

func (n *Network) DynamixelAddr() string {
	return "http://" + n.dynamixelIP + ":" + n.dynamixelPort
}

func (n *Network) WizWebAddr() string {
	return "http://" + n.wizWebIP + ":" + n.wizWebPort
}

const (
	BackendIP     = "localhost"
	DynamixelIP   = "localhost"
	WizWebIP      = "localhost"
	BackendPort   = "3303"
	DynamixelPort = "3333"
	WizWebPort    = "5551"
)
