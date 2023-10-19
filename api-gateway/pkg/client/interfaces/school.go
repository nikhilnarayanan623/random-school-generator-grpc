package interfaces

type SchoolClient interface {
	GetOneInJSON(name string) ([]byte, error)
}
