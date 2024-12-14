package repository

import (
	"context"
	"database/sql"

	"github.com/kadzany/closure-table-go/model/domain"
)

type NodeRepository interface {
	Create(ctx context.Context, tx *sql.Tx, node domain.Node) domain.Node
	Update(ctx context.Context, tx *sql.Tx, id string, node domain.Node) domain.Node
	DeleteByDescendantIds(ctx context.Context, tx *sql.Tx, descendantIds []string) error
	GetRootList(ctx context.Context, db *sql.DB) []domain.Node
	CheckByID(ctx context.Context, db *sql.DB, id string) bool
	DetailByID(ctx context.Context, db *sql.DB, id string) domain.Node
	GetDescendantList(ctx context.Context, db *sql.DB, nodeId string) []domain.Node
}
