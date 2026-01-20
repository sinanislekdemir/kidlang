package interpreter

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// parseStackIndex extracts stack name and index from "name[index]" or "name(index)"
// Returns: stackName, index, isStack
func parseStackIndex(text string) (string, string, bool) {
	// Match pattern: name[index] or name(index)
	re := regexp.MustCompile(`^(.+?)[\[\(](.+?)[\]\)]$`)
	matches := re.FindStringSubmatch(text)
	if len(matches) == 3 {
		return strings.TrimSpace(matches[1]), strings.TrimSpace(matches[2]), true
	}
	return "", "", false
}

// getNumberArgument handles getting a number argument that might be split by +/- sign
// Returns the float value and number of tokens consumed (1 or 2)
func getNumberArgument(arguments []VariableBox, index int) (float64, int, bool) {
	if index >= len(arguments) {
		return 0, 0, false
	}

	arg := arguments[index]

	// Direct number
	if arg.VariableType == TYPE_INTEGER || arg.VariableType == TYPE_FLOAT {
		return arg.ToFloat(), 1, true
	}

	// Sign followed by number (e.g., "- 5")
	if (arg.String == "-" || arg.String == "+") && index+1 < len(arguments) {
		nextArg := arguments[index+1]
		if nextArg.VariableType == TYPE_INTEGER || nextArg.VariableType == TYPE_FLOAT {
			val := nextArg.ToFloat()
			if arg.String == "-" {
				val = -val
			}
			return val, 2, true
		}
	}

	return 0, 0, false
}

// EvaluateInlineFunctions evaluates the inline functions in the arguments.
// Does not overwrite the original arguments.
func evaluateInlineFunctions(arguments []VariableBox) []VariableBox {
	result := make([]VariableBox, 0, len(arguments))
	for index := 0; index < len(arguments); index++ {
		if arguments[index].VariableType != TYPE_STRING {
			result = append(result, arguments[index])
			continue
		}
		if strings.EqualFold(arguments[index].String, getTranslation("RANDOM")) {
			result = append(result, VariableBox{
				VariableType: TYPE_INTEGER,
				Integer:      rand.Int63(),
			})
			continue
		}
		if strings.EqualFold(arguments[index].String, getTranslation("NOW")) {
			now := time.Now()
			result = append(result, VariableBox{
				VariableType: TYPE_STRING,
				String:       now.Format("Monday, January 2, 2006 15:03:05"),
			})
			continue
		}
		if strings.EqualFold(arguments[index].String, getTranslation("SQRT")) {
			if val, consumed, ok := getNumberArgument(arguments, index+1); ok {
				res := math.Sqrt(val)
				result = append(result, VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        res,
				})
				index += consumed
				continue
			}
			result = append(result, arguments[index])
			continue
		}
		if strings.EqualFold(arguments[index].String, getTranslation("ABS")) {
			if val, consumed, ok := getNumberArgument(arguments, index+1); ok {
				res := math.Abs(val)
				result = append(result, VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        res,
				})
				index += consumed
				continue
			}
			result = append(result, arguments[index])
			continue
		}
		if strings.EqualFold(arguments[index].String, getTranslation("SQR")) {
			if val, consumed, ok := getNumberArgument(arguments, index+1); ok {
				res := val * val
				result = append(result, VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        res,
				})
				index += consumed
				continue
			}
			result = append(result, arguments[index])
			continue
		}
		if strings.EqualFold(arguments[index].String, "SIN") {
			if val, consumed, ok := getNumberArgument(arguments, index+1); ok {
				res := math.Sin(val)
				result = append(result, VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        res,
				})
				index += consumed
				continue
			}
			result = append(result, arguments[index])
			continue
		}
		if strings.EqualFold(arguments[index].String, "COS") {
			if val, consumed, ok := getNumberArgument(arguments, index+1); ok {
				res := math.Cos(val)
				result = append(result, VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        res,
				})
				index += consumed
				continue
			}
			result = append(result, arguments[index])
			continue
		}
		if strings.EqualFold(arguments[index].String, "TAN") {
			if val, consumed, ok := getNumberArgument(arguments, index+1); ok {
				res := math.Tan(val)
				result = append(result, VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        res,
				})
				index += consumed
				continue
			}
			result = append(result, arguments[index])
			continue
		}
		if strings.EqualFold(arguments[index].String, "LOG") {
			if val, consumed, ok := getNumberArgument(arguments, index+1); ok {
				res := math.Log(val)
				result = append(result, VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        res,
				})
				index += consumed
				continue
			}
			result = append(result, arguments[index])
			continue
		}
		if strings.EqualFold(arguments[index].String, "ASIN") {
			if val, consumed, ok := getNumberArgument(arguments, index+1); ok {
				res := math.Asin(val)
				result = append(result, VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        res,
				})
				index += consumed
				continue
			}
			result = append(result, arguments[index])
			continue
		}
		if strings.EqualFold(arguments[index].String, "ACOS") {
			if val, consumed, ok := getNumberArgument(arguments, index+1); ok {
				res := math.Acos(val)
				result = append(result, VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        res,
				})
				index += consumed
				continue
			}
			result = append(result, arguments[index])
			continue
		}
		result = append(result, arguments[index])
	}
	return result
}

func evaluateMulDiv(arguments []VariableBox) ([]VariableBox, error) {
	result := make([]VariableBox, 0, len(arguments))
	for index := 0; index < len(arguments); index++ {
		if arguments[index].VariableType != TYPE_STRING {
			result = append(result, arguments[index])
			continue
		}
		if arguments[index].String != "*" && arguments[index].String != "/" {
			result = append(result, arguments[index])
			continue
		}

		if index == 0 || index == len(arguments)-1 {
			return nil, fmt.Errorf("invalid expression")
		}

		if arguments[index].String == "*" {
			res, err := result[len(result)-1].Mul(arguments[index+1])
			if err != nil {
				return nil, err
			}
			result[len(result)-1] = res
			index += 1
		} else {
			res, err := result[len(result)-1].Div(arguments[index+1])
			if err != nil {
				return nil, err
			}
			result[len(result)-1] = res
			index += 1
		}
	}
	return result, nil
}

func evalSumSubModXor(arguments []VariableBox) ([]VariableBox, error) {
	result := make([]VariableBox, 0, len(arguments))
	for index := 0; index < len(arguments); index++ {
		if arguments[index].VariableType != TYPE_STRING {
			result = append(result, arguments[index])
			continue
		}
		if arguments[index].String != "+" && arguments[index].String != "-" && arguments[index].String != "%" && arguments[index].String != "^" {
			result = append(result, arguments[index])
			continue
		}

		// Handle unary minus/plus at start
		if index == 0 {
			if arguments[index].String == "-" && index+1 < len(arguments) {
				// Unary minus: negate the next value
				nextVal := arguments[index+1]
				if nextVal.VariableType == TYPE_INTEGER {
					result = append(result, VariableBox{
						VariableType: TYPE_INTEGER,
						Integer:      -nextVal.Integer,
					})
					index++
					continue
				} else if nextVal.VariableType == TYPE_FLOAT {
					result = append(result, VariableBox{
						VariableType: TYPE_FLOAT,
						Float:        -nextVal.Float,
					})
					index++
					continue
				}
			} else if arguments[index].String == "+" && index+1 < len(arguments) {
				// Unary plus: just use the next value as-is
				result = append(result, arguments[index+1])
				index++
				continue
			}
			return nil, fmt.Errorf("invalid expression")
		}

		if index == len(arguments)-1 {
			return nil, fmt.Errorf("invalid expression")
		}

		if arguments[index].String == "+" {
			res := result[len(result)-1].Sum(arguments[index+1])
			result[len(result)-1] = res
			index++
		} else if arguments[index].String == "-" {
			res, err := result[len(result)-1].Sub(arguments[index+1])
			if err != nil {
				return nil, err
			}
			result[len(result)-1] = res
			index++
		} else if arguments[index].String == "%" {
			res, err := result[len(result)-1].Mod(arguments[index+1])
			if err != nil {
				return nil, err
			}
			result[len(result)-1] = res
			index++
		} else {
			res, err := result[len(result)-1].Xor(arguments[index+1])
			if err != nil {
				return nil, err
			}
			result[len(result)-1] = res
			index++
		}
	}
	return result, nil
}

// processArguments processes and evaluates a list of VariableBox arguments.
// It performs the following steps:
// 1. Processes the arguments using the provided memory.
// 2. Evaluates the processed arguments.
// 3. Evaluates any dynamic boxes within the arguments using the provided memory.
// 4. Converts the arguments to their string representations.
//
// Parameters:
// - memory: The KLMemory instance used for processing and evaluating arguments.
// - arguments: A slice of VariableBox representing the arguments to be processed.
//
// Returns:
// - A slice of VariableBox containing the processed and evaluated arguments.
// - An error if any step in the processing or evaluation fails.
func processArguments(memory KLMemory, arguments []VariableBox) ([]VariableBox, error) {
	processedArguments, err := prepareArguments(memory, arguments)
	if err != nil {
		return nil, err
	}
	processedArguments, err = evalMathematicalOperations(processedArguments)
	if err != nil {
		return nil, err
	}
	processedArguments = stringifyArguments(processedArguments)
	return processedArguments, nil
}

func stringifyArguments(arguments []VariableBox) []VariableBox {
	// Filter out TYPE_UNKNOWN (placeholders)
	filtered := make([]VariableBox, 0, len(arguments))
	for _, val := range arguments {
		if val.VariableType != TYPE_UNKNOWN {
			filtered = append(filtered, val)
		}
	}

	isString := false
	for _, val := range filtered {
		if val.VariableType == TYPE_STRING {
			isString = true
			break
		}
	}
	if isString {
		stringList := make([]string, 0, len(filtered))
		for _, val := range filtered {
			stringList = append(stringList, val.ToString())
		}
		return []VariableBox{
			{
				VariableType: TYPE_STRING,
				String:       strings.Join(stringList, " "),
			},
		}
	}

	return filtered
}

func evalMathematicalOperations(arguments []VariableBox) ([]VariableBox, error) {
	afterInlines := evaluateInlineFunctions(arguments)
	afterMulDiv, err := evaluateMulDiv(afterInlines)
	if err != nil {
		return nil, err
	}
	afterSubSumModXor, err := evalSumSubModXor(afterMulDiv)
	if err != nil {
		return nil, err
	}

	return afterSubSumModXor, nil
}

// prepareArguments happens during the runtime. That is why this is a separate function. Not handled in the GetArguments.
// GetArguments is a compile time function.
// Important: This function is not thread safe. Do not use it in concurrent environments.
// Important: This function is not safe for recursive calls. Do not use it in recursive environments.
// Important: This function does not modify the original arguments as it's possible to run the same statement multiple times.
//
//	Consider goto. If you have a goto statement, you can run the same statement multiple times.
func prepareArguments(memory KLMemory, arguments []VariableBox) ([]VariableBox, error) {
	result := make([]VariableBox, len(arguments))
	boxName := getTranslation(BOX)
	fileName := getTranslation(FILE)

OUTER:
	for index := 0; index < len(arguments); index++ {
		// Check if this is "box" or "file" followed by another argument
		if index+1 < len(arguments) && arguments[index].VariableType == TYPE_STRING {
			if strings.ToUpper(arguments[index].String) == boxName || strings.ToUpper(arguments[index].String) == fileName {
				// Get the string representation of next argument
				nextStr := arguments[index+1].String
				if arguments[index+1].VariableType == TYPE_INTEGER {
					nextStr = strconv.FormatInt(arguments[index+1].Integer, 10)
				} else if arguments[index+1].VariableType == TYPE_FLOAT {
					nextStr = strconv.FormatFloat(arguments[index+1].Float, 'f', -1, 64)
				}

				combinedName := arguments[index].String + " " + nextStr
				if val, exists := memory[strings.ToUpper(combinedName)]; exists {
					result[index] = val
					// Mark next element as processed
					index++
					result[index] = VariableBox{VariableType: TYPE_UNKNOWN}
					continue
				}
			}
		}

		// Check if this argument (as string) exists in memory first
		argStr := arguments[index].String
		if arguments[index].VariableType == TYPE_INTEGER {
			argStr = strconv.FormatInt(arguments[index].Integer, 10)
		} else if arguments[index].VariableType == TYPE_FLOAT {
			argStr = strconv.FormatFloat(arguments[index].Float, 'f', -1, 64)
		}
		if val, exists := memory[strings.ToUpper(argStr)]; exists {
			result[index] = val
			continue
		}

		switch arguments[index].VariableType {
		case TYPE_REFERENCE:
			value, err := Resolve(memory, arguments[index].String)
			if err != nil {
				return nil, err
			}
			result[index].VariableType = value.VariableType
			result[index].Integer = value.Integer
			result[index].String = value.String
			result[index].Float = value.Float
		case TYPE_STRING:
			answerKw := ADDRESS_ANSWER

			if _, exists := memory[ADDRESS_LANGUAGE]; exists {
				answerKw = getTranslation(ADDRESS_ANSWER)
			}

			if strings.EqualFold(arguments[index].String, answerKw) {
				answer, err := Resolve(memory, answerKw)
				if err != nil {
					return nil, err
				}
				result[index].VariableType = answer.VariableType
				result[index].Integer = answer.Integer
				result[index].String = answer.String
				result[index].Float = answer.Float
				result[index].Bool = answer.Bool
				continue
			}

			for _, special := range Specials {
				pattern := special.Pattern
				if _, exists := memory[ADDRESS_LANGUAGE]; exists {
					pattern = getTranslation(special.Pattern)
				}
				if strings.EqualFold(pattern, arguments[index].String) {
					val := special.Function()
					result[index].VariableType = val.VariableType
					result[index].Float = val.Float
					result[index].Integer = val.Integer
					result[index].String = val.String
					result[index].Bool = val.Bool
					continue OUTER
				}
			}
			result[index].VariableType = TYPE_STRING
			result[index].String = arguments[index].String
		default:
			// do nothing
			result[index].VariableType = arguments[index].VariableType
			result[index].Float = arguments[index].Float
			result[index].Integer = arguments[index].Integer
			result[index].String = arguments[index].String
			result[index].Bool = arguments[index].Bool
		}
	}
	return result, nil
}

func stringsToArguments(memory KLMemory, args []string) []VariableBox {
	result := make([]VariableBox, 0, len(args))
	boxName := getTranslation(BOX)
	fileName := getTranslation(FILE)
	stackName := getTranslation("STACK")

	for i := 0; i < len(args); i++ {
		argument := args[i]
		if len(argument) == 0 {
			continue
		}

		// Check for "stack X[index]" or "stack X(index)" patterns
		if i+1 < len(args) && strings.ToUpper(argument) == stackName {
			// Next token might have indexing
			nextToken := args[i+1]
			_, _, isIndexed := parseStackIndex(nextToken)

			if isIndexed {
				// This is stack varname[index]
				// Return as a REFERENCE that will be resolved at execution time
				// Format: "stack name[index]" as a reference
				result = append(result, VariableBox{
					VariableType: TYPE_REFERENCE,
					String:       argument + " " + nextToken, // Keep full "stack toys[1]" format
				})
				i++ // Skip next token
				continue
			}

			// Check if next next token has indexing (e.g., "stack toys (3)" where (3) is separate)
			if i+2 < len(args) {
				thirdToken := args[i+2]
				// Check if this looks like an index: starts and ends with brackets
				if (strings.HasPrefix(thirdToken, "(") || strings.HasPrefix(thirdToken, "[")) &&
					(strings.HasSuffix(thirdToken, ")") || strings.HasSuffix(thirdToken, "]")) {
					// Combine as "stack toys(3)"
					result = append(result, VariableBox{
						VariableType: TYPE_REFERENCE,
						String:       argument + " " + nextToken + thirdToken,
					})
					i += 2 // Skip next two tokens
					continue
				}
			}

			// Not indexed or doesn't exist, treat as "stack name" reference
			combinedName := argument + " " + args[i+1]
			// For stacks, always keep as reference (don't resolve to value)
			// This allows commands like READ and WRITE to detect stack targets
			if strings.ToUpper(argument) == stackName {
				result = append(result, VariableBox{
					VariableType: TYPE_REFERENCE,
					String:       combinedName,
				})
				i++
				continue
			}
			// Check with combined name for other cases
			if val, exists := memory[strings.ToUpper(combinedName)]; exists {
				result = append(result, val)
				i++
				continue
			} else {
				result = append(result, VariableBox{
					VariableType: TYPE_REFERENCE,
					String:       combinedName,
				})
				i++
				continue
			}
		}

		// Check for "box X" or "file X" patterns - treat as reference even if not in memory yet
		if i+1 < len(args) && (strings.ToUpper(argument) == boxName || strings.ToUpper(argument) == fileName) {
			combinedName := argument + " " + args[i+1]
			if val, exists := memory[strings.ToUpper(combinedName)]; exists {
				result = append(result, val)
				i++ // Skip next token as it's part of the variable name
				continue
			} else {
				// Variable doesn't exist yet, but store as reference for runtime lookup
				result = append(result, VariableBox{
					VariableType: TYPE_REFERENCE,
					String:       combinedName,
				})
				i++ // Skip next token
				continue
			}
		}

		// Check if this argument itself has stack indexing (varname[index] without "stack" prefix)
		if stackVar, index, isIndexed := parseStackIndex(argument); isIndexed {
			// Try to find this as a stack variable
			if stackVal, exists := memory[strings.ToUpper(stackVar)]; exists && stackVal.VariableType == TYPE_STACK {
				// Resolve the index
				indexArgs := stringsToArguments(memory, []string{index})
				var indexKey string
				if len(indexArgs) > 0 {
					indexKey = indexArgs[0].ToString()
				} else {
					indexKey = index
				}
				// Get value from stack
				val := stackVal.GetFromStack(indexKey)
				result = append(result, val)
				continue
			}
		}

		if val, exists := memory[strings.ToUpper(argument)]; exists {
			result = append(result, val)
			continue
		}
		if val, err := strconv.Atoi(argument); err == nil {
			result = append(result, VariableBox{
				VariableType: TYPE_INTEGER,
				Integer:      int64(val),
			})
			continue
		}
		if val, err := strconv.ParseFloat(argument, 64); err == nil {
			result = append(result, VariableBox{
				VariableType: TYPE_FLOAT,
				Float:        val,
			})
			continue
		}
		// TODO: Add support for translations
		if strings.ToUpper(argument) == "TRUE" {
			result = append(result, VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         true,
			})
			continue
		}
		if strings.ToUpper(argument) == "FALSE" {
			result = append(result, VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         false,
			})
			continue
		}
		result = append(result, VariableBox{
			VariableType: TYPE_STRING,
			String:       argument,
		})
	}
	return result
}
