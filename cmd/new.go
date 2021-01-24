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
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	componentName string
	baseDir       string
)

// Comp is a component
type Comp struct {
	Component string
	Workspace string
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "new",
	Short: "Generates new from template",
	Long: `Generation tool for MVC.
Default settings takes current working directory as basedir for determining where the templates are.
If customizing project workspace to make use of this without specifying template dir:

<dir>
|-> templates
	|-> controller.templ
	|-> repo.templ
	|-> service.templ

Format must be as above.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		CheckFlag(cmd, componentName)
		Info("Generate component %s", componentName)
		wd, err := os.Getwd()
		if err != nil {
			Error("%s", err)
		}
		if baseDir != "" {
			wd = baseDir
		}
		generateComponent(componentName, wd)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&componentName, "component", "c", "", "Generates MVC component with Controller-Service-Repository")
	generateCmd.Flags().StringVarP(&baseDir, "baseDir", "b", "", "Base dir acting as parent folder for template dir")
}

func generateComponent(name, wd string) {

	templateFiles := map[string]string{
		"repository.templ": "repositories",
		"service.templ":    "services",
		"controller.templ": "controllers",
	}

	newComp := Comp{
		Component: name,
		Workspace: "_", // todo
	}

	for templateName, folderName := range templateFiles {
		path := filepath.Join(wd, "templates", templateName)
		templ := template.Must(template.ParseFiles(path))

		var filename string
		filename = strings.Title(name) + strings.Title(strings.Split(templateName, ".")[0]) + ".go"

		_ = os.Mkdir(filepath.Join(wd, folderName), 0700)
		f, err := os.Create(filepath.Join(wd, folderName, filename))
		if err != nil {
			Error("%s", err)
		}

		CheckIfError(templ.Execute(f, newComp))
		f.Close()
	}

}
