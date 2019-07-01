// Copyright © 2019 Red Hat
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"path"

	"github.com/fusor/cpma/pkg/env"
	"github.com/fusor/cpma/pkg/transform"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()

	// Get config file from CLI argument an save it to viper config
	rootCmd.PersistentFlags().StringVar(&env.ConfigFile, "config", "", "config file (Default searches ./cpma.yaml, $HOME/cpma.yml)")

	// Set log level from CLI argument
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "show debug ouput")
	env.Config().BindPFlag("Debug", rootCmd.PersistentFlags().Lookup("debug"))

	// Output logs to console if true
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	env.Config().BindPFlag("Verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	// Allow insecure host key if true
	rootCmd.PersistentFlags().BoolP("allow-insecure-host", "i", false, "allow insecure ssh host key ")
	env.Config().BindPFlag("InsecureHostKey", rootCmd.PersistentFlags().Lookup("allow-insecure-host"))

	// Get OCP3 source cluster and save it to viper config
	rootCmd.PersistentFlags().StringP("hostname", "n", "", "OCP3 cluster hostname")
	env.Config().BindPFlag("Hostname", rootCmd.PersistentFlags().Lookup("hostname"))

	// Get OCP3 source cluster name that is used in kubeconfig and save it to viper config
	rootCmd.PersistentFlags().StringP("cluster-name", "c", "", "OCP3 cluster kubeconfig name")
	env.Config().BindPFlag("ClusterName", rootCmd.PersistentFlags().Lookup("cluster-name"))

	// Get ssh config values from CLI argument
	rootCmd.PersistentFlags().StringVarP(&env.Login, "ssh-login", "l", "", "OCP3 ssh login")
	rootCmd.PersistentFlags().StringVarP(&env.PrivateKey, "ssh-keyfile", "k", "", "OCP3 ssh keyfile path")
	rootCmd.PersistentFlags().StringVarP(&env.Port, "ssh-port", "p", "", "OCP3 ssh port")

	// Get config file from CLI argument an save to viper config
	rootCmd.PersistentFlags().StringP("output-dir", "o", path.Dir("."), "set the directory to store extracted configuration.")
	env.Config().BindPFlag("OutputDir", rootCmd.PersistentFlags().Lookup("output-dir"))
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cpma",
	Short: "Helps migration cluster configuration of a OCP 3.x cluster to OCP 4.x",
	Long:  `Helps migration cluster configuration of a OCP 3.x cluster to OCP 4.x`,
	Run: func(cmd *cobra.Command, args []string) {
		env.InitLogger()

		if err := env.InitConfig(); err != nil {
			logrus.Fatal(err)
		}

		transform.Start()
	},
	Args: cobra.MaximumNArgs(0),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// It only needs to happen once.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
	}
}
