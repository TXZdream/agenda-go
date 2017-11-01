# Agenda —— Go version
![build status](https://travis-ci.org/freakkid/agenda-go.svg?branch=master)

## Team member
- 唐玄昭 xuanzhaotang@gmail.com
- 黄楠绚 threequarters@qq.com
- 夏显茁 xiaxzh2015@163.com

## Usage

* get help

    If you want to know more about parameters of the command, you can input
   
   > $ agenda-go help

    or

   > $ agenda-go help

    or

   > $ agenda-go help

   for details:

    ```
    Usage:
    agenda [command]

    Available Commands:
    help        Help about any command
    meeting     Manage meetings
    user        Manage user account

    Flags:
    -h, --help   help for agenda

    Use "agenda [command] --help" for more information about a command.
    ```

    If you want to know more details about command of user or meeting, you can input:

    > $ agenda-go user

    ```
    Usage:
    agenda user [flags]
    agenda user [command]

    Available Commands:
    create      create user account
    delete      Delete user account
    login       user login
    logout      Sign out
    show        Show user account

    Flags:
    -h, --help   help for user

    Use "agenda user [command] --help" for more information about a command.
    ```

    > $ agenda-go meeting

    ```
    Usage:
    agenda meeting [flags]
    agenda meeting [command]

    Available Commands:
    clear       Clear all meetings
    create      create meeting
    delete      Delete meeting
    leave       Leave meeting
    manage      Manage meeting
    show        Show meeting information

    Flags:
    -h, --help   help for meeting

    Use "agenda meeting [command] --help" for more information about a command.
    ```

    You can know about the details of every command by inputting

    > $ agenda meeting [command] -h

    or 

    > $ agenda user [command] -h

* register

    > $ agenda-go user create -uusername -e904811062@qq.com -p10086

    And input password correctly.

* log in
    > $ agenda-go user login -uusername

    And input your password.

* log out
    > $ agenda-go user logout

* show information of all users
    > $  agenda-go user show
    
    ```
    @username:
    Username       E-mail                   phone number
    llguser        fgh                      dfghjkl
    gdfghjkuser    fgh                      dfghjkl
    guser          fgh                      dfghjkl
    UUU            904811062@qq.com         123
    a              904811062@qq.com         13719342025
    ss             saaaaaaaa                sssssss
    aaa            dghjkl                   dfghjk
    username       904811062@qq.com         10086

    Total number is 8
    ```

* delete your account

    > $ agenda-go user delete


* create meeting
    > $ agenda-go meeting create --name meetingname
    ```
        @username:
        1. llguser
        2. gdfghjkuser
        3. guser
        4. UUU
        5. a
        6. ss
        7. aaa
        8. username
        Please choose the number of them to join your meeting(seprate with space): 1 2 3
        Please input start time(format: YYYY-MM-DD/HH:MM): 2017-11-11/11:22
        Please input end time(format: YYYY-MM-DD/HH:MM): 2017-11-11/12:22
        Create meeting meetingname finished.
    ```

* leave meeting you participate in
    > $ agenda-go meeting leave --name meetingname
    ```
    @user:
    Finish leaving meeting meetingname.
    ```

* manage meeting created by you
    + delete one of participators in your meeting
    > $ agenda-go meeting manage -d --name meetingname
    ```
    @username:
    Participators:
    1. llguser
    2. gdfghjkuser
    3. guser
    4. UUU
    5. a
    6. ss
    Please input the number you want to remove: 1
    llguser was removed.
    ```

    + add someone as a partipator in your meeting

    > $ agenda-go meeting manage --name meetingname
    ```
    @username:
    You can choose some of them to add to your meeting:
    1. llguser
    2. gdfghjkuser
    3. guser
    4. UUU
    5. a
    6. ss
    7. aaa
    8. username
    Please input the number of users you want to add(separate with blank): 1
    llguser was added.
    ```

* delete meeting created by you
    > $ agenda-go meeting delete --name meetingname
    ```
    @username:
    Delete meetingname finished.
    ```

* clear all meetings created by you
    > $ agenda-go meeting clear
    ```
    @username:
    Are you sure you want to clear all of your meetings? (y/n) y
    All of the meeting have been deleted.
    ```

* show information of all meetings you sponsored or participate in
    > $ agenda-go meeting show -s2015-11-11/11:00 -e2018-11-11/11:00
    ```
    @username:
    --·--·--·--·--·--·--·--·--·--·--
    Theme: meetingname
    Sponsor: username
    Start time: 2017-11-11/11:22
    End time: 2017-11-11/12:22
    Participator: llguser, gdfghjkuser, guser
    --·--·--·--·--·--·--·--·--·--·--
    ```

## data
All user data of our program include _user.json_, _meeting.json_, _curUser.txt_, _agenda.log_ is put in _HOME/.agenda_.


## code

Our code consists by _cmd_ and _entity_. _entity_ is responsible for low-level storage and logical processes._cmd_ is used for processing user input.


## At last, thanks for reading!