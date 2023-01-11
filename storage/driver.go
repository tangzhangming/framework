package storage


type driverface interface {

	Name(path string) string

}