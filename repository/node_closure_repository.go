package repository

import (
	"context"
	"database/sql"

	"github.com/kadzany/closure-table-go/model/domain"
)

type NodeClosureRepository interface {
	Save(ctx context.Context, tx *sql.Tx, nodeClosures domain.NodeClosure) domain.NodeClosure
	DeleteByDescendantIds(ctx context.Context, tx *sql.Tx, descendantIds []string) error
	FindDescendantIdsByAncestor(ctx context.Context, tx *sql.Tx, ancestorId string) []string
	FindByDescendant(ctx context.Context, db *sql.DB, nodeID string) []domain.NodeClosure
	GetNewClosures(ctx context.Context, tx *sql.Tx, nodeId string, newAncestorId string) []domain.NodeClosure
}
