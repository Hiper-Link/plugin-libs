package shared

import (
	"github.com/Hiper-Link/plugin-libs/proto"

	"golang.org/x/net/context"
)

// GRPCClient is an implementation of KV that talks over RPC.
type GRPCClient struct {
	client proto.PluginClient
}

func (m *GRPCClient) OnLoad(pluginInterface string) ([]byte, error) {
	resp, err := m.client.OnLoad(context.Background(), &proto.EventsRequest{
		PluginInterface: pluginInterface,
	})
	if err != nil {
		return nil, err
	}

	return resp.Value, nil
}

func (m *GRPCClient) OnUnload(pluginInterface string) ([]byte, error) {
	resp, err := m.client.OnUnload(context.Background(), &proto.EventsRequest{
		PluginInterface: pluginInterface,
	})
	if err != nil {
		return nil, err
	}

	return resp.Value, nil
}

func (m *GRPCClient) OnInstall(pluginInterface string) ([]byte, error) {
	resp, err := m.client.OnInstall(context.Background(), &proto.EventsRequest{
		PluginInterface: pluginInterface,
	})
	if err != nil {
		return nil, err
	}

	return resp.Value, nil
}

func (m *GRPCClient) OnUninstall(pluginInterface string) ([]byte, error) {
	resp, err := m.client.OnUninstall(context.Background(), &proto.EventsRequest{
		PluginInterface: pluginInterface,
	})
	if err != nil {
		return nil, err
	}

	return resp.Value, nil
}

func (m *GRPCClient) OnStart(pluginInterface string) ([]byte, error) {
	resp, err := m.client.OnStart(context.Background(), &proto.EventsRequest{
		PluginInterface: pluginInterface,
	})
	if err != nil {
		return nil, err
	}

	return resp.Value, nil
}

func (m *GRPCClient) OnStop(pluginInterface string) ([]byte, error) {
	resp, err := m.client.OnStop(context.Background(), &proto.EventsRequest{
		PluginInterface: pluginInterface,
	})
	if err != nil {
		return nil, err
	}

	return resp.Value, nil
}

func (m *GRPCClient) Interaction(pluginInterface string, function string) (string, error) {
	resp, err := m.client.Interaction(context.Background(), &proto.InteractionRequest{
		PluginInterface: pluginInterface,
		Function:        function,
	})
	if err != nil {
		return "", err
	}

	return resp.Value, nil
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl API
}

func (m *GRPCServer) OnLoad(
	ctx context.Context,
	req *proto.EventsRequest) (*proto.EventsResponse, error) {
	v, err := m.Impl.OnLoad(req.PluginInterface)
	return &proto.EventsResponse{Value: v}, err
}

func (m *GRPCServer) OnUnload(
	ctx context.Context,
	req *proto.EventsRequest) (*proto.EventsResponse, error) {
	v, err := m.Impl.OnUnload(req.PluginInterface)
	return &proto.EventsResponse{Value: v}, err
}

func (m *GRPCServer) OnInstall(
	ctx context.Context,
	req *proto.EventsRequest) (*proto.EventsResponse, error) {
	v, err := m.Impl.OnInstall(req.PluginInterface)
	return &proto.EventsResponse{Value: v}, err
}

func (m *GRPCServer) OnUninstall(
	ctx context.Context,
	req *proto.EventsRequest) (*proto.EventsResponse, error) {
	v, err := m.Impl.OnUninstall(req.PluginInterface)
	return &proto.EventsResponse{Value: v}, err
}

func (m *GRPCServer) OnStart(
	ctx context.Context,
	req *proto.EventsRequest) (*proto.EventsResponse, error) {
	v, err := m.Impl.OnStart(req.PluginInterface)
	return &proto.EventsResponse{Value: v}, err
}

func (m *GRPCServer) OnStop(
	ctx context.Context,
	req *proto.EventsRequest) (*proto.EventsResponse, error) {
	v, err := m.Impl.OnStop(req.PluginInterface)
	return &proto.EventsResponse{Value: v}, err
}

func (m *GRPCServer) Interaction(
	ctx context.Context,
	req *proto.InteractionRequest) (*proto.InteractionResponse, error) {
	v, err := m.Impl.Interaction(req.PluginInterface, req.Function)
	return &proto.InteractionResponse{Value: v}, err
}
