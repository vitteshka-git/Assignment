package main
import "log"

func info(message string) {
	log.Println("[INFO]", message)
}

func warning(message string) {
	log.Println("[WARNING]", message)
}

func errorLog(message string) {
	log.Println("[ERROR]", message)
}

func main() {

	info("Application started")
	warning("Employee salary is low")
	errorLog("Employee not found")
	info("Application completed")
}