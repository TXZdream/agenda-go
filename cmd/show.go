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

// showCmd represents the show command
var ushowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show user account",
	Long: `Use this command to show every user's information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("show called")
	},
}

var mshowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show meeting information",
	Long: `Use this command to show meeting information.`,
	Run: func(cmd *cobra.Command, args []string) {
		var Service service.Service
		service.StartAgenda(&Service)

		ok, name := Service.AutoUserLogin()
		if !ok {
			fmt.Fprintln(os.Stderr, "error: No current logged user.")
			os.Exit(0)
		}
		
		if startTime == "" || endTime == "" {
			fmt.Fprintln(os.Stderr, "error: Start time and end time is required.")
			os.Exit(0)
		}
		meetingList := Service.MeetingQueryByUserAndTime(name, startTime, endTime)

		// print all meetings with the given name
		if len(meetingList) == 0 {
			fmt.Println("--·--·--·--·--·--·--·--·--·--·--")
			fmt.Println("No matching meeting")
			fmt.Println("--·--·--·--·--·--·--·--·--·--·--")
		} else {
			fmt.Println("--·--·--·--·--·--·--·--·--·--·--")
			for _, v := range meetingList {
				fmt.Printf("Theme: %s\n", v.GetTitle())
				fmt.Printf("Sponsor: %s\n", v.GetSponsor())
				fmt.Printf("Start time: %s\n", v.GetStartDate())
				fmt.Printf("End time: %s\n", v.GetEndDate())
				fmt.Printf("Participator: %s\n", v.GetParticipators())
				fmt.Println("--·--·--·--·--·--·--·--·--·--·--")
			}
		}
	},
}

func init() {
	userCmd.AddCommand(ushowCmd)
	meetingCmd.AddCommand(mshowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	mshowCmd.Flags().StringVarP(&startTime, "start", "s", "", "the start time of a meeting")
	mshowCmd.Flags().StringVarP(&endTime, "end", "e", "", "the end time of a meeting")
}
