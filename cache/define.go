package cache

type (
	RedisConnOpt struct {
		Enable   bool
		Host     string
		Port     string
		Password string
		Index    int32
		TTL      int32
	}
)
