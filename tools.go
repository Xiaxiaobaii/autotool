package autotool

import (
	"bytes"
	"crypto/aes"
	"crypto/md5"
	crand "crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"math/rand"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/kardianos/service"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"

	"github.com/google/uuid"

	regexp "github.com/wasilibs/go-re2"
)

func Error(text string) error {
	return &errorString{text}
}

func (e *errorString) Error() string {
	return e.s
}

func Itoa(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func Atoi(a string) (int, error) {
	return strconv.Atoi(a)
}

func BtoS(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

func StoB(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func SubString(father, son string) bool {
	return strings.Contains(father, son)
}

func Btom(b []byte) map[string]interface{} {
	var JsonR map[string]interface{}
	json.Unmarshal(b, &JsonR)
	return JsonR
}

func StructToMap(su any) (map[string]any, error) {
	jso, e := json.Marshal(su)
	if e != nil {
		return nil, e
	}
	data := make(map[string]any)
	e = json.Unmarshal(jso, &data)
	return data, e
}

func Logi() {
	*logs += 1
	LogPrint(Itoa(*logs), DEBUG, 2)
}

func ReGet(str string, re string) string {
	RE, _ := regexp.Compile(re)
	return RE.FindString(str)
}

// func CmdColorPrint(i int, s string, v ...any) {
// 	kernel32 := syscall.NewLazyDLL("kernel32.dll")
// 	proc := kernel32.NewProc("SetConsoleTextAttribute")
// 	proc.Call(uintptr(syscall.Stdout), uintptr(i))
// 	fmt.Printf(s, v...)
// 	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(7))
// 	CloseHandle := kernel32.NewProc("CloseHandle")
// 	CloseHandle.Call(handle)
// }

func LogPrint(str interface{}, Type LogType, caller int) {
	nowTime := time.Now()
	if caller == 9 {
		fmt.Printf("[Error %s] %s\n", nowTime.Format("2006-01-02 15:04:05"), str)
		Print(string(debug.Stack()))
		return
	}
	if _, file, line, ok := runtime.Caller(caller); ok {

		if sysGOOS == "windows" {
			switch Type {
			case ERROR:
				fmt.Printf("[Error %s %s:%d] %s\n", nowTime.Format("2006-01-02 15:04:05"), file, line, str)
				return
			case INFO:
				fmt.Printf("[Info %s] %s\n", nowTime.Format("2006-01-02 15:04:05"), str)
				return
			case WARNING:
				fmt.Printf("[Warning %s %s] %s\n", nowTime.Format("2006-01-02 15:04:05"), file, str)
				return
			case DEBUG:
				fmt.Printf("[Debug %s:%d] %s\n", file, line, str)
				return
			case PANIC:
				fmt.Printf("[PanicError! %s %s:%d] %s\n", nowTime.Format("2006-01-02 15:04:05"), file, line, str)
				os.Exit(0)
			}
		} else {
			switch Type {
			case ERROR:
				fmt.Printf("\033[1;%s;40m[Error %s %s:%d] %s\033[0m\n", Colors[Red], nowTime.Format("2006-01-02 15:04:05"), file, line, str)
				return
			case INFO:
				fmt.Printf("\033[0;%s;40m[Info %s] %s\033[0m\n", Colors[Green], nowTime.Format("2006-01-02 15:04:05"), str)
				return
			case WARNING:
				fmt.Printf("\033[0;%s;40m[Warning %s %s] %s\033[0m\n", Colors[Yellow], nowTime.Format("2006-01-02 15:04:05"), file, str)
				return
			case DEBUG:
				fmt.Printf("\033[0;%s;40m[Debug %s:%d] %s\033[0m\n", Colors[Blue], file, line, str)
				return
			case PANIC:
				fmt.Printf("\033[1;%s;40m[PanicError! %s %s:%d] %s\033[0m\n", Colors[Red], nowTime.Format("2006-01-02 15:04:05"), file, line, str)
				os.Exit(0)
			}
		}

		fmt.Println("[error] TypeError:", Type)
	} else {
		fmt.Println("[error] create new error log grow runtime.Caller error,skip error")
		return
	}
}

func Ftoi(f float64) int {
	return int(math.Ceil(f - 0.5))
}

func LogSprint(str interface{}, Type LogType, caller int) string {
	if _, file, line, ok := runtime.Caller(caller); ok {
		switch Type {
		case ERROR:
			return fmt.Sprintf("[Error %s:%d] %s\n", file, line, str)
		case INFO:
			return fmt.Sprintf("[Info] %s\n", str)
		case WARNING:
			return fmt.Sprintf("[Warning %s] %s\n", file, str)
		case DEBUG:
			return fmt.Sprintf("[Debug %s:%d] %s\n", file, line, str)
		case PANIC:
			return fmt.Sprintf("[PanicError! %s:%d] %s\n", file, line, str)
		}
		return "[error] Type Error"
	} else {
		return "[error] create new error log grow runtime.Caller error,skip error"
	}
}

func ReadAll(body io.ReadCloser) []byte {
	if b, e := io.ReadAll(body); e == nil {
		return b
	} else {
		LogPrint("ReadError:"+e.Error(), "Info", 2)
	}
	return []byte("")
}

func ReadAllString(body io.ReadCloser) string {
	if b, e := io.ReadAll(body); e == nil {
		return string(b)
	} else {
		LogPrint("ReadError:"+e.Error(), "Info", 2)
	}
	return ""
}

func ReverseString(s string) string {
	a := func(s string) *[]rune {
		var b []rune
		for _, k := range s {
			defer func(v rune) {
				b = append(b, v)
			}(k)
		}
		return &b
	}(s)
	return string(*a)
}

func Encrypt(Type CryptType, str []byte, args ...string) string {
	var cry string

	switch Type {
	case MD5:
		var Ka []byte = make([]byte, 16)
		kas := md5.Sum(str)
		copy(Ka, kas[:])
		cry = hex.EncodeToString(Ka)
	case SHA256:
		var Ka []byte = make([]byte, 32)
		kas := sha256.Sum256(str)
		copy(Ka, kas[:])
		cry = hex.EncodeToString(Ka)
	case SHA512:
		var Ka []byte = make([]byte, 64)
		kas := sha512.Sum512(str)
		copy(Ka, kas[:])
		cry = hex.EncodeToString(Ka)
	case BASE64:
		cry = base64.StdEncoding.EncodeToString(str)
	case UUID:
		cry = uuid.New().String()
	case AES:
		cipher, _ := aes.NewCipher([]byte(args[0]))
		out := make([]byte, len(str))
		cipher.Encrypt(out, str)
		cry = string(out)
	}

	return cry
}

func Decrypt(Type CryptType, str string, args ...string) string {
	switch Type {
	case "BASE64":
		cry, e := base64.StdEncoding.DecodeString(str)
		if e != nil {
			cry = []byte{}
		}
		return string(cry)
	case AES:
		cipher, _ := aes.NewCipher([]byte(args[0]))
		out := make([]byte, len(str))
		cipher.Decrypt(out, []byte(str))

		return string(out)
	}
	return "Not Decrypt"
}

func FileNameDelSuffix(OldFileName string) string {
	SuffixNumber := len(strings.Split(OldFileName, ".")[len(strings.Split(OldFileName, "."))-1])
	return OldFileName[:len(OldFileName)-SuffixNumber-1]
}

func LC() string {
	if runtime.GOOS == "linux" {
		return "/"
	} else {
		return "\\"
	}
}

func ReadJson(path string) (map[string]interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var data = make(map[string]interface{})
	err = json.Unmarshal(ReadAll(file), &data)
	return data, err
}

func ReadIo(r io.ReadCloser) ([]byte, error) {
	b := make([]byte, 0, 512)
	for {
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return b, err
		}

		if len(b) == cap(b) {
			b = append(b, 0)[:len(b)]
		}
	}
}

func Mtoj(m map[string]interface{}) string {
	Json, _ := json.Marshal(m)
	return string(Json)
}

func FileNameGetSuffix(FileName string) string {
	if len(strings.Split(FileName, ".")) <= 1 {
		return ""
	}
	return strings.Split(FileName, ".")[len(strings.Split(FileName, "."))-1]
}

func I6toS(i int64) string {
	return strconv.FormatInt(i, 10)
}

func GetSuffixType(Suffix string) string {
	for k, v := range SuffixTypes {
		for _, i := range v {
			if i == Suffix {
				return k
			}
		}
	}
	return "其他"
}

func RandString(i int) string {
	var ks string
	for {
		ks += string(letters[rand.Intn(len(letters))])
		i -= 1
		if i == 0 {
			return ks
		}
	}

}

func Stof(str string) float64 {
	r, _ := strconv.ParseFloat(str, 64)
	return r
}

func Atof(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

func Ftos(flo float64) string {
	return strconv.FormatFloat(flo, 'f', 2, 64)
}

func Jtom(str string) map[string]interface{} {
	Re := make(map[string]interface{})
	if str == "" {
		return Re
	}
	_ = json.Unmarshal([]byte(str), &Re)
	return Re
}

func Stom(str string) map[string]interface{} {
	var A interface{} = str
	return A.(map[string]interface{})
}

func UrlFormat(path string) string {
	if path[len(path)-1:] != "/" && path[len(path)-1:] != "*" {
		path = path + "/"
	}
	if path[0:1] != "/" {
		path = "/" + path
	}
	return path
}

func RandIntT(ra int64) int64 {
	if ra <= 0 {
		return 0
	}
	result, _ := crand.Int(crand.Reader, big.NewInt(ra))
	return result.Int64()
}

var RandC = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandInt(ra int) int {
	return RandC.Intn(ra)
}

func HashAdd(key string) int {
	var hash int64
	var modulus int64 = 10000000991
	for _, b := range []byte(key) {
		hash = (hash + int64(b)) % modulus
	}
	return int(hash)
}

func Stof3(str string) float32 {
	r, _ := strconv.ParseFloat(str, 32)
	return float32(r)
}

func Dsend(str *string, i int) {
	if len(*str) < i {
		return
	}
	*str = (*str)[0 : len(*str)-i]
}

func Fsend(str *string, i int) {
	if len(*str) < i {
		return
	}
	*str = (*str)[i:]
}

func IntoS(in any) string {
	switch v := in.(type) {
	case bool:
		return Btos(v)
	case int:
		return Itoa(v)
	case float64:
		return Ftos(v)
	case float32:
		return Ftos(float64(v))
	case string:
		return v
	case int64:
		return I6toS(v)
	case int32:
		return I6toS(int64(v))
	case map[string]any:
		by, _ := json.Marshal(v)
		return string(by)
	case []byte:
		return string(v)
	default:
		return ""
	}
}

func Btos(b bool) string {
	return strconv.FormatBool(b)
}

func Print(str ...any) {
	_, _ = fmt.Println(str...)
}

func NewKvO[T any](key, value T) Kvalue[T] {
	return Kvalue[T]{
		key:   key,
		value: value,
	}
}

var Flags map[string]func(str string) = map[string]func(str string){}

func FlagCall(cmd string, call func(str string)) {
	for i := range os.Args {
		if os.Args[i] != cmd {
			continue
		}
		if len(os.Args) <= i+1 {
			call("")
			return
		}
		call(os.Args[i+1])
		return
	}
}

func FlagVar(cmd string, arg *string, def string) {
	cmd = "-" + cmd
	for i := range os.Args {
		if os.Args[i] != cmd {
			continue
		}
		if len(os.Args) <= i+1 {
			*arg = def
			return
		}
		*arg = os.Args[i+1]
		return
	}
	*arg = def
}

func RandSliceValue[T interface{}](xs []T) T {
	return xs[rand.Intn(len(xs))]
}

func HelpTool(commands map[string][]string) {
	for k, v := range commands {
		fmt.Println()
		ll := len("  " + k)
		query := "  " + k
		for {
			if ll < 20 {
				query += " "
				ll += 1
			} else {
				break
			}
		}
		fmt.Printf("%s%s\n", query, v[0])
		v = v[1:]
		for _, str := range v {
			fmt.Printf("                    %s\n", str)
		}
	}
}

func ReadStructConfig[T any](filePath string, Stc *T) error {
	File, err := os.Open(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(ReadAll(File), Stc)
	return err
}

func GetPathName(path string) string {
	path = strings.Replace(path, "\\", "/", -1)
	Path := strings.Split(path, "/")
	return Path[len(Path)-1:][0]
}

func PathNameNc(path string) string {
	path = GetPathName(path)
	return strings.Split(path, ".")[:len(strings.Split(path, "."))][0]
}

func PathNameNd(path string) string {
	path = GetPathName(path)
	return strings.Split(path, ".")[:len(strings.Split(path, "."))][1]
}

func Sleep(i int) {
	for i > 0 {
		time.Sleep(time.Millisecond * 100)
		i -= 1
	}
}

func CopySlice[T any](Slice []T) []T {
	result := make([]T, len(Slice))
	copy(result, Slice)
	return result
}

func (kv *Kvalue[T]) Set(key, value T) {
	kv.key = key
	kv.value = value
}

func (kv *Kvalue[T]) Get() (T, T) {
	return kv.key, kv.value
}

func ReadConfig(Path string) (map[string]interface{}, error) {
	ReturnMap := make(map[string]interface{})
	File, err := os.Open(Path)
	if err != nil {
		return nil, Error("No Config")
	}
	FileByte := ReadAll(File)
	err = json.Unmarshal(FileByte, &ReturnMap)
	if err != nil {
		return nil, err
	}
	return ReturnMap, nil
}

func Atoix(str string) int {
	i, _ := Atoi(str)
	return i
}

func ReadConFigGet(Path string, OptionConfig *map[string]*string, AppConfig *map[string]map[string]*map[string]interface{}) error {
	Config, err := ReadConfig(Path)
	if err != nil {
		return err
	}
	for k, v := range Config {
		for g := range *OptionConfig {
			if k == g {
				*(*OptionConfig)[g] = IntoS(v)
			}
		}
		if *AppConfig == nil {
			for AppK, AppV := range *AppConfig {
				if k == AppK {
					App := v.(map[string]interface{})
					for k, v := range App {
						for g, i := range AppV {
							if k == g {
								for k, v := range v.(map[string]interface{}) {
									for g := range *i {
										if k == g {
											(*i)[g] = v
										}
									}
								}
								*i = v.(map[string]interface{})
							}
						}
					}
				}
			}
		}
	}
	return nil
}

func RunService(Name, Description string, Run func()) {
	svcConfig := &service.Config{
		Name:        Name,
		DisplayName: Name,
		Description: Description,
	}
	prg := &ServiceLess{}
	prg.Run = Run
	s, err := service.New(prg, svcConfig)
	if err != nil {
		LogPrint(err, PANIC, 2)
	}
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}

// func bucketSort(nums []float64) {
// 	k := len(nums) / 2
// 	buckets := make([][]float64, k)
// 	for i := 0; i < k; i++ {
// 		buckets[i] = make([]float64, 0)
// 	}
// 	for _, num := range nums {
// 		i := int(num * float64(k))
// 		buckets[i] = append(buckets[i], num)
// 	}
// }

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
