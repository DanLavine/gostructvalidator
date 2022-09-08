package internal

import "github.com/DanLavine/gostructwalker"

type structWalker struct{}

func NewStructWalker() *structWalker {
	return &structWalker{}
}

func (sw *structWalker) FieldCallback(structParser *gostructwalker.StructParser) {

}
