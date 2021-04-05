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
	"os"
	"os/exec"
	"math/rand"
	"time"
)


func main() {
	// Main flags that will determine what action will be triggered
	start := flag.Bool("start", false, "start a daemon")
	add := flag.Bool("add", false, "add new person")
	edit := flag.Int("edit", -1, "edit existing person")
	read := flag.Int("read", -1, "read data about existing person, by ID")
	search := flag.Bool("search", false, "search person by parameters")
	list := flag.Bool("list", false, "list all persons")
	remove := flag.Int("remove", -1, "remove existing person")
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
	trackBirthdate := flag.Int("trackbirthdate", -1, "")
	trackBirthdateShort := flag.Int("tb", -1, "")

	// Optional flags for categories
	categoryNew := flag.String("cnew", "", "add new category")
	categoryNewShort := flag.String("cn", "", "add new category")
	categoryEdit := flag.String("cedit", "", "edit existing category")
	categoryEditShort := flag.String("ce", "", "edit existing category")
	categoryRemove := flag.String("cremove", "", "remove a category")
	categoryRemoveShort := flag.String("cr", "", "remove a category")
	categoryList := flag.Bool("clist", false, "list all categories")
	categoryListShort := flag.Bool("cl", false, "list all categories")

	flag.Parse()

	if *start == true {
		cmd := exec.Command("./pcrmdaemon.exe")
		err := cmd.Start()
		if err != nil {
			fmt.Println("Could not start pcrmdaemon.exe!")
			os.Exit(-1)
		}
		os.Exit(1)
	}

	if *add == true && *edit < 0 && *read < 0 && *search == false &&
		*list == false && *remove < 0 && *categories == false {
			ActionAdd(*forename, *forenameShort, *surname, *surnameShort,
				*birthdate, *birthdateShort, *nickname, *nicknameShort,
				*trackBirthdate, *trackBirthdateShort)
	} else if *add == false && *edit >= 0 && *read < 0 && *search == false &&
		*list == false && *remove < 0 && *categories == false {
			ActionEdit(*edit, *forename, *forenameShort, *surname, *surnameShort,
				*birthdate, *birthdateShort, *nickname, *nicknameShort,
				*trackBirthdate, *trackBirthdateShort)
	} else if *add == false && *edit < 0 && *read >= 0 && *search == false &&
		*list == false && *remove < 0 && *categories == false {
			ActionRead(*read, *forename, *forenameShort, *surname, *surnameShort,
				*birthdate, *birthdateShort, *nickname, *nicknameShort,
				*trackBirthdate, *trackBirthdateShort)
	} else if *add == false && *edit < 0 && *read < 0 && *search == true &&
		*list == false && *remove < 0 && *categories == false {
			ActionSearch(*forename, *forenameShort, *surname, *surnameShort,
				*birthdate, *birthdateShort, *nickname, *nicknameShort,
				*trackBirthdate, *trackBirthdateShort)
	} else if *add == false && *edit < 0 && *read < 0 && *search == false &&
		*list == true && *remove < 0 && *categories == false {
			ActionList()
	} else if *add == false && *edit < 0 && *read < 0 && *search == false &&
		*list == false &&  *remove >= 0 && *categories == false {
			ActionRemove(*remove)
	} else if *add == false && *edit < 0 && *read < 0 && *search == false &&
		*list == false && *remove < 0 && *categories == true {
			ActionCategory(*categoryNew, *categoryNewShort, *categoryEdit, *categoryEditShort,
				*categoryRemove, *categoryRemoveShort, *categoryList, *categoryListShort)
	} else {
		fmt.Println("wrong parameters; you need to -add, OR -edit, OR -read, OR -search, OR -list, OR -remove, OR -category.")
	}
}


func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
