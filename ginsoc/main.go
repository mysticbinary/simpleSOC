package main

import (
	"ginsoc/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// init DB

	// init gin
	r := gin.Default()
	r = router.CollectRoute(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
