package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"job-application/dtos"
	helpercontext "job-application/helpers/helper_context"
	"job-application/models"
)

type UserRepository interface {
	IsEmailAlreadyRegistered(context.Context, string) bool
	PostUser(context.Context, dtos.RequestRegisterUser, string) (*models.User, error)
	PostUserProfile(context.Context, dtos.RequestRegisterUser, int64) (*models.UserProfile, error)
	GetUserByEmail(context.Context, string) (*models.User, error)
	GetUserById(context.Context, int64) (*models.User, error)
	GetUserProfileById(context.Context, int64) (*models.UserProfile, error)
}

type UserRepositoryImplementation struct {
	db *sql.DB
}

func NewUserRepositoryImplementation(db *sql.DB) *UserRepositoryImplementation {
	return &UserRepositoryImplementation{
		db: db,
	}
}

func (us *UserRepositoryImplementation) IsEmailAlreadyRegistered(ctx context.Context, emailInput string) bool {
	sql := `
	SELECT
		id,
		name,
		email,
		created_at
	FROM users
	WHERE email = $1 AND deleted_at IS NULL;
`
	var user models.User
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		txFromCtx.QueryRowContext(ctx, sql, emailInput).Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
		)
	} else {
		us.db.QueryRowContext(ctx, sql, emailInput).Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
		)
	}
	return user.ID != 0
}

func (us *UserRepositoryImplementation) GetUserById(ctx context.Context, userId int64) (*models.User, error) {
	sql := `
	SELECT
		*
	FROM users
	WHERE id = $1 AND deleted_at IS NULL;
	`
	var user models.User
	var err error
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err = txFromCtx.QueryRowContext(ctx, sql, userId).Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeleteAt,
		)
	} else {
		err = us.db.QueryRowContext(ctx, sql, userId).Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeleteAt,
		)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}

func (us *UserRepositoryImplementation) GetUserProfileById(ctx context.Context, userId int64) (*models.UserProfile, error) {
	sql := `
	SELECT
		*
	FROM UserProfiles
	WHERE user_id = $1 AND deleted_at IS NULL;
	`
	var userProfile models.UserProfile
	var err error
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err = txFromCtx.QueryRowContext(ctx, sql, userId).Scan(
			&userProfile.ID,
			&userProfile.UserId,
			&userProfile.Age,
			&userProfile.CurrentJob,
			&userProfile.Address,
			&userProfile.CreatedAt,
			&userProfile.UpdatedAt,
			&userProfile.DeleteAt,
		)
	} else {
		err = us.db.QueryRowContext(ctx, sql, userId).Scan(
			&userProfile.ID,
			&userProfile.UserId,
			&userProfile.Age,
			&userProfile.CurrentJob,
			&userProfile.Address,
			&userProfile.CreatedAt,
			&userProfile.UpdatedAt,
			&userProfile.DeleteAt,
		)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &userProfile, nil
}

func (us *UserRepositoryImplementation) PostUser(ctx context.Context, reqBody dtos.RequestRegisterUser, hashPassword string) (*models.User, error) {
	sql := `
		INSERT INTO users (name, email, password, role, created_at, updated_at) VALUES 
		($1, $2, $3, $4, NOW(), NOW())
		RETURNING *;
	`
	var user models.User
	var err error
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err = txFromCtx.QueryRowContext(ctx, sql, reqBody.Name, reqBody.Email, hashPassword, "applicant").Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeleteAt,
		)
	} else {
		err = us.db.QueryRowContext(ctx, sql, reqBody.Name, reqBody.Email, hashPassword, "applicant").Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeleteAt,
		)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}

func (us *UserRepositoryImplementation) PostUserProfile(ctx context.Context, reqBody dtos.RequestRegisterUser, userId int64) (*models.UserProfile, error) {
	sql := `
		INSERT INTO UserProfiles (user_id, age, current_job, address, created_at, updated_at) VALUES 
		($1, $2, $3, $4, NOW(), NOW())
		RETURNING *;
	`
	var userProfile models.UserProfile
	var err error
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err = txFromCtx.QueryRowContext(ctx, sql, userId, reqBody.Age, reqBody.CurrentJob, reqBody.Address).Scan(
			&userProfile.ID,
			&userProfile.UserId,
			&userProfile.Age,
			&userProfile.CurrentJob,
			&userProfile.Address,
			&userProfile.CreatedAt,
			&userProfile.UpdatedAt,
			&userProfile.DeleteAt,
		)
	} else {
		err = us.db.QueryRowContext(ctx, sql, userId, reqBody.Age, reqBody.CurrentJob, reqBody.Address).Scan(
			&userProfile.ID,
			&userProfile.UserId,
			&userProfile.Age,
			&userProfile.CurrentJob,
			&userProfile.Address,
			&userProfile.CreatedAt,
			&userProfile.UpdatedAt,
			&userProfile.DeleteAt,
		)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &userProfile, nil
}

func (us *UserRepositoryImplementation) GetUserByEmail(ctx context.Context, emailInput string) (*models.User, error) {
	sql := `
		SELECT
		*
		FROM users
		WHERE email = $1 AND deleted_at IS NULL;
	`
	var user models.User
	var err error
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err = txFromCtx.QueryRowContext(ctx, sql, emailInput).Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeleteAt,
		)
	} else {
		err = us.db.QueryRowContext(ctx, sql, emailInput).Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeleteAt,
		)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}
