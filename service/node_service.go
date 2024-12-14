package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kadzany/closure-table-go/model/dto"
)

type NodeService interface {
	Create(ctx *fiber.Ctx, request dto.NodeCreateRequest) (dto.NodeCreatedResponse, error)
	RootList(ctx *fiber.Ctx) ([]dto.NodeResponse, error)
	DetailNode(ctx *fiber.Ctx, nodeId string) (dto.NodeResponse, error)
	UpdateNode(ctx *fiber.Ctx, nodeId string, request dto.NodeUpdateRequest) (dto.NodeResponse, error)
	DeleteNode(ctx *fiber.Ctx, nodeId string) error
	DescendantList(ctx *fiber.Ctx, nodeId string) ([]dto.NodeResponse, error)
	MoveNode(ctx *fiber.Ctx, nodeId string, request dto.NodeMoveRequest) error
}
