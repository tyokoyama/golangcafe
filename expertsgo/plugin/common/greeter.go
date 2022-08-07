package common

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type Greeter interface {
	Greet() (string, error)
}

type GreeterRPC struct { client *rpc.Client}

func (g *GreeterRPC) Greet() (string, error) {
	var resp string
	err := g.client.Call("Plugin.Greet", new(interface{}), &resp)
	if err != nil {
		return "", err
	}

	return resp, nil
}

type GreeterRPCServer struct {
	Impl Greeter
}

func (s *GreeterRPCServer) Greet(args interface{}, resp *string) error {
	var err error
	*resp, err = s.Impl.Greet()
	return err
}

var HandshakeConfig = plugin.HandshakeConfig {
	MagicCookieKey: "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

type GreeterPlugin struct {
	Impl Greeter
}

func (GreeterPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &GreeterRPC{client: c}, nil
}

func (p *GreeterPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &GreeterRPCServer{Impl: p.Impl}, nil
}