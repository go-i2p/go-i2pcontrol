package i2pcontrol

// Going to need different fields for each type of service, but for now just a placeholder struct
// maybe an interface? Could help solve the problem of different types of services?

// TODO: Once we get everything working on the golang end, I need to ensure the
// 	JSON RPC is foolproof. I need to make sure that the API is consistent and that the error handling is robust.

type TunnelOptions struct {
	Length   int
	Variance int // -2 to 2
	Quantity int
	Backup   int
	Profile  string // (interactive or default)

	// possibly another struct for the following fields?
	ReduceIdle   bool
	ReducedCount int
	IdleTime     int
	SigType      string
	EncType      string
}

type CommonConfig struct {
	Name           string
	Description    string
	AutoStart      bool
	SSL            bool
	Port           int
	Type           string
	CustomOptions  string
	PrivateKeyFile string
	TunnelSettings TunnelOptions
}

type ServiceConfig struct {
	CommonSettings CommonConfig
	ReachableBy    string
	TargetHost     string // ONLY FOR STREAMR TUNNELS -- need to test this some more
}

type ClientConfig struct{}

// TODO: When a function is stopped, for some services we may need to set destination and b32 dest to null

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
	retpre, err := Call("TunnelManager", map[string]interface{}{
		"Action":      "create",
		"Name":        service.CommonSettings.Name,
		"Description": service.CommonSettings.Description,
		"StartOnLoad": service.CommonSettings.AutoStart,
		"UseSSL":      service.CommonSettings.SSL,
		"Port":        service.CommonSettings.Port,
		"Type":        service.CommonSettings.Type,
		//"CustomOptions": service.CommonSettings.CustomOptions,
		//"PrivKeyFile":   service.CommonSettings.PrivateKeyFile, // NOT AVAILABLE ON HTTP - GOING TO NEED TO FIND A GOOD WAY
		"ReachableBy": service.ReachableBy,

		"TunnelLength":         service.CommonSettings.TunnelSettings.Length,
		"TunnelVariance":       service.CommonSettings.TunnelSettings.Variance,
		"TunnelQuantity":       service.CommonSettings.TunnelSettings.Quantity,
		"TunnelBackupQuantity": service.CommonSettings.TunnelSettings.Backup,
		"Profile":              service.CommonSettings.TunnelSettings.Profile,
		//"Reduce":               service.CommonSettings.TunnelSettings.ReduceIdle,
		"ReduceCount": service.CommonSettings.TunnelSettings.ReducedCount,
		"ReducedTime": service.CommonSettings.TunnelSettings.IdleTime,
		"SigType":     service.CommonSettings.TunnelSettings.SigType,
		"EncType":     service.CommonSettings.TunnelSettings.EncType,
	})
	if err != nil {
		return "", err
	}
	result := retpre["status"].(string)

	return result, nil
}

func AddClientTunnel(client ClientConfig) (string, error) {
	return "", nil
}
