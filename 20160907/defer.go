package main

import ()

func createFile() {
	out, err := os.Create(path)
	if err != nil {
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return
	}
}
