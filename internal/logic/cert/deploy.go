package cert

import (
	"cert-auto-cdn/internal/service"
	"context"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/grand"

	cdn20180510 "github.com/alibabacloud-go/cdn-20180510/v5/client"

	util "github.com/alibabacloud-go/tea-utils/v2/service"
)

func (s *sCert) Deploy(ctx context.Context, customDomain string, certPath string, isDelOld bool) (err error) {
	client := s.CdnClient
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	// 上传证书
	certName := gstr.Join([]string{customDomain, grand.S(6)}, "-")
	g.Log().Debug(ctx, "certName:", certName)
	setCdnDomainSSLCertificateRequest := &cdn20180510.SetCdnDomainSSLCertificateRequest{
		SSLProtocol: tea.String("on"),
		DomainName:  tea.String(customDomain),
		CertName:    tea.String(certName),
		CertType:    tea.String("upload"),
		SSLPub:      tea.String(gfile.GetContents(certPath + "/fullchain.pem")),
		SSLPri:      tea.String(gfile.GetContents(certPath + "/privkey.pem")),
	}
	runtime := &util.RuntimeOptions{}
	err = func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		_, err = client.SetCdnDomainSSLCertificateWithOptions(setCdnDomainSSLCertificateRequest, runtime)
		if err != nil {
			return err
		}

		return nil
	}()

	if err == nil && isDelOld {
		service.Cert().Del(ctx)
	}
	return err
}
