package server

import "github.com/fitrah-firdaus/notes/config"

func Init() {
	conf := config.GetConfig()
	r := NewRouter()
	r.Run(conf.GetString("server.port"))
}
