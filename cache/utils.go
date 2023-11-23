package cache

// GetKey 拼接Redis Key
func GetKey(keys ...string) string {
	key := ""
	for i := 0; i < len(keys); i++ {
		key += keys[i]
		if i != len(keys)-1 {
			key += ":"
		}
	}
	return key
}
