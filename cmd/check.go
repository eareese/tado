// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"fmt"

	"github.com/eareese/todo/canvas"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check against Canvas for new tasks.",
	Long: `The check command looks for ungraded assignments to convert
into task items.

Requires a valid access token for the Canvas instance to be
configured in the TODO_TOKEN environment variable.`,
	Run: func(cmd *cobra.Command, args []string) {
		got, err := canvas.QueryCourses()
		if err != nil {
			panic(err)
		}
		fmt.Println("Got this many courses from query: ", len(got))
	},
}

// func queryAssignments(courseId int) ([]assignmentData, error) {
// 	resp, err := http.Get("")
// }

func init() {
	RootCmd.AddCommand(checkCmd)
}
