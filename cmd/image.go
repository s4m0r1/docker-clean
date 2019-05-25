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
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "delete docker image",
	Long:  `delete docker image`,
	Args: func(cmd *cobra.Command, args []string) error {
		// 引数になにもない場合
		if len(args) < 1 {
			fmt.Print("Delete Docker image?[y/n]: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			//　[y/n]
			if scanner.Text() == "y" || scanner.Text() == "Y" {
				// docker image 全削除
				cmdstr := "docker images -aq | xargs docker rmi -f"
				out, err := exec.Command("sh", "-c", cmdstr).Output()
				if out != nil {
					fmt.Printf("%s", out)
				} else {
					fmt.Printf("%s", err)
				}
			} else if scanner.Text() == "n" || scanner.Text() == "N" {
				fmt.Printf("Exit...\n")
			}
		// 	//引数がある場合
		// } else {
		// 	//引数のパース
		// 	flag.Parse()
		// 	args := flag.Args()
		// 	fmt.Println(args)
		// }
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// imageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// imageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
