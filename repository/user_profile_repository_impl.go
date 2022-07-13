package repository

import (
	"context"
	"database/sql"
	"errors"
	"programmerzamannow/belajar-golang-restful-api/helper"
	"programmerzamannow/belajar-golang-restful-api/model/domain"
)

type UserProfileRepositoryImpl struct {
}

func NewUserProfileRepository() UserProfileRepository {
	return &UserProfileRepositoryImpl{}
}

func (repository *UserProfileRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, userProfile domain.UserProfile) domain.UserProfile {
	SQL := "insert into user_profile(email, password, name, gender, assigned_to_hotel) values (?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, userProfile.Email, userProfile.Password, userProfile.Name, userProfile.Gender, userProfile.AssignedToHotel)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	userProfile.Id = int(id)
	return userProfile
}

func (repository *UserProfileRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, userProfile domain.UserProfile) domain.UserProfile {
	SQL := "update user_profile set name = ?, gender = ?, email = ?, password = ?, assigned_to_hotel = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, userProfile.Email, userProfile.Password, userProfile.Name, userProfile.Gender, userProfile.AssignedToHotel, userProfile.Id)
	helper.PanicIfError(err)

	return userProfile
}

func (repository *UserProfileRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userProfile domain.UserProfile) {
	SQL := "delete from user_profile where id = ?"
	_, err := tx.ExecContext(ctx, SQL, userProfile.Id)
	helper.PanicIfError(err)
}

func (repository *UserProfileRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userProfileId int) (domain.UserProfile, error) {
	SQL := "select id, name, gender, email, password, assigned_to_hotel from user_profile where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userProfileId)
	helper.PanicIfError(err)
	defer rows.Close()

	userProfile := domain.UserProfile{}
	if rows.Next() {
		err := rows.Scan(&userProfile.Id, &userProfile.Name, &userProfile.Gender, &userProfile.Email, &userProfile.Password, &userProfile.AssignedToHotel)
		helper.PanicIfError(err)
		return userProfile, nil
	} else {
		return userProfile, errors.New("userProfile is not found")
	}
}

func (repository *UserProfileRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.UserProfile {
	SQL := "select id, name, gender, email, password, assigned_to_hotel from user_profile"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var userProfiles []domain.UserProfile
	for rows.Next() {
		userProfile := domain.UserProfile{}
		err := rows.Scan(&userProfile.Id, &userProfile.Name, &userProfile.Gender, &userProfile.Email, &userProfile.Password, &userProfile.AssignedToHotel)
		helper.PanicIfError(err)
		userProfiles = append(userProfiles, userProfile)
	}
	return userProfiles
}
