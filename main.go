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
	"math/rand"
	"time"
)


func main() {
	// Main flags that will determine what action will be triggered
	add := flag.Bool("add", false, "add new person")
	edit := flag.Int("edit", -1, "edit existing person")
	read := flag.Int("read", -1, "read data about existing person, by ID")
	search := flag.Bool("search", false, "search person by parameters")
	categories := flag.Bool("category", false, "manage categories")

	// Optional flags for people
	forename := flag.String("forename", "", "")
	forenameShort := flag.String("fn", "", "")
	surname := flag.String("surname", "", "")
	surnameShort := flag.String("sn", "", "")
	birthdate := flag.String("birthdate", "", "")
	birthdateShort := flag.String("bd", "", "")
	nickname := flag.String("nickname", "", "")
	nicknameShort := flag.String("nn", "", "")

	// Optional flags for categories
	categoryNew := flag.String("cnew", "", "add new category")
	categoryNewShort := flag.String("cn", "", "add new category")
	categoryEdit := flag.String("cedit", "", "edit existing category")
	categoryEditShort := flag.String("ce", "", "edit existing category")

	flag.Parse()

	if *add == true && *edit < 0 && *read < 0 && *search == false && *categories == false {
		ActionAdd(*forename, *forenameShort, *surname, *surnameShort,
			*birthdate, *birthdateShort, *nickname, *nicknameShort)
	} else if *add == false && *edit >= 0 && *read < 0 && *search == false && *categories == false {
		ActionEdit(*edit, *forename, *forenameShort, *surname, *surnameShort,
			*birthdate, *birthdateShort, *nickname, *nicknameShort)
	} else if *add == false && *edit < 0 && *read >= 0 && *search == false && *categories == false {
		ActionRead(*read, *forename, *forenameShort, *surname, *surnameShort,
			*birthdate, *birthdateShort, *nickname, *nicknameShort)
	} else if *add == false && *edit < 0 && *read < 0 && *search == true && *categories == false {
		ActionSearch(*forename, *forenameShort, *surname, *surnameShort,
			*birthdate, *birthdateShort, *nickname, *nicknameShort)
	} else if *add == false && *edit < 0 && *read < 0 && *search == false && *categories == true {
		ActionCategory(*categoryNew, *categoryNewShort, *categoryEdit, *categoryEditShort)
	} else {
		fmt.Println("wrong parameters; you need to -add, OR -edit, OR -read, OR -search.")
	}
}


func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
