// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"os"
	"fmt"
	service "github.com/txzdream/agenda-go/entity/service"
	"github.com/spf13/cobra"
)

// manageCmd represents the manage command
var manageCmd = &cobra.Command{
	Use:   "manage",
	Short: "Manage meeting",
	Long: `Create meeting.`,
	Run: func(cmd *cobra.Command, args []string) {
		var Service service.Service
		service.StartAgenda(&Service)

		ok, name := Service.AutoUserLogin()
		if !ok {
			fmt.Fprintln(os.Stderr, "error: No current logged user.")
			os.Exit(0)
		}
		
		if meetingName == "" {
			fmt.Fprintln(os.Stderr, "error: Meeting theme is required.")
			os.Exit(0)
		}
		//TODO: manage meeting
		fmt.Println(name)
	},
}

func init() {
	meetingCmd.AddCommand(manageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// manageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// manageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	manageCmd.Flags().StringVarP(&meetingName, "name", "", "", "the name of meeting to be managed")
}
