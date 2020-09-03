package collection

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

const examplePics = "testFolder"

func TestRead(t *testing.T) {
	if err := createTestFolder(); err != nil {
		fmt.Println(err)
	}
	if err := createValidImages(5); err != nil {
		fmt.Println(err)
	}
}

func TestSet_Name(t *testing.T) {
	s, err := Read(examplePics)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s.Name)
}
func TestReadAll(t *testing.T) {
	sets, errs, err := ReadAll("../logo")
	if err != nil {
		t.Error(err)
	}
	for _, e := range errs {
		fmt.Println(e)
	}
	for _, s := range sets {
		fmt.Println(s.Name)
	}
	for _, s := range sets {
		fmt.Println(s)
	}
}
func createTestFolder() error {
	if _, err := os.Stat(examplePics); os.IsNotExist(err) {
		if err := os.Mkdir(examplePics, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create test folder")
		}
	} else if err == nil {
		fis, err := ioutil.ReadDir(examplePics)
		if err != nil {
			return err
		}
		if len(fis) > 0 {
			//TODO check whether it's test logo
			fmt.Println("test folder isn't empty")
		}
		return nil
	}
	return nil
}

func deleteTestFolder() error {
	err := os.RemoveAll(examplePics)
	return err
}

func createValidImages(n int) error {
	for i := 0; i < n; i++ {
		file, err := os.Create(filepath.Join(examplePics, fmt.Sprintf("valid%d.jpg", i)))
		if err != nil {
			return err
		}
		if err := file.Close(); err != nil {
			return err
		}
	}
	return nil
}
