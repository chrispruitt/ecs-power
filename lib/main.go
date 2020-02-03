package lib

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var (
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
)

func PowerOn(cluster string) {
	min, max, desired := getAutoScaleValues(cluster)
	scaleCluster(cluster, min, max, desired)
}

func PowerOff(cluster string) {
	scaleCluster(cluster, 0, 0, 0)
}

func scaleCluster(cluster string, min, max, desired int64) {
	autoScalingGroupName := cluster + "-ecs"
	fmt.Printf("Updating autoscaling group \"%v\" desired: %v, min: %v, max: %v\n", autoScalingGroupName, desired, min, max)

	var autoScaleService = autoscaling.New(sess)

	input := &autoscaling.UpdateAutoScalingGroupInput{
		AutoScalingGroupName: aws.String(autoScalingGroupName),
		MinSize:              aws.Int64(min),
		MaxSize:              aws.Int64(max),
		DesiredCapacity:      aws.Int64(desired),
	}

	_, err := autoScaleService.UpdateAutoScalingGroup(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case autoscaling.ErrCodeScalingActivityInProgressFault:
				fmt.Println(autoscaling.ErrCodeScalingActivityInProgressFault, aerr.Error())
			case autoscaling.ErrCodeResourceContentionFault:
				fmt.Println(autoscaling.ErrCodeResourceContentionFault, aerr.Error())
			case autoscaling.ErrCodeServiceLinkedRoleFailure:
				fmt.Println(autoscaling.ErrCodeServiceLinkedRoleFailure, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}
	fmt.Println("Done.")
}

func getAutoScaleValues(cluster string) (int64, int64, int64) {
	min := getSSMParaValue(fmt.Sprintf("/%s/ecs-cluster/AUTOSCALE_MIN", cluster))
	max := getSSMParaValue(fmt.Sprintf("/%s/ecs-cluster/AUTOSCALE_MAX", cluster))
	desired := getSSMParaValue(fmt.Sprintf("/%s/ecs-cluster/AUTOSCALE_DESIRED", cluster))

	return stringToInt64(min), stringToInt64(max), stringToInt64(desired)
}

func getSSMParaValue(key string) string {
	var ssmService = ssm.New(sess)

	input := &ssm.GetParameterInput{
		Name: aws.String(key),
	}

	result, err := ssmService.GetParameter(input)

	if err != nil {
		fmt.Println(err.Error())
		panic(fmt.Sprintf("Unable to get ssm parameter \"%v\".", key))
	}

	return *result.Parameter.Value
}

func stringToInt64(s string) int64 {
	n, err := strconv.Atoi(s)

	if err != nil {
		panic(fmt.Sprintf("\"%v\" is not an integer.", s))
	}

	return int64(n)
}
