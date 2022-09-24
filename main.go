package main

import (
	"github.com/suryaadi44/go-training-restful/config"
	"github.com/suryaadi44/go-training-restful/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
