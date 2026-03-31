package i2pcontrol

import "fmt"

// ParticipatingTunnels gets the number of participating tunnels the router has currently
func ParticipatingTunnels() (int, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.net.tunnels.participating": nil,
		"Token":                                token,
	})
	if err != nil {
		return -1, err
	}
	result := int(retpre["i2p.router.net.tunnels.participating"].(float64))
	return result, nil
}

// Status queries the status of the router
func Status() (string, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.status": nil,
		"Token":             token,
	})
	if err != nil {
		return "", err
	}
	result := retpre["i2p.router.status"].(string)
	return result, nil
}

// NetStatus queries the status of the network connection
func NetStatus() (string, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.net.status": nil,
		"Token":                 token,
	})
	if err != nil {
		return "", err
	}
	result := int(retpre["i2p.router.net.status"].(float64))
	switch result {
	case 0:
		return "OK", nil
	case 1:
		return "TESTING", nil
	case 2:
		return "FIREWALLED", nil
	case 3:
		return "HIDDEN", nil
	case 4:
		return "WARN_FIREWALLED_AND_FAST", nil
	case 5:
		return "WARN_FIREWALLED_AND_FLOODFILL", nil
	case 6:
		return "WARN_FIREWALLED_WITH_INBOUND_TCP", nil
	case 7:
		return "WARN_FIREWALLED_WITH_UDP_DISABLED", nil
	case 8:
		return "ERROR_I2CP", fmt.Errorf("I2CP error")
	case 9:
		return "ERROR_CLOCK_SKEW", fmt.Errorf("Clock skew error")
	case 10:
		return "ERROR_PRIVATE_TCP_ADDRESS", fmt.Errorf("Private TCP address error")
	case 11:
		return "ERROR_SYMMETRIC_NAT", fmt.Errorf("Symmetric NAT error")
	case 12:
		return "ERROR_UDP_PORT_IN_USE", fmt.Errorf("UDP port in use error")
	case 13:
		return "ERROR_NO_ACTIVE_PEERS_CHECK_CONNECTION_AND_FIREWALL", fmt.Errorf("No active peers check connection and firewall")
	case 14:
		return "ERROR_UDP_DISABLED_AND_TCP_UNSET", fmt.Errorf("UDP disabled and TCP unset")
	default:
		return "unexpected result", fmt.Errorf("unexpected result  %d", result)
	}
}

// Reseeding checks if the I2P Router is reseeding
func Reseeding() (bool, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.netdb.isreseeding": nil,
		"Token":                        token,
	})
	if err != nil {
		return false, err
	}
	result := retpre["i2p.router.netdb.isreseeding"].(bool)
	return result, nil
}

// IncomingBW bandwidth per second
func IncomingBW() (int, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.net.bw.inbound.1s": nil,
		"Token":                        token,
	})
	if err != nil {
		return -1, err
	}
	result := int(retpre["i2p.router.net.bw.inbound.1s"].(float64))
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
	result := int(retpre["i2p.router.net.bw.outbound.1s"].(float64))
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
	result := int64(retpre["i2p.router.uptime"].(float64))
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
	result := int(retpre["i2p.router.netdb.knownpeers"].(float64))
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

	result := int(retpre["i2p.router.netdb.activepeers"].(float64))
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

	result := retpre["i2p.router.version"].(string)
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

	result := retpre["i2p.router.id"].(string)
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
	result := retpre["i2p.router.info"].(string)
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

	result := int(retpre["i2p.router.clockskew"].(float64))
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

	result := int(retpre["i2p.router.net.tunnels.exploratory.inbound"].(float64))
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

	result := int(retpre["i2p.router.net.tunnels.exploratory.outbound"].(float64))
	return result, nil
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

	result := int(retpre["i2p.router.net.tunnels.client.inbound"].(float64))
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

	result := int(retpre["i2p.router.net.tunnels.client.outbound"].(float64))
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

	resultInterface := retpre["i2p.router.netdb.peers"].([]interface{})
	result := make([]string, len(resultInterface))
	for i, v := range resultInterface {
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

	resultInterface := retpre["i2p.router.netdb.activepeers.list"].([]interface{})
	result := make([]string, len(resultInterface))
	for i, v := range resultInterface {
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

	resultInterface := retpre["i2p.router.netdb.activepeers.info"].([]interface{})
	result := make([]string, len(resultInterface))
	for i, v := range resultInterface {
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

	resultInterface := retpre["i2p.router.netdb.peers.list"].([]interface{})
	result := make([]string, len(resultInterface))
	for i, v := range resultInterface {
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

	resultInterface := retpre["i2p.router.netdb.peers.info"].([]interface{})
	result := make([]string, len(resultInterface))
	for i, v := range resultInterface {
		result[i] = v.(string)
	}
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

	result := retpre["i2p.router.netdb.activepeers.stats"].([]interface{})
	return result, nil
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

	return retpre["i2p.router.addressbook.private.list"].([]interface{}), nil
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

	return retpre["i2p.router.addressbook.local.list"].([]interface{}), nil
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
	return retpre["i2p.router.addressbook.router.list"].([]interface{}), nil
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
	return retpre["i2p.router.addressbook.published.list"].([]interface{}), nil
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

	raw := retpre["i2p.router.addressbook.config"].(map[string]interface{})
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

	raw := retpre["i2p.router.addressbook.subscriptions"].(map[string]interface{})
	path := raw["path"]
	entries := raw["entries"]

	pathAndEntries := []interface{}{path, entries}

	return pathAndEntries, nil
}
