package public

import (
	"fmt"
	"src/github.com/gomodule/redigo/redis"
)

func redisCon() (pool *redis.Pool) {
	return &redis.Pool{
		MaxIdle:   3, /*最大的空闲连接数*/
		MaxActive: 8, /*最大的激活连接数*/
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}

func Push() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("连接失败 err = ", err)
	}

	defer conn.Close()

	_, err = conn.Do("set", 100, "abc")
	if err != nil {
		fmt.Println("set key err = ", err)
	}

}

func RedisGet(userID int, userPwd string) (code int) {
	conn := redisCon().Get()
	defer conn.Close()
	pwd, err := redis.String(conn.Do("get", 100))
	if err != nil || pwd != userPwd {
		code = 500
	} else {
		code = 200
	}

	return code
}
func Del() {

}

func Modify() {

}
