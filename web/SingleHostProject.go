package web

import (
	"net/http"
	"strings"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"go.lafronz.com/vanago/tools/logger"
)

// SingleHostProjectSetup is the Setup script to build all the server routes to handle http requests
func (s *ServerSettings) SingleHostProjectSetup() *http.Server {
	// register hello function to handle all requests
	server := mux.NewRouter()
	server.HandleFunc("/", s.indexRedirectHandler)
	server.HandleFunc("/{package}", s.singleHostPackageHandler)
	server.HandleFunc(`/{package}/{sub:[a-zA-Z0-9=\-\/]+}`, s.singleHostPackageHandler)
	server.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "OK", http.StatusOK)
		// log.Println("Health Check Hit")
		return
	})

	// Server Settings
	svr := &http.Server{
		Handler:      server,
		Addr:         ":" + s.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return svr

}

func (s *ServerSettings) indexRedirectHandler(w http.ResponseWriter, r *http.Request) {

	if err := indexTmpl.Execute(w, struct {
		Redirect string
	}{
		Redirect: s.RedirectSettings.IndexRedirect,
	}); err != nil {
		http.Error(w, "cannot render the page", http.StatusInternalServerError)
		logger.Error("%s", err)
	}
	return
}

func (s *ServerSettings) singleHostPackageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if len(vars["package"]) <= 0 {
		s.indexRedirectHandler(w, r)
		return
	}

	if strings.EqualFold(vars["package"], "favicon.ico") {
		w.WriteHeader(404)
		return
	}

	logger.Info("Path Vars - Project: %s - Package: %s", s.RedirectSettings.ProjectName, vars["package"])

	w.Header().Set("Cache-Control", "86400")
	if err := vanityTmpl.Execute(w, struct {
		Import  string
		Subpath string
		Repo    string
		Display string
		VCS     string
	}{
		Import:  r.Host + "/" + vars["package"],
		Subpath: vars["sub"],
		Repo:    s.RedirectSettings.VCSBase + s.RedirectSettings.VCSTeamName + s.RedirectSettings.VCSLinkMiddle + vars["package"],
		VCS:     "git",
	}); err != nil {
		http.Error(w, "cannot render the page", http.StatusInternalServerError)
		logger.Error("%s", err)
	}
	return
}

var indexTmpl = template.Must(template.New("index").Parse(`<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta http-equiv="refresh" content="0; url={{.Redirect}}">
</head>
<body>
Nothing to see here; <a href="{{.Redirect}}">Visit this page for more</a>.
</body>
</html>`))

var vanityTmpl = template.Must(template.New("vanity").Parse(`<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="{{.Import}} {{.VCS}} {{.Repo}}">
<meta name="go-source" content="{{.Import}} {{.Repo}} {{.Repo}}/tree/master{/dir} {{.Repo}}/blob/master{/dir}/{file}#L{line}">
<meta http-equiv="refresh" content="0; url=https://pkg.go.dev/{{.Import}}/{{.Subpath}}">
</head>
<body>
Nothing to see here; <a href="https://pkg.go.dev/{{.Import}}/{{.Subpath}}">see the package on pkg.go.dev</a>.
</body>
</html>`))
