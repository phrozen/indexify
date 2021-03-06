package main

import (
	"fmt"
	"os"
	"strings"
)

// FormatBytes returns a string representing the size in
// readable format, either KB, MB, or GB.
func FormatBytes(n int64) string {
	if n < 1e3 {
		return fmt.Sprintf("%d B", n)
	}

	f := float64(n)
	if (f / 1024) < 1000 {
		return fmt.Sprintf("%.2F KB", f/1024)
	}

	if (f / (1024 * 1024)) < 1000 {
		return fmt.Sprintf("%.2F MB", f/(1024*1024))
	}

	if (f / (1024 * 1024 * 1024)) < 1000 {
		return fmt.Sprintf("%.2F GB", f/(1024*1024*1024))
	}

	return fmt.Sprintf("%.2F TB", f/(1024*1024*1024*1024))
}

// DetectType checks file extension to check for
// known file types to select an icon class.
func DetectType(filename string) string {
	str := strings.Split(filename, ".")
	ext := strings.ToLower(str[len(str)-1])
	switch ext {
	case "txt", "rtf", "md", "doc", "docx", "pdf", "html":
		return "text"
	case "jpg", "jpeg", "gif", "png", "tiff", "bmp":
		return "image"
	case "mp3", "ogg", "oga", "m4a", "acc", "wma", "wav", "flac":
		return "audio"
	case "mp4", "m4v", "mov", "avi", "mkv", "wmv", "mpg", "flv", "mpeg":
		return "video"
	case "zip", "rar", "7z", "cab", "iso", "tar", "gz", "bz2":
		return "compressed"
	default:
		return "file"
	}
}

// NameFromPath extracts the name of the current Folder
func NameFromPath(path string) string {
	name, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	return name.Name()
}
