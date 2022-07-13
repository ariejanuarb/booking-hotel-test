package repository

import (
	"context"
	"database/sql"
	"programmerzamannow/belajar-golang-restful-api/model/domain"
)

type UserRoleRepository interface {
	Save(ctx context.Context, tx *sql.Tx, userRole domain.UserRole) domain.UserRole
	Update(ctx context.Context, tx *sql.Tx, userRole domain.UserRole) domain.UserRole
	Delete(ctx context.Context, tx *sql.Tx, userRole domain.UserRole)
	FindById(ctx context.Context, tx *sql.Tx, userRoleId int) (domain.UserRole, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.UserRole
}
