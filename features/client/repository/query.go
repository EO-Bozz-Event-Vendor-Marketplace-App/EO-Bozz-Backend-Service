package repository

import (
	cfg "capstone-alta1/config"
	client "capstone-alta1/features/client"
	"capstone-alta1/utils/helper"
	"capstone-alta1/utils/thirdparty"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type clientRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) client.RepositoryInterface {
	return &clientRepository{
		db: db,
	}
}

// Create implements user.Repository
func (repo *clientRepository) Create(input client.Core) error {
	clientGorm := fromCore(input)
	tx := repo.db.Create(&clientGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

// GetAll implements user.Repository
func (repo *clientRepository) GetAll(query string) (data []client.Core, err error) {
	var clients []Client
	// var users []User
	if query != "" {
		tx := repo.db.Where("`clients`.`users`.`name` LIKE ?", query).Find(&clients)
		// tx := repo.db.Raw("SELECT * FROM `clients` JOIN `users` ON `clients`.`user_id` = `users`.`id` WHERE name LIKE ?", query).Scan(&clients)
		if tx.Error != nil {
			return nil, tx.Error
		}
		var dataCore = toCoreList(clients)
		return dataCore, nil
	} else {
		tx := repo.db.Preload("User").Find(&clients)
		if tx.Error != nil {
			return nil, tx.Error
		}
		var dataCore = toCoreList(clients)
		return dataCore, nil
	}
}

// GetById implements user.RepositoryInterface
func (repo *clientRepository) GetById(id uint) (data client.Core, err error) {
	var client Client

	tx := repo.db.Preload("User").First(&client, id)

	if tx.Error != nil {
		return data, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = client.toCore()
	return dataCore, nil
}

// Update implements user.Repository
func (repo *clientRepository) Update(input client.Core, clientID uint, userID uint) error {
	clientGorm := fromCore(input)
	var client Client
	var user User

	tx := repo.db.Model(&user).Where("ID = ?", userID).Updates(&clientGorm.User)
	yx := repo.db.Model(&client).Where("ID = ?", clientID).Updates(&clientGorm)
	if tx.Error != nil && yx.Error != nil {
		return errors.New("failed update client")
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

// Delete implements user.Repository
func (repo *clientRepository) Delete(clientID uint, userID uint) error {
	var client Client
	var user User
	tx := repo.db.Delete(&client, clientID)
	yx := repo.db.Delete(&user, userID)
	if tx.Error != nil && yx.Error != nil {
		return errors.New("failed update client")
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

func (repo *clientRepository) FindUser(email string) (result client.Core, err error) {
	var clientData Client
	tx := repo.db.Where("email = ?", email).First(&clientData.User)
	if tx.Error != nil {
		return client.Core{}, tx.Error
	}

	result = clientData.toCore()

	return result, nil
}

func (repo *clientRepository) GetOrderById(clientId uint) (data []client.Order, err error) {
	var clientorder []Order

	tx := repo.db.Where("client_id = ?", clientId).Find(&clientorder)

	if tx.Error != nil {
		helper.LogDebug("client-query-Getorder | Error execute query. Error :", tx.Error)
		return data, tx.Error
	}

	helper.LogDebug("client-query-Getorder  | Row Affected : ", tx.RowsAffected)
	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = toCoreListOrder(clientorder)
	return dataCore, nil
}

func (repo *clientRepository) UpdateCompleteOrder(input client.Order, orderId, clientId uint) error {
	orderGorm := fromOrder(input)
	var ModelDataOrder Order

	// check status yang ada
	tx := repo.db.First(&ModelDataOrder, orderId)
	if tx.Error != nil {
		helper.LogDebug("Partner-query-UpdateOrderStatusComplete | Error execute query check order. Error :", tx.Error)
		return tx.Error
	}

	helper.LogDebug("Partner-query-UpdateOrderStatusComplete | Row Affected query check order: ", tx.RowsAffected)
	if tx.RowsAffected == 0 {
		return tx.Error
	}

	if ModelDataOrder.OrderStatus == cfg.ORDER_STATUS_ORDER_CONFIRMED {
		// proses update
		ty := repo.db.Model(&ModelDataOrder).Where("ID = ?", orderId).Where("client_id = ?", clientId).Updates(&orderGorm)
		if ty.Error != nil {
			helper.LogDebug("Partner-query-UpdateOrderStatusComplete | Error execute query update status. Error :", ty.Error)
			return errors.New("failed update client")
		}
		helper.LogDebug("Partner-query-UpdateOrderStatusComplete | Row Affected update status: ", ty.RowsAffected)
		if ty.RowsAffected == 0 {
			return errors.New("update failed")
		}
	} else {
		helper.LogDebug("Partner-query-UpdateOrderStatusComplete | modelData.OrderStatus : ", ModelDataOrder.OrderStatus)
		return errors.New("Order data no need to complete.")
	}

	//send email
	var client Client
	yx := repo.db.Preload("User").First(&client, clientId)
	if yx.Error != nil {
		helper.LogDebug("Partner-query-UpdateOrderStatusComplete | Failed get client data. Error ", yx.Error)
	}

	if ModelDataOrder.OrderStatus == cfg.ORDER_STATUS_COMPLETE_ORDER {
		if client.User.Email != "" {
			clientEmail := client.User.Email
			vaString := thirdparty.GetVABankTitle(ModelDataOrder.PaymentMethod)
			var dataBody = map[string]interface{}{
				"name":           client.User.Name,
				"event_name":     ModelDataOrder.EventName,
				"payment_method": vaString,
				"gross_ammount":  helper.FormatCurrencyIDR(ModelDataOrder.GrossAmmount),
			}

			thirdparty.SendEmailCompleteOrder2(clientEmail, fmt.Sprintf("Order Confirmed for Event : %s", ModelDataOrder.EventName), dataBody)
		} else {
			helper.LogDebug("Partner-query-UpdateOrderStatusComplete | Failed Sent Email. Client email not found.")
		}
	}

	return nil
}
