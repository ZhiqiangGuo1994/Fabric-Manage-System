package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct{}

type agency struct {
	Patentid        string `json:"patentid"`
	Patentname      string `json:"patentname"`
	Powerofattorney string `json:"powerofattorney"`
	Applicationid   string `json:"applicationid"`
	Submitid        string `json:"submitid"`
	Submittime      string `json:"submittime"`

	Flag            string `json:"flag"`
	Powermarkleroot string `json:"powermarkleroot"`
	Reason          string `json:"reason"`
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	if function == "initAgency" {
		//create a new agency
		return t.initAgency(stub, args)
	} else if function == "resubmitAgency" {
		return t.resubmitAgency(stub, args)
	} else if function == "alter" {
		return t.alter(stub, args)
	} else if function == "submit" {
		return t.submit(stub, args)
	} else if function == "refuse" {
		return t.refuse(stub, args)
	} else if function == "readAgency" {
		return t.readAgency(stub, args)
	} else if function == "queryByApplicationId" {
		return t.queryByApplicationId(stub, args)
	}
	return shim.Success(nil)
}

//============================================================
// initAgency - create a new agency, store into chaincode state
// ============================================================
func (t *SimpleChaincode) initAgency(stub shim.ChaincodeStubInterface, args []string) pb.Response {
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

	//check if agency already exists
	agencyAsBytes, err := stub.GetState(patentid)
	if err != nil {
		return shim.Error("Failed to get agency: " + err.Error())
	} else if agencyAsBytes != nil {
		return shim.Error("Patent has exited!")
	}

	//Create agency object and marshal to JSON
	agency := &agency{patentid, patentname, "", applicationid, "", "", flag, "", ""}
	agencyJSONasBytes, err := json.Marshal(agency)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(patentid, agencyJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *SimpleChaincode) resubmitAgency(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguements.")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	patentid := args[0]

	agencysAsBytes, err := stub.GetState(patentid)
	if err != nil {
		return shim.Error("Failed to get patent:" + err.Error())
	} else if agencysAsBytes == nil {
		return shim.Error("Patent does not exist")
	}
	agencyToTransfer := agency{}
	err = json.Unmarshal(agencysAsBytes, &agencyToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}

	agencyToTransfer.Flag = "0"

	agencyJSONasBytes, _ := json.Marshal(agencyToTransfer)
	err = stub.PutState(patentid, agencyJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end transferPatent (success)")
	return shim.Success(nil)
}

func (t *SimpleChaincode) alter(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 9 {
		return shim.Error("Incorrect number of arguements.")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	patentid := args[0]
	patentname := args[1]
	powerofattorney := args[2]
	applicationid := args[3]
	submitid := args[4]
	submittime := args[5]
	flag := args[6]
	powermarkleroot := args[7]
	reason := args[8]

	agencysAsBytes, err := stub.GetState(patentid)
	if err != nil {
		return shim.Error("Failed to get patent:" + err.Error())
	} else if agencysAsBytes == nil {
		return shim.Error("Patent does not exist")
	}
	agencyToTransfer := agency{}
	err = json.Unmarshal(agencysAsBytes, &agencyToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	agencyToTransfer.Patentname = patentname
	agencyToTransfer.Powerofattorney = powerofattorney
	agencyToTransfer.Applicationid = applicationid
	agencyToTransfer.Submitid = submitid
	agencyToTransfer.Submittime = submittime
	agencyToTransfer.Flag = flag
	agencyToTransfer.Powermarkleroot = powermarkleroot
	agencyToTransfer.Reason = reason

	agencyJSONasBytes, _ := json.Marshal(agencyToTransfer)
	err = stub.PutState(patentid, agencyJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end transferPatent (success)")
	return shim.Success(nil)
}

func (t *SimpleChaincode) submit(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguements.")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	patentid := args[0]
	submittime := args[1]

	agencysAsBytes, err := stub.GetState(patentid)
	if err != nil {
		return shim.Error("Failed to get patent:" + err.Error())
	} else if agencysAsBytes == nil {
		return shim.Error("Patent does not exist")
	}
	agencyToTransfer := agency{}
	err = json.Unmarshal(agencysAsBytes, &agencyToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}

	agencyToTransfer.Submittime = submittime
	agencyToTransfer.Flag = "1"
	agencyJSONasBytes, _ := json.Marshal(agencyToTransfer)
	err = stub.PutState(patentid, agencyJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end transferPatent (success)")
	return shim.Success(nil)
}

func (t *SimpleChaincode) refuse(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting name of the patent to query")
	}
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	patentid := args[0]
	reason := args[1]

	agencysAsBytes, err := stub.GetState(patentid)
	if err != nil {
		return shim.Error("Failed to get patent:" + err.Error())
	} else if agencysAsBytes == nil {
		return shim.Error("Patent does not exist")
	}
	agencyToTransfer := agency{}
	err = json.Unmarshal(agencysAsBytes, &agencyToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}

	agencyToTransfer.Reason = reason
	agencyToTransfer.Flag = "2"
	agencyJSONasBytes, _ := json.Marshal(agencyToTransfer)
	err = stub.PutState(patentid, agencyJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end transferPatent (success)")
	return shim.Success(nil)
}

func (t *SimpleChaincode) readAgency(stub shim.ChaincodeStubInterface, args []string) pb.Response {
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
