// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ICert interface {
		Del(ctx context.Context) (err error)
		Deploy(ctx context.Context, customDomain string, certPath string, isDelOld bool) (err error)
	}
)

var (
	localCert ICert
)

func Cert() ICert {
	if localCert == nil {
		panic("implement not found for interface ICert, forgot register?")
	}
	return localCert
}

func RegisterCert(i ICert) {
	localCert = i
}
