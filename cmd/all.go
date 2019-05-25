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

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "All delete",
	Long:  `All delete`,
	Args: func(cmd *cobra.Command, args []string) error {
		// 引数になにもない場合
		if len(args) < 1 {
			fmt.Print("All delete?[y/n]: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			//　[y/n]
			if scanner.Text() == "y" || scanner.Text() == "Y" {
				// docker container 全削除
				cmdstr := "docker ps -aq | xargs docker rm -f"
				outCon, errCon := exec.Command("sh", "-c", cmdstr).Output()
				if outCon != nil {
					fmt.Printf("%s", outCon)
				} else {
					fmt.Printf("%s", errCon)
				}
				// docker image 全削除
				cmdstrImage := "docker images -aq | xargs docker rmi -f"
				outImg, errImg := exec.Command("sh", "-c", cmdstrImage).Output()
				if outImg != nil {
					fmt.Printf("%s", outImg)
				} else {
					fmt.Printf("%s", errImg)
				}
				// docker volume 全削除
				cmdstrVolume := "docker volume ls -q | xargs docker volume rm -f"
				outVol, errVol := exec.Command("sh", "-c", cmdstrVolume).Output()
				if outVol != nil {
					fmt.Printf("%s", outVol)
				} else {
					fmt.Printf("%s", errVol)
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
	rootCmd.AddCommand(allCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
