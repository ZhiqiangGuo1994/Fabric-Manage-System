package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pt "github.com/hyperledger/fabric/protos/peer"
)

// PatentChaincode example simple Chaincode implementation
type PatentChaincode struct {
}

type patent struct {
	Uuid string `json:"uuid"`

	Patentname      string `json:"patentname"`
	Applicationdate string `json:"applicationdate"`

	Applicationid string `json:"applicationid"`

	Inventorname      string `json:"inventorname"`
	Inventorsex       string `json:"inventorsex"`
	Inventortelephone string `json:"inventortelephone"`
	Inventoremail     string `json:"inventoremail"`

	Patentapplication string `json:"patentapplication"`
	Instructionmanual string `json:"instructionmanual"`
	Instructimage     string `json:"instructimage"`
	Claim             string `json:"claim"`
	Summary           string `json:"summary"`
	Claimimage        string `json:"claimimage"`
	Summit            string `json:"submit"`
	Summitflag        string `json:"submitflag"`
	Markleroot        string `json:"markleroot"`
	Flagdel           string `json:"flagdel"`
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {
	//Create a new Smart Contract
	err := shim.Start(new(PatentChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *PatentChaincode) Init(stub shim.ChaincodeStubInterface) pt.Response {
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *PatentChaincode) Invoke(stub shim.ChaincodeStubInterface) pt.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	//Handle different functions
	if function == "initPatent" {
		//create a new patent
		return t.initPatent(stub, args)
	} else if function == "transferPatent" {
		return t.transferPatent(stub, args)
	} else if function == "refuse" {
		return t.refuse(stub, args)
	} else if function == "readPatent" {
		return t.readPatent(stub, args)
	} else if function == "delete" {
		return t.delete(stub, args)
	} else if function == "queryByApplicationId" {
		return t.queryByApplicationId(stub, args)
	}

	return shim.Error("Received unknown function invocation")
}

// ============================================================
// initPatent - create a new patent, store into chaincode state
// ============================================================
func (t *PatentChaincode) initPatent(stub shim.ChaincodeStubInterface, args []string) pt.Response {
	var err error

	if len(args) != 18 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init patent")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	uuid := args[0]

	//专利基本信息
	patentname := args[1]
	applicationdate := args[2]
	//申请人信息
	applicationid := args[3]
	//发明人信息
	inventorname := args[4]
	inventorsex := args[5]
	inventortelephone := args[6]
	inventoremail := args[7]
	//专利信息
	patentapplication := args[8]
	instructionmanual := args[9]
	instructimage := args[10]
	claim := args[11]
	summary := args[12]
	claimimage := args[13]
	//授权书
	submit := args[14]
	submitflag := args[15]

	markleroot := args[16]
	//当前用户签名
	flagdel := args[17]

	// ==== Check if patent already exists ====
	patentAsBytes, err := stub.GetState(uuid)
	if err != nil {
		return shim.Error("Failed to get patent: " + err.Error())
	} else if patentAsBytes != nil {
		patentToTransfer := patent{}
		err = json.Unmarshal(patentAsBytes, &patentToTransfer) //unmarshal it aka JSON.parse()
		if err != nil {
			return shim.Error(err.Error())
		}
		patentToTransfer.Patentname = patentname
		patentToTransfer.Applicationdate = applicationdate
		patentToTransfer.Applicationid = applicationid

		patentToTransfer.Inventorname = inventorname
		patentToTransfer.Inventorsex = inventorsex
		patentToTransfer.Inventortelephone = inventortelephone
		patentToTransfer.Inventoremail = inventoremail
		patentToTransfer.Patentapplication = patentapplication
		patentToTransfer.Instructionmanual = instructionmanual
		patentToTransfer.Instructimage = instructimage
		patentToTransfer.Claim = claim
		patentToTransfer.Summary = summary
		patentToTransfer.Claimimage = claimimage
		patentToTransfer.Summit = submit
		patentToTransfer.Summitflag = submitflag

		patentToTransfer.Markleroot = markleroot
		patentToTransfer.Flagdel = flagdel

		patentJSONasBytes, _ := json.Marshal(patentToTransfer)
		err = stub.PutState(uuid, patentJSONasBytes) //rewrite the patent
		if err != nil {
			return shim.Error(err.Error())
		}

		fmt.Println("- end transferPatent (success)")
		return shim.Success(nil)
	} else {
		// ==== Create patent object and marshal to JSON ====
		patent := &patent{uuid, patentname, applicationdate, applicationid, inventorname, inventorsex, inventortelephone, inventoremail, patentapplication, instructionmanual, instructimage, claim, summary, claimimage, submit, submitflag, markleroot, flagdel}
		patentJSONasBytes, err := json.Marshal(patent)
		if err != nil {
			return shim.Error(err.Error())
		}

		// === Save patent to state ===
		err = stub.PutState(uuid, patentJSONasBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

		return shim.Success(nil)
	}

}

// ===========================================================
// transfer a patent by setting a new owner name on the patent
// ===========================================================
func (t *PatentChaincode) transferPatent(stub shim.ChaincodeStubInterface, args []string) pt.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments.")
	}

	uuid := args[0]
	fmt.Println("- start transferPatent ", uuid)

	patentsAsBytes, err := stub.GetState(uuid)
	if err != nil {
		return shim.Error("Failed to get patent:" + err.Error())
	} else if patentsAsBytes == nil {
		return shim.Error("Patent does not exist")
	}

	patentToTransfer := patent{}
	err = json.Unmarshal(patentsAsBytes, &patentToTransfer) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	patentToTransfer.Summitflag = "1" //0代表未提交，１代表已提交

	patentJSONasBytes, _ := json.Marshal(patentToTransfer)
	err = stub.PutState(uuid, patentJSONasBytes) //rewrite the patent
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end transferPatent (success)")
	return shim.Success(nil)
}

// ===============================================
// readPatent - read a patent from chaincode state
// ===============================================
func (t *PatentChaincode) readPatent(stub shim.ChaincodeStubInterface, args []string) pt.Response {
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

// ==================================================
// delete - remove a patent key/value pair from state
// ==================================================
func (t *PatentChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pt.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	uuid := args[0]

	patentsAsBytes, err := stub.GetState(uuid)
	if err != nil {
		return shim.Error("Failed to get patent:" + err.Error())
	} else if patentsAsBytes == nil {
		return shim.Error("Patent does not exist")
	}

	patentToTransfer := patent{}
	err = json.Unmarshal(patentsAsBytes, &patentToTransfer) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	if patentToTransfer.Summit == "1" {
		return shim.Error("This patent has submitted .")
	} else {
		patentToTransfer.Flagdel = "1" //1代表已经删除，０代表未删除
		patentJSONasBytes, _ := json.Marshal(patentToTransfer)
		err = stub.PutState(uuid, patentJSONasBytes) //rewrite the patent
		if err != nil {
			return shim.Error(err.Error())
		}

		fmt.Println("- end transferPatent (success)")
		return shim.Success(nil)
	}
}

// ==================================================
// refuse - refuse a patent for submitflag is "2"
// ==================================================
func (t *PatentChaincode) refuse(stub shim.ChaincodeStubInterface, args []string) pt.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	uuid := args[0]

	patentsAsBytes, err := stub.GetState(uuid)
	if err != nil {
		return shim.Error("Failed to get patent:" + err.Error())
	} else if patentsAsBytes == nil {
		return shim.Error("Patent does not exist")
	}

	patentToTransfer := patent{}
	err = json.Unmarshal(patentsAsBytes, &patentToTransfer) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	patentToTransfer.Summitflag = "2" //1代表已经删除，０代表未删除
	patentJSONasBytes, _ := json.Marshal(patentToTransfer)
	err = stub.PutState(uuid, patentJSONasBytes) //rewrite the patent
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end transferPatent (success)")
	return shim.Success(nil)

}

//用户查询个人申请的专利
func (t *PatentChaincode) queryByApplicationId(stub shim.ChaincodeStubInterface, args []string) pt.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	applicationid := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"applicationid\":\"%s\",\"flagdel\":\"0\"}}", applicationid)
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
