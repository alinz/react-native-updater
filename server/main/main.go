package main

import (
	"flag"
	"os"
	"syscall"
	"time"

	"github.com/alinz/react-native-updater/server/api/bundles"
	"github.com/alinz/react-native-updater/server/api/releases"
	"github.com/alinz/react-native-updater/server/config"
	"github.com/alinz/react-native-updater/server/lib/crypto"
	"github.com/alinz/react-native-updater/server/lib/logme"
	"github.com/alinz/react-native-updater/server/middleware"
	"github.com/pressly/cji"
	"github.com/zenazn/goji/graceful"
)

var (
	flags       = flag.NewFlagSet("server", flag.ExitOnError)
	configFile  = flags.String("config", "", "path to configuration file")
	generateKey = flags.Int("generate", -1, "generate key with size")
)

func main() {
	flags.Parse(os.Args[1:])

	//genearating public and private keys
	if *generateKey != -1 {
		crypto.Generate(*generateKey, "../key")
		return
	}

	if *configFile == "" {
		*configFile = os.Getenv("CONFIG")
	}

	configuration, err := config.New(*configFile)

	if err != nil {
		logme.Warn("configuration file is not found", *configFile)
		logme.Fatal(err)
	}

	route := cji.NewRouter()

	//global middlewares
	route.Use(middleware.LogHTTP)

	//global handlers
	route.Mount("/bundles", bundles.New())
	route.Mount("/releases", releases.New())

	graceful.AddSignal(syscall.SIGINT, syscall.SIGTERM)
	graceful.Timeout(10 * time.Second)

	logme.Info("Server started at", configuration.Server.Bind)

	graceful.ListenAndServe(configuration.Server.Bind, route)

	graceful.Wait()
}
