package database

import (
	"io"
	"os"
	"testing"
)

func CreateTo(writer io.WriteCloser) Option {
	return func(database *Database) error {
		database.create = func(string) (io.WriteCloser, error) {
			return writer, nil
		}
		return nil
	}
}

func OpenTo(reader io.ReadCloser) Option {
	return func(database *Database) error {
		database.open = func(string) (io.ReadCloser, error) {
			return reader, nil
		}
		return nil
	}
}

func OpenFileTo(reader io.ReadWriteCloser) Option {
	return func(database *Database) error {
		database.openFile = func(string, int, os.FileMode) (io.ReadWriteCloser, error) {
			return reader, nil
		}
		return nil
	}
}

func TestDatabase(t *testing.T) {
}
