package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/viper"

	"ec2control/internal/config"
	"ec2control/pkg/controller"
)

func main() {
	loadConfig()

	instanceID := flag.String("i", "", "The ID of the instance to start or stop")
	state := flag.String("s", "STOP", "The state to put the instance in: START or STOP")
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
		err := controller.StartInstance(svc, instanceID)
		if err != nil {
			fmt.Println("Got an error starting instance")
			fmt.Println(err)
			return
		}

		fmt.Println("Started instance with ID " + *instanceID)
	} else if *state == "STOP" {
		err := controller.StopInstance(svc, instanceID)
		if err != nil {
			fmt.Println("Got an error stopping the instance")
			fmt.Println(err)
			return
		}

		fmt.Println("Stopped instance with ID " + *instanceID)
	}
}

func loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	viper.AddConfigPath(".")
	if runtime.GOOS == "linux" {
		viper.AddConfigPath("$HOME/.config/ec2")
	}
	if runtime.GOOS == "windows" {
		roamingDir, _ := os.UserConfigDir()
		viper.AddConfigPath(roamingDir + "\\ec2")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Assign to a new var because of escaped new line error in println
			exampleConfig := config.ExampleConfig
			log.Fatalln("Configuration file not found. Please create config.toml using the below example.", exampleConfig)
		} else {
			log.Fatalln("Error occurred while reading the config file: ", err)
		}
	}
	log.Println(viper.Get("aws_profile"))
}
