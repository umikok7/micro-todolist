package wrappers

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
)

type userWrapper struct {
	client.Client
}

// Call 通过 hystrix 熔断器对服务调用进行保护，防止因服务故障导致系统崩溃。
func (wrapper *userWrapper) Call(ctx context.Context, req client.Request, resp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()
	config := hystrix.CommandConfig{
		Timeout:                30000,
		RequestVolumeThreshold: 20,   // 熔断器请求阈值，默认20，意思是有20个请求才能进行错误百分比计算
		ErrorPercentThreshold:  50,   // 错误百分比，当错误超过百分比时，直接进行降级处理，直至熔断器再次 开启，默认50%
		SleepWindow:            5000, // 过多长时间，熔断器再次检测是否开启，单位毫秒ms（默认5秒）
	}
	hystrix.ConfigureCommand(cmdName, config) // 应用熔断器配置
	return hystrix.Do(cmdName, func() error {
		return wrapper.Client.Call(ctx, req, resp) // 正常的服务调用逻辑
	}, func(err error) error {
		return err // 降级处理逻辑，当服务调用失败时执行
	})
}

// NewUserWrapper 初始化Wrapper
func NewUserWrapper(c client.Client) client.Client {
	return &userWrapper{c}
}
