package api

import (
	"github.com/plumos/breakfast/auth"
	"github.com/plumos/breakfast/web"
)

func HttpStart() {
	handle := web.NewHandle("/api")
	handle.PostOri("/login", login)
	registryUrls(handle)
	handle.Run(":80")
}

func registryUrls(handle *web.Handle) {
	handle.RG.Use(auth.JWTAuth())
}
