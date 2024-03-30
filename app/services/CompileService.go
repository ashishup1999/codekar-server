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

func RunJavafn(req models.CompileReq) (string, error) {

	//create a java file
	tempFile := "Main.java"
	file, err := os.Create(tempFile)
	if err != nil {
		return "", nil
	}

	//write java file
	io.WriteString(file, req.Code)

	//create input and output buffer
	var inpBuff bytes.Buffer
	var outBuff bytes.Buffer
	var errBuff bytes.Buffer

	//execute javac command to create class file
	cmd := exec.Command("javac", tempFile)

	//attaching error buffer to terminal to detect if any error comes out
	cmd.Stderr = &errBuff

	//run cmd
	err = cmd.Run()
	if err != nil {
		return errBuff.String(), nil
	}

	//deletion of tempfile and all .class files with regex
	tempFileParams := strings.Split(tempFile, ".")
	tempFileNameWOExt := strings.Join(tempFileParams[:len(tempFileParams)-1], "")
	classFilePatterns := fmt.Sprintf(`^%s.*%s$`, tempFileNameWOExt, `.class`)
	classFileRegex, err := regexp.Compile(classFilePatterns)
	if err != nil {
		return "", err
	}
	allFiles, err := filepath.Glob("*")
	if err != nil {
		return "", err
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

	//attaching all kinds of buffer to terminal
	cmd.Stdin = &inpBuff
	cmd.Stdout = &outBuff
	cmd.Stderr = &errBuff

	// Run the command
	err = cmd.Run()
	if err != nil {
		fmt.Println(err, outBuff.String())
		return errBuff.String(), nil
	}

	//return value
	return outBuff.String(), nil
}

func RunPythonfn(req models.CompileReq) (string, error) {

	//create a py file
	tempFile := req.FileName + ".py"
	file, err := os.Create(tempFile)
	if err != nil {
		return "", err
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
	var errBuff bytes.Buffer

	//assign cmd command to execute
	cmd := exec.Command("python3", tempFile)

	//send all inputs to buffer
	for _, inp := range req.Inputs {
		inpBuff.WriteString(fmt.Sprintf("%s\n", inp))
	}

	//attaching all kinds of buffer to terminal
	cmd.Stdin = &inpBuff
	cmd.Stdout = &outBuff
	cmd.Stderr = &errBuff

	// Run the command
	err = cmd.Run()
	if err != nil {
		return errBuff.String(), nil
	}

	//return value
	return outBuff.String(), nil
}

func RunCppfn(req models.CompileReq) (string, error) {

	//create a cpp file
	tempFile := req.FileName + ".cpp"
	file, err := os.Create(tempFile)
	if err != nil {
		return "", err
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
	var errBuff bytes.Buffer

	//assign cmd command to create executable
	cmd := exec.Command("g++", tempFile, "-o", "out")

	// Run the command to create executable file
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	//assign cmd command to rrun executable file
	cmd = exec.Command("./out")

	//attaching all kinds of buffer to terminal
	cmd.Stdin = &inpBuff
	cmd.Stdout = &outBuff
	cmd.Stderr = &errBuff

	//send all inputs to buffer
	for _, inp := range req.Inputs {
		inpBuff.WriteString(fmt.Sprintf("%s\n", inp))
	}

	// Run the command to run executable file
	err = cmd.Run()
	if err != nil {
		return errBuff.String(), nil
	}

	//return value
	return outBuff.String(), nil
}

func RunGofn(req models.CompileReq) (string, error) {

	//create a go file
	tempFile := req.FileName + ".go"
	file, err := os.Create(tempFile)
	if err != nil {
		return "", err
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
	var errBuff bytes.Buffer

	//assign cmd command to execute
	cmd := exec.Command("go", "run", tempFile)

	//send all inputs to buffer
	for _, inp := range req.Inputs {
		inpBuff.WriteString(fmt.Sprintf("%s\n", inp))
	}

	//attaching all kinds of buffer to terminal
	cmd.Stdin = &inpBuff
	cmd.Stdout = &outBuff
	cmd.Stderr = &errBuff

	// Run the command
	err = cmd.Run()
	if err != nil {
		return errBuff.String(), nil
	}

	//return value
	return outBuff.String(), nil
}

func RunJSfn(req models.CompileReq) (string, error) {

	//create a JS file
	tempFile := req.FileName + ".js"
	file, err := os.Create(tempFile)
	if err != nil {
		return "", err
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
	var errBuff bytes.Buffer

	//assign cmd command to execute
	cmd := exec.Command("node", tempFile)

	//send all inputs to buffer
	for _, inp := range req.Inputs {
		inpBuff.WriteString(fmt.Sprintf("%s\n", inp))
	}

	//attaching all kinds of buffer to terminal
	cmd.Stdin = &inpBuff
	cmd.Stdout = &outBuff
	cmd.Stderr = &errBuff

	// Run the command
	err = cmd.Run()
	if err != nil {
		return errBuff.String(), nil
	}

	//return value
	return outBuff.String(), nil
}
