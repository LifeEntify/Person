package person_ctl

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	person_common "github.com/lifeentify/person/common"
	person_config "github.com/lifeentify/person/config"
	person "github.com/lifeentify/person/person/v1"
	person_repo "github.com/lifeentify/person/repository"
	person_db "github.com/lifeentify/person/repository/db"
	"golang.org/x/exp/slices"
)

type Controller struct {
	DB     person_repo.Repository
	Config *person_config.Config
}

const (
	Mongo    = "MONGODB"
	MySQL    = "MYSQL"
	PostGres = "POSTGRES"
)

func NewController(config *person_config.Config) (*Controller, error) {
	dbType := config.DatabaseType

	if dbType == Mongo {
		uri := config.MongoUrl
		if uri == "" {
			log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
		}
		db := person_db.NewMongoDB(config)
		return &Controller{
			db,
			config,
		}, nil
	}
	return nil, fmt.Errorf("no match fund for db type :%s", dbType)
}

func (c *Controller) Login(phoneNumber string, password string, category string) (*person.Person, error) {
	dataByte, err := c.DB.FindPersonByPhone(context.TODO(), phoneNumber)
	if err != nil {
		return nil, err
	}
	var p person.Person
	err = json.Unmarshal(dataByte, &p)
	idx := slices.IndexFunc(p.Credential, func(c *person.Credential) bool { return c.Category == category })
	if err != nil {
		return nil, err
	}
	if idx == -1 {
		return nil, fmt.Errorf("not permitted")
	}
	if !person_common.CheckPasswordHash(password, p.Credential[idx].Password) {
		return nil, fmt.Errorf("wrong password")
	}
	return &p, err
}
func (c *Controller) CreateAccount(newPerson *person.Person) (any, error) {
	sPersonByte, err := c.DB.FindPersonByPhone(context.TODO(), newPerson.Profile.PhoneNumber)
	if err != nil {
		return nil, err
	}
	var p person.Person
	err = json.Unmarshal(sPersonByte, &p)
	if err != nil {
		return nil, err
	}
	if p.AccountId == "" {
		return nil, fmt.Errorf("person with phone %s already exists", newPerson.Profile.PhoneNumber)
	}
	passHash, err := person_common.HashPassword(newPerson.Credential[0].Password)
	if err != nil {
		return nil, err
	}
	newPerson.Credential[0].Password = passHash
	personJson, err := newPerson.ToJson()
	if err != nil {
		return nil, err
	}
	a, err := c.DB.Save(context.TODO(), personJson)
	if err != nil {
		return nil, err
	}
	return a.InsertedID, nil
}
