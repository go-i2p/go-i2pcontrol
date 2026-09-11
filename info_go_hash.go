package i2pcontrol

// RouterHash queries the hash of the router, which is a base64-encoded string that uniquely identifies the router on the I2P network.
// this is a go-i2p only function, as the router hash is not available in the I2P Control Protocol but it is available in Go as an
// unofficial extension. It uses prop170 i2p.router.hash, which is non-negotiable for a go-i2p router console, but is not guaranteed
// to be available in other implementations of the I2P Control Protocol. It is recommended to use this function only with go-i2p
// routers.
func RouterHash() (string, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.hash": nil,
		"Token":           token,
	})
	if err != nil {
		return "", err
	}
	value, err := responseValue(retpre, "i2p.router.hash")
	if err != nil {
		return "", err
	}
	result := value.(string)
	return result, nil
}
