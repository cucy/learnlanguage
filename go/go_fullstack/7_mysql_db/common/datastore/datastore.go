package datastore

import "7_mysql_db/models"
import "errors"

// 定义存储数据接口
type Datastore interface {
	CreateUser(user *models.User) error
	GetUser(username string) (*models.User, error)
	Close()
}

const (
	MYSQL = iota
	MONGODB
	REDIS
)

func NewDatastore(datastoreType int, dbConnectionString string) (Datastore, error) {

	switch datastoreType {
	case MYSQL:
		return NewMySQLDatastore(dbConnectionString)
	case MONGODB:
		//return NewMongoDBDatastore(dbConnectionString)
	case REDIS:
		//return NewRedisDatastore(dbConnectionString)
	default:
		return nil, errors.New("The datastore you specified does not exist!")
	}
	return nil, errors.New("The datastore you specified does not exist!")

}
