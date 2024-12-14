package repository

import (
	"context"
	"database/sql"

	"github.com/kadzany/closure-table-go/model/domain"
	"github.com/kadzany/closure-table-go/pkg"
	"github.com/lib/pq"
)

type NodeRepositoryImpl struct {
}

func NewNodeRepository() NodeRepository {
	return &NodeRepositoryImpl{}
}

func (repository *NodeRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, node domain.Node) domain.Node {
	// Save Root Node
	SQL := `INSERT INTO nodes (id, title, type, description, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := tx.QueryRowContext(ctx, SQL,
		node.ID,
		node.Title,
		node.Type,
		node.Description,
		node.CreatedAt,
	).Scan(&node.ID)

	// Panic if error
	pkg.PanicIfError(err)

	// return root node
	return node
}

func (repository *NodeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, id string, node domain.Node) domain.Node {
	// Update Node
	SQL := `UPDATE nodes SET title = $1, type = $2, description = $3, updated_at = $4 WHERE id = $5`
	_, err := tx.ExecContext(ctx, SQL,
		node.Title,
		node.Type,
		node.Description,
		node.UpdatedAt,
		id,
	)

	// Panic if error
	pkg.PanicIfError(err)

	// return updated node
	return node
}

func (repository *NodeRepositoryImpl) DeleteByDescendantIds(ctx context.Context, tx *sql.Tx, descendantIds []string) error {
	// Delete Node By Descendant Ids
	SQL := `DELETE FROM nodes WHERE id = ANY($1)`
	_, err := tx.ExecContext(ctx, SQL, pq.Array(descendantIds))

	// Panic if error
	pkg.PanicIfError(err)

	// return nil
	return nil
}

func (repository *NodeRepositoryImpl) GetRootList(ctx context.Context, db *sql.DB) []domain.Node {
	// Get Root List
	SQL := `SELECT n.id, n.title, n.type, n.description, n.created_at, n.updated_at
			FROM nodes n
			    JOIN node_closure nc ON n.id = nc.descendant
			WHERE nc.ancestor = nc.descendant
			  AND nc.depth = 0
			  AND NOT EXISTS (SELECT 1
			                  FROM node_closure nc2
			                  WHERE nc2.descendant = nc.descendant
			                    AND nc2.ancestor != nc.descendant)
			ORDER BY n.created_at DESC`
	rows, err := db.QueryContext(ctx, SQL)

	// Panic if error
	pkg.PanicIfError(err)

	// Close rows
	defer pkg.CloseRows(rows)

	// Create nodes slice
	var nodes []domain.Node

	// Loop through rows
	for rows.Next() {
		// Create node
		node := domain.Node{}

		// Scan rows to node
		err := rows.Scan(
			&node.ID,
			&node.Title,
			&node.Type,
			&node.Description,
			&node.CreatedAt,
			&node.UpdatedAt,
		)

		// Panic if error
		pkg.PanicIfError(err)

		// Append node to nodes
		nodes = append(nodes, node)
	}

	// Return nodes
	return nodes
}

func (repository *NodeRepositoryImpl) CheckByID(ctx context.Context, db *sql.DB, id string) bool {
	// Check Node By ID
	SQL := `SELECT id FROM nodes WHERE id = $1`
	rows, err := db.QueryContext(ctx, SQL, id)

	// Panic if error
	pkg.PanicIfError(err)

	// Close rows
	defer pkg.CloseRows(rows)

	// Check if rows next exist
	return rows.Next()
}

func (repository *NodeRepositoryImpl) GetNodeByID(ctx context.Context, db *sql.DB, id string) (domain.Node, error) {
	// Get Node By ID
	SQL := `SELECT id, title, type, description, created_at, updated_at FROM nodes WHERE id = $1`
	row := db.QueryRowContext(ctx, SQL, id)

	// Create node
	node := domain.Node{}

	// Scan row to node
	err := row.Scan(
		&node.ID,
		&node.Title,
		&node.Type,
		&node.Description,
		&node.CreatedAt,
		&node.UpdatedAt,
	)

	// Return empty node if error
	if err != nil {
		return domain.Node{}, err
	}

	// Return node
	return node, nil
}

func (repository *NodeRepositoryImpl) GetDescendantList(ctx context.Context, db *sql.DB, nodeId string) ([]domain.Node, error) {
	// Get Descendant List
	SQL := `SELECT n.id, n.title, n.type, n.description, n.created_at, n.updated_at
			FROM nodes n
			    JOIN node_closure nc ON n.id = nc.descendant
			WHERE nc.ancestor = $1
			  AND nc.depth > 0
			ORDER BY n.created_at DESC`
	rows, err := db.QueryContext(ctx, SQL, nodeId)

	// Panic if error
	if err != nil {
		return nil, err
	}

	// Close rows
	defer pkg.CloseRows(rows)

	// Create nodes slice
	var nodes []domain.Node

	// Loop through rows
	for rows.Next() {
		// Create node
		node := domain.Node{}

		// Scan rows to node
		err := rows.Scan(
			&node.ID,
			&node.Title,
			&node.Type,
			&node.Description,
			&node.CreatedAt,
			&node.UpdatedAt,
		)

		// Panic if error
		if err != nil {
			return nil, err
		}

		// Append node to nodes
		nodes = append(nodes, node)
	}

	// Return nodes
	return nodes, nil
}
