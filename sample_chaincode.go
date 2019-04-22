package main 
import (   
	"fmt"       
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)
type SampleChaincode struct {}
func (t *SampleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {    
	return nil, nil
}
func (t *SampleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {    
	return nil, nil
}
func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {    
	return nil, nil
}
func CreateLoanApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {    
	fmt.Println("Entering CreateLoanApplication")    
	if len(args) < 2 {        
		fmt.Println("Invalid number of args")        
		return nil, errors.New("Expected atleast two arguments for loan application creation")    
	}    
	var loanApplicationId = args[0]    
	var loanApplicationInput = args[1]    
	//TODO: Include schema validation here    
	err := stub.PutState(loanApplicationId, []byte(loanApplicationInput))    
	if err != nil {        
		fmt.Println("Could not save loan application to ledger", err)        
		return nil, err    
	}    
	fmt.Println("Successfully saved loan application")    
	return []byte(loanApplicationInput), nil
}
func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {    
	fmt.Println("Entering Invoke")    
	return nil, errors.New("unauthorized user")
}
func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {    
	fmt.Println("Entering Invoke")    
	ubytes, _ := stub.ReadCertAttribute("username")    
	rbytes, _ := stub.ReadCertAttribute("role")    
	username := string(ubytes)    
	role := string(rbytes)    
	if role != "Bank_Admin" {        
		return nil, errors.New("caller with " + username + " and role " + role + " does not have         access to invoke CreateLoanApplication")    
	}    return nil, nil
}