package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var tools = map[string]string{
	".mp4": "ffmpeg", ".mkv": "ffmpeg", ".webm": "ffmpeg", ".mov": "ffmpeg",
	".mp3": "ffmpeg", ".flac": "ffmpeg", ".wav": "ffmpeg", ".ogg": "ffmpeg",
	".gif": "ffmpeg",
	".png": "magick", ".jpg": "magick", ".jpeg": "magick", ".webp": "magick",
	".md": "pandoc", ".docx": "pandoc", ".html": "pandoc", ".pdf": "pandoc",
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println(os.Stderr, "usage: conv <input> <output>")
		os.Exit(1)
	}
	in, out := os.Args[1], os.Args[2]		// input and output paths

	ext := strings.ToLower((filepath.Ext(out)))
	tool, ok := tools[ext]		// get proper tooling piece
	if !ok {
		fmt.Fprint(os.Stderr, "conv: don't know how to make %q files\n", ext)	// dunno what that is
		os.Exit(1)
	}

	var args []string
	switch tool {
	case "ffmpeg":
		args = []string{"-i", in , out}
	case "magick": 
		args = []string{in, out}
	case "pandoc":
		args = []string{in, "-o", out}
	}

	cmd := exec.Command(tool, args...)
	cmd.Stdout = os.Stdout				// show tool's own output
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "conv:", err)
		os.Exit(1)
	}

}
