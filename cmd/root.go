/*
Copyright © 2024 crazyoptimist <hey@crazyoptimist.net>
*/
package cmd

import (
	"log"
	"os"

	"ec2control/pkg/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ec2control",
	Short: "Control AWS EC2 VMs",
	Long: `Turn on and off AWS EC2 instances as needed.

#############################################################################
To use this application, you need to create AWS credentials:
https://docs.aws.amazon.com/keyspaces/latest/devguide/access.credentials.html

Your key pair needs to have permissions:
"ec2:StartInstances"
"ec2:StopInstances"
#############################################################################
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var cfgFile string

const DEFAULT_CONFIG_FILE_NAME = "ec2control"

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/ec2control.toml or $APPDATA\\ec2control\\ec2control.toml on Windows)")

}

func initConfig() {
	// Don't forget to read config either from cfgFile or from default config directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		configDir, err := utils.GetConfigDir()
		if err != nil {
			log.Fatalln(err)
		}
		viper.AddConfigPath(configDir)

		viper.SetConfigName(DEFAULT_CONFIG_FILE_NAME)
		viper.SetConfigType("toml")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Config file not found, please create one using 'config' command.")
	}
}

func getAWSProfile() *AWSProfile {
	return &AWSProfile{
		Region:             viper.GetString("AWSProfile.Region"),
		AWSAccessKeyId:     viper.GetString("AWSProfile.AWSAccessKeyId"),
		AWSSecretAccessKey: viper.GetString("AWSProfile.AWSSecretAccessKey"),
	}
}

func getEC2Client(awsProfile AWSProfile) *ec2.EC2 {
	sess := session.Must(
		session.NewSession(&aws.Config{
			Region: aws.String(awsProfile.Region),
			Credentials: credentials.NewStaticCredentials(
				awsProfile.AWSAccessKeyId,
				awsProfile.AWSSecretAccessKey,
				"",
			),
		}),
	)

	return ec2.New(sess)
}
