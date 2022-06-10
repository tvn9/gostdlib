// Logging customization
// Besides the logging with the default logging from the package,
// The standard library also privide a way to create the custom loger,
// accoding to the needs of the application or package.
package main

import (
	"log"
	"os"
)

func main() {
	custLogger := log.New(os.Stdout, "custom1: ", log.Ldate|log.Ltime)
	custLogger.Println("Hello I'am customized looger 1")
	custLoggerEnh := log.New(os.Stdout, "custom2: ", log.Ldate|log.Lshortfile)
	custLoggerEnh.Println("Hello I'am customized logger 2")
}
