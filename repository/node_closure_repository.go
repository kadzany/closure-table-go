package repository

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/kadzany/closure-table-go/model/domain"
)

type NodeClosureRepository interface {
	Save(ctx *fiber.Ctx, tx *sql.Tx, nodeClosures domain.NodeClosure) domain.NodeClosure
	DeleteByDescendantIds(ctx *fiber.Ctx, tx *sql.Tx, descendantIds []string) error
	FindDescendantIdsByAncestor(ctx *fiber.Ctx, tx *sql.Tx, ancestorId string) []string
	FindByDescendant(ctx *fiber.Ctx, db *sql.DB, nodeID string) []domain.NodeClosure
	GetNewClosures(ctx *fiber.Ctx, tx *sql.Tx, nodeId string, newAncestorId string) []domain.NodeClosure
}
