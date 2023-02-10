package pluginmod

import (
	"context"
	"fmt"

	"github.com/grafana/grafana-plugin-sdk-go/backend"

	"github.com/grafana/dskit/services"

	"github.com/grafana/grafana/pkg/plugins"
)

var _ services.Service = (*Client)(nil)
var _ PluginManager = (*Client)(nil)

type Client struct {
	*services.BasicService
}

func NewClient() *Client {
	c := &Client{}
	c.BasicService = services.NewBasicService(c.start, c.run, c.stop)
	fmt.Println("Creating client service...")
	return c
}

func (c *Client) start(ctx context.Context) error {
	fmt.Println("Starting client...")
	return nil
}

func (c *Client) run(ctx context.Context) error {
	fmt.Println("Running client...")
	<-ctx.Done()
	return nil
}

func (c *Client) stop(err error) error {
	fmt.Println("Stopping client...")
	return err
}

func (c *Client) Add(ctx context.Context, pluginID, version string, opts plugins.CompatOpts) error {
	panic("implement me")
}

func (c *Client) Remove(ctx context.Context, pluginID string) error {
	panic("implement me")
}

func (c *Client) Plugin(ctx context.Context, pluginID string) (plugins.PluginDTO, bool) {
	panic("implement me")
}

func (c *Client) Plugins(ctx context.Context, pluginTypes ...plugins.Type) []plugins.PluginDTO {
	panic("implement me")
}

func (c *Client) Renderer(ctx context.Context) *plugins.Plugin {
	panic("implement me")
}

func (c *Client) SecretsManager(ctx context.Context) *plugins.Plugin {
	panic("implement me")
}

func (c *Client) Routes() []*plugins.StaticRoute {
	panic("implement me")
}

func (c *Client) PluginErrors() []*plugins.Error {
	panic("implement me")
}

func (c *Client) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	panic("implement me")
}

func (c *Client) CallResource(ctx context.Context, req *backend.CallResourceRequest, sender backend.CallResourceResponseSender) error {
	panic("implement me")
}

func (c *Client) CheckHealth(ctx context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	panic("implement me")
}

func (c *Client) CollectMetrics(ctx context.Context, req *backend.CollectMetricsRequest) (*backend.CollectMetricsResult, error) {
	panic("implement me")
}

func (c *Client) SubscribeStream(ctx context.Context, req *backend.SubscribeStreamRequest) (*backend.SubscribeStreamResponse, error) {
	panic("implement me")
}

func (c *Client) PublishStream(ctx context.Context, req *backend.PublishStreamRequest) (*backend.PublishStreamResponse, error) {
	panic("implement me")
}

func (c *Client) RunStream(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender) error {
	panic("implement me")
}
