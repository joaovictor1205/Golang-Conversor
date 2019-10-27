package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

func converter(data string) string {

	encoded := base64.StdEncoding.EncodeToString([]byte(data)) // CONVERTING STRING TO BASE64

	decoded, _ := base64.StdEncoding.DecodeString(encoded) // CONVERTING BASE64 TO STRING

	b64, _ := base64.StdEncoding.DecodeString(string(decoded)) // CONVERTING STRING TO HEXADECIMAL

	hexa := hex.EncodeToString(b64)

	fmt.Println(hexa)

	return hexa

}

func base64ToObject(information string) {

	var type_string_1 string
	var chanel_value_1 int
	var sensor_value_1 string

	/*
		var type_string_2 string
		var chanel_value_2 int

		var type_string_3 string
		var chanel_value_3 int
	*/

	position_0 := string(information[0])
	position_1 := string(information[1])
	position_2 := string(information[2]) // *****
	position_3 := string(information[3]) // *****
	position_4 := string(information[4])
	position_5 := string(information[5])
	position_6 := string(information[6])
	position_7 := string(information[7])
	position_8 := string(information[8])
	position_9 := string(information[9])
	position_10 := string(information[10]) // *****
	position_11 := string(information[11]) // *****
	position_12 := string(information[12])
	position_13 := string(information[13])
	position_14 := string(information[14])
	position_15 := string(information[15])
	position_16 := string(information[16])
	position_17 := string(information[17])
	position_18 := string(information[18]) // *****
	position_19 := string(information[19]) // *****
	position_20 := string(information[20])
	position_21 := string(information[21])

	fmt.Println(position_0, position_1, position_2, position_3, position_4, position_5, position_6, position_7, position_8, position_9, position_10, position_11, position_12, position_13, position_14, position_15, position_16, position_17, position_18, position_19, position_20, position_21)

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
		log.Println("Error")
	}

	if position_2 == "7" && position_3 == "3" {
		type_string_1 = "Barometer"
	} else if position_2 == "6" && position_3 == "7" {
		type_string_1 = "Temperature"
	} else if position_3 == "6" && position_4 == "8" {
		type_string_1 = "Humidity"
	} else {
		type_string_1 = "Error"
	}

	sensor_value_1 = string(information[4:8])
	fmt.Println(sensor_value_1)
	sensor_hexa_value_1 := hex.EncodeToString([]byte(sensor_value_1))
	fmt.Println(sensor_hexa_value_1)

	bs, _ := hex.DecodeString(sensor_hexa_value_1)
	fmt.Println(string(bs))

	n, _ := strconv.ParseUint(string(bs), 16, 32)

	fmt.Println("Chanel: ", chanel_value_1)

	fmt.Println("Sensor Type: ", type_string_1)

	fmt.Println("Sensor Value: ", float64(n)*0.1)
}

func main() {

	var base string

	fmt.Print("Insert Base64 information: ")
	fmt.Scan(&base)

	value := converter(base)
	base64ToObject(value)

}
