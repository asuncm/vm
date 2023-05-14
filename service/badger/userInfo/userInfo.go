package userInfo

import (
	"github.com/asuncm/vm/service/config"
	"github.com/dgraph-io/badger/v4"
	"log"
	"strings"
)

// 创建数据库实例
func Badger(options config.ComConf, config Authorization) (UserInfo, error) {
	// 初始化实例数据
	userInfo := UserInfo{
		Status: true,
		Origin: "*",
	}
	pathname := strings.Join([]string{options.DataDir, "badgerDB"}, "/")
	opts := badger.DefaultOptions(pathname)
	// 设置badger缓存
	opts.IndexCacheSize = 100 << 20
	db, err := badger.Open(opts)
	if err != nil {
		userInfo.Status = false
	}
	err = db.View(func(txn *badger.Txn) error {
		users := query(txn, config)
		log.Print(users)
		return err
	})
	defer db.Close()
	return userInfo, err
}

// 查询用户数据
func query(txn *badger.Txn, opts Authorization) interface{} {
	// 获取用户数据
	user, err := txn.Get([]byte(opts.Verify))
	// 声明value参数实例
	var value, valueCopy []byte
	// 未获取到用户信息
	if err != nil {
		return err
	}

	err = user.Value(func(val []byte) error {
		value = val
		valueCopy = append([]byte{}, val...)
		return nil
	})
	log.Print(err, "99990000====", valueCopy)
	if err != nil {
		return err
	}
	return value
}

// 更新用户数据
func update() {}

// 删除冗余数据
func delData() {

}
