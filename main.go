package main

import (
	"os"

	"github.com/chrispruitt/ecs-power/cmd"
	ld "github.com/chrispruitt/ecs-power/lambda"
)

// Version of ecs-power. Overwritten during build
var Version = "development"

func main() {
	if os.Getenv("_LAMBDA_SERVER_PORT") != "" {
		ld.Start()
	} else {
		os.Setenv("AWS_SDK_LOAD_CONFIG", "true")
		cmd.Execute(Version)
	}
}
