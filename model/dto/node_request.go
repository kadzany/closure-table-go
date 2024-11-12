package dto

type NodeCreateRequest struct {
	Title       string  `json:"title" form:"title" validate:"required"`
	Type        string  `json:"type" form:"type" validate:"required,oneof=note task reminder"`
	Description *string `json:"description,omitempty" form:"description,omitempty"`
	AncestorID  *string `json:"ancestor_id,omitempty" form:"ancestor_id,omitempty"`
}

type NodeUpdateRequest struct {
	Title       string  `json:"title" form:"title" validate:"required"`
	Type        string  `json:"type" form:"type" validate:"required,oneof=note task reminder"`
	Description *string `json:"description,omitempty" form:"description,omitempty"`
}

type NodeMoveRequest struct {
	ToAncestorID string `json:"to_ancestor_id" form:"to_ancestor_id" validate:"required"`
}
