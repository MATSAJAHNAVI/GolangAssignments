package config

import (
	"strconv"
	"sync"
)

// var app *config.AppConfig

// func NewVideos(a *config.AppConfig) {
// 	app = a
// }

var mutex sync.Mutex

// func (app *AppConfig) SeeViews(id string, m *sync.Mutex) int {
// 	//defer wg.Done()
// 	m.Lock()
// 	v := app.Viewsmap[id].Views
// 	m.Unlock()
// 	return v
// }

func (app *AppConfig) GetViews(v_id string) string {

	_, present := app.Viewsmap[v_id]
	mutex.Lock()
	defer mutex.Unlock()
	if !present {
		res := "video not found"
		return res
	} else {
		//views := app.SeeViews(v_id, &mutex)
		views := app.Viewsmap[v_id].Views
		res := strconv.Itoa(views)
		res1 := "Total views for this video :" + res
		return res1
	}

}

func (app *AppConfig) CreateAndResetViews(v_id string) (string, int) {

	_, present := app.Viewsmap[v_id]

	mutex.Lock()
	defer mutex.Unlock()

	if !present {
		temp_v := Video{
			Id:    v_id,
			Views: 1,
		}
		app.Viewsmap[v_id] = &temp_v
		app.Videos = append(app.Videos, temp_v)
		res1 := "video created successfully"
		res2 := temp_v.Views
		return res1, res2
	} else {
		//views := app.SeeViews(v_id, &mutex)
		//views++
		//app.Viewsmap[v_id].Views = views
		views := app.Viewsmap[v_id].Views
		views++
		for i, j := range app.Videos {
			if j.Id == v_id {
				app.Videos[i].Views++
				break
			}
		}
		res1 := "views incremented"
		res2 := views
		return res1, res2
	}
}

//wg.Wait()
