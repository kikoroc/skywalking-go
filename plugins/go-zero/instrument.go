/*
 * @Description:
 * @version: 1.0.0
 * @Author: kikoroc@gmail.com
 * @Date: 2023-09-12 15:40:20
 */

package gozero

import (
	"embed"

	"github.com/apache/skywalking-go/plugins/core/instrument"
)

//go:embed *
var fs embed.FS

type Instrument struct {
}

func NewInstrument() *Instrument {
	return &Instrument{}
}

func (i *Instrument) Name() string {
	return "go-zero"
}

func (i *Instrument) BasePackage() string {
	return "github.com/zeromicro/go-zero"
}

func (i *Instrument) VersionChecker(version string) bool {
	return true
}

func (i *Instrument) Points() []*instrument.Point {
	return []*instrument.Point{
		// http transport
		{
			PackagePath: "rest",
			At:          instrument.NewStructEnhance("Server"),
		},
		{
			PackagePath: "rest",
			At: instrument.NewStaticMethodEnhance("NewServer",
				instrument.WithArgsCount(2),
				instrument.WithArgType(0, "RestConf"), instrument.WithArgType(1, "...RunOption"),
				instrument.WithResultCount(2),
				instrument.WithResultType(0, "*Server"), instrument.WithResultType(1, "error")),
			Interceptor: "NewServerInterceptor",
		},
		/*{
			PackagePath: "rest",
			At: instrument.NewMethodEnhance("*Server", "Use",
				instrument.WithArgsCount(1),
				instrument.WithArgType(0, "Middleware"),
				instrument.WithResultCount(0)),
			Interceptor: "ServerMiddlewareInterceptor",
		},*/
	}
}

func (i *Instrument) FS() *embed.FS {
	return &fs
}
