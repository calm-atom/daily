/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// weatherCmd represents the weather command
var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: weatherData,
}

func init() {
	rootCmd.AddCommand(weatherCmd)
	weatherCmd.Flags().StringP("key", "k", "none", "Open weather api key")

	//weatherCmd.Flags().StringP("key", "k")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// weatherCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// weatherCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func weatherData(cmd *cobra.Command, args []string) {
	key, _ := cmd.Flags().GetString("key")
	url := "https://api.openweathermap.org/data/2.5/weather?lat=39.204940&lon=-94.532330&appid=" + key
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)

	fmt.Println(sb)
}

// func getWeatherData(api_key string) string {
// 	get_url := ("https://api.openweathermap.org/data/2.5/weather?lat=39.204940&lon=-94.532330&appid=%s", api_key)
// 	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?lat=39.204940&lon=-94.532330&appid=%s", api_key)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
//
// 	return sb
// }
