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
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	snippetFilename string
)

// snippetGenCmd represents the snippetGen command
var snippetGenCmd = &cobra.Command{
	Use:   "snippetGen",
	Short: "generates vs code snippet body from a file",
	Long: `since body code requires the snippets to be in an array of strings,
	this code converts:
	Hello world
		This should be indented
			This should be indented twice

	into:
	[
		"Hello world",
		"\nThis should be indented",
		"\n\nThis should be indented twice",
	]

	and saves it into a <filename>.snippet.json file
	`,
	Run: func(cmd *cobra.Command, args []string) {
		CheckFlag(cmd, snippetFilename)
		Info("Generating from file %s", snippetFilename)
		wd, err := os.Getwd()
		if err != nil {
			Error("%s", err)
		}

		path := filepath.Join(wd, snippetFilename)
		file, err := os.Open(path)
		CheckIfError(err)
		defer file.Close()
		// Read file line by line
		scanner := bufio.NewScanner(file)

		// Create file
		created := path + ".snippet.json"
		f, _ := os.Create(created)
		defer f.Close()

		w := bufio.NewWriter(f)
		fmt.Fprint(w, "[\n")
		for scanner.Scan() {
			line := "\t\"" + scanner.Text() + "\",\n"
			strings.ReplaceAll(line, "\t", "\\t")
			fmt.Fprint(w, line)
		}
		fmt.Fprint(w, "]")
		w.Flush()

		Info("generated in %s", created)
	},
}

func init() {
	rootCmd.AddCommand(snippetGenCmd)
	snippetGenCmd.Flags().StringVarP(
		&snippetFilename,
		"filename",
		"f",
		"",
		"Snippet filename with the code to be used for generation",
	)
}
