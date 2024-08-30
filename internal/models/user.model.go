package models

var TableUser = `
CREATE TABLE public.users (
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	first_name varchar(20) NULL,
	last_name varchar(20) NULL,
	email varchar(50) NULL,
	"password" varchar(255) NULL,
	"role" varchar NULL,
	image text NULL,
	phone_number varchar(15) NULL,
	point varchar(50) NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	is_deleted bool NULL,
	CONSTRAINT unique_email UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
`

type User struct {
	Id           string `db:"id" json:"id" form:"id" valid:"-"`
	First_name   string `db:"first_name" json:"first_name" form:"first_name" valid:"-"`
	Last_name    string `db:"last_name" json:"last_name" form:"last_name" valid:"-"`
	Email        string `db:"email" json:"email" form:"email" valid:"email"`
	Password     string `db:"password" json:"password" form:"password" valid:"stringlength(5|256)~Password minimal 5 karakter"`
	Role         string `db:"role" json:"role" form:"role" valid:"-"`
	Image        string `db:"image" json:"image" valid:"-"`
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
	Id       string  `db:"id" json:"id" form:"id" valid:"-"`
	Email    string  `db:"email" json:"email" form:"email" valid:"email"`
	Password string  `db:"password" json:"password" form:"password" valid:"stringlength(5|256)~Password minimal 5 karakter"`
	Role     string  `db:"role" json:"role" form:"role" valid:"-"`
	Image    *string `db:"image" json:"image" form:"image" valid:"-"`
}

type Users []UserAll
