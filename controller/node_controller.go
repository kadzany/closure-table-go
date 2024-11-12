package controller

import (
	"github.com/gofiber/fiber/v2"
)

type NodeController interface {
	Create(ctx *fiber.Ctx) error
	RootList(ctx *fiber.Ctx) error
	DetailNode(ctx *fiber.Ctx) error
	UpdateNode(ctx *fiber.Ctx) error
	DeleteNode(ctx *fiber.Ctx) error
	DescendantList(ctx *fiber.Ctx) error
	MoveNode(ctx *fiber.Ctx) error
}
