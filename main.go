package main

import (
	"fmt"
	"github.com/reujab/wallpaper"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"github.com/schollz/progressbar/v3"
	"github.com/ttacon/chalk"
  "runtime"
  "time"
  "os/exec"
)

// open opens the specified URL in the default browser of the user.
func open(url string) error {
  var cmd string
  var args []string

  switch runtime.GOOS {
  case "windows":
      cmd = "cmd"
      args = []string{"/c", "start"}
  case "darwin":
      cmd = "open"
  default: // "linux", "freebsd", "openbsd", "netbsd"
      cmd = "xdg-open"
  }
  args = append(args, url)
  return exec.Command(cmd, args...).Start()
}
func getPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}
func downloadFile(filepath string, url string, downloadText string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"Downloading our wallpaper",
	)
	// Writer the body to file
	_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
	if err != nil {
		return err
	}

	return nil
}
func main() {
  fmt.Println(chalk.Red.Color("A love virus was detected on your computer. It will explode"))
  time.Sleep(30 * time.Second)
	fmt.Println(chalk.Magenta.Color("Just kidding, I love you 🥰\nThis a little program I wrote for you"))
	var cacheDir string = getPath() + "/.cache"

  // Download Wallpaper
	if _, err := os.Stat(cacheDir); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
      os.Mkdir(cacheDir, os.ModePerm);
      downloadFile(cacheDir + "/background.png", "https://docs.cloud.kabeers.network/c/v/643fd05bd08eb---Group%201.png", "Downloading Wallpaper");
		}
	}

  // Set Wallpaper
  wallpaper.SetFromFile(cacheDir + "/background.png")
	fmt.Println(chalk.Magenta.Color("Go check your wallpaper cutie"))

  // Open Website Image
  time.AfterFunc(2 * time.Minute, func () {
    open("https://i.pinimg.com/564x/d4/14/88/d4148825522fcdf4c99298d5954c5ea2.jpg")
    fmt.Println(chalk.Green.Color("btw the image thing was also me <3"))
  })
}
