package i2pcontrol

// Going to need different fields for each type of service, but for now just a placeholder struct
// maybe an interface? Could help solve the problem of different types of services?

// TODO: Once we get everything working on the golang end, I need to ensure the
// 	JSON RPC is foolproof. I need to make sure that the API is consistent and that the error handling is robust.

type TunnelLengthOptions struct {
	Length   int
	Variance int
}

type TunnelQuantityOptions struct {
	Quantity int
	Backup   int
}

type TunnelCryptographyOptions struct {
	SigType string
	EncType string
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
	ReduceIdle     bool
	ReducedCount   int
	IdleTime       int
	TunnelLength   *TunnelLengthOptions
	TunnelQuantity *TunnelQuantityOptions
	TunnelCrypto   *TunnelCryptographyOptions
}

type SOCKSConfig struct {
	OutProxyType      string
	OutProxies        string // HTTP, SOCKS, CONNECT
	UseOutProxyPlugin bool
	Auth              *Authentication
}

type CONNECTConfig struct {
	OutProxies        string // HTTP, SOCKS, CONNECT
	UseOutProxyPlugin bool
	Auth              *Authentication
}

type HTTPFiltering struct {
	SpoofUserAgent     bool
	BlockAcceptHeaders bool
	BlockReferers      bool
	AllowSSLI2P        bool
}

type HTTPConfig struct {
	Filtering         *HTTPFiltering
	Auth              *Authentication
	OutProxies        string // HTTP, SOCKS, CONNECT
	UseOutProxyPlugin bool
	SSLOutProxies     string
	JumpURLs          string
}

type STREAMRConfig struct {
	TargetHost        string
	TargetDestination string
}

type IRCConfig struct {
	EnableDCC bool
}

type ProfileConfig struct {
	Profile      string // (interactive or default)
	DelayConnect bool
}

type Authentication struct {
	RequireLocalAuth     bool
	LocalAuthUsername    string
	LocalAuthPassword    string
	RequireOutProxyAuth  bool
	OutProxyAuthUsername string
	OutProxyAuthPassword string
}

type ClientConfig struct {
	CommonSettings CommonConfig
	ReachableBy    string
	SharedClient   bool
	DelayOpen      bool

	TunnelDestination string

	// close when idle section - maybe make struct later
	CloseWhenIdle bool
	IdlePeriod    int
	NewKeysOnOpen bool

	// persistent private keys section

	GenerateKeys bool // when this is true, NewKeysOnOpen becomes null / ignored

	ProfileConfig *ProfileConfig
	CONNECTConfig *CONNECTConfig
	HTTPConfig    *HTTPConfig
	SOCKSConfig   *SOCKSConfig
	STREAMRConfig *STREAMRConfig
	IRCConfig     *IRCConfig
}

type LeaseSetConfig struct {
	LeaseSetEnc      string
	LeaseSetPassword string
}

type ServerThrottling struct {
	MaxConcurrentConnections int
	PerClientMinuteLimit     int
	PerClientHourLimit       int
	PerClientDayLimit        int

	TotalPerMinuteLimit int
	TotalPerHourLimit   int
	TotalPerDayLimit    int
}

type STREAMRServerConfig struct { // & also no SSL?
	ReachableBy string
}

type PostLimitConfig struct {
	ClientPerPeriod   int
	ClientBanDuration int
	TotalPerPeriod    int
	TotalBanDuration  int
	POSTLimit         int
}

type TunnelAccessControlConfig struct {
	AccessList            string
	AccessListDescription string
	AccessListFile        string
	UniquePerClient       bool
	MultiHoming           bool
}

type HTTPServerConfig struct {
	WebHost              string
	POSTLimitConfig      PostLimitConfig
	BlockReferers        bool
	BlockUserAgents      bool
	UserAgents           string
	BlockAccessInProxies bool
}

type ServiceConfig struct {
	CommonSettings            CommonConfig
	LeaseSetConfig            LeaseSetConfig
	ServerThrottling          ServerThrottling
	TunnelAccessControlConfig TunnelAccessControlConfig
	PrivateFileKey            string
	Host                      string
	LeaseSetEnc               string
	Profile                   string // (Used for everything besides STREAMR)
	STREAMRServerConfig       *STREAMRServerConfig
	HTTPServerConfig          *HTTPServerConfig
}

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

// AddClientTunnel creates a new hidden service.
func AddClientTunnel(client ClientConfig) (string, error) {
	params := map[string]interface{}{
		"Action":      "create",
		"Name":        client.CommonSettings.Name,
		"StartOnLoad": client.CommonSettings.AutoStart,
		"UseSSL":      client.CommonSettings.SSL,
		"Port":        client.CommonSettings.Port,
		"Type":        client.CommonSettings.Type,
		"Shared":      client.SharedClient,
		"DelayOpen":   client.DelayOpen,
		"Close":       client.CloseWhenIdle,
	}

	addString := func(key, value string) {
		if value != "" {
			params[key] = value
		}
	}

	addInt := func(key string, value int) {
		if value != 0 {
			params[key] = value
		}
	}

	putAuth := func(auth *Authentication) {
		if auth == nil {
			return
		}

		params["ProxyAuth"] = auth.RequireLocalAuth
		addString("ProxyUsername", auth.LocalAuthUsername)
		addString("ProxyPassword", auth.LocalAuthPassword)

		params["OutproxyAuth"] = auth.RequireOutProxyAuth
		addString("OutproxyUsername", auth.OutProxyAuthUsername)
		addString("OutproxyPassword", auth.OutProxyAuthPassword)
	}

	addString("Description", client.CommonSettings.Description)
	addString("CustomOptions", client.CommonSettings.CustomOptions)
	addString("PrivKeyFile", client.CommonSettings.PrivateKeyFile)
	addString("ReachableBy", client.ReachableBy)
	addString("TargetDestination", client.TunnelDestination)
	addInt("CloseTime", client.IdlePeriod)
	params["Reduce"] = client.CommonSettings.ReduceIdle
	addInt("ReduceCount", client.CommonSettings.ReducedCount)
	addInt("ReduceTime", client.CommonSettings.IdleTime)

	if client.GenerateKeys {
		params["NewDest"] = 0
		params["PersistentClientKey"] = false
	} else if client.NewKeysOnOpen {
		params["NewDest"] = 1
	} else if client.CommonSettings.PrivateKeyFile != "" {
		params["NewDest"] = 2
		params["PersistentClientKey"] = true
	}

	// These fields should never be null, we need to have the JAVA API return default values if someone does not pass them in
	if tunnelLength := client.CommonSettings.TunnelLength; tunnelLength != nil {
		addInt("TunnelLength", tunnelLength.Length)
		addInt("TunnelVariance", tunnelLength.Variance)
	}

	if tunnelQuantity := client.CommonSettings.TunnelQuantity; tunnelQuantity != nil {
		addInt("TunnelQuantity", tunnelQuantity.Quantity)
		addInt("TunnelBackupQuantity", tunnelQuantity.Backup)
	}

	if tunnelCrypto := client.CommonSettings.TunnelCrypto; tunnelCrypto != nil {
		addString("SigType", tunnelCrypto.SigType)
		addString("EncType", tunnelCrypto.EncType)
	}

	if profile := client.ProfileConfig; profile != nil {
		addString("Profile", profile.Profile)
		params["ConnectDelay"] = profile.DelayConnect
	}

	if connect := client.CONNECTConfig; connect != nil {
		addString("ProxyList", connect.OutProxies)
		params["UseOutproxyPlugin"] = connect.UseOutProxyPlugin
		putAuth(connect.Auth)
	}

	if http := client.HTTPConfig; http != nil {
		addString("ProxyList", http.OutProxies)
		addString("SSLProxies", http.SSLOutProxies)
		addString("JumpList", http.JumpURLs)
		params["UseOutproxyPlugin"] = http.UseOutProxyPlugin

		if filtering := http.Filtering; filtering != nil {
			params["AllowUserAgent"] = !filtering.SpoofUserAgent
			params["AllowAccept"] = !filtering.BlockAcceptHeaders
			params["AllowReferer"] = !filtering.BlockReferers
			params["AllowInternalSSL"] = filtering.AllowSSLI2P
		}

		putAuth(http.Auth)
	}

	if socks := client.SOCKSConfig; socks != nil {
		addString("OutproxyType", socks.OutProxyType)
		addString("ProxyList", socks.OutProxies)
		params["UseOutproxyPlugin"] = socks.UseOutProxyPlugin
		putAuth(socks.Auth)
	}

	if streamr := client.STREAMRConfig; streamr != nil {
		addString("TargetHost", streamr.TargetHost)
		addString("TargetDestination", streamr.TargetDestination)
	}

	if irc := client.IRCConfig; irc != nil {
		params["DCC"] = irc.EnableDCC
	}

	retpre, err := Call("TunnelManager", params)
	if err != nil {
		return "", err
	}

	result := retpre["status"].(string)
	return result, nil
}

func AddHiddenService(client ClientConfig) (string, error) {
	return "", nil
}
