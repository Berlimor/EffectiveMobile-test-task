package main

import (
	"EffectiveMobile/m/routers"
)

func main() {
	r := routers.InitRouter()

	r.GET("/music/all", routers.GetMusicList)
	r.GET("/music/filter", routers.GetMusicFiltered)
	r.POST("/music", routers.PostMusic)
	r.DELETE("/music", routers.DeleteMusic)

	r.GET("/music/text", routers.GetMusicText)
	r.GET("/music/details", routers.GetMusicDetails)
	r.POST("/music/details", routers.UpdateSongDetails)

	r.Run(":8080")
}