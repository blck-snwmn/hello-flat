// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package sample

import (
	"strconv"
	flatbuffers "github.com/google/flatbuffers/go"
)

type ItemContent byte

const (
	ItemContentNONE   ItemContent = 0
	ItemContentPosion ItemContent = 1
	ItemContentDrink  ItemContent = 2
)

var EnumNamesItemContent = map[ItemContent]string{
	ItemContentNONE:   "NONE",
	ItemContentPosion: "Posion",
	ItemContentDrink:  "Drink",
}

var EnumValuesItemContent = map[string]ItemContent{
	"NONE":   ItemContentNONE,
	"Posion": ItemContentPosion,
	"Drink":  ItemContentDrink,
}

func (v ItemContent) String() string {
	if s, ok := EnumNamesItemContent[v]; ok {
		return s
	}
	return "ItemContent(" + strconv.FormatInt(int64(v), 10) + ")"
}

type ItemContentT struct {
	Type ItemContent
	Value interface{}
}

func (t *ItemContentT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	switch t.Type {
	case ItemContentPosion:
		return t.Value.(*PosionT).Pack(builder)
	case ItemContentDrink:
		return t.Value.(*DrinkT).Pack(builder)
	}
	return 0
}

func (rcv ItemContent) UnPack(table flatbuffers.Table) *ItemContentT {
	switch rcv {
	case ItemContentPosion:
		var x Posion
		x.Init(table.Bytes, table.Pos)
		return &ItemContentT{ Type: ItemContentPosion, Value: x.UnPack() }
	case ItemContentDrink:
		var x Drink
		x.Init(table.Bytes, table.Pos)
		return &ItemContentT{ Type: ItemContentDrink, Value: x.UnPack() }
	}
	return nil
}