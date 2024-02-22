package ylog

import (
	"fmt"
	"strconv"
	"time"

	"github.com/json-iterator/go"
)

type JsonFormatter struct {
	IgnoreBasicFields bool
}

func (f *JsonFormatter) Format(e *Entry) error {
	if f.IgnoreBasicFields {
		switch e.Format {
		case "":
			for _, arg := range e.Args {
				err := jsoniter.NewEncoder(e.Buffer).Encode(arg)
				if err != nil {
					return err
				}
			}
		default:
			e.Buffer.WriteString(fmt.Sprintf(e.Format, e.Args...))
		}

		return nil
	}

	e.Map["level"] = LevelNameMapping[e.Level]
	e.Map["time"] = e.Time.Format(time.RFC3339)
	if e.File != "" {
		e.Map["file"] = e.File + ":" + strconv.Itoa(e.Line)
		e.Map["func"] = e.Func
	}

	switch e.Format {
	case "":
		e.Map["message"] = fmt.Sprint(e.Args...)
	default:
		e.Map["message"] = fmt.Sprintf(e.Format, e.Args...)
	}

	return jsoniter.NewEncoder(e.Buffer).Encode(e.Map)
}
