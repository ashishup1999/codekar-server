package services

import (
	"bytes"
	"codekar/app/models"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func RunJavafn(req models.CompileReq) models.CompileResp {
	var errorResponse = models.CompileResp{
		Status: "ERROR",
	}

	//create a java file
	tempFile := req.FileName
	file, err := os.Create(tempFile)
	if err != nil {
		errorResponse.Message = "File Creation Failed"
		return errorResponse
	}

	//write java file
	io.WriteString(file, req.Code)

	//create input and output buffer
	var inpBuff bytes.Buffer
	var outBuff bytes.Buffer

	//execute javac command to create class file
	cmd := exec.Command("javac", tempFile)
	err = cmd.Run()
	if err != nil {
		errorResponse.Message = "Javac command failed"
		return errorResponse
	}

	//deletion of tempfile and all .class files with regex
	tempFileParams := strings.Split(tempFile, ".")
	tempFileNameWOExt := strings.Join(tempFileParams[:len(tempFileParams)-1], "")
	classFilePatterns := fmt.Sprintf(`^%s.*%s$`, tempFileNameWOExt, `.class`)
	classFileRegex, err := regexp.Compile(classFilePatterns)
	if err != nil {
		errorResponse.Message = "Regex failed"
		return errorResponse
	}
	allFiles, err := filepath.Glob("*")
	if err != nil {
		errorResponse.Message = "global file search failed"
		return errorResponse
	}

	defer func() {
		file.Close()
		for _, eachFile := range allFiles {
			if classFileRegex.MatchString(eachFile) {
				os.Remove(eachFile)
			}
		}
		os.Remove(tempFile)
	}()

	//execute the compiled java code
	cmd = exec.Command("java", "-classpath", ".", tempFileNameWOExt)

	//send all inputs to buffer
	for _, inp := range req.Inputs {
		inpBuff.WriteString(fmt.Sprintf("%s\n", inp))
	}

	//assign cmd.Stdin the input buffer reference
	cmd.Stdin = &inpBuff

	//assign cmd.Stdout the output buffer reference
	cmd.Stdout = &outBuff

	// Run the command
	err = cmd.Run()
	if err != nil {
		errorResponse.Message = "File execution failed"
		return errorResponse
	}

	//return value
	return models.CompileResp{
		Status: "SUCCESS",
		Output: outBuff.String(),
	}
}

func RunPythonfn(req models.CompileReq) models.CompileResp {
	var errorResponse = models.CompileResp{
		Status: "ERROR",
	}

	//create a py file
	tempFile := req.FileName
	file, err := os.Create(tempFile)
	if err != nil {
		errorResponse.Message = "File Creation Failed"
		return errorResponse
	}

	//write python file
	io.WriteString(file, req.Code)

	//delete file after execution
	defer func() {
		file.Close()
		os.Remove(tempFile)
	}()

	//create input and output buffer
	var inpBuff bytes.Buffer
	var outBuff bytes.Buffer

	//assign cmd command to execute
	cmd := exec.Command("python3", tempFile)

	//send all inputs to buffer
	for _, inp := range req.Inputs {
		inpBuff.WriteString(fmt.Sprintf("%s\n", inp))
	}

	//assign cmd.Stdin the input buffer
	cmd.Stdin = &inpBuff

	//assign cmd.Stdout the output buffer
	cmd.Stdout = &outBuff

	// Run the command
	err = cmd.Run()
	if err != nil {
		errorResponse.Message = "File execution failed"
		return errorResponse
	}

	//get the output as string from output buffer
	strOutput := outBuff.String()

	//return value
	return models.CompileResp{
		Status: "SUCCESS",
		Output: strOutput,
	}
}

func RunCppfn(req models.CompileReq) models.CompileResp {
	var errorResponse = models.CompileResp{
		Status: "ERROR",
	}

	//create a cpp file
	tempFile := req.FileName
	file, err := os.Create(tempFile)
	if err != nil {
		errorResponse.Message = "File Creation Failed"
		return errorResponse
	}

	//write cpp file
	io.WriteString(file, req.Code)

	// delete file after execution
	defer func() {
		file.Close()
		os.Remove(tempFile)
		os.Remove("out.exe")
	}()

	//create input and output buffer
	var inpBuff bytes.Buffer
	var outBuff bytes.Buffer

	//assign cmd command to create executable
	cmd := exec.Command("g++", tempFile, "-o", "out")

	// Run the command to create executable file
	err = cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		errorResponse.Message = "Creation Failed"
		return errorResponse
	}

	//assign cmd command to rrun executable file
	cmd = exec.Command("./out")

	//assign cmd.Stdin the input buffer
	cmd.Stdin = &inpBuff

	//assign cmd.Stdout the output buffer
	cmd.Stdout = &outBuff
	
	//send all inputs to buffer
	for _, inp := range req.Inputs {
		inpBuff.WriteString(fmt.Sprintf("%s\n", inp))
	}

	// Run the command to run executable file
	err = cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		errorResponse.Message = "Execution Failed"
		return errorResponse
	}

	//get the output as string from output buffer
	strOutput := outBuff.String()

	//return value
	return models.CompileResp{
		Status: "SUCCESS",
		Output: strOutput,
	}
}


func RunGofn(req models.CompileReq) models.CompileResp {
	var errorResponse = models.CompileResp{
		Status: "ERROR",
	}

	//create a go file
	tempFile := req.FileName
	file, err := os.Create(tempFile)
	if err != nil {
		errorResponse.Message = "File Creation Failed"
		return errorResponse
	}

	//write go file
	io.WriteString(file, req.Code)

	//delete file after execution
	defer func() {
		file.Close()
		os.Remove(tempFile)
	}()

	//create input and output buffer
	var inpBuff bytes.Buffer
	var outBuff bytes.Buffer

	//assign cmd command to execute
	cmd := exec.Command("go","run", tempFile)

	//send all inputs to buffer
	for _, inp := range req.Inputs {
		inpBuff.WriteString(fmt.Sprintf("%s\n", inp))
	}

	//assign cmd.Stdin the input buffer
	cmd.Stdin = &inpBuff

	//assign cmd.Stdout the output buffer
	cmd.Stdout = &outBuff

	// Run the command
	err = cmd.Run()
	if err != nil {
		errorResponse.Message = "File execution failed"
		return errorResponse
	}

	//get the output as string from output buffer
	strOutput := outBuff.String()

	//return value
	return models.CompileResp{
		Status: "SUCCESS",
		Output: strOutput,
	}
}

func RunJSfn(req models.CompileReq) models.CompileResp {
	var errorResponse = models.CompileResp{
		Status: "ERROR",
	}

	//create a JS file
	tempFile := req.FileName
	file, err := os.Create(tempFile)
	if err != nil {
		errorResponse.Message = "File Creation Failed"
		return errorResponse
	}

	//write JS file
	io.WriteString(file, req.Code)

	//delete file after execution
	defer func() {
		file.Close()
		os.Remove(tempFile)
	}()

	//create input and output buffer
	var inpBuff bytes.Buffer
	var outBuff bytes.Buffer

	//assign cmd command to execute
	cmd := exec.Command("node","-e", tempFile)

	//send all inputs to buffer
	for _, inp := range req.Inputs {
		inpBuff.WriteString(fmt.Sprintf("%s\n", inp))
	}

	//assign cmd.Stdin the input buffer
	cmd.Stdin = &inpBuff

	//assign cmd.Stdout the output buffer
	cmd.Stdout = &outBuff

	// Run the command
	err = cmd.Run()
	if err != nil {
		errorResponse.Message = "File execution failed"
		return errorResponse
	}

	//get the output as string from output buffer
	strOutput := outBuff.String()

	//return value
	return models.CompileResp{
		Status: "SUCCESS",
		Output: strOutput,
	}
}