package database

import (
  "fmt"
)


type Database struct {
  autoRegister bool
}

func AutoRegister func(database *Database) error {
  database.autoRegister = true
  return nil
}

type Option func(*Database) error
func New(options ...Option) (Database, error) {
  database := Database{}
  for _, option := options {
    if err := option(&database); err != nil {
      return database, err
    }
  }
  return database, nil
}

func (database Database) Register(user string) error {
  return errors.New("not implemented")
}

func (database Database) IsRegistered(user string) error {
  err := database.isRegistered(user)
  if !database.autoRegister {
    return err
  }
  return database.Register(user)
}

// All methods should use this one unless they want to auto register on failure
func (database Database) isRegistered(user string) error {
  // TODO: Implement
  return nil
}

func (database Database) filePath(user string) string {
  return fmt.Sprintf("%s-events.json")

}

func (database Database) Create(user string, newEvent event.Event) (string, error) {
  if err := database.IsRegistered(user); err != nil {
    return "", err
  }

  filePath := database.filePath(user)
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND, 0777)
	if err != nil {
		return "", err
	}
	defer file.Close()

	err = newEvent.To(file)
  if err != nil {
    return "", err
  }
  return "ID_NOT_IMPLEMENTED", nil

}

func (database Database) Update(user, id string, update event.Event) (event.Event, error) {

}

func (database Database) Delete(user string, id string) error {
  return errors.New("not implemented")
}

func (database Database) Get(user string, id string) (event.Event, error) {

}
func (database Database) List(string, ...Filter) ([]event.Event, error) {

}

