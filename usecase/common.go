package usecase

// UseCase Port for Driven Adapters.
type UseCase interface {
	CreateAccount(request []byte) (interface{}, error)
}
