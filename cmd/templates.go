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
	"os"

	"github.com/gophersumit/goscaffold/pkg/templates"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// templatesCmd represents the templates command
var templatesCmd = &cobra.Command{
	Use:   "templates",
	Short: "list available templates",
	Long:  `templates command lists all the available templates.`,
	Run: func(cmd *cobra.Command, args []string) {
		templateFile, err := templates.GetGroups()
		if err != nil {
			fmt.Println(err)
			return
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Group Name", "Group Description", "Template Name", "Template Description"})

		for _, group := range templateFile.Groups {
			for _, template := range group.Templates {
				row := []string{group.Name, group.Description, template.Name, template.Description}
				table.Append(row)
			}
		}

		table.Render()

	},
}

func init() {
	rootCmd.AddCommand(templatesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// templatesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// templatesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
