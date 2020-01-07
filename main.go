package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type sensor struct {
	NAME                    string
	LPP_DATA_TYPE           int
	HEXA                    int
	DATA_SIZE               int
	DATA_RESOLUTION_PER_BIT float64
}

func createSensor(NAME string, LPP_DATA_TYPE int, HEXA int, DATA_SIZE int, DATA_RESOLUTION_PER_BIT float64) *sensor {

	s := sensor{
		NAME:                    NAME,
		LPP_DATA_TYPE:           LPP_DATA_TYPE,
		HEXA:                    HEXA,
		DATA_SIZE:               DATA_SIZE,
		DATA_RESOLUTION_PER_BIT: DATA_RESOLUTION_PER_BIT,
	}

	return &s

}

var ILLUMINANCE_SENSOR = createSensor("Illuminance Sensor", 101, 65, 2, 1.0)
var PRESENCE_SENSOR = createSensor("Presence Sensor", 102, 66, 1, 1.0)
var TEMPERATURE_SENSOR = createSensor("Temperature Sensor", 103, 67, 2, 0.1)
var HUMIDITY_SENSOR = createSensor("Humidity Sensor", 104, 68, 1, 0.5)
var ACCELEROMETER_SENSOR = createSensor("Accelerometer Sensor", 113, 71, 6, 0.001)
var BAROMETER_SENSOR = createSensor("Barometer Sensor", 115, 73, 2, 0.1)
var GYROMETER_SENSOR = createSensor("Gyrometer Sensor", 134, 86, 6, 0.01)
var GPS_LOCATION_SENSOR = createSensor("GPS Location Sensor", 136, 88, 9, 0.0001)

func converter(data string) string {

	encoded := base64.StdEncoding.EncodeToString([]byte(data)) // CONVERTING STRING TO BASE64
	decoded, _ := base64.StdEncoding.DecodeString(encoded)     // CONVERTING BASE64 TO STRING
	b64, _ := base64.StdEncoding.DecodeString(string(decoded)) // CONVERTING STRING TO HEXADECIMAL
	hexa := hex.EncodeToString(b64)

	return hexa

}

// values to test:
// AHMCbQFnAC0CaCQ=
// AHMAAAFnAAACaAA=

func objectToJson(information string, finishingRoutine chan string) {

	if len(string(information)) == 0 {
		fmt.Println("ERROR! Please verify your value")
		os.Exit(-1)
	}

	var type_sensor_1 string
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

	var next_string_position int

	//position_8 := string(information[8])
	//position_9 := string(information[9])
	position_10 := string(information[10])
	position_11 := string(information[11])
	position_16 := string(information[16])
	position_17 := string(information[17])
	position_18 := string(information[18])
	position_19 := string(information[19])

	///////////////////////// FIRST SENSOR INFORMATION ///////////////////
	position_0 = string(information[0])                                  // FIRST POSITION
	position_1 = string(information[1])                                  // SECOND POSITION
	position_2 = string(information[2])                                  // THIRD POSITION
	position_3 = string(information[3])                                  // FOURTH POSITION
	chanel_value_1 = chanelValue(position_0, position_1)                 // WITH THE FIRST AND SECOND POSITION (1 BYTE) I CAN DEFINE THE FIRST CHANEL VALUE
	type_sensor_1 = sensorType(position_2, position_3)                   // WITH THE THIRD AND FOURTH POSITION (1 BYTE) I CAN DEFINE THE FIRST SENSOR TYPE
	sensor_bytes = numOfBytes(type_sensor_1)                             // THE SENSOR TYPE WILL DEFINE NUMBER OF BYTES TO BE READ
	sensor_value_1 = bytesReading(sensor_bytes, information)             // WITH THIS FUNCTION I HAVE THE 'PIECE' OF THE INFORMATION THAT CONTAINS THE SENSOR VALUE
	first_sensor_value = sensorConversion(sensor_value_1, type_sensor_1) // SO NOW, WITH THIS 'PIECE' I CAN TRANSLATE THE VALUE TO INT TO SHOW TO USER
	if sensor_bytes == -1 {
		fmt.Println("Error with Sensor")
		os.Exit(-1)
	}
	//////////////////////////////////////////////////////////////////////

	// WHEN WE FINISHED READING THE FIRST SENSOR INFORMATION (CHANEL VALUE + SENSOR TYPE + SENSOR VALUE)
	// WE NEED TO KNOW IF WE HAVE MORE SENSORS ON THE PAYLOAD
	// WITH THE HASSUFFIX FUNCTION WE KNOW IF THE SENSOR VALUE IS THE LAST THING ON THE STRING
	// IF THE FIRST SENSOR VALUE ISN'T THE LAST THING ON THE STRING, WE NEED TO KEEP GOING THROUGH THE REST OF THE STRING
	boolean_end_string := strings.HasSuffix(information, sensor_value_1)
	if boolean_end_string == true {
		os.Exit(-1)
	} else {
		fmt.Println("test")
	}
	if sensor_bytes == 2 {
		next_string_position = strings.Index(information, sensor_value_1) + 5 // THIS FUNCTION SHOWS THE POSITION OF THE SENSOR VALUE ON THE STRING INFORMATION
		fmt.Println(next_string_position)
	} else if sensor_bytes == 4 {
		next_string_position = strings.Index(information, sensor_value_1) + 13
	} else if sensor_bytes == 6 {
		next_string_position = strings.Index(information, sensor_value_1) + 17
	} else if sensor_bytes == 8 {
		next_string_position = strings.Index(information, sensor_value_1) + 19
	} else {
		next_string_position = strings.Index(information, sensor_value_1) + 3
	}

	//next_position := strconv.Itoa(next_string_position)

	////////////////////// SECOND SENSOR INFORMATION ////////////////////
	chanel_value_2 = chanelValue("0", string(information[int(next_string_position)]))
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
	createJSON(chanel_value_1, type_sensor_1, first_sensor_value,
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
	case ILLUMINANCE_SENSOR.HEXA:
		return ILLUMINANCE_SENSOR.NAME
	case PRESENCE_SENSOR.HEXA:
		return PRESENCE_SENSOR.NAME
	case TEMPERATURE_SENSOR.HEXA:
		return TEMPERATURE_SENSOR.NAME
	case HUMIDITY_SENSOR.HEXA:
		return HUMIDITY_SENSOR.NAME
	case ACCELEROMETER_SENSOR.HEXA:
		return ACCELEROMETER_SENSOR.NAME
	case BAROMETER_SENSOR.HEXA:
		return BAROMETER_SENSOR.NAME
	case GYROMETER_SENSOR.HEXA:
		return GYROMETER_SENSOR.NAME
	case GPS_LOCATION_SENSOR.HEXA:
		return GPS_LOCATION_SENSOR.NAME
	default:
		return "Please verify your Sensor Type"
	}

}

func sensorConversion(first_parameter string, sensor_type string) float64 {

	var sensor_value float64

	sensor_hexa_value := hex.EncodeToString([]byte(first_parameter))    // STRING FOR HEXA
	hexa_to_string, err := hex.DecodeString(sensor_hexa_value)          // HEXA FOR STRING
	int_value, err := strconv.ParseUint(string(hexa_to_string), 16, 32) // STRING FOR INT64

	if err != nil {
		fmt.Println("Error on Sensor value conversion")
		os.Exit(-1)
	}

	switch sensor_type {
	case ILLUMINANCE_SENSOR.NAME:
		sensor_value = float64(int_value) * ILLUMINANCE_SENSOR.DATA_RESOLUTION_PER_BIT
	case PRESENCE_SENSOR.NAME:
		sensor_value = float64(int_value) * PRESENCE_SENSOR.DATA_RESOLUTION_PER_BIT
	case TEMPERATURE_SENSOR.NAME:
		sensor_value = float64(int_value) * TEMPERATURE_SENSOR.DATA_RESOLUTION_PER_BIT
	case HUMIDITY_SENSOR.NAME:
		sensor_value = float64(int_value) * HUMIDITY_SENSOR.DATA_RESOLUTION_PER_BIT
	case ACCELEROMETER_SENSOR.NAME:
		sensor_value = float64(int_value) * ACCELEROMETER_SENSOR.DATA_RESOLUTION_PER_BIT
	case BAROMETER_SENSOR.NAME:
		sensor_value = float64(int_value) * BAROMETER_SENSOR.DATA_RESOLUTION_PER_BIT
	case GYROMETER_SENSOR.NAME:
		sensor_value = float64(int_value) * GYROMETER_SENSOR.DATA_RESOLUTION_PER_BIT
	case GPS_LOCATION_SENSOR.NAME:
		sensor_value = float64(int_value) * GPS_LOCATION_SENSOR.DATA_RESOLUTION_PER_BIT

	}

	return sensor_value
}

func numOfBytes(sensor_type string) int {

	switch sensor_type {
	case ILLUMINANCE_SENSOR.NAME:
		return ILLUMINANCE_SENSOR.DATA_SIZE
	case PRESENCE_SENSOR.NAME:
		return PRESENCE_SENSOR.DATA_SIZE
	case TEMPERATURE_SENSOR.NAME:
		return TEMPERATURE_SENSOR.DATA_SIZE
	case HUMIDITY_SENSOR.NAME:
		return HUMIDITY_SENSOR.DATA_SIZE
	case ACCELEROMETER_SENSOR.NAME:
		return ACCELEROMETER_SENSOR.DATA_SIZE
	case BAROMETER_SENSOR.NAME:
		return BAROMETER_SENSOR.DATA_SIZE
	case GYROMETER_SENSOR.NAME:
		return GYROMETER_SENSOR.DATA_SIZE
	case GPS_LOCATION_SENSOR.NAME:
		return GPS_LOCATION_SENSOR.DATA_SIZE
	default:
		return -1
	}

}

func bytesReading(sensor_bytes int, information string) string {

	switch sensor_bytes {
	case 1:
		return string(information[4:6])
	case 2:
		return string(information[4:8])
	case 6:
		return string(information[4:16])
	case 9:
		return string(information[4:22])
	default:
		return "Err"
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
