package decode

import (
	"fmt"

	eth_abi "github.com/ethereum/go-ethereum/accounts/abi"
)

func getInputArguments(abi eth_abi.ABI, name string, data []byte) (eth_abi.Arguments, error) {

	var args eth_abi.Arguments
	if method, ok := abi.Methods[name]; ok {
		if len(data)%32 != 0 {
			return nil, fmt.Errorf("abi: improperly formatted output: %q - Bytes: %+v", data, data)
		}
		args = method.Inputs
	}
	if event, ok := abi.Events[name]; ok {
		args = event.Inputs
	}
	if args == nil {
		return nil, fmt.Errorf("abi: cannot locate named method: %s", name)
	}
	return args, nil
}

func getOutputArguments(abi eth_abi.ABI, name string, data []byte) (eth_abi.Arguments, error) {

	var args eth_abi.Arguments
	if method, ok := abi.Methods[name]; ok {
		if len(data)%32 != 0 {
			return nil, fmt.Errorf("abi: improperly formatted output: %q - Bytes: %+v", data, data)
		}
		args = method.Outputs
	}
	if args == nil {
		return nil, fmt.Errorf("abi: cannot locate named method: %s", name)
	}
	return args, nil
}

func decodeTxInputData(abi eth_abi.ABI, name string, data []byte) (v map[string]interface{}, e error) {

	args, err := getInputArguments(abi, name, data)
	if err != nil {
		return nil, err
	}
	v = make(map[string]interface{})
	e = args.UnpackIntoMap(v, data)
	return
}

func decodeTxOutputData(abi eth_abi.ABI, name string, data []byte) (v map[string]interface{}, e error) {

	args, err := getOutputArguments(abi, name, data)
	if err != nil {
		return nil, err
	}
	v = make(map[string]interface{})
	e = args.UnpackIntoMap(v, data)
	return
}

func GetInputArguments(abi eth_abi.ABI, name string, data []byte) (eth_abi.Arguments, error) {

	return getInputArguments(abi, name, data)
}

func GetOutputArguments(abi eth_abi.ABI, name string, data []byte) (eth_abi.Arguments, error) {

	return getOutputArguments(abi, name, data)
}

func DecodeTxOutputDataByMethodName(abi eth_abi.ABI, name string, data []byte) (map[string]interface{}, error) {

	return decodeTxOutputData(abi, name, data)
}

func DecodeTxInputDataByMethodName(abi eth_abi.ABI, name string, data []byte) (map[string]interface{}, error) {

	return decodeTxInputData(abi, name, data)
}

func DecodeTxInputDataByOriginData(abi eth_abi.ABI, data []byte) (v map[string]interface{}, e error) {
	method, err := abi.MethodById(data)
	if err != nil {
		return nil, err
	}
	if len(data[4:])%32 != 0 {
		return nil, fmt.Errorf("abi: invalid data length to decode: %s", data)
	}
	v = make(map[string]interface{})
	e = method.Inputs.UnpackIntoMap(v, data[4:])
	return
}

func DecodeTxOutputDataByOriginData(abi eth_abi.ABI, data []byte) (v map[string]interface{}, e error) {
	method, err := abi.MethodById(data)
	if err != nil {
		return nil, err
	}
	if len(data[4:])%32 != 0 {
		return nil, fmt.Errorf("abi: invalid data length to decode: %s", data)
	}
	v = make(map[string]interface{})
	e = method.Outputs.UnpackIntoMap(v, data[4:])
	return
}
