package filter

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"log"
)

// CommonMiddlewareFilter 请求头配置 /*
func CommonMiddlewareFilter() beego.FilterFunc{
	return func(ctx *context.Context) {
		ctx.ResponseWriter.Header().Set("Content-type","application/json")
	}
}

// OriginFilter 跨域配置 /*
func OriginFilter() beego.FilterFunc{
	return func(ctx *context.Context) {
		cors.Allow(&cors.Options{
			AllowAllOrigins: true,
			AllowMethods: []string{
				"GET",
				"POST",
				"PUT",
				"DELETE",
				"OPTIONS",
			},
			AllowHeaders: []string{
				"Origin",
				"Access-Control-Allow-Origin",
				"Access-Control-Allow-Headers",
				"Content-Type",
			},
			ExposeHeaders: []string{
				"Origin",
				"Access-Control-Allow-Origin",
				"Access-Control-Allow-Headers",
				"Content-Type",
			},
			AllowCredentials: true,
		})
	}
}

func RequestParameter() beego.FilterFunc{
	return func(ctx *context.Context) {
		log.Println("Request Method: ",ctx.Request.Method)
		log.Println("Request URI: ",ctx.Input.URI())
	}
}