/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strconv"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		i := Command{}
		i.Add(args[0])
	},
}

func (c *Command) Done( s string) {
	g, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	readJson, err := ioutil.ReadFile("todolist.csv")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(readJson, &Support)

	for i, _ := range Support {
		if i + 1 == g{
			Support[i].Status = true
		}
	}
	c, err:= json.Marshal(Support)

	ioutil.WriteFile("todolist.csv", c, 0666)

	fmt.Println("It is a success")


}








func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

/*func (t *ListItem) Done(serlNo string) []ListItem {
//unmarshalCsv()
serialNo, err := strconv.Atoi(serlNo)
if err != nil {
	panic(err)
}
if serialNo != 0 {
	for i := range list {
		if i == (serialNo - 1) {
			list[i].Status = true
			fmt.Println("item status updated")
		}
	}
	t.savetoCsv()
	return list
}
fmt.Println("please input a valid serial no above 0")
return list*/
func unc (t *ListItem) Done(serlNo string) []ListItem {
//unmarshalCsv()
serialNo, err := strconv.Atoi(serlNo)
if err != nil {
panic(err)
}
if serialNo != 0 {
for i := range list {
if i == (serialNo - 1) {
list[i].Status = true
fmt.Println("item status updated")
}
}
t.savetoCsv()
return list
}
fmt.Println("please input a valid serial no above 0")
return list