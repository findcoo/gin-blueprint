package tests

import (
	"testing"

	"github.com/findcoo/gin-blueprint/api"
	"github.com/findcoo/gin-blueprint/conf"
	"github.com/findcoo/goTest"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	caseOne = conf.NewCaseOne("develop")
	App     = setApp()
)

func setApp() *gin.Engine {
	r := testSuit.GetGinEngine()
	api.NewApp(caseOne)
	api.SetRouter(r)
	return r
}

func TestPing(*testing.T) {
	suit := &testSuit.TestSuit{
		Router: App,
		Method: "GET",
		URL:    "/ping",
	}

	resp := suit.Do()
	assert.Equal(t, 200, resp.Code)
}
