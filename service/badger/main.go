package badger

import (
	"github.com/asuncm/vm/service/config"
	"github.com/dgraph-io/badger/v4"
	"regexp"
	"runtime"
	"strings"
)

// 建立badger数据库连接
func database(filename string, fn func(body interface{})) {
	// 初始化实例数据
	options, err := config.Config(filename)
	if err != nil {
		fn(err)
	} else {
		pathname := strings.Join([]string{options.DataDir, "badgerDB"}, "/")
		platform := runtime.GOOS
		reg, _ := regexp.MatchString("^window", platform)
		if reg {
			pathname = strings.Replace(pathname, "/", "\\", -1)
		}
		opts := badger.DefaultOptions(pathname)
		// 设置badger缓存
		opts.IndexCacheSize = 100 << 20
		db, err := badger.Open(opts)
		if err != nil {
			fn(err)
		}
		err = db.View(func(txn *badger.Txn) error {
			fn(txn)
			return nil
		})
		if err != nil {
			fn(err)
		}
		defer db.Close()
	}
}

// 查询数据
func Query(id string, filename string) interface{} {
	var body interface{}
	data := func(txn interface{}) {
		user, err := txn.Get([]byte(id))
	}
	database(filename, data)
	return body
}

// 增加数据
func Add() {

}

// 更新数据
func Update() {

}

// 删除数据
func Remove() {

}
