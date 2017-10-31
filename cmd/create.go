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
	"strconv"
	"strings"
	"os"
	"fmt"
	"bufio"
	"github.com/spf13/cobra"
	service "github.com/txzdream/agenda-go/entity/service"
	log "github.com/txzdream/agenda-go/entity/tools"
)

// createCmd represents the create command
var ucreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create user account",
	Long: `Use this command to create a new user account.`,
	Run: func(cmd *cobra.Command, args []string) {
		// get service
		var Service service.Service
		service.StartAgenda(&Service)
		// get createUser information
		createUsername, _ := cmd.Flags().GetString("username")
		createEmail, _ := cmd.Flags().GetString("email")
		createPhone, _ := cmd.Flags().GetString("phone")
		// check whether username, password, email or phone empty
		if createUsername == "" || createEmail == "" ||
		   createPhone == "" {
			   fmt.Fprintln(os.Stderr, "error : Username, Email or Phone is(are) empty")
				os.Exit(1)
			}
		// validator ? not necessary
		// get the password
		var createPassword string
		var prePassword string
		times := 1
		for {
			if times == 1 {
				fmt.Print("Please enter the password you want to create: ")
				fmt.Scanf("%s", &createPassword)
			} else {
				fmt.Print("Please enter the password again: ")
				fmt.Scanf("%s", &createPassword)
				if createPassword == prePassword {
					break
				} else {
					fmt.Println("The two passwords entered are not consistent. \nPlease restart setting password.")
				}
			}
			times *= -1
			prePassword = createPassword			
		}
		// check whether User is registed		
		ok := Service.UserRegister(createUsername, createPassword, createEmail, createPhone)
		if ok == false {
			fmt.Println(createUsername," has been registered")
			os.Exit(1)
		}
		fmt.Println("Sucess : Register ", createUsername)
		fmt.Println("You have logged in automatically.")
		Service.QuitAgenda(createUsername)
		os.Exit(0)
	},
}

var mcreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create meeting",
	Long: `Use this command to create a new meeting.`,
	Run: func(cmd *cobra.Command, args []string) {
		var Service service.Service
		service.StartAgenda(&Service)

		ok, name := Service.AutoUserLogin()
		if !ok {
			fmt.Fprintln(os.Stderr, "error: No current logged user.")
			log.LogInfoOrErrorIntoFile("any", false, fmt.Sprintf("Create meeting with no user login."))
			os.Exit(0)
		}

		if meetingName == "" {
			fmt.Fprintln(os.Stderr, "error: Meeting name is required.")
			log.LogInfoOrErrorIntoFile(name, false, fmt.Sprintf("create meeting %s.", meetingName))
			os.Exit(0)
		}
		var begin, end string
		var participator []string

		// show all users
		userList := Service.ListAllUsers()
		for i, v := range userList {
			fmt.Printf("%d. %s\n", i + 1, v.GetUserName())
		}

		fmt.Printf("Please choose the number of them to join your meeting(seprate with space): ")
		var chosenUsers string
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		chosenUsers = string(data)
		chosenList := strings.Split(chosenUsers, " ")
		
		for _, v := range chosenList {
			i, err := strconv.Atoi(v)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error: Invalid input.")
				log.LogInfoOrErrorIntoFile(name, false, fmt.Sprintf("Can not recognize %s when creating meeting.", v))
				os.Exit(0)
			}
			participator = append(participator, userList[i - 1].GetUserName())
		}

		// scan start time and end time
		var tmp int
		fmt.Printf("Please input start time(format: YYYY-MM-DD/HH:MM): ")
		fmt.Scanf("%s", &begin)
		fmt.Scanf("%d", &tmp)
		fmt.Printf("Please input end time(format: YYYY-MM-DD/HH:MM): ")
		fmt.Scanf("%s", &end)
		fmt.Scanf("%d", &tmp)
		
		ok = Service.CreateMeeting(name, meetingName, begin, end, participator)
		if ok {
			fmt.Printf("Create meeting %s finished.", meetingName)
			log.LogInfoOrErrorIntoFile(name, true, fmt.Sprintf("Finish creating meeting %s.", meetingName))
		} else {
			fmt.Printf("Can not create meeting %s.\n", meetingName)
			log.LogInfoOrErrorIntoFile(name, false, fmt.Sprintf("Fail to create meeting %s.", meetingName))
		}
	},
}

func init() {
	userCmd.AddCommand(ucreateCmd)
	meetingCmd.AddCommand(mcreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	mcreateCmd.Flags().StringVarP(&meetingName, "name", "", "", "Name for meeting you want to create.")

	// xiaxzh's part:
	ucreateCmd.Flags().StringP("username", "u", "", "Create Username")
	ucreateCmd.Flags().StringP("email", "e", "", "Create Email")
	ucreateCmd.Flags().StringP("phone", "p", "", "Create Phone")
}
