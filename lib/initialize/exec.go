package init

import (
	"os"
)

func Exec() error {
	if err := os.MkdirAll("interfaces/handlers", 0777); err != nil {
		return err
	}

	if err := os.MkdirAll("interfaces/presenters", 0777); err != nil {
		return err
	}

	if err := os.MkdirAll("applications/usecases", 0777); err != nil {
		return err
	}

	if err := os.MkdirAll("applications/usecases/inputs/", 0777); err != nil {
		return err
	}

	if err := os.MkdirAll("applications/usecases/outputs", 0777); err != nil {
		return err
	}

	if err := os.MkdirAll("applications/interactors", 0777); err != nil {
		return err
	}
	return nil
}
