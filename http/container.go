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

	// service layer
	sysSvc := svc.NewSystemService(sysRepo)

	// SYSTEM routes
	sys := controllers.NewSystemController(sysSvc)
	grp.GET("/root", sys.Root)
	grp.GET("/h34l7h", sys.Health)
}
