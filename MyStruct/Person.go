// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package MyStruct

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Person struct {
	_tab flatbuffers.Table
}

func GetRootAsPerson(buf []byte, offset flatbuffers.UOffsetT) *Person {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Person{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsPerson(buf []byte, offset flatbuffers.UOffsetT) *Person {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Person{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Person) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Person) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Person) FirstName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Person) LastName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func PersonStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PersonAddFirstName(builder *flatbuffers.Builder, firstName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(firstName), 0)
}
func PersonAddLastName(builder *flatbuffers.Builder, lastName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(lastName), 0)
}
func PersonEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
