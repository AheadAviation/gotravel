package db

import (
  "errors"
  "flag"
  "fmt"
  "os"
  "stay/internal/app/model"
)

type Database interface {
  Init() error
  CreateStay(stay *model.Stay) error
  Ping() error
}

var (
  database              string
  DefaultDb             Database
  DBTypes               = map[string]Database{}
  ErrNoDatabaseFound    = "No database with %v registered"
  ErrNoDatabaseSelected = errors.New("No DB Selected")
)

func init() {
  flag.StringVar(&database, "database", os.Getenv("STAY_DATABASE"), "Database for trip stays")
}

func Init() error {
  if database == "" {
    return ErrNoDatabaseSelected
  }
  err := Set()
  if err != nil {
    return err
  }
  return DefaultDb.Init()
}

func Set() error {
  if v, ok := DBTypes[database]; ok {
    DefaultDb = v
    return nil
  }
  return fmt.Errorf(ErrNoDatabaseFound, database)
}

func Register(name string, db Database) {
  DBTypes[name] = db
}

func CreateStay(s *model.Stay) error {
  return DefaultDb.CreateStay(s)
}

func Ping() error {
  return DefaultDb.Ping()
}
