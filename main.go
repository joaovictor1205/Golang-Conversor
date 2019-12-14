package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

func converter(data string) string {

	encoded := base64.StdEncoding.EncodeToString([]byte(data)) // CONVERTING STRING TO BASE64

	decoded, _ := base64.StdEncoding.DecodeString(encoded) // CONVERTING BASE64 TO STRING

	b64, _ := base64.StdEncoding.DecodeString(string(decoded)) // CONVERTING STRING TO HEXADECIMAL

	hexa := hex.EncodeToString(b64)

	return hexa

}

func objectToJson(information string, finishingRoutine chan string) {

	if len(string(information)) == 0 {
		fmt.Println("ERROR! Please verify your value")
		os.Exit(-1)
	}

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

	var sensor_bytes int

	var position_0 string
	var position_1 string
	var position_2 string
	var position_3 string

	position_8 := string(information[8])
	position_9 := string(information[9])
	position_10 := string(information[10])
	position_11 := string(information[11])
	position_16 := string(information[16])
	position_17 := string(information[17])
	position_18 := string(information[18])
	position_19 := string(information[19])

	for i := 0; i <= 3; i++ {
		position_0 = string(information[0])
		position_1 = string(information[1])
		position_2 = string(information[2])
		position_3 = string(information[3])
	}

	///////////////////////// FIRST SENSOR INFORMATION ///////////////////
	chanel_value_1 = chanelValue(position_0, position_1)
	type_string_1 = sensorType(position_2, position_3)
	sensor_bytes = numOfBytes(type_string_1)
	if sensor_bytes == 2 {
		sensor_value_1 = string(information[4:8])
		first_sensor_value = sensorConversion(sensor_value_1, type_string_1)
	}
	//////////////////////////////////////////////////////////////////////

	////////////////////// SECOND SENSOR INFORMATION ////////////////////
	chanel_value_2 = chanelValue(position_8, position_9)
	type_string_2 = sensorType(position_10, position_11)
	sensor_value_2 = string(information[12:16])
	second_sensor_value = sensorConversion(sensor_value_2, type_string_2)
	////////////////////////////////////////////////////////////////////

	////////////////// THIRDY SENSOR INFORMATION ///////////////////////
	chanel_value_3 = chanelValue(position_16, position_17)
	type_string_3 = sensorType(position_18, position_19)
	sensor_value_3 = string(information[20:22])
	third_sensor_value = sensorConversion(sensor_value_3, type_string_3)
	////////////////////////////////////////////////////////////////////

	////////////// CREATING JSON WITH SENSOR INFORMATIONS //////////////
	createJSON(chanel_value_1, type_string_1, first_sensor_value,
		chanel_value_2, type_string_2, second_sensor_value,
		chanel_value_3, type_string_3, third_sensor_value)
	////////////////////////////////////////////////////////////////////

	finishingRoutine <- "Finishing Go Routine"

}

func chanelValue(first_parameter, second_parameter string) int {

	chanel_value, err := strconv.Atoi(first_parameter + second_parameter)

	if err != nil {
		fmt.Println("Error on Chanel")
		os.Exit(-1)
	}

	return chanel_value
}

func sensorType(first_parameter, second_parameter string) string {

	type_string, err := strconv.Atoi(first_parameter + second_parameter)

	if err != nil {
		fmt.Println("Error on Sensor")
		os.Exit(-1)
	}

	switch type_string {
	case 65:
		return "Illuminance"
	case 66:
		return "Presence"
	case 67:
		return "Temperature"
	case 68:
		return "Humidity"
	case 71:
		return "Accelerometer"
	case 73:
		return "Barometer"
	case 86:
		return "Gyrometer"
	case 88:
		return "GPS Location"
	default:
		return "Please verify your Sensor Type"
	}

}

func sensorConversion(first_parameter string, sensor_type string) float64 {

	var sensor_value float64

	sensor_hexa_value_1 := hex.EncodeToString([]byte(first_parameter))    // STRING FOR HEXA
	hexa_to_string_1, _ := hex.DecodeString(sensor_hexa_value_1)          // HEXA FOR STRING
	int_value_1, _ := strconv.ParseUint(string(hexa_to_string_1), 16, 32) // STRING FOR INT64

	switch sensor_type {
	case "Iluminance":
		sensor_value = float64(int_value_1) * 1
	case "Presence":
		sensor_value = float64(int_value_1) * 1
	case "Temperature":
		sensor_value = float64(int_value_1) * 0.1
	case "Humidity":
		sensor_value = float64(int_value_1) * 0.5
	case "Accelerometer":
		sensor_value = float64(int_value_1) * 0.001
	case "Barometer":
		sensor_value = float64(int_value_1) * 0.1
	case "Gyrometer":
		sensor_value = float64(int_value_1) * 0.01
	case "GPS Location":
		sensor_value = float64(int_value_1) * 0.0001

	}

	return sensor_value
}

func numOfBytes(sensor_type string) int {

	switch sensor_type {
	case "Iluminance":
		return 2
	case "Presence":
		return 1
	case "Temperature":
		return 2
	case "Humidity":
		return 1
	case "Accelerometer":
		return 6
	case "Barometer":
		return 2
	case "Gyrometer":
		return 6
	case "GPS Location":
		return 9
	default:
		return -1
	}

}

func createJSON(chanel_value_1 int, type_string_1 string, sensor_value_1 float64,
	chanel_value_2 int, type_string_2 string, sensor_value_2 float64,
	chanel_value_3 int, type_string_3 string, sensor_value_3 float64) {

	fmt.Println("[")
	fmt.Printf(`{"Chanel One"`+": "+"%d,\n"+`"Sensor One Type"`+": "+`"`+"%s"+`"`+",\n"+`"Sensor One Value"`+": "+"%.1f}\n", chanel_value_1, type_string_1, sensor_value_1)
	fmt.Print(",")
	fmt.Print("\n")
	fmt.Printf(`{"Chanel Two"`+": "+"%d,\n"+`"Sensor Two Type"`+": "+`"`+"%s"+`"`+",\n"+`"Sensor Two Value"`+": "+"%.1f}\n", chanel_value_2, type_string_2, sensor_value_2)
	fmt.Print(",")
	fmt.Print("\n")
	fmt.Printf(`{"Chanel Three"`+": "+"%d,\n"+`"Sensor Three Type"`+": "+`"`+"%s"+`"`+",\n"+`"Sensor Three Value"`+": "+"%.1f}\n", chanel_value_3, type_string_3, sensor_value_3)
	fmt.Println("]")
}

func main() {

	var base string
	var value string
	finishingRoutine := make(chan string)

	fmt.Print("Insert Base64 information: ")
	fmt.Scan(&base)

	value = converter(base)

	go objectToJson(value, finishingRoutine)
	fmt.Println(<-finishingRoutine)
}
