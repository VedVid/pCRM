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
	"flag"
	"fmt"
)


func main() {
	// Main flags that will determine what action will be triggered
	add := flag.Bool("add", false, "add new person")
	edit := flag.Bool("edit", false, "edit existing person")
	read := flag.Bool("read", false, "read data about existing person")
	search := flag.Bool("search", false, "search person by parameters")

	// Optional flags
	forename := flag.String("forename", "", "")
	forenameShort := flag.String("fn", "", "")
	surname := flag.String("surname", "", "")
	surnameShort := flag.String("sn", "", "")
	birthdate := flag.String("birthdate", "", "")
	birthdateShort := flag.String("bd", "", "")

	flag.Parse()

	if *add == true && *edit == false && *read == false && *search == false {
		ActionAdd(*forename, *forenameShort, *surname, *surnameShort, *birthdate, *birthdateShort)
	} else if *add == false && *edit == true && *read == false && *search == false {
		fmt.Println("editing a person...")
	} else if *add == false && *edit == false && *read == true && *search == false {
		fmt.Println("reading about a person...")
	} else if *add == false && *edit == false && *read == false && *search == true {
		fmt.Println("searching for a person...")
	} else {
		fmt.Println("wrong parameters; you need to -add, OR -edit, OR -read.")
	}
}
