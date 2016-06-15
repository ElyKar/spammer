package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

// A QuestionFactory's job is to create random questions
type QuestionFactory interface {
	CreateQuestion() Question
}

// Questions is defined by its content (json produced) and its expected answer
type Question interface {
	// Content of the question in the form of json
	Content() string
	// Answer to the question
	Answer() interface{}
}

// Factories registered in the game
var factories []QuestionFactory = []QuestionFactory{
	new(AddQuestionFactory),
	new(SubQuestionFactory),
	new(MulQuestionFactory),
	new(DivQuestionFactory),
}

// Pseudo-random generator
var gen *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// Creates a random question
func NewQuestion() Question {
	factory := factories[gen.Intn(len(factories))]
	return factory.CreateQuestion()
}

// Implements a Question representing basic arithmetic operations
type Operation struct {
	// Type of the operation (ADD, SUB, MUL, DIV)
	Operation string
	// Left operand
	OperandA float64
	// Right operand
	OperandB float64
	// Expected result (is not exported)
	result interface{}
}

// Get the json for the Question
func (op *Operation) Content() string {
	content, err := json.Marshal(op)
	if err != nil {
		fmt.Println("An error occurred while creating a new question")
		return "{}"
	}
	return string(content)
}

// Get the expected answer
func (op *Operation) Answer() interface{} {
	return op.result
}

// Factory for additions
type AddQuestionFactory struct{}

// Creates a random addition
func (f *AddQuestionFactory) CreateQuestion() Question {
	a := rand.Float64()
	b := rand.Float64()
	operation := &Operation{
		Operation: "ADD",
		OperandA:  a,
		OperandB:  b,
		result:    a + b,
	}
	return operation
}

// Substraction factory
type SubQuestionFactory struct{}

// Creates a random substraction
func (s *SubQuestionFactory) CreateQuestion() Question {
	a := rand.Float64()
	b := rand.Float64()
	operation := &Operation{
		Operation: "SUB",
		OperandA:  a,
		OperandB:  b,
		result:    a - b,
	}
	return operation
}

// Multiplication factory
type MulQuestionFactory struct{}

// Creates a random multiplication
func (m *MulQuestionFactory) CreateQuestion() Question {
	a := rand.Float64()
	b := rand.Float64()
	operation := &Operation{
		Operation: "MUL",
		OperandA:  a,
		OperandB:  b,
		result:    a * b,
	}
	return operation
}

// Division factory
type DivQuestionFactory struct{}

// Creates a random division
func (d *DivQuestionFactory) CreateQuestion() Question {
	a := rand.Float64()
	b := rand.Float64()
	operation := &Operation{
		Operation: "DIV",
		OperandA:  a,
		OperandB:  b,
		result:    a / b,
	}
	return operation
}
