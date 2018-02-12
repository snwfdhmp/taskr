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

	"github.com/snwfdhmp/taskr/pkg"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Install the taskr components in your current repository",
	Long: `Install the taskr components in your current repository:
.taskr/
  history.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := taskr.MkInit(); err != nil {
			fmt.Println("fatal mk:", err)
			return
		}

		if err := taskr.SampleTask("taskr"); err != nil {
			fmt.Println("fatal:", err)
			return
		}

		if err := taskr.AddPreCommitHook(); err != nil {
			fmt.Println("fatal:", err)
			return
		}

		if err := taskr.AddPostCommitHook(); err != nil {
			fmt.Println("fatal:", err)
			return
		}

		fmt.Println("taskr inited successfully.")
	},
}

func init() {
	RootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
