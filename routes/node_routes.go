package routes

import (
	"closure-table-go/controller"
	"closure-table-go/repository"
	"closure-table-go/service"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func InitNodeRoutes(server *fiber.App, db *sql.DB, validate *validator.Validate) {
	// Setup Node API
	nodeRepository := repository.NewNodeRepository()
	nodeClosureRepository := repository.NewNodeClosureRepository()
	nodeService := service.NewNodeService(nodeRepository, nodeClosureRepository, db, validate)
	nodeController := controller.NewNodeController(nodeService)

	// Set Routes
	v1NodesAPI := server.Group("/v1/nodes")
	v1NodesAPI.Post("/", nodeController.Create)
	v1NodesAPI.Get("/", nodeController.RootList)
	v1NodesAPI.Get("/:nodeId", nodeController.DetailNode)
	v1NodesAPI.Put("/:nodeId", nodeController.UpdateNode)
	v1NodesAPI.Delete("/:nodeId", nodeController.DeleteNode)
	v1NodesAPI.Get("/:nodeId/descendants", nodeController.DescendantList)
	v1NodesAPI.Put("/:nodeId/move", nodeController.MoveNode)
}
