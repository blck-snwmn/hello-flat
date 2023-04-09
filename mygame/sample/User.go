// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package sample

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type User struct {
	_tab flatbuffers.Table
}

func GetRootAsUser(buf []byte, offset flatbuffers.UOffsetT) *User {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &User{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsUser(buf []byte, offset flatbuffers.UOffsetT) *User {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &User{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *User) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *User) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *User) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *User) Pos(obj *Position) *Position {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Position)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *User) Color() Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return Color(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 2
}

func (rcv *User) MutateColor(n Color) bool {
	return rcv._tab.MutateInt8Slot(8, int8(n))
}

func UserStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func UserAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func UserAddPos(builder *flatbuffers.Builder, pos flatbuffers.UOffsetT) {
	builder.PrependStructSlot(1, flatbuffers.UOffsetT(pos), 0)
}
func UserAddColor(builder *flatbuffers.Builder, color Color) {
	builder.PrependInt8Slot(2, int8(color), 2)
}
func UserEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}