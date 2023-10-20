package usecase

import (
	"errors"
	"log"

	model "github.com/tufee/codepix/domain/model/Bank"
)

type TransactionUseCase struct {
	TransactionRepository model.TransactionRepositoryInterface
	PixKeyRepository      model.PixKeyRepositoryInterface
}

func (transactionUseCase *TransactionUseCase) Register(accountID string, amount float64, pixKeyto string, pixKeyKindTo string, description string) (*model.Transaction, error) {
	account, err := transactionUseCase.PixKeyRepository.FindAccount(accountID)

	if err != nil {
		return nil, err
	}

	pixKey, err := transactionUseCase.PixKeyRepository.FindKeyByKind(pixKeyto, pixKeyKindTo)

	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, amount, pixKey, description)

	if err != nil {
		return nil, err
	}

	transactionUseCase.TransactionRepository.Save(transaction)

	if transaction.ID != "" {
		return transaction, nil
	}

	return nil, errors.New("unable to process this transaction")
}

func (transactionUseCase *TransactionUseCase) Confirm(transactionID string) (*model.Transaction, error) {
	transaction, err := transactionUseCase.TransactionRepository.Find(transactionID)

	if err != nil {
		log.Println("Transaction not found", transactionID)
		return nil, err
	}

	transaction.Status = model.TransactionConfirmed

	err = transactionUseCase.TransactionRepository.Save(transaction)

	if err != nil {
		return transaction, nil
	}

	return nil, errors.New("unable to confirm this transaction")
}

func (transactionUseCase *TransactionUseCase) Complete(transactionID string) (*model.Transaction, error) {
	transaction, err := transactionUseCase.TransactionRepository.Find(transactionID)

	if err != nil {
		log.Println("Transaction not found", transactionID)
		return nil, err
	}

	transaction.Status = model.TransactionCompleted

	err = transactionUseCase.TransactionRepository.Save(transaction)

	if err != nil {
		return transaction, nil
	}

	return nil, errors.New("unable to complete this transaction")
}

func (transactionUseCase *TransactionUseCase) Error(transactionID string, reason string) (*model.Transaction, error) {
	transaction, err := transactionUseCase.TransactionRepository.Find(transactionID)

	if err != nil {
		log.Println("Transaction not found", transactionID)
		return nil, err
	}

	transaction.Status = model.TransactionError
	transaction.CancelDescription = reason

	err = transactionUseCase.TransactionRepository.Save(transaction)

	if err != nil {
		return transaction, nil
	}

	return nil, errors.New("transaction error")
}
