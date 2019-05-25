// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"os/exec"

	"github.com/spf13/cobra"
)

// containerCmd represents the container command
var containerCmd = &cobra.Command{
	Use:   "container",
	Short: "All delete docker container",
	Long:  `All delete docker container`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			fmt.Print("All Delete Docker container?[y/n]: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			//　[y/n]
			if scanner.Text() == "y" || scanner.Text() == "Y" {
				// docker container 全削除
				cmdstr := "docker ps -aq | xargs docker rm -f"
				out, err := exec.Command("sh", "-c", cmdstr).Output()
				if out != nil {
					fmt.Printf("%s", out)
				} else {
					fmt.Printf("%s", err)
				}
			} else if scanner.Text() == "n" || scanner.Text() == "N" {
				fmt.Printf("Exit...\n")
			}
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(containerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// containerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// containerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
