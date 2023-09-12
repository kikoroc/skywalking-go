/*
 * @Description:
 * @version: 1.0.0
 * @Author: kikoroc@gmail.com
 * @Date: 2023-09-12 17:25:52
 */
package rest

import (
	"github.com/apache/skywalking-go/plugins/core/operator"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/zeromicro/go-zero/rest"

	"github.com/apache/skywalking-go/plugins/core/tracing"
)

var ignoreServerMiddlewareKey = "ignoreServerMiddleware"

type NewServerInterceptor struct {
}

func (n *NewServerInterceptor) BeforeInvoke(invocation operator.Invocation) error {
	return nil
}

func (n *NewServerInterceptor) AfterInvoke(invocation operator.Invocation, results ...interface{}) error {
	serverEnhanced, ok := results[0].(operator.EnhancedInstance)
	if !ok || serverEnhanced.GetSkyWalkingDynamicField() == true {
		return nil
	}
	server, ok := results[0].(*rest.Server)
	if !ok {
		return nil
	}
	// adding the middleware to the server
	tracing.SetRuntimeContextValue(ignoreServerMiddlewareKey, true)
	http.Middleware(serverMiddleware)(server)
	serverEnhanced.SetSkyWalkingDynamicField(true)
	return nil
}
