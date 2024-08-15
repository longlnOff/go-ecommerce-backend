package user

import "github.com/gin-gonic/gin"

type ProductRouter struct{

}

func (pr *ProductRouter) InitProductRouter(r *gin.RouterGroup) {
	// public router
	productRouterPublic := r.Group("/product")
	{
		productRouterPublic.GET("/search")
		productRouterPublic.GET("/detail/:id")
	}

	// private router
	
}