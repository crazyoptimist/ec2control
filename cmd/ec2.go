package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	StartStopInstances "start-stop-ec2/pkg/StartStopInstances"
)

func main() {
	instanceID := flag.String("i", "", "The ID of the instance to start or stop")
	state := flag.String("s", "", "The state to put the instance in: START or STOP")
	flag.Parse()

	if (*state != "START" && *state != "STOP") || *instanceID == "" {
		fmt.Println("You must supply a START or STOP state and an instance ID")
		fmt.Println("(-s START | STOP -i INSTANCE-ID")
		return
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := ec2.New(sess)

	if *state == "START" {
		err := StartStopInstances.StartInstance(svc, instanceID)
		if err != nil {
			fmt.Println("Got an error starting instance")
			fmt.Println(err)
			return
		}

		fmt.Println("Started instance with ID " + *instanceID)
	} else if *state == "STOP" {
		err := StartStopInstances.StopInstance(svc, instanceID)
		if err != nil {
			fmt.Println("Got an error stopping the instance")
			fmt.Println(err)
			return
		}

		fmt.Println("Stopped instance with ID " + *instanceID)
	}
}
