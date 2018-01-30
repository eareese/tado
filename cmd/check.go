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
		fmt.Println("check called")
	},
}

type courseData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type assignmentData struct {
	ID                int    `json:"id"`
	Description       string `json:"description"`
	NeedsGradingCount int    `json:"needs_grading_count"`
}

func init() {
	RootCmd.AddCommand(checkCmd)
}
