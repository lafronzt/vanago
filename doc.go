/*
1. Set the Environment variables:

	| Name          | Default                 | Description                                       |
	| PORT 	        | 8080                    | Port to Open Server On                            |
	| vcsBase       | https://github.com      | URL base that your Git Repos are hosted on        |
	| vcsTeamName   | lafronzt                | Team or User Name                                 |
	| vcsLinkMiddle | /                       | used to separate the team name from the repo name |
	| projectName   | Personal                | Used to Log the Project name in requests          |
	| indexRedirect | http://TylerLaFronz.com | URL to redirect the naked index request to        |

2. Run the server
    go install go.lafronz.com/cmd/vanago

*/

package cmd