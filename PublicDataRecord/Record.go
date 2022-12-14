// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package PublicDataRecord

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RecordT struct {
	Code string `json:"code"`
	Authors string `json:"authors"`
	DateOfPublication string `json:"date_of_publication"`
	Title string `json:"title"`
	PlaceOfPublication string `json:"place_of_publication"`
	PrimaryUbis string `json:"primary_ubis"`
	SecondaryUbis string `json:"secondary_ubis"`
	IfMultiple string `json:"if_multiple"`
	Region string `json:"region"`
	Country string `json:"country"`
	AnalyticalFramework string `json:"analytical_framework"`
	MultipleFrameworks string `json:"multiple_frameworks"`
}

func (t *RecordT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	codeOffset := builder.CreateString(t.Code)
	authorsOffset := builder.CreateString(t.Authors)
	dateOfPublicationOffset := builder.CreateString(t.DateOfPublication)
	titleOffset := builder.CreateString(t.Title)
	placeOfPublicationOffset := builder.CreateString(t.PlaceOfPublication)
	primaryUbisOffset := builder.CreateString(t.PrimaryUbis)
	secondaryUbisOffset := builder.CreateString(t.SecondaryUbis)
	ifMultipleOffset := builder.CreateString(t.IfMultiple)
	regionOffset := builder.CreateString(t.Region)
	countryOffset := builder.CreateString(t.Country)
	analyticalFrameworkOffset := builder.CreateString(t.AnalyticalFramework)
	multipleFrameworksOffset := builder.CreateString(t.MultipleFrameworks)
	RecordStart(builder)
	RecordAddCode(builder, codeOffset)
	RecordAddAuthors(builder, authorsOffset)
	RecordAddDateOfPublication(builder, dateOfPublicationOffset)
	RecordAddTitle(builder, titleOffset)
	RecordAddPlaceOfPublication(builder, placeOfPublicationOffset)
	RecordAddPrimaryUbis(builder, primaryUbisOffset)
	RecordAddSecondaryUbis(builder, secondaryUbisOffset)
	RecordAddIfMultiple(builder, ifMultipleOffset)
	RecordAddRegion(builder, regionOffset)
	RecordAddCountry(builder, countryOffset)
	RecordAddAnalyticalFramework(builder, analyticalFrameworkOffset)
	RecordAddMultipleFrameworks(builder, multipleFrameworksOffset)
	return RecordEnd(builder)
}

func (rcv *Record) UnPackTo(t *RecordT) {
	t.Code = string(rcv.Code())
	t.Authors = string(rcv.Authors())
	t.DateOfPublication = string(rcv.DateOfPublication())
	t.Title = string(rcv.Title())
	t.PlaceOfPublication = string(rcv.PlaceOfPublication())
	t.PrimaryUbis = string(rcv.PrimaryUbis())
	t.SecondaryUbis = string(rcv.SecondaryUbis())
	t.IfMultiple = string(rcv.IfMultiple())
	t.Region = string(rcv.Region())
	t.Country = string(rcv.Country())
	t.AnalyticalFramework = string(rcv.AnalyticalFramework())
	t.MultipleFrameworks = string(rcv.MultipleFrameworks())
}

func (rcv *Record) UnPack() *RecordT {
	if rcv == nil { return nil }
	t := &RecordT{}
	rcv.UnPackTo(t)
	return t
}

type Record struct {
	_tab flatbuffers.Table
}

func GetRootAsRecord(buf []byte, offset flatbuffers.UOffsetT) *Record {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Record{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsRecord(buf []byte, offset flatbuffers.UOffsetT) *Record {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Record{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Record) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Record) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Record) Code() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Record) Authors() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Record) DateOfPublication() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Record) Title() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Record) PlaceOfPublication() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Record) PrimaryUbis() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Record) SecondaryUbis() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Record) IfMultiple() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Record) Region() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Record) Country() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Record) AnalyticalFramework() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Record) MultipleFrameworks() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func RecordStart(builder *flatbuffers.Builder) {
	builder.StartObject(12)
}
func RecordAddCode(builder *flatbuffers.Builder, code flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(code), 0)
}
func RecordAddAuthors(builder *flatbuffers.Builder, authors flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(authors), 0)
}
func RecordAddDateOfPublication(builder *flatbuffers.Builder, dateOfPublication flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(dateOfPublication), 0)
}
func RecordAddTitle(builder *flatbuffers.Builder, title flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(title), 0)
}
func RecordAddPlaceOfPublication(builder *flatbuffers.Builder, placeOfPublication flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(placeOfPublication), 0)
}
func RecordAddPrimaryUbis(builder *flatbuffers.Builder, primaryUbis flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(primaryUbis), 0)
}
func RecordAddSecondaryUbis(builder *flatbuffers.Builder, secondaryUbis flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(secondaryUbis), 0)
}
func RecordAddIfMultiple(builder *flatbuffers.Builder, ifMultiple flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(ifMultiple), 0)
}
func RecordAddRegion(builder *flatbuffers.Builder, region flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(region), 0)
}
func RecordAddCountry(builder *flatbuffers.Builder, country flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(country), 0)
}
func RecordAddAnalyticalFramework(builder *flatbuffers.Builder, analyticalFramework flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(analyticalFramework), 0)
}
func RecordAddMultipleFrameworks(builder *flatbuffers.Builder, multipleFrameworks flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(multipleFrameworks), 0)
}
func RecordEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
