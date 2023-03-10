package client

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID              uint
	Gender          string
	Address         string
	City            string
	Phone           string
	ClientImageFile string
	UserID          uint
	User            User
	Order           []Order
}

type User struct {
	ID       uint
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Role     string
}

type Order struct {
	ID                    uint
	EventName             string
	StartDate             time.Time
	EndDate               time.Time
	EventLocation         string
	EventAddress          string
	NotesForPartner       string
	ServiceName           string
	ServicePrice          uint
	GrossAmmount          uint
	PaymentMethod         string
	OrderStatus           string
	PayoutRecieptFile     string
	PayoutDate            time.Time
	MidtransTransactionID string
	MidtransVaNumber      string
	MidtransExpiredTime   string
	ServiceID             uint
	ClientID              uint
}

type ServiceInterface interface {
	GetAll(query string) (data []Core, err error)
	GetById(id uint) (data Core, err error)
	Create(input Core, c echo.Context) error
	Update(input Core, clientID uint, userID uint, c echo.Context) error
	Delete(clientID uint, userID uint) error
	GetOrderById(id uint) (data []Order, err error)
	UpdateCompleteOrder(input Order, orderId, clientId uint) error
}

type RepositoryInterface interface {
	GetAll(query string) (data []Core, err error)
	GetById(id uint) (data Core, err error)
	Create(input Core) error
	Update(input Core, clientID uint, userID uint) error
	Delete(clientID uint, userID uint) error
	FindUser(email string) (data Core, err error)
	GetOrderById(id uint) (data []Order, err error)
	UpdateCompleteOrder(input Order, orderId, clientId uint) error
}
