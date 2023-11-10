/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

type weather struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		MinTemp   float64 `json:"temp_min"`
		MaxTemp   float64 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

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
	url := "https://api.openweathermap.org/data/2.5/weather?lat=39.204940&lon=-94.532330&units=imperial&appid=" + key
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Request for weather data was not successful")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var CurrentWeather weather
	err = json.Unmarshal(body, &CurrentWeather)
	if err != nil {
		log.Fatalln(err)
	}

	description := CurrentWeather.Weather[0].Description

	fmt.Printf("It is currently %s and %.2f\n", description, CurrentWeather.Main.Temp)
	fmt.Printf("The high today is %.2f and the low today is %.2f with a wind speed of %.2f", CurrentWeather.Main.MaxTemp, CurrentWeather.Main.MinTemp, CurrentWeather.Wind.Speed)
}
