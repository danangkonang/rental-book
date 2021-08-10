package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/danangkonang/rental-book/db/migration"
	"github.com/danangkonang/rental-book/db/seeder"
)

type Tables struct {
	NameTable []string
}

var (
	MigrationFolder = "db/migration"
	SeederFolder    = "db/seeder"
)
func main() {
	if len(os.Args[1:]) == 0 {
		PrintHelper()
		return
	}
	runCmd()
}

func runCmd() {
	var t Tables
	switch os.Args[1] {
	case "run":
		migrationOrSeeder()
	case "reset":
		if os.Args[2] != "" {
			t.NameTable = strings.Split(os.Args[2], ",")
		}
		ResetTables(&t)
	case "drop":
		if os.Args[2] != "" {
			t.NameTable = strings.Split(os.Args[2], ",")
		}
		DropTables(&t)
	default:
		PrintHelper()
	}
}

func migrationOrSeeder() {
	var t Tables
	switch os.Args[2] {
	case "migration":
		if os.Args[3] == "" {
			RuningMigration(&t)
		} else {
			t.NameTable = strings.Split(os.Args[3], ",")
			RuningMigration(&t)
		}
	case "seeder":
		if os.Args[3] == "" {
			RuningSeeder(&t)
		} else {
			t.NameTable = strings.Split(os.Args[3], ",")
			RuningSeeder(&t)
		}
	default:
		PrintHelper()
	}
}

func RuningMigration(tbl *Tables) {
	files, err := ioutil.ReadDir(MigrationFolder)
	if err != nil {
		os.Exit(0)
	}
	if len(tbl.NameTable) == 0 {
		newFile := []string{}
		for _, file := range files {
			filename := file.Name()
			list := strings.Split(filename, "_migration_")
			if list[0] != "0.go" {
				name := list[1]
				tb_name := strings.Split(name, ".go")
				newFile = append(newFile, tb_name[0])
			}
		}
		tbl.NameTable = newFile
	}
	m := migration.Migration{}
	for _, migrate := range tbl.NameTable {
		meth := reflect.ValueOf(&m).MethodByName(strings.Title(migrate))
		meth.Call(nil)
	}
}

func RuningSeeder(tbl *Tables) {
	files, err := ioutil.ReadDir(SeederFolder)
	if err != nil {
		os.Exit(0)
	}
	if len(tbl.NameTable) == 0 {
		newFile := []string{}
		for _, file := range files {
			filename := file.Name()
			list := strings.Split(filename, "_seeder_")
			if list[0] != "0.go" {
				name := list[1]
				tb_name := strings.Split(name, ".go")
				newFile = append(newFile, tb_name[0])
			}
		}
		tbl.NameTable = newFile
	}
	s := seeder.Seeder{}
	for _, migrate := range tbl.NameTable {
		meth := reflect.ValueOf(&s).MethodByName(strings.Title(migrate))
		meth.Call(nil)
	}
}

type Gomig struct {
	Name    string
	Version string
}

var MyTemplate = `{{.Name}} Version {{.Version}}

Usage: {{.Name}} [command] [options]

Options:
	-v, --version                       output the version number
	-h, --help                          output usage information

Commands:
	- migrate
	- run
	- down

{{- /* end */ -}}
{{- "" }}`

var VersionTemplate = `{{.Name}} Version {{.Version}}
{{- /* end */ -}}
{{- "" }}
`

func PrintHelper() {
	data := Gomig{"Danang", "0.0.5"}
	tmpl, err := template.New("test").Parse(MyTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

func PrintVersion() {
	data := Gomig{"Danang", "0.0.5"}
	tmpl, err := template.New("test").Parse(MyTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

func DropTables(tb *Tables) {
	files, err := ioutil.ReadDir(MigrationFolder)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	if len(tb.NameTable) > 0 {
		for _, ntb := range tb.NameTable {
			query := "DROP TABLE IF EXISTS " + ntb + " CASCADE;"
			_, err := migration.Connection().Db.Exec(query)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			fmt.Println("success DROP TABLE " + ntb)
		}
	} else {
		for _, file := range files {
			filename := file.Name()
			if filename != "0.go" {
				names := strings.Split(filename, "_migration_")
				tb_name := strings.Split(names[1], ".go")
				query := "DROP TABLE IF EXISTS " + tb_name[0] + " CASCADE;"
				_, err := migration.Connection().Db.Exec(query)
				if err != nil {
					fmt.Println(err)
					os.Exit(0)
				}
				fmt.Println("success DROP TABLE " + tb_name[0])
			}
		}
	}
}

func ResetTables(tb *Tables) {
	files, err := ioutil.ReadDir(MigrationFolder)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	if len(tb.NameTable) > 0 {
		for _, ntb := range tb.NameTable {
			query := "TRUNCATE " + ntb + ";"
			_, err := migration.Connection().Db.Exec(query)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			fmt.Println("success delete row")
		}
	} else {
		for _, file := range files {
			filename := file.Name()
			if filename != "0.go" {
				names := strings.Split(filename, "_migration_")
				tb_name := strings.Split(names[1], ".go")
				query := "TRUNCATE " + tb_name[0] + ";"
				_, err := migration.Connection().Db.Exec(query)
				if err != nil {
					fmt.Println(err)
					os.Exit(0)
				}
				fmt.Println("success delete row")
			}
		}
	}
}
