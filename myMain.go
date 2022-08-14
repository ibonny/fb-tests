package main

import (
	"fmt"
	"io/ioutil"
	"os"

	people "fb-tests/v2/MyStruct"
	person "fb-tests/v2/MyStruct"

	"github.com/edsrzf/mmap-go"
	flatbuffers "github.com/google/flatbuffers/go"
)

func createRecord(builder *flatbuffers.Builder, firstName string, lastName string) flatbuffers.UOffsetT {
	fn := builder.CreateString(firstName)
	ln := builder.CreateString(lastName)

	person.PersonStart(builder)
	person.PersonAddFirstName(builder, fn)
	person.PersonAddLastName(builder, ln)
	temp := person.PersonEnd(builder)

	builder.Finish(temp)

	return temp
}

func main() {
	builder := flatbuffers.NewBuilder(0)

	// firstName := builder.CreateString("Ian")
	// lastName := builder.CreateString("Bonnycastle")
	// firstName2 := builder.CreateString("Shannan")
	// lastName2 := builder.CreateString("Stewart")

	// person.PersonStart(builder)
	// person.PersonAddFirstName(builder, firstName)
	// person.PersonAddLastName(builder, lastName)
	// ian := person.PersonEnd(builder)

	// builder.Finish(ian)

	// person.PersonStart(builder)
	// person.PersonAddFirstName(builder, firstName2)
	// person.PersonAddLastName(builder, lastName2)
	// shannan := person.PersonEnd(builder)

	// builder.Finish(ian)

	// people.PeopleStart(builder)

	// This has to be done in this order, or else you get an error.
	// panic: Incorrect creation order: object must not be nested.
	var thePeople []flatbuffers.UOffsetT

	thePeople = append(thePeople, createRecord(builder, "Shannan", "Stewart"))
	thePeople = append(thePeople, createRecord(builder, "Ian", "Bonnycastle"))

	people.PeopleStartPeopleVector(builder, 2)

	for _, person := range thePeople {
		builder.PrependUOffsetT(person)
	}

	// builder.PrependUOffsetT(createRecord(builder, "Shannan", "Stewart"))
	// builder.PrependUOffsetT(createRecord(builder, "Ian", "Bonnycastle"))
	peopleVec := builder.EndVector(2)

	people.PeopleStart(builder)
	people.PeopleAddPeople(builder, peopleVec)
	allPeople := people.PeopleEnd(builder)
	builder.Finish(allPeople)

	bytes := builder.FinishedBytes()

	err := ioutil.WriteFile("testout.dat", bytes, 0644)

	if err != nil {
		panic(err)
	}

	// newPerson := person.GetRootAsPerson(bytes, 0)

	// fmt.Println(string(newPerson.FirstName()))
	// fmt.Println(string(newPerson.LastName()))

	f, _ := os.OpenFile("testout.dat", os.O_RDWR, 0644)

	mmap, _ := mmap.Map(f, mmap.RDWR, 0)

	defer mmap.Unmap()

	newPeople := people.GetRootAsPeople(mmap, 0)

	fmt.Println(newPeople.PeopleLength())

	person := new(person.Person)

	// newPeople.People(person, 0)

	// fmt.Println(string(person.FirstName()))
	// fmt.Println(string(person.LastName()))

	for i := 0; i < newPeople.PeopleLength(); i++ {
		newPeople.People(person, i)

		fmt.Println(string(person.FirstName()))
		fmt.Println(string(person.LastName()))
	}
}
