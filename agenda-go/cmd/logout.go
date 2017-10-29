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

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Sign out",
	Long: `Use this command to sign out`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logout called")
		// get service
		var Service service.Service
		service.StartAgenda(&Service)
		// check Whether CurUser exits
		ok, CurUsername := Service.AutoUserLogin()
		if ok == false {
			fmt.Fprintln(os.Stderr, "error : Current User not exits")
			os.Exit(1)
		}
		fmt.Println("Success : ", CurUsername, " Logout")
		ok = Service.QuitAgenda("")
		if ok == false {
			fmt.Fprintln(os.Stderr, "error : Some mistakes happend in QuitAgenda")
		}
		os.Exit(0)
	},
}

func init() {
	userCmd.AddCommand(logoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
