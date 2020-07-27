package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func Session() (sess *session.Session) {
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	}))
	return
}

func main() {

	instanceid := flag.String("i", "", "description")
	flag.Parse()
	//fmt.Println(*instanceid)

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

	fmt.Println(result)

}
