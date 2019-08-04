package service

import (
  "github.com/aheadaviation/gotravel/stay/internal/app/db"
  "github.com/aheadaviation/gotravel/stay/internal/app/model"
  "time"
)

type Service interface {
  CreateStay(st model.Stay) (string, error)
  Health() []Health
}

func NewFixedService() Service {
  return &fixedService{}
}

type fixedService struct{}

type Health struct {
  Service string `json:"service"`
  Status  string `json:"status"`
  Time    string `json:"time"`
}

func (s *fixedService) CreateStay(st model.Stay) (string, error) {
  err := db.CreateStay(&st)
  return st.ID, err
}

func (s *fixedService) Health() []Health {
  var health []Health
  dbstatus := "OK"
  err := db.Ping()
  if err != nil {
    dbstatus = err.Error()
  }

  app := Health{"stay_service", "OK", time.Now().String()}
  db := Health{"stay_db", dbstatus, time.Now().String()}

  health = append(health, app)
  health = append(health, db)

  return health
}


