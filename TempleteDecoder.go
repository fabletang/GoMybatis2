package GoMybatis

import (
	"github.com/beevik/etree"
	"reflect"
)

type TemplateDecoder interface {
	SetPrintElement(print bool)
	DecodeTree(tree map[string]etree.Token, beanType reflect.Type) error
}
