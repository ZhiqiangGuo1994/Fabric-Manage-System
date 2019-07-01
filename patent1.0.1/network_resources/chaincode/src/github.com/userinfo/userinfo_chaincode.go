package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct{}

type userinfo struct {
	Uuid         string `json:"uuid"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Loginname    string `json:"loginname"`
	Workspace    string `json:"workspace"`
	Usertype     string `json:"usertype"`
	Telephone    string `json:"telephone"`
	Email        string `json:"email"`
	Organization string `json:"organization"`
	Mspid        string `json:"mspid"`
	Delflag      string `json:"delflag"`
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)
	if function == "initUserinfo" {
		return t.initUserinfo(stub, args)
	} else if function == "getCreator" {
		return t.getCreator(stub)
	} else if function == "readUserinfo" {
		return t.readUserinfo(stub, args)
	} else if function == "alter" {
		return t.alter(stub, args)
	} else if function == "delUser" {
		return t.delUser(stub, args)
	}
	return shim.Success(nil)
}

func (t *SimpleChaincode) getCreator(stub shim.ChaincodeStubInterface) pb.Response {
	userinfoAsBytes, err := stub.GetCreator()
	if err != nil {
		shim.Error("Failed to get userinfo: " + err.Error())
	} else if userinfoAsBytes == nil {
		shim.Error("Patent does not exist")
	}
	return shim.Success(userinfoAsBytes)
}

func (t *SimpleChaincode) initUserinfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 10 {
		return shim.Error("Incorrect number of arguments. Expecting 10")
	}

	uuid := args[0]
	username := args[1]
	password := args[2]
	loginname := args[3]
	workspace := args[4]
	usertype := args[5]
	telephone := args[6]
	email := args[7]
	organization := args[8]
	mspid := args[9]
	delflag := "0"
	userinfoAsBytes, err := stub.GetState(uuid)
	if err != nil {
		return shim.Error("Failed to get userinfo: " + err.Error())
	} else if userinfoAsBytes != nil {
		return shim.Error("Patent has exist")
	}
	userinfo := &userinfo{uuid, username, password, loginname, workspace, usertype, telephone, email, organization, mspid, delflag}
	userinfoJSONasBytes, err := json.Marshal(userinfo)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(uuid, userinfoJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *SimpleChaincode) readUserinfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var uuid, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the patent to query")
	}

	uuid = args[0]
	valAsbytes, err := stub.GetState(uuid) //get the patent from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + uuid + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Patent does not exist: " + uuid + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

func (t *SimpleChaincode) alter(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 10 {
		return shim.Error("Incorrect number of arguments. Expecting 10")
	}
	uuid := args[0]
	username := args[1]
	password := args[2]
	loginname := args[3]
	workspace := args[4]
	usertype := args[5]
	telephone := args[6]
	email := args[7]
	organization := args[8]
	mspid := args[9]

	userinfoAsBytes, err := stub.GetState(uuid)
	if err != nil {
		shim.Error("Failed to get userinfo: " + err.Error())
	} else if userinfoAsBytes == nil {
		shim.Error("Patent does not exist")
	}
	userinfoToTransfer := userinfo{}
	err = json.Unmarshal(userinfoAsBytes, &userinfoToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	userinfoToTransfer.Username = username
	userinfoToTransfer.Password = password
	userinfoToTransfer.Loginname = loginname
	userinfoToTransfer.Workspace = workspace
	userinfoToTransfer.Usertype = usertype
	userinfoToTransfer.Telephone = telephone
	userinfoToTransfer.Email = email
	userinfoToTransfer.Organization = organization
	userinfoToTransfer.Mspid = mspid

	userinfoJSONasBytes, _ := json.Marshal(userinfoToTransfer)
	err = stub.PutState(uuid, userinfoJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *SimpleChaincode) delUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	uuid := args[0]
	userinfoAsBytes, err := stub.GetState(uuid)
	if err != nil {
		shim.Error("Failed to get userinfo: " + err.Error())
	} else if userinfoAsBytes == nil {
		shim.Error("Patent does not exist")
	}
	userinfoToTransfer := userinfo{}
	err = json.Unmarshal(userinfoAsBytes, &userinfoToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}

	userinfoToTransfer.Delflag = "1"

	userinfoJSONasBytes, _ := json.Marshal(userinfoToTransfer)
	err = stub.PutState(uuid, userinfoJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *SimpleChaincode) userHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	uuid := args[0]
	historyIterator, err := stub.GetHistoryForKey(uuid)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer historyIterator.Close()
	var buffer bytes.Buffer
	buffer.WriteString("[")

	//	bArrayMemberAlreadyWritten := false
	//	for historyIterator.HasNext() {
	//		queryResponse, err := historyIterator.Next()
	//		if err != nil {
	//			return shim.Error(err.Error())
	//		}
	//		// Add a comma before array members, suppress it for the first array member
	//		if bArrayMemberAlreadyWritten == true {
	//			buffer.WriteString(",")
	//		}
	//		buffer.WriteString("{\"Key\":")
	//		buffer.WriteString("\"")
	//		buffer.WriteString(queryResponse.Key)
	//		buffer.WriteString("\"")

	//		buffer.WriteString(", \"Record\":")
	//		// Record is a JSON object, so we write as-is
	//		buffer.WriteString(string(queryResponse.Value))
	//		buffer.WriteString("}")
	//		bArrayMemberAlreadyWritten = true
	//	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Print("Error starting Simple chaincode: %s", err)
	}
}
