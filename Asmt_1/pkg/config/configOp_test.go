package config

import "testing"

func TestAppConfig_GetViews(t *testing.T) {
	app := &AppConfig{
		Videos:   []Video{},
		Viewsmap: make(map[string]*Video),
	}

	res1 := "video not found"
	res2 := "Total views for this video :"

	type arg struct {
		v_id string
	}
	tests := []struct {
		name     string
		args     arg
		wantdata string
	}{
		{name: "test1", args: arg{v_id: "first"}, wantdata: res1},
		{name: "test2", args: arg{v_id: "first"}, wantdata: res2 + "1"},
	}
	for _, tt := range tests {
		got := app.GetViews(tt.args.v_id)
		app.CreateAndResetViews(tt.args.v_id)
		if got != tt.wantdata {
			t.Errorf("GetViews() = %v, want %v", got, tt.wantdata)
		}

	}
}

func TestAppConfig_CreateAndResetViews(t *testing.T) {
	app := &AppConfig{
		Videos:   []Video{},
		Viewsmap: make(map[string]*Video),
	}

	res1 := "video created successfully"
	res2 := "views incremented"
	type arg struct {
		v_id string
	}
	tests := []struct {
		name      string
		args      arg
		wantdata1 string
		wantdata2 int
	}{
		{name: "test1", args: arg{v_id: "first"}, wantdata1: res1, wantdata2: 1},
		{name: "test2", args: arg{v_id: "first"}, wantdata1: res2, wantdata2: 1},
	}
	for _, tt := range tests {
		got, value := app.CreateAndResetViews(tt.args.v_id)
		if got != tt.wantdata1 {
			t.Errorf("wrong message printed want %s but got %s", tt.wantdata1, got)
		}
		if value != tt.wantdata2 {
			t.Errorf("false view count wanted %v but got %v", tt.wantdata2, value)
		}
	}
}
