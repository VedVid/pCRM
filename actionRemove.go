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


func ActionRemove(id int) {
	if id < 0 {
		fmt.Println("ERROR: Wrong id passed.")
		os.Exit(-1)
	}

	removePerson(id)
}


func removePerson(id int) {
	var err error
	var personsOld = &[]Person{}
	var personsNew = &[]Person{}

	f, err := ioutil.ReadFile("./data/people.json")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OK")
	}

	err = json.Unmarshal([]byte(f), personsOld)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OK")
	}

	validID := false
	for _, v := range *personsOld {
		if v.Id != id {
			*personsNew = append(*personsNew, v)
		} else {
			validID = true
		}
	}

	if validID == false {
		fmt.Println("ERROR: ID", id, "not found.")
		os.Exit(-1)
	}

	data, err := json.MarshalIndent(personsNew, "", "    ")
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
