# gooki

Sync users from google groups to the nuki user database.

## Usage
[embedmd]:# (tmp/help.txt)
```txt
A command line tool to enable you to synchronise your Google
Apps (Google Workspace) users to nuki users

Usage:
  gooki [flags]

Flags:
  -d, --debug                       enable verbose / debug logging
  -u, --google-admin string         Google Workspace admin user email
  -c, --google-credentials string   path to Google Workspace credentials file (default "credentials.json")
  -h, --help                        help for gooki
      --ignore-users strings        ignores these Google Workspace users
      --listen string               Address to listen on for internal server (default ":8080")
      --log-format string           log format (default "text")
      --log-level string            log level (default "info")
  -t, --nuki-token string           Nuki Bearer Token for authentication with the nuki API.
      --smartlock-id int            Nuki Smartlock ID; if 0 no smartlock auths will be created.
  -p, --sync-period duration        Time between two syncronizations, will only run once if set to 0
  -m, --user-match string           Google Workspace Users filter query parameter, example: 'name:John* email:admin*', see: https://developers.google.com/admin-sdk/directory/v1/guides/search-users
  -v, --version                     version for gooki
```
