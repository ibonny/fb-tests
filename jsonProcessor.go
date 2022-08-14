package main

import (
	"encoding/json"
	jsonPeople "fb-tests/v2/JsonPeople"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/edsrzf/mmap-go"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/oliveagle/jsonpath"
)

const endpoint = "https://jsonplaceholder.typicode.com/users"

func getMap(data interface{}, field string) map[string]interface{} {
	v, _ := jsonpath.JsonPathLookup(data, field)

	return v.(map[string]interface{})
}

func getString(data interface{}, field string) string {
	v, _ := jsonpath.JsonPathLookup(data, field)

	return v.(string)
}

func getInt(data interface{}, field string) int32 {
	v, _ := jsonpath.JsonPathLookup(data, field)

	return int32(v.(float64))
}

func getFloat(data interface{}, field string) float32 {
	v, _ := jsonpath.JsonPathLookup(data, field)

	return v.(float32)
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "write" {
		resp, err := http.Get(endpoint)

		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			panic(err)
		}

		var json_data interface{}

		json.Unmarshal(body, &json_data)

		v, _ := jsonpath.JsonPathLookup(json_data, "$.[:]")

		builder := flatbuffers.NewBuilder(0)

		var people []flatbuffers.UOffsetT

		for _, entry := range v.([]interface{}) {
			name := builder.CreateString(getString(entry, "$.name"))
			username := builder.CreateString(getString(entry, "$.username"))
			email := builder.CreateString(getString(entry, "$.email"))
			phone := builder.CreateString(getString(entry, "$.phone"))
			website := builder.CreateString(getString(entry, "$.website"))

			street := builder.CreateString(getString(entry, "$.address.street"))
			suite := builder.CreateString(getString(entry, "$.address.suite"))
			city := builder.CreateString(getString(entry, "$.address.city"))
			zipcode := builder.CreateString(getString(entry, "$.address.zipcode"))

			lat := builder.CreateString(getString(entry, "$.address.geo.lat"))
			lng := builder.CreateString(getString(entry, "$.address.geo.lng"))

			jsonPeople.GeoStart(builder)
			jsonPeople.GeoAddLat(builder, lat)
			jsonPeople.GeoAddLng(builder, lng)
			geo := jsonPeople.GeoEnd(builder)

			jsonPeople.AddressStart(builder)
			jsonPeople.AddressAddStreet(builder, street)
			jsonPeople.AddressAddSuite(builder, suite)
			jsonPeople.AddressAddCity(builder, city)
			jsonPeople.AddressAddCity(builder, zipcode)
			jsonPeople.AddressAddGeo(builder, geo)
			address := jsonPeople.AddressEnd(builder)

			companyName := builder.CreateString(getString(entry, "$.company.name"))
			catchPhrase := builder.CreateString(getString(entry, "$.company.catchPhrase"))
			bs := builder.CreateString(getString(entry, "$.company.bs"))

			jsonPeople.CompanyStart(builder)
			jsonPeople.CompanyAddName(builder, companyName)
			jsonPeople.CompanyAddCatchPhrase(builder, catchPhrase)
			jsonPeople.CompanyAddBs(builder, bs)
			company := jsonPeople.CompanyEnd(builder)

			jsonPeople.PersonStart(builder)
			jsonPeople.PersonAddId(builder, getInt(entry, "$.id"))
			jsonPeople.PersonAddName(builder, name)
			jsonPeople.PersonAddUsername(builder, username)
			jsonPeople.PersonAddEmail(builder, email)
			jsonPeople.PersonAddAddress(builder, address)
			jsonPeople.PersonAddPhone(builder, phone)
			jsonPeople.PersonAddWebsite(builder, website)
			jsonPeople.PersonAddCompany(builder, company)
			p := jsonPeople.PersonEnd(builder)

			people = append(people, p)
		}

		jsonPeople.PeopleStartPeopleVector(builder, len(people))

		for _, person := range people {
			builder.PrependUOffsetT(person)
		}

		peopleVec := builder.EndVector(len(people))

		jsonPeople.PeopleStart(builder)
		jsonPeople.PeopleAddPeople(builder, peopleVec)
		allPeople := jsonPeople.PeopleEnd(builder)
		builder.Finish(allPeople)

		bytes := builder.FinishedBytes()

		err = ioutil.WriteFile("people.dat", bytes, 0644)

		if err != nil {
			panic(err)
		}

		fmt.Printf("Wrote out %d people in %d bytes.\n", len(people), len(bytes))

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

		f, err := os.Open("people.dat")

		if err != nil {
			panic(err)
		}

		mmap, _ := mmap.Map(f, os.O_RDONLY, 0)

		defer mmap.Unmap()

		newPeople := jsonPeople.GetRootAsPeople(mmap, 0)

		if i > int64(newPeople.PeopleLength())-1 {
			fmt.Println("Index out-of-bounds.")

			os.Exit(1)
		}

		person := new(jsonPeople.Person)

		newPeople.People(person, int(i))

		fmt.Println(string(person.Name()))

		os.Exit(0)
	}

	fmt.Println("Unknown command.")

	os.Exit(0)
}
