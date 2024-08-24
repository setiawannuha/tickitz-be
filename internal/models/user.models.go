package models

var TableUser = `
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(20),
    last_name VARCHAR(20),
    email VARCHAR(50),
    password VARCHAR(255),
    role VARCHAR,
    image TEXT,
    phone_number VARCHAR(15),
    point VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`

type User struct {
	Id           string `db:"id" json:"id" form:"id"`
	First_name   string `db:"first_name" json:"first_name" form:"first_name" valid:"-"`
	Last_name    string `db:"last_name" json:"last_name" form:"last_name"`
	Email        string `db:"email" json:"email" form:"email"`
	Password     string `db:"password" json:"password" form:"password"`
	Role         string `db:"role" json:"role" form:"role"`
	Image        string `db:"image" json:"image" form:"image"`
	Phone_number string `db:"phone_number" json:"phone_number" form:"phone_number"`
	Point        string `db:"point" json:"point" form:"point"`
}

type UserDetails struct {
	Id           string  `db:"id" json:"id" form:"id"`
	First_name   *string `db:"first_name" json:"first_name" form:"first_name" valid:"-"`
	Last_name    *string `db:"last_name" json:"last_name" form:"last_name"`
	Email        string  `db:"email" json:"email" form:"email"`
	Image        *string `db:"image" json:"image" form:"image"`
	Phone_number *string `db:"phone_number" json:"phone_number" form:"phone_number"`
	Point        *string `db:"point" json:"point" form:"point"`
}

type UserAll struct {
	Id         string  `db:"id" json:"id" form:"id"`
	First_name *string `db:"first_name" json:"first_name" form:"first_name"`
	Last_name  *string `db:"last_name" json:"last_name" form:"last_name"`
	Email      string  `db:"email" json:"email" form:"email"`
}

type UserLogin struct {
	Id       string `db:"id" json:"id" form:"id"`
	Email    string `db:"email" json:"email" form:"email"`
	Password string `db:"password" json:"password" form:"password"`
}

type Users []UserAll