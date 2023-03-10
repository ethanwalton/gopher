/*
Copyright © 2023 Ethan Walton
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "This command will get the desired gopher wantes",
	Long:  `This command under the hood will call another repo to return the gopher.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
		var gopherName = "gophygoph"

		if len(args) >= 1 && args[0] != "" {
			gopherName = args[0]
		}

		URL := "https://github.com/scraly/gophers/raw/main/" + gopherName + ".png"

		fmt.Println("Trying to get '" + gopherName + "' Gopher....")

		//getting the gopher data now
		response, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
		}

		defer response.Body.Close()

		if response.StatusCode == 200 {
			out, err := os.Create(gopherName + ".png")
			if err != nil {
				fmt.Println(err)
			}
			defer out.Close()

			//write the body to file
			_, err = io.Copy(out, response.Body)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Perfect! Just saved in" + out.Name() + "!")
		} else {
			fmt.Println("Error: " + gopherName + " does not exists! :-(")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
