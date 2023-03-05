package internal

// Contains code for the parser from input data format to internal document
// format.

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	ojson "gitlab.com/c0b/go-ordered-json"
)

const inputTokensDefaultCount = 10

func ParseInput(data []byte) (*InputDoc, error) {
	inputDoc := NewInputDoc()
	om := ojson.NewOrderedMap()

	err := json.Unmarshal(data, om)
	if err != nil {
		return nil, err
	}

	iter := om.EntriesIter()
	for {
		pair, ok := iter()
		if !ok {
			break
		}

		switch oType := reflect.TypeOf(pair.Value).Kind(); oType {
		case reflect.String:
			token, err := NewVarToken(pair.Key, pair.Value)
			if err != nil {
				return nil, err
			}
			inputDoc.Tokens = append(inputDoc.Tokens, token)

		// function definition
		case reflect.Slice:
			sliceValues := reflect.ValueOf(pair.Value)
			commands := make([]*CmdDef, 0, sliceValues.Len())

			var isInitFunc bool
			if strings.Compare(pair.Key, "init") == 0 {
				isInitFunc = true
				inputDoc.Meta.InitFuncIsPresent = true
			}

			for i := 0; i < sliceValues.Len(); i++ {
				sliceItemPtr := sliceValues.Index(i).Elem()
				funcDefinition, ok := sliceItemPtr.Interface().(*ojson.OrderedMap)
				if !ok {
					return nil, fmt.Errorf(
						"internal error: function definition is not a (*ojson.OrderedMap) type",
					)
				}

				cmdDef, err := parseCmd(funcDefinition.EntriesIter())
				if err != nil {
					return nil, fmt.Errorf("error parsing cmd definition: %v", err)
				}

				if isInitFunc {
					_, isDefault := IsDefaultCmd(cmdDef.Call)
					if !isDefault {
						key := strings.TrimPrefix(cmdDef.Call, "#")
						inputDoc.Meta.InitRequired[key] = struct{}{}
					}
				}

				commands = append(commands, cmdDef)
			}
			token := NewFuncToken(pair.Key, commands)
			inputDoc.Tokens = append(inputDoc.Tokens, token)

		default:
			fmt.Printf("warning: unexpected object %v (%v)\n", pair.Key, oType)
		}
	}

	return inputDoc, nil
}

func NewCmdDef() *CmdDef {
	return &CmdDef{
		OperandRefs: make([]*ArgDef, 0, 3),
	}
}

func parseCmd(iter func() (*ojson.KVPair, bool)) (*CmdDef, error) {
	ErrUnexpectedValue := errors.New("unexpected type of value")
	badValueMsg := "bad %v value : %v:%v\n"
	cmd := NewCmdDef()

	for {
		cmdItemPair, ok := iter()
		if !ok {
			break
		}

		switch typeOfKey(cmdItemPair.Key) {
		case CmdKeyTypeCmd:
			cmd.Call, ok = cmdItemPair.Value.(string)
			if !ok {
				fmt.Printf(badValueMsg, "cmd", cmdItemPair.Key, cmdItemPair.Value)
				return nil, ErrUnexpectedValue
			}

		case CmdKeyTypeID:
			cmd.Target, ok = cmdItemPair.Value.(string)
			if !ok {
				fmt.Printf(badValueMsg, "id", cmdItemPair.Key, cmdItemPair.Value)
				return nil, ErrUnexpectedValue
			}

		case CmdKeyTypeValue:
			ref, ok := cmdItemPair.Value.(string)
			if ok {
				cmd.OperandRefs = append(cmd.OperandRefs, NewArgValueRef(ref))
				continue
			}

			_, ok = cmdItemPair.Value.(json.Number)
			if ok {
				vf, err := getVarValue(cmdItemPair.Value)
				if err != nil {
					fmt.Printf(badValueMsg, "value", cmdItemPair.Key, cmdItemPair.Value)
					return nil, err
				}
				cmd.OperandRefs = append(cmd.OperandRefs, NewArgValue(vf))
				continue
			}

		case CmdKeyTypeOperand:
			val, ok := cmdItemPair.Value.(string)
			if !ok {
				fmt.Printf(badValueMsg, "operand value", cmdItemPair.Key, cmdItemPair.Value)
				return nil, ErrUnexpectedValue
			}
			arg := &ArgDef{
				ValType:  ArgTypeOperandRef,
				ValueRef: val,
			}
			cmd.OperandRefs = append(cmd.OperandRefs, arg)

		default:
			return nil, fmt.Errorf("unexpected cmd key: %v", cmdItemPair.Key)
		}
	}
	return cmd, nil
}

type CmdKeyType int

const (
	CmdKeyTypeUndefined CmdKeyType = iota
	CmdKeyTypeCmd
	CmdKeyTypeID
	CmdKeyTypeValue
	CmdKeyTypeOperand
)

var (
	validCmdCmd      = regexp.MustCompile(`^cmd$`)
	validCmdID       = regexp.MustCompile(`^id$`)
	validCmdValueRef = regexp.MustCompile(`^value[0-9]*$`)
	validCmdOperand  = regexp.MustCompile(`^operand[0-9]*$`)
)

func typeOfKey(key string) CmdKeyType {
	if validCmdCmd.MatchString(key) {
		return CmdKeyTypeCmd
	}
	if validCmdID.MatchString(key) {
		return CmdKeyTypeID
	}
	if validCmdValueRef.MatchString(key) {
		return CmdKeyTypeValue
	}
	if validCmdOperand.MatchString(key) {
		return CmdKeyTypeOperand
	}
	return CmdKeyTypeUndefined
}
