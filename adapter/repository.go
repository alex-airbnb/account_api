package adapter

// RepositoryPort Port for a Driven Adapter to manage Databases.
type RepositoryPort interface {
	Create(model interface{}) error
}
