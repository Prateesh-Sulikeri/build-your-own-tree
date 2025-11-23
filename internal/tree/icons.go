package tree

import "strings"

func IconFor(name string, isDir bool) string {
	if isDir {
		return "📁 "
	}

	lower := strings.ToLower(name)

	switch {
	case strings.HasSuffix(lower, ".go"):
		return "🐹 "
	case strings.HasSuffix(lower, ".js"):
		return "📟 "
	case strings.HasSuffix(lower, ".ts"):
		return "🔷 "
	case strings.HasSuffix(lower, ".rs"):
		return "🦀 "
	case strings.HasSuffix(lower, ".py"):
		return "🐍 "
	case strings.HasSuffix(lower, ".c"), strings.HasSuffix(lower, ".h"):
		return "💻 "
	case strings.HasSuffix(lower, ".cpp"), strings.HasSuffix(lower, ".hpp"):
		return "💻 "
	case strings.HasSuffix(lower, ".html"), strings.HasSuffix(lower, ".htm"):
		return "🌐 "
	case strings.HasSuffix(lower, ".md"), strings.HasSuffix(lower, ".txt"):
		return "📝 "
	case strings.HasSuffix(lower, ".json"), strings.HasSuffix(lower, ".yaml"), strings.HasSuffix(lower, ".yml"), strings.HasSuffix(lower, ".toml"):
		return "📝 "
	case strings.HasSuffix(lower, ".png"), strings.HasSuffix(lower, ".jpg"), strings.HasSuffix(lower, ".jpeg"), strings.HasSuffix(lower, ".svg"), strings.HasSuffix(lower, ".gif"), strings.HasSuffix(lower, ".ico"):
		return "🖼️ "
	case strings.HasSuffix(lower, ".mp3"), strings.HasSuffix(lower, ".wav"), strings.HasSuffix(lower, ".flac"):
		return "🎵 "
	case strings.HasSuffix(lower, ".mp4"), strings.HasSuffix(lower, ".mov"), strings.HasSuffix(lower, ".mkv"):
		return "🎬 "
	case strings.HasSuffix(lower, ".zip"), strings.HasSuffix(lower, ".tar"), strings.HasSuffix(lower, ".gz"), strings.HasSuffix(lower, ".rar"):
		return "📦 "
	case strings.HasSuffix(lower, ".exe"), strings.HasSuffix(lower, ".bin"), strings.HasSuffix(lower, ".sh"):
		return "⚙️ "
	default:
		return "📄 "
	}
}

