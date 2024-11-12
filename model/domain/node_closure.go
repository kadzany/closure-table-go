package domain

import "github.com/google/uuid"

type NodeClosure struct {
	Ancestor   uuid.UUID `db:"ancestor" json:"ancestor"`
	Descendant uuid.UUID `db:"descendant" json:"descendant"`
	Depth      int       `db:"depth" json:"depth"`
}
