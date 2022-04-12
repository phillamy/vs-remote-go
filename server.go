package main

import "base/router"

func main() {
	router.SetupRouter().Run(":8080")
}
