package main

import (
  "os"
  "regexp"
  "path/filepath"
  "os/exec"
)

func main() {
  out, err := exec.Command("ginkgo", "-cover", "-r").Output()
  if nil != err {
    print(err)
  } else {
    print(string(out))
  }
  const allCover string = "all_coverage.coverprofile"
  var coverprofiles []string
  path, _ := filepath.Abs(filepath.Dir(""))
  filepath.Walk(path, func (path string, f os.FileInfo, err error) error {
    if nil != err {
      return err
    }
    if !f.IsDir() && allCover != f.Name() && ".coverprofile" == filepath.Ext(path) {
      coverprofiles = append(coverprofiles, path)
    }
    return nil
  })
  out, err = exec.Command("cat", coverprofiles...).Output()
  if nil != err {
    panic(err)
  }
  allCoverFile, err := os.OpenFile(path + string(os.PathSeparator) + allCover,
    os.O_CREATE | os.O_TRUNC | os.O_WRONLY, 0655)
  defer allCoverFile.Close()
  reg := regexp.MustCompile(`mode: atomic\n`)
  allCoverFile.WriteString("mode: atomic\n" + reg.ReplaceAllString(string(out), ""))
  err = exec.Command("go", "tool",
    "cover", "-html=" + path + string(os.PathSeparator) + allCover, "-o", "coverage.html").Run()
  if nil != err {
    panic(out)
  }
}
