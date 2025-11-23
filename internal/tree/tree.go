package tree

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Options controls PrintTree behavior.
type Options struct {
	All         bool   // show hidden files
	DirsOnly    bool   // show directories only
	MaxDepth    int    // 0 = unlimited, 1 = current directory only
	FullPath    bool   // print full path instead of name
	IgnoreGlob  string // single glob pattern to ignore (shell-style)
	DirsFirst   bool   // list directories before files
	FollowLinks bool   // follow symlinks to directories
}

// PrintTree is the entrypoint
func PrintTree(root string, opts Options) error {
	// start with depth 1 (like real tree: -L 1 means only top-level)
	return printLevel(root, "", 1, opts)
}

func printLevel(root, prefix string, depth int, opts Options) error {
	entries, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	// STEP 1 — Filter hidden files + ignore patterns
	filtered := make([]os.DirEntry, 0, len(entries))
	for _, e := range entries {
		name := e.Name()

		// skip hidden (unless -a)
		if !opts.All && strings.HasPrefix(name, ".") {
			continue
		}

		// match ignore glob (e.g., -I "*.log")
		if opts.IgnoreGlob != "" {
			match, _ := filepath.Match(opts.IgnoreGlob, name)
			if match {
				continue
			}
		}

		filtered = append(filtered, e)
	}

	// STEP 2 — If dirsOnly, filter out non-directories
	if opts.DirsOnly {
		tmp := filtered[:0]
		for _, e := range filtered {
			isDir := e.IsDir()
			if !isDir && opts.FollowLinks {
				full := filepath.Join(root, e.Name())
				if fi, err := os.Stat(full); err == nil {
					if fi.IsDir() {
						isDir = true
					}
				}
			}
			if isDir {
				tmp = append(tmp, e)
			}
		}
		filtered = tmp
	}

	// STEP 3 — Sorting (alphabetical, or dirs first)
	sort.SliceStable(filtered, func(i, j int) bool {
		if opts.DirsFirst {
			di := isDirMaybe(filtered[i], root, opts)
			dj := isDirMaybe(filtered[j], root, opts)
			if di != dj {
				return di && !dj // directories come before files
			}
		}
		return strings.ToLower(filtered[i].Name()) < strings.ToLower(filtered[j].Name())
	})

	// STEP 4 — Print each entry
	for i, e := range filtered {
		isLast := i == len(filtered)-1
		conn := "├── "
		if isLast {
			conn = "└── "
		}

		// Determine if directory (follow symlinks if -l)
		isDir := e.IsDir()
		if !isDir && opts.FollowLinks {
			full := filepath.Join(root, e.Name())
			if fi, err := os.Stat(full); err == nil {
				if fi.IsDir() {
					isDir = true
				}
			}
		}

		// ICONS HERE ↓↓↓
		icon := IconFor(e.Name(), isDir)

		// Full path or only name?
		name := e.Name()
		if opts.FullPath {
			name = filepath.Join(root, name)
		}

		display := icon + name
		fmt.Println(prefix + conn + display)

		// STEP 5 — Recurse (if directory + depth allowed)
		if isDir && (opts.MaxDepth == 0 || depth < opts.MaxDepth) {
			nextPrefix := prefix
			if isLast {
				nextPrefix += "    "
			} else {
				nextPrefix += "│   "
			}

			if err := printLevel(filepath.Join(root, e.Name()), nextPrefix, depth+1, opts); err != nil {
				return err
			}
		}
	}

	return nil
}

// isDirMaybe checks whether entry is a directory
// optionally following symbolic links if -l
func isDirMaybe(e os.DirEntry, parent string, opts Options) bool {
	if e.IsDir() {
		return true
	}
	if opts.FollowLinks {
		full := filepath.Join(parent, e.Name())
		if fi, err := os.Stat(full); err == nil {
			return fi.IsDir()
		}
	}
	return false
}

