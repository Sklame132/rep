package domain

import (
	"fmt"
	"regexp"
	"time"

	core_errors "github.com/Sklame132/rep/internal/core/errors"
)

type User struct {
	ID          string
	Username    string
	Password    string
	FirstName   string
	LastName    string
	Address     *string
	Email       *string
	PhoneNumber *string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	Rating      *int16
	Role        *string
	ImageURL    *string
}

func NewUser(
	id string,
	username string,
	password string,
	firstName string,
	lastName string,
	address *string,
	email *string,
	phoneNumber *string,
	createdAt time.Time,
	updatedAt *time.Time,
	rating *int16,
	role *string,
	imageURL *string,
) User {
	return User{
		ID:          id,
		Username:    username,
		Password:    password,
		FirstName:   firstName,
		LastName:    lastName,
		Address:     address,
		Email:       email,
		PhoneNumber: phoneNumber,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		Rating:      rating,
		Role:        role,
		ImageURL:    imageURL,
	}
}

func NewUserUninitialized(
	username string,
	password string,
	firstName string,
	lastName string,
	address *string,
	email *string,
	phoneNumber *string,
	rating *int16,
	role *string,
) User {
	return NewUser(
		UninitializedString,
		username,
		password,
		firstName,
		lastName,
		address,
		email,
		phoneNumber,
		UninitializedTime,
		&UninitializedTime,
		rating,
		role,
		&UninitializedString,
	)
}

func (u *User) Validate() error {
	if u.Email == nil && u.PhoneNumber == nil {
		return fmt.Errorf("`Email` and `PhoneNumber` is undefined: %w", core_errors.ErrInvalidArgument)
	}

	if u.Email != nil {
		emailLen := len([]rune(*u.Email))
		if emailLen > 32 {
			return fmt.Errorf("invalid `Email` length: %d: %w", emailLen, core_errors.ErrInvalidArgument)
		}

		re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

		if !re.MatchString(*u.Email) {
			return fmt.Errorf("invalid `Email` format: %w", core_errors.ErrInvalidArgument)
		}
	}

	if u.PhoneNumber != nil {
		phoneNumberLen := len([]rune(*u.PhoneNumber))
		if phoneNumberLen < 10 || phoneNumberLen > 15 {
			return fmt.Errorf("invalid `PhoneNumber` length: %d: %w", phoneNumberLen, core_errors.ErrInvalidArgument)
		}

		re := regexp.MustCompile(`^\+[0-9]+$`)

		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf(
				"invalid `PhoneNumber` format: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if u.Rating == nil {
		u.Rating = &UnitializedRating
	} else if *u.Rating < 0 {
		return fmt.Errorf(
			"`Rating` value must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if u.Role == nil {
		u.Role = &UninitializedRole
	}

	return nil
}

type UserPatch struct {
	Username    Nullable[string]
	Password    Nullable[string]
	FirstName   Nullable[string]
	LastName    Nullable[string]
	Address     Nullable[string]
	Email       Nullable[string]
	PhoneNumber Nullable[string]
	Rating      Nullable[int16]
	Role        Nullable[string]
}

func NewUserPatch(
	username    Nullable[string],
	password    Nullable[string],
	firstName   Nullable[string],
	lastName    Nullable[string],
	address     Nullable[string],
	email       Nullable[string],
	phoneNumber Nullable[string],
	rating      Nullable[int16],
	role        Nullable[string],
) UserPatch {
	return UserPatch {
		Username:    username,
		Password:    password,
		FirstName:   firstName,
		LastName:    lastName,
		Address:     address,
		Email:       email,
		PhoneNumber: phoneNumber,
		Rating:      rating,
		Role:        role,
	}
}

func (p *UserPatch) Validate() error {
	if p.Username.Set && p.Username.Value == nil {
		return fmt.Errorf(
			"Username can't be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	if p.Password.Set && p.Password.Value == nil {
		return fmt.Errorf(
			"Password can't be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	if p.FirstName.Set && p.FirstName.Value == nil {
		return fmt.Errorf(
			"FirstName can't be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	if p.LastName.Set && p.LastName.Value == nil {
		return fmt.Errorf(
			"LastName can't be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	if (p.Email.Set && p.Email.Value == nil) && (p.PhoneNumber.Set && p.PhoneNumber.Value == nil) {
		return fmt.Errorf(
			"Email and PhoneNumber can't be patched to NULL at the same time: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	if p.Rating.Set && p.Rating.Value == nil {
		return fmt.Errorf(
			"Email and PhoneNumber can't be patched to NULL at the same time: %w",
			core_errors.ErrInvalidArgument,
		)
	}

		return nil
}

func (u *User) ApplyPatch(patch UserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate user patch: %w", err)
	}

	tmp := *u
	
	if patch.Username.Set{
		tmp.Username = *patch.Username.Value
	}
	if patch.Password.Set{
		tmp.Password = *patch.Password.Value
	}
	if patch.FirstName.Set{
		tmp.FirstName = *patch.FirstName.Value
	}
	if patch.LastName.Set{
		tmp.LastName = *patch.LastName.Value
	}
	if patch.Address.Set {
		tmp.Address = patch.Address.Value
	}
	if patch.Email.Set{
		tmp.Email = patch.Email.Value
	}
	if patch.PhoneNumber.Set{
		tmp.PhoneNumber = patch.PhoneNumber.Value
	}
	if patch.Rating.Set{
		tmp.Rating = patch.Rating.Value
	}
	if patch.Role.Set {
		tmp.Role = patch.Role.Value
	}


	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched user: %w", err)
	}

	*u = tmp

	return nil
}
