/*
Copyright © 2023 Sumit Agrawal <gophersumit@gmail.com>

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
	"fmt"

	"github.com/gophersumit/goscaffold/pkg/gonew"
	"github.com/gophersumit/goscaffold/pkg/templates"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"golang.org/x/mod/module"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		allgroups, err := templates.GetGroups()
		if err != nil {
			fmt.Printf("Error getting project template groups: %v\n", err)
			return
		}

		type TemplateDetails struct {
			Name string
			URL  string
		}

		var list []TemplateDetails
		for _, group := range allgroups.Groups {
			for _, template := range group.Templates {
				list = append(list, TemplateDetails{Name: template.Name, URL: template.URL})
			}
		}

		var names []string
		for _, template := range list {
			names = append(names, template.Name)
		}
		selectTemplate := promptui.Select{
			Label: "Select Project Template",
			Items: names,
		}

		_, selected, err := selectTemplate.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You choose template%q\n", selected)

		selectName := promptui.Prompt{
			Label: "What will be the new project name?",
			Validate: func(input string) error {
				if len(input) < 3 {
					return fmt.Errorf("name must be at least 3 characters")
				}
				if err := module.CheckPath(input); err != nil {
					return fmt.Errorf("invalid destination module name: %v", err)
				}
				return nil
			},
		}

		givenName, err := selectName.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		fmt.Printf("You choose name%q\n", givenName)

		templateUrl := ""
		for _, template := range list {
			if template.Name == selected {
				templateUrl = template.URL
			}
		}

		err = gonew.Create(templateUrl, givenName, "")
		if err != nil {
			fmt.Printf("Error creating new project: %v\n", err)
			return
		}
		fmt.Printf("Project created successfully\n")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
