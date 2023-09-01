package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Version godoc
//
//	@Summary		Get API version and release information
//	@Description	View Version app
//	@Tags			go-rest-template-app
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	vesionOut
//	@Router			/system/version [get]
func (srv *server) Version(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.JSON(http.StatusOK, &vesionOut{
		Code:           "0",
		Message:        "OK",
		ReleaseTime:    srv.config.Release.ReleaseTime,
		ReleaseVersion: srv.config.Release.ReleaseVersion,
	},
	)
}

type vesionOut struct {
	Code           string `json:"code"`
	ReleaseTime    string `json:"releaseTime"`
	ReleaseVersion string `json:"releaseVersion"`
	Message        string `json:"message"`
}
