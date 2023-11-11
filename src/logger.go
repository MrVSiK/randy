package generations

import (
	"bytes"
	"log"
	"os"
);

var buffer bytes.Buffer;
var FileLogger log.Logger = *log.New(&buffer, "Logger: ", log.Llongfile);
var ConsoleLogger log.Logger = *log.New(os.Stdout, "Randy: ", log.Lshortfile);

func LogToFile(output string){
	FileLogger.Print(output);
}

func LogToConsole(output string){
	go LogToFile(output);
	ConsoleLogger.Print(output);
}