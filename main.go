package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	flag "github.com/spf13/pflag"
)

func Session() (sess *session.Session) {
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	}))
	return
}

var (
	helpFlag   = flag.BoolP("help", "h", false, "show help message")
	instanceid = flag.StringP("instanceid", "i", "", "ec2 instanceid")
	publicipFlag = flag.BoolP("publicip", "p", false, "show public ip address only")
)

func main() {

	flag.Parse()

	if *helpFlag {
		flag.PrintDefaults()
		return
	}

	svc := ec2.New(Session())
	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			aws.String(*instanceid),
		},
	}
	result, err := svc.DescribeInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}

	if *publicipFlag {
		fmt.Println(*result.Reservations[0].Instances[0].NetworkInterfaces[0].Association.PublicIp)
		return
	}

	fmt.Println(result)

}
