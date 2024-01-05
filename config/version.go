package config

import (
	"net/http"
	"prototype/lib/env"
	"time"

	"github.com/gin-gonic/gin"
)

type Version struct {
	VersionRelease string `json:"version_release"`
	ServiceName    string
	ServiceType    string
	GitCommit      string
	BranchName     string
	DateTime       string
	VersionType    string
	Notes          string
}

var (
	DateTime string
)

func init() {
	DateTime = time.Now().Format("2006-01-02 15:04:05 Z0700")
}

func HandleVersion(c *gin.Context) {
	res := Version{
		VersionRelease: env.String("Version.ReleaseVersion", ""),
		VersionType:    env.String("Version.VersionType", ""),
		ServiceName:    env.String("MainSetup.ServiceName", ""),
		ServiceType:    env.String("MainSetup.ServiceType", ""),
		Notes:          env.String("Version.Notes", ""),
		BranchName:     env.String("Version.BranchName", ""),
		DateTime:       DateTime,
	}

	c.JSON(http.StatusOK, res)
}
