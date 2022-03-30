package service

import (
	"pos-is-backend/internal/domain/entity"
	"pos-is-backend/internal/domain/repository"
	"pos-is-backend/pkg/auth"
	"strconv"
)

type VerifyService struct {
	verifyRepo  repository.IVerifyRepository
	cashierRepo repository.ICashierRepository
}

type IVerifyService interface {
	LoginPasscode(cashierId int, passcode string) (*auth.TokenDetails, error)
	GetPasscode(cashierId int) (*entity.Verify, error)
	LogoutPasscode(cashierId int, passcode string) error
}

func NewVerifyService(verifyRepo repository.IVerifyRepository, cashierRepo repository.ICashierRepository) *VerifyService {
	return &VerifyService{
		verifyRepo:  verifyRepo,
		cashierRepo: cashierRepo,
	}
}

func (vs *VerifyService) LoginPasscode(cashierId int, passcode string) (*auth.TokenDetails, error) {

	_, err := vs.cashierRepo.GetCashierByIdAndPasscode(cashierId, passcode)
	if err != nil {
		return nil, err
	}

	tot, err := vs.verifyRepo.LoginPasscode(cashierId, passcode)
	if err != nil {
		return nil, err
	}

	if tot == 0 {
		return nil, err
	}

	token, err := auth.CreateToken(strconv.Itoa(cashierId))
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (vs *VerifyService) GetPasscode(cashierId int) (*entity.Verify, error) {
	_, err := vs.cashierRepo.GetDetailCashier(cashierId)
	if err != nil {
		return nil, err
	}

	res, err := vs.verifyRepo.GetPasscode(cashierId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (vs *VerifyService) LogoutPasscode(cashierId int, passcode string) error {

	_, err := vs.cashierRepo.GetCashierByIdAndPasscode(cashierId, passcode)
	if err != nil {
		return err
	}

	tot, err := vs.verifyRepo.LoginPasscode(cashierId, passcode)
	if err != nil {
		return err
	}

	if tot == 0 {
		return err
	}

	return nil
}
