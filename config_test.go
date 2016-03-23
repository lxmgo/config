package config

import (
	"testing"
)

func TestGet(t *testing.T){
	config, err := NewConfig("testdata/testini.ini")
	if err != nil {
		t.Errorf("Configuration file loading failed, err:%v",err.Error())
		t.Fatalf("err: %v",err)
	}

	// default::key test
	if v, err := config.Bool("debug"); err != nil || v != true {
		t.Errorf("Get failure: expected different value for debug (expected: [%#v] got: [%#v])",true,v)
		t.Fatalf("err: %v",err)
	}
	if v := config.String("url"); v != "act.wiki"{
		t.Errorf("Get failure: expected different value for url (expected: [%#v] got: [%#v])","act.wiki",v)
	}

	// reids::key test
	if v := config.Strings("redis::redis.key"); len(v) != 2 || v[0] != "push1" || v[1] != "push2"{
		t.Errorf("Get failure: expected different value for redis::redis.key (expected: [%#v] got: [%#v])","[]string{push1,push2}",v)
	}
	if v := config.String("mysql::mysql.dev.host"); v != "127.0.0.1" {
		t.Errorf("Get failure: expected different value for mysql::mysql.dev.host (expected: [%#v] got: [%#v])","127.0.0.1",v)
	}
	if v := config.String("mysql::mysql.master.host"); v != "10.0.0.1" {
		t.Errorf("Get failure: expected different value for mysql::mysql.master.host (expected: [%#v] got: [%#v])","10.0.0.1",v)
	}

	// math::key test
	if v, err := config.Int64("math::math.i64"); err != nil || v != 64 {
		t.Errorf("Get failure: expected different value for math::math.i64 (expected: [%#v] got: [%#v])",64,v)
		t.Fatalf("err: %v",err)
	}
	if v, err := config.Float64("math::math.f64"); err != nil || v != 64.1 {
		t.Errorf("Get failure: expected different value for math::math.f64 (expected: [%#v] got: [%#v])",64.1,v)
		t.Fatalf("err: %v",err)
	}

	// other::key test
	if v := config.String("other::name"); v != "ATC自动化测试^-^&($#……#" {
		t.Errorf("Get failure: expected different value for other::name (expected: [%#v] got: [%#v])","ATC自动化测试^-^&($#……#",v)
		t.Fatalf("err: %v",err)
	}
	if v := config.String("other::key1"); v != "test key" {
		t.Errorf("Get failure: expected different value for other::key1 (expected: [%#v] got: [%#v])","test key",v)
		t.Fatalf("err: %v",err)
	}
}
