package main

func stats(email string) {
  commits := processRepositories(email)
  printCommitsStats(commits)
}

func processRepositories(email string) map[int]int {
  filePath := getDotFilePath()
  repos := parseFileLinesToSlice(filePath)
  daysInMap := daysInLastSixMonths

  commits := make(map[int]int, daysInMap)
  for i := daysInMap; i > 0; i-- {
    commits[i] = 0
  }

  for _, path := range repos {
    commits = fillCommits(email, path, commits)
  }

  return commits
}

func fillCommits(email string, path string, commits map[int]int) map[int]int {
  // instantiate a git repo object from path 
  repo, err := git.PlainOpen(path)
  if err != nil {
    panic(err)
  }
  // get the Head reference
}
