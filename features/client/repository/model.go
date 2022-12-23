package repository

import (
	client "capstone-alta1/features/client"
	core "capstone-alta1/features/user"
	user "capstone-alta1/features/user/repository"
)

// struct gorm model
type Client struct {
	User           user.User
	Gender         string
	Address        string
	City           string
	Phone          string
	ClientImageUrl string
}

// mapping

// mengubah struct core ke struct model gorm
func fromCore(dataCore client.Core) Client {
	clientGorm := Client{
		User: user.User{
			Name:     dataCore.User.Name,
			Email:    dataCore.User.Email,
			Password: dataCore.User.Password,
			Role:     dataCore.User.Role,
		},
		Gender:         dataCore.Gender,
		Address:        dataCore.Address,
		City:           dataCore.City,
		Phone:          dataCore.Phone,
		ClientImageUrl: dataCore.ClientImageUrl,
	}
	return clientGorm
}

// mengubah struct model gorm ke struct core
func (dataModel *Client) toCore() client.Core {
	return client.Core{
		User: core.Core{
			ID:       dataModel.User.ID,
			Name:     dataModel.User.Name,
			Email:    dataModel.User.Email,
			Password: dataModel.User.Password,
			Role:     dataModel.User.Role,
		},
		Gender:         dataModel.Gender,
		Address:        dataModel.Address,
		City:           dataModel.City,
		Phone:          dataModel.Phone,
		ClientImageUrl: dataModel.ClientImageUrl,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Client) []client.Core {
	var dataCore []client.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
