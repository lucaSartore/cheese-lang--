package test

import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/parser"
	"cheese-lang/internal/tokenizer"
	"testing"

	"github.com/go-errors/errors"
)

func PrintErrorWithStack(message string, t *testing.T, err error) {
	err2, ok := err.(*errors.Error)
	if ok {
		t.Errorf("%v: %v\nTrace: %v", message, err2, err2.ErrorStack())
	} else {
		t.Errorf("%v: %v", message, err)
	}
}

func DoTestOnString(code string, variableToTest string, expectedValue expressions.VariableContainer, t *testing.T) {

	tokens, err := tokenizer.Tokenize(code)

	if err != nil {
		PrintErrorWithStack("Error while tokenizing", t, err)
		return
	}

	context := expressions.MakeContext()

	parser := parser.MakeParser(tokens)

	returnValue := parser.ParseAnything(false)

	if returnValue.Error != nil {
		PrintErrorWithStack("Error while parsing", t, returnValue.Error)
		return
	}

	if returnValue.Expression == nil {
		t.Errorf("Expect to parse an expression, but got nil")
		return
	}

	_, err = returnValue.Expression.Evaluate(&context, &context)

	if err != nil {
		PrintErrorWithStack("Error while evaluating", t, err)
		return
	}

	variable, ok := context.GetVariable(variableToTest)
	if !ok {
		t.Errorf("Variable not found")
		return
	}

	if !VerifyValueEquivalence(variable.Value, expectedValue) {
		t.Errorf("Expected value: %v, got: %v", expectedValue, variable.Value)
		return
	}
}

func TestMozzarellaVariableParser(t *testing.T) {
	code := "Mozzarella x = \"hello\" + \" \"  + \"world\""
	DoTestOnString(code, "x", &expressions.MozzarellaVariable{Value: "hello world"}, t)
}

func TestParmesanVariableParser(t *testing.T) {
	code := "Parmesan x = (100-((5 + 3) * 2)) * (11 - 1)"
	DoTestOnString(code, "x", &expressions.ParmesanVariable{Value: 840}, t)
}

func TestCodeBlock(t *testing.T) {
	code := `
	{
		Mozzarella x = "hello";
		Mozzarella y = "world";
		Mozzarella z = "undefined";
		z = x + " " + y;
	}
	`
	DoTestOnString(code, "z", &expressions.MozzarellaVariable{Value: "hello world"}, t)
}

func TestTasteBlock(t *testing.T) {
	code := `
		{
        	Parmesan x = 100;
        	Milk c1 = spoiled;
        	Milk c2 = spoiled;
        	taste parmesan > 50 {
            	c1 = fresh;
        	}
        	taste parmesan <= 50 {
            	c2 = fresh;
        	}
        
        	Milk success = (c1 == fresh) && (c2 == spoiled)
		}
    `
	DoTestOnString(code, "success", &expressions.MilkVariable{Value: true}, t)
}