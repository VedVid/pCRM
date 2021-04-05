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


func ActionSearch(forename, fn, surname, sn, birthdate, bd, nickname, nn string) {
	if forename == "" && fn == "" && surname == "" && sn == "" &&
		birthdate == "" && bd == "" && nickname == "" && nn == "" {
		fmt.Println("INFO: Switching to interactive mode.")
		interactiveSearch()
		os.Exit(1)
	}

	switch {
	case len(forename) == 0 && len(fn) == 0:
		break
	case forename == fn:
		fmt.Println("WARNING: You used two flags to mark forename.")
		break
	case len(forename) > 0 && len(fn) == 0:
		break
	case len(forename) == 0 && len(fn) > 0:
		forename = fn
		break
	default:
		break
	}

	switch {
	case len(surname) == 0 && len(sn) == 0:
		break
	case surname == sn:
		fmt.Println("WARNING: You used two flags to mark surname.")
		break
	case len(surname) > 0 && len(sn) == 0:
		break
	case len(surname) == 0 && len(sn) > 0:
		surname = sn
		break
	default:
		break
	}

	switch {
	case len(birthdate) == 0 && len(bd) == 0:
		break
	case birthdate == bd:
		fmt.Println("WARNING: You used two flags to mark birthdate.")
		break
	case len(birthdate) > 0 && len(bd) == 0:
		break
	case len(birthdate) == 0 && len(bd) > 0:
		birthdate = bd
		break
	default:
		break
	}

	switch {
	case len(nickname) == 0 && len(nn) == 0:
		break
	case nickname == nn:
		fmt.Println("WARNING: You used two flags to mark nickname.")
		break
	case len(nickname) > 0 && len(nn) == 0:
		break
	case len(nickname) == 0 && len(nn) > 0:
		nickname = nn
		break
	default:
		break
	}

	searchPerson(forename, surname, birthdate, nickname)
}


func interactiveSearch() {
	fmt.Println("WERR: Interactive mode not implemented yet.")
}


func searchPerson(forename, surname, birthdate, nickname string) {
	var err error
	var persons = &[]Person{}
	var tempPersons = []Person{}

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

	validForename := false
	validSurname := false
	validBirthdate := false
	validNickname := false
	for _, v := range *persons {
		if forename == "" {
			validForename = true
		}
		if forename != "" && forename == v.Forename {
			validForename = true
		}
		if surname == "" {
			validSurname = true
		}
		if surname != "" && surname == v.Surname {
			validSurname = true
		}
		if birthdate == "" {
			validBirthdate = true
		}
		if birthdate != "" && birthdate == v.Birthdate {
			validBirthdate = true
		}
		if nickname == "" {
			validNickname = true
		}
		if nickname != "" && nickname == v.Nickname {
			validNickname = true
		}
		if validForename == true && validSurname == true &&
			validBirthdate == true && validNickname == true {
				tempPersons = append(tempPersons, v)
		}
	}

	for _, v := range tempPersons {
		fmt.Println("Id:       ", v.Id)
		fmt.Println("Forename: ", v.Forename)
		fmt.Println("Surname:  ", v.Surname)
		fmt.Println("Birthdate:", v.Birthdate)
		fmt.Println("Nickname: ", v.Nickname)
	}
}
