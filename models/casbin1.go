package main

import (
	"github.com/casbin/casbin"
	"github.com/gin-contrib/authz"
	"github.com/gin-gonic/gin"
)

func main() {
	// load the casbin model and policy from files, database is also supported.
	e := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")

	// define your router, and use the Casbin authz middleware.
	// the access that is denied by authz will return HTTP 403 error.
	router := gin.New()
	router.Use(authz.NewAuthorizer(e))

}
