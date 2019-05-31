package common

import (
	"fmt"
	"regexp"
	"strings"
)

var expectKeys = []string{"limit", "fields"}

type includes struct {
	items []*includeItem
}

type includeItem struct {
	name   string
	limit  uint
	fields []string
	items  []*includeItem
}

func (this *includes) AddItem(item *includeItem) {
	this.items = append(this.items, item)
}

func (this *includeItem) AddItem(item *includeItem) {
	this.items = append(this.items, item)
}

func (this *includeItem) New(name string) *includeItem {
	this.name = name
	return this
}

func (this *includeItem) SetLimit(limit uint) *includeItem {
	this.limit = limit
	return this
}

func (this *includeItem) SetFields(fields []string) *includeItem {
	this.fields = fields
	return this
}

type resloveInclude struct {
	params string
	data   []map[string]interface{}
}

func NewResloveInclude(params string) *resloveInclude {
	return &resloveInclude{
		params: params,
	}
}

func (this *resloveInclude) Reslove() {
	fields := strings.Split(this.params, ",")

	for _, relation := range fields {
		// 第一个:分割关系
		item := make(map[string]interface{})
		i := strings.Index(relation, ":")

		if i >= 0 {
			with := StrToCamel(relation[:i])
			fmt.Println(with)
			item[with] = make(map[string]interface{})
			//再次分割子关联
			subParams := strings.Split(relation[i+1:], ".")
			regSubWith := regexp.MustCompile(`:?(\w+):?`)
			if len(subParams) > 0 {
				item[with]["include"] = []string{}
			}
			for _, sub := range subParams {
				submatch := regSubWith.FindStringSubmatch(sub)
				if len(submatch) > 1 {
					if this.allowKey(submatch[1]) {
						fmt.Println(submatch[1])
					}
				}
			}
			// 匹配条件 limit fields

			//regLimit:=regexp.MustCompile(`limit\((\d+)\)`)
			//fmt.Println(relation[i:])
			//submatch := regLimit.FindStringSubmatch(relation[i:])
			//fmt.Println(submatch)
		}
	}
}

func (this *resloveInclude) allowKey(key string) bool {
	for _, v := range expectKeys {
		if key == v {
			return true
		}
	}
	return false
}
