package handlers

import (
	"github.com/eduardor2m/go-mastery/src/archive/testing_user_examples/database"
	"github.com/eduardor2m/go-mastery/src/archive/testing_user_examples/entity/user"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

type UserRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	BirthDate string `json:"birth_date"`
}

type UserDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	BirthDate string `json:"birth_date,omitempty"`
}

type ErrorDTO struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (instance UserHandler) Create(context echo.Context) error {
	var userRequest UserRequest

	if err := context.Bind(&userRequest); err != nil {
		return context.JSON(400, ErrorDTO{
			Message: "Invalid request body",
			Code:    400,
		})
	}

	userUUID, err := uuid.NewUUID()
	if err != nil {
		return context.JSON(500, ErrorDTO{
			Message: err.Error(),
			Code:    500,
		})
	}

	userFormatted, err := user.NewBuilder().WithID(userUUID).WithName(userRequest.Name).WithEmail(userRequest.Email).WithPassword(userRequest.Password).WithBirthDate(userRequest.BirthDate).Build()
	if err != nil {
		return context.JSON(400, ErrorDTO{
			Message: err.Error(),
			Code:    400,
		})
	}

	sqlite := database.SQLite{}

	conn, err := sqlite.Connect()
	if err != nil {
		return context.JSON(500, ErrorDTO{
			Message: err.Error(),
			Code:    500,
		})
	}

	defer conn.Close()

	result, err := conn.ExecContext(context.Request().Context(), "INSERT INTO users (id, name, email, password, birth_date) VALUES (?, ?, ?, ?, ?)", userFormatted.ID().String(), userFormatted.Name(), userFormatted.Email(), userFormatted.Password(), userFormatted.BirthDate())
	if err != nil {
		return context.JSON(500, ErrorDTO{
			Message: err.Error(),
			Code:    500,
		})
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return context.JSON(500, ErrorDTO{
			Message: err.Error(),
			Code:    500,
		})
	}

	if rowsAffected == 0 {
		return context.JSON(500, ErrorDTO{
			Message: "User not created",
			Code:    500,
		})
	}

	return context.JSON(201, UserDTO{
		ID:        userFormatted.ID().String(),
		Name:      userFormatted.Name(),
		Email:     userFormatted.Email(),
		BirthDate: userFormatted.BirthDate(),
	})
}

func (instance UserHandler) FindAll(context echo.Context) error {
	conn, err := database.SQLite{}.Connect()

	if err != nil {
		return context.JSON(500, ErrorDTO{
			Message: err.Error(),
			Code:    500,
		})
	}

	defer conn.Close()

	rows, err := conn.QueryContext(context.Request().Context(), "SELECT id, name, email, birth_date FROM users")
	if err != nil {
		return context.JSON(500, ErrorDTO{
			Message: err.Error(),
			Code:    500,
		})
	}

	var users []UserDTO

	for rows.Next() {
		var id string
		var name string
		var email string
		var birthDate string

		if err := rows.Scan(&id, &name, &email, &birthDate); err != nil {
			return context.JSON(500, ErrorDTO{
				Message: err.Error(),
				Code:    500,
			})
		}

		users = append(users, UserDTO{
			ID:        id,
			Name:      name,
			Email:     email,
			BirthDate: birthDate,
		})
	}

	return context.JSON(200, users)
}
