package model

import "time"

type Stay struct {
  ID            string
  Hotel         string
  StartDate     time.Time
  EndDate       time.Time
  currentState  string
  previousState string
}

func New() Stay {
  return Stay{}
}
