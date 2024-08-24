package repository

import (
	"database/sql"
	"fmt"
	"khalifgfrz/coffee-shop-be-go/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepositoryInterface interface {
	CreateData(body *models.User)(string , error)
    UpdateData(data *models.User, id string) (*models.User, error)
	GetAllData()(*models.Users , error)
	GetDetailData(id string)(*models.UserDetails , error)
	DeleteData(id string)(string , error) 
}


type UserRepository struct {
	*sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateData(body *models.User)(string , error){
	query := `INSERT INTO users (email , password , role ) VALUES ( $1 , $2 , $3 )`

	_, err := r.Exec(query,body.Email , body.Password , body.Role )
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("congratulations account %s has been registered",body.Email) , nil
}

func (r *UserRepository) UpdateData(data *models.User, id string) (*models.User, error) {
    query := `UPDATE users SET `
    var values []interface{}
    condition := false

    if data.First_name != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`first_name = $%d`, len(values)+1)
        values = append(values, data.First_name)
        condition = true
    }

    if data.Last_name != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`last_name = $%d`, len(values)+1)
        values = append(values, data.Last_name)
        condition = true
    }

    if data.Email != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`email = $%d`, len(values)+1)
        values = append(values, data.Email)
        condition = true
    }

    if data.Password != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`password = $%d`, len(values)+1)
        values = append(values, data.Password)
        condition = true
    }

    if data.Image != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`image = $%d`, len(values)+1)
        values = append(values, data.Image)
        condition = true
    }

    if data.Phone_number != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`phone_number = $%d`, len(values)+1)
        values = append(values, data.Phone_number)
        condition = true
    }

    if data.Point != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`point = $%d`, len(values)+1)
        values = append(values, data.Point)
        condition = true
    }

    if !condition {
        return nil, fmt.Errorf("no fields to update")
    }

    query += fmt.Sprintf(` WHERE id = $%d RETURNING first_name, last_name, email, password, image, phone_number, point`, len(values)+1)
    values = append(values, id)

    row := r.DB.QueryRow(query, values...)

    var user models.User
    err := row.Scan(
        &user.First_name,
        &user.Last_name,
        &user.Email,
        &user.Password,
        &user.Image,
        &user.Phone_number,
        &user.Point,
    )

    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("user with id = %s not found", id)
        }
        return nil, fmt.Errorf("query execution error: %w", err)
    }

    return &user, nil
}

func (r *UserRepository) GetAllData()(*models.Users , error){
	query := `select id , first_name , last_name , email  FROM users where is_deleted = FALSE`
	data := models.Users{}

	err := r.Select(&data, query)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *UserRepository) GetDetailData(id string)(*models.UserDetails , error){
	query := `select id , first_name , last_name , email , image , phone_number , point FROM users WHERE id = $1 AND is_deleted = FALSE `
	data := models.UserDetails{}

	err := r.Get(&data, query , id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *UserRepository) DeleteData(id string)(string , error) {
	query := `UPDATE users SET is_deleted = true WHERE id = $1`
	_, err := r.Exec(query,id )
	if err != nil {
		return "", err
	}

	return "Delete successful" , nil
}