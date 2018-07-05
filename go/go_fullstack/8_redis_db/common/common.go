package common

import "8_redis_db/common/datastore"

type Env struct {
	DB datastore.Datastore
}
