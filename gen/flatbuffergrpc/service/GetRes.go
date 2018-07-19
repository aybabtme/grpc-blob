// automatically generated by the FlatBuffers compiler, do not modify

package service

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type GetRes struct {
	_tab flatbuffers.Table
}

func GetRootAsGetRes(buf []byte, offset flatbuffers.UOffsetT) *GetRes {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &GetRes{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *GetRes) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *GetRes) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *GetRes) Blob(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *GetRes) BlobLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *GetRes) BlobBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func GetResStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func GetResAddBlob(builder *flatbuffers.Builder, blob flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(blob), 0)
}
func GetResStartBlobVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func GetResEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}