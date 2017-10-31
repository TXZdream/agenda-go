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
	service "github.com/txzdream/agenda-go/entity/service"
	"github.com/spf13/cobra"
	log "github.com/txzdream/agenda-go/entity/tools"
)

// leaveCmd represents the leave command
var leaveCmd = &cobra.Command{
	Use:   "leave",
	Short: "Leave meeting",
	Long: `Use this command to leave a meeting.`,
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
			log.LogInfoOrErrorIntoFile(name, false, fmt.Sprintf("Leave meeting with no user login."))
			os.Exit(0)
		}
		
		if meetingName == "" {
			fmt.Fprintln(os.Stderr, "error: Meeting theme is required.")
			log.LogInfoOrErrorIntoFile(name, false, fmt.Sprintf("Leave meeting with no title."))
			os.Exit(0)
		}
		
		ok = Service.QuitMeeting(name, meetingName)
		if ok {
			fmt.Printf("Finish leaving meeting %s.\n", meetingName)
			log.LogInfoOrErrorIntoFile(name, true, fmt.Sprintf("Leave meeting %s.", meetingName))
		} else {
			fmt.Printf("Can not leave this meeting %s.\n", meetingName)
			log.LogInfoOrErrorIntoFile(name, true, fmt.Sprintf("Can not leave meeting %s.", meetingName))
		}
	},
}

func init() {
	meetingCmd.AddCommand(leaveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// leaveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// leaveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	leaveCmd.Flags().StringVarP(&meetingName, "name", "", "", "the name of meeting to be managed")
}
