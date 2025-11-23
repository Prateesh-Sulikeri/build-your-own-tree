package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Prateesh-Sulikeri/build-your-own-tree/internal/tree"
)

func main() {
	// Flags
	all := flag.Bool("a", false, "All files are printed (include hidden files).")
	dirsOnly := flag.Bool("d", false, "List directories only.")
	maxDepth := flag.Int("L", 0, "Max display depth of the directory tree (1 = current dir only). 0 = unlimited.")
	fullPath := flag.Bool("f", false, "Print the full path prefix for each file.")
	ignore := flag.String("I", "", "Ignore files that match the glob pattern (e.g. \"*.log\").")
	dirsFirst := flag.Bool("dirsfirst", false, "List directories before files.")
	followLinks := flag.Bool("l", false, "Follow symbolic links (when they point to directories).")

	flag.Parse()

	path := "."
	if flag.NArg() > 0 {
		path = flag.Arg(0)
	}

	opts := tree.Options{
		All:        *all,
		DirsOnly:   *dirsOnly,
		MaxDepth:   *maxDepth,
		FullPath:   *fullPath,
		IgnoreGlob: *ignore,
		DirsFirst:  *dirsFirst,
		FollowLinks: *followLinks,
	}

	if err := tree.PrintTree(path, opts); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

