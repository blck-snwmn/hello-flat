// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package sample

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PosionT struct {
	Name string `json:"name"`
	EffectType string `json:"effect_type"`
}

func (t *PosionT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	effectTypeOffset := flatbuffers.UOffsetT(0)
	if t.EffectType != "" {
		effectTypeOffset = builder.CreateString(t.EffectType)
	}
	PosionStart(builder)
	PosionAddName(builder, nameOffset)
	PosionAddEffectType(builder, effectTypeOffset)
	return PosionEnd(builder)
}

func (rcv *Posion) UnPackTo(t *PosionT) {
	t.Name = string(rcv.Name())
	t.EffectType = string(rcv.EffectType())
}

func (rcv *Posion) UnPack() *PosionT {
	if rcv == nil { return nil }
	t := &PosionT{}
	rcv.UnPackTo(t)
	return t
}

type Posion struct {
	_tab flatbuffers.Table
}

func GetRootAsPosion(buf []byte, offset flatbuffers.UOffsetT) *Posion {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Posion{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsPosion(buf []byte, offset flatbuffers.UOffsetT) *Posion {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Posion{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Posion) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Posion) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Posion) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Posion) EffectType() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func PosionStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PosionAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func PosionAddEffectType(builder *flatbuffers.Builder, effectType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(effectType), 0)
}
func PosionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
