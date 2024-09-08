package test

import (
	"bronya.com/gin-gorm/src/conf"
	"bronya.com/gin-gorm/src/util"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"time"
	// "io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var rootDir string
var jsonFname = "./tree.json"
var sep = string(filepath.Separator)

func mkdir(dirname string) {
	if dirname == "" {
		return
	}
	log.Printf("dirname = %s\n", dirname)
	// err := os.MkdirAll(rootDir+sep+dirname, fs.ModePerm /* 0777 */)
	// if err != nil {
	// 	panic(err.Error())
	// }
}

func loadJson(jsonMap *map[string]any) {
	currDir, _ := os.Getwd()
	rootDir = currDir[0:strings.LastIndex(currDir, sep)]
	log.Printf("currDir = %s\n", currDir)
	log.Printf("rootDir = %s\n", rootDir)
	jsonBytes, _ := os.ReadFile(currDir + sep + jsonFname)
	err := json.Unmarshal(jsonBytes, jsonMap)
	if err != nil {
		panic(err.Error())
	}
}

func parseMap(jsonMap *map[string]any, prefix string) {
	for _, v := range *jsonMap {
		//! 类型 switch
		switch v.(type) {
		case string:
			{
				dirname, _ := v.(string) // 类型断言
				if dirname == "" {
					continue
				}
				if prefix != "" {
					dirname = prefix + sep + dirname
				}
				prefix = dirname
				mkdir(dirname)
			}
		case []any:
			parseArr(v.([]any), prefix)
		}
	}
}

func parseArr(jsonArr []any, prefix string) {
	for _, v := range jsonArr {
		mapV, _ := v.(map[string]any)
		parseMap(&mapV, prefix)
	}
}

// ! go test -run TestBuildByMap
func TestBuildByMap(t *testing.T) {
	rootDir = ""
	var jsonMap map[string]any
	loadJson(&jsonMap)
	parseMap(&jsonMap, "")
	log.Println("Done!")
}

type Dir struct {
	DirName string `json:"dirname"`
	SubDirs []Dir  `json:"subdirs"`
}

func (dir *Dir) loadJson() {
	currDir, _ := os.Getwd()
	rootDir = currDir[0:strings.LastIndex(currDir, sep)]
	log.Printf("currDir = %s\n", currDir)
	log.Printf("rootDir = %s\n", rootDir)
	jsonBytes, _ := os.ReadFile(currDir + sep + jsonFname)
	err := json.Unmarshal(jsonBytes, dir)
	if err != nil {
		panic(err.Error())
	}
}

func (dir *Dir) parseDir(prefix string) {
	if dir.DirName != "" {
		if prefix != "" {
			dir.DirName = prefix + sep + dir.DirName
		}
		prefix = dir.DirName
		mkdir(dir.DirName)
	}

	if dir.SubDirs != nil {
		for _, subNode := range dir.SubDirs {
			subNode.parseDir(prefix)
		}
	}
}

// ! go test -run TestBuildByDir
func TestBuildByDir(t *testing.T) {
	rootDir = ""
	var dir Dir
	dir.loadJson()
	dir.parseDir("")
	log.Println("Done!")
}

// ! 从一个已关闭的空通道中读，返回通道元素类型的零值和 false（表示读失败）
// ! go test -run TestNotify
func TestNotify(t *testing.T) {
	ch := make(chan int) // make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		close(ch)
	}()
	fmt.Println(<-ch) // 0
}

// ! go test -run TestRedisCli
func TestRedisCli(t *testing.T) {
	redisCli, err := conf.ConnRedis() //! 连接 redis，创建表
	if err != nil {
		panic(fmt.Sprintf("Connect redis server error %s", err.Error()))
	}

	expiration := viper.GetDuration("db.redis.expiration") // 10 分钟

	//! redisCli.Set(context.Background(), key, value, expiration).Err()
	err = redisCli.Set(context.Background(), "username", "root", expiration*time.Second).Err()
	if err != nil {
		log.Printf("Redis set error %s\n", err.Error())
	}

	//! redisCli.Set(context.Background(), key, value, expiration).Err()
	err = redisCli.Set(context.Background(), "password", "0228", expiration*time.Second).Err()
	if err != nil {
		log.Printf("Redis set error %s\n", err.Error())
	}

	//! redisCli.Get(context.Background(), key).Result()
	username, err := redisCli.Get(context.Background(), "username").Result()
	if err != nil {
		log.Printf("Redis get error %s\n", err.Error())
	}
	log.Printf("Redis get username = %s", username)

	//! redisCli.Del(context.Background(), key1, key2, ...).Err()
	err = redisCli.Del(context.Background(), "username", "password").Err()
	if err != nil {
		log.Printf("Redis delete error %s\n", err.Error())
	}
}

// ! go test -run TestToken
func TestToken(t *testing.T) {
	viper.AddConfigPath("../")
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Read in config error %s", err.Error()))
	}

	log.Println("==================== Generate token ====================")
	tokStr, err := util.GenToken(1, "root")
	if err != nil {
		panic(fmt.Sprintf("Generate token error %s", err.Error()))
	}
	log.Printf("token = %s\n", tokStr)
	log.Println("==================== Parse token ====================")
	payload, err, isValid := util.ParseToken(tokStr) // tokStr + "suffix"
	if err != nil {
		panic(fmt.Sprintf("Parse token error %s, isValid = %v", err.Error(), isValid))
	}
	log.Printf("payload = %v, isValid = %v\n", payload, isValid)
}
