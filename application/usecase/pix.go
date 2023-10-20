package usecase

import model "github.com/tufee/codepix/domain/model/Bank"

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (pixUseCase *PixUseCase) RegisterKey(kind string, accountID string, key string) (*model.PixKey, error) {
	account, err := pixUseCase.PixKeyRepository.FindAccount(accountID)

	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)

	if err != nil {
		return nil, err
	}

	pixUseCase.PixKeyRepository.RegisterKey(pixKey)

	if pixKey.ID == "" {
		return nil, err
	}

	return pixKey, nil
}

func (pixUseCase *PixUseCase) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	pixKey, err := pixUseCase.PixKeyRepository.FindKeyByKind(key, kind)

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
