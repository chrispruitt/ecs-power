package ld

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/chrispruitt/ecs-power/lib"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	ErrInvalidInputProvided = errors.New("invalid input - ensure \"Cluster\" and \"Power\" is provided\"")
)

type PowerInput struct {
	Cluster string
	Power   string
}

type PowerOutput struct {
	Success bool
}

func Start() {
	lambda.Start(Handler)
}

func Handler(powerInput PowerInput) (s string, err error) {
	var powerOutput PowerOutput
	if !inputValidation(powerInput) {
		return s, ErrInvalidInputProvided
	}

	fmt.Printf("Update autoscaling requested for group: %v\n", powerInput.Cluster)

	switch powerInput.Power {
	case "ON":
		err = lib.PowerOn(powerInput.Cluster)
	case "OFF":
		err = lib.PowerOff(powerInput.Cluster)
	default:
		fmt.Println("Invalid input for Power. Must be \"ON\" or \"OFF\".")
	}

	if err != nil {
		return s, err
	}

	powerOutput.Success = true

	res, err := json.Marshal(powerOutput)
	s = string(res)

	return s, err
}

func inputValidation(powerInput PowerInput) bool {
	if powerInput.Cluster == "" {
		return false
	}

	if powerInput.Power != "OFF" && powerInput.Power != "ON" {
		return false
	}

	return true
}
