package logutils

import (
	"fmt"
	"log"
)

func Printf(format string, v ...any) {
	log.Printf(format, v...)
}

func Debugf(format string, v ...any) {
	log.Printf("[DEBUG] "+format, v...)
}

func Warnf(format string, v ...any) {
	log.Printf("[WARN] "+format, v...)
}

func Errorf(format string, v ...any) {
	log.Printf("[ERROR] "+format, v...)
}

func Fatalf(format string, v ...any) {
	log.Fatalf("[ERROR] "+format, v...)
}

//

func Print(v ...any) {
	log.Print(v...)
}

func Debug(v ...any) {
	log.Print("[DEBUG] ", fmt.Sprint(v...))
}

func Warn(v ...any) {
	log.Print("[WARN] ", fmt.Sprint(v...))
}

func Error(v ...any) {
	log.Print("[ERROR] ", fmt.Sprint(v...))
}

func Fatal(v ...any) {
	log.Fatal("[ERROR] ", fmt.Sprint(v...))
}

//

func Println(v ...any) {
	log.Println(v...)
}

func Debugln(v ...any) {
	log.Println("[DEBUG] ", fmt.Sprint(v...))
}

func Warnln(v ...any) {
	log.Println("[WARN] ", fmt.Sprint(v...))
}

func Errorln(v ...any) {
	log.Println("[ERROR] ", fmt.Sprint(v...))
}

func Fatalln(v ...any) {
	log.Fatalln("[ERROR] ", fmt.Sprint(v...))
}
