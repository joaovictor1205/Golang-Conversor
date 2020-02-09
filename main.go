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

	encoded := base64.StdEncoding.EncodeToString([]byte(data))   // CONVERTING STRING TO BASE64
	decoded, err := base64.StdEncoding.DecodeString(encoded)     // CONVERTING BASE64 TO STRING
	b64, err := base64.StdEncoding.DecodeString(string(decoded)) // CONVERTING STRING TO HEXADECIMAL
	hexa := hex.EncodeToString(b64)

	if err != nil {
		fmt.Println("Error with payload")
		os.Exit(-1)
	}

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

	var chanel_value int
	var type_sensor string
	var sensor_bytes int
	var sensor_value string
	var conversion_sensor float64

	var position_0 string
	var position_1 string
	var position_2 string
	var position_3 string

	var next_string_position int

	var finishing_reading bool
	var end_of_reading int

	///////////////////////// FIRST SENSOR INFORMATION ///////////////////
	position_0 = string(information[0])                                  // FIRST POSITION
	position_1 = string(information[1])                                  // SECOND POSITION
	position_2 = string(information[2])                                  // THIRD POSITION
	position_3 = string(information[3])                                  // FOURTH POSITION
	chanel_value_1 = chanelValue(position_0, position_1)                 // WITH THE FIRST AND SECOND POSITION (1 BYTE) I CAN DEFINE THE FIRST CHANEL VALUE
	type_sensor_1 = sensorType(position_2, position_3)                   // WITH THE THIRD AND FOURTH POSITION (1 BYTE) I CAN DEFINE THE FIRST SENSOR TYPE
	sensor_bytes = numOfBytes(type_sensor_1)                             // THE SENSOR TYPE WILL DEFINE NUMBER OF BYTES TO BE READ
	sensor_value_1 = bytesReading(sensor_bytes, information, 4)          // WITH THIS FUNCTION I HAVE THE 'PIECE' OF THE INFORMATION THAT CONTAINS THE SENSOR VALUE
	first_sensor_value = sensorConversion(sensor_value_1, type_sensor_1) // SO NOW, WITH THIS 'PIECE' I CAN TRANSLATE THE VALUE TO INT TO SHOW TO USER
	if sensor_bytes == -1 {
		fmt.Println("Error with Sensor")
		os.Exit(-1)
	}
	fmt.Println(chanel_value_1, type_sensor_1, first_sensor_value)
	//////////////////////////////////////////////////////////////////////

	end_of_reading = strings.Index(information, sensor_value_1)
	end_of_reading = end_of_reading + 1
	if end_of_reading == 1 {
		fmt.Println("Finishing Go Routine")
		os.Exit(0)
	} else {
		next_string_position = next_position(sensor_bytes, information, sensor_value_1)

		for i := 0; finishing_reading != true; i++ {

			chanel_value = chanelValue(string(information[int(next_string_position)]), string(information[int(next_string_position)+1]))
			type_sensor = sensorType(string(information[int(next_string_position)+2]), string(information[int(next_string_position)+3]))
			sensor_bytes = numOfBytes(type_sensor)
			sensor_value = bytesReading(sensor_bytes, information, next_string_position+4)
			conversion_sensor = sensorConversion(sensor_value, type_sensor)

			fmt.Println(chanel_value, type_sensor, conversion_sensor)

			finishing_reading = strings.HasSuffix(information, sensor_value)

			if finishing_reading == true {
				break
			} else {
				next_string_position = next_position(sensor_bytes, information, sensor_value)
				if sensor_value == sensor_value_1 {
					next_string_position += next_position(sensor_bytes, information, sensor_value)
				}
				continue
			}
		}
	}

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

func bytesReading(sensor_bytes int, information string, actual_position int) string {

	switch sensor_bytes {
	case 1:
		return string(information[actual_position : actual_position+sensor_bytes+1])
	case 2:
		return string(information[actual_position : actual_position+sensor_bytes+2])
	case 6:
		return string(information[actual_position : actual_position+sensor_bytes+6])
	case 9:
		return string(information[actual_position : actual_position+sensor_bytes+9])
	default:
		return "Err"
	}
}

func next_position(sensor_bytes int, information string, sensor_value string) int {

	switch sensor_bytes {
	case 2:
		return strings.Index(information, sensor_value) + 4
	case 4:
		return strings.Index(information, sensor_value) + 12
	case 6:
		return strings.Index(information, sensor_value) + 16
	case 8:
		return strings.Index(information, sensor_value) + 18
	default:
		return strings.Index(information, sensor_value) + 2
	}
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
