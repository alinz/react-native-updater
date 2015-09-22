package main

import (
	"flag"
	"os"
	"syscall"
	"time"

	"github.com/alinz/react-native-updater/server/api/bundles"
	"github.com/alinz/react-native-updater/server/api/releases"
	"github.com/alinz/react-native-updater/server/config"
	"github.com/alinz/react-native-updater/server/lib/logme"
	"github.com/pressly/cji"
	"github.com/zenazn/goji/graceful"
)

var (
	flags      = flag.NewFlagSet("server", flag.ExitOnError)
	configFile = flags.String("config", "", "path to configuration file")
)

func main() {
	flags.Parse(os.Args[1:])

	if *configFile == "" {
		*configFile = os.Getenv("CONFIG")
	}

	configuration, err := config.New(*configFile)

	if err != nil {
		logme.Warn("configuration file is not found", *configFile)
		logme.Fatal(err)
	}

	route := cji.NewRouter()

	route.Mount("/bundles", bundles.New())
	route.Mount("/releases", releases.New())

	graceful.AddSignal(syscall.SIGINT, syscall.SIGTERM)
	graceful.Timeout(10 * time.Second)

	logme.Info("Server started at", configuration.Server.Bind)

	graceful.ListenAndServe(configuration.Server.Bind, route)

	graceful.Wait()
}