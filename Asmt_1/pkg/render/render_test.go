package render

import "testing"

// func TestGetViews(t *testing.T) {
// 	//var w http.ResponseWriter
// 	var app config.AppConfig
// 	app.Viewsmap = make(map[string]*config.Video)

// 	w := httptest.NewRecorder()
// 	testid := "first"
// 	GetViews(w, testid)
// 	res := w.Result()
// 	defer res.Body.Close()
// 	data, _ := ioutil.ReadAll(res.Body)
// 	s1 := "Video not found"
// 	if string(data) != s1 {
// 		t.Errorf("expected %s got %s", s1, string(data))
// 	}

//}

// func TestGetViews(t *testing.T) {
// 	videos := []config.Video{}
// 	viewsmap := make(map[string]*config.Video)
// 	res1 := "video not found"
// 	res2 := "Total views for this video :"
// 	type args struct {
// 		v_id string
// 		app  *config.AppConfig
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		{name: "test1",
// 			args: args{v_id: "first", app: &config.AppConfig{Videos: videos, Viewsmap: viewsmap}},
// 			want: res1,
// 		},

// 	{name: "test2",
// 		args: args{v_id: "first", app: &config.AppConfig{Videos: videos, Viewsmap: viewsmap}},
// 		want: res2 + "1",
// 	},
// 	// TODO: Add test cases.
// }
// for _, tt := range tests {
// 	t.Run(tt.name, func(t *testing.T) {
// 		if got := GetViews(tt.args.v_id, tt.args.app); got != tt.want {
// 			t.Errorf("GetViews() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestGetViews(t *testing.T) {
	res1 := "video not found"
	res2 := "Total views for this video :"
	type args struct {
		v_id string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{v_id: "first"}, want: res1},
		{name: "test2", args: args{v_id: "first"}, want: res2 + "1"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetViews(tt.args.v_id); got != tt.want {
				t.Errorf("GetViews() = %v, want %v", got, tt.want)
			}
		})
	}
}
