package queries

import (
	"context"
	"myapp/internal/user/domain"
)

type GetUserByEmailQuery struct {
	Email string
}

type GetUserByEmailHandler struct {
	repo domain.UserRepository
}

func NewGetUserByEmailHandler(r domain.UserRepository) *GetUserByEmailHandler {
	return &GetUserByEmailHandler{repo: r}
}

func (h *GetUserByEmailHandler) Handle(ctx context.Context, q GetUserByEmailQuery) (*domain.User, error) {
	return h.repo.FindByEmail(ctx, q.Email)
}
