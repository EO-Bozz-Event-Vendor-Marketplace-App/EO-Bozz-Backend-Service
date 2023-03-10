package discussion

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	Comment   string
	PartnerID uint
	ClientID  uint
	ServiceID uint
	CreatedAt time.Time
}

type ServiceInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core, c echo.Context) error
	GetById(id uint) (data Core, err error)
	Update(input Core, id uint, c echo.Context) error
	Delete(id uint) error
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) error
	GetById(id uint) (data Core, err error)
	Update(input Core, id uint) error
	Delete(id uint) error
}
