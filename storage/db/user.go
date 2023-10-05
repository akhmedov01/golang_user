package db

import (
	"context"
	"fmt"
	"main/models"
	"main/pkg/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {

	return &userRepo{
		db: db,
	}

}

func (u *userRepo) Create(ctx context.Context, req models.CreateUser) (string, error) {

	id := uuid.NewString()

	query := `
	INSERT INTO 
		users(id,name,age,login,password) 
	VALUES($1,$2,$3,$4,$5)`

	_, err := u.db.Exec(ctx, query,
		id,
		req.Name,
		req.Age,
		req.Login,
		req.Password,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "", err
	}

	return id, nil

}

func (u *userRepo) GetByLoging(ctx context.Context, req models.GetByLoginReq) (models.GetByLoginRes, error) {

	query := `
	SELECT
	id,
	password
	FROM users WHERE login = $1`

	resp := u.db.QueryRow(ctx, query, req.Login)

	var password models.GetByLoginRes

	err := resp.Scan(
		&password.Id,
		&password.Password,
	)

	if err != nil {
		return models.GetByLoginRes{}, err
	}

	return password, nil
}

func (u *userRepo) Update(ctx context.Context, id string, req models.UpdateUser) (string, error) {

	query := `
	UPDATE users
	SET name=$2,age=$3,login=$4,updated_at=NOW()
	WHERE id=$1`

	resp, err := u.db.Exec(ctx, query,
		id,
		req.Name,
		req.Age,
		req.Login,
	)
	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}
	return "OK", nil
}

func (u *userRepo) Get(ctx context.Context, req models.IdRequest) (models.User, error) {

	query := `
	SELECT
	id,
	name,
	age,
	login,
	created_at::text,
	updated_at::text
	FROM users WHERE id = $1`

	resp := u.db.QueryRow(ctx, query, req.Id)

	var user models.User

	err := resp.Scan(
		&user.Id,
		&user.Name,
		&user.Age,
		&user.Login,
		&user.Created_at,
		&user.Updated_at,
	)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *userRepo) GetAll(ctx context.Context, req models.GetAllUserRequest) (models.GetAllUser, error) {

	var (
		params  = make(map[string]interface{})
		filter  = "WHERE true"
		offsetQ = " OFFSET 0 "
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
	)

	s := `
	SELECT
	id,
	name,
	age,
	login,
	created_at::text,
	updated_at::text
	FROM users
	`
	if req.Name != "" {
		filter += ` AND name ILIKE '%' || @name || '%' `
		params["name"] = req.Name
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf(" OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ

	q, pArr := helper.ReplaceQueryParams(query, params)
	rows, err := u.db.Query(ctx, q, pArr...)
	if err != nil {
		fmt.Println("eeeeeeeeeeeeeeeerrrrrrrrrrrrrrrrrrrrooooooooooooooooorrrrrrrrrrrrrrrr")
		return models.GetAllUser{}, err
	}

	defer rows.Close()

	result := []models.User{}

	for rows.Next() {

		var user models.User

		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Age,
			&user.Login,
			&user.Created_at,
			&user.Updated_at,
		)
		if err != nil {
			return models.GetAllUser{}, err
		}

		result = append(result, user)

	}

	return models.GetAllUser{Users: result, Count: len(result)}, nil

}

func (u *userRepo) Delete(ctx context.Context, req models.IdRequest) (string, error) {

	query := `DELETE FROM users WHERE id = $1`

	resp, err := u.db.Exec(ctx, query,
		req.Id,
	)

	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}

	return "Deleted suc", nil
}
