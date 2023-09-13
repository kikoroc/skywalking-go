package rest

import (
	"fmt"
	"net/http"

	"github.com/apache/skywalking-go/plugins/core/log"
	"github.com/apache/skywalking-go/plugins/core/tracing"
)

var serverMiddleware = func(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		span, err := tracing.CreateEntrySpan(fmt.Sprintf("%s:%s", r.Method, r.URL.Path), func(key string) (string, error) {
			return r.Header.Get(key), nil
		}, tracing.WithComponent(5020),
			tracing.WithLayer(tracing.SpanLayerRPCFramework),
			tracing.WithTag("transport", "HTTP"),
			tracing.WithTag(tracing.TagURL, r.Host+r.URL.Path),
			tracing.WithTag(tracing.TagHTTPMethod, r.Method))
		if err != nil {
			log.Warnf("cannot create entry span: %v", err)
			next(w, r)
			return
		}

		defer span.End()

		next(w, r)
	}
}
