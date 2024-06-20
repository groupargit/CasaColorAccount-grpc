package utilsvalidate

import (
	"fmt"
	"strings"

	protoAccount "github.com/groupargit/casacoloraccount-grpc/rpc/types/account/v1alpha1"
)

//var regexName string = `^[a-z][-_0-9a-z]+`

func ValidateDataDeletePayment(request *protoAccount.AccountLoginRequest) error {

	stringArray := []string{}

	// organizationId := request.GetOrganizationId()
	// cardId := request.GetCardId()

	// if organizationId == "" {
	// 	stringArray = append(stringArray, "organizationId is required")
	// }

	// if cardId == "" {
	// 	stringArray = append(stringArray, "cardId is required")
	// }

	if len(stringArray) > 0 {
		justString := strings.Join(stringArray, " ")
		return fmt.Errorf(justString)
	}
	return nil

}

// func ValidateDataListPayment(request *protoPayment.ListPaymentRequest) error {

// 	stringArray := []string{}

// 	organizationId := request.GetOrganizationId()
// 	if organizationId == "" {
// 		stringArray = append(stringArray, "organizationId is required")
// 	}

// 	customerId := request.GetCustomerId()
// 	if customerId == "" {
// 		stringArray = append(stringArray, "customerId is required")
// 	}

// 	if len(stringArray) > 0 {
// 		justString := strings.Join(stringArray, " ")
// 		return fmt.Errorf(justString)
// 	}
// 	return nil

// }
