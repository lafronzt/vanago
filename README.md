# Vanago

By Tyler La Fronz

`go get go.lafronz.com/vanago`

## Description

Vanago is a simple webserver that allows anyone to host custom vanity urls for their Go import paths.

## How to Use

1. Set the Environment variables:

    | Name | Default | Description |
    |------|---------|-------------|
    | PORT | 8080    | Port to Open Server On |
    | vcsBase | https://github.com | URL base that your Git Repos are hosted on |
    | vcsTeamName | lafronzt | Team or User Name |
    | vcsLinkMiddle | / | used to separate the team name from the repo name |
    | projectName | Personal | Used to Log the Project name in requests |
    | indexRedirect | http://TylerLaFronz.com | URL to redirect the naked index request to |

2. Run the server

```bash
    go install go.lafronz.com/vanago/cmd/vanago
```

## Why I made it

As part of my push to have a better understand of Go, I started reading about how the `go get` command works.
After learning how it functioned, I was searching around for any tool that would allow me to customize my go import paths.
I was unable to find anything that really fit my needs or was dynamic enough for me to use for both public and private packages.
So I created Vanago, which I now use to run my Go import path.

### Future

I am looking to expand this in the future to support multiple Teams/Users on the same URL.
