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
func ToArray(items []*domain.User) []*utils.UserCollection {
	var userCollection []*utils.UserCollection

	for _, item := range items {
		userCollection = append(userCollection, &utils.UserCollection{
			ID:      item.ID,
			Name:    item.Email,
			Email:   item.Email,
			Address: item.Address,
			Age:     item.Age,
			Status:  item.Status,
		})
	}

	return userCollection
}
