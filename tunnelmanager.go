package i2pcontrol

// Going to need different fields for each type of service, but for now just a placeholder struct
// maybe an interface? Could help solve the problem of different types of services?

// TODO: Once we get everything working on the golang end, I need to ensure the
// 	JSON RPC is foolproof. I need to make sure that the API is consistent and that the error handling is robust.

// TODO: Post check, ban, and total needs to be x * 60 (Pretty sure)

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

type STANDARDConfig struct {
	ProfileConfig ProfileConfig
}

type ProfileConfig struct {
	Profile      string // (interactive or default)
	ConnectDelay bool
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
	PersistentPrivateKeyFile string

	GenerateKeys bool // when this is true, NewKeysOnOpen becomes null / ignored

	STANDARDConfig *STANDARDConfig
	CONNECTConfig  *CONNECTConfig
	HTTPConfig     *HTTPConfig
	SOCKSConfig    *SOCKSConfig
	STREAMRConfig  *STREAMRConfig
	IRCConfig      *IRCConfig
}

type LeaseSetClientAuth struct {
	Name string
	Key  string
}

type LeaseSetConfig struct {
	LeaseSetEnc         string
	LeaseSetPassword    string
	LeaseSetClientAuths []LeaseSetClientAuth
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
	AccessOption    string
	AccessList      string
	AccessListFile  string
	UniquePerClient bool
	MultiHoming     bool
}

type HTTPServerConfig struct {
	WebHost              string
	POSTLimitConfig      PostLimitConfig
	BlockReferers        bool
	BlockUserAgents      bool
	UserAgents           string
	BlockAccessInProxies bool
	HTTPBidirConfig      *HTTPBidirConfig
}

type HTTPBidirConfig struct {
	ReachableBy string
}

type ServiceConfig struct {
	CommonSettings            CommonConfig
	LeaseSetConfig            LeaseSetConfig
	ServerThrottling          ServerThrottling
	TunnelAccessControlConfig TunnelAccessControlConfig
	PrivateKeyFile            string
	Host                      string
	TargetPort                int
	Profile                   string // (Used for everything besides STREAMR)
	ConnectDelay              bool
	STREAMRServerConfig       *STREAMRServerConfig
	HTTPServerConfig          *HTTPServerConfig
}

// if the edit structs grow, create a EditOptions struct and add it to all the edit functions

type EditServiceConfig struct {
	NewName       *string
	ServiceConfig ServiceConfig
}

type EditClientConfig struct {
	NewName      *string
	ClientConfig ClientConfig
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
	// a get action returns the tunnel options; this is the only action that returns them
	if action == "get" {
		tunnelOptions = retpre["info"].(map[string]interface{})
	}

	return result, tunnelOptions, nil
}

// EditClientTunnel is the same. Only difference is looking for a NewName param
func EditClientTunnel(client EditClientConfig) (string, error) {
	params := setCommonParams(client.ClientConfig.CommonSettings, "edit")
	if client.NewName != nil {
		addString(params, "NewName", *client.NewName)
	}
	setClientParams(client.ClientConfig, params)
	retpre, err := Call("TunnelManager", params)
	if err != nil {
		return "", err
	}

	result := retpre["status"].(string)
	return result, nil
}

func EditHiddenService(service EditServiceConfig) (string, error) {
	params := setCommonParams(service.ServiceConfig.CommonSettings, "edit")
	if service.NewName != nil {
		addString(params, "NewName", *service.NewName)
	}
	setServiceParams(service.ServiceConfig, params)
	retpre, err := Call("TunnelManager", params)
	if err != nil {
		return "", err
	}

	result := retpre["status"].(string)
	return result, nil
}

// AddClientTunnel creates a new hidden service.
func AddClientTunnel(client ClientConfig) (string, error) {
	params := setCommonParams(client.CommonSettings, "create")

	setClientParams(client, params)

	retpre, err := Call("TunnelManager", params)
	if err != nil {
		return "", err
	}

	result := retpre["status"].(string)
	return result, nil
}

func AddHiddenService(service ServiceConfig) (string, error) {
	params := setCommonParams(service.CommonSettings, "create")

	setServiceParams(service, params)

	retpre, err := Call("TunnelManager", params)
	if err != nil {
		return "", err
	}

	result := retpre["status"].(string)
	return result, nil
}

func setCommonParams(common CommonConfig, action string) map[string]interface{} {
	params := map[string]interface{}{
		"Action":      action,
		"Name":        common.Name,
		"StartOnLoad": common.AutoStart,
		"UseSSL":      common.SSL,
		"Port":        common.Port,
		"Type":        common.Type,
		"Reduce":      common.ReduceIdle,
	}
	addString(params, "Description", common.Description)
	addString(params, "CustomOptions", common.CustomOptions)

	params["ReduceCount"] = common.ReducedCount
	params["ReduceTime"] = common.IdleTime

	if tunnelLength := common.TunnelLength; tunnelLength != nil {
		params["TunnelLength"] = tunnelLength.Length
		params["TunnelVariance"] = tunnelLength.Variance
	}

	if tunnelQuantity := common.TunnelQuantity; tunnelQuantity != nil {
		params["TunnelQuantity"] = tunnelQuantity.Quantity
		params["TunnelBackupQuantity"] = tunnelQuantity.Backup
	}

	if tunnelCrypto := common.TunnelCrypto; tunnelCrypto != nil {
		addString(params, "SigType", tunnelCrypto.SigType)
		addString(params, "EncType", tunnelCrypto.EncType)
	}

	return params
}

func setServiceParams(service ServiceConfig, params map[string]interface{}) {
	addString(params, "PrivKeyFile", service.PrivateKeyFile)
	addString(params, "TargetHost", service.Host)
	addString(params, "Profile", service.Profile)
	params["ConnectDelay"] = service.ConnectDelay

	params["TargetPort"] = service.TargetPort

	addString(params, "EncryptLeaseSet", service.LeaseSetConfig.LeaseSetEnc)
	addString(params, "OptionalLookup", service.LeaseSetConfig.LeaseSetPassword)
	if len(service.LeaseSetConfig.LeaseSetClientAuths) > 0 {
		auths := make([]map[string]string, 0, len(service.LeaseSetConfig.LeaseSetClientAuths))
		for _, auth := range service.LeaseSetConfig.LeaseSetClientAuths {
			auths = append(auths, map[string]string{
				"Name": auth.Name,
				"Key":  auth.Key,
			})
		}
		params["LeaseSetClientAuths"] = auths
	} else {
		delete(params, "LeaseSetClientAuths")
	}

	params["MaxConcurrentConns"] = service.ServerThrottling.MaxConcurrentConnections
	params["ClientPerMinute"] = service.ServerThrottling.PerClientMinuteLimit
	params["ClientPerHour"] = service.ServerThrottling.PerClientHourLimit
	params["ClientPerDay"] = service.ServerThrottling.PerClientDayLimit
	params["TotalInPerMinute"] = service.ServerThrottling.TotalPerMinuteLimit
	params["TotalInPerHour"] = service.ServerThrottling.TotalPerHourLimit
	params["TotalInPerDay"] = service.ServerThrottling.TotalPerDayLimit

	addString(params, "AccessOption", service.TunnelAccessControlConfig.AccessOption)
	addString(params, "AccessList", service.TunnelAccessControlConfig.AccessList)
	addString(params, "FilterFilePath", service.TunnelAccessControlConfig.AccessListFile)
	params["UniqueLocalAddressPerClient"] = service.TunnelAccessControlConfig.UniquePerClient
	params["MultiHoming"] = service.TunnelAccessControlConfig.MultiHoming

	if streamr := service.STREAMRServerConfig; streamr != nil {
		addString(params, "ReachableBy", streamr.ReachableBy)
	}

	if http := service.HTTPServerConfig; http != nil {
		addString(params, "WebsiteHostname", http.WebHost)
		params["BlockReferers"] = http.BlockReferers
		params["BlockUserAgents"] = http.BlockUserAgents
		params["BlockAccessInProxies"] = http.BlockAccessInProxies
		addString(params, "UserAgents", http.UserAgents)

		params["PostLimit"] = http.POSTLimitConfig.POSTLimit
		params["PostLimitTime"] = http.POSTLimitConfig.ClientBanDuration
		params["PerClientPeriod"] = http.POSTLimitConfig.ClientPerPeriod
		params["TotalPeriod"] = http.POSTLimitConfig.TotalPerPeriod
		params["TotalBanTime"] = http.POSTLimitConfig.TotalBanDuration
		if http.HTTPBidirConfig != nil {
			addString(params, "ReachableBy", http.HTTPBidirConfig.ReachableBy)
		}
	}
}

func setClientParams(client ClientConfig, params map[string]interface{}) {
	isStreamrClient := client.CommonSettings.Type == "streamrclient"

	params["Shared"] = client.SharedClient && !isStreamrClient
	params["DelayOpen"] = client.DelayOpen && !isStreamrClient

	putAuth := func(auth *Authentication) {
		if auth == nil {
			return
		}

		params["ProxyAuth"] = auth.RequireLocalAuth
		addString(params, "ProxyUsername", auth.LocalAuthUsername)
		addString(params, "ProxyPassword", auth.LocalAuthPassword)

		params["OutproxyAuth"] = auth.RequireOutProxyAuth
		addString(params, "OutproxyUsername", auth.OutProxyAuthUsername)
		addString(params, "OutproxyPassword", auth.OutProxyAuthPassword)
	}

	addString(params, "PrivKeyFile", client.PersistentPrivateKeyFile)

	addString(params, "ReachableBy", client.ReachableBy)
	addString(params, "TargetDestination", client.TunnelDestination)

	params["CloseTime"] = client.IdlePeriod

	params["Close"] = client.CloseWhenIdle

	if client.GenerateKeys {
		params["NewDest"] = 0
		params["PersistentClientKey"] = false
	} else if client.NewKeysOnOpen {
		params["NewDest"] = 1
	} else if client.PersistentPrivateKeyFile != "" &&
		client.CommonSettings.Type != "httpclient" &&
		client.CommonSettings.Type != "connectclient" &&
		client.CommonSettings.Type != "streamrclient" {
		params["NewDest"] = 2
		params["PersistentClientKey"] = true
	}

	if standard := client.STANDARDConfig; standard != nil {
		addString(params, "Profile", standard.ProfileConfig.Profile)
		params["ConnectDelay"] = standard.ProfileConfig.ConnectDelay
	}

	if connect := client.CONNECTConfig; connect != nil {
		addString(params, "ProxyList", connect.OutProxies)
		params["UseOutproxyPlugin"] = connect.UseOutProxyPlugin
		putAuth(connect.Auth)
	}

	if http := client.HTTPConfig; http != nil {
		addString(params, "ProxyList", http.OutProxies)
		addString(params, "SSLProxies", http.SSLOutProxies)
		addString(params, "JumpList", http.JumpURLs)
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
		addString(params, "OutproxyType", socks.OutProxyType)
		addString(params, "ProxyList", socks.OutProxies)
		params["UseOutproxyPlugin"] = socks.UseOutProxyPlugin
		putAuth(socks.Auth)
	}

	if streamr := client.STREAMRConfig; streamr != nil {
		addString(params, "TargetHost", streamr.TargetHost)
		addString(params, "TargetDestination", streamr.TargetDestination)
	}

	if irc := client.IRCConfig; irc != nil {
		params["DCC"] = irc.EnableDCC
	}
}

// addString, using this as a validation technique, if a field is empty we omit it and act as if it is null,
// this is so the JAVA API will react to them as null.
func addString(params map[string]interface{}, key, value string) {
	if value != "" {
		params[key] = value
	} else {
		delete(params, key)
	}
}
