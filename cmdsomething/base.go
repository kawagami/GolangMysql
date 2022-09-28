package cmdsomething

import (
	"fmt"
	"os/exec"
	"time"

	cb "github.com/atotto/clipboard"
)

// Open frequently visited sites
func Ofvs() {
	fmt.Println("")
	start := time.Now().UnixMicro()
	//
	cmd := "cmd"
	web1 := []string{"/c", "start", `http://www.wnacg.org/albums.html`}
	web2 := []string{"/c", "start", `https://e-hentai.org/?f_search=language%3Achinese%24+-yaoi`}
	web3 := []string{"/c", "start", `https://nhentai.net/search/?q=-yaoi+chinese`}
	exec.Command(cmd, web1...).Start()
	exec.Command(cmd, web2...).Start()
	exec.Command(cmd, web3...).Start()
	//
	end := time.Now().UnixMicro()
	timeResult := end - start
	processTime := float64(timeResult) / 1000000
	fmt.Printf("經過 %v 秒\n", processTime)
	fmt.Println("")

}

func WebInfoWithCopy() {
	text, _ := cb.ReadAll()
	cmd := "cmd"
	web2String := fmt.Sprintf("https://e-hentai.org/?f_search=%s", text)
	web3String := fmt.Sprintf("https://nhentai.net/search/?q=%s", text)
	web1String := fmt.Sprintf("http://www.wnacg.org/search/?q=%s&f=_all&s=create_time_DESC&syn=yes", text)
	web1 := []string{"/c", "start", web1String}
	web2 := []string{"/c", "start", web2String}
	web3 := []string{"/c", "start", web3String}
	exec.Command(cmd, web1...).Start()
	exec.Command(cmd, web2...).Start()
	exec.Command(cmd, web3...).Start()
}

func VideoInfoWithCopy() {
	text, _ := cb.ReadAll()
	cmd := "cmd"
	web1String := fmt.Sprintf("https://javdb.com/search?q=%s&f=all", text)
	web1 := []string{"/c", "start", web1String}
	exec.Command(cmd, web1...).Start()
}
