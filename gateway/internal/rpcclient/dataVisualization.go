package rpcclient

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"

	"github.com/zhufuyi/sponge/pkg/consulcli"
	"github.com/zhufuyi/sponge/pkg/etcdcli"
	"github.com/zhufuyi/sponge/pkg/grpc/grpccli"
	"github.com/zhufuyi/sponge/pkg/logger"
	"github.com/zhufuyi/sponge/pkg/nacoscli"
	"github.com/zhufuyi/sponge/pkg/servicerd/registry/consul"
	"github.com/zhufuyi/sponge/pkg/servicerd/registry/etcd"
	"github.com/zhufuyi/sponge/pkg/servicerd/registry/nacos"

	"gateway/internal/config"
)

var (
	dataVisualizationConn *grpc.ClientConn
	dataVisualizationOnce sync.Once
)

// NewDataVisualizationRPCConn instantiate rpc client connection
func NewDataVisualizationRPCConn() {
	cfg := config.Get()

	serverName := "dataVisualization"
	var grpcClientCfg config.GrpcClient
	for _, cli := range cfg.GrpcClient {
		if strings.EqualFold(cli.Name, serverName) {
			grpcClientCfg = cli
			break
		}
	}
	if grpcClientCfg.Name == "" {
		panic(fmt.Sprintf("not found grpc service name '%v' in configuration file(yaml), "+
			"please add gprc service configuration in the configuration file(yaml) under the field grpcClient.", serverName))
	}

	var cliOptions = []grpccli.Option{
		grpccli.WithEnableRequestID(),
		grpccli.WithEnableLog(logger.Get()),
	}

	//if grpcClientCfg.Timeout > 0 {
	//	cliOptions = append(cliOptions, grpccli.WithTimeout(time.Second*time.Duration(grpcClientCfg.Timeout)))
	//}

	// load balance
	if grpcClientCfg.EnableLoadBalance {
		cliOptions = append(cliOptions, grpccli.WithEnableLoadBalance())
	}

	// secure
	cliOptions = append(cliOptions, grpccli.WithSecure(
		grpcClientCfg.ClientSecure.Type,
		grpcClientCfg.ClientSecure.ServerName,
		grpcClientCfg.ClientSecure.CaFile,
		grpcClientCfg.ClientSecure.CertFile,
		grpcClientCfg.ClientSecure.KeyFile,
	))

	// token
	cliOptions = append(cliOptions, grpccli.WithToken(
		grpcClientCfg.ClientToken.Enable,
		grpcClientCfg.ClientToken.AppID,
		grpcClientCfg.ClientToken.AppKey,
	))

	// if service discovery is not used, connect directly to the rpc service using the ip and port
	endpoint := fmt.Sprintf("%s:%d", grpcClientCfg.Host, grpcClientCfg.Port)

	isUseDiscover := false
	switch grpcClientCfg.RegistryDiscoveryType {
	// discovering services using consul
	case "consul":
		endpoint = "discovery:///" + grpcClientCfg.Name // connecting to grpc services by service name
		cli, err := consulcli.Init(cfg.Consul.Addr, consulcli.WithWaitTime(time.Second*5))
		if err != nil {
			panic(fmt.Sprintf("consulcli.Init error: %v, addr: %s", err, cfg.Consul.Addr))
		}
		iDiscovery := consul.New(cli)
		cliOptions = append(cliOptions, grpccli.WithDiscovery(iDiscovery))
		isUseDiscover = true
	// discovering services using etcd
	case "etcd":
		endpoint = "discovery:///" + grpcClientCfg.Name // Connecting to grpc services by service name
		cli, err := etcdcli.Init(cfg.Etcd.Addrs, etcdcli.WithDialTimeout(time.Second*5))
		if err != nil {
			panic(fmt.Sprintf("etcdcli.Init error: %v, addr: %v", err, cfg.Etcd.Addrs))
		}
		iDiscovery := etcd.New(cli)
		cliOptions = append(cliOptions, grpccli.WithDiscovery(iDiscovery))
		isUseDiscover = true
	// discovering services using nacos
	case "nacos":
		// example: endpoint = "discovery:///serverName.scheme"
		endpoint = "discovery:///" + grpcClientCfg.Name + ".grpc" // Connecting to grpc services by service name
		cli, err := nacoscli.NewNamingClient(
			cfg.NacosRd.IPAddr,
			cfg.NacosRd.Port,
			cfg.NacosRd.NamespaceID)
		if err != nil {
			panic(fmt.Sprintf("nacoscli.NewNamingClient error: %v, ipAddr: %s, port: %d",
				err, cfg.NacosRd.IPAddr, cfg.NacosRd.Port))
		}
		iDiscovery := nacos.New(cli)
		cliOptions = append(cliOptions, grpccli.WithDiscovery(iDiscovery))
		isUseDiscover = true
	}

	if cfg.App.EnableTrace {
		cliOptions = append(cliOptions, grpccli.WithEnableTrace())
	}
	if cfg.App.EnableCircuitBreaker {
		cliOptions = append(cliOptions, grpccli.WithEnableCircuitBreaker())
	}
	if cfg.App.EnableMetrics {
		cliOptions = append(cliOptions, grpccli.WithEnableMetrics())
	}

	msg := "dialing grpc server"
	if isUseDiscover {
		msg += " with discovery from " + grpcClientCfg.RegistryDiscoveryType
	}
	logger.Info(msg, logger.String("name", serverName), logger.String("endpoint", endpoint))

	var err error
	dataVisualizationConn, err = grpccli.Dial(context.Background(), endpoint, cliOptions...)
	if err != nil {
		panic(fmt.Sprintf("dial rpc server failed: %v, name: %s, endpoint: %s", err, serverName, endpoint))
	}
	logger.Infof("dialed grpc server (%s) successfully", serverName+", "+endpoint)
}

// GetDataVisualizationRPCConn get client conn
func GetDataVisualizationRPCConn() *grpc.ClientConn {
	if dataVisualizationConn == nil {
		dataVisualizationOnce.Do(func() {
			NewDataVisualizationRPCConn()
		})
	}

	return dataVisualizationConn
}

// CloseDataVisualizationRPCConn Close tears down the ClientConn and all underlying connections.
func CloseDataVisualizationRPCConn() error {
	if dataVisualizationConn == nil {
		return nil
	}

	return dataVisualizationConn.Close()
}
