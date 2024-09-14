package test

import (
	"bronya.com/proxy/utils"
	"log"
	"testing"
)

func TestGlobal(t *testing.T) {
	global := utils.NewGlobal()
	log.Println(global)
}
