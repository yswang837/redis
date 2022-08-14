package iRedis

type IRedis interface {
	//KEYS(patten string) ([]string, error)
	EXISTS(key string) (bool, error)
	GET(key string) (string, error)
	//HGET(key, field string) (string, error)
	//HGETALL(key string) (map[string]string, error)
	//SCAN(patten string) ([]string, error)
	SET(key, value string) (int64, error)
	EXPIRE(key, value, ttl string) (int64, error)

	//HSET(key, field, value string) (int64, error)

}
