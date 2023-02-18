package http

import (
	"core/http/controllers"
	"core/infra/conn"
	repo "core/repository"
	"core/svc"
	"github.com/gin-gonic/gin"
)

func Init(g interface{}) {
	const VersionPrefix = "/v1"
	grp := g.(*gin.RouterGroup)
	grp = grp.Group(VersionPrefix)
	db := conn.Db()
	//conn.Migrate()

	// repo layer
	sysRepo := repo.NewSystemRepository(db)
	//prodRepo := repo.NewSystemRepository(db)

	// service layer
	sysSvc := svc.NewSystemService(sysRepo)
	//prodSvc := svc.NewSystemService(prodRepo)

	// SYSTEM routes
	sys := controllers.NewSystemController(sysSvc)
	grp.GET("/root", sys.Root)
	grp.GET("/h34l7h", sys.Health)

	// Product routes
	//sys := controllers.NewSystemController(prodSvc)
	//grp.GET("/products", sys.Health)
}
