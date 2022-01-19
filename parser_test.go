package main

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/stretchr/testify/assert"
)

func TestParserWithEmptyQuery(t *testing.T) {
	const query = ""

	var errorListener ErrorListenerImpl
	var listener ListenerImpl
	p := CreateParser(&errorListener, query)

	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Query())

	assert.Equal(t, 1, errorListener.Count, "Empty query should be an error")
	assert.Equal(t, ListenerImpl{}, listener)
}

func TestParserWithoutWhereClause(t *testing.T) {
	const query = "SELECT * FROM pods"

	var errorListener ErrorListenerImpl
	var listener ListenerImpl
	p := CreateParser(&errorListener, query)

	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Query())

	assert.Equal(t, 0, errorListener.Count, "Found errors in input")
	assert.Equal(t, ListenerImpl{}, listener)
}

func TestParserWithNamespace(t *testing.T) {
	const query = "SELECT * FROM pods WHERE namespace=default"

	var errorListener ErrorListenerImpl
	var listener ListenerImpl
	p := CreateParser(&errorListener, query)

	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Query())

	assert.Equal(t, 0, errorListener.Count, "Found errors in input")
	expected := map[string]string{
		"namespace": "default",
	}
	assert.Equal(t, expected, listener.Fields)
}

func TestParserWithName(t *testing.T) {
	const query = "SELECT * FROM pods WHERE name=blargle1-flargle2.example.com"

	var errorListener ErrorListenerImpl
	var listener ListenerImpl
	p := CreateParser(&errorListener, query)

	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Query())

	assert.Equal(t, 0, errorListener.Count, "Found errors in input")
	expected := map[string]string{
		"name": "blargle1-flargle2.example.com",
	}
	assert.Equal(t, expected, listener.Fields)
}

func TestParserWithNameAndNamespace(t *testing.T) {
	const query = "SELECT * FROM pods WHERE name=blargle AND namespace=flargle"

	var errorListener ErrorListenerImpl
	var listener ListenerImpl
	p := CreateParser(&errorListener, query)

	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Query())

	assert.Equal(t, 0, errorListener.Count, "Found errors in input")
	expected := map[string]string{
		"name":      "blargle",
		"namespace": "flargle",
	}
	assert.Equal(t, expected, listener.Fields)
}
