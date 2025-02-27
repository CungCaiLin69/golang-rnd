package main

import (
	"golang-rnd/initializers"
	"golang-rnd/routes"
)

func init(){
	initializers.LoadEnv()
	initializers.ConnectDb()
	initializers.SyncDb()
}

func main() {
	routes.Router()
}