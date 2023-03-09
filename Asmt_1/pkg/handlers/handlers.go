package handlers

import (
	"fmt"
	"net/http"
	"platform/pkg/config"
	"sync"
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
	var result string
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		result = m.App.GetViews(v_id)
	}()
	wg.Wait()
	//fmt.Fprintf(w, fmt.Sprintf("%s", result))
	fmt.Fprintf(w, result)
}

//wg.Wait()

func (m *Repository) Set(w http.ResponseWriter, r *http.Request) {
	v_id := r.URL.Query().Get(":id")
	var result string
	var views int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		result, views = m.App.CreateAndResetViews(v_id)
	}()
	wg.Wait()
	fmt.Fprintf(w, fmt.Sprintf("%s for given id and final views are : %v", result, views))

}
