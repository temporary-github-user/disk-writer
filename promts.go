package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

func AskForFilesystem(partitions []FS) (int, string, error) {

	templates := &promptui.SelectTemplates{
		Active:   "{{ `Device:`| green}}{{ .Device | green }}\t{{ `MountedOn:`| green}}{{ .MountedOn | green }}\t{{ `Available:`| green}}{{ .FreeSpaceInMB | green }}{{ `MB`| green}}",
		Inactive: "Device:{{ .Device | white }}\tMountedOn:{{ .MountedOn | white }}\tAvailable:{{ .FreeSpaceInMB | white }}MB",
		Selected: "{{ `Device:`| green}}{{ .Device | green }}\t{{ `MountedOn:`| green}}{{ .MountedOn | green }}\t{{ `Available:`| green}}{{ .FreeSpaceInMB | green }}{{ `MB`| green}}",
	}

	prompt2 := promptui.Select{
		Label:     "Select Filesystem",
		Items:     partitions,
		Templates: templates,
	}

	return prompt2.Run()
}

func AskForNumberOfFiles() (int, error) {

	validate := func(input string) error {
		_, err := strconv.Atoi(input)
		if err != nil {
			return errors.New("invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Enter the number of files to be created on the filesystem",
		Validate: validate,
	}

	res, err := prompt.Run()
	if err != nil {
		return 0, err
	}

	nFiles, err := strconv.Atoi(res)
	if err != nil {
		return 0, err
	}

	return nFiles, nil
}

func askForSize() (int, error) {

	validate := func(input string) error {
		_, err := strconv.Atoi(input)
		if err != nil {
			return errors.New("invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Enter the size for every file (in Megabytes as a whole number)",
		Validate: validate,
	}

	res, err := prompt.Run()
	if err != nil {
		return 0, err
	}

	size, err := strconv.Atoi(res)
	if err != nil {
		return 0, err
	}

	return size, nil
}

func confirm(nFiles, size int, fs, mountedOn string) bool {
	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("%d files of size %d will be created in \"%s\" filesystem which is mounted on %s", nFiles, size, fs, mountedOn),
		IsConfirm: true,
	}

	res, _ := prompt.Run()

	return strings.ToLower(res) == "y"
}
