package test

import (
	"bronya.com/proxy/utils"
	"log"
	"testing"
)

func TestGlobal(t *testing.T) {
	log.Println(utils.Global)
}
