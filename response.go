package i2pcontrol

import "fmt"

func responseValue(values map[string]interface{}, key string) (interface{}, error) {
	if values == nil {
		return nil, fmt.Errorf("missing RPC response field %q: response is nil", key)
	}

	value, ok := values[key]
	if !ok {
		return nil, fmt.Errorf("missing RPC response field %q", key)
	}
	if value == nil {
		return nil, fmt.Errorf("RPC response field %q is nil", key)
	}

	return value, nil
}
