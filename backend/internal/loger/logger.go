package loger

import (
	"log"
	"os"
)

type ErrStr struct {
	fileName string
}

func NewLogger() *ErrStr {
	return &ErrStr{
		fileName: "log.txt",
	}
}

func (e *ErrStr) Output(newErr string) {
	log.Println(newErr)
	e.Save(newErr)
}

func (e *ErrStr) Save(newErr string) {
	f, err := os.OpenFile(e.fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		e.Output(err.Error())
	}
	defer f.Close()
	if _, err := f.Write([]byte(newErr)); err != nil {
		log.Println(err)
	}
}
