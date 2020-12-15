

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the car structure, with 5 properties.  Structure tags are used by encoding/json library
type Car struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Colour string `json:"colour"`
	Owner  string `json:"owner"`
	Srno   string 'json:"srno"'
}

// Define the carowner structure, with 4 properties.  Structure tags are used by encoding/json library
type Carowner struct {
	Name   string `json:"name"`
	Id  string `json:"id"`
	Gender string `json:"gender"`
	Address  string `json:"address"`
}


func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}


func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryCar" {
		return s.queryCar(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createCar" {
		return s.createCar(APIstub, args)
	} else if function == "changeCarOwner" {
		return s.changeCarOwner(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	carAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(carAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	cars := []Car{
		Car{Make: "Toyota", Model: "Prius", Colour: "blue", Owner: "Tomoko", Srno: "1"},
		Car{Make: "Ford", Model: "Mustang", Colour: "red", Owner: "Brad", Srno: "2"},
		Car{Make: "Hyundai", Model: "Tucson", Colour: "green", Owner: "Jin Soo", Srno: "3"},
		Car{Make: "Volkswagen", Model: "Passat", Colour: "yellow", Owner: "Max", Srno: "4"},
		Car{Make: "Tesla", Model: "S", Colour: "black", Owner: "Adriana", Srno: "5"},
		Car{Make: "Peugeot", Model: "205", Colour: "purple", Owner: "Michel", Srno: "6"},
		Car{Make: "Chery", Model: "S22L", Colour: "white", Owner: "Aarav", Srno: "7"},
		Car{Make: "Fiat", Model: "Punto", Colour: "violet", Owner: "Pari", Srno: "8"},
		Car{Make: "Tata", Model: "Nano", Colour: "indigo", Owner: "Valeria", Srno: "9"},
		Car{Make: "Holden", Model: "Barina", Colour: "brown", Owner: "Shotaro", Srno: "10"},
	}

	i := 0
	for i < len(cars) {
		fmt.Println("i is ", i)
		carAsBytes, _ := json.Marshal(cars[i])
		APIstub.PutState("CAR"+strconv.Itoa(i), carAsBytes)
		fmt.Println("Added", cars[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var car = Car{Make: args[1], Model: args[2], Colour: args[3], Owner: args[4], Srno: args[5]}

	carAsBytes, _ := json.Marshal(car)
	APIstub.PutState(args[0], carAsBytes)

	return shim.Success(nil)
}



func (s *SmartContract) changeCarOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	carAsBytes, _ := APIstub.GetState(args[0])
	car := Car{}

	json.Unmarshal(carAsBytes, &car)
	car.Owner = args[1]

	carAsBytes, _ = json.Marshal(car)
	APIstub.PutState(args[0], carAsBytes)

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}