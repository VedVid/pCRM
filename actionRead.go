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


func ActionRead(id int, forename, fn, surname, sn, birthdate, bd, nickname, nn string, trackBirthdate, tb int) {
	if id < 0 {
		fmt.Println("ERROR: Wrong id passed.")
		os.Exit(-1)
	}

	if forename == "" && fn == "" && surname == "" && sn == "" &&
		birthdate == "" && bd == "" && nickname == "" && nn == "" &&
		trackBirthdate == -1 && tb == -1 {
		fmt.Println("INFO: Show all fields.")
		readPerson(id, true, "", "", "", "", -1)
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

	switch {
	case trackBirthdate < 0 && tb < 0:
		break
	case trackBirthdate == tb:
		fmt.Println("WARNING: You used two flags to set birthdate tracking.")
		break
	case trackBirthdate >= 0 && tb < 0:
		break
	case trackBirthdate < 0 && tb >= 0:
		trackBirthdate = tb
		break
	case (trackBirthdate == 1 && tb == 0) || (trackBirthdate == 0 && tb == 1):
		fmt.Println("ERROR: You used two flags to set birthdate tracking, each with different data.")
		os.Exit(-1)
	default:
		fmt.Println("ERROR: Something went wrong with setting up the birthdate tracking.")
		fmt.Println("       Please leave empty or -1 to left it in current state, 0 to disable, 1 to enable.")
		os.Exit(-1)
	}

	readPerson(id, false, forename, surname, birthdate, nickname, trackBirthdate)
}


func readPerson(id int, all bool, forename, surname, birthdate, nickname string, trackBirthdate int) {
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

	validID := false
	for _, v := range *persons {
		if id == v.Id {
			validID = true
			if all == true {
				fmt.Println("Id:             ", v.Id)
				fmt.Println("Forename:       ", v.Forename)
				fmt.Println("Surname:        ", v.Surname)
				fmt.Println("Birthdate:      ", v.Birthdate)
				fmt.Println("Nickname:       ", v.Nickname)
				fmt.Println("Track birthdate:", v.TrackBirthdate)
			} else {
				fmt.Println("Id:             ", v.Id)
				if forename != "" {
					fmt.Println("Forename:       ", v.Forename)
				}
				if surname != "" {
					fmt.Println("Surname:        ", v.Surname)
				}
				if birthdate != "" {
					fmt.Println("Birthdate:      ", v.Birthdate)
				}
				if nickname != "" {
					fmt.Println("Nickname:       ", v.Nickname)
				}
				if trackBirthdate == 1 {
					fmt.Println("Track birthdate:", v.TrackBirthdate)
				}
			}
			break
		}
	}

	if validID == false {
		fmt.Println("ERROR: ID", id, "not found.")
	}
}
