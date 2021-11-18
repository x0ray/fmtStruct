package main

// test fmtStruct() structure printing function
/* Example output:
Sample Movie Data:
[000]  &main.Lib{
[001]    Movies:[]main.Movie{
[002]      main.Movie{
[003]        Name:"Moby Dick",
[004]        UPC:2761686294,
[005]        ISBN:0-7928-5014-9,
[006]        pubYear:1956,
[007]        director:"John Huston",
[008]        cost:19.95,
[009]        copies:3
[010]      },
[011]      main.Movie{
[012]        Name:"Firewall",
[013]        UPC:1256959410,
[014]        ISBN:1-4198-0220-8,
[015]        pubYear:2006,
[016]        director:"Richard Loncraine",
[017]        cost:10.45,
[018]        copies:1
[019]      },
[020]      main.Movie{
[021]        Name:"The Exorcist",
[022]        UPC:8539186322,
[023]        ISBN:0-7907-5167-4,
[024]        pubYear:1973,
[025]        director:"William Friedkin",
[026]        cost:12.98,
[027]        copies:1
[028]      }
[029]    },
[030]    totalCount:5,
[031]    maxIncome:54.32
[032]  }
*/

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"unsafe"
)

// fmtStruct - formats any variable / struct in a readable format
func fmtStruct(val interface{}, opts ...string) string {
	lineNumForm := "[%04d] "
	title := "Structure dump. . ."
	switch len(opts) {
	case 1:
		title = opts[0]
	case 2:
		title = opts[0]
		lineNumForm = opts[1]
	}
	rstr := fmt.Sprintf("\n%s:\n", title)
	str := fmt.Sprintf("%#v", val)

	fmt.Println(str)

	// handy regexp creation tool: https://regexr.com/
	re := regexp.MustCompile(`{wall:[0-9a-f]x([0-9a-f]+),\sext:([0-9]+),\sloc:\(\*time.Location\)\([0-9a-f]x([0-9a-f]+)\)}`)
	repr := func(in string) string {
		fmt.Println(in)
		sl := re.FindAllStringSubmatch(in, -1)
		fmt.Printf("%q\n", sl)

		var wall uint64
		fmt.Sscanf(sl[0][1], "%x", &wall)
		var ext int64
		fmt.Sscanf(sl[0][2], "%d", &ext)

		var ploc unsafe.Pointer
		fmt.Sscanf(sl[0][3], "%x", &ploc)
		var loc *time.Location
		loc = (*time.Location)(ploc)

		fmt.Printf("wall: 0x%x ext: %d ploc: 0x%x loc: %v\n", wall, ext, ploc, loc)
		return in
	}
	fmt.Println(re.ReplaceAllStringFunc(str, repr))

	fmt.Println(str)

	// insert some split characters (which are not in the data!)
	str = strings.Replace(str, ",", ",^", -1)
	str = strings.Replace(str, "{}", "~", -1) // hide {} pairs
	str = strings.Replace(str, "{", "{^", -1)
	str = strings.Replace(str, "}", "^}", -1)
	// split up the long string
	strs := strings.Split(str, "^")
	pad := 1
	for i, v := range strs {
		cl := strings.Count(v, "}") * 2
		pad = pad - cl
		out := strings.Trim(v, " ")
		out = strings.Replace(out, "~", "{}", -1) // put {} pairs back
		rstr = rstr + fmt.Sprintf(lineNumForm+"%s%s\n", i, strings.Repeat(" ", pad), out)
		op := strings.Count(v, "{") * 2
		pad = pad + op
	}
	return rstr
}

type Lib struct {
	Movies     []Movie
	totalCount int
	maxIncome  float64
}

type Isbn [13]byte

func (i Isbn) GoString() string { return string(i[:]) }
func toIsbn(s string) Isbn      { var i Isbn; copy(i[:], s[:13]); return i }

type Year int16

func (y Year) GoString() string { return string(y) }

type DateTime time.Time

func (t *DateTime) String() string   { return time.Time(*t).Format(time.RFC3339) }
func (t *DateTime) GoString() string { return time.Time(*t).String() }

type Movie struct {
	Name     string
	UPC      int64
	ISBN     Isbn
	pubYear  Year
	director string
	cost     float32
	viewed   DateTime
	copies   int
}

func main() {
	simple := struct {
		count int
		start time.Time
		size  int8
		name  string
		part  uint8
		cost  float64
	}{
		count: 45,
		start: time.Now(),
		size:  29,
		name:  "Billy",
		part:  64,
		cost:  28.95,
	}
	// print out simple after formatting with fmtStruct()
	fmt.Print(fmtStruct(simple))

	data := Lib{
		Movies: []Movie{
			{Name: "Moby Dick", UPC: 2761686294, ISBN: toIsbn("0-7928-5014-9"),
				pubYear: 1956, director: "John Huston", cost: 19.95, copies: 3},
			{Name: "Firewall", UPC: 1256959410, ISBN: toIsbn("1-4198-0220-8"),
				pubYear: 2006, director: "Richard Loncraine", cost: 10.45, copies: 1},
			{Name: "The Exorcist", UPC: 8539186322, ISBN: toIsbn("0-7907-5167-4"),
				pubYear: 1973, director: "William Friedkin", cost: 12.98, copies: 1},
		},
		totalCount: 0,
		maxIncome:  54.32,
	}
	for i, v := range data.Movies {
		data.Movies[i].viewed = DateTime(time.Now())
		data.totalCount += v.copies
	}
	adata := &data

	// print out adata after formatting with fmtStruct()
	fmt.Print(fmtStruct(adata))
}
