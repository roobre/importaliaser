package importaliaser

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const HTML_TEMPLATE = `<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<meta name="go-import" content="%s %s %s">
%s
</head>
<body>
%s
</body>
`
const META_TEMPLATE = `<meta http-equiv="refresh" content="0; url=%s">`
const BODY_TEMPLATE = `See <a href="%s">this</a> for details`

type Aliaser struct {
	storer Store
}

func NewAliaser(s Store) *Aliaser {
	return &Aliaser{storer: s}

}

func (a *Aliaser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.Trim(r.RequestURI, "/") == "" {
		a.RootPath(w, r)
		return
	}

	a.Path(w, r)
}

func (a *Aliaser) RootPath(w http.ResponseWriter, r *http.Request) {
	conf := a.storer.Config()
	if conf.RootURL != "" {
		w.Header().Add("Location", conf.RootURL)
		w.WriteHeader(http.StatusFound)
	} else {
		http.NotFound(w, r)
		return
	}
}

func (a *Aliaser) Path(w http.ResponseWriter, r *http.Request) {
	repo := strings.Split(strings.Trim(r.RequestURI, "/"), "/")[0]
	name := r.Host + "/" + repo
	conf := a.storer.Config()

	alias, found := a.storer.Alias(name)
	if !found && conf.Speculative {
		found = true
		alias = Alias{
			Protocol: conf.SpeculativeProtocol,
			URI:      fmt.Sprintf(conf.SpeculativeFormat, repo),
		}
	}

	if found {
		var metaRedir, body string
		if strings.HasPrefix(alias.URI, "https://") {
			metaRedir = fmt.Sprintf(META_TEMPLATE, alias.URI)
			body = fmt.Sprintf(BODY_TEMPLATE, alias.URI)
		}

		_, _ = fmt.Fprintf(w, HTML_TEMPLATE, name, alias.Protocol, alias.URI, metaRedir, body)
	} else {
		log.Printf("Unknown alias %s", name)
		http.NotFound(w, r)
		return
	}
}
