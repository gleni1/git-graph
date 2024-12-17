package main

import (
)

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

func fillCommits(email string, path string, commmits map[int]int) map[int]int {
  // instantiate a git repo object from path 
  repo, err := git.PlainOpen(path)
  if err != nil {
    panic(err)
  }

  // get the HEAD reference 
  ref, err := repo.Head()
  if err != nil {
    panic(err)
  }

  // get the commmits history starting from HEAD 
  iterator, err := repo.Log(&git.LogOptions{From: ref.Hash()})
  if err != nil {
    panic(err)
  }

  // iterate the commits
  offset := calcOffset()
  err = iterator.ForEach(func(c *object.Commit) error) {
    daysAgo := countDaysSinceDate(c.Author.When) + offset 

    if c.Author.Email != email {
      return nil 
    }

    if daysAgo != outOfRange {
      commits[daysAgo]++
    }

    return nil
  })
  
  if err != nil {
    panic(err)
  }
}

func getBeginningOfDay(t time.Time) time.Time {
  year, month, day := t.Date()
  startOfDay := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
  return startOfDay
}

func countDaysSinceDate(date time.Time) int {
  days := 0
  now := getBeginningOfDay(time.Now())
  for date.Before(now) {
    date = date.Add(time.Hour * 24)
    days++
    if days > daysInLastSixMonths {
      return outOfRange 
    }
  }
  return days
}
