package main

import (
	"fmt"

	"struct-basics/structs"
)

type IntGrades struct {
	visualEffects int
	soundMixing   int
	costumDesign  int
	direction     int
}

func (g IntGrades) getSumGrades() int {
	return g.visualEffects + g.soundMixing + g.costumDesign + g.direction
}

func main() {
	type Movie struct {
		name, director  string
		durationMinutes float32
		good            bool
	}

	catchMeIfYouCan := Movie{
		name:            "Catch me if you can",
		director:        "Steven Spielberg",
		durationMinutes: 141,
		good:            true,
	}

	// another way of initialize the struct, this way you need to provide all fields
	// values in order
	greenbook := Movie{"Greenbook", "Peter Farrelly", 124, true}

	// you can skip some fields of a struct
	django := Movie{
		name:     "Django",
		director: "Tarantino",
		good:     true,
	}

	// anonymous struct
	ironMan3 := struct {
		name, director  string
		durationMinutes float32
		good            bool
	}{
		name:            "Iron Man 3",
		director:        "Shane Black",
		durationMinutes: 130,
		good:            false,
	}

	// pointer to a struct
	_21 := &Movie{
		name:            "21",
		director:        "Robert Luketic",
		durationMinutes: 123,
		good:            true,
	}
	fmt.Println("------------------- Part 1")
	fmt.Println("normal:", catchMeIfYouCan)
	fmt.Println("another way:", greenbook)
	fmt.Println("skip fields:", django)
	fmt.Println("anonymous:", ironMan3)
	fmt.Println("pointer to a struct:", "name", (_21.name), "good movie", (_21).good)

	// anonymous fields
	// you can define a struct type without declaring any field names
	type Data struct {
		string
		int
		bool
	}
	sample1 := Data{"John Locke", 50, true}

	// nested struct
	type Grades struct {
		visualEffects int
		soundMixing   int
		costumDesign  int
		direction     int
	}

	type NestedMovie struct {
		name, director  string
		durationMinutes float32
		good            bool
		grades          Grades
	}

	inception := NestedMovie{
		name:            "Inception",
		director:        "Christopher Nolan",
		durationMinutes: 155,
		good:            true,
		grades:          Grades{10, 9, 9, 10},
	}

	// anonymous nested struct
	type AnoGrades struct {
		visualEffects int
		soundMixing   int
		costumDesign  int
		direction     int
	}

	type AnoNestedMovie struct {
		name, director  string
		durationMinutes float32
		good            bool
		AnoGrades
	}

	_365dni := AnoNestedMovie{
		name:      "365 dni",
		director:  "Barbara Białowąs",
		AnoGrades: AnoGrades{5, 5, 5, 5},
	}

	// when anonymous nested struct is used, all the nested struct fields are automatically
	// available on the parent struct. this is called
	//FIELD PROMOTION //
	_365dni.costumDesign = 7
	_365dni.direction = 6
	_365dni.durationMinutes = 120

	fmt.Println("------------------ Part 2")
	fmt.Println("anonymous fields:", sample1)
	fmt.Println("nested struct:", inception)
	fmt.Println("nested struct access direction grade field:", inception.grades.direction)
	fmt.Println("anonymous nested struct:", _365dni)
	fmt.Println("field promotions:", _365dni)

	///// PART 3: INTERFACES
	type SumGrades interface {
		getSumGrades() int
	}

	type IntNestedMovie struct {
		name, director  string
		durationMinutes float32
		good            bool
		grades          IntGrades
	}

	// this func and the intGrades need to be outside Main func

	fiftyShadesOfGrey := IntNestedMovie{
		name:            "Fifty Shades of Grey",
		director:        "Sam Taylor-Johnson & James Foley",
		durationMinutes: 140,
		grades:          IntGrades{5, 5, 5, 5},
	}

	// anonymous interface
	type AnoIntNestedMovie struct {
		name, director  string
		durationMinutes float32
		good            bool
		IntGrades
	}

	blackPanther := AnoIntNestedMovie{
		name:            "Black Panther",
		director:        "Ryan Coogler",
		durationMinutes: 135,
		good:            true,
		IntGrades:       IntGrades{10, 9, 10, 10},
	}

	fmt.Println("------------------ Part 3: Interfaces")
	fmt.Println("methods nested interface:", fiftyShadesOfGrey)
	fmt.Println("methods nested intercafe/sum of grades:", fiftyShadesOfGrey.grades.getSumGrades())
	fmt.Println("methods anonymous nested interface:", blackPanther)
	fmt.Println("methods anonymous nested interface/sum of grades:", blackPanther.IntGrades.getSumGrades())

	// part 4: imported struct

	backToTheFuture := structs.ImpMovie{
		Name:            "Back to The Future",
		Director:        "Robert Zemeckis",
		DurationMinutes: 140,
		Good:            true,
	}

	jack := structs.ImpPerson{
		FirstName: "Jack",
		LastName:  "Shephard",
		FullName: func(firstName string, lastName string) string {
			return firstName + " " + lastName
		},
	}

	fmt.Println("------------------ Part 4")
	fmt.Println("imported struct:", backToTheFuture)
	fmt.Println("imported struct with function:", jack.FullName(jack.FirstName, jack.LastName))
}
