package users_transport_http

import "github.com/Alv1ol/Todoapp/internal/core/domain"

type UserDTOResponce struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func userDTOFromDomain(user domain.User) UserDTOResponce{
	return UserDTOResponce{
		ID: user.ID,
		Version: user.Version,
		FullName: user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}

func usersDTOFromDomains(users []domain.User) []UserDTOResponce{
	usersDTO := make([]UserDTOResponce, len(users))

	for i, user := range users{
		usersDTO[i] = userDTOFromDomain(user)
	}

	return usersDTO
}