package main 

import (
  "fmt"
  "os"
  "bufio"
  "io"
  "io/ioutil"
  "log"
  "os/user"
  "strings"
)

func scanGitFolders(folders []string, folder string) []string {
  // trim the last '/'
  folder = strings.TrimSuffix(folder, "/")
 
  f, err := os.Open(folder)
  if err != nil {
    log.Fatal(err)
  }

  files, err := f.Readdir(-1)
  f.Close()
  if err != nil {
    log.Fatal(err)
  }

  var path string 

  for _, file := range files {
    if file.IsDir() {
      path = folder + "/" + file.Name()
      fi file.Name() == "git" {
        path = strings.TrimSuffix(path, "/.git")
        fmt.Println(path)
        folders = append(folders, path)
        continue
      }
      if file.Name() == "vendor" || file.Name() == "node_modules" {
        continue 
      }
      folders = scanGitFolders(folders, path)
    }
  }
  return folders
}


func recursiveScanFolder(folder string) []string {
  return scanGitFolders(make([]string, 0), folder)
}

// getDotFilePath returns the dot file for the repos list
// Creates it and the enclosing flder if it does not exist
func getDotFilePath() string {
  user, err := user.Current()
  if err != nil {
    log.Fatal(err)
  }

  dotFile := usr.HomeDir + "/.gogitlocalstats"
  
  return dotFile 
}

func addNewSliceElementsToFile(filePath string, newRepos []string) {
  existingRepos := parseFileLinesToSlice(filePath)
  repos := joinSlices(newRepos, existingRepos)
  dumpStringsSliceToFile(repos, filePath)
}

func parseFileLinesToSlice(filePath string) []string {
  f := openFile(filePath)
  defer f.Close()

  var lines []string 
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  if err := scanner.Err(); err != nil {
    if err != io.EOF {
      panic(err)
    }
  }
  return lines
}

func openFile(filePath string) *os.File {
  f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0755)
  if err != nil {
    if os.IsNotExist(err) {
      // file does not exist 
      _, err = os.Create(filePath)
      if err != nil {
        panic(err)
      }
    }else {
      panic(err)
    }
  }
  return f
}

func joinSlices(new []string, existing []string) []string {
  for _, i := range new {
    if !sliceContains(existing, i) {
      existing = append(existing, i)
    }
  }
  return existing
}

// sliceContains returns true if slice contains value
func sliceContains(slice []string, value string) bool {
  for _, v := range slice {
    if v == value {
      return true
    }
  }
  return false
}

// dumpStringsSliceToFile writes content to the file in path `filePath` (overwriting existing content) 
func dumpStringsSliceToFile(repos []string, filePath string){
  content := strings.Join(repos, "\n")
  ioutil.WriteFile(filePath, []byte(content), 0755)
}