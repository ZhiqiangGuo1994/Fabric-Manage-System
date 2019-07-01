package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct{}

type country struct {
	Patentid      string `json:"patentid"`
	Patentname    string `json:"patentname"`
	Applicationid string `json:"applicationid"`

	Approvetime string `json:"approvetime"`
	Reason      string `json:"reason"`
	Flag        string `json:"flag"`

	Acceptance     string `json:"acceptance"`
	Acceptancetime string `json:"acceptancetime"`
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	if function == "initCountry" {
		return t.initCountry(stub, args)
	} else if function == "readCountry" {
		return t.readCountry(stub, args)
	} else if function == "acceptanceCountry" {
		return t.acceptanceCountry(stub, args)
	} else if function == "approveCountry" {
		return t.approveCountry(stub, args)
	} else if function == "queryByApplicationId" {
		return t.queryByApplicationId(stub, args)
	}
	return shim.Success(nil)
}

//============================================================
// initCountry - create a new country, store into chaincode state
//============================================================
func (t *SimpleChaincode) initCountry(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguements.")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	patentid := args[0]
	patentname := args[1]
	applicationid := args[2]
	flag := args[3]

	//check if province already exists
	countryAsBytes, err := stub.GetState(patentid)
	if err != nil {
		return shim.Error("Failed to get province: " + err.Error())
	} else if countryAsBytes != nil {
		return shim.Error("Province has exited!")
	}

	//Create province object and marshal to JSON
	country := &country{patentid, patentname, applicationid, "", "", flag, "0", ""}
	countryJSONasBytes, err := json.Marshal(country)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(patentid, countryJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *SimpleChaincode) readCountry(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var patentid, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the patent to query")
	}

	patentid = args[0]
	valAsbytes, err := stub.GetState(patentid) //get the patent from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + patentid + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Patent does not exist: " + patentid + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

func (t *SimpleChaincode) acceptanceCountry(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting name of the country to approveCountry")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	patentid := args[0]
	acceptancetime := args[1]
	countryAsBytes, err := stub.GetState(patentid)
	if err != nil {
		return shim.Error("Failed to get country: " + err.Error())
	} else if countryAsBytes == nil {
		return shim.Error("Country has exited!")
	}

	countryToTransfer := country{}
	err = json.Unmarshal(countryAsBytes, &countryToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	countryToTransfer.Acceptancetime = acceptancetime
	countryToTransfer.Acceptance = "1"
	countryJSONasBytes, _ := json.Marshal(countryToTransfer)
	err = stub.PutState(patentid, countryJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end transferPatent (success)")
	return shim.Success(nil)
}

func (t *SimpleChaincode) approveCountry(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting name of the country to approveCountry")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	patentid := args[0]
	approvetime := args[1]
	countryAsBytes, err := stub.GetState(patentid)
	if err != nil {
		return shim.Error("Failed to get country: " + err.Error())
	} else if countryAsBytes == nil {
		return shim.Error("Country has exited!")
	}

	countryToTransfer := country{}
	err = json.Unmarshal(countryAsBytes, &countryToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	countryToTransfer.Approvetime = approvetime
	countryToTransfer.Flag = "1"
	countryJSONasBytes, _ := json.Marshal(countryToTransfer)
	err = stub.PutState(patentid, countryJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end transferPatent (success)")
	return shim.Success(nil)
}

//用户查询个人申请的专利
func (t *SimpleChaincode) queryByApplicationId(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	applicationid := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"applicationid\":\"%s\"}}", applicationid)
	queryResults, err := getQueryResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Print("Error starting Simple chaincode: %s", err)
	}
}
