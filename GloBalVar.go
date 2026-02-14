package autotool

import "runtime"

const (
	MD5    CryptType = "MD5"
	SHA256 CryptType = "SHA256"
	SHA512 CryptType = "SHA512"
	BASE64 CryptType = "BASE64"
	UUID   CryptType = "UUID"
	AES    CryptType = "AES"
)

const (
	ERROR   LogType = "Error"
	INFO    LogType = "Info"
	WARNING LogType = "Warning"
	DEBUG   LogType = "Debug"
	PANIC   LogType = "Panic"
)

const (
	Black  ColorType = "Black"
	White  ColorType = "White"
	Red    ColorType = "Red"
	Green  ColorType = "Green"
	Yellow ColorType = "Yellow"
	Blue   ColorType = "Blue"
	Purple ColorType = "Pueple"
)

var logs *int = new(int)

const letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const sysGOOS string = runtime.GOOS

var Colors = map[ColorType]string{
	Black:  "30",
	White:  "37",
	Red:    "31",
	Green:  "32",
	Yellow: "33",
	Blue:   "34",
	Purple: "35",
}

var SuffixTypes map[string][]string = map[string][]string{
	"应用": {
		"exe", "", "bin", "jar",
	},
	"视频": {
		"mp4", "avi", "dv", "mpeg", "mov", "m4v", "flv", "wmv", "mkv", "asf", "rm", "vob",
	},
	"音频": {
		"mp3", "wma", "wav", "midi", "ape", "flac", "ogg", "au", "m4a", "mka", "aiff",
	},
	"图片": {
		"xbm", "tif", "pjp", "svgz", "jpg", "jpeg", "ico", "tiff", "gif", "svg", "jfif", "webp", "png", "bmp", "pjpeg", "avif",
	},
	"文本": {
		"txt",
	},
	"压缩文件": {
		"zip", "rar", "rar4", "tar", "gz",
	},
	"字体": {
		"woff2", "woff", "ttf",
	},
	"临时文件": {
		"br",
	},
}
