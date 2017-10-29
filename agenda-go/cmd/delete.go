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
	"fmt"
	"os"
	service "github.com/txzdream/agenda-go/entity/service"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var udeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete user account",
	Long: `Use this command to delete your account, meetings included.`,
	Run: func(cmd *cobra.Command, args []string) {
		var Service service.Service
		service.StartAgenda(&Service)

		// check username empty
		username, _ := cmd.Flags().GetString("username")
		if username == "" {
			fmt.Fprintln(os.Stderr, "error : Username is empty")
			os.Exit(1)
		}
		// check whether User Login
		_, loginUsername := Service.AutoUserLogin()
		if username == loginUsername {
			// hints to ensure and enter password to delete User
			var password string
			fmt.Println("Ensure to delete User : ", username)
			fmt.Println("Plase enter password :")
			fmt.Scanf("%s", &password)
			// chech the password
			ok := Service.UserLogin(username, password)
			if ok == false {
				fmt.Fprintln(os.Stderr, "error : Wrong password")
				os.Exit(1)
			}
			// delete user and meetings it participate
			Service.DeleteUser(loginUsername)
			fmt.Println("Success : delete ", loginUsername)
			Service.QuitAgenda("")
			os.Exit(0)
		} else {
			fmt.Fprintln(os.Stderr, "Please Login in First")
			os.Exit(1)
		}
	},
}

var mdeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete meeting",
	Long: `Use this command to delete specific meeting.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
		var Service service.Service
		service.StartAgenda(&Service)

		ok, name := Service.AutoUserLogin()
		if !ok {
			fmt.Fprintln(os.Stderr, "error: No current logged user.")
			os.Exit(0)
		}

		if meetingName == "" {
			fmt.Fprintln(os.Stderr, "error: Meeting name is required.")
			os.Exit(0)
		}
		ok = Service.DeleteMeeting(name, meetingName)
		if ok {
			fmt.Printf("Delete %s finished.\n", meetingName)
		} else {
			fmt.Printf("Can not delete the meeting called %s.\n", meetingName)
		}
	},
}

func init() {
	userCmd.AddCommand(udeleteCmd)
	meetingCmd.AddCommand(mdeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	mdeleteCmd.Flags().StringVarP(&meetingName, "name", "", "", "meeting name to be deleted")
	udeleteCmd.Flags().StringP("username", "u", "", "Delete user")
}
