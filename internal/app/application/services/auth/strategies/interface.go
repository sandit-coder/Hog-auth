package strategies

type Strategy interface {
	NormalizeCredential(credential string) (string, error)
}
