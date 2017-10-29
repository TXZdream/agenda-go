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
	"strings"
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
		var Service service.Service
		service.StartAgenda(&Service)
		// check whether user login
		ok, CurUsername := Service.AutoUserLogin()
		if ok == false {
			fmt.Fprintln(os.Stderr, "error : No User has Logined in")
			os.Exit(1)
		}
		// get email and phone by username
		ok, email, phone := Service.ListUserInformation(CurUsername)
		if ok == false {
			fmt.Fprintln(os.Stderr, "Some mistakes happend in ListUserInformation")
			os.Exit(1)	
		}
		fmt.Println("Username : ", CurUsername)
		fmt.Println("Email : ", email)
		fmt.Println("Phone : ", phone)
		// get meetings by username
		meetings := Service.ListAllMeetings(CurUsername)
		for _, item := range meetings {
			fmt.Println(strings.Join([]string{item.GetTitle(), item.GetStartDate(), item.GetEndDate()}, " "))
		}
		os.Exit(0)
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
		
		if meetingName == "" {
			fmt.Fprintln(os.Stderr, "error: Meeting theme is required.")
			os.Exit(0)
		}
		meetingList := Service.MeetingQueryByTitle(name, meetingName)

		// print all meetings with the given name
		if len(meetingList) == 0 {
			fmt.Println("--·--·--·--·--·--·--·--·--·--·--")
			fmt.Println("No matching meeting")
			fmt.Println("--·--·--·--·--·--·--·--·--·--·--")
		} else {
			fmt.Println("--·--·--·--·--·--·--·--·--·--·--")
			fmt.Printf("Theme: %s\n", meetingList[1].GetTitle())
			fmt.Printf("Sponsor: %s\n", meetingList[0].GetSponsor())
			fmt.Printf("Start time: %s\n", meetingList[0].GetStartDate())
			fmt.Printf("End time: %s\n", meetingList[0].GetEndDate())
			fmt.Printf("Participator: %s\n", meetingList[0].GetParticipators())
			fmt.Println("--·--·--·--·--·--·--·--·--·--·--")
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
	mshowCmd.Flags().StringVarP(&meetingName, "name", "", "", "the name of meeting to be managed")

	// xiaxzh's part:
}
