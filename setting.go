package i2pcontrol

// Checks our UPnp
func Upnp() (string, error) {
	retpre, err := Call("NetworkSetting", map[string]interface{}{
		"i2p.router.net.upnp": nil,
		"Token":               token,
	})
	if err != nil {
		return "", err
	}
	value, err := responseValue(retpre, "i2p.router.net.upnp")
	if err != nil {
		return "Upnp Disabled", nil
	}
	result := value.(string)

	return result, nil
}
