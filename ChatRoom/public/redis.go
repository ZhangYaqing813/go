package public

import (
	"fmt"
	"src/github.com/gomodule/redigo/redis"
)

// 创建redis 链接，并返回连接池地址

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
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("RedisPush close conneted err= ", err)
		}
	}()
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
	// 从redis 连接池中获取一个链接
	conn := redisCon().Get()
	//延迟关闭
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("redisget close conneted err= ", err)
		}
	}()
	// 判断用户ID和密码是否匹配
	// 1、先判断用户是否存在
	res, _ := redis.String(conn.Do("keys", userID))
	if len(res) == 0 { //如果redis 存在UID则长度不为零，
		// 获取userID中的密码
		pwd, err := redis.String(conn.Do("get", userID))
		// 判断获取密码字段是是否有错误，或者与输入的用户名不同
		if err != nil || pwd != userPwd {
			// 条件成立则code 值设置成500 ，表示验证未通过
			code = 500
		} else {
			//条件不成立则表示验证通过，返回200
			code = 200
		}
	}

	return code
}

/*
后期扩展
func Del() {

}

func Modify() {

}
*/
