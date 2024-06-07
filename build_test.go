package wms_test

import (
	"github.com/Domingor/go-blackbox/buildscript"
	"testing"
)

/**
* @Author: Connor
* @Date:   24.6.7 14:45
* @Description:
 */

func TestName(t *testing.T) {
	buildscript.Generate("wms-test", "cloudbyte.top", "cmd/main/main.go", false)
}
