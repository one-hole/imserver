package logs

import (
	"github.com/sirupsen/logrus"
	"os"
)

var (
	WebSocketLogger = logrus.New()
)

func init()  {
	var socketFile *os.File
	var err error
	
	if socketFile, err =createOrOpenFile("./logs/websocket.log"); err != nil {
		panic(err)
	}
	
	WebSocketLogger.Out = socketFile
}

func createOrOpenFile(path string) (*os.File, error) {

	var file *os.File
	var err error

	if file, err = os.OpenFile(path, os.O_RDWR, os.ModeAppend); err != nil {
		file, err = os.Create(path)
	}

	return file, err
}
