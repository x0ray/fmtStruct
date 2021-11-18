# func fmtStruct()
Create readable multiline formatted dump of any Golang structure using the fmtStruct() func.

## Introduction

When debugging a go program it can be handy to dump an entire data structure with as much detail as possible. This program demonstrates a function that provides detailed output in a string, so it is easy to log, print or even display on a web page. Here is some example output:

``` 
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
```

## How it works

The standard Golang fmt package does the heavy lifting, extracting the details of the data structure, and formating it using the Go-syntax representation of the value. This is then post processed to add new line characters, line numbers and indentation. Because the go-syntax formatting is used it is possible to add GoString stringers to user defined types to improve readability even further, as demonstrated in the main program. 

## Usage

The following go program excerpt shows the use of fmtStruct():

``` go	data := Lib{
	simple := struct{
		count int
		name string
		cost floar64
	}{
		count: 45
		name: "Billy"
		cost "28.95"
	}
	// print out simple after formatting with fmtStruct()
	fmt.Print(fmtStruct(adata))
```
This produces the following output:
```
Structure dump. . .:
[0000]  struct {
[0001]    count int; name string; cost float64
[0002]  }{
[0003]    count:45,
[0004]    name:"Billy",
[0005]    cost:28.95
[0006]  }
```

The demonstration program shows this example and a more complex example which demonstrates the use of user defined GoStringers to enhance readability even further.
