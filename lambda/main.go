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
	Message string
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
	case "STATUS":
		groups, err := lib.Status(powerInput.Cluster)

		if err != nil {
			return s, err
		}

		message := ""
		for _, group := range groups {
			message += fmt.Sprintf("Group: %s - Desired: %v - Min: %v - Max: %v\n", *group.AutoScalingGroupName, *group.DesiredCapacity, *group.MinSize, *group.MaxSize)
		}
		fmt.Println(message)
		powerOutput.Message = message
	default:
		fmt.Println("Invalid input for Power. Must be \"ON\", \"OFF\" or \"STATUS\".")
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

	if !(powerInput.Power == "OFF" || powerInput.Power == "ON" || powerInput.Power == "STATUS") {
		return false
	}

	return true
}
