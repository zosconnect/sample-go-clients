// (C) Copyright IBM Corp. 2021 All Rights Reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// This sample shows how to create an orchestration API using Go.
// This program can be executed on any platform that supports Go including
// z/OS.
//
// This sample uses the phonebook application (IVTNO) provided with IMS.
// The phonebook application can add a contact, delete a contact, display
// a contact and update the contact. The data is stored in an IMS database. 
// In this sample, we will use the 'display a contact' function to list
// information about the contact based on last name. This IMS function can be 
// called as a REST API using z/OS Connect Enterprise Edition. The following
// fields are returned:
//
//	* Last Name
//	* First Name
//	* Extension Number
//	* Zip Code
//
// After retrieving the contact information, it will extract additional
// information related to the zip code of the contact using a postal code 
// API provided by geocoder.ca
//
// The final JSON result will include additional information returned by
// the geocoder.ca postal code API like:
//
//	* City
//	* Province
//	* Longitude
//	* Latitude
//

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"os"
)

// struct that will be used to read names from input.json
type Input struct {
	Names	[]string	`json:"names"`	
}

// the next 2 structs will be used to hold the response from the IMS REST API
type FirstResponse struct {
	Output		OutputArea	`json:"OUTPUT_AREA"`	
}

type OutputArea struct {
	ZipCode		string	`json:"OUT_ZIP_CODE"`
	FirstName	string	`json:"OUT_FIRST_NAME"`
	LastName	string	`json:"OUT_LAST_NAME"`
	Extension	string	`json:"OUT_EXTENSION"`
	Message		string	`json:"OUT_MESSAGE"`
}

// the next 2 structs will be used to hold response from the geocoder.ca REST API
type SecondResponse struct {
	PlaceData	Place	`json:"standard"`
	Longitude	string	`json:"longt"`
	Latitude	string	`json:"latt"`
}

type Place struct {
	Province	string	`json:"prov"`
	City		string	`json:"city"`
}

// this struct will be used to construct the final output
type Result struct {
	lastname	string
	firstname	string
	extension	string
	zipcode		string
	latitude	string
	longitude	string
	province	string
	city		string
}

//
// This function retrieves the contact information by calling 2 REST APIs
// - API to get phonebook contact from IMS
// - API to get postal code or zip code information
// and combines the results into a single JSON object
//
func getContactInfo(lastName string) {
	//
	// The REST endpoint below is using the IBM Cloud secure gateway service to enable
	// access to a private network where the original REST API is located
	//
	reqUrl := "http://cap-sg-prd-4.securegateway.appdomain.cloud:20522/ims/phonebook/contact?lastname=" + lastName
	response, err := http.Get(reqUrl)

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	// Handling the response from the IMS Phonebook REST API
	responseData, err := ioutil.ReadAll(response.Body)
	response.Body.Close()	// important to do this once done reading

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	var phoneBook FirstResponse
	json.Unmarshal(responseData, &phoneBook)

	postal := phoneBook.Output.ZipCode

	//
	// Check http://geocoder.ca for details on the postal or zip code
	// REST API used in this program
	//
	addrUrl := "http://geocoder.ca/?locate=" + postal + "&geoit=XML&json=1"
	
	response2, err := http.Get(addrUrl)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	// Handling the response from the geocoder REST API
	responseData2, err := ioutil.ReadAll(response2.Body)
	response2.Body.Close()
	
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	var addr SecondResponse
	json.Unmarshal(responseData2, &addr)
	
	//
	// Check if phonebook response was not empty and if so, build
	// resulting object and populate fields with data before
	// returning the final object
	//
	if (phoneBook.Output.Message != "SPECIFIED PERSON WAS NOT FOUND") {
		var result Result
		
		result.lastname = phoneBook.Output.LastName
		result.firstname = phoneBook.Output.FirstName
		result.extension = phoneBook.Output.Extension
		result.zipcode = phoneBook.Output.ZipCode
		result.latitude = addr.Latitude
		result.longitude = addr.Longitude
		result.province = addr.PlaceData.Province
		result.city = addr.PlaceData.City

		fmt.Printf("%s\n", result.lastname)
		fmt.Printf("%s\n", result.firstname)
		fmt.Printf("%s\n", result.extension)
		fmt.Printf("%s\n", result.zipcode)
		fmt.Printf("%s\n", result.latitude)
		fmt.Printf("%s\n", result.longitude)
		fmt.Printf("%s\n", result.province)
		fmt.Printf("%s\n", result.city)

	} else {
		fmt.Printf("Contact record not found\n")
		os.Exit(1)
	}

}


//
// Opens input.json file, populates struct with its data and then iterates over names
//
func main() {

	file, err := os.Open("input.json")
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	var usrInput Input
	json.Unmarshal(data, &usrInput)

	for i := 0; i < len(usrInput.Names); i++ {
		getContactInfo(usrInput.Names[i])
	}

	file.Close()
}
