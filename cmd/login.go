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
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	service "github.com/txzdream/agenda-go/entity/service"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "User login",
	Long:  `Use this command to sign in to the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		// get Service
		var Service service.Service
		service.StartAgenda(&Service)
		// check whether other user logged in
		ok, name := Service.AutoUserLogin()
		if ok == true {
			fmt.Println(strings.Join([]string{name, "@:"}, ""))
		}
		// get username
		username, _ := cmd.Flags().GetString("username")
		// check whether username or password empty
		if username == "" {
			fmt.Fprintln(os.Stderr, "error : Username is empty")
			os.Exit(1)
		}
		// wait for password
		var password string
		fmt.Printf("Please enter the password: ")
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		password = string(data)
		// check whether user is registed
		ok = Service.IsRegisteredUser(username)
		if ok == false {
			fmt.Fprintln(os.Stderr, "error : This user not exists")
			os.Exit(1)
		}
		// check whether has Login in
		ok, CurUserName := Service.AutoUserLogin()
		if CurUserName == username {
			fmt.Fprintln(os.Stderr, "error : You have Logined in as ", CurUserName)
			os.Exit(1)
		}
		// check the password
		var times int
		for {
			ok = Service.UserLogin(username, password)
			if ok == false {
				if times < 2 {
					times++
					fmt.Print("Wrong password, Please try again: ")
					fmt.Scanf("%s", &password)
				} else {
					fmt.Fprintln(os.Stderr, "error : Wrong password")
					os.Exit(1)
				}
			} else {
				break
			}
		}
		// Succeed in Login as {username}
		fmt.Println("success : You have Logined in as ", username)
		fmt.Println("Welcome to use Agenda!")
		Service.QuitAgenda()
		os.Exit(0)
	},
}

func init() {
	userCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	loginCmd.Flags().StringP("username", "u", "", "Login username")
}
