package main

import (
	"todo_list/cache"
	"todo_list/conf"
	"todo_list/routes"
)

func main(){
	conf.Init()
	cache.ConnRedis()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}