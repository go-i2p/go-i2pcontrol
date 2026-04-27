package i2pcontrol

// Going to need different fields for each type of service, but for now just a placeholder struct
// maybe an interface? Could help solve the problem of different types of services?

type CommonConfig struct {
	Name        string
	Description string
	AutoStart   bool
	SSL         bool
}

type ServiceConfig struct{}

type ClientConfig struct{}

// ServiceAction performs an action on a tunnel.
func ServiceAction(name, action string, toAll bool) (string, map[string]interface{}, error) {
	retpre, err := Call("TunnelManager", map[string]interface{}{
		"Name":   name,
		"Action": action,
		"All":    toAll,
	})
	if err != nil {
		return "", nil, err
	}
	var tunnelOptions map[string]interface{}
	result := retpre["status"].(string)

	// a get action returns the tunnel options, this is the only action that returns them
	if action == "get" {
		tunnelOptions = retpre["i2p.router.net.tunnels.i2ptunnel.options"].(map[string]interface{})
	}

	return result, tunnelOptions, nil
}

// AddHiddenService creates a new hidden service.
func AddHiddenService(service ServiceConfig) (string, error) {
	return "", nil
}

func AddClientTunnel(client ClientConfig) (string, error) {
	return "", nil
}
