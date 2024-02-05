package api

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

func GetCubelifyToken() (string, error) {
	configDir, _ := os.UserConfigDir()
	db, err := leveldb.OpenFile(filepath.Join(configDir+"/cubelify-overlay/Local Storage/leveldb"), &opt.Options{ReadOnly: true})
	if err != nil {
		return "", err
	}
	defer db.Close()

	// iter over all keys until key contains token
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		if strings.Contains(string(key), "token") {
			return string(iter.Value()[1:]), nil
		}
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		panic(err)
	}

	return "", err
}
