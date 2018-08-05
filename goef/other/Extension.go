package RedigoEFcore

type Container struct {
	Action string
	DB     int
	Input  []interface{}
	Twice  []interface{}
}

func MergeValue(com interface{}, table interface{}, key interface{}, val ...interface{}) ([]interface{}, bool) {

	var urls []interface{}

	// table , key ,key
	if com == "HMGET" {
		//fmt.Println(" table , key ,key")
		urls = append(urls, table.(string))
		for _, i := range key.([]string) {
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
	if com == "SADD" || com == "SREM" {
		//fmt.Println(" key  VALUE1  VALUE1 ")
		urls = append(urls, table)
		for _, i := range val {
			urls = append(urls, i)
		}
		return urls, true

	}
	if com == "HSET" {
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
	//  table key
	if com == "HGET" || com == "HEXISTS" || com == "HDEL" {
		//fmt.Println("table key")
		urls = append(urls, table)
		urls = append(urls, key.(string))

		return urls, true
	}
	// table
	if com == "HGETALL" || com == "SMEMBERS" || com == "DEL" {
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

	return nil, false

}
