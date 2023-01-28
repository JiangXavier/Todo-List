package main

import (
	"todo_list/conf"
	"todo_list/routes"
)

func main(){
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}