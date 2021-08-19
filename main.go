package main

import "fmt"

type profile struct {
	name    string
	address string
}

type datastorers struct {
	id         int
	dataholder map[int]*profile
}

// a method in go is a function on a struct

// addStuff takes in attributes of a person and adds them to our datastore
func addStuff(id int, name2, address2 string, datastores datastorers) {
	datastores.dataholder[id] = &profile{
		name:    name2,
		address: address2,
	}
}

func (d datastorers) Add(id int, name2, address2 string) {
	d.dataholder[d.id] = &profile{
		name:    name2,
		address: address2,
	}
	d.id++
}

func (d datastorers) Remove(id int) {
	delete(d.dataholder, d.id)
}

//Removes by name
// Loops over map, if key is found and name is same then delete!
func (d datastorers) removeByName(name2 string) {
	for key := range d.dataholder {
		element, found := d.dataholder[key]
		if found == true && element.name == name2 {
			delete(d.dataholder, key)

		}

	}
}

func (d datastorers) Read(id int) {
	_, found := d.dataholder[d.id]
	if found == true {
		fmt.Println(d.dataholder[d.id])
	}

}

func (d datastorers) Update(id int, name2 string) {
	d.dataholder[d.id].name = name2
}

func main() {
	// numid := 1

	data1 := datastorers{
		dataholder: map[int]*profile{
			1: {
				name:    "Jess",
				address: "address1",
			},
		},
	}

	data1.Add(2, "Zade", "address2")
	data1.Read(2)
	data1.removeByName("Zade")
	// data1.Remove("2")
	data1.Read(2)
	// data1.Update(2, "Jim")
	// data1.Read(2)
	data1.Read(3)

	// fmt.Println(data1)

}
