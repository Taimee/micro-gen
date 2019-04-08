package cmd

import (
	ini "github.com/Taimee/micro-gen/lib/initialize"
	"github.com/micro/cli"
)

func Init(c *cli.Context) {
	ini.Exec()
}
