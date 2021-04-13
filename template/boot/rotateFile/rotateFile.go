package rotateFile

import (
	"qkgo-template/boot/config"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"time"
)

func Open(filename string) (*rotatelogs.RotateLogs, error) {
	reverseDuration := config.LogConfig().Rotate.ReverseTime
	rotationDuration := config.LogConfig().Rotate.RotateTime
	fileSuffixFormat := config.LogConfig().Rotate.FileSuffixFormat
	reverseTime, err := time.ParseDuration(reverseDuration)
	if err != nil {
		fmt.Println(err)
		// parse err
	}
	rotationTime, err := time.ParseDuration(rotationDuration)
	if err != nil {
		fmt.Println(err)
		// parse err
	}

	return rotatelogs.New(
		filename +fileSuffixFormat,
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(reverseTime),
		rotatelogs.WithRotationTime(rotationTime),
	)
}
