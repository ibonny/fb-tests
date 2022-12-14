// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package PublicDataRecord

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RecordsT struct {
	Recs []*RecordT `json:"recs"`
}

func (t *RecordsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	recsOffset := flatbuffers.UOffsetT(0)
	if t.Recs != nil {
		recsLength := len(t.Recs)
		recsOffsets := make([]flatbuffers.UOffsetT, recsLength)
		for j := 0; j < recsLength; j++ {
			recsOffsets[j] = t.Recs[j].Pack(builder)
		}
		RecordsStartRecsVector(builder, recsLength)
		for j := recsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(recsOffsets[j])
		}
		recsOffset = builder.EndVector(recsLength)
	}
	RecordsStart(builder)
	RecordsAddRecs(builder, recsOffset)
	return RecordsEnd(builder)
}

func (rcv *Records) UnPackTo(t *RecordsT) {
	recsLength := rcv.RecsLength()
	t.Recs = make([]*RecordT, recsLength)
	for j := 0; j < recsLength; j++ {
		x := Record{}
		rcv.Recs(&x, j)
		t.Recs[j] = x.UnPack()
	}
}

func (rcv *Records) UnPack() *RecordsT {
	if rcv == nil { return nil }
	t := &RecordsT{}
	rcv.UnPackTo(t)
	return t
}

type Records struct {
	_tab flatbuffers.Table
}

func GetRootAsRecords(buf []byte, offset flatbuffers.UOffsetT) *Records {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Records{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsRecords(buf []byte, offset flatbuffers.UOffsetT) *Records {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Records{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Records) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Records) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Records) Recs(obj *Record, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Records) RecsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func RecordsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func RecordsAddRecs(builder *flatbuffers.Builder, recs flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(recs), 0)
}
func RecordsStartRecsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func RecordsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
