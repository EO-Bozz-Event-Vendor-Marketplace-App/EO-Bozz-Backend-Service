package service

import (
	cfg "capstone-alta1/config"
	_service "capstone-alta1/features/service"
	"capstone-alta1/utils/helper"
	"capstone-alta1/utils/thirdparty"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type serviceService struct {
	serviceRepository _service.RepositoryInterface
	validate          *validator.Validate
}

func New(repo _service.RepositoryInterface) _service.ServiceInterface {
	return &serviceService{
		serviceRepository: repo,
		validate:          validator.New(),
	}
}

func (service *serviceService) Create(input _service.Core, c echo.Context) (err error) {
	var errUpload error
	input.ServiceImageFile, errUpload = thirdparty.Upload(c, cfg.SERVICE_IMAGE_FILE, cfg.SERVICE_FOLDER)
	if errUpload != nil {
		return errUpload
	}

	errCreate := service.serviceRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

func (service *serviceService) GetAll(queryName, queryCategory, queryCity, queryMinPrice, queryMaxPrice string) (data []_service.Core, err error) {
	data, err = service.serviceRepository.GetAll(queryName, queryCategory, queryCity, queryMinPrice, queryMaxPrice)
	if err != nil {
		helper.LogDebug(err)
		return nil, helper.ServiceErrorMsg(err)
	}
	return data, err
}

func (service *serviceService) GetById(id uint) (data _service.ServiceDetailJoinPartner, err error) {
	data, err = service.serviceRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		return _service.ServiceDetailJoinPartner{}, err
	}
	return data, err
}

func (service *serviceService) Update(input _service.Core, id uint, c echo.Context) error {
	var errUpload error
	input.ServiceImageFile, errUpload = thirdparty.UploadForUpdate(c, cfg.SERVICE_IMAGE_FILE, cfg.SERVICE_FOLDER)
	if errUpload != nil {
		return errUpload
	}
	err := service.serviceRepository.Update(input, id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

func (service *serviceService) Delete(id uint) error {
	err := service.serviceRepository.Delete(id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}
	return nil
}

func (service *serviceService) GetServiceAdditionalById(id uint) (data []_service.JoinServiceAdditional, err error) {
	data, err = service.serviceRepository.GetServiceAdditionalById(id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return data, err

}

func (service *serviceService) AddAdditionalToService(input []_service.ServiceAdditional) (err error) {

	errCreate := service.serviceRepository.AddAdditionalToService(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

func (service *serviceService) GetReviewById(id uint) (data []_service.Review, err error) {
	data, err = service.serviceRepository.GetReviewById(id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return data, err

}

func (service *serviceService) GetDiscussionById(id uint) (data []_service.Discussion, err error) {
	data, err = service.serviceRepository.GetDiscussionById(id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return data, err

}

func (service *serviceService) CheckAvailability(serviceId uint, queryStart, queryEnd string) (data _service.Order, err error) {

	startDate := helper.GetDateTimeFormatedToTime(queryStart + " 00:00:00")
	endtDate := helper.GetDateTimeFormatedToTime(queryEnd + " 00:00:00")

	helper.LogDebug("Input param Data startdate ", queryStart, " enddate ", queryEnd)
	helper.LogDebug("Data startdate ", startDate, " enddate ", endtDate)

	data, err = service.serviceRepository.CheckAvailability(serviceId, queryStart, queryEnd)
	if err != nil {
		return _service.Order{}, helper.ServiceErrorMsg(err)
	}

	return data, nil
}
