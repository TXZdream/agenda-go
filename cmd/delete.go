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
	"strings"
	"bufio"
	service "github.com/txzdream/agenda-go/entity/service"
	"github.com/spf13/cobra"
	log "github.com/txzdream/agenda-go/entity/tools"
)

// deleteCmd represents the delete command
var udeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete user account",
	Long: `Use this command to delete your account, meetings included.`,
	Run: func(cmd *cobra.Command, args []string) {
		var Service service.Service
		service.StartAgenda(&Service)
		// check whether other user logged in
		ok, curUsername := Service.AutoUserLogin()
		if ok == true {
			fmt.Println(strings.Join([]string{curUsername, "@:"}, ""))
		} else {
			fmt.Fprintln(os.Stderr, "Please Login in First")
			os.Exit(1)
		}

		// hints to ensure and enter password to delete User
		var password string
		fmt.Print("Plase enter password: ")
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		password = string(data)
		// delete user and meetings it participate
		if Service.DeleteUser(curUsername, password) == false {
			fmt.Fprintln(os.Stderr, "Some mistakes happend in DeleteUser.")
			os.Exit(1)
		}
		fmt.Println("Success : delete ", curUsername)
		ok = Service.UserLogout()
		if ok == false {
			fmt.Fprintln(os.Stderr, "some mistake happend in UserLogout")
			os.Exit(1)
		}
		Service.QuitAgenda()
		os.Exit(0)
	},
}

var mdeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete meeting",
	Long: `Use this command to delete specific meeting.`,
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
			log.LogInfoOrErrorIntoFile(name, true, fmt.Sprintf("Delete meeting with no user login."))
			os.Exit(0)
		}

		if meetingName == "" {
			fmt.Fprintln(os.Stderr, "error: Meeting name is required.")
			log.LogInfoOrErrorIntoFile(name, false, fmt.Sprintf("Delete meeting with no title login."))
			os.Exit(0)
		}
		ok = Service.DeleteMeeting(name, meetingName)
		if ok {
			fmt.Printf("Delete %s finished.\n", meetingName)
			log.LogInfoOrErrorIntoFile(name, true, fmt.Sprintf("Delete %s finished.", meetingName))
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
