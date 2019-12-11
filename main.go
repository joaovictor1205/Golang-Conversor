package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"time"
)

// ILLUMINANCE SENSOR -> LPP = 101 / HEX = 65 / DATA SIZE = 2 / 1 UNSIGNED MSB
// PRESENCE SENSOR -> LPP = 102 / HEX = 66 / DATA SIZE = 1 / 1
// TEMPERATURE SENSOR -> LPP = 103 / HEX = 67 / DATA SIZE = 2 / 0.1 SIGNED MSB
// HUMIDITY SENSOR -> LPP = 104 / HEX = 68 / DATA SIZE = 1 / 0.5 UNSIGNED
// ACCELEROMETER SENSOR -> LPP = 113 / HEX = 71 / DATA SIZE = 6 / 0.001 SIGNED MSB
// BAROMETER SENSOR -> LPP = 115 / HEX = 73 / DATA SIZE = 2 / 0.1 UNSIGNED MSB
// GYROMETER SENSOR -> LPP = 134 / HEX = 86 / DATA SIZE = 6 / 0.01 SIGNED MSB

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
	position_8 := string(information[8])
	position_9 := string(information[9])
	position_10 := string(information[10])
	position_11 := string(information[11])
	position_16 := string(information[16])
	position_17 := string(information[17])
	position_18 := string(information[18])
	position_19 := string(information[19])

	///////////////////////// FIRST SENSOR INFORMATION ///////////////////
	chanel_value_1 = chanelValue(position_0, position_1)
	type_string_1 = sensorType(position_2, position_3)
	sensor_value_1 = string(information[4:8])
	first_sensor_value = sensorConversion(sensor_value_1)
	//////////////////////////////////////////////////////////////////////

	////////////////////// SECOND SENSOR INFORMATION ////////////////////
	chanel_value_2 = chanelValue(position_8, position_9)
	type_string_2 = sensorType(position_10, position_11)
	sensor_value_2 = string(information[12:16])
	second_sensor_value = sensorConversion(sensor_value_2)
	////////////////////////////////////////////////////////////////////
	
	////////////////// THIRDY SENSOR INFORMATION ///////////////////////
	chanel_value_3 = chanelValue(position_16, position_17)
	type_string_3 = sensorType(position_18, position_19)

	sensor_value_3 = string(information[20:22])
	if type_string_3 == "Humidity" {
		third_sensor_value = sensorConversionHumidity(sensor_value_3)
	} else {
		third_sensor_value = sensorConversion(sensor_value_3)
	}
	////////////////////////////////////////////////////////////////////

	////////////// CREATING JSON WITH SENSOR INFORMATIONS //////////////
	createJSON(chanel_value_1, type_string_1, first_sensor_value,
		chanel_value_2, type_string_2, second_sensor_value,
		chanel_value_3, type_string_3, third_sensor_value)
	////////////////////////////////////////////////////////////////////

}

// sensor vai receber mais valores ou apenas 3 por leitura?
func chanelValue(first_parameter, second_parameter string) int {

	var chanel_value int

	if first_parameter == "0" && second_parameter == "0" {
		chanel_value = 0
	} else if first_parameter == "0" && second_parameter == "1" {
		chanel_value = 1
	} else if first_parameter == "0" && second_parameter == "2" {
		chanel_value = 2
	} else {
		chanel_value = -1
	}

	if chanel_value == -1 {
		fmt.Println("Error on Chanel")
		os.Exit(-1)
	}

	return chanel_value
}

// feature -> terminar de inserir os outros tipos de sensores para
// seus respectivos valores
// ex.: Accelerometer = 71
func sensorType(first_parameter, second_parameter string) string {

	var type_string string

	if first_parameter == "7" && second_parameter == "3" {
		type_string = "Barometer"
	} else if first_parameter == "6" && second_parameter == "7" {
		type_string = "Temperature"
	} else if first_parameter == "6" && second_parameter == "8" {
		type_string = "Humidity"
	} else {
		type_string = "Error"
		os.Exit(-1)
		fmt.Println("Error on Sensor Type")
	}

	return type_string
}

// problema: temperatura e barometro multiplicam por 0.1 mas
// humidade multiplica por 0.5
// feature -> resolver com 1 funcao apenas
func sensorConversion(first_parameter string) float64 {

	sensor_hexa_value_1 := hex.EncodeToString([]byte(first_parameter))    // STRING FOR HEXA
	hexa_to_string_1, _ := hex.DecodeString(sensor_hexa_value_1)          // HEXA FOR STRING
	int_value_1, _ := strconv.ParseUint(string(hexa_to_string_1), 16, 32) // STRING FOR INT64
	sensor_value := float64(int_value_1) * 0.1

	return sensor_value
}

func sensorConversionHumidity(first_parameter string) float64 {

	sensor_hexa_value_1 := hex.EncodeToString([]byte(first_parameter))    // STRING FOR HEXA
	hexa_to_string_1, _ := hex.DecodeString(sensor_hexa_value_1)          // HEXA FOR STRING
	int_value_1, _ := strconv.ParseUint(string(hexa_to_string_1), 16, 32) // STRING FOR INT64
	sensor_value := float64(int_value_1) * 0.5

	return sensor_value
}

// como melhorar essa formar de criar JSON? fazer um objeto e aplicar um
// marshal nele?
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

	fmt.Print("Insert Base64 information: ")
	fmt.Scan(&base)

	value = converter(base)
	go objectToJson(value)

// rotina nao pode depender do Millisecond, pq se a maquina tiver um
// desempenho ruim 1 millisecond nao e o suficiente para a rotina terminar
// feature -> criar um Chanel para comunicar com a Routine
	time.Sleep(time.Millisecond)
	fmt.Println("Finishing Go Routine")
}
