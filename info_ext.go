package i2pcontrol

import "fmt"

// IncomingBW bandwidth per second
func IncomingBW() (int, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.net.bw.inbound.1s": nil,
		"Token":                        token,
	})
	if err != nil {
		return -1, err
	}
	value, err := responseValue(retpre, "i2p.router.net.bw.inbound.1s")
	if err != nil {
		return -1, err
	}
	result := int(value.(float64))
	return result, nil
}

// OutgoingBw bandwidth per second
func OutgoingBw() (int, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.net.bw.outbound.1s": nil,
		"Token":                         token,
	})
	if err != nil {
		return -1, err
	}
	value, err := responseValue(retpre, "i2p.router.net.bw.outbound.1s")
	if err != nil {
		return -1, err
	}
	result := int(value.(float64))
	return result, nil
}

// TotalOutgoingBW total outgoing bandwidth
func TotalOutgoingBW() (int, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.net.bw.used.outbound.total": nil,
		"Token":                                 token,
	})
	if err != nil {
		return -1, err
	}
	value, err := responseValue(retpre, "i2p.router.net.bw.used.outbound.total")
	if err != nil {
		return -1, err
	}
	result := int(value.(float64))
	return result, nil
}

// TotalIncomingBW total incoming bandwidth
func TotalIncomingBW() (int, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.net.bw.used.inbound.total": nil,
		"Token":                                token,
	})
	if err != nil {
		return -1, err
	}
	value, err := responseValue(retpre, "i2p.router.net.bw.used.inbound.total")
	if err != nil {
		return -1, err
	}
	result := int(value.(float64))
	return result, nil
}

// UpTime of the router
func UpTime() (int64, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.uptime": nil,
		"Token":             token,
	})
	if err != nil {
		return -1, err
	}
	value, err := responseValue(retpre, "i2p.router.uptime")
	if err != nil {
		return -1, err
	}
	result := int64(value.(float64))
	return result, nil
}

// KnownPeers All the known peers
func KnownPeers() (int, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.netdb.knownpeers": nil,
		"Token":                       token,
	})
	if err != nil {
		return -1, err
	}
	value, err := responseValue(retpre, "i2p.router.netdb.knownpeers")
	if err != nil {
		return -1, err
	}
	result := int(value.(float64))
	return result, nil
}

// ActivePeers All Active peers
func ActivePeers() (int, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.netdb.activepeers": nil,
		"Token":                        token,
	})
	if err != nil {
		return -1, err
	}

	value, err := responseValue(retpre, "i2p.router.netdb.activepeers")
	if err != nil {
		return -1, err
	}
	result := int(value.(float64))
	return result, nil
}

// Version the current router version
func Version() (string, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.version": nil,
		"Token":              token,
	})
	if err != nil {
		return "", err
	}

	value, err := responseValue(retpre, "i2p.router.version")
	if err != nil {
		return "", err
	}
	result := value.(string)
	return result, nil
}

// RouterID the current router ID
func RouterID() (string, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.id": nil,
		"Token":         token,
	})
	if err != nil {
		return "", err
	}

	value, err := responseValue(retpre, "i2p.router.id")
	if err != nil {
		return "", err
	}
	result := value.(string)
	return result, nil
}

// RouterInfo the current router ID in base64 format
func RouterInfo() (string, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.info": nil,
		"Token":           token,
	})
	if err != nil {
		return "", err
	}
	value, err := responseValue(retpre, "i2p.router.info")
	if err != nil {
		return "", err
	}
	result := value.(string)
	return result, nil
}

// ClockSkew the current clock skew of the router
func ClockSkew() (int, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.clockskew": nil,
		"Token":                token,
	})
	if err != nil {
		return -1, err
	}

	value, err := responseValue(retpre, "i2p.router.clockskew")
	if err != nil {
		return -1, err
	}
	result := int(value.(float64))
	return result, nil
}

// InboundExploratoryTunnels gets the number of inbound exploratory tunnels
func InboundExploratoryTunnels() (int, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.net.tunnels.exploratory.inbound": nil,
		"Token": token,
	})
	if err != nil {
		return -1, err
	}

	value, err := responseValue(retpre, "i2p.router.net.tunnels.exploratory.inbound")
	if err != nil {
		return -1, err
	}
	result := int(value.(float64))
	return result, nil
}

// OutboundExploratoryTunnels gets the number of outbound exploratory tunnels
func OutboundExploratoryTunnels() (int, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.net.tunnels.exploratory.outbound": nil,
		"Token": token,
	})
	if err != nil {
		return -1, err
	}

	value, err := responseValue(retpre, "i2p.router.net.tunnels.exploratory.outbound")
	if err != nil {
		return -1, err
	}
	result := int(value.(float64))
	return result, nil
}

// ExploratoryTunnelsInfo gets the exploratory tunnels info
func ExploratoryTunnelsInfo() ([]interface{}, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.net.tunnels.exploratory.info.list": nil,
		"Token": token,
	})
	if err != nil {
		return nil, err
	}
	value, err := responseValue(retpre, "i2p.router.net.tunnels.exploratory.info.list")
	if err != nil {
		return nil, err
	}
	return value.([]interface{}), nil
}

// ClientTunnelsInfo gets the client tunnels info
func ClientTunnelsInfo() ([]interface{}, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.net.tunnels.client.info.list": nil,
		"Token": token,
	})
	if err != nil {
		return nil, err
	}
	value, err := responseValue(retpre, "i2p.router.net.tunnels.client.info.list")
	if err != nil {
		return nil, err
	}
	return value.([]interface{}), nil
}

// ShareRatio gets the share ratio of the router
func ShareRatio() (float64, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.net.tunnels.shareratio": nil,
		"Token":                             token,
	})
	if err != nil {
		return -1, err
	}
	value, err := responseValue(retpre, "i2p.router.net.tunnels.shareratio")
	if err != nil {
		return -1, err
	}
	result := value.(float64)
	return result, nil
}

// ParticipatingTunnelsCount gets the number of participating tunnels (inactive + active)
func ParticipatingTunnelsCount() (int, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.net.tunnels.participating": nil,
		"Token":                                token,
	})
	if err != nil {
		return -1, err
	}
	value, err := responseValue(retpre, "i2p.router.net.tunnels.participating")
	if err != nil {
		return -1, err
	}
	result := int(value.(float64))
	return result, nil
}

// ParticipatingTunnelsInfo gets the participating tunnels info
func ParticipatingTunnelsInfo() ([]interface{}, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.net.tunnels.participating.info": nil,
		"Token": token,
	})
	if err != nil {
		return nil, err
	}
	value, err := responseValue(retpre, "i2p.router.net.tunnels.participating.info")
	if err != nil {
		return nil, err
	}
	return value.([]interface{}), nil
}

// InboundClientTunnels gets the number of inbound client tunnels
func InboundClientTunnels() (int, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.net.tunnels.client.inbound": nil,
		"Token":                                 token,
	})
	if err != nil {
		return -1, err
	}

	value, err := responseValue(retpre, "i2p.router.net.tunnels.client.inbound")
	if err != nil {
		return -1, err
	}
	result := int(value.(float64))
	return result, nil
}

// OutboundClientTunnels gets the number of outbound client tunnels
func OutboundClientTunnels() (int, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.net.tunnels.client.outbound": nil,
		"Token":                                  token,
	})
	if err != nil {
		return -1, err
	}

	value, err := responseValue(retpre, "i2p.router.net.tunnels.client.outbound")
	if err != nil {
		return -1, err
	}
	result := int(value.(float64))
	return result, nil
}

// KnownPeersList gets a list of known peers
func KnownPeersList() ([]string, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.netdb.peers": nil,
		"Token":                  token,
	})
	if err != nil {
		return nil, err
	}

	value, err := responseValue(retpre, "i2p.router.netdb.peers")
	if err != nil {
		return nil, err
	}
	resultInterface := value.([]interface{})
	result := make([]string, len(resultInterface))
	for i, v := range resultInterface {
		if v == nil {
			return nil, fmt.Errorf("RPC response field %q[%d] is nil", "i2p.router.netdb.peers", i)
		}
		result[i] = v.(string)
	}
	return result, nil
}

// ActivePeersList gets a list of active peers
func ActivePeersList() ([]string, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.netdb.activepeers.list": nil,
		"Token":                             token,
	})
	if err != nil {
		return nil, err
	}

	value, err := responseValue(retpre, "i2p.router.netdb.activepeers.list")
	if err != nil {
		return nil, err
	}
	resultInterface := value.([]interface{})
	result := make([]string, len(resultInterface))
	for i, v := range resultInterface {
		if v == nil {
			return nil, fmt.Errorf("RPC response field %q[%d] is nil", "i2p.router.netdb.activepeers.list", i)
		}
		result[i] = v.(string)
	}
	return result, nil
}

// ActivePeersInfo gets a list of active peers info
func ActivePeersInfo() ([]string, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.netdb.activepeers.info": nil,
		"Token":                             token,
	})
	if err != nil {
		return nil, err
	}

	value, err := responseValue(retpre, "i2p.router.netdb.activepeers.info")
	if err != nil {
		return nil, err
	}
	resultInterface := value.([]interface{})
	result := make([]string, len(resultInterface))
	for i, v := range resultInterface {
		if v == nil {
			return nil, fmt.Errorf("RPC response field %q[%d] is nil", "i2p.router.netdb.activepeers.info", i)
		}
		result[i] = v.(string)
	}
	return result, nil
}

// AllPeersList gets a list of all known peer hashes
func AllPeersList() ([]string, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.netdb.peers.list": nil,
		"Token":                       token,
	})
	if err != nil {
		return nil, err
	}

	value, err := responseValue(retpre, "i2p.router.netdb.peers.list")
	if err != nil {
		return nil, err
	}
	resultInterface := value.([]interface{})
	result := make([]string, len(resultInterface))
	for i, v := range resultInterface {
		if v == nil {
			return nil, fmt.Errorf("RPC response field %q[%d] is nil", "i2p.router.netdb.peers.list", i)
		}
		result[i] = v.(string)
	}
	return result, nil
}

// AllPeersInfo gets the raw base64-encoded RouterInfo blobs for all known peers
func AllPeersInfo() ([]string, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.netdb.peers.info": nil,
		"Token":                       token,
	})
	if err != nil {
		return nil, err
	}

	value, err := responseValue(retpre, "i2p.router.netdb.peers.info")
	if err != nil {
		return nil, err
	}
	resultInterface := value.([]interface{})
	result := make([]string, len(resultInterface))
	for i, v := range resultInterface {
		if v == nil {
			return nil, fmt.Errorf("RPC response field %q[%d] is nil", "i2p.router.netdb.peers.info", i)
		}
		result[i] = v.(string)
	}
	return result, nil
}

// NTCPLimit gets the NTCP limit
func NTCPLimit() (int, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.netdb.ntcp.limit": nil,
		"Token":                       token,
	})
	if err != nil {
		return -1, err
	}
	value, err := responseValue(retpre, "i2p.router.netdb.ntcp.limit")
	if err != nil {
		return -1, err
	}
	result := int(value.(float64))
	return result, nil
}

// SSULimit gets the SSU limit
func SSULimit() (int, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.netdb.ssu.limit": nil,
		"Token":                      token,
	})
	if err != nil {
		return -1, err
	}
	value, err := responseValue(retpre, "i2p.router.netdb.ssu.limit")
	if err != nil {
		return -1, err
	}
	result := int(value.(float64))
	return result, nil
}

// NewsEntries gets the latest news entries
func NewsEntries() (string, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.news": nil,
		"Token":           token,
	})
	if err != nil {
		return "", err
	}
	value, err := responseValue(retpre, "i2p.router.news")
	if err != nil {
		return "", err
	}
	result := value.(string)
	return result, nil
}

// BannedPeers gets the banned peers
func BannedPeers() (map[string]interface{}, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.netdb.bannedpeers": nil,
		"Token":                        token,
	})
	if err != nil {
		return nil, err
	}
	value, err := responseValue(retpre, "i2p.router.netdb.bannedpeers")
	if err != nil {
		return nil, err
	}
	return value.(map[string]interface{}), nil
}

func RouterLogs() ([]string, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.logs": nil,
		"Token":           token,
	})
	if err != nil {
		return nil, err
	}
	value, err := responseValue(retpre, "i2p.router.logs")
	if err != nil {
		return nil, err
	}
	resultInterface := value.([]interface{})
	result := make([]string, len(resultInterface))

	for i, v := range resultInterface {
		if v == nil {
			return nil, fmt.Errorf("RPC response field %q[%d] is nil", "i2p.router.logs", i)
		}
		result[i] = v.(string)
	}

	return result, nil
}

func ClearRouterLogs() (string, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.logs.clear": nil,
		"Token":                 token,
	})
	if err != nil {
		return "", err
	}
	value, err := responseValue(retpre, "i2p.router.logs.clear")
	if err != nil {
		return "", err
	}
	result := value.(string)
	return result, nil
}

// ActivePeersStats gets detailed live stats for all active peer connections
func ActivePeersStats() ([]interface{}, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.netdb.activepeers.stats": nil,
		"Token":                              token,
	})
	if err != nil {
		return nil, err
	}

	value, err := responseValue(retpre, "i2p.router.netdb.activepeers.stats")
	if err != nil {
		return nil, err
	}
	return value.([]interface{}), nil
}

// PrivateAddressBook gets the list of private address book entries
func PrivateAddressBook() ([]interface{}, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.addressbook.private.list": nil,
		"Token":                               token,
	})
	if err != nil {
		return nil, err
	}

	value, err := responseValue(retpre, "i2p.router.addressbook.private.list")
	if err != nil {
		return nil, err
	}
	return value.([]interface{}), nil
}

// LocalAddressBook gets the list of local address book entries
func LocalAddressBook() ([]interface{}, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.addressbook.local.list": nil,
		"Token":                             token,
	})
	if err != nil {
		return nil, err
	}

	value, err := responseValue(retpre, "i2p.router.addressbook.local.list")
	if err != nil {
		return nil, err
	}
	return value.([]interface{}), nil
}

// RouterAddressBook gets the list of router address book entries
func RouterAddressBook() ([]interface{}, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.addressbook.router.list": nil,
		"Token":                              token,
	})
	if err != nil {
		return nil, err
	}
	value, err := responseValue(retpre, "i2p.router.addressbook.router.list")
	if err != nil {
		return nil, err
	}
	return value.([]interface{}), nil
}

// PublishedAddressBook gets the list of published address book entries
func PublishedAddressBook() ([]interface{}, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.addressbook.published.list": nil,
		"Token":                                 token,
	})
	if err != nil {
		return nil, err
	}
	value, err := responseValue(retpre, "i2p.router.addressbook.published.list")
	if err != nil {
		return nil, err
	}
	return value.([]interface{}), nil
}

// AddressBookConfig gets the address book configuration, including the path and entries
func AddressBookConfig() ([]interface{}, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.addressbook.config": nil,
		"Token":                         token,
	})
	if err != nil {
		return nil, err
	}

	value, err := responseValue(retpre, "i2p.router.addressbook.config")
	if err != nil {
		return nil, err
	}
	raw := value.(map[string]interface{})
	path := raw["path"]
	entries := raw["entries"]

	pathAndEntries := []interface{}{path, entries}

	return pathAndEntries, nil
}

// AddressBookSubscriptions gets the address book subscriptions, including the path and entries
func AddressBookSubscriptions() ([]interface{}, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.addressbook.subscriptions": nil,
		"Token":                                token,
	})
	if err != nil {
		return nil, err
	}

	value, err := responseValue(retpre, "i2p.router.addressbook.subscriptions")
	if err != nil {
		return nil, err
	}
	raw := value.(map[string]interface{})
	path := raw["path"]
	entries := raw["entries"]

	pathAndEntries := []interface{}{path, entries}

	return pathAndEntries, nil
}

func HiddenServices() ([]interface{}, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.net.tunnels.i2ptunnel": nil,
		"Token":                            token,
	})
	if err != nil {
		return nil, err
	}
	value, err := responseValue(retpre, "i2p.router.net.tunnels.i2ptunnel")
	if err != nil {
		return nil, err
	}
	return value.([]interface{}), nil
}
