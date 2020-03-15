// Package web is used to create and set up the routes used to host the server.
//
// Single Project Host - func (*ServerSettings) SingleHostProjectSetup
//
package web

import (
	"os"

	"go.lafronz.com/vanago/tools/logger"
)

//ServerSettings a struck that holds variables used to create and direct the server
type ServerSettings struct {
	// Port is used to hold the port value that the server will run on
	// Loaded from Environment variable PORT
	Port string
	// RedirectSettings uses the interface RedirectSettings.
	// It holds the settings needed for the service to redirect the go get to the correct location.
	RedirectSettings RedirectSettings
}

// S exposes the Server Env Variables to the main function
var S ServerSettings

// RedirectSettings stores the data used to redirect the requester to the proper go source host
type RedirectSettings struct {
	// IndexRedirect is the host you would like the index of this service to be redirect to.
	// Loaded from Environment variable indexRedirect
	IndexRedirect string
	// VCSBase is the host name of the Version Control System - ie 'https://github.com/'
	// Loaded from Environment variable vcsBase
	VCSBase string

	// VCSTeamName is the team or user name at the VCS - ie 'lafronzt'
	// Loaded from Environment variable vcsTeamName
	VCSTeamName string

	// VCSLinkMiddle is the part of the url that is in between VCSTeamName and the repo name - ie '/'
	// Loaded from Environment variable vcsLinkMiddle
	VCSLinkMiddle string

	// ProjectName is used in Logging only
	// Loaded from Environment variable projectName
	ProjectName string
}

// init loads initial data from the system environment variables
func init() {

	// use PORT environment variable, or default to 8080
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	// vcsBase environment variable, or default to "https://github.com"
	vcsBase := "https://github.com/"
	if fromEnv := os.Getenv("vcsBase"); fromEnv != "" {
		vcsBase = fromEnv
	}

	// vcsTeamName environment variable, or default to "lafronzt"
	vcsTeamName := "lafronzt"
	if fromEnv := os.Getenv("vcsTeamName"); fromEnv != "" {
		vcsTeamName = fromEnv
	}

	// vcsLinkMiddle environment variable, or default to "/"
	// used to separate the team name from the repo name
	vcsLinkMiddle := "/"
	if fromEnv := os.Getenv("vcsLinkMiddle"); fromEnv != "" {
		vcsLinkMiddle = fromEnv
	}

	// projectName environment variable, or default to "Personal"
	projectName := "Personal"
	if fromEnv := os.Getenv("projectName"); fromEnv != "" {
		projectName = fromEnv
	}

	// indexRedirect environment variable, or default to "http://TylerLaFronz.com"
	indexRedirect := "http://TylerLaFronz.com"
	if fromEnv := os.Getenv("indexRedirect"); fromEnv != "" {
		projectName = fromEnv
	}

	S = ServerSettings{
		Port: port,
		RedirectSettings: RedirectSettings{
			IndexRedirect: indexRedirect,
			VCSBase:       vcsBase,
			VCSTeamName:   vcsTeamName,
			VCSLinkMiddle: vcsLinkMiddle,
			ProjectName:   projectName,
		},
	}

	logger.Info("Loaded Data: %+v", S)

}
