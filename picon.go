package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	logger "github.com/Sirupsen/logrus"
	config "github.com/mgreau/picon/config"
	pipeline "github.com/mgreau/picon/pipeline"
	types "github.com/mgreau/picon/types"
	workers "github.com/mgreau/picon/workers"
	cli "gopkg.in/urfave/cli.v1"
)

var (
	// Version is the version of the software
	Version string
	// BuildStmp is the build date
	BuildStmp string
	// GitHash is the git build hash
	GitHash string

	logLevel = "warning"
	// path to config file
	cfgFile = "picon.yaml"

	// The pipeline description
	picon *types.Picon
)

// preload initializes any global options and configuration
// before the main or sub commands are run.
func preload(c *cli.Context) (err error) {
	if c.GlobalBool("debug") {
		logger.SetLevel(logger.DebugLevel)
	}

	return nil
}

func main() {

	// set timezone as UTC for bson/json time marshalling
	time.Local = time.UTC

	// new app
	app := cli.NewApp()
	app.Name = "picon"
	app.Usage = "Your build tools companion"

	timeStmp, err := strconv.Atoi(BuildStmp)
	if err != nil {
		timeStmp = 0
	}
	app.Version = Version + ", build on " + time.Unix(int64(timeStmp), 0).String() + ", git hash " + GitHash
	app.Author = "@mgreau"
	app.Email = "contact@mgreau.com"
	app.Before = preload

	// command line flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Value: logLevel,
			Name:  "logl",
			Usage: "Set the output log level (debug, info, warning, error)",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"rm"},
			Usage:   "Initialize the project with picon.yml file",
			Action: func(c *cli.Context) error {
				workers.ListContainers(false)
				return nil
			},
		},
		{
			Name:    "show",
			Aliases: []string{"sw"},
			Usage:   "Display the pipeline content ",
			Action: func(c *cli.Context) error {

				byt, err := ioutil.ReadFile(cfgFile)
				if err != nil {
					return fmt.Errorf("error reading configuration: %s", err)
				}
				picon, err = config.Parse(byt)
				if err != nil {
					return fmt.Errorf("error reading configuration: %s", err)
				}
				pipeline.ListWorkers(picon)
				return nil
			},
		},
		{
			Name:    "worker",
			Aliases: []string{"wk"},
			Usage:   "Display workers defined for the pipeline",
			Action: func(c *cli.Context) error {
				workers.ListContainers(false)
				return nil
			},
		},
		{
			Name:    "task",
			Aliases: []string{"tk"},
			Usage:   "Display tasks defined in the pipeline",
			Action: func(c *cli.Context) error {
				workers.ListContainers(false)
				return nil
			},
		},
	}

	// run the appcd
	err = app.Run(os.Args)
	if err != nil {
		logger.Fatalf("Run error %q\n", err)
	}
}
