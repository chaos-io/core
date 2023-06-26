package main

import (
	"net/http"

	"github.com/chaos-io/core/vanity/internal/repos"
	template2 "github.com/chaos-io/core/vanity/web/template"
)

func handleIndex(w http.ResponseWriter, _ *http.Request) {
	err := template2.Index.Execute(w, repos.Repos)
	if err != nil {
		logs.Error("error while rendering template", "error", err)
	}
}

func handleGoGet(w http.ResponseWriter, r *http.Request) {
	relpath := chi.URLParam(r, "*")

	repo, ok := repos.Repos[relpath]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data := struct {
		Relpath string
		Repo    *repos.Repo
	}{
		Relpath: relpath,
		Repo:    repo,
	}

	err := template2.GoGet.Execute(w, data)
	if err != nil {
		logs.Error("error while rendering template", "error", err)
	}
}

func handlePing(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("pong"))
}
