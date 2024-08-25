package models

var TableUser = `
CREATE TABLE public.users (
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
	constraint user_pk primary key(id)
);`

type User struct {
	Id           string `db:"id" json:"id" form:"id" valid:"-"`
	First_name   string `db:"first_name" json:"first_name" form:"first_name" valid:"-"`
	Last_name    string `db:"last_name" json:"last_name" form:"last_name" valid:"-"`
	Email        string `db:"email" json:"email" form:"email" valid:"email"`
	Password     string `db:"password" json:"password" form:"password" valid:"stringlength(5|256)~Password minimal 5 karakter"`
	Role         string `db:"role" json:"role" form:"role" valid:"-"`
	Image        string `db:"image" json:"image" form:"image" valid:"-"`
	Phone_number string `db:"phone_number" json:"phone_number" form:"phone_number" valid:"-"`
	Point        string `db:"point" json:"point" form:"point" valid:"-"`
}

type UserDetails struct {
	Id           string  `db:"id" json:"id" form:"id"`
	First_name   *string `db:"first_name" json:"first_name" form:"first_name" valid:"-"`
	Last_name    *string `db:"last_name" json:"last_name" form:"last_name" valid:"-"`
	Email        string  `db:"email" json:"email" form:"email" valid:"email"`
	Image        *string `db:"image" json:"image" form:"image" valid:"-"`
	Phone_number *string `db:"phone_number" json:"phone_number" form:"phone_number" valid:"-"`
	Point        *string `db:"point" json:"point" form:"point" valid:"-"`
}

type UserAll struct {
	Id         string  `db:"id" json:"id" form:"id" valid:"-"`
	First_name *string `db:"first_name" json:"first_name" form:"first_name" valid:"-"`
	Last_name  *string `db:"last_name" json:"last_name" form:"last_name" valid:"-"`
	Email      string  `db:"email" json:"email" form:"email" valid:"email"`
}

type Auth struct {
	Id       string `db:"id" json:"id" form:"id" valid:"-"`
	Email    string `db:"email" json:"email" form:"email" valid:"email"`
	Password string `db:"password" json:"password" form:"password" valid:"stringlength(5|256)~Password minimal 5 karakter"`
	Role     string `db:"role" json:"role" form:"role" valid:"-"`
}

type Users []UserAll