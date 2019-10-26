package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
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

	position_0 := string(information[0])
	position_1 := string(information[1])
	position_2 := string(information[2])
	position_3 := string(information[3])
	position_4 := string(information[4])
	position_5 := string(information[5])
	position_6 := string(information[6])
	position_7 := string(information[7])
	position_8 := string(information[8])
	position_9 := string(information[9])
	position_10 := string(information[10])
	position_11 := string(information[11])
	position_12 := string(information[12])
	position_13 := string(information[13])
	position_14 := string(information[14])
	position_15 := string(information[15])
	position_16 := string(information[16])
	position_17 := string(information[17])
	position_18 := string(information[18])
	position_19 := string(information[19])
	position_20 := string(information[20])
	position_21 := string(information[21])

	fmt.Println(position_0, position_1, position_2, position_3, position_4, position_5, position_6, position_7, position_8, position_9, position_10, position_11, position_12, position_13, position_14, position_15, position_16, position_17, position_18, position_19, position_20, position_21)

}

func main() {

	var base string

	fmt.Print("Insert Base64 information: ")
	fmt.Scan(&base)

	value := converter(base)
	base64ToObject(value)

}
