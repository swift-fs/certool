package cert

import (
	"cert-auto-cdn/internal/service"
	"context"

	cas20200407 "github.com/alibabacloud-go/cas-20200407/v2/client"
	cdn20180510 "github.com/alibabacloud-go/cdn-20180510/v5/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/frame/g"
)

type sCert struct {
	CdnClient *cdn20180510.Client
	CasClient *cas20200407.Client
}

func New(cdnC *cdn20180510.Client, casC *cas20200407.Client) *sCert {

	return &sCert{
		CdnClient: cdnC,
		CasClient: casC,
	}
}

func init() {
	cdnClient, err := CreateCdnClient(context.Background())
	if err != nil {
		panic(err)
	}

	casClient, err := CreateCasClient(context.Background())
	if err != nil {
		panic(err)
	}
	service.RegisterCert(New(cdnClient, casClient))

}

func CreateCdnClient(ctx context.Context) (_result *cdn20180510.Client, _err error) {
	config := &openapi.Config{
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID。
		AccessKeyId: tea.String(g.Cfg().MustGet(ctx, "aliyun.cdnAccessKeyId").String()),
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
		AccessKeySecret: tea.String(g.Cfg().MustGet(ctx, "aliyun.cdnAccessKeySecret").String()),
	}
	config.Endpoint = tea.String("cdn.aliyuncs.com")
	_result, _err = cdn20180510.NewClient(config)
	return _result, _err
}

func CreateCasClient(ctx context.Context) (_result *cas20200407.Client, _err error) {
	config := &openapi.Config{
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID。
		AccessKeyId: tea.String(g.Cfg().MustGet(ctx, "aliyun.certAccessKeyId").String()),
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
		AccessKeySecret: tea.String(g.Cfg().MustGet(ctx, "aliyun.certAccessKeySecret").String()),
	}

	config.Endpoint = tea.String("cas.aliyuncs.com")
	_result, _err = cas20200407.NewClient(config)
	return _result, _err
}
