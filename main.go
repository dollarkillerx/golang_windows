package main

//
import "C"
import (
	"github.com/axgle/mahonia"
	"io/ioutil"
	"log"
	"net/http"
)

//export HttpGet
func HttpGet(url string,c **C.char) {
	body, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer body.Body.Close()

	all, err := ioutil.ReadAll(body.Body)
	if err != nil {
		log.Println(err)
		return
	}

	ToNewGBKCStr(string(all),c)
}

//export ToNewGBKCStr
func ToNewGBKCStr(str string,c **C.char) {
	// generate GBK str
	enc := mahonia.NewEncoder("GBK")
	output := enc.ConvertString(str)
	*c = C.CString(output)
}

func main() {}  // main 函数必须有
