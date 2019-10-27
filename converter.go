package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func converter(data string) string {

	encoded := base64.StdEncoding.EncodeToString([]byte(data)) // CONVERTING STRING TO BASE64

	decoded, _ := base64.StdEncoding.DecodeString(encoded) // CONVERTING BASE64 TO STRING

	b64, _ := base64.StdEncoding.DecodeString(string(decoded)) // CONVERTING STRING TO HEXADECIMAL

	hexa := hex.EncodeToString(b64)

	return hexa

}

func objectToJson(information string) {

	var type_string_1 string
	var chanel_value_1 int
	var sensor_value_1 string
	var first_sensor_value float64

	var type_string_2 string
	var chanel_value_2 int
	var sensor_value_2 string
	var second_sensor_value float64

	var type_string_3 string
	var chanel_value_3 int
	var sensor_value_3 string
	var third_sensor_value float64

	position_0 := string(information[0])
	position_1 := string(information[1])
	position_2 := string(information[2])
	position_3 := string(information[3])
	position_4 := string(information[4])
	//position_5 := string(information[5])
	//position_6 := string(information[6])
	//position_7 := string(information[7])
	position_8 := string(information[8])
	position_9 := string(information[9])
	position_10 := string(information[10])
	position_11 := string(information[11])
	//position_12 := string(information[12])
	//position_13 := string(information[13])
	//position_14 := string(information[14])
	//position_15 := string(information[15])
	position_16 := string(information[16])
	position_17 := string(information[17])
	position_18 := string(information[18])
	position_19 := string(information[19])
	//position_20 := string(information[20])
	//position_21 := string(information[21])

	if position_0 == "0" && position_1 == "0" {
		chanel_value_1 = 0
	} else if position_0 == "0" && position_1 == "1" {
		chanel_value_1 = 1
	} else if position_0 == "0" && position_1 == "2" {
		chanel_value_1 = 2
	} else {
		chanel_value_1 = -1
	}

	if chanel_value_1 == -1 {
		fmt.Println("Error on Chanel")
		os.Exit(-1)
	}

	if position_2 == "7" && position_3 == "3" {
		type_string_1 = "Barometer"
	} else if position_2 == "6" && position_3 == "7" {
		type_string_1 = "Temperature"
	} else if position_3 == "6" && position_4 == "8" {
		type_string_1 = "Humidity"
	} else {
		type_string_1 = "Error"
		os.Exit(-1)
		fmt.Println("Error on Sensor Type")
	}

	sensor_value_1 = string(information[4:8])
	sensor_hexa_value_1 := hex.EncodeToString([]byte(sensor_value_1))     // STRING FOR HEXA
	hexa_to_string_1, _ := hex.DecodeString(sensor_hexa_value_1)          // HEXA FOR STRING
	int_value_1, _ := strconv.ParseUint(string(hexa_to_string_1), 16, 32) // STRING FOR INT64
	first_sensor_value = float64(int_value_1) * 0.1

	if position_8 == "0" && position_9 == "0" {
		chanel_value_2 = 0
	} else if position_8 == "0" && position_9 == "1" {
		chanel_value_2 = 1
	} else if position_8 == "0" && position_9 == "2" {
		chanel_value_2 = 2
	} else {
		chanel_value_2 = -1
	}

	if chanel_value_2 == -1 {
		fmt.Println("Error on Chanel")
		os.Exit(-1)
	}

	if position_10 == "7" && position_11 == "3" {
		type_string_2 = "Barometer"
	} else if position_10 == "6" && position_11 == "7" {
		type_string_2 = "Temperature"
	} else if position_10 == "6" && position_11 == "8" {
		type_string_2 = "Humidity"
	} else {
		type_string_2 = "Error"
		os.Exit(-1)
		fmt.Println("Error on Sensor Type")
	}

	sensor_value_2 = string(information[12:16])
	sensor_hexa_value_2 := hex.EncodeToString([]byte(sensor_value_2))     // STRING FOR HEXA
	hexa_to_string_2, _ := hex.DecodeString(sensor_hexa_value_2)          // HEXA FOR STRING
	int_value_2, _ := strconv.ParseUint(string(hexa_to_string_2), 16, 32) // STRING FOR INT64
	second_sensor_value = float64(int_value_2) * 0.1

	if position_16 == "0" && position_17 == "0" {
		chanel_value_3 = 0
	} else if position_16 == "0" && position_17 == "1" {
		chanel_value_3 = 1
	} else if position_16 == "0" && position_17 == "2" {
		chanel_value_3 = 2
	} else {
		chanel_value_3 = -1
	}

	if chanel_value_3 == -1 {
		log.Println("Error")
	}

	if position_18 == "7" && position_19 == "3" {
		type_string_3 = "Barometer"
	} else if position_18 == "6" && position_19 == "7" {
		type_string_3 = "Temperature"
	} else if position_18 == "6" && position_19 == "8" {
		type_string_3 = "Humidity"
	} else {
		type_string_3 = "Error"
		os.Exit(-1)
		fmt.Println("Error on Sensor Type")
	}

	sensor_value_3 = string(information[20:22])
	sensor_hexa_value_3 := hex.EncodeToString([]byte(sensor_value_3))     // STRING FOR HEXA
	hexa_to_string_3, _ := hex.DecodeString(sensor_hexa_value_3)          // HEXA FOR STRING
	int_value_3, _ := strconv.ParseUint(string(hexa_to_string_3), 16, 32) // STRING FOR INT64
	third_sensor_value = float64(int_value_3) * 0.5

	fmt.Println("[")
	fmt.Printf(`{"Chanel One"`+": "+"%d,\n"+`"Sensor One Type"`+": "+`"`+"%s"+`"`+",\n"+`"Sensor One Value"`+": "+"%.1f}\n", chanel_value_1, type_string_1, first_sensor_value)
	fmt.Print(",")
	fmt.Print("\n")
	fmt.Printf(`{"Chanel Two"`+": "+"%d,\n"+`"Sensor Two Type"`+": "+`"`+"%s"+`"`+",\n"+`"Sensor Two Value"`+": "+"%.1f}\n", chanel_value_2, type_string_2, second_sensor_value)
	fmt.Print(",")
	fmt.Print("\n")
	fmt.Printf(`{"Chanel Three"`+": "+"%d,\n"+`"Sensor Three Type"`+": "+`"`+"%s"+`"`+",\n"+`"Sensor Three Value"`+": "+"%.1f}\n", chanel_value_3, type_string_3, third_sensor_value)
	fmt.Println("]")

}

func main() {

	var base string
	var value string

	fmt.Print("Insert Base64 information: ")
	fmt.Scan(&base)

	value = converter(base)
	go objectToJson(value)

	time.Sleep(time.Millisecond)
	fmt.Println("Finishing Go Routine")
}
