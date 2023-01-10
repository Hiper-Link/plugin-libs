package shared

import (
	"net/rpc"
)

// RPCClient is an implementation of KV that talks over RPC.
type RPCClient struct {
	client *rpc.Client
}

func (m *RPCClient) OnLoad(pluginInterface string) ([]byte, error) {
	var resp []byte
	err := m.client.Call("Plugin.OnLoad", pluginInterface, &resp)
	return resp, err
}

func (m *RPCClient) OnUnload(pluginInterface string) ([]byte, error) {
	var resp []byte
	err := m.client.Call("Plugin.OnUnload", pluginInterface, &resp)
	return resp, err
}

func (m *RPCClient) OnInstall(pluginInterface string) ([]byte, error) {
	var resp []byte
	err := m.client.Call("Plugin.OnInstall", pluginInterface, &resp)
	return resp, err
}

func (m *RPCClient) OnUninstall(pluginInterface string) ([]byte, error) {
	var resp []byte
	err := m.client.Call("Plugin.OnUninstall", pluginInterface, &resp)
	return resp, err
}

func (m *RPCClient) OnStart(pluginInterface string) ([]byte, error) {
	var resp []byte
	err := m.client.Call("Plugin.OnStart", pluginInterface, &resp)
	return resp, err
}

func (m *RPCClient) OnStop(pluginInterface string) ([]byte, error) {
	var resp []byte
	err := m.client.Call("Plugin.OnStop", pluginInterface, &resp)
	return resp, err
}

func (m *RPCClient) Interaction(pluginInterface string, function string) (string, error) {
	var resp string
	err := m.client.Call("Plugin.Interaction", map[string]interface{}{
		"pluginInterface": pluginInterface,
		"function":        function,
	}, &resp)
	return resp, err
}

// Here is the RPC server that RPCClient talks to, conforming to
// the requirements of net/rpc
type RPCServer struct {
	Impl API
}

func (m *RPCServer) OnLoad(pluginInterface string, resp *[]byte) error {
	v, err := m.Impl.OnLoad(pluginInterface)
	*resp = v
	return err
}

func (m *RPCServer) OnUnload(pluginInterface string, resp *[]byte) error {
	v, err := m.Impl.OnUnload(pluginInterface)
	*resp = v
	return err
}

func (m *RPCServer) OnInstall(pluginInterface string, resp *[]byte) error {
	v, err := m.Impl.OnInstall(pluginInterface)
	*resp = v
	return err
}

func (m *RPCServer) OnUninstall(pluginInterface string, resp *[]byte) error {
	v, err := m.Impl.OnUninstall(pluginInterface)
	*resp = v
	return err
}

func (m *RPCServer) OnStart(pluginInterface string, resp *[]byte) error {
	v, err := m.Impl.OnStart(pluginInterface)
	*resp = v
	return err
}

func (m *RPCServer) OnStop(pluginInterface string, resp *[]byte) error {
	v, err := m.Impl.OnStop(pluginInterface)
	*resp = v
	return err
}

func (m *RPCServer) Interaction(args map[string]interface{}, resp *string) error {
	v, err := m.Impl.Interaction(args["pluginInterface"].(string), args["function"].(string))
	*resp = v
	return err
}
