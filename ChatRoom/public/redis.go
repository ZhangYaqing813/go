package public

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
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

func RedisPush(userID int, userPwd string) (code int) {
	conn := redisCon().Get()
	defer conn.Close()
	//判断下userID 是否已经存在，如果存在，则返回一个状态码或错误信息
	res, _ := redis.String(conn.Do("keys", userID))
	if len(res) == 0 { //如果redis 存在UID则长度不为零，
		_, err := conn.Do("set", userID, userPwd)
		if err != nil {
			fmt.Println("写入redis err= ", err)
			code = 500
		} else {
			code = 200
		}
	} else {
		fmt.Println("UID已被占用")
		code = 600
	}
	return code
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
