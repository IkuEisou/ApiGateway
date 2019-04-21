package main

import (
	"database/sql"
	"github.com/goadesign/goa"
	"github.com/ikueisou/apigateway/auth/repositories"
	"github.com/ikueisou/apigateway/auth/src/app"
	"github.com/ikueisou/apigateway/auth/utils/crypto"
	"github.com/ikueisou/apigateway/auth/utils/jwt"
)

// UserController implements the User resource.
type UserController struct {
	*goa.Controller
	*sql.DB
}

// NewUserController creates a User controller.
func NewUserController(service *goa.Service, db *sql.DB) *UserController {
	return &UserController{
		Controller: service.NewController("UserController"),
		DB:         db,
	}
}

// Login runs the login action.
func (c *UserController) Login(ctx *app.LoginUserContext) error {
	// UserController_Login: start_implement
	payload := ctx.Payload
	u, err := repositories.GetUserByEmail(c.DB, payload.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.BadRequest(goa.ErrBadRequest("Invalid email or password"))
		}
		c.Service.LogError("Login User", "err", err)
		return ctx.InternalServerError()
	}
	hashedPassword := crypto.HashPassword(payload.Password, *u.Salt)
	if *u.Password != hashedPassword {
		return ctx.BadRequest(goa.ErrBadRequest("Invalid email or password"))
	}
	token, err := jwt.CreateJWTToken(*u.Email)
	if err != nil {
		c.Service.LogError("Login User", "err", err)
		return ctx.InternalServerError()
	}

	res := &app.Token{Token: &token}
	return ctx.OK(res)
	// UserController_Login: end_implement
}

// Register runs the register action.
func (c *UserController) Register(ctx *app.RegisterUserContext) error {
	// UserController_Register: start_implement
	payload := ctx.Payload
	exists, err := repositories.CheckEmailExists(c.DB, payload.Email)
	if err != nil {
		c.Service.LogError("Register User", "err", err)
		return ctx.InternalServerError()
	}
	if exists {
		return ctx.BadRequest(goa.ErrBadRequest("Email already exists"))
	}
	err = repositories.AddUserToDatabase(c.DB, payload.FirstName, payload.LastName, payload.Email, payload.Password)
	if err != nil {
		c.Service.LogError("Register User", "err", err)
		return ctx.InternalServerError()
	}
	token, err := jwt.CreateJWTToken(payload.Email)
	if err != nil {
		c.Service.LogError("Register User", "err", err)
		return ctx.InternalServerError()
	}

	res := &app.Token{Token: &token}
	return ctx.OK(res)
	// UserController_Register: end_implement
}
