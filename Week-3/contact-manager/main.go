// Import main package
package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Data structure to hold the data
type Contact struct{
	Name string `json:"name"`
	PhoneNumber int `json:"phonenumber"`
	Email string `json:"email"`
}

// Constructor for Contact data structure
func New(name string, phonenumber int, email string) Contact{
	return Contact{
		Name: name,
		PhoneNumber: phonenumber,
		Email: email,
	}
}


// Save contact
func saveContact(contact []interface{}) error{
	if len(contact) == 0{
		return errors.New("Contact is empty ... Please Enter details")
	}else {
		name := contact[0].(string)
		phonenumber := contact[1].(int)
		email := contact[2].(string)

		cont := New(name, phonenumber, email)
		jsonData , err := json.MarshalIndent(cont, ""," ")
		if err != nil{
			return errors.New("Error in saving the contact")
		}

		fmt.Println(jsonData)
		return nil
	}
}
// Add contact
func addContact(){
	var contact []interface{}

	var userName string
	fmt.Println("Enter User Name :")
	fmt.Scanf("%s", &userName)

	var phoneNumber int
	fmt.Println("Enter Phone Number :")
	fmt.Println("%d", &phoneNumber)

	var email string
	fmt.Println("Enter Email:")
	fmt.Scanf("%s", &email)

	contact = append(contact, userName)
	contact = append(contact, phoneNumber)
	contact = append(contact, email)

	err := saveContact(contact)

	if err != nil{
		fmt.Println("Failed to Save the contact .....\n Please try again")
	}else{
		fmt.Println("Contact Saved Successfully .....")
	}
}

// List contact
func displayContact(){

}

// Search contact
func searchContact(){

}

// Delete contact
func deleteContact(){

}


// Main function consiting of user input and menu driven logic
func main(){

	i :=0
	for i >= 0{
		fmt.Println("Contact Manager App :")
		var choice int

		fmt.Println("Enter your choice :")
		fmt.Println("Press 1 : Add Contact")
		fmt.Println("Press 2 : List Contact")
		fmt.Println("Press 3 : Search Contact")
		fmt.Println("Press 4 : Delete Contact")
		fmt.Println("Press 5 : Exit Application")
		fmt.Scanf("%d", &choice)

		if (choice == 1){
			addContact()
		}else if (choice ==2){
			displayContact()
		}else if (choice == 3){
			searchContact()
		}else if (choice == 4){
			deleteContact()
		}else if (choice == 5){
			fmt.Println("Bye! See you later")
			break
		}else{
			fmt.Println("Please enter a valid number from 1 - 5")
		}
		i++
	}
}