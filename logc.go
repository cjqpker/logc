package logc

import (
	"fmt"
	gl "log"
	"os"
	"time"
)

func appendToFile(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("/tmp/logc.txt file create failed. err: " + err.Error())
	} else {
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
	return err
}

//Logc Custom log method
func Logc(f string, args ...interface{}) {
	s := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05 ")+f+"\n", args...)

	if err := appendToFile("/tmp/logc.txt", s); err != nil {
		gl.Fatal(err)
	}
}
