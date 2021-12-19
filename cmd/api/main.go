package main

import (
	"context"
	"flag"
	. "github.com/v8tix/kit/app"
	. "github.com/v8tix/kit/handler"
	. "github.com/v8tix/kit/log"
	. "github.com/v8tix/kit/server/api"
	. "github.com/v8tix/kit/server/db"
	"strings"
	"sync"
	"time"
)

func main() {
	var app App
	var wg sync.WaitGroup
	var apiServer Server
	var mongoServer MongoServer
	var idleTimeout int
	var readTimeout int
	var writeTimeout int
	var cors Cors
	app.Wg = &wg
	ctx := context.TODO()
	clogs := NewLog()
	app.Log = clogs

	flag.IntVar(&apiServer.Port, "port", 4000, "API server port")
	flag.IntVar(&idleTimeout, "idle-timeout", 60, "Idle Timeout in seconds")
	flag.IntVar(&readTimeout, "read-timeout", 10, "Read Timeout in seconds")
	flag.IntVar(&writeTimeout, "write-timeout", 30, "Write Timeout in seconds")
	flag.StringVar(&apiServer.Env, "env", "development", "Environment (testing|development|staging|production)")
	flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
		cors.TrustedOrigins = strings.Fields(val)
		return nil
	})
	flag.StringVar(&mongoServer.ConFile, "con-file", "/home/v8tix/Public/Projects/tests/db/db-uri", "MongoDB connection URI file path")
	flag.StringVar(&mongoServer.Db, "db", "bender_dev", "Database name")
	flag.Parse()

	mongoServer.Ctx = ctx
	err := mongoServer.Connect()
	if err != nil {
		app.Log.E.Printf(err.Error())
		return
	}
	apiServer.App = app
	apiServer.IdleTimeout = time.Duration(idleTimeout) * time.Second
	apiServer.ReadTimeout = time.Duration(readTimeout) * time.Second
	apiServer.WriteTimeout = time.Duration(writeTimeout) * time.Second
	apiServer.Handler = Routes(NewHandler(app, cors))
	app.Log.I.Printf("database connection pool established")
	//repo := NewRepositoryI(app)

	if err := apiServer.Serve(); err != nil {
		app.Log.E.Fatal(err)
	}
}
