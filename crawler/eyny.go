package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
)

func GetEyny() (result []string) {
	c := colly.NewCollector() // 在colly中使用 Collector 這類物件 來做事情

	c.OnResponse(func(r *colly.Response) { // 當Visit訪問網頁後，網頁響應(Response)時候執行的事情
		fmt.Println(string(r.Body)) // 返回的Response物件r.Body 是[]Byte格式，要再轉成字串
	})

	// c.OnHTML("a[href].box", func(e *colly.HTMLElement) {
	// 	if link := e.Attr("href"); link != "" {
	// 		fmt.Println(link)
	// 	}
	// })

	// c.OnRequest(func(r *colly.Request) {
	//     r.Headers.Set("Cookie","BIGipServerPools_Web_ssl=2135533760.47873.0000; Hm_lvt_c01558ab05fd344e898880e9fc1b65c4=1577432018; qimo_seosource_578c8dc0-6fab-11e8-ab7a-fda8d0606763=%E7%BB%94%E6%AC%8F%E5%94%B4; qimo_seokeywords_578c8dc0-6fab-11e8-ab7a-fda8d0606763=; accessId=578c8dc0-6fab-11e8-ab7a-fda8d0606763; pageViewNum=3; Hm_lpvt_c01558ab05fd344e898880e9fc1b65c4=1577432866")
	//     r.Headers.Add("referer", "https://www.quanjing.com/search.aspx?q=%E5%8D%A1%E9%80%9A")
	//     r.Headers.Add("sec-fetch-mode", "cors")
	//     r.Headers.Add("sec-fetch-site", "same-origin")
	//     r.Headers.Add("accept", "text/javascript, application/javascript, application/ecmascript, application/x-ecmascript, */*; q=0.01")
	//     r.Headers.Add("accept-encoding", "gzip, deflate, br")
	//     r.Headers.Add("accept-language", "en,zh-CN;q=0.9,zh;q=0.8")
	//     r.Headers.Add("X-Requested-With", "XMLHttpRequest")
	// })

	c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
		r.Headers.Set("Cookie", "632e55XbD_e8d7_smile=1D1; 632e55XbD_e8d7_secqaaSp0e2i00=3b21aHfN7YQasu1AX9yX45jzH5mkSUaAWBXgq92eudwPdmf6ItWWjMwY49pfP3B3Gj2cif8q98Lqr1noIWfofTgYVMSzB56CpDzeLg%2FaAc6rRLNyLtys%2B8kl; 632e55XbD_e8d7_home_diymode=1; 632e55XbD_e8d7_secqaaS71WJUz0=e4c8yNUgv07inm8O4T1QoHrHM%2BcwtT2nu%2FIvJYJ8uP%2BNeDk%2FXqwCdxZ87sBZqrY0xwZxqQ1efIsXRIlNeFtqtwWUqcYMDdXhR5HeS3hZGqxWreYLpKmw04CD; 632e55XbD_e8d7_secqaaSC25mHu0=0ee21Xqerv2phFyL1iRCLwJIkrJVPMwu43oAp%2FVa4sCGY7OZRHVTSGV8wsQkQUy4TTp3ZkWWOJdj7ifZx%2FFhdHitb94YzNXzXG76Gm7Si7firmNN4%2F3Q%2BlO9; 632e55XbD_e8d7_ulastactivity=ad08iRE%2FVA5Ri%2BbQ05S3t%2FoShLP1anh19U5MjcvUSmc3wqhaVA55; 632e55XbD_e8d7_secqaaSFPHX030=9624v5Uxl0Gl00PCxYY1nGgJdLIWmXrlP3Act6WHkxKs7ZhYFb3uMEHZ%2B4%2FzHv8gHzJwZ8miuQmbpmnVR6a9Zl7DKK%2BY1VPVfF9tB8SgNOhr%2F8clLZjdeGEI; 633e55XbD_e8d7_home_diymode=1; 633e55XbD_e8d7_smile=1D1; 633e55XbD_e8d7_secqaaS79E0O70=80a7l%2BqeGLtltOxVPZSaDLDG0qoyxXkU4F%2Flki4d1%2FQqMrlOL2s54MkYnU9i2YvWYE0lE7TrN5Oj5Hh1H1JtiL3WWfx3gNLL9NEUvk9MRPim%2FC3QZICpQx8%2B; 633e55XbD_e8d7_secqaaSdF00dq0=f0dcu%2BLQaH1eH7%2B79KWK04B2GUFGKmwSwmdXWs%2FiZDwWGc%2BVK0ExiR2GiAlWlr3NOImDdk1pwGDwELgykx1dvRMJY132o3jeb%2BLk1%2FcXeJy4b13BUBivyTwS; 633e55XbD_e8d7_secqaaSpXgrz80=3b42KvFdWAKUeJM%2FvxXOa272Ps%2BDteRaS0JJWxXz6dwGXEgDWNneB9gPBex4yWCvvChuEc9Sv6%2BPzV7eNUgm8rVS%2Bmoam2yFQa8nJy%2F0mb%2F2fR%2FIQf7PT4kt; 633e55XbD_e8d7_secqaaSu9sWtz0=3a14sJKYkDVUa5eXCL9N4nEubtsMAfXZJHsWvoPoQgi2aN3FTSRidRGVcU4CWOzEv5ZLQyVDModfVQN76zryqTM3po7tY1DfX%2FHlvNDAz3Z0BNQEAOfZJsDc; 633e55XbD_e8d7_secqaaSxWfYyQ0=d0a8vNG0J1kRXKwh32ul23fnWQv2zewHQRSUddk0Msdi%2FCWWgZ4bId3Bt8F%2FBxyB7zEJuh4MHrsDi6OqlZHkzy%2BAyue8RDZ%2BZBXoiDz3VMHIllmucjqITkZ3; 633e55XbD_e8d7_ulastactivity=ec2adisZrQrft7tYr20ssipvY8Io2Dz4L7PJS5Iq0YDaz7V04Pl5; 633e55XbD_e8d7_secqaaSbd09590=e50adT6BLA%2BRzKpo0NJtnmRE%2Fq6pvN9Ve4v1MeK3h1wxdOsgqdlXlotRsgnfchgfF86yeJi8vXzzp1TwA%2FZG%2FXqVwE%2B9YN%2FQfrQQJzFIlvQNScEhsacrfVWR; 634e55XbD_e8d7_home_diymode=1; 634e55XbD_e8d7_smile=1D1; 634e55XbD_e8d7_secqaaSNXV4T00=d271bVhz1xbuznHnVtpTYs15eX%2FNjVHSApbgP52%2B2cuL8CHg%2BTWF96dYGH8VNtBqTO5EhdQh0Mztt6HerzV1N2pePAJNwfki5CiLAZ4SBoXjll%2BniGP9eAgG; 634e55XbD_e8d7_ulastactivity=6febWY2i3yksVBlHz3UqkjLyp%2FN8QyNk6mT02lmMN5oH9kBz1wpY; 634e55XbD_e8d7_secqaaSK7RNiS0=927cSaALgGNnvyW7TdgCvNiIKreYOqI%2BY6S5VzY%2B74K3HRHaLt6J6y0CjC5syuaxdmacUEnvipWyz2Zi6IZbNQ%2Bg%2Fwhkf6ja3YatxPR%2FNcwPaLzFX%2FNpoqn8; 635e55XbD_e8d7_home_diymode=1; 635e55XbD_e8d7_ulastactivity=d6b2hLUCUXzPETdql9jiWxSVNemPpLNjzImbuFqxE%2Fd3yl8EeLwL; 635e55XbD_e8d7_smile=1D1; 635e55XbD_e8d7_secqaaScs0d2M0=0824aNZg7DhNHemaY4LcDHc3F9Y3q%2Fqt5XurM7Bq5Z87oNpoX%2Fne2kZhYPdZIUBG9qgkyajBu3qJJs9%2BQY2daUXt6rRnDnZa%2BJT21GScK0WSzttmMvn%2FH5Yz; 636e55XbD_e8d7_auth=565ekb9HE%2FsZ2CbxdBL70GHWCsJTI4MAXFAab7RhpzSYw6sbV15Xwv8dWA; 636e55XbD_e8d7_lastvisit=1649554727; 636e55XbD_e8d7_home_diymode=1; 636e55XbD_e8d7_smile=1D1; 636e55XbD_e8d7_secqaaS1fzb1c0=c4a1SRQutbGLcShQPj2NUAdwhBr5FzDXjcj0Ae1i3o0WXc4Na3PNUsoVTMdUCbysTSaJBEWdjjvMYazLwxtZ0Q2L32Cm9epnWseL9EaKr%2F0h8EtF1IPpz1Jb; 636e55XbD_e8d7_favorite=a%3A1%3A%7Bi%3A4635%3Bs%3A12%3A%22MG%E8%BC%89%E9%BB%9E+FHD%22%3B%7D; 636e55XbD_e8d7_ulastactivity=8aa4XKSvWdk6mlqI3xQiu5Dg%2Buy2hzSV07YHewc9ftnDzR5XATxE; 636e55XbD_e8d7_visitedfid=2D4635; 636e55XbD_e8d7_secqaaSw5Q6N50=b7e3gm8jwDb%2Bt0dXwS%2BtUjawPyqPXljSHgZ5SSfkCxC0AL%2FoEyFcofcxWHYGQ8VhtWZEwVcAdLfiA45jF0F8Cfw4SqGEgIrP930SsKTaPl48JHyx2vbgxTLB; username=playerkiller77; 637e55XbD_e8d7_auth=4310q0lYE4hTAMGTJbsQSfjtHjLiNQcDdWmKHEs%2BumsAFzn37A7lQwSkhg; 637e55XbD_e8d7_lastvisit=1651368143; 637e55XbD_e8d7_home_diymode=1; 637e55XbD_e8d7_smile=1D1; 637e55XbD_e8d7_secqaaS55zeEb0=5f0fyOT4kVpmg039WXzm7vceodXIVwESeXv4uaTpKQrE1rTlurbeLM6gzW1Jw4lUEK4lOlIXb3mbGxmQKnWYVApwYEaFhYpDk%2Bm8IdH3pEhuRgzPm%2FsB4md%2B; 637e55XbD_e8d7_secqaaSmas11t0=00bf3OD5OQ0LRK%2FQIWD59FaND5HW37CO5GSMOcMMI3FAMTsqmZ1P5qJFlW4DfhGkNSWtnAhSIdEYeoZc0nMaOhFMs6gY1Rt5dFdVQgVTgRoxcsyhwgGSMT%2Fk; 637e55XbD_e8d7_secqaaSlNBvRZ0=0315uI2ifWNhvvnE4PzSvdyN%2FgpcEeZcok%2FCc6GVVjfQD%2BHsyVMI0J3ftiLrkP%2BYXqY9lum9HwJ%2BkIvDJ22fkjI3G8O6xTNHEWqJfXb5aC6GNVsWkm%2FHN3uX; 637e55XbD_e8d7_secqaaSKwfX9A0=6691ArXguXx6SIkAkMk%2BnaPtYY5VtvpoBX6RIVwGa43P%2Bf9gA1rRBF07UiK4Td9wqVuq%2FNzQV8BKq6ewWDD9CXVYG6u2d1FkMg%2FRAprtigDA915a44tCkPuL; 637e55XbD_e8d7_favorite=a%3A2%3A%7Bi%3A431%3Bs%3A31%3A%22H+%E5%8B%95%E7%95%AB%E4%B8%8B%E8%BC%89%E5%8D%80%28%E4%B8%8A%E5%82%B3%E7%A9%BA%E9%96%93%29%22%3Bi%3A4635%3Bs%3A12%3A%22MG%E8%BC%89%E9%BB%9E+FHD%22%3B%7D; 637e55XbD_e8d7_secqaaSJLU9zj0=2c56UFF1zsSVq68FgZ%2BzWzmsaqDGkbCknRJxH%2F1kt48474xNaPvjenNnxR%2By6BtfntD4JPqdMHSm63hvEoops%2Bw%2FlCkMdEAhmZ54AdMKXM1bQgBPWegGI8DO; 637e55XbD_e8d7_forum_lastvisit=D_431_1651478178D_4635_1651624653; 637e55XbD_e8d7_visitedfid=2D4635D431; 637e55XbD_e8d7_secqaaS2F0Jhj0=0134RCenMLxri5xxOLcvJhF0iBgXK21wDCrcwF9aSREPUxdmCk%2BhHwDDNYxXunRfsZuK2eajKczILGSGLUkEAZhpHQ1dmSNTuMDrNHZFloe4DE5K8qs55zV5; 637e55XbD_e8d7_secqaaS093xe50=0dc5GLQb%2FrU%2FzriORmzuCmg2Riuv%2BmLwMXxzXY%2BgW3QtawdmLCISSyVCuSP2ChCPKq39jURdsduYNqEC2iZkCGPSgip%2BXbjqoDb7j8rr8L8LAArscC7kebMU; 637e55XbD_e8d7_inlang=zh; 637e55XbD_e8d7_txlang=0; 637e55XbD_e8d7_ulastactivity=f162BCYQB2Zz9A%2FXI0Wac17HdrxC%2BEy7pGMtqGb2eM5rx%2F9g0%2F%2FV; 637e55XbD_e8d7_guestsid=pefSAA; 637e55XbD_e8d7_videoadult=1; 637e55XbD_e8d7_secqaaSX0XMg30=99d6lpkd1BluMbHN7tNT9Rbpp3FQGhg2MHPOTUlBFpf7I7Idd7H1tbf5GTcB6qlCFQ2rC1DgflZfczAdtvK0NcxHJdJPLHqBjGBpTM0ajRvyumozWG6B5%2Fo2; 637e55XbD_e8d7_agree=576; 637e55XbD_e8d7_secqaaSEe6bf10=b19ep72CnN8od7DaqWP3WbZcEvPq67xS4ecwGGFoZ%2FtseiOpnfpcqySU1xKLoYg9s2Mydj9BaKbCQVv6ZOl2wttrk5lpCqfGusg%2B0RY6RiTpbq47BfyFF9bH; 637e55XbD_e8d7_secqaaSTHrYbx0=199fv1iCPspAIift4EGjxSS7%2BCRLxz4PRL5r65EAzhYoq7lFhdGGdndfbUquHU9juaj4YpY%2B0DYgWV0LVlZ181Wl1xlxuP28C728FVv0ZNtgqqP74Y9kInid; 637e55XbD_e8d7_sid=44d6zd; 637e55XbD_e8d7_lastact=1651917817%09home.php%09misc")
		r.Headers.Add("sec-fetch-mode", "cors")
		r.Headers.Add("sec-fetch-site", "same-origin")
		r.Headers.Add("accept", "text/javascript, application/javascript, application/ecmascript, application/x-ecmascript, */*; q=0.01")
		r.Headers.Add("accept-encoding", "gzip, deflate, br")
		r.Headers.Add("accept-language", "en,zh-TW;q=0.9,zh;q=0.8")
		r.Headers.Add("X-Requested-With", "XMLHttpRequest")
	})

	// url := fmt.Sprintf("%v/search?q=%v&f=all", baseUrl, videoName)
	// url := `http://www.eyny.com/home.php?mod=space&uid=4706903&do=thread&view=me&from=space`
	url := `http://www.eyny.com/thread-13404675-1-1.html`
	// c.Visit(url) // Visit 要放最後

	data := map[string]string{
		"agree":  "yes",
		"submit": "是，我已年滿18歲。Yes, I am.",
	}
	// c.Visit(url)
	c.Post(url, data)

	return
}
