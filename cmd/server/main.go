package main

import (
	"github.com/longln/go-ecommerce-backend/internal/routers"

)

func main() {
	r := routers.NewRouter()
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run("0.0.0.0:8080")
}