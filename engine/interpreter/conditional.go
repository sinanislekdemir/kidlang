package interpreter

import (
	"fmt"
	"slices"
	"strings"
)

func If(memory KLMemory, stack *KLStack, arguments []VariableBox) error {
	stack.ExitScope = false

	gotoLabel := ""
	andKeyword := getTranslation("AND")
	orKeyword := getTranslation("OR")

	nonOpCount := 0
	operatorCount := 0
	for _, token := range arguments {
		if slices.Contains([]string{"=", "!", ">", "<"}, token.String) {
			operatorCount++
		} else {
			nonOpCount++
		}
	}

	if nonOpCount >= 3 && operatorCount >= 1 && len(arguments) > 0 {
		lastArg := arguments[len(arguments)-1]
		if lastArg.VariableType == TYPE_STRING && !slices.Contains([]string{"=", "!", ">", "<"}, lastArg.String) {
			gotoLabel = lastArg.String
			arguments = arguments[:len(arguments)-1]
		}
	}

	conditions := [][]VariableBox{}
	logicOps := []string{}
	currentCondition := []VariableBox{}

	for _, token := range arguments {
		tokenUpper := strings.ToUpper(token.String)
		if tokenUpper == andKeyword || tokenUpper == orKeyword {
			if len(currentCondition) > 0 {
				conditions = append(conditions, currentCondition)
				if tokenUpper == andKeyword {
					logicOps = append(logicOps, "AND")
				} else {
					logicOps = append(logicOps, "OR")
				}
				currentCondition = []VariableBox{}
			}
		} else {
			currentCondition = append(currentCondition, token)
		}
	}
	if len(currentCondition) > 0 {
		conditions = append(conditions, currentCondition)
	}

	results := []bool{}
	for _, condition := range conditions {
		result, err := evaluateCondition(memory, condition)
		if err != nil {
			return err
		}
		results = append(results, result)
	}

	finalResult := results[0]
	for i, logicOp := range logicOps {
		if logicOp == "AND" {
			finalResult = finalResult && results[i+1]
		} else if logicOp == "OR" {
			finalResult = finalResult || results[i+1]
		}
	}

	if finalResult {
		if gotoLabel != "" {
			stack.JumpLabel = &gotoLabel
		}
		return nil
	}
	stack.ExitScope = true
	return nil
}

func evaluateCondition(memory KLMemory, arguments []VariableBox) (bool, error) {
	leftTokens := make([]VariableBox, 0)
	rightTokens := make([]VariableBox, 0)
	operator := ""

	for _, token := range arguments {
		if slices.Contains([]string{"=", "!", ">", "<"}, token.String) {
			operator = operator + token.String
			continue
		}
		if operator == "" {
			leftTokens = append(leftTokens, token)
		} else {
			rightTokens = append(rightTokens, token)
		}
	}

	leftArguments, err := processArguments(memory, leftTokens)
	if err != nil {
		return false, err
	}
	rightArguments, err := processArguments(memory, rightTokens)
	if err != nil {
		return false, err
	}

	if len(leftArguments) != 1 || len(rightArguments) != 1 {
		return false, fmt.Errorf("if statement must have two arguments")
	}

	first := leftArguments[0]
	second := rightArguments[0]

	var result VariableBox
	switch operator {
	case "=":
		result, err = first.EqualTo(second)
	case "!=", "=!":
		result, err = first.EqualTo(second)
		result.Bool = !result.Bool
	case ">":
		result, err = first.GreaterThan(second)
	case "<":
		result, err = first.LessThan(second)
	case "<=", "=<":
		result, err = first.LessThan(second)
		secondary, err := first.EqualTo(second)
		if err != nil {
			return false, err
		}
		result.Bool = result.Bool || secondary.Bool
	case ">=", "=>":
		result, err = first.GreaterThan(second)
		secondary, err := first.EqualTo(second)
		if err != nil {
			return false, err
		}
		result.Bool = result.Bool || secondary.Bool
	}
	if err != nil {
		return false, err
	}
	return result.Bool, nil
}
