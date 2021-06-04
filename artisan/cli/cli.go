package cli

import (
	"github.com/wuyan94zl/go-api/artisan/command"
	"github.com/wuyan94zl/go-api/artisan/crud"
	"github.com/wuyan94zl/go-api/artisan/model"
	"github.com/wuyan94zl/go-api/artisan/queue"
)

type Command interface {
	Run() error
	GetDir() string
}

func NewArtisan(command []string) Command {
	return getCommand(command[1], command[2])
}

func getCommand(method string, name string) Command {
	switch method {
	case "model":
		return &model.Command{
			Name: name,
		}
	case "api":
		return &crud.Command{
			Name: name,
		}
	case "console":
		return &command.Command{
			Name: name,
		}
	case "queue":
		return &queue.Command{
			Name: name,
		}
	}
	return nil
}
