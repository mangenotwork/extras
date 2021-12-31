package boltdb

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/boltdb/bolt"
)

type BoltDB struct {
	dbFile string
	db *bolt.DB
}

func NewBoltDB(fileName string) (bo *BoltDB, err error) {
	bo = &BoltDB{
		dbFile: fileName,
	}
	bo.db, err = bolt.Open(bo.dbFile, 0600, nil)
	return
}

func (bo *BoltDB) Conn() (*bolt.DB, error) {
	var err error
	if bo.db == nil {
		bo.db, err = bolt.Open(bo.dbFile, 0600, nil)
	}
	return bo.db, err
}

func (bo *BoltDB) Close() {
	bo.db.Close()
}

func (bo *BoltDB) CreateTable(tableName string) error {
	//2.创建表
	return bo.db.Update(func(tx *bolt.Tx) error {

		//判断要创建的表是否存在
		b := tx.Bucket([]byte(tableName))
		if b == nil {
			//创建叫"MyBucket"的表
			_, err := tx.CreateBucket([]byte(tableName))
			return err
		}
		//一定要返回nil
		return nil
	})
}

func (bo *BoltDB) GetTable() (list []string, err error){
	list = make([]string, 0)
	err = bo.db.Update(func(tx *bolt.Tx) error {
		_=tx.ForEach(func(name []byte, b *bolt.Bucket) error {
			list = append(list, string(name))
			return nil
		})
		return nil
	})
	return
}

func (bo *BoltDB) Insert(table, key, value string) error {
	return bo.db.Update(func(tx *bolt.Tx) error {

		//取出叫"MyBucket"的表
		b := tx.Bucket([]byte(table))

		//往表里面存储数据
		if b != nil {
			//插入的键值对数据类型必须是字节数组
			return b.Put([]byte(key), []byte(value))
		}

		//一定要返回nil
		return nil
	})
}

func (bo *BoltDB) Select(table, key string) (data string, err error){
	err = bo.db.View(func(tx *bolt.Tx) error {

		//取出叫"MyBucket"的表
		b := tx.Bucket([]byte(table))

		//往表里面存储数据
		if b != nil {

			bData := b.Get([]byte(key))
			data = string(bData)
			return err
		}

		//一定要返回nil
		return nil
	})
	return
}

// 前缀扫描
func (bo *BoltDB) SelectPrefix(table, key string) (data map[string]string, err error) {
	data = make(map[string]string)
	err = bo.db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(table)).Cursor()
		prefix := []byte(key)
		for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
			data[string(k)] = string(v)
		}
		return nil
	})
	return
}

// 区间扫描
// use : SelectInterval(table, "2021-01-01 00:00:00", "2022-01-01 00:00:00")
func (bo *BoltDB) SelectInterval(table, keyA, keyB string) (data map[string]string, err error) {
	data = make(map[string]string)
	err = bo.db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(table)).Cursor()
		for k, v := c.Seek([]byte(keyA)); k != nil && bytes.Compare(k, []byte(keyB)) <= 0; k, v = c.Next() {
			data[string(k)] = string(v)
		}
		return nil
	})
	return
}

// SelectFront   从前面获取多少个
func (bo *BoltDB) SelectFront(table string, count int) (data map[string]string) {
	data = make(map[string]string)
	_=bo.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		i := 0
		_=b.ForEach(func(k, v []byte) error {
			if i > count {
				return nil
			}
			data[string(k)] = string(v)
			i++
			return nil
		})
		return nil
	})
	return
}

// Fro
func (bo *BoltDB) SelectFor(table string, f func(k,v []byte) (map[string]string, error), data map[string]string ) error {
	return bo.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var err error
			data, err = f(k,v)
			return err
		}
		return nil
	})
}

// 数据库备份
func (bo *BoltDB) BackupFromHttp(w http.ResponseWriter, req *http.Request) {
	err := bo.db.View(func(tx *bolt.Tx) error {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", `attachment; filename="`+bo.dbFile+`"`)
		w.Header().Set("Content-Length", strconv.Itoa(int(tx.Size())))
		_, err := tx.WriteTo(w)
		return err
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}



/*

## 方法总结

#### 打开数据库：
```
db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
```

#### 读写数据库：
```
err := db.Update(func(tx *bolt.Tx) error {
    ...
    return nil
})
```

#### 只读数据库：
```
err := db.View(func(tx *bolt.Tx) error {
    ...
    return nil
})
```

#### 批量读写数据库：
```
err := db.Batch(func(tx *bolt.Tx) error {
    ...
    return nil
})
```

#### 事务:
```
// Start a writable transaction.
tx, err := db.Begin(true)
if err != nil {
    return err
}
defer tx.Rollback()

// Use the transaction...
_, err := tx.CreateBucket([]byte("MyBucket"))
if err != nil {
    return err
}

// Commit the transaction and check for error.
if err := tx.Commit(); err != nil {
    return err
}
```

#### 将key/value对写入Bucket：
```
db.Update(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte("MyBucket"))
    err := b.Put([]byte("answer"), []byte("42"))
    return err
})
```

#### 根据key遍历：
```
db.View(func(tx *bolt.Tx) error {
    // Assume bucket exists and has keys
    b := tx.Bucket([]byte("MyBucket"))

    c := b.Cursor()

    for k, v := c.First(); k != nil; k, v = c.Next() {
        fmt.Printf("key=%s, value=%s\n", k, v)
    }
    return nil
})


cursor支持的操作方法：
First()  Move to the first key.
Last()   Move to the last key.
Seek()   Move to a specific key.
Next()   Move to the next key.
Prev()   Move to the previous key.
```

#### 统计：
```
go func() {
    // Grab the initial stats.
    prev := db.Stats()

    for {
        // Wait for 10s.
        time.Sleep(10 * time.Second)

        // Grab the current stats and diff them.
        stats := db.Stats()
        diff := stats.Sub(&prev)

        // Encode stats to JSON and print to STDERR.
        json.NewEncoder(os.Stderr).Encode(diff)

        // Save stats for the next loop.
        prev = stats
    }
}()
```
#### 只读模式：
```
db, err := bolt.Open("my.db", 0666, &bolt.Options{ReadOnly: true})
if err != nil {
    log.Fatal(err)
}
```
*/



