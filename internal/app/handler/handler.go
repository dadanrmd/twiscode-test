package handler

import (
	"twiscode-test/internal/app/commons"
	"twiscode-test/internal/app/service"
)

// HandlerOption option for handler, including all service
type HandlerOption struct {
	commons.Options
	*service.Services
}
