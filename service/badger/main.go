package badger

import (
	"github.com/asuncm/vm/service/config"
	"github.com/dgraph-io/badger/v4"
	"github.com/samber/lo"
	"github.com/vmihailenco/msgpack"
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
	return ctx, err
}

// 查询数据
func Query(id []byte, filename string) ([]byte, error) {
	db, err := database(filename)
	var body []byte
	err = db.View(func(txn *badger.Txn) error {
		var item *badger.Item
		item, err = txn.Get(id)
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			body = val
			return nil
		})
		return err
	})
	defer db.Close()
	return body, err
}

// 增加数据
func Add(id []byte, payload []byte, filename string) ([]byte, error) {
	db, err := database(filename)
	if err != nil {
		panic(err)
	}
	err = db.Update(func(txn *badger.Txn) error {
		err = txn.Set(id, payload)
		return err
	})
	defer db.Close()
	return nil, err
}

// 更新map数据
func Update(id []byte, payload map[string]interface{}, filename string) ([]byte, error) {
	db, err := database(filename)
	var body []byte
	err = db.View(func(txn *badger.Txn) error {
		var item *badger.Item
		item, err = txn.Get(id)
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			err = db.Update(func(txn *badger.Txn) error {
				var mType map[string]interface{}
				err = msgpack.Unmarshal(val, &mType)
				if err == nil {
					newData := lo.Assign(mType, payload)
					ctx, _ := msgpack.Marshal(newData)
					body = []byte(ctx)
					err = txn.Set(id, body)
				}
				return err
			})
			return err
		})
		return err
	})
	defer db.Close()
	return body, err
}

// 更新string数据

func UpdateToString(id []byte, payload string, filename string) ([]byte, error) {
	db, err := database(filename)
	var body []byte
	err = db.Update(func(txn *badger.Txn) error {
		err = txn.Set(id, []byte(payload))
		return err
	})
	defer db.Close()
	return body, err
}

// 删除数据
func Remove(id []byte, filename string) ([]byte, error) {
	db, err := database(filename)
	err = db.Update(func(txn *badger.Txn) error {
		err = txn.Delete(id)
		return err
	})
	defer db.Close()
	return nil, err
}
