/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	componentName string
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "new",
	Short: "Generates new from template",
	Long:  `Generation tool for MVC`,
	Run: func(cmd *cobra.Command, args []string) {
		CheckFlag(cmd, componentName)
		Info("Generate component %s", componentName)
		generateComponent(componentName)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&componentName, "component", "c", "", "Generates MVC component with Controller-Service-Repository")
}

func generateComponent(name string) {

}
