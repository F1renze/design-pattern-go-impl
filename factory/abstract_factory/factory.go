package abstract_factory

import (
	"errors"
	"os"
	"reflect"
)

type IDataAccess interface {
	CreateUser() (IUser, error)
	CreateDepartment() (IDepartment, error)
}

type IUser interface {
	SetUserName(name string) error
	GetUserName() string
}

type IDepartment interface {
	AddMember(user IUser) error
}

const DBEnv = "DP_DB"

var (
	typeRegistry = make(map[string]reflect.Type)

	ErrDBTypeNotSupport = errors.New("DB type Not Support")
)

func registerType(typedNil interface{}) {
	t := reflect.TypeOf(typedNil).Elem()
	typeRegistry[t.Name()] = t
}

func init() {
	registerType((*MySQLUser)(nil))
	registerType((*MysqlDepartment)(nil))
	registerType((*MongoDBUser)(nil))
	registerType((*MongoDBDepartment)(nil))
}
func makeInstance(name string) interface{} {
	got, ok := typeRegistry[name]
	if !ok {
		return nil
	}
	return reflect.New(got).Interface()
}

func NewDataAccess() IDataAccess {
	return DataAccess{
		db: os.Getenv(DBEnv),
	}
}

type DataAccess struct {
	db string
}

func (a DataAccess) CreateUser() (IUser, error) {
	u, ok := makeInstance(a.db+"User").(IUser)
	if !ok {
		return nil, ErrDBTypeNotSupport
	}
	return u, nil
}

func (a DataAccess) CreateDepartment() (IDepartment, error) {
	d, ok := makeInstance(a.db+"Department").(IDepartment)
	if !ok {
		return nil, ErrDBTypeNotSupport
	}
	return d, nil
}

type MySQLUser struct {
	name string
}

func (u *MySQLUser) SetUserName(name string) error {
	u.name = name
	return nil
}

func (u *MySQLUser) GetUserName() string {
	return u.name
}

type MysqlDepartment struct {
}

func (d *MysqlDepartment) AddMember(u IUser) error {
	return nil
}

type MongoDBUser struct {
	name string
}

func (u *MongoDBUser) SetUserName(name string) error {
	u.name = name
	return nil
}

func (u *MongoDBUser) GetUserName() string {
	return u.name
}

type MongoDBDepartment struct {
}

func (d *MongoDBDepartment) AddMember(u IUser) error {
	return nil
}
