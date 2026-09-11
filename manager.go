package i2pcontrol

// Echo test the API
func Echo(echo string) (string, error) {
	retpre, err := Call("Echo", map[string]interface{}{
		"Echo":  echo,
		"Token": token,
	})
	if err != nil {
		return "", err
	}
	value, err := responseValue(retpre, "Result")
	if err != nil {
		return "", err
	}
	result := value.(string)
	return result, nil
}

// RestartGraceful initiates a graceful restart, which will occur in around 11 minutes.
func RestartGraceful() (string, error) {
	_, err := Call("RouterManager", map[string]interface{}{
		"RestartGraceful": nil,
		"Token":           token,
	})
	if err != nil {
		return "", err
	}
	return "Graceful Restart Initiated", nil
}

// Restart the router immediately
func Restart() (string, error) {
	_, err := Call("RouterManager", map[string]interface{}{
		"Restart": nil,
		"Token":   token,
	})
	if err != nil {
		return "", err
	}
	return "Restart Initiated", nil
}

func Reseed() (string, error) {
	_, err := Call("RouterManager", map[string]interface{}{
		"Reseed": nil,
		"Token":  token,
	})
	if err != nil {
		return "", err
	}
	return "Reseed Initiated", nil
}

// Shutdown initiates a graceful restart, which will occur in around 11 minutes.
func ShutdownGraceful() (string, error) {
	_, err := Call("RouterManager", map[string]interface{}{
		"ShutdownGraceful": nil,
		"Token":            token,
	})
	if err != nil {
		return "", err
	}
	return "Graceful Shutdown Initiated", nil
}

// Shutdown the router immediately
func Shutdown() (string, error) {
	_, err := Call("RouterManager", map[string]interface{}{
		"Shutdown": nil,
		"Token":    token,
	})
	if err != nil {
		return "", err
	}
	return "Shutdown Initiated", nil
}

// FindUpdates pings the server to check for updates
func FindUpdates() (bool, error) {
	retpre, err := Call("RouterManager", map[string]interface{}{
		"FindUpdates": nil,
		"Token":       token,
	})
	if err != nil {
		return false, err
	}
	value, err := responseValue(retpre, "FindUpdates")
	if err != nil {
		return false, err
	}
	result := value.(bool)
	return result, nil
}

// Starts an update.
func Update() (string, error) {
	_, err := Call("RouterManager", map[string]interface{}{
		"Update": nil,
		"Token":  token,
	})
	if err != nil {
		return "", err
	}
	return "Update initiated", nil
}
