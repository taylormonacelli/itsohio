/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/taylormonacelli/itsohio/test2"
)

// test2Cmd represents the test2 command
var test2Cmd = &cobra.Command{
	Use:   "test2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := test2.Test2()
		if err != nil {
			fmt.Println("error running test2:", err)
		}
	},
}

var (
	userCount int
	batchSize int
)

func init() {
	rootCmd.AddCommand(test2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// test2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// test2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	test2Cmd.PersistentFlags().IntVar(&userCount, "user-count", 500_000, "number of users to insert")
	err := viper.BindPFlag("user-count", test2Cmd.PersistentFlags().Lookup("user-count"))
	if err != nil {
		fmt.Println("error binding user-count flag")
		os.Exit(1)
	}

	test2Cmd.PersistentFlags().IntVar(&batchSize, "batch-size", 1_000, "sqlite batch size")
	err = viper.BindPFlag("batch-size", test2Cmd.PersistentFlags().Lookup("batch-size"))
	if err != nil {
		fmt.Println("error binding batch-size flag")
		os.Exit(1)
	}
}