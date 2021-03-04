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
// This program shows how Go can be used as a lightweight rules engine.
//
// This sample rule handles insurance claims. The following rules are used
// when processing a claim:
//
//	Drug claim - amount execeeded claim limit of $1000
//	Dental claim - amount exceeded claim limit of $800
//	Medical claim - amount exceeded claim limit of $500
//

package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

// this struct will be used to read from input.json
type ClaimDetails struct {
	ClaimType	[]string	`json:"types"`
	MedClaims	[]string	`json:"medAmounts"`
	DentalClaims	[]string	`json:"dentalAmounts"`
	DrugClaims	[]string	`json:"drugAmounts"`
}

// this struct will be used to form the final result object
type ClaimResult struct {
	Type	string
	Amount	float64
	Status	string
	Reason	string
}

//
// This function gets the result of the claim based on the claim type
// and claim amount. The following fields are passed as query parameters
// - claimType
// - claimAmount
//
func getClaimResult(claimType string, claimAmount float64) {
	
	var results ClaimResult
	
	results.Type = claimType
	results.Amount = claimAmount

	// switch statement to handle the different claim types
	switch claimType{
	case "MEDICAL":
		if claimAmount > 100.00 {
			results.Status = "Rejected"
			results.Reason = "Amount exceeded $100. Claim requires further review"
		} else {
			results.Status = "Accepted"
			results.Reason = "Normal claim"
		}
	case "DENTAL":
		if claimAmount > 800.00 {
			results.Status = "Rejected"
			results.Reason = "Amount exceeded $800. Claim requires further review"
		} else {
			results.Status = "Accepted"
			results.Reason = "Normal claim"
		}
	case "DRUG":
		if claimAmount > 1000.00 {
			results.Status = "Rejected"
			results.Reason = "Amount exceeded $1000. Claim requires further review"
		} else {
			results.Status = "Accepted"
			results.Reason = "Normal claim"
		}
	default:
		fmt.Printf("invalid claim type specified. Please enter a valid claim type\n")
		os.Exit(1)
	}

	fmt.Printf("%s\n", results.Type)
	fmt.Printf("%f\n", results.Amount)
	fmt.Printf("%s\n", results.Status)
	fmt.Printf("%s\n", results.Reason)
}

//
// populates struct with data from input.json and then reads from resulting struct
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

	var claim ClaimDetails
	json.Unmarshal(data, &claim)

	// iterator for each claim type
	for i := 0; i < len(claim.ClaimType); i++ {
		if claim.ClaimType[i] == "MEDICAL" {
			for j := 0; j < len(claim.MedClaims); j++ {
				amount, err := strconv.ParseFloat(claim.MedClaims[j], 64)
				if err != nil {
					fmt.Printf("%v\n", err)
					os.Exit(1)
				}
				getClaimResult(claim.ClaimType[i], amount)
			}
		} else if claim.ClaimType[i] == "DENTAL" {
			for j := 0; j < len(claim.DentalClaims); j++ {
				amount, err := strconv.ParseFloat(claim.DentalClaims[j], 64)
				if err != nil {
					fmt.Printf("%v\n", err)
					os.Exit(1)
				}
				getClaimResult(claim.ClaimType[i], amount)
			}
		} else if claim.ClaimType[i] == "DRUG" {
			for j := 0; j < len(claim.DrugClaims); j++ {
				amount, err := strconv.ParseFloat(claim.DrugClaims[j], 64)
				if err != nil {
					fmt.Printf("%v\n", err)
					os.Exit(1)
				}
				getClaimResult(claim.ClaimType[i], amount)
			}
		}
	}

	file.Close()
}
