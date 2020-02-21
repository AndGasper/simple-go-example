package main

import (
	"io/ioutil"

	"github.com/ghodss/yaml"

	hcltoken "github.com/hashicorp/hcl/hcl/token"
	hashicorpJsonParser "github.com/hashicorp/hcl/json/parser"
	hashicorpJsonScanner "github.com/hashicorp/hcl/json/scanner"
	hashicorpJsonTokens "github.com/hashicorp/hcl/json/token"

	// because I am too much of a pleb to implement my own structs
	"fmt"
	"os"
	"path/filepath"
)

// ast
// node1
// node

func check(e error) {
	// check for errors and panic if so. Productive...
	if e != nil {
		panic(e)
	}
}

func convertYamlToJSON(yamlFile []byte) []byte {
	// take in yaml file as bytes and return some wicked json bytes
	jsonBytes, err := yaml.YAMLToJSON(yamlFile)
	check(err)
	return jsonBytes
}

func print(thing interface{}) {
	// because I keep typing print jfc
	fmt.Print(thing)
}


// ah. so *thing => pointer
// *grumbles in half forgotten pointer logic*
func convertJSONToHcl(jsonBytes []byte) []hcltoken.Token {
	// take in some json bytes and use the hashicorp hcl package
	// to do stuff and things to get hcl out
	// stuff -lex-> token stream -parse-> ast -compile-> other thing
	// It's coming back to me...Dormammu, I've come to bargain.
	// so I parse these mystical jsonBytes, and I have some abstract syntax tree waiting for me
	// func (receiver) Function/Method? returnValue
	// the hcl method seems to jump from the lex check to parsing hence why the parse method returns an ast?
	// really a pointer for astFile
	astFile, err := hashicorpJsonParser.Parse(jsonBytes) // it checks if it should lex from token to token Stream for json or hcl
	check(err)
	tokenPosition := astFile.Pos()
	fmt.Print("tokenPosition\n")
	fmt.Print(tokenPosition)
	// so now I have this stream of tokens; maybe
	scanner := hashicorpJsonScanner.New(jsonBytes)

	var hclTokens []hcltoken.Token
	for {
		// pick off the return values of Scan
		// token has attributes of: Type, Position, Text
		token := scanner.Scan()
		if token.Type == hashicorpJsonTokens.EOF {
			break
		}
		fmt.Print("\ntoken: ")
		fmt.Print(token)
		// don't turn braces into hcl tokens?
		// because they're not special to hcl
		// if not a literalbegin trying to build it
		if !token.type.IsLiteral() {

		}
		if token.Type.IsLiteral() {
			hclToken := token.HCLToken() // convert the json token to an HCL token
			// accumulate the hcl token
			hclTokens = append(hclTokens, hclToken)
		}
	}

	return hclTokens
}

func main() {
	var yamlFiles []string
	const sourceDirectory = "./fixtures/"
	absolutePath, absPathErr := filepath.Abs(sourceDirectory)
	// there's gotta be a better way...
	if absPathErr != nil {
		// uh oh.
		panic(absPathErr)
	}
	// path which is of type string
	err := filepath.Walk(absolutePath, func(path string, info os.FileInfo, err error) error {
		fileExtension := filepath.Ext(path)
		// only  add yaml files to the lis
		if fileExtension == ".yaml" {
			fmt.Println("yaml file found")
			yamlFiles = append(yamlFiles, path)
		}
		return nil
	})
	if err != nil {
		// Uh oh.
		panic(err)
	}
	for _, yamlFile := range yamlFiles {
		fmt.Println(yamlFile)
		// cheat and just use ioutil.ReadFile; don't want to deal with allocating the buffer right now
		// ALL THE MEMORY!
		yamlBytes, err := ioutil.ReadFile(yamlFile)
		check(err)
		jsonBytes := convertYamlToJSON(yamlBytes)
		hclTokens := convertJSONToHcl(jsonBytes)
		fmt.Print(hclTokens)
		// fileContent := []byte(hclTokens)

	}
}

// https://gobyexample.com/reading-files

// token-ify
// stream(tokens) -into-> |parser| -> AST
