package storage

type file_storage struct {
}

func (d file_storage) Name(path string) string {
	return "file_storage"
}
