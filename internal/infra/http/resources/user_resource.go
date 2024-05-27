package resources

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type UserDto struct {
	Id         uint64      `json:"id"`
	FirstName  string      `json:"firstName"`
	SecondName string      `json:"secondName"`
	Email      string      `json:"email"`
	Role       domain.Role `json:"role,omitempty"`
}

type AuthDto struct {
	Token string  `json:"token"`
	User  UserDto `json:"user"`
}

type UsersDto struct {
	Items []UserDto `json:"items"`
	Total uint64    `json:"total"`
	Pages uint      `json:"pages"`
}

func (d UserDto) DomainToDto(user domain.User) UserDto {
	return UserDto{
		Id:         user.Id,
		FirstName:  user.FirstName,
		SecondName: user.SecondName,
		Email:      user.Email,
		Role:       user.Role,
	}
}

func (d UserDto) DomainToDtoCollection(users []domain.User) []UserDto {
	result := make([]UserDto, len(users))
	for i, u := range users {
		result[i] = d.DomainToDto(u)
	}
	return result
}

func (d AuthDto) DomainToDto(token string, user domain.User) AuthDto {
	var userDto UserDto
	return AuthDto{
		Token: token,
		User:  userDto.DomainToDto(user),
	}
}
