package interpreter

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/talon-one/talang/token"
)

var coreFunctions []TaFunction

func RegisterCoreFunction(signatures ...TaFunction) error {
	for i := 0; i < len(signatures); i++ {
		signature := signatures[i]
		signature.sanitize()
		if getFunction(coreFunctions, &signature) != nil {
			return fmt.Errorf("Function `%s' is already registered", signature.Name)
		}
		coreFunctions = append(coreFunctions, signature)
	}
	return nil
}

func (interp *Interpreter) RegisterFunction(signatures ...TaFunction) error {
	for i := 0; i < len(signatures); i++ {
		signature := signatures[i]
		signature.sanitize()
		if interp.GetFunction(&signature) != nil {
			return errors.Errorf("Function `%s' is already registered", signature.Name)
		}
		interp.Functions = append(interp.Functions, signature)
	}
	return nil
}

func (interp *Interpreter) MustRegisterFunction(signatures ...TaFunction) {
	if err := interp.RegisterFunction(signatures...); err != nil {
		panic(err)
	}
}

func (interp *Interpreter) UpdateFunction(signature TaFunction) error {
	signature.sanitize()
	if s := interp.GetFunction(&signature); s != nil {
		*s = signature
		return nil
	}
	return errors.Errorf("Function `%s' is not registered", signature.Name)
}

func (interp *Interpreter) RemoveFunction(signature TaFunction) error {
	signature.sanitize()
	for i := 0; i < len(interp.Functions); i++ {
		if interp.Functions[i].Equal(&signature) {
			fns := interp.Functions[:i]
			interp.Functions = append(fns, interp.Functions[i+1:]...)
			return nil
		}
	}
	return errors.Errorf("Function `%s' is not registered", signature.Name)
}

func getFunction(functions []TaFunction, signature *TaFunction) *TaFunction {
	signature.sanitize()
	for i := 0; i < len(functions); i++ {
		if functions[i].Equal(signature) {
			return &functions[i]
		}
	}
	return nil
}

func (interp *Interpreter) GetFunction(signature *TaFunction) *TaFunction {
	return getFunction(interp.Functions, signature)
}

func (interp *Interpreter) registerCoreFunctions() error {
	bindingSignature.sanitize()
	setBindingSignature.sanitize()
	templateSignature.sanitize()

	// binding
	interp.Functions = append(interp.Functions, bindingSignature)
	interp.Functions = append(interp.Functions, setBindingSignature)

	// template
	interp.Functions = append(interp.Functions, setTemplateSignature)
	interp.Functions = append(interp.Functions, templateSignature)

	interp.Functions = append(interp.Functions, coreFunctions...)

	// sanitize name
	for i := range interp.Functions {
		interp.Functions[i].sanitize()
	}
	return nil
}

func (interp *Interpreter) RemoveAllFunctions() error {
	interp.Functions = []TaFunction{}
	return nil
}

var bindingSignature = TaFunction{
	CommonSignature: CommonSignature{
		Name:       ".",
		IsVariadic: true,
		Arguments: []token.Kind{
			token.Atom,
			token.Atom,
		},
		Returns:     token.Any,
		Description: "Access a variable in the binding",
		Example: `
(. Key1)                                                         ; returns the data assigned to Key1
(. Key2 SubKey1)                                                 ; returns the data assigned to SubKey1 in the Map Key2
`,
	},
	Func: bindingFunc,
}

func bindingFunc(interp *Interpreter, args ...*token.TaToken) (*token.TaToken, error) {
	argc := len(args)
	if interp.Binding != nil {
		value := interp.Binding
		for i := 0; i < argc; i++ {
			if !value.IsMap() {
				break
			}
			value = value.MapItem(args[i].String)
			if value.IsNull() {
				break
			}
		}

		if !value.IsNull() {
			return value, nil
		}
	}

	// lookup in parent
	if interp.Parent != nil {
		value, err := bindingFunc(interp.Parent, args...)
		if err == nil {
			return value, nil
		}
	}

	qualifiers := make([]string, argc)
	for j, arg := range args {
		qualifiers[j] = arg.String
	}
	return nil, errors.Errorf("Unable to find `%s'", strings.Join(qualifiers, "."))
}

var setBindingSignature = TaFunction{
	CommonSignature: CommonSignature{
		Name:       "set",
		IsVariadic: true,
		Arguments: []token.Kind{
			token.String,
			token.Atom | token.Collection,
			token.Atom | token.Collection,
		},
		Returns:     token.Null,
		Description: "Set a variable in the binding",
		Example: `
(set Key1 "Hello World")                                         ; sets Key1 to "Hello World"
(set Key2 SubKey1 true)                                          ; sets SubKey1 in map Key2 to true
`,
	},
	Func: setBindingFunc,
}

func setBindingFunc(interp *Interpreter, args ...*token.TaToken) (*token.TaToken, error) {
	argc := len(args)
	if argc < 2 {
		return nil, errors.New("invalid or missing arguments")
	}
	if interp.Binding == nil {
		interp.Binding = token.NewMap(map[string]*token.TaToken{})
	}

	value := interp.Binding
	for i := 0; i < argc-2; i++ {
		child := value.MapItem(args[i].String)
		if child.IsNull() {
			child = token.NewMap(map[string]*token.TaToken{})
			value.SetMapItem(args[i].String, child)
		}
		value = child
	}
	value.SetMapItem(args[argc-2].String, args[argc-1])
	return token.NewNull(), nil
}
