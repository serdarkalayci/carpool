package mappers

import (
	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/dto"
	"github.com/serdarkalayci/carpool/api/domain"
)

// MapAddUserRequest2User maps dto UserRequest to domain User
func MapAddUserRequest2User(ur dto.AddUserRequest) domain.User {
	return domain.User{
		Name:     ur.Name,
		Password: ur.Password,
		Email:    ur.Email,
	}
}

// MapUser2SUserResponse maps domain User to dto ShortUserResponse
func MapUser2SUserResponse(u domain.User) dto.ShortUserResponse {
	return dto.ShortUserResponse{
		ID:   u.ID,
		Name: u.Name,
	}
}

// MapUser2LUserResponse maps domain User to dto ShortUserResponse
func MapUser2LUserResponse(u domain.User) dto.LongUserResponse {
	return dto.LongUserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Phone: u.Phone,
	}
}
