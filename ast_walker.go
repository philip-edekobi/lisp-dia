package main

import (
	"fmt"
	"strconv"
)

var builtins = map[string]func([]Value, map[string]any) any{}

func initializeBuiltins() {
	builtins["if"] = func(args []Value, ctx map[string]any) any {
		condition := Walk(args[0], ctx)
		then := args[1]
		_else := args[2]

		if condition.(bool) == true {
			return Walk(then, ctx)
		}

		return Walk(_else, ctx)
	}

	builtins["+"] = func(args []Value, ctx map[string]any) any {
		var i int
		for _, arg := range args {
			i += Walk(arg, ctx).(int)
		}

		return i
	}

	builtins["-"] = func(args []Value, ctx map[string]any) any {
		i := Walk(args[0], ctx).(int)
		for _, arg := range args[1:] {
			i -= Walk(arg, ctx).(int)
		}

		return i
	}

	builtins["*"] = func(args []Value, ctx map[string]any) any {
		var i int = 1
		for _, arg := range args {
			i *= Walk(arg, ctx).(int)
		}

		return i
	}

	builtins["/"] = func(args []Value, ctx map[string]any) any {
		i := Walk(args[0], ctx).(int)
		for _, arg := range args[1:] {
			i /= Walk(arg, ctx).(int)
		}

		return i
	}

	builtins["define"] = func(args []Value, ctx map[string]any) any {
		if len(args) != 2 {
			msg := fmt.Sprintf("define should come with 2 arguments")
			panic(msg)
		}
		if args[0].Kind == literalValue {
			ctx[(*args[0].Literal).Value] = Walk(args[1], ctx)
		}

		if args[0].Kind == listValue {
			funcList := []Value(*args[0].List)
			varArray := []string{}

			fName := (*funcList[0].Literal).Value

			for _, val := range funcList[1:] {
				if val.Kind != literalValue {
					panic("can not pass an expr as an identifier or argument to a function")
				}

				varArray = append(varArray, (*val.Literal).Value)
			}

			builtins[fName] = func(fArgs []Value, fCtx map[string]any) any {
				internalFuncCtx := mapCopy(fCtx)

				if len(varArray) != len(fArgs) {
					msg := fmt.Sprintf("func %s expects %d args", fName, len(varArray))
					panic(msg)
				}

				for i := 0; i < len(fArgs); i++ {
					internalFuncCtx[varArray[i]] = Walk(fArgs[i], internalFuncCtx)
				}

				return Walk(args[1], internalFuncCtx)
			}
		}

		return nil
	}

	builtins["<"] = func(args []Value, ctx map[string]any) any {
		val1 := Walk(args[0], ctx).(int)
		val2 := Walk(args[1], ctx).(int)

		return val1 < val2
	}

	builtins[">"] = func(args []Value, ctx map[string]any) any {
		val1 := Walk(args[0], ctx).(int)
		val2 := Walk(args[1], ctx).(int)

		return val1 > val2
	}
}

func mapCopy(mp map[string]any) map[string]any {
	newMap := make(map[string]any)

	for k, v := range mp {
		newMap[k] = v
	}

	return newMap
}

/*
 * (+ 13 (- 12 1))
 * (+ 13 11)
 * 24
 */

func interpret(v Value, ctx map[string]any) any {
	if v.Kind == literalValue {
		return Walk(v, ctx)
	}

	return WalkList(*v.List, ctx)
}

func Walk(v Value, ctx map[string]any) any {
	if v.Kind == literalValue {
		lit := *v.Literal
		switch lit.Kind {
		case integerToken:
			integer, err := strconv.Atoi(lit.Value)
			if err != nil {
				fmt.Printf("%v should be an integer\n", lit.Value)
				panic("something went horribly wrong")
			}

			return integer

		case identifierToken:
			if val, ok := ctx[lit.Value]; ok {
				return val
			}
			msg := fmt.Sprintf("identifier %s is not defined", lit.Value)
			panic(msg)

		default:
			return lit
		}
	}

	return WalkList(*v.List, ctx)
}

func WalkList(ast []Value, ctx map[string]any) any {
	functionName := (ast[0].Literal).Value

	if function, ok := builtins[functionName]; ok {
		return function(ast[1:], ctx)
	}

	// Calling something that isn't built in
	msg := fmt.Sprintf("%s is not defined", functionName)
	panic(msg)
}
