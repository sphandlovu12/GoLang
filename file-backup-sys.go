// File backup system that prompts the user for the source and backup directories, then resumes the backup to the given paths

package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return destFile.Sync()
}

func backupFiles(sourceDir, backupDir string) error {
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}

		backupPath := filepath.Join(backupDir, relativePath)

		if info.IsDir() {
			return os.MkdirAll(backupPath, info.Mode())
		}

		return copyFile(path, backupPath)
	})

	return err
}

func main() {
	var sourceDir, backupDir string

	fmt.Print("Enter the source directory: ")
	fmt.Scanln(&sourceDir)

	fmt.Print("Enter the backup directory: ")
	fmt.Scanln(&backupDir)

	err := backupFiles(sourceDir, backupDir)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Backup completed successfully.")
	}
}
