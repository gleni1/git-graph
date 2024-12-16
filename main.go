package main 

import (
  "fmt"
  "os"
)

func main() {
  folderPath := "/Users/mariglenpoleshi/workspace/github.com/gleni1/pokedex"
  
  // Read the Folder1
  entries, err := os.ReadDir(folderPath)
  if err != nil {
    fmt.Println("Error reading directory:", err)
    return
  }
  
  for _, entry := range entries {
    if entry.IsDir() {
      fmt.Println("Directory: ", entry.Name())
    } else {
      fmt.Println("File: ", entry.Name())
    }
  }
}
