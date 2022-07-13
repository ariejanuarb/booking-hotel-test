package repository

import (
	"context"
	"database/sql"
	"errors"
	"programmerzamannow/belajar-golang-restful-api/helper"
	"programmerzamannow/belajar-golang-restful-api/model/domain"
)

type UserRoleRepositoryImpl struct {
}

func NewUserRoleRepository() UserRoleRepository {
	return &UserRoleRepositoryImpl{}
}

func (repository *UserRoleRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, userRole domain.UserRole) domain.UserRole {
	SQL := "insert into user_role(role_id, user_profile_id) values (?,?)"
	result, err := tx.ExecContext(ctx, SQL, userRole.RoleId, userRole.UserProfileId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	userRole.Id = int(id)
	return userRole
}

func (repository *UserRoleRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, userRole domain.UserRole) domain.UserRole {
	SQL := "update user_role set role_id = ?, user_profile_id = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, userRole.RoleId, userRole.UserProfileId, userRole.Id)
	helper.PanicIfError(err)

	return userRole
}

func (repository *UserRoleRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userRole domain.UserRole) {
	SQL := "delete from user_role where id = ?"
	_, err := tx.ExecContext(ctx, SQL, userRole.Id)
	helper.PanicIfError(err)
}

func (repository *UserRoleRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userRoleId int) (domain.UserRole, error) {
	SQL := "select id, role_id, user_profile_id from user_role where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userRoleId)
	helper.PanicIfError(err)
	defer rows.Close()

	userRole := domain.UserRole{}
	if rows.Next() {
		err := rows.Scan(&userRole.Id, &userRole.RoleId, &userRole.UserProfileId)
		helper.PanicIfError(err)
		return userRole, nil
	} else {
		return userRole, errors.New("userRole is not found")
	}
}

func (repository *UserRoleRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.UserRole {
	SQL := "select id, role_id, user_profile_id from user_role"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var userRoles []domain.UserRole
	for rows.Next() {
		userRole := domain.UserRole{}
		err := rows.Scan(&userRole.Id, &userRole.RoleId, &userRole.UserProfileId)
		helper.PanicIfError(err)
		userRoles = append(userRoles, userRole)
	}
	return userRoles
}
