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
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"
)

var (
	filename    string
	word        string
	replacement string
)

// replaceCmd represents the replace command
var replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "Replaces word 'w' with word 'r'",
	Long:  `See flags for replacing words w with word r in file f`,
	Run: func(cmd *cobra.Command, args []string) {
		CheckFlag(cmd, filename)
		CheckFlag(cmd, word)
		CheckFlag(cmd, replacement)

		fi, err := ioutil.ReadFile(filename)
		CheckIfError(err)
		output := strings.Replace(string(fi), word, replacement, -1)

		err = ioutil.WriteFile(filename, []byte(output), 0)
		CheckIfError(err)
		Info("Replaced %s with %s in %s", word, replacement, filename)
	},
}

func init() {
	rootCmd.AddCommand(replaceCmd)
	replaceCmd.Flags().StringVarP(&filename, "filename", "f", "", "File to execute replacement")
	replaceCmd.Flags().StringVarP(&word, "word", "w", "", "Word to replace")
	replaceCmd.Flags().StringVarP(&replacement, "replacement", "r", "", "Replace with")
}
