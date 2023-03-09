package handlers

import (
	"fmt"
	"net/http"
	"platform/pkg/config"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func Newhandlers(r *Repository) {
	Repo = r
}

//var wg sync.WaitGroup

// wg.Add(1)
func (m *Repository) Get(w http.ResponseWriter, r *http.Request) {
	v_id := r.URL.Query().Get(":id")
	result := m.App.GetViews(v_id)
	fmt.Fprintf(w, fmt.Sprintf("%s", result))
}

//wg.Wait()

func (m *Repository) Set(w http.ResponseWriter, r *http.Request) {
	v_id := r.URL.Query().Get(":id")
	result, views := m.App.CreateAndResetViews(v_id)
	fmt.Fprintf(w, fmt.Sprintf("%s for given id and final views are : %v", result, views))

}
