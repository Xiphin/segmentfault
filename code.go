package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

const (
	zero  = byte('0')
	one   = byte('1')
	lsb   = byte('[') // left square brackets
	rsb   = byte(']') // right square brackets
	space = byte(' ')
)

var uint8arr [8]uint8

// ErrBadStringFormat represents a error of input string's format is illegal .
var ErrBadStringFormat = errors.New("bad string format")

// ErrEmptyString represents a error of empty input string.
var ErrEmptyString = errors.New("empty string")

func init() {
	uint8arr[0] = 128
	uint8arr[1] = 64
	uint8arr[2] = 32
	uint8arr[3] = 16
	uint8arr[4] = 8
	uint8arr[5] = 4
	uint8arr[6] = 2
	uint8arr[7] = 1
}

// append bytes of string in binary format.
func appendBinaryString(bs []byte, b byte) []byte {
	var a byte
	for i := 0; i < 8; i++ {
		a = b
		b <<= 1
		b >>= 1
		switch a {
		case b:
			bs = append(bs, zero)
		default:
			bs = append(bs, one)
		}
		b <<= 1
	}
	return bs
}

// ByteToBinaryString get the string in binary format of a byte or uint8.
func ByteToBinaryString(b byte) string {
	buf := make([]byte, 0, 8)
	buf = appendBinaryString(buf, b)
	return string(buf)
}

// BytesToBinaryString get the string in binary format of a []byte or []int8.
func BytesToBinaryString(bs []byte) string {
	l := len(bs)
	bl := l*8 + l + 1
	buf := make([]byte, 0, bl)
	buf = append(buf, lsb)
	for _, b := range bs {
		buf = appendBinaryString(buf, b)
		buf = append(buf, space)
	}
	buf[bl-1] = rsb
	return string(buf)
}

// regex for delete useless string which is going to be in binary format.
var rbDel = regexp.MustCompile(`[^01]`)

// BinaryStringToBytes get the binary bytes according to the
// input string which is in binary format.
func BinaryStringToBytes(s string) (bs []byte) {
	if len(s) == 0 {
		panic(ErrEmptyString)
	}

	s = rbDel.ReplaceAllString(s, "")
	l := len(s)
	if l == 0 {
		panic(ErrBadStringFormat)
	}

	mo := l % 8
	l /= 8
	if mo != 0 {
		l++
	}
	bs = make([]byte, 0, l)
	mo = 8 - mo
	var n uint8
	for i, b := range []byte(s) {
		m := (i + mo) % 8
		switch b {
		case one:
			n += uint8arr[m]
		}
		if m == 7 {
			bs = append(bs, n)
			n = 0
		}
	}
	return
}

func get(urlStr string) (string, error) {
	var err error
	var resp *http.Response
	client := &http.Client{}

	req, err := http.NewRequest("GET", urlStr, nil)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	req.Header.Add("Referer", "https://1111.segmentfault.com")
	req.Header.Add("User-agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36")

	resp, err = client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bodyByte), nil
}

func post(urlStr string, params map[string]string, headers map[string]string) (string, error) {
	var err error
	var resp *http.Response
	client := &http.Client{}

	v := url.Values{}
	for key, val := range params {
		v.Add(key, val)
	}
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(v.Encode()))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36")
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	} else {
		req.Header.Add("Referer", "https://1111.segmentfault.com")
	}

	resp, err = client.Do(req)

	if err != nil || resp == nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(bodyByte), nil
}

func getResponseHeaderK(urlStr string, params map[string]string, headers map[string]string) (string, error) {
	var err error
	var resp *http.Response
	client := &http.Client{}

	v := url.Values{}
	for key, val := range params {
		v.Add(key, val)
	}
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(v.Encode()))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Referer", "https://1111.segmentfault.com")
	req.Header.Add("User-agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36")
	resp, err = client.Do(req)

	if err != nil || resp == nil {
		fmt.Println(err)
		return "", err
	}

	k := resp.Header.Get("The-Key-Is")
	return k, nil
}

var hurdle = map[string]func(urlStr string){
	"1": hurdleOne,
	"2": hurdleTwo,
	"3": hurdleThree,
	"4": hurdleFour,
	"5": hurdleFive,
	"6": hurdleSix,
	"7": hurdleSeven,
	"8": hurdleEight,
	"9": hurdleNine,
}

func hurdleOne(urlStr string) {
	segmentfaultHTML, err := get(urlStr)
	if err != nil {
		fmt.Println(err)
	}
	re := regexp.MustCompile(`<a style="color: #172024" href="\?k=(.+?)">`)
	find := re.FindStringSubmatch(segmentfaultHTML)
	fmt.Println("[=>]通往第 2 关的密码：", find[1])
	nextURL := "https://1111.segmentfault.com/?k=" + find[1]
	hurdleTwo(nextURL)
}

func hurdleTwo(urlStr string) {
	segmentfaultHTML, err := get(urlStr)
	if err != nil {
		fmt.Println(err)
	}
	re := regexp.MustCompile(`<!-- 不错嘛,密码在此:(.+?) -->`)
	find := re.FindStringSubmatch(segmentfaultHTML)
	fmt.Println("[=>]通往第 3 关的密码：", find[1])
	nextURL := "https://1111.segmentfault.com/?k=" + find[1]
	hurdleThree(nextURL)
}

func hurdleThree(urlStr string) {
	splitStr := strings.Split(urlStr, "?k=")

	k, err := getResponseHeaderK(splitStr[0], map[string]string{"k": splitStr[1]}, map[string]string{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("[=>]通往第 4 关的密码：", k)
	nextURL := "https://1111.segmentfault.com/?k=" + k
	hurdleFour(nextURL)
}

func hurdleFour(urlStr string) {
	splitStr := strings.Split(urlStr, "?k=")
	k := splitStr[1]
	h := md5.New()
	io.WriteString(h, "4")
	if k == hex.EncodeToString(h.Sum(nil)) {
		h.Reset()
		io.WriteString(h, "5")
		k = hex.EncodeToString(h.Sum(nil))
	}
	fmt.Println("[=>]通往第 5 关的密码：", k)
	nextURL := "https://1111.segmentfault.com/?k=" + k
	hurdleFive(nextURL)
}

func hurdleFive(urlStr string) {
	splitStr := strings.Split(urlStr, "?k=")
	segmentfaultHTML, err := get(urlStr)
	if err != nil {
		fmt.Println(err)
	}
	re := regexp.MustCompile(`<img src="(.+?)" />`)
	find := re.FindStringSubmatch(segmentfaultHTML)
	imgData, err := get(splitStr[0] + "/" + find[1])
	if err != nil {
		fmt.Println(err)
	}
	re = regexp.MustCompile(`\/KEY:(.+?)\/`)
	find = re.FindStringSubmatch(imgData)
	fmt.Println("[=>]通往第 6 关的密码：", find[1])
	nextURL := "https://1111.segmentfault.com/?k=" + find[1]
	hurdleSix(nextURL)
}

func hurdleSix(urlStr string) {
	segmentfaultHTML, err := get(urlStr)
	if err != nil {
		fmt.Println(err)
	}
	re := regexp.MustCompile(`<code>(.+?)</code>`)
	find := re.FindStringSubmatch(segmentfaultHTML)
	segmentfaultHTML, err = get("https://www.baidu.com/s?wd=" + find[1] + "&ie=UTF-8")
	re = regexp.MustCompile(`key: (.*?)腾讯微博,与其在别处仰望 不...`)
	find = re.FindStringSubmatch(segmentfaultHTML)
	fmt.Println("[=>]通往第 7 关的密码：", find[1])
	nextURL := "https://1111.segmentfault.com/?k=" + find[1]
	hurdleSeven(nextURL)
}

func hurdleSeven(urlStr string) {
	segmentfaultHTML, err := get(urlStr)
	if err != nil {
		fmt.Println(err)
	}
	re := regexp.MustCompile(`<code>(.+?)</code>`)
	find := re.FindStringSubmatch(segmentfaultHTML)
	fmt.Println("[=>]通往第 8 关的密码：", find[1])
	nextURL := "https://1111.segmentfault.com/?k=" + find[1]
	hurdleEight(nextURL)
}

func hurdleEight(urlStr string) {
	segmentfaultHTML, err := get(urlStr)
	if err != nil {
		fmt.Println(err)
	}
	re := regexp.MustCompile(`<input type="text" name="k" value="(.+?)" />`)
	find := re.FindStringSubmatch(segmentfaultHTML)

	nextURL := urlStr + "$post_k=" + find[1]
	fmt.Println("[=>]通往第 9 关的密码：", find[1])
	fmt.Print("\n[**]注意：第 9 关不能直接通过密码访问，需要制下面链接才能访问通关:\n\n" + nextURL + "\n\n")

	hurdleNine(nextURL)
}

func hurdleNine(urlStr string) {
	if strings.Contains(urlStr, "$post_k=") {
		splitStr := strings.Split(urlStr, "$post_k=")
		segmentfaultHTML, err := post(splitStr[0], map[string]string{"k": splitStr[1]}, map[string]string{"Referer": splitStr[0]})
		if err != nil {
			fmt.Println(err)
		}

		re := regexp.MustCompile(`<pre>([\s\S]+?)</pre>`)
		find := re.FindStringSubmatch(segmentfaultHTML)
		binaryData := strings.Replace(find[1], "____", "1111", -1)
		fileData, err := base64.StdEncoding.DecodeString(string(BinaryStringToBytes(binaryData)))
		if err != nil {
			fmt.Println(err)
		}
		err = ioutil.WriteFile("./segmentfault.tar.gz", fileData, 0666)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("[=>]通往第 10 关的文件已生成，请自行解压通关！")
			fmt.Print("\n[**]你也可以复制下面链接直接通行第 10 关：\n\nhttps://1111.segmentfault.com/?k=e4a4a96a69a1b2b530b3bec6734cdf52\n")
		}
	} else {
		fmt.Println("第 9 关不能直接访问，选从第 8 关开始获取通往第 9 关的链接")
	}
}

func main() {
	segmentfaultURL := flag.String("sf", "https://1111.segmentfault.com/?k=1573402aa6086d9ce42cfd5991027022", "Use -sf <SegmentFault 1111 URL>")
	flag.Parse()
	fmt.Printf("\nSegmentFault 1111 URL: %s\n", *segmentfaultURL)
	if strings.Contains(*segmentfaultURL, "$post_k=") {
		fmt.Println("\n[=>]你从第 9 关开始的 =>")
		hurdle["9"](*segmentfaultURL)
	} else {
		segmentfaultHTML, err := get(*segmentfaultURL)
		if err != nil {
			fmt.Println(err)
		}

		re := regexp.MustCompile(`<title>[\S]+?(\d|,\s)+[\S]+</title>`)
		find := re.FindStringSubmatch(segmentfaultHTML)
		if find != nil && find[1] != "" && strings.Contains(find[1], ",") != true {
			fmt.Println("\n[=>]你从第", find[1], "关开始的 =>")
			hurdle[find[1]](*segmentfaultURL)
		} else if *segmentfaultURL == "https://1111.segmentfault.com/?k=e4a4a96a69a1b2b530b3bec6734cdf52" {
			fmt.Println("\n恭喜, 你已经通过了所有关卡！")
		} else {
			fmt.Println("\n对不起, 你提供的链接不能识别，请核对后再试！")
		}
	}
}
