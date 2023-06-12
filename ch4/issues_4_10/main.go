// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

// !+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	issues := make(map[string][]*github.Issue)
	issues["less than a month old"] = make([]*github.Issue, 0)
	issues["less than a year old"] = make([]*github.Issue, 0)
	issues["older than a year"] = make([]*github.Issue, 0)
	now := time.Now()
	month := now.Add(-30 * 24 * time.Hour)
	year := now.Add(-365 * 24 * time.Hour)
	for _, item := range result.Items {
		if item.CreatedAt.After(month) {
			issues["less than a month old"] = append(issues["less than a month old"], item)
		} else if item.CreatedAt.Before(month) && item.CreatedAt.After(year) {
			issues["less than a year old"] = append(issues["less than a year old"], item)
		} else {
			issues["older than a year"] = append(issues["older than a year"], item)
		}
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Print("\n[Issues less than a month old]\n")
	for _, item := range issues["less than a month old"] {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Print("\n[Issues less than a year old]\n")
	for _, item := range issues["less than a year old"] {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Print("\n[Issues older than a year]\n")
	for _, item := range issues["older than a year"] {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

//!-

/*
//!+textoutput
$ go build github.com/KipsasJaujoj/gopl.io/ch4/issues_4_10
$ ./issues_4_10 repo:golang/go is:open json decoder
82 issues:

[Issues less than a month old]
#60497  VirrageS x/vuln: make progress-like output opt-in for -json flag

[Issues less than a year old]
#56733 rolandsho encoding/json: add (*Decoder).SetLimit
#59053   joerdav proposal: encoding/json: add a generic Decode function
#58649 nabokihms encoding/json: show nested fields path if DisallowUnkno
#56332    gansvv encoding/json: clearer error message for boolean like p
#58655    Fazt01 proposal: encoding/json: wrap error from TextUnmarshale

[Issues older than a year]
#48298     dsnet encoding/json: add Decoder.DisallowDuplicateFields
#29035    jaswdr proposal: encoding/json: add error var to compare  the
#36225     dsnet encoding/json: the Decoder.Decode API lends itself to m
#42571     dsnet encoding/json: clarify Decoder.InputOffset semantics
#11046     kurin encoding/json: Decoder internally buffers full input
#43716 ggaaooppe encoding/json: increment byte counter when using decode
#34543  maxatome encoding/json: Unmarshal & json.(*Decoder).Token report
#43513 Alexander encoding/json: add line number to SyntaxError
#32779       rsc encoding/json: memoize strings during decode
#48950 Alexander encoding/json: calculate correct SyntaxError.Offset in
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#31701    lr1980 encoding/json: second decode after error impossible
#29750  jacoelho cmd/vet: stdmethods check gets confused if run on a pac
#40128  rogpeppe proposal: encoding/json: garbage-free reading of tokens
#40982   Segflow encoding/json: use different error type for unknown fie
#16212 josharian encoding/json: do all reflect work before decoding
#40127  rogpeppe encoding/json: add Encoder.EncodeToken method
#5901        rsc encoding/json: allow per-Encoder/per-Decoder registrati
#41144 alvaroale encoding/json: Unmarshaler breaks DisallowUnknownFields
#33854     Qhesz encoding/json: unmarshal option to treat omitted fields
#14750 cyberphon encoding/json: parser ignores the case of member names
#34564  mdempsky go/internal/gcimporter: single source of truth for deco
#26946    deuill encoding/json: clarify what happens when unmarshaling i
#19858 mrajashre proposal: encoding/json: add mechanism to mark fields a
//!-textoutput
*/
