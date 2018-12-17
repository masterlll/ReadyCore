package RedigoEFcore


func MergeValue(com interface{}, table interface{}, key interface{}, val ...interface{}) ([]interface{}, bool) {

	var urls []interface{}

	// table , key ,key
	if com == "HMGET" {
		//fmt.Println(" table , key ,key")
		urls = append(urls, table.(string))
		for _, i := range val {
			urls = append(urls, i)
		}
		return urls, true
	}
	//  table ,key ,value , key ,value
	if com == "HMSET" {
		//	fmt.Println(" table ,key ,value , key ,value")
		for b, i := range key.([]string) {
			urls = append(urls, i)
			urls = append(urls, val[b])
		}
		return urls, true
	}
	//table (key)  VALUE1  VALUE1
	if com == "SADD" || com == "SREM" || com == "LPUSH" || com == "RPUSH"{
		//fmt.Println(" key  VALUE1  VALUE1 ")
		urls = append(urls, table)
		for _, i := range val {
			urls = append(urls, i)
		}

		return urls, true

	}
	if com == "HSET" || com == "LSET" {
		//	fmt.Println("table key value")
		urls = append(urls, table)

		//a := key.([]string)
		a := key.(string)
		urls = append(urls, a)
		urls = append(urls, val[0])
		// for i := 0; i <= len(a)-1; i++ {
		// 	urls = append(urls, a[i])
		// 	urls = append(urls, a[i])
		// }
		//	fmt.Println(len(urls))

		return urls, true
	}
	//  table key(int)
	if com == "INCRBY" || com == "GETSET" || com == "DECRBY" {
		//fmt.Println("table key")
		urls = append(urls, table)
		urls = append(urls, key.(int))

		return urls, true
	}
	//  table key
	if com == "HGET" || com == "HEXISTS" || com == "HDEL" || com == "LINDEX" || com == "SISMEMBER" || com == "SET" {
		//fmt.Println("table key")
		urls = append(urls, table)
		urls = append(urls, key.(string))

		return urls, true
	}
	// tables
	if com == "HGETALL" || com == "SMEMBERS" || com == "DEL" || com == "LLEN" || com == "RPOP" || com == "LPOP" || com == "GET" || com == "EXISTS" || com == "HLEN" || com == "HKEYS" || com == "SPOP" || com == "SCARD" || com == "HVALS"{
		//	fmt.Println("table ")
		urls = append(urls, table)

		return urls, true

	}
	if com == "SCAN" {

		urls = append(urls, "0")
		urls = append(urls, "MATCH")
		urls = append(urls, "*"+table.(string)+"*")
		urls = append(urls, "COUNT")
		urls = append(urls, "10000")
		return urls, true
	}

	if com == "EXPIRE" {

		urls = append(urls, table)
		urls = append(urls, val[0].(int))
		return urls, true
	}

	return nil, false

}
