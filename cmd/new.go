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
	workspace     string
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

Specify template dir if possible.
By default, assumes current working dir has a structure of:

<dir>
|-> templates
	|-> controller.templ
	|-> repo.templ
	|-> service.templ
	|-> model.templ

Generated files will always be in current working dir.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// Pre-parses arguments before generating
		CheckFlag(cmd, componentName)
		Info("Generate component %s", componentName)
		wd, err := os.Getwd()
		if err != nil {
			Error("%s", err)
		}
		if baseDir != "" {
			wd = baseDir
		} else {
			wd = filepath.Join(wd, "templates")
		}
		if workspace == "" {
			splitted := strings.Split(wd, "/")
			workspace = splitted[len(splitted)-1]
		}
		generateComponent(componentName, wd, workspace)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&componentName, "component", "c", "", "Generates MVC component with Controller-Service-Repository")
	generateCmd.Flags().StringVarP(&baseDir, "templates", "t", "", "Template dir holding all templates")
	generateCmd.Flags().StringVarP(&workspace, "workspace", "w", "", "Workspace for import template purposes, defaults to same as baseDir")
}

func generateComponent(name, templateDir, workspace string) {
	templateFiles := map[string]string{
		"repository.templ": "repositories",
		"service.templ":    "services",
		"controller.templ": "controllers",
		"model.templ":      "models",
	}

	newComp := Comp{
		Component: name,
		Workspace: workspace,
	}

	for templateName, folderName := range templateFiles {
		path := filepath.Join(templateDir, templateName)
		templ := template.Must(template.ParseFiles(path))

		filename := strings.Title(name) + strings.Title(strings.Split(templateName, ".")[0]) + ".go"
		cwd, _ := os.Getwd()

		_ = os.Mkdir(filepath.Join(cwd, folderName), 0700)
		f, err := os.Create(filepath.Join(cwd, folderName, filename))
		if err != nil {
			Error("%s", err)
		}

		CheckIfError(templ.Execute(f, newComp))
		f.Close()
	}

}
