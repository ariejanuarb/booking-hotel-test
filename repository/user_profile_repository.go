package repository

import (
	"context"
	"database/sql"
	"programmerzamannow/belajar-golang-restful-api/model/domain"
)

type UserProfileRepository interface {
	Save(ctx context.Context, tx *sql.Tx, userProfile domain.UserProfile) domain.UserProfile
	Update(ctx context.Context, tx *sql.Tx, userProfile domain.UserProfile) domain.UserProfile
	Delete(ctx context.Context, tx *sql.Tx, userProfile domain.UserProfile)
	FindById(ctx context.Context, tx *sql.Tx, userProfileId int) (domain.UserProfile, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.UserProfile
}
