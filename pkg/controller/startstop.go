package controller

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// StartInstance starts an Amazon EC2 instance.
// Inputs:
//
//	svc is an Amazon EC2 service client
//	instanceID is the ID of the instance
//
// Output:
//
//	If success, nil
//	Otherwise, an error from the call to StartInstances
func StartInstance(svc ec2iface.EC2API, instanceID *string) error {
	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			instanceID,
		},
		DryRun: aws.Bool(true),
	}
	_, err := svc.StartInstances(input)
	awsErr, ok := err.(awserr.Error)

	if ok && awsErr.Code() == "DryRunOperation" {
		// Set DryRun to be false to enable starting the instances
		input.DryRun = aws.Bool(false)
		_, err = svc.StartInstances(input)
		if err != nil {
			return err
		}

		return nil
	}

	return err
}

// StopInstance stops an Amazon EC2 instance.
// Inputs:
//
//	svc is an Amazon EC2 service client
//	instance ID is the ID of the instance
//
// Output:
//
//	If success, nil
//	Otherwise, an error from the call to StopInstances
func StopInstance(svc ec2iface.EC2API, instanceID *string) error {
	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			instanceID,
		},
		DryRun: aws.Bool(true),
	}
	_, err := svc.StopInstances(input)
	awsErr, ok := err.(awserr.Error)
	if ok && awsErr.Code() == "DryRunOperation" {
		input.DryRun = aws.Bool(false)
		_, err = svc.StopInstances(input)
		if err != nil {
			return err
		}

		return nil
	}

	return err
}
