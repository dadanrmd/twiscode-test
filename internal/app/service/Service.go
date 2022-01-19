package service

import (
	"twiscode-test/internal/app/commons"
	"twiscode-test/internal/app/service/funcService"
)

// Option anything any service object needed
type Option struct {
	commons.Options
}

type Services struct {
	FuncService funcService.IFuncService
}
