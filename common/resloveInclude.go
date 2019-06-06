package common

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
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

func (this includes) String() string {
	items := bytes.NewBufferString("")
	for _, item := range this.items {
		items.WriteString(item.String())
		items.WriteString("\n")
	}
	return items.String()
}

func (this *includes) AddItem(item *includeItem) {
	this.items = append(this.items, item)
}

func (this includeItem) String() string {
	items := bytes.NewBufferString("[")
	for _, item := range this.items {
		items.WriteString("\t")
		items.WriteString(item.String())
		items.WriteString("\n")
	}
	items.WriteString("]")
	return fmt.Sprintf("name:%s\n limit:%d\n fields:%s\n items:%s\n", this.name, this.limit, this.fields, items.String())
}

func (this *includes) newItem(name string) *includeItem {
	inc := &includeItem{name: name}
	this.items = append(this.items, inc)
	return inc
}

func (this *includeItem) AddItem(item *includeItem) {
	this.items = append(this.items, item)
}

func (this *includeItem) newItem(name string) *includeItem {
	inc := &includeItem{name: name}
	this.items = append(this.items, inc)
	return inc
}

func NewIncludeItem(name string) *includeItem {
	return &includeItem{name: name}
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
	data   *includes
}

func NewResloveInclude(params string) *resloveInclude {
	return &resloveInclude{
		params: params,
	}
}

func (this *resloveInclude) Reslove() {
	fields := strings.Split(this.params, ",")
	inc := &includes{}
	for _, relation := range fields {
		// 第一个:分割关系
		i := strings.Index(relation, ":")

		if i >= 0 {
			with := StrToCamel(relation[:i])
			fmt.Println(with)
			item := inc.newItem(with)

			// 匹配条件 limit fields
			regLimit := regexp.MustCompile(`limit\((\d+)\)`)

			submatch := regLimit.FindStringSubmatch(relation[i:])

			if len(submatch) >= 2 {
				fmt.Println("limit: ", submatch)
				fmt.Println(submatch[1])
				atoi, _ := strconv.Atoi(submatch[1])
				item.SetLimit(uint(atoi))
			}

			// 匹配fields

			//再次分割子关联
			subParams := strings.Split(relation[i+1:], ".")

			regSubWith := regexp.MustCompile(`:?(\w+):?`)

			for _, sub := range subParams {
				submatch := regSubWith.FindStringSubmatch(sub)
				if len(submatch) >= 2 {
					if this.allowKey(submatch[1]) {
						item.newItem(submatch[1])
						fmt.Println(submatch[1])
					}
				}
			}

		}
	}

	fmt.Println(inc)
}

func (this *resloveInclude) allowKey(key string) bool {
	for _, v := range expectKeys {
		if key == v {
			return false
		}
	}
	return true
}
