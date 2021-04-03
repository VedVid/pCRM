/*
Copyright (c) 2021, Tomasz "VedVid" Nowakowski
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/


package main


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)


func ActionAdd(forename, fn, surname, sn, birthdate, bd string) {

	if forename == "" && fn == "" && surname == "" && sn == "" &&
		birthdate == "" && bd == "" {
		fmt.Println("INFO: Switching to interactive mode.")
		interactiveAdd()
		os.Exit(1)
	}

	switch {
	case len(forename) == 0 && len(fn) == 0:
		break
	case forename == fn:
		fmt.Println("WARNING: You used two flags to set forename.")
		break
	case len(forename) > 0 && len(fn) == 0:
		fmt.Println("INFO: Proper data entered.")
		break
	case len(forename) == 0 && len(fn) > 0:
		fmt.Println("INFO: Proper data entered.")
		forename = fn
		break
	default:
		fmt.Println("ERROR: You used two flags to set forename, each with different data.")
		os.Exit(-1)
	}

	switch {
	case len(surname) == 0 && len(sn) == 0:
		break
	case surname == sn:
		fmt.Println("WARNING: You used two flags to set surname.")
		break
	case len(surname) > 0 && len(sn) == 0:
		fmt.Println("INFO: Proper data entered.")
		break
	case len(surname) == 0 && len(sn) > 0:
		fmt.Println("INFO: Proper data entered.")
		surname = sn
		break
	default:
		fmt.Println("ERROR: You used two flags to set surname, each with different data.")
		os.Exit(-1)
	}

	switch {
	case len(birthdate) == 0 && len(bd) == 0:
		break
	case birthdate == bd:
		fmt.Println("WARNING: You used two flags to set birthdate.")
		break
	case len(birthdate) > 0 && len(bd) == 0:
		fmt.Println("INFO: Proper data entered.")
		break
	case len(birthdate) == 0 && len(bd) > 0:
		fmt.Println("INFO: Proper data entered.")
		birthdate = bd
		break
	default:
		fmt.Println("ERROR: You used two flags to set birthdate, each with different data.")
		os.Exit(-1)
	}

	addPerson(forename, surname, birthdate)
}


func interactiveAdd() {
	fmt.Println("WERR: Interactive mode not implemented yet.")
}


func addPerson(forename, surname, birthdate string) {
	var err error
	var persons = &[]Person{}

	f, err := ioutil.ReadFile("./data/people.json")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OK")
	}

	err = json.Unmarshal([]byte(f), persons)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OK")
	}

	newPerson := &Person{forename, surname, birthdate}
	*persons = append(*persons, *newPerson)
	fmt.Println(persons)

	data, err := json.MarshalIndent(persons, "", "    ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OK")
	}

	err = ioutil.WriteFile("./data/people.json", data, 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OK")
	}
}
