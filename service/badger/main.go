package badger

import (
	"fmt"
	"github.com/asuncm/vm/service/config"
	"github.com/dgraph-io/badger/v4"
	"github.com/samber/lo"
	"github.com/vmihailenco/msgpack"
	"log"
	"regexp"
	"runtime"
	"strings"
	"time"
)

/*@package database		初始化数据库
* @param   badger		数据库
* @param   platform		平台类型
* @param   pathname		数据库路径
* @param   filename		项目路径关键词
 */
func database(filename string) (*badger.DB, error) {
	var ctx *badger.DB
	// 初始化实例数据
	options, err := config.Config(filename)
	if err != nil {
		return ctx, err
	}
	pathname := strings.Join([]string{options.DataDir, "badgerDB"}, "/")
	// 获取平台类型
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

/*@package Query		查询数据
* @param   id			查询数据uuid
* @param   filename		项目路径关键词
 */
func Query(id []byte, filename string) ([]byte, error) {
	db, err := database(filename)
	if err != nil {
		defer db.Close()
		return nil, fmt.Errorf("数据库异常")
	}
	var body []byte
	err = db.View(func(txn *badger.Txn) error {
		var item *badger.Item
		// 获取uuid的value值
		item, err = txn.Get(id)
		if err != nil {
			return fmt.Errorf("未查询到相关数据")
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

/*@package	Add			数据库新增数据
* @param   id			查询数据uuid
* @param   filename		项目路径关键词
* @param   payload		新增数据map结构
* @param   timestamp	过期时间TTL，单位s
 */
func Add(id []byte, payload []byte, filename string, timestamp interface{}) ([]byte, error) {
	db, err := database(filename)
	var (
		entry   *badger.Entry
		nowTime time.Duration
		newTime time.Time
	)
	if err != nil {
		defer db.Close()
		return nil, fmt.Errorf("数据库异常")
	}
	if timestamp != nil {
		timeStr := strings.Join([]string{timestamp.(string), "s"}, "")
		nowTime, err = time.ParseDuration(timeStr)
		if err != nil {
			defer db.Close()
			return nil, fmt.Errorf("时间戳转换异常")
		}
		date := time.Now()
		newTime = date.Add(nowTime)
		if err != nil {
			entry = badger.NewEntry(id, payload)
			entry.WithTTL(time.Duration(newTime.UnixNano()))
		}
	}
	err = db.Update(func(txn *badger.Txn) error {
		if timestamp != nil {
			err = txn.SetEntry(entry)
		} else {
			err = txn.Set(id, payload)
		}
		return err
	})

	defer db.Close()
	if err != nil {
		err = fmt.Errorf("数据存储发生异常")
	}
	return nil, err
}

/*@package	Update		数据库新增数据
* @param   id			查询数据uuid
* @param   filename		项目路径关键词
* @param   payload		新增数据map结构
* @param   timestamp	过期时间TTL，单位s
 */
func Update(id []byte, payload map[string]interface{}, filename string, timestamp interface{}) ([]byte, error) {
	db, err := database(filename)
	var (
		entry   *badger.Entry
		nowTime time.Duration
		newTime time.Time
		body    []byte
	)
	if err != nil {
		defer db.Close()
		return nil, fmt.Errorf("数据库异常")
	}
	log.Print(nowTime)
	if timestamp != nil {
		timeStr := strings.Join([]string{timestamp.(string), "s"}, "")
		nowTime, err = time.ParseDuration(timeStr)
		if err != nil {
			defer db.Close()
			return nil, fmt.Errorf("时间戳转换异常")
		}
		date := time.Now()
		newTime = date.Add(nowTime)
	}
	err = db.View(func(txn *badger.Txn) error {
		var item *badger.Item
		item, err = txn.Get(id)
		if err != nil {
			return fmt.Errorf("获取数据异常")
		}
		err = item.Value(func(val []byte) error {
			err = db.Update(func(txn *badger.Txn) error {
				var mType map[string]interface{}
				err = msgpack.Unmarshal(val, &mType)
				if err == nil {
					newData := lo.Assign(mType, payload)
					ctx, _ := msgpack.Marshal(newData)
					body = []byte(ctx)
					if timestamp != nil {
						entry = badger.NewEntry(id, body)
						entry.WithTTL(time.Duration(newTime.UnixNano()))
						err = txn.SetEntry(entry)
					} else {
						err = txn.Set(id, body)
					}
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

/*
*
 */

func UpdateToString(id []byte, payload string, filename string) ([]byte, error) {
	db, err := database(filename)
	if err != nil {
		defer db.Close()
		return nil, fmt.Errorf("数据库异常")
	}
	var body []byte
	err = db.Update(func(txn *badger.Txn) error {
		err = txn.Set(id, []byte(payload))
		return err
	})
	defer db.Close()
	return body, err
}

/*
*
 */
func Remove(id []byte, filename string) ([]byte, error) {
	db, err := database(filename)
	if err != nil {
		defer db.Close()
		return nil, fmt.Errorf("数据库异常")
	}
	err = db.Update(func(txn *badger.Txn) error {
		err = txn.Delete(id)
		return err
	})
	defer db.Close()
	return nil, err
}
