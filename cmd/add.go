/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"io/ioutil"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
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

type Command struct {
	List   string
	Status bool
}

var Support []Command

func (c *Command) Add(s string) {
	a := Command{List: s, Status: false}

	red, _ := ioutil.ReadFile("todolist.csv")
	json.Unmarshal(red, &Support)

	Support = append(Support, a)

	json_red, _ := json.Marshal(Support)

	ioutil.WriteFile("todolist.csv", json_red, 0666)

}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*func WriteToFile(item string) {
f, err := os.Create("todolist.csv")
if err != nil {
	panic(err)
}
val, err := f.WriteString(item)
if err != nil {
	fmt.Println(err)
	f.Close()
	return
}
fmt.Println(val, "added to list")
err = f.Close()
if err != nil {
	fmt.Println(err)
	return
}*/
