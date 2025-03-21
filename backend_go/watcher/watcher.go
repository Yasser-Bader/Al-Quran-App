package watcher

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/fsnotify/fsnotify"
)

func Fun_watcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer watcher.Close()

	// مجلد المشروع الحالي
	dir := "./"

	// أضف المجلد إلى الـ Watcher
	err = watcher.Add(dir)
	if err != nil {
		fmt.Println("Error adding directory:", err)
		return
	}

	fmt.Println("Watching for file changes...")

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			// تأكد أن الملف تم تعديله
			if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Println("File modified:", event.Name)
				restartServer()
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("Watcher error:", err)
		}
	}
}

func restartServer() {
	fmt.Println("Restarting server...")

	// إيقاف أي عملية قديمة
	exec.Command("pkill", "-f", "main").Run()

	// انتظر قليلًا ثم شغل التطبيق مرة أخرى
	time.Sleep(1 * time.Second)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error restarting server:", err)
	}
}
