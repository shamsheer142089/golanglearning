package main

import "fmt"

//Normal Struct
type person struct {
	firstName string
	lastName  string
	age       int
	gender    string
}

//Struct within Struct
type personwithcontact struct {
	firstName string
	lastName  string
	age       int
	gender    string
	contacts  contactinfo
}

//embedded struct without variable name
type personwithoutnameforinternalstruct struct {
	firstName   string
	lastName    string
	age         int
	gender      string
	contactinfo //this is same as personwithcontact
}

func main() {
	//Assigning struct values - method one
	shamsheer := person{"Mohammed Shamsheer", "Shaik", 30, "M"}
	//Assigning struct values - method 2
	var vihaan person
	vihaan.firstName = "Vihaan"
	fmt.Println(vihaan)
	fmt.Println(shamsheer)
	//getting individual values from struct
	fmt.Println(shamsheer.firstName)
	fmt.Println(shamsheer.lastName)
	fmt.Println(shamsheer.age)
	fmt.Println(shamsheer.gender)
	//assigning struct values for embedded struct
	shamsheerWithContactDetails := personwithcontact{
		"Shamsheer",
		"Shaik",
		30,
		"M",
		contactinfo{"sham@gmail.com", "123456678", "9876543s"},
	}
	fmt.Println(shamsheerWithContactDetails)

	shamsheerWithoutVariableNameForContact := personwithoutnameforinternalstruct{
		"Shamsheer",
		"Shaik",
		30,
		"M",
		contactinfo{"sham@gmail.com", "123456678", "9876543s"},
	}

	fmt.Println(shamsheerWithoutVariableNameForContact)
	//struct with receiver
	shamsheer.printPerson()

	//pointers
	structPointer:=&shamsheer
	fmt.Println(structPointer)
	structPointer.updatePersonFirstName("S Md Shamsheer")



}

//receiver and struct
func (per person) printPerson() {
	fmt.Println(per)

}

func(per *person) updatePersonFirstName(firstName string){
	(*per).firstName=firstName
	 fmt.Println(&per)


}
