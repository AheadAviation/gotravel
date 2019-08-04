package mongodb

import (
  "errors"
  "flag"
  "log"
  "os"
  "stay/internal/app/model"

  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

var (
  username        string
  password        string
  host            string
  db              = "stays"
  ErrInvalidHexID = errors.New("Invalid Id Hex")
)

func init () {
  flag.StringVar(&username, "mongo-user", os.Getenv("MONGO_USER"), "Mongo Username")
  flag.StringVar(&password, "mongo-password", os.Getenv("MONGO_PASSWORD"), "Mongo Possword")
  flag.StringVar(&host, "mongo-host", os.Getenv("MONGO_HOST"), "Mongo Host")
}

type Mongo struct {
  Session *mgo.Session
}

func (m *Mongo) Init() error {
  log.Printf("Connecting to DB: %s", host)
  var err error
  m.Session, err = mgo.Dial(host)
  if err != nil {
    return err
  }
  return nil
}

type MongoStay struct {
  model.Stay `bson:",inline"`
  ID bson.ObjectId `bson:"_id"`
}

func New() MongoStay {
  s := model.New()
  return MongoStay{
    Stay: s,
  }
}

func (m *Mongo) CreateStay(st *model.Stay) error {
  s := m.Session.Copy()
  defer s.Close()
  id := bson.NewObjectId()
  ms := New()
  ms.Stay = *st
  ms.ID = id
  c := s.DB("").C("stays")
  _, err := c.UpsertId(ms.ID, ms)
  if err != nil {
    return err
  }
  return nil
}

func (m *Mongo) Ping() error {
  s := m.Session.Copy()
  defer s.Close()
  return s.Ping()
}

