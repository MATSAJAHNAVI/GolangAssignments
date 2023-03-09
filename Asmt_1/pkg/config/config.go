package config

type Video struct {
	Id    string
	Views int
}

type AppConfig struct {
	Videos   []Video
	Viewsmap map[string]*Video
}
