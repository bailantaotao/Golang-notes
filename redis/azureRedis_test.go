package redis

import (
	"encoding/json"
	"fmt"
	"gopkg.in/redis.v3"
	"testing"
)

type testRedisModel struct {
	UserName string
	Tel      string
	Address  string
}

func TestNewClient(t *testing.T) {
	//for i := 0; i < 10; i++ {
	go azureCachePressTest()
	//}
	c := make(chan struct{})
	<-c
	//for i := 0; i < 100; i++ {
	//	err = client.Set(fmt.Sprintf("key%d", i), "value", 0).Err()
	//	if err != nil {
	//		panic(err)
	//	}
	//	_, err := client.Get(fmt.Sprintf("key%d", i)).Result()
	//	if err != nil {
	//		panic(err)
	//	}
	//	err = client.Del(fmt.Sprintf("key%d", i)).Err()
	//	if err != nil {
	//		panic(err)
	//	}
	//}
}

func azureCachePressTest() {
	client := redis.NewClient(&redis.Options{
		//Addr:     "edwinderedisserver.redis.cache.windows.net:6379",
		//Password: "OyPTtUq7/na/o5cx6XPfcIoDXgodOZgDb7rcimiwvZs=", // no password set
		DB:       0, // use default DB
		Addr:     "edwinderedisserver2.redis.cache.windows.net:6379",
		Password: "KSMMXHvyf0QmNMKAZqD2/7aEdoHO6IErB7LiWZiDHGU=", // no password set
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	redisModel := &testRedisModel{
		UserName: "usaear",
		Tel:      "22212",
		Address:  "ffffff",
	}

	marshalModel, _ := json.Marshal(redisModel)

	err = client.Set("key", marshalModel, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Bytes()
	if err != nil {
		panic(err)
	}

	unmarshalRedis := &testRedisModel{}
	err = json.Unmarshal(val, unmarshalRedis)
	if err != nil {
		panic(err)
	}

	fmt.Println("unmarshal: ", unmarshalRedis)
}
