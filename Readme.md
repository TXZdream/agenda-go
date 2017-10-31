# Agenda —— Go version
![build status](https://travis-ci.org/freakkid/agenda-go.svg?branch=master)

## Team member
- 唐玄昭 xuanzhaotang@gmail.com
- 黄楠绚 threequarters@qq.com
- 夏显茁 xiaxzh2015@163.com

## Usage
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