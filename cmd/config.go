/*
Copyright © 2024 crazyoptimist <hey@crazyoptimist.net>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"ec2control/pkg/utils"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the application",
	Long: `Provide necessary config values, then the application will
generate a configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		config := readinConfig()

		configDir, err := utils.GetConfigDir()
		if err != nil {
			log.Fatalln(err)
		}

		err = generateConfigFile(*config, filepath.Join(configDir, DEFAULT_CONFIG_FILE_NAME)+".toml")
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

type AWSProfile struct {
	Region             string
	AWSAccessKeyId     string
	AWSSecretAccessKey string
}

type appConfig struct {
	AWSProfile    AWSProfile
	EC2InstanceID string
}

func readinConfig() *appConfig {
	stdinReader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your AWS region name: ")
	Region := utils.GetUserInput(stdinReader)
	fmt.Print("Enter your AWS Access Key ID: ")
	AWSAccessKeyId := utils.GetUserInput(stdinReader)
	fmt.Print("Enter your AWS Secret Access Key: ")
	AWSSecretAccessKey := utils.GetUserInput(stdinReader)

	fmt.Print("Enter your EC2 Instance ID: ")
	EC2InstanceID := utils.GetUserInput(stdinReader)

	return &appConfig{
		AWSProfile: AWSProfile{
			Region,
			AWSAccessKeyId,
			AWSSecretAccessKey,
		},
		EC2InstanceID: EC2InstanceID,
	}
}

func generateConfigFile(config appConfig, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = toml.NewEncoder(file).Encode(config)
	if err != nil {
		return err
	}

	fmt.Printf("Config file created successfully: %s\n", path)
	return nil
}
