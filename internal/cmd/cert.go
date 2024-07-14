// This file is auto-generated, don't edit it. Thanks.
package cmd

import (
	"cert-auto-cdn/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

const (
	// 需要部署的cdn域名
	CdnDomian = "cdn.oss.tobmm.com"
)

var Cert = gcmd.Command{
	Name:  "cert",
	Usage: "cert -domain=xxx.example.com -path=证书绝对路径 -del=是否删除过期证书",
	Brief: "cert",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		customDomain := parser.GetOpt("domain", CdnDomian).String()
		err = g.Validator().Data(customDomain).Rules("domain").Messages("请输入合法域名格式,格式: xxx.example.com").Run(ctx)
		if err != nil {
			g.Log().Error(ctx, err)
			return
		}
		certPath := parser.GetOpt("path", "").String()
		err = g.Validator().Data(certPath).Rules("min-length:1").Messages("请输入证书所在的绝对路径").Run(ctx)
		if err != nil {
			g.Log().Error(ctx, err)
			return
		}
		isDelOld := parser.GetOpt("del", "true").Bool()
		g.Log().Info(ctx, "为域名:", customDomain, "部署证书")
		g.Log().Info(ctx, "证书路径:", certPath)
		err = service.Cert().Deploy(ctx, customDomain, certPath, isDelOld)
		if err != nil {
			g.Log().Error(ctx, err)
		}
		return
	},
}
