package routers

import (
	"github.com/longln/go-ecommerce-backend-api/internal/routers/manager"
	"github.com/longln/go-ecommerce-backend-api/internal/routers/user"
)


type RouterGroup struct {
	User user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}


var RouterGroupApp = new(RouterGroup)