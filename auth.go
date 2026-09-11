package i2pcontrol

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"github.com/ybbus/jsonrpc/v2"
)

var (
	rpcClient jsonrpc.RPCClient
	RPCOpts   *jsonrpc.RPCClientOpts
	token     string
)

// Initialize set up the rpcClient with the specified host, port, and path
func Initialize(host, port, path string) {
	RPCOpts = &jsonrpc.RPCClientOpts{
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
	}
	addr := net.JoinHostPort(host, port)
	path = strings.TrimLeft(path, "/")
	rpcClient = jsonrpc.NewClientWithOpts("http://"+addr+"/"+path+"/", RPCOpts)
}

// InitializeWithSelfSignedCert will set up an rpcClient which will accept a self-signed certificate, at the path speciied by cert.
func InitializeWithSelfSignedCert(host, port, path, cert string) error {
	caCert, err := ioutil.ReadFile(cert)
	if err != nil {
		return err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	RPCOpts = &jsonrpc.RPCClientOpts{
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs: caCertPool,
				},
			},
		},
	}
	rpcClient = jsonrpc.NewClientWithOpts("https://"+host+":"+port+"/"+path+"/", RPCOpts)
	return nil
}

// Call an RPC method with params
func Call(method string, params interface{}) (map[string]interface{}, error) {
	if rpcClient == nil {
		return nil, fmt.Errorf("RPC client is not initialized")
	}

	response, err := rpcClient.Call(method, params)
	if err != nil {
		return nil, err
	}
	if response == nil {
		return nil, fmt.Errorf("empty response from RPC call")
	}
	if response.Error != nil {
		return nil, response.Error
	}
	// var retv string
	var retpre map[string]interface{}
	err = response.GetObject(&retpre)
	if err != nil {
		return nil, err
	}

	if method == "AddressBook" || method == "RouterInfo" || method == "TunnelManager" {
		if len(retpre) == 0 {
			return nil, fmt.Errorf("error - empty response from RPC call, implementation pending Prop170 I2PControl Extensions")
		}
	}

	return retpre, nil
}

// Authenticate to the RPC API with a password
func Authenticate(password string) (int, error) {
	retpre, err := Call("Authenticate", map[string]interface{}{
		"API":      1,
		"Password": password,
	})
	if err != nil {
		return -1, err
	}
	value, err := responseValue(retpre, "Token")
	if err != nil {
		return -1, err
	}
	nextToken := value.(string)
	value, err = responseValue(retpre, "API")
	if err != nil {
		return -1, err
	}
	version := int(value.(float64))
	token = nextToken
	return version, nil
}
