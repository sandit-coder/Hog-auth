package strategies

import "Hog-auth/internal/app/domain/vo"

type PhoneNumberNormalizer struct {
}

func NewPhoneNumberNormalizer() *EmailNormalizer {
	return &EmailNormalizer{}
}

func (p *PhoneNumberNormalizer) NormalizeCredential(credential string) (string, error) {
	phone, err := vo.NewPhoneNumber(credential)
	if err != nil {
		return "", err
	}

	return phone.Value, nil
}
