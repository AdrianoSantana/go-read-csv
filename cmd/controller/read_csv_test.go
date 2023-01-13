package controller

import "testing"

func TestReadCsv(t *testing.T) {
	t.Run("Should return an error if no file is found it", func(t *testing.T) {
		filePath := "../movies"
		_, err := ReadCSV(filePath)

		if err == nil {
			t.Errorf("it was expected an error")
		}
	})
}
