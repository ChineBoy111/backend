package test

import (
	"encoding/json"
	"fmt"
	"time"

	// "io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// ! go test -run TestClean
func TestClean(t *testing.T) {
	// cmd := exec.Command("rm", "-rf", "../src")
	// cmd.Run()
	log.Println("Done!")
}

var rootDir string
var jsonFname = "./tree.json"
var sep = string(filepath.Separator)

func mkdir(dirname string) {
	if dirname == "" {
		return
	}
	log.Println(dirname)
	// err := os.MkdirAll(rootDir+sep+dirname, fs.ModePerm /* 0777 */)
	// if err != nil {
	// 	panic(err.Error())
	// }
}

// ! go test -run TestBuildByMap
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

// ! go test -run TestBuildByDirNode
type DirNode struct {
	DirName  string    `json:"dirname"`
	SubNodes []DirNode `json:"subdirs"`
}

func (dirNode *DirNode) loadJson() {
	currDir, _ := os.Getwd()
	rootDir = currDir[0:strings.LastIndex(currDir, sep)]
	log.Printf("currDir = %s\n", currDir)
	log.Printf("rootDir = %s\n", rootDir)
	jsonBytes, _ := os.ReadFile(currDir + sep + jsonFname)
	err := json.Unmarshal(jsonBytes, dirNode)
	if err != nil {
		panic(err.Error())
	}
}

func (dirNode *DirNode) parseNode(prefix string) {
	if dirNode.DirName != "" {
		if prefix != "" {
			dirNode.DirName = prefix + sep + dirNode.DirName
		}
		prefix = dirNode.DirName
		mkdir(dirNode.DirName)
	}

	if dirNode.SubNodes != nil {
		for _, subNode := range dirNode.SubNodes {
			subNode.parseNode(prefix)
		}
	}
}

// ! go test -run TestBuildByDirNode
func TestBuildByDirNode(t *testing.T) {
	rootDir = ""
	var dirNode DirNode
	dirNode.loadJson()
	dirNode.parseNode("")
	log.Println("Done!")
}

// ! 从一个已关闭的空通道中读，返回通道元素类型的零值和 false（表示读失败）
func TestNotify(t *testing.T) {
	ch := make(chan int) // make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		close(ch)
	}()
	fmt.Println(<-ch) // 0
}
