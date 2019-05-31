package test

import (
	"github.com/zhouchang2017/epp/common"
	"testing"
)

// 特定关键字

var KEYWORDS = []string{
	"limit",
	"fields",
	//"sort_by",
	//"order",
}

// http://api.test/hubs/1?include=hot_users:limit(3).posts:fields(id|title):limit(3)
func TestResloveInclude(t *testing.T) {
	str := "hot_users:limit(3).posts:fields(id|title):limit(3),likes"

	common.NewResloveInclude(str).Reslove()
}
