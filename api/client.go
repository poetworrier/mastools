package api

import (
	"github.com/imroc/req/v3"
)

func NewClient(origin string, accessToken string, debug bool) (*req.Client, func()) {
	// TODO: could do input validtion on params
	c := req.C().
		SetBaseURL(origin).
		SetCommonBearerAuthToken(accessToken)
	if debug {
		c = c.EnableDumpAll().EnableDebugLog()
	}
	return c, func() {
		c.CloseIdleConnections()
	}
}
