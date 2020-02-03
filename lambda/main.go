package lambda

import (
	"errors"
	"log"

	lib "github.com/chrispruitt/ecs-power/lib"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	ErrInvalidInputProvided = errors.New("invalid input - ensure \"cluster\" and \"up\" is provided\"")
)

type InputOpts struct {
	Cluster string
	Power   string
}

func Start() {
	lambda.Start(Handler)
}

func Handler(inputOpts InputOpts) (err error) {

	if inputOpts.Cluster == "" {
		return ErrInvalidInputProvided
	}

	if inputOpts.Power == "OFF" || inputOpts.Power == "ON" {
		return ErrInvalidInputProvided
	}

	log.Printf("Update autoscaling requested for group: %v\n", inputOpts.Cluster)

	switch inputOpts.Power {
	case "ON":
		lib.PowerOn(inputOpts.Cluster)
	case "OFF":
		lib.PowerOff(inputOpts.Cluster)
	default:
		panic("Invalid input for Power. Must be \"ON\" or \"OFF\".")
	}

	return nil
}
