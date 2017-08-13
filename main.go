package main

import (
	"log"
	"net/http"
	"time"

	cors "gopkg.in/gin-contrib/cors.v1"
	pprof "gopkg.in/gin-contrib/pprof.v1"

	"github.com/findcoo/gin-blueprint/api"
	"github.com/findcoo/gin-blueprint/conf"
	"github.com/gin-gonic/gin"
)

func main() {
	caseOne := conf.NewCaseOne("production")
	env := caseOne.Env
	api.NewApp(caseOne)
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	/*
		redisHost := env.GetString("Store")
		sessionStore, err := sessions.NewRedisStore(10, "tcp", redisHost, env.GetString("StorePassword"), []byte("secret"))

		if err != nil {
			log.Fatal(err)
		}
		log.Println("session type: redis")
	*/
	// r.Use(sessions.Sessions("appsession", sessionStore))

	r.Use(cors.Default())
	api.SetRouter(r)

	pprof.Register(r, nil)

	server := &http.Server{
		Addr:           env.GetString("Listen"),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   35 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Start server\n\tport: %s\n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
