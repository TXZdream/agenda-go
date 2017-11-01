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
	"bufio"
	service "github.com/txzdream/agenda-go/entity/service"
	"github.com/spf13/cobra"
	"strconv"
	log "github.com/txzdream/agenda-go/entity/tools"
)

// manageCmd represents the manage command
var manageCmd = &cobra.Command{
	Use:   "manage",
	Short: "Manage meeting",
	Long: `Create meeting.`,
	Run: func(cmd *cobra.Command, args []string) {
		var Service service.Service
		service.StartAgenda(&Service)
		// check whether other user logged in
		ok, name := Service.AutoUserLogin()
		if ok == true {
			fmt.Println(strings.Join([]string{name,"@:"}, ""))
		}
		if !ok {
			fmt.Fprintln(os.Stderr, "error: No current logged user.")
			log.LogInfoOrErrorIntoFile(name, false, fmt.Sprintf("Manage meeting with no user login."))
			os.Exit(0)
		}
		
		if meetingName == "" {
			fmt.Fprintln(os.Stderr, "error: Meeting theme is required.")
			log.LogInfoOrErrorIntoFile(name, false, fmt.Sprintf("Manage  meeting %s with no title.", meetingName))
			os.Exit(0)
		}
		
		meetingList := Service.MeetingQueryByTitle(name, meetingName)
		if len(meetingList) == 0 {
			fmt.Println("No matching meeting with the given theme.")
			os.Exit(1)
		}

		// delete users
		if isDelete {
			var participator []string
			fmt.Println("Participators:")
			for i, v := range meetingList[0].GetParticipators() {
				participator = append(participator, v)
				fmt.Printf("%d. %s\n", i + 1, v)
			}
			fmt.Print("Please input the number you want to remove: ")
			var inputNums string
			reader := bufio.NewReader(os.Stdin)
			data, _, _ := reader.ReadLine()
			inputNums = string(data)
			chosenList := strings.Split(inputNums, " ")
			var toBeRemovedParticipators []string
			for _, v := range chosenList {
				num, err := strconv.Atoi(v)
				if err != nil || num <= 0 || num > len(participator) {
					fmt.Fprintln(os.Stderr, "error: Invalid input")
					os.Exit(1)
				}
				toBeRemovedParticipators = append(toBeRemovedParticipators, participator[num - 1])
			}
			for _, v := range toBeRemovedParticipators {
				ok := Service.DeleteParticipatorByTitle(name, meetingName, v)
				if ok {
					fmt.Printf("%s was removed.\n", v)
					log.LogInfoOrErrorIntoFile(name, true, fmt.Sprintf("Remove %s from meeting %s.", v, meetingName))
				} else {
					fmt.Printf("%s can not be removed.\n", v)
					log.LogInfoOrErrorIntoFile(name, false, fmt.Sprintf("Can not remove %s from meeting %s.", v, meetingName))
				}
			}
		} else {
			// add users
			fmt.Println("You can choose some of them to add to your meeting:")
			userList := Service.ListAllUsers()
			for i, v := range userList {
				fmt.Printf("%d. %s\n", i + 1, v.GetUserName())
			}
			fmt.Print("Please input the number of users you want to add(separate with blank): ")
			var userNums string
			reader := bufio.NewReader(os.Stdin)
			data, _, _ := reader.ReadLine()
			userNums = string(data)
			userNumList := strings.Split(userNums, " ")
			for _, v := range userNumList {
				if len(v) == 0 {
					continue
				}
				i, ok := strconv.Atoi(v)
				if ok != nil || i > len(userList) || i <= 0 {
					fmt.Fprintln(os.Stderr, "error: Invalid input.")
					os.Exit(0)
				}
				if Service.AddParticipatorByTitle(name, meetingName, userList[i - 1].GetUserName()) {
					fmt.Printf("%s was added.\n", userList[i - 1].GetUserName())
					log.LogInfoOrErrorIntoFile(name, true, fmt.Sprintf("Add %s to meeting %s.", userList[i - 1].GetUserName(), meetingName))
				} else {
					fmt.Printf("%s can not be added.\n", userList[i - 1].GetUserName())
					log.LogInfoOrErrorIntoFile(name, false, fmt.Sprintf("Can not add %s to meeting %s.", userList[i - 1].GetUserName(), meetingName))
				}
			}
		}
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
	manageCmd.Flags().BoolVarP(&isDelete, "", "d", false, "Delete user(s) from a meeting")
}
