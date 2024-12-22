package domain

import (
	"ddd-kata-golang/ecommerce/internal/shared/domain"
	"ddd-kata-golang/ecommerce/pkg/errors"

	"github.com/google/uuid"
	"time"
)

type Category struct {
	id          CategoryID
	name        string
	description string
}

func NewCategory(name string, description string) (*Category, error) {
	if name == "" {
		return nil, errors.ErrInvalidCategoryName
	}

	return &Category{
		id:          NewCategoryID(uuid.New().String()),
		name:        name,
		description: description,
	}, nil
}

func (c *Category) ID() CategoryID {
	return c.id
}

func (c *Category) Name() string {
	return c.name
}

func (c *Category) Description() string {
	return c.description
}

type CategoryCreatedEvent struct {
	baseEvent    domain.BaseEvent
	CategoryID   string `json:"category_id"`
	CategoryName string `json:"name"`
}

func NewCategoryCreatedEvent(category *Category) *CategoryCreatedEvent {
	return &CategoryCreatedEvent{
		baseEvent:    domain.NewBaseEvent("category.created"),
		CategoryID:   category.ID().String(),
		CategoryName: category.Name(),
	}
}

func (e *CategoryCreatedEvent) Name() string {
	return e.baseEvent.Name()
}

func (e *CategoryCreatedEvent) Timestamp() time.Time {
	return e.baseEvent.Timestamp()
}
