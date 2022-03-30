package service

import (
	"pos-is-backend/internal/domain/entity"
	"pos-is-backend/internal/domain/repository"
	"time"
)

type CashierService struct {
	cashierRepo repository.ICashierRepository
}

type ICashierService interface {
	CreateCashier(*entity.CashierRequest) (*entity.CashierResponse, error)
	GetAllCashier(limit, skip int) (*entity.ListCashierResponse, error)
	GetDetailCashier(cashierId int) (*entity.CashierView, error)
	UpdateCashier(cashierId int, name string) error
	DeleteCashier(cashierId int, name string) error
}

func NewCashierService(cashierRepo repository.ICashierRepository) *CashierService {
	return &CashierService{
		cashierRepo: cashierRepo,
	}
}

func (cs *CashierService) CreateCashier(request *entity.CashierRequest) (*entity.CashierResponse, error) {
	req := entity.Cashier{
		Name:     request.Name,
		Passcode: "123456",
	}

	res, err := cs.cashierRepo.CreateCashier(&req)
	if err != nil {
		return nil, err
	}

	return &entity.CashierResponse{
		Id:        res.Id,
		Name:      res.Name,
		Passcode:  res.Passcode,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (cs *CashierService) GetAllCashier(limit, skip int) (*entity.ListCashierResponse, error) {
	var (
		res         entity.ListCashierResponse
		err         error
		total       int
		cashierView []entity.CashierView
	)

	res.Meta.Limit = limit
	res.Meta.Skip = skip
	total, err = cs.cashierRepo.TotalCashier()
	if err != nil {
		return &res, err
	}

	res.Meta.Total = total
	cashiers, err := cs.cashierRepo.GetAllCashier(limit, skip)
	if err != nil {
		res.Cashiers = make([]entity.CashierView, 0)
		return &res, nil
	}

	for _, v := range *cashiers {
		cashierView = append(cashierView, entity.CashierView{
			CashierId: v.Id,
			Name:      v.Name,
		})
	}

	res.Cashiers = cashierView

	return &res, nil
}

func (cs *CashierService) GetDetailCashier(cashierId int) (*entity.CashierView, error) {
	var (
		res entity.CashierView
		err error
	)

	cashier, err := cs.cashierRepo.GetDetailCashier(cashierId)
	if err != nil {
		return &res, err
	}

	res.CashierId = cashier.Id
	res.Name = cashier.Name

	return &res, nil
}

func (cs *CashierService) UpdateCashier(cashierId int, name string) error {
	err := cs.cashierRepo.UpdateCashier(&entity.Cashier{
		Id:        cashierId,
		Name:      name,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (cs *CashierService) DeleteCashier(cashierId int, name string) error {
	err := cs.cashierRepo.DeleteCashier(cashierId, name)
	if err != nil {
		return err
	}

	return nil
}
