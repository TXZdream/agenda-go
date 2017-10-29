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
	// "strconv"
	// "strings"
	"os"
	"fmt"
	"github.com/spf13/cobra"
	// model "github.com/txzdream/agenda-go/entity/model"
	service "github.com/txzdream/agenda-go/entity/service"
)

// createCmd represents the create command
var ucreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create user account",
	Long: `Use this command to create a new user account.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
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
				fmt.Println("Please enter the password you want to create :")
				fmt.Scanf("%s", &createPassword)
			} else {
				fmt.Println("Please enter the password again :")
				fmt.Scanf("%s", &createPassword)
				if createPassword == prePassword {
					break
				} else {
					fmt.Println("\nThe two passwords entered are not consistent. \nAnd restart setting password.\n")
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
		Service.QuitAgenda(createUsername)
		os.Exit(0)
	},
}

var mcreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create meeting",
	Long: `Use this command to create a new meeting.`,
	Run: func(cmd *cobra.Command, args []string) {
		// var Service service.Service
		// service.StartAgenda(&Service)

		// ok, name := Service.AutoUserLogin()
		// if !ok {
		// 	fmt.Fprintln(os.Stderr, "error: No current logged user.")
		// 	os.Exit(0)
		// }

		// if meetingName == "" {
		// 	fmt.Fprintln(os.Stderr, "error: Meeting name is required.")
		// 	os.Exit(0)
		// }
		// var begin, end string
		// var participator []model.User

		// // show all users
		// userList := Service.ListAllUsers()
		// for i, v := range userList {
		// 	fmt.Printf("%d. %s\n", i, v.GetUserName())
		// }

		// fmt.Printf("Please choose some of them to join your meeting(seprate with space): ")
		// var chosenUsers, tmp string
		// fmt.Scanf("%s", &chosenUsers)
		// fmt.Scanf("%d", &tmp)
		// chosenList := strings.Split(" ", chosenUsers)
		
		// for _, v := range chosenList {
		// 	i, err := strconv.Atoi(v)
		// 	if err != nil {
		// 		fmt.Fprintln(os.Stderr, "error: Invalid input.")
		// 	}
		// 	participator = append(participator, userList[i])
		// }

		// // scan start time and end time
		// fmt.Printf("Please input start time(format: YYYY-MM-DD/HH:MM): ")
		// fmt.Scanf("%s", &begin)
		// fmt.Scanf("%d", &tmp)
		// fmt.Scanf("%s", &end)
		// fmt.Scanf("%d", &tmp)
		
		// //  small tips from xiaxzh : 
		// // 'participators' segement required by 
		// //  func (service *Service) CreateMeeting(sponsor string, title string,  
		// //                     startDateString string, endDateString string, participators []string) bool
		// //  should be []string but not [] model.User
	    // Service.CreateMeeting(name, meetingName, begin, end, participator)
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
