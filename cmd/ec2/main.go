package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/viper"

	"ec2control/internal/config"
	"ec2control/pkg/controller"
)

func main() {
	err := loadConfig()
	if err != nil {
		fmt.Println("Got an error reading the config file(config.toml). ", err)
		return
	}

	instanceID := viper.GetString("ec2_instance.id")

	start := flag.Bool("start", false, "Start the instance.")
	stop := flag.Bool("stop", false, "Stop the instance.")
	flag.Parse()

	if !*start && !*stop {
		fmt.Println("You must supply a <start> or <stop> flag.")
		fmt.Println("(-start | -stop)")
		return
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := ec2.New(sess)

	if *start {
		err := controller.StartInstance(svc, &instanceID)
		if err != nil {
			fmt.Println("Got an error starting instance")
			fmt.Println(err)
			return
		}

		fmt.Println("Started instance with ID " + instanceID)
	} else if *stop {
		err := controller.StopInstance(svc, &instanceID)
		if err != nil {
			fmt.Println("Got an error stopping the instance")
			fmt.Println(err)
			return
		}

		fmt.Println("Stopped instance with ID " + instanceID)
	}
}

func loadConfig() error {
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
			fmt.Println("Configuration file not found. Please create config.toml using the below example.", exampleConfig)
			return err
		} else {
			return err
		}
	}

	return nil
}
