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
	"os"

	"github.com/snwfdhmp/taskr/pkg"
	"github.com/spf13/cobra"
)

// preCommitCmd represents the preCommit command
var preCommitCmd = &cobra.Command{
	Use:   "pre-commit",
	Short: "Run pre-commit flow",
	Long: `taskr will:
- parse 'taskr.yaml' for tasks and './taskr/history.yaml' as history
- run tasks and verify if no regression is introduced`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := taskr.ParseTasks("taskr.yaml")
		if err != nil {
			fmt.Println("fatal parse:", err)
			os.Exit(-1)
		}

		history, err := taskr.OpenHistory()
		if err != nil {
			fmt.Println("fatal read:", err)
			os.Exit(-1)
		}

		if _, err = history.Run(tasks...); err != nil {
			fmt.Println("fatal run:", err)
			os.Exit(1)
		}

		os.Exit(0)
	},
}

func init() {
	RootCmd.AddCommand(preCommitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// preCommitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// preCommitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
