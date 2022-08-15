package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	pdr "fb-tests/v2/PublicDataRecord"
	"fb-tests/v2/excelize"

	"github.com/edsrzf/mmap-go"
	flatbuffers "github.com/google/flatbuffers/go"
)

func createRecord2(builder *flatbuffers.Builder, args ...string) flatbuffers.UOffsetT {
	codeStr := builder.CreateString(args[0])
	authorsStr := builder.CreateString(args[1])
	dopStr := builder.CreateString(args[2])
	titleStr := builder.CreateString(args[3])
	popStr := builder.CreateString(args[4])
	primaryStr := builder.CreateString(args[5])
	secondaryStr := builder.CreateString(args[6])
	ifMult1Str := builder.CreateString(args[7])
	regionStr := builder.CreateString(args[8])
	countryStr := builder.CreateString(args[9])
	afStr := builder.CreateString(args[10])
	ifMult2Str := builder.CreateString(args[11])

	pdr.RecordStart(builder)
	pdr.RecordAddCode(builder, codeStr)
	pdr.RecordAddAuthors(builder, authorsStr)
	pdr.RecordAddDateOfPublication(builder, dopStr)
	pdr.RecordAddTitle(builder, titleStr)
	pdr.RecordAddPlaceOfPublication(builder, popStr)
	pdr.RecordAddPrimaryUbis(builder, primaryStr)
	pdr.RecordAddSecondaryUbis(builder, secondaryStr)
	pdr.RecordAddIfMultiple(builder, ifMult1Str)
	pdr.RecordAddRegion(builder, regionStr)
	pdr.RecordAddCountry(builder, countryStr)
	pdr.RecordAddAnalyticalFramework(builder, afStr)
	pdr.RecordAddMultipleFrameworks(builder, ifMult2Str)
	temp := pdr.RecordEnd(builder)

	builder.Finish(temp)

	return temp
}

func main() {
	builder := flatbuffers.NewBuilder(0)

	if len(os.Args) > 1 && os.Args[1] == "write" {
		f, err := excelize.OpenFile("metadata.xlsx")

		if err != nil {
			panic(err)
		}

		// for _, entry := range f.GetSheetList() {
		// 	fmt.Println(entry)
		// }

		rows, _ := f.GetRows("Master list - articles")

		var recs []flatbuffers.UOffsetT

		for _, row := range rows[1:] {
			rec := createRecord2(builder, row...)

			recs = append(recs, rec)

			// for _, colCell := range row {
			// 	fmt.Println(colCell)

			// 	// fmt.Print(colCell, "\t")

			// }

			// fmt.Println()
		}

		pdr.RecordsStartRecsVector(builder, len(recs))

		for i := len(recs) - 1; i > 0; i-- {
			builder.PrependUOffsetT(recs[i])
		}

		recVec := builder.EndVector(len(recs))

		pdr.RecordsStart(builder)
		pdr.RecordsAddRecs(builder, recVec)
		allRecs := pdr.RecordsEnd(builder)
		builder.Finish(allRecs)

		bytes := builder.FinishedBytes()

		err = ioutil.WriteFile("recs.dat", bytes, 0644)

		if err != nil {
			panic(err)
		}

		fmt.Printf("Wrote out %d people in %d bytes.\n", len(recs), len(bytes))

		os.Exit(0)
	}

	if len(os.Args) > 1 && os.Args[1] == "read" {
		if len(os.Args) < 2 {
			fmt.Println("Please provide an index number, zero offset.")

			os.Exit(1)
		}

		i, err := strconv.ParseInt(os.Args[2], 10, 64)

		if err != nil {
			panic(err)
		}

		f, err := os.Open("recs.dat")

		if err != nil {
			panic(err)
		}

		mmap, _ := mmap.Map(f, os.O_RDONLY, 0)

		defer mmap.Unmap()

		recs := pdr.GetRootAsRecords(mmap, 0)

		if i > int64(recs.RecsLength())-1 {
			fmt.Println("Index out-of-bounds.")

			os.Exit(1)
		}

		record := new(pdr.Record)

		recs.Recs(record, int(i))

		fmt.Println(string(record.Code()))
		fmt.Println(string(record.Authors()))
		fmt.Println(string(record.DateOfPublication()))

		os.Exit(0)
	}

	if len(os.Args) > 1 && os.Args[1] == "modify" {
		if len(os.Args) < 3 {
			fmt.Println("Please provide an index number, zero offset, and a new author name.")

			os.Exit(1)
		}

		i, err := strconv.ParseInt(os.Args[2], 10, 64)

		if err != nil {
			panic(err)
		}

		newAuthor := os.Args[3]

		f, err := os.Open("recs.dat")

		if err != nil {
			panic(err)
		}

		mmap, _ := mmap.Map(f, os.O_RDWR, 0)

		defer mmap.Unmap()

		recs := pdr.GetRootAsRecords(mmap, 10)

		unpackedRecs := recs.UnPack().Recs

		unpackedRecs[i].Authors = newAuthor

		pr := unpackedRecs[i].Pack(builder)

		builder.Finish(pr)
	}

	fmt.Println("Unknowm command.")

	os.Exit(1)
}
