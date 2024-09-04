package test

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

/////////////////////////////////////////
// Description: Test Clean             //
// Author:      Tiancheng              //
// Run command: go test -run TestClean //
/////////////////////////////////////////

func TestClean(t *testing.T) {
	cmd := exec.Command("rm", "-rf", "../src")
	cmd.Run()
	fmt.Println("Done!")
}

var rootDir string
var jsonFname = "./tree.json"
var sep = string(filepath.Separator)

func mkdir(dirname string) {
	if dirname == "" {
		return
	}
	fmt.Println(dirname)
	err := os.MkdirAll(rootDir+sep+dirname, fs.ModePerm /* 0777 */)
	if err != nil {
		panic(err.Error())
	}
}

//////////////////////////////////////////////
// Description: Test build by Map           //
// Author:      Tiancheng                   //
// Run command: go test -run TestBuildByMap //
//////////////////////////////////////////////

func loadJson(jsonMap *map[string]any) { // load a json file to a map
	currDir, _ := os.Getwd()
	rootDir = currDir[0:strings.LastIndex(currDir, sep)]
	fmt.Printf("currDir = %s\nrootDir = %s\n", currDir, rootDir)
	jsonBytes, _ := os.ReadFile(currDir + sep + jsonFname)
	err := json.Unmarshal(jsonBytes, jsonMap)
	if err != nil {
		panic(err.Error())
	}
}

func parseMap(jsonMap *map[string]any, prefix string) {
	for _, v := range *jsonMap {
		//! switch type
		switch v.(type) {
		case string:
			{
				dirname, _ := v.(string) // type assertion
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

//! go test -run TestBuildByMap
func TestBuildByMap(t *testing.T) {
	rootDir = ""
	var jsonMap map[string]any
	loadJson(&jsonMap)
	parseMap(&jsonMap, "")
	fmt.Println("Done!")
}

//////////////////////////////////////////////////
// Description: Test build by DirNode           //
// Author:      Tiancheng                       //
// Run command: go test -run TestBuildByDirNode //
//////////////////////////////////////////////////

type DirNode struct {
	DirName string `json:"dirname"`
	SubNodes []DirNode `json:"subdirs"`
}

func (dirNode *DirNode) loadJson() {
	currDir, _ := os.Getwd()
	rootDir = currDir[0:strings.LastIndex(currDir, sep)]
	fmt.Printf("currDir = %s\nrootDir = %s\n", currDir, rootDir)
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

//! go test -run TestBuildByDirNode
func TestBuildByDirNode(t *testing.T) {
	rootDir = ""
	var dirNode DirNode
	dirNode.loadJson()
	dirNode.parseNode("")
	fmt.Println("Done!")
}
