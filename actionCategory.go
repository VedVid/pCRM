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
	//"encoding/json"
	"flag"
	"fmt"
	//"io/ioutil"
	"os"
)


func ActionCategory(cnew, cn, cedit, ce string) {

	if cnew == "" && cn == "" && cedit == "" && ce == "" {
		fmt.Println("INFO: Switching to interactive mode.")
		interactiveCategory()
		os.Exit(1)
	}

	switch {
	case len(cnew) == 0 && len(cn) == 0:
		break
	case cnew == cn:
		fmt.Println("WARNING: You used two flags to set new category.")
		newCategory(cnew)
		os.Exit(1)
	case len(cnew) > 0 && len(cn) == 0:
		newCategory(cnew)
		os.Exit(1)
	case len(cnew) == 0 && len(cn) > 0:
		newCategory(cn)
		os.Exit(1)
	default:
		fmt.Println("ERROR: You used two flags to set new category, each with different data.")
		os.Exit(-1)
	}

	switch {
	case len(cedit) == 0 && len(ce) == 0:
		break
	case cedit == ce:
		fmt.Println("WARNING: You used two flags to edit category.")
		editCategory(cedit)
		os.Exit(1)
	case len(cedit) > 0 && len(ce) == 0:
		editCategory(cedit)
		os.Exit(1)
	case len(cedit) == 0 && len(ce) > 0:
		editCategory(ce)
		os.Exit(1)
	default:
		fmt.Println("ERROR: You used two flags to edit category, each with different data.")
		os.Exit(-1)
	}
}


func interactiveCategory() {
	fmt.Println("WERR: Interactive mode not implemented yet.")
}


func newCategory(cnew string) {
	fmt.Println(cnew)
}


func editCategory(cedit string) {
	if flag.NArg() != 1 {
		fmt.Println("ERROR: Invalid number of arguments:", flag.NArg(),
			"; should be 1.")
		os.Exit(-1)
	}
	fmt.Println("Old name:", cedit, "New name:", flag.Arg(len(flag.Args())-1))
}
