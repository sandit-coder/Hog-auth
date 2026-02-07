package strategies

import "Hog-auth/internal/app/domain/vo"

type EmailNormalizer struct{}

func NewEmailNormalizer() *EmailNormalizer {
	return &EmailNormalizer{}
}

func (e *EmailNormalizer) NormalizeCredential(credential string) (string, error) {
	email, err := vo.NewEmail(credential)
	if err != nil {
		return "", err
	}

	return email.Value, nil
}
