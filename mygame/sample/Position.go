// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package sample

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PositionT struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

func (t *PositionT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	return CreatePosition(builder, t.X, t.Y, t.Z)
}
func (rcv *Position) UnPackTo(t *PositionT) {
	t.X = rcv.X()
	t.Y = rcv.Y()
	t.Z = rcv.Z()
}

func (rcv *Position) UnPack() *PositionT {
	if rcv == nil { return nil }
	t := &PositionT{}
	rcv.UnPackTo(t)
	return t
}

type Position struct {
	_tab flatbuffers.Struct
}

func (rcv *Position) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Position) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Position) X() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *Position) MutateX(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *Position) Y() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}
func (rcv *Position) MutateY(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

func (rcv *Position) Z() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(8))
}
func (rcv *Position) MutateZ(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(8), n)
}

func CreatePosition(builder *flatbuffers.Builder, x float32, y float32, z float32) flatbuffers.UOffsetT {
	builder.Prep(4, 12)
	builder.PrependFloat32(z)
	builder.PrependFloat32(y)
	builder.PrependFloat32(x)
	return builder.Offset()
}
