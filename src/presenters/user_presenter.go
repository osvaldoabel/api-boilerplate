package presenters

import (
	"osvaldoabel/users-api/src/domain"
	"osvaldoabel/users-api/utils"
)

/*
	@desctription: Receives a slice of User
		formats it as necessary and returns a Slice of UsesrPayloads

	@return: userCollection
*/
func ToCollection(items []*domain.User) []*utils.UserCollection {
	var userCollection []*utils.UserCollection

	for _, item := range items {
		userCollection = append(userCollection, ToArray(item))
	}

	return userCollection
}

func ToArray(item *domain.User) *utils.UserCollection {

	return &utils.UserCollection{
		ID:      item.ID,
		Name:    item.Email,
		Email:   item.Email,
		Address: item.Address,
		Age:     item.Age,
		Status:  item.Status,
	}
}
