package cert

import (
	"context"

	cas20200407 "github.com/alibabacloud-go/cas-20200407/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sCert) Del(ctx context.Context) (err error) {
	g.Log().Infof(ctx, "删除已过期证书\n")

	client := s.CasClient

	listUserCertificateOrderRequest := &cas20200407.ListUserCertificateOrderRequest{
		OrderType: tea.String("UPLOAD"),
		ShowSize:  tea.Int64(10000),
		Status:    tea.String("EXPIRED"),
	}

	err = func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		// 查询过期证书列表
		res, err := client.ListUserCertificateOrder(listUserCertificateOrderRequest)
		if err != nil {
			return err
		}

		g.Log().Debug(ctx, "过期证书列表>>", res)
		g.Log().Info(ctx, "开始删除过期证书")

		for _, v := range res.Body.CertificateOrderList {
			// 删除过期证书
			deleteUserCertificateRequest := &cas20200407.DeleteUserCertificateRequest{
				CertId: v.CertificateId,
			}
			runtime := &util.RuntimeOptions{}
			defer func() {
				if r := tea.Recover(recover()); r != nil {
					_e = r
				}
			}()

			_, err = client.DeleteUserCertificateWithOptions(deleteUserCertificateRequest, runtime)
			if err != nil {
				g.Log().Error(ctx, err)
			} else {
				g.Log().Info(ctx, "已删除过期证书 >>", v.Name, v.CertificateId)
			}

		}

		return nil
	}()

	return err

}
