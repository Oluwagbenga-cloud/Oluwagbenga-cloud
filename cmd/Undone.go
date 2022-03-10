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

// undoneCmd represents the undone command
var undoneCmd = &cobra.Command{
	Use:   "undone",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("undone called")
	},
}

func init() {
	rootCmd.AddCommand(undoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// undoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// undoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func (c *Command) Undone(s string) {
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
		if i+1 == g {
			Support[i].Status = false
		}
	}
	c, _ := json.Marshal(Support)

	ioutil.WriteFile("todolist.csv", c, 0666)

	fmt.Println("It is a success")

}
