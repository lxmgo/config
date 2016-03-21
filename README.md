# config

Package config is a Configuration file parser for INI format

包 config 是一个简洁方的,支持注释的Go语言配置文件解析器,类似于Windows下的INI文件.

配置文件形式为`[section]` 的段构成, 内部使用 `name=value`键值对
如果为指定段节点,则默认放入名为`[default]`的段当中.
"#"为注释的开头,可以放置于任意的单独一行中.


## 安装

        go get github.com/lxmgo/config
        
## 示例

请查看 [conf.ini](testdata/testini.ini) 文件作为使用示例
        
## 使用规范

示例配置文件:
        
        [DEFAULT]
        host = act.wiki
        port = 8080
        f64 = 64.1
        
        [mysql]
        host = 127.0.0.1
        [mongodb]
        host = 127.0.0.2
        [redis]
        host = 127.0.0.3
        push_key = key1,key2,key3,...
        [memcache]
        host = 127.0.0.4
    
加载配置文件:
        
        config, err := NewConfig("testdata/testini.ini")
        c.Int("port")
        // result is int 8080
        
        c.Int64("port")
        // result is int64 8080
        
        c.Float64("f64")
        // result is float64 64.1
                
        c.String("host")
        // result is string "act.wiki"
    
        c.String("mysql.host")
        // result is string "127.0.0.1"
        
        c.String("redis.host")
        // result is string "127.0.0.3"
        
        c.Strings("redis.key")
        // result is []string{"key1","key2","key3",...}

## config APIS:
    
        String(key string) string
        Strings(key string) []string
        Bool(key string) (bool, error)
        Int(key string) (int, error)
        Int64(key string) (int64, error)
        Float64(key string) (float64,error)
        Set(key string, value string) error
        
## 更多信息

- 所有字符解析均使用小写的!