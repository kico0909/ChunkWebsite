package reids
//
import (
"github.com/garyburd/redigo/redis"
"encoding/json"
"time"
"log"

)

type DatabaseRedis struct {
	dbInfo map[string] interface{}
	connPool *redis.Pool
	connectionName string
}

/*
私有方法, 用于连接 redis 数据库
*/
func (_self *DatabaseRedis) connectionPool() bool {

	pwd, _ :=  _self.dbInfo["Password"].(string)
	host, _ :=  _self.dbInfo["Host"].(string)
	port, _ :=  _self.dbInfo["Port"].(string)
	dbname, _ :=  _self.dbInfo["Dbname"].(string)

	_self.connPool = &redis.Pool{
		// TODO 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     1000,
		MaxActive:   500,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host+`:`+port, redis.DialPassword(pwd))
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", dbname)
			return c, nil
		},
	}

	return true

}

/*
私有方法, 用于关闭数据库
*/
func (_self *DatabaseRedis) closeDB(){
	//_self.conn.Close()
}

/*
初始化,获得数据库配置信息
*/
func (_self *DatabaseRedis) Init(dbInfo map[string] interface{}){
	_self.dbInfo = dbInfo
	_self.connectionPool()
}

/*
链接数据库, 根据配置文件中的配置信息去对数据库进行连接
*/
func (_self *DatabaseRedis) Connection(connectionName string) {
	_self.connectionName = connectionName
	_self.connectionPool()
}

/*
REDIS 设置数据
*/
func (_self *DatabaseRedis) Set(key string, val interface{}) bool {
	_v ,_ := json.Marshal(val)
	_rc := _self.connPool.Get()

	_, err := _rc.Do("SET", key, string(_v))
	if err != nil {
		log.Print("redis error[", err , "]")
		return false
	}

	defer _rc.Close()
	return true
}

/*
读取数据
*/
func (_self *DatabaseRedis) Get(key string, toString bool) interface{} {
	var _res interface{}
	_rc := _self.connPool.Get()
	res, err := _rc.Do("GET", key)
	_r, _ := redis.Bytes(res,err)
	var mapResult map[string]interface{}
	if _err := json.Unmarshal(_r, &mapResult); _err!=nil {
		_r2,_ := redis.String(res,err)
		_res = _r2
	}else{
		_res = mapResult
	}
	defer _rc.Close()
	return _res
}

