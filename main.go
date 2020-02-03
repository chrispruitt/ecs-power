package main

import (
	"github.com/chrispruitt/ecs-power/cmd"
	"github.com/chrispruitt/ecs-power/lambda"
	"os"
)

// Version of ecs-power. Overwritten during build
var Version = "development"

func main() {
	if os.Getenv("_LAMBDA_SERVER_PORT") != "" {
		lambda.Start()
	} else {
		os.Setenv("AWS_SDK_LOAD_CONFIG", "true")
		cmd.Execute(Version)
	}
}
