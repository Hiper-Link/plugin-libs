package shared

import (
	"net/rpc"

	"github.com/Hiper-Link/plugin-libs/proto"

	"github.com/Hiper-Link/go-plugin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// 传入数据
type PluginInterface struct {
	Config          string
	EnvironmentInfo struct {
		BuildType string
		Platform  string
		Arch      string
	}
}

// 插件可调用接口
type API interface {
	OnLoad(config string) ([]byte, error)                       // 加载
	OnUnload(config string) ([]byte, error)                     // 取消启用
	OnInstall(config string) ([]byte, error)                    // 安装插件
	OnUninstall(config string) ([]byte, error)                  // 卸载插件
	OnStart(config string) ([]byte, error)                      // HL 启动
	OnStop(config string) ([]byte, error)                       // HL 停止
	Interaction(config string, function string) (string, error) // 前后端交互
	// Config(config string) ([]byte, error)      // 配置文件读取写入 返回 !null 将写入配置
}

// netRPC 配置
type Plugin struct {
	Impl API
}

// netRPC Server 配置
func (p *Plugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RPCServer{Impl: p.Impl}, nil
}

// netRPC Client 配置
func (*Plugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RPCClient{client: c}, nil
}

// gRPC 配置
type GRPCPlugin struct {
	plugin.Plugin
	Impl API
}

// gRPC Server 配置
func (p *GRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterPluginServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

// gRPC Client 配置
func (p *GRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: proto.NewPluginClient(c)}, nil
}
