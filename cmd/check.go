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
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/eareese/todo/canvas"
	"github.com/spf13/cobra"
)

type assignment struct {
	Name   string
	Course string
}

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check against Canvas for new tasks.",
	Long: `The check command looks for ungraded assignments to convert
into task items.

Requires a valid access token for the Canvas instance to be
configured in the TODO_TOKEN environment variable.`,
	Run: func(cmd *cobra.Command, args []string) {
		// collect ungraded assignments:
		var ungraded []assignment

		// first query up all available courses
		courses, err := canvas.QueryCourses()
		if err != nil {
			panic(err)
		}

		// look up the names of ungraded assignments for each course found
		for _, course := range courses {
			assignments, err := canvas.QueryAssignmentsUngraded(course.ID)
			if err != nil {
				fmt.Printf("failed to query assignments for course with ID %d:%v\n", course.ID, err)
				continue
			}
			for _, assn := range assignments {
				ungraded = append(ungraded, assignment{Name: assn.Name, Course: course.Name})
			}
		}

		// prompt to add the ungraded as tasks
		fmt.Printf("Found %d ungraded assignments in %d courses.\n", len(ungraded), len(courses))
		fmt.Println("Add them all as Grading tasks? [Y/n]")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "", "y", "Y", "yes", "Yes":
			for _, u := range ungraded {
				task := fmt.Sprintf("Grade %s", u.Name)
				Add(task)
			}
		default:
			fmt.Println("No grading tasks were added.")
		}
	},
}

func init() {
	RootCmd.AddCommand(checkCmd)
}
