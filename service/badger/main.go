package badger

import (
	"github.com/asuncm/vm/service/config"
	"github.com/dgraph-io/badger/v4"
	"regexp"
	"runtime"
	"strings"
)

// 建立badger数据库连接
func database(filename string) (*badger.DB, error) {
	var ctx *badger.DB
	// 初始化实例数据
	options, err := config.Config(filename)
	if err != nil {
		return ctx, err
	}
	pathname := strings.Join([]string{options.DataDir, "badgerDB"}, "/")
	platform := runtime.GOOS
	reg, _ := regexp.MatchString("^window", platform)
	if reg {
		pathname = strings.Replace(pathname, "/", "\\", -1)
	}
	opts := badger.DefaultOptions(pathname)
	// 设置badger缓存
	opts.IndexCacheSize = 100 << 20
	ctx, err = badger.Open(opts)
	defer ctx.Close()
	return ctx, err
}

// 查询数据
func Query(id string, filename string) (interface{}, error) {
	db, err := database(filename)
	var body interface{}
	err = db.View(func(txn *badger.Txn) error {
		var item *badger.Item
		item, err = txn.Get([]byte(id))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			body = val
			return nil
		})
		return err
	})
	return body, err
}

// 增加数据
func Add(id string, payload interface{}, filename string) (interface{}, error) {
	db, err := database(filename)
	if err != nil {
		panic(err)
	}
	err = db.Update(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		
		err = txn.NewKeyIterator([]byte(id), opts)
	})
	return nil, err
}

// 更新数据
func Update(options interface{}, filename string) (interface{}, error) {
	txn, err := database(filename)
	item, err := txn.Get([]byte(options.id))
	if err != nil {
		panic(err)
	}
	var body interface{}
	err = item.Value(func(val []byte) error {
		body = val
		return nil
	})
	if err != nil {
		panic(err)
	}
	err = txn.Set([]byte(options.id), []byte(options), 0)
	return body, err
}

// 删除数据
func Remove(id string, filename string) {
	txn, err := database(filename)
	err = txn.Delete([]byte(id))
	if err != nil {
		panic(err)
	}
}
