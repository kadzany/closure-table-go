package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/kadzany/closure-table-go/model/domain"
	"github.com/kadzany/closure-table-go/model/dto"
	"github.com/kadzany/closure-table-go/pkg"
	"github.com/kadzany/closure-table-go/repository"
)

type NodeServiceImpl struct {
	NodeRepository        repository.NodeRepository
	NodeClosureRepository repository.NodeClosureRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewNodeService(
	nodeRepository repository.NodeRepository,
	nodeClosureRepository repository.NodeClosureRepository,
	db *sql.DB,
	validate *validator.Validate,
) NodeService {
	return &NodeServiceImpl{
		NodeRepository:        nodeRepository,
		NodeClosureRepository: nodeClosureRepository,
		DB:                    db,
		Validate:              validate,
	}
}

func (service *NodeServiceImpl) Create(ctx context.Context, request dto.NodeCreateRequest) (dto.NodeCreatedResponse, error) {
	// Validate request
	err := service.Validate.Struct(request)
	if err != nil {
		return dto.NodeCreatedResponse{}, err
	}

	// Check Ancestor Node
	if request.AncestorID != nil {
		isAncestorNodeExist := service.NodeRepository.CheckByID(ctx, service.DB, *request.AncestorID)
		if !isAncestorNodeExist {
			return dto.NodeCreatedResponse{}, fmt.Errorf("ancestor node is not found")
		}
	}

	// Start transaction
	tx, err := service.DB.Begin()
	if err != nil {
		return dto.NodeCreatedResponse{}, err
	}

	// Defer commit or rollback
	defer pkg.CommitOrRollback(tx)

	// Save node
	description := sql.NullString{Valid: false}
	if request.Description != nil {
		description = sql.NullString{String: *request.Description, Valid: true}
	}
	node := domain.Node{
		ID:          uuid.New(),
		Title:       request.Title,
		Type:        request.Type,
		Description: description,
		CreatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
	}
	createdNode := service.NodeRepository.Create(ctx, tx, node)

	// Save NodeClosure : Self Reference
	closure := domain.NodeClosure{
		Ancestor:   createdNode.ID,
		Descendant: createdNode.ID,
		Depth:      0,
	}
	service.NodeClosureRepository.Save(ctx, tx, closure)

	// When Node Have Ancestor
	if request.AncestorID != nil {
		// Get Ancestor Closures
		ancestorClosures := service.NodeClosureRepository.FindByDescendant(ctx, service.DB, *request.AncestorID)

		// Save NodeClosure : Ancestor Reference
		depth := 1
		for _, ancestorClosure := range ancestorClosures {
			closure := domain.NodeClosure{
				Ancestor:   ancestorClosure.Ancestor,
				Descendant: createdNode.ID,
				Depth:      depth,
			}
			service.NodeClosureRepository.Save(ctx, tx, closure)
			depth++
		}
	}

	// return response
	return dto.ToNodeCreatedResponse(createdNode), nil
}

func (service *NodeServiceImpl) RootList(ctx context.Context) ([]dto.NodeResponse, error) {
	// Get Root Nodes
	rootNodes := service.NodeRepository.GetRootList(ctx, service.DB)

	// return response
	return dto.ToNodePaginationResponse(rootNodes), nil
}

func (service *NodeServiceImpl) DetailNode(ctx context.Context, nodeId string) (dto.NodeResponse, error) {
	// Get Node By ID
	node, _ := service.NodeRepository.GetNodeByID(ctx, service.DB, nodeId)
	if node.ID == uuid.Nil {
		return dto.NodeResponse{}, fmt.Errorf("node not found")
	}

	// return response
	return dto.ToNodeDetailResponse(node), nil
}

func (service *NodeServiceImpl) UpdateNode(ctx context.Context, nodeId string, request dto.NodeUpdateRequest) (dto.NodeResponse, error) {
	// Get Detail Node By ID
	node, _ := service.NodeRepository.GetNodeByID(ctx, service.DB, nodeId)
	if node.ID == uuid.Nil {
		return dto.NodeResponse{}, fmt.Errorf("node not found")
	}

	// Validate request
	err := service.Validate.Struct(request)
	if err != nil {
		return dto.NodeResponse{}, err
	}

	// Start transaction
	tx, err := service.DB.Begin()
	if err != nil {
		return dto.NodeResponse{}, err
	}

	// Defer commit or rollback
	defer pkg.CommitOrRollback(tx)

	// Update Node
	node.Title = request.Title
	node.Type = request.Type
	if request.Description != nil {
		node.Description = sql.NullString{String: *request.Description, Valid: true}
	}
	node.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	updatedNode := service.NodeRepository.Update(ctx, tx, nodeId, node)

	// return response
	return dto.ToNodeDetailResponse(updatedNode), nil
}

func (service *NodeServiceImpl) DeleteNode(ctx context.Context, nodeId string) error {
	// Check Node By ID
	isNodeExist := service.NodeRepository.CheckByID(ctx, service.DB, nodeId)
	if !isNodeExist {
		return fmt.Errorf("node not found")
	}

	// Start transaction
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}

	// Defer commit or rollback
	defer pkg.CommitOrRollback(tx)

	// Get Descendant IDs
	descendantIds := service.NodeClosureRepository.FindDescendantIdsByAncestor(ctx, tx, nodeId)

	// Delete Node Closure : Self with All Descendants
	_ = service.NodeClosureRepository.DeleteByDescendantIds(ctx, tx, descendantIds)

	// Delete Node with All Descendants
	_ = service.NodeRepository.DeleteByDescendantIds(ctx, tx, descendantIds)

	// return response
	return nil
}

func (service *NodeServiceImpl) DescendantList(ctx context.Context, nodeId string) ([]dto.NodeResponse, error) {
	// Check Node By ID
	isNodeExist := service.NodeRepository.CheckByID(ctx, service.DB, nodeId)
	if !isNodeExist {
		return []dto.NodeResponse{}, fmt.Errorf("node not found")
	}

	// Get Descendant Nodes
	descendantNodes := service.NodeRepository.GetDescendantList(ctx, service.DB, nodeId)

	// return response
	return dto.ToNodePaginationResponse(descendantNodes), nil
}

func (service *NodeServiceImpl) MoveNode(ctx context.Context, nodeId string, request dto.NodeMoveRequest) error {
	// Check Node By ID
	isNodeExist := service.NodeRepository.CheckByID(ctx, service.DB, nodeId)
	if !isNodeExist {
		return fmt.Errorf("node not found")
	}

	// Validate request
	err := service.Validate.Struct(request)
	if err != nil {
		return err
	}

	// Check Ancestor Node
	isAncestorNodeExist := service.NodeRepository.CheckByID(ctx, service.DB, request.ToAncestorID)
	if !isAncestorNodeExist {
		return fmt.Errorf("ancestor node is not found")
	}

	// Start transaction
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}

	// Defer commit or rollback
	defer pkg.CommitOrRollback(tx)

	// Get New Path For Node
	newClosures := service.NodeClosureRepository.GetNewClosures(ctx, tx, nodeId, request.ToAncestorID)

	// Get Descendant IDs
	descendantIds := service.NodeClosureRepository.FindDescendantIdsByAncestor(ctx, tx, nodeId)

	// Delete Node Closure : Self with All Descendants
	_ = service.NodeClosureRepository.DeleteByDescendantIds(ctx, tx, descendantIds)

	// Save New Node Closure For Self and All Descendants Under New Ancestor
	for _, closure := range newClosures {
		service.NodeClosureRepository.Save(ctx, tx, closure)
	}

	// return success
	return nil
}
