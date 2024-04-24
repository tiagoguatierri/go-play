package repository

type Creatable[T any] interface {
	Create(entity T) error
}

type Getable[T any] interface {
	Get(entity T) (T, error)
}

type Updatable[T any] interface {
	Update(entity T) error
}

type Deletable[T any] interface {
	Delete(entity T) error
}

type Repository[T any] interface {
	Creatable[T]
	Getable[T]
	Updatable[T]
	Deletable[T]
}
