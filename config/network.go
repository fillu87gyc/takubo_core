package config

type network struct {
	BackendIP     string
	DynamixelIP   string
	WizWebIP      string
	BackendPort   int
	DynamixelPort int
	WizWebPort    int
}

var ApiVersion = "/v1/"

func Network() network {
	return network{
		BackendIP:     BackendIP,
		DynamixelIP:   DynamixelIP,
		WizWebIP:      WizWebIP,
		BackendPort:   BackendPort,
		DynamixelPort: DynamixelPort,
		WizWebPort:    WizWebPort,
	}
}

const (
	BackendIP     = "0.0.0.0"
	DynamixelIP   = "0.0.0.0"
	WizWebIP      = "0.0.0.0"
	BackendPort   = 3303
	DynamixelPort = 5550
	WizWebPort    = 5551
)
