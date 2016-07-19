package main

import (
        "os"
	"io/ioutil"
        "strings"
        "flag"
        "fmt"
        "time"

        "github.com/chuckpreslar/inflect"
)

const (
        sample = `-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
-- FIXME add your script here
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back`
)

func main() {
        var output string
        var migrationName string
        flag.Usage = func() {
                fmt.Println("Usage: " + os.Args[0] + " [options]\n\nOptions:")
                fmt.Println("\t-o the output path.")
                fmt.Println("\t-n the migration name.")
        }
        flag.StringVar(&output, "o", "", "/path/to/foo/bar")
        flag.StringVar(&migrationName, "n", "", "AddNewColumnToTable[add_new_column_to_table]")
        flag.Parse()

        if migrationName == "" {
                flag.Usage()
                os.Exit(0)
        }

        migrationName = inflect.Underscore(migrationName)

        if output != "" && !strings.HasSuffix(output, "/") {
                output += "/"
        }

        f := fmt.Sprintf("%d_%s.sql",time.Now().Unix(), migrationName)

        err := ioutil.WriteFile(fmt.Sprintf("%s%s", output, f), []byte(sample), 0644)
        if err != nil {
                fmt.Println(err.Error())
                os.Exit(1)
        } else {
                fmt.Println("done.")
        }
}
