package main

import (
	"bufio"
	"bytes"
	"delivery/src/controller/cls"
	"fmt"
	"golang.org/x/crypto/ssh"
	"image/color"
	"image/png"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	textrank "github.com/DavidBelicza/TextRank"
	"github.com/pkg/sftp"
	"github.com/psykhi/wordclouds"
)

const (
	RECENT = iota   // 최근 리뷰
	LAST_MONTH      // 지난달 리뷰
	OLD_MONTH       // 과거 4달 이전 리뷰
)
/*
func renderNow(restId string, text map[string]int) {
	//需要写入的文本数组
	//textList := []string{"恭喜", "发财", "万事", "如意"}
	var textList []string

	for k := range text{
		textList = append(textList, k)
	}

	//文本角度数组
	angles := []int{0, 15, -15, 90}
	//文本颜色数组

	colors := []*wordcloud.Color{
		// water
		{0x00,	0x00,	0x80},
		{0x57,	0xa0,	0xd2},
		{0x0f,	0x52,	0xba},
		{0x00,	0x80,	0xfe},
		{0x11,	0x34,	0xa6},
	}

	//设置对应的字体路径，和输出路径
	render := wordcloud.NewWordCloudRender(60, 8,
		"./src/fonts/gulim.ttc",
		"./src/img/water.png", textList, angles, colors, fmt.Sprintf("./src/%s/out.png", restId))
	//开始渲染
	render.Render()
}
 */

func GetCompInfo()(int, []map[string]string){

	var params map[string]string
	//params = make(map[string]string)

	// 달아요 캐시 가맹점
	compInfo, err := cls.GetSelectData2(SelectCompInfo, params)
	if err != nil {
		lprintf(1, "[ERROR] get comp err(%s)\n", err.Error())
		return -1, nil
	}

	return 1, compInfo
}

func GetStoreInfo(restId string)(int, []map[string]string){

	var params map[string]string

	var query string
	if len(restId) > 0{
		params = make(map[string]string)
		params["restId"] = restId

		query = SelectStoreInfo
	}else{
		query = SelectStoreInfoAll
	}

	store, err := cls.SelectData(query, params)
	if err != nil{
		lprintf(1, "[ERROR] get store err(%s)\n", err.Error())
		return -1, nil
	}

	return 1, store
}

func GetBaeminReivew(baeminId string)(int, []map[string]string){

	var params map[string]string
	params = make(map[string]string)
	params["baeminId"] = baeminId

	dt := time.Now().AddDate(0, 0, -7)
	startWeek := cls.GetFirstOfWeek(dt)
	endWeek := cls.GetEndOfWeek(dt)
	wStartDt := startWeek.Format("20060102")
	wEndDt := endWeek.Format("20060102")

	params["startDt"] = wStartDt
	params["endDt"] = wEndDt


	review, err := cls.SelectData(SelectBaeminReview, params)
	if err != nil{
		lprintf(1, "[ERROR] get review err(%s)\n", err.Error())
		return -1, nil
	}

	return 1, review
}

// 달아요 캐시 가맹점 - 배민, 네이버, 요기요 맵핑
func SetStore(restId, bizNum, compNm, baeminId, yogiyoId, naverId, coupangId string) int {

	lprintf(4, "[INFO] restId(%s), bizNum(%s), compNm(%s), bId(%s), yId(%s), nId(%s), cId(%s)\n",restId, bizNum, compNm, baeminId, yogiyoId, naverId, coupangId)

	rst, comps := GetStoreInfo(restId)
	if rst > 0 && len(comps) > 0{

		if len(comps[0]["baemin_id"]) > 0{
			baeminId = comps[0]["baemin_id"]
		}

		if len(comps[0]["yogiyo_id"]) > 0{
			yogiyoId = comps[0]["yogiyo_id"]
		}

		if len(comps[0]["naver_id"]) > 0{
			naverId = comps[0]["naver_id"]
		}

		if len(comps[0]["coupang_id"]) > 0{
			coupangId = comps[0]["coupang_id"]
		}

		lprintf(4, "[INFO] store info restId(%s), bizNum(%s), compNm(%s), bId(%s), yId(%s), nId(%s), cId(%s)\n",restId, bizNum, compNm, baeminId, yogiyoId, naverId, coupangId)
	}

	query := "REPLACE INTO b_store(REST_ID, BIZ_NUM, COMP_NM, BAEMIN_ID, YOGIYO_ID, NAVER_ID, COUPANG_ID) VALUES(?,?,?,?,?,?,?);"

	var params []interface{}
	params = append(params, restId)
	params = append(params, bizNum)
	params = append(params, compNm)
	params = append(params, baeminId)
	params = append(params, yogiyoId)
	params = append(params, naverId)
	params = append(params, coupangId)

	_, err := cls.ExecDBbyParam(query, params)
	if err != nil {
		lprintf(1, "[ERROR] cls.ExecDBbyParam error(%s) \n", err.Error())
		return -1
	}

	return 1
}

func ReplaceAddr(addr string) string {

	if len(addr) == 0{
		return addr
	}

	var newAddr string
	addrs := strings.Split(addr, " ")

	if strings.Contains(addrs[0], "서울") || strings.Contains(addrs[0], "서울시"){
		newAddr = "서울특별시"
	}else if strings.Contains(addrs[0], "경기"){
		newAddr = "경기도"
	}else if strings.Contains(addrs[0], "강원"){
		newAddr = "강원도"
	}else if strings.Contains(addrs[0], "인천"){
		newAddr = "인천광역시"
	}else if strings.Contains(addrs[0], "울산"){
		newAddr = "울산광역시"
	}else if strings.Contains(addrs[0], "대전"){
		newAddr = "대전광역시"
	}else if strings.Contains(addrs[0], "경남"){
		newAddr = "경상남도"
	}else if strings.Contains(addrs[0], "경북"){
		newAddr = "경상북도"
	}else if strings.Contains(addrs[0], "충북"){
		newAddr = "충청북도"
	}else if strings.Contains(addrs[0], "충남"){
		newAddr = "충청남도"
	}else if strings.Contains(addrs[0], "부산"){
		newAddr = "부산광역시"
	}else if strings.Contains(addrs[0], "전북"){
		newAddr = "전라북도"
	}else if strings.Contains(addrs[0], "전남"){
		newAddr = "전라남도"
	}else if strings.Contains(addrs[0], "광주"){
		newAddr = "광주광역시"
	}else if strings.Contains(addrs[0], "제주"){
		newAddr = "제주특별시"
	}

	if len(newAddr) > 0{
		for idx, v := range addrs{
			if idx == 0{
				continue
			}

			newAddr += " " + v
		}
		return newAddr
	}

	return addr
}

func IsReviewFinished(date string) bool{

	var td string // 20210514

	if len(date) != 8 {
		if strings.Contains(date, "."){
			td = strings.ReplaceAll(date, ".", "-")
		}else if strings.Contains(date, ":"){
			dates := strings.Split(date, " ")
			td = dates[0]
		}else{
			td = krToDate(date)
		}
	}

	now := time.Now()
	limitDate := now.AddDate(0, -4, -now.Day())

	thisDate,err := time.Parse("20060102", td)
	if err != nil{
		thisDate,err = time.Parse("2006-01-02", td)
		if err != nil{
			lprintf(1, "[ERROR] time parse err(%s)\n", err.Error())
			return true
		}
	}

	if limitDate.Sub(thisDate).Hours() >= 0{
		lprintf(4, "[INFO] review limitDate(%s), thisDate(%s)\n", limitDate.Format("20060102"), thisDate.Format("20060102"))
		return true
	}

	return false
}

// 배민, 쿠팡이츠
func krToDate(date string) string{

	n := time.Now()
	var td string

	if strings.Contains(date, "오늘"){
		td = n.Format("20060102")
	}else if strings.Contains(date, "어제") || strings.Contains(date, "1일 전"){
		td = n.AddDate(0,0,-1).Format("20060102")
	}else if strings.Contains(date, "그제") || strings.Contains(date, "2일 전"){
		td = n.AddDate(0,0,-2).Format("20060102")
	}else if strings.Contains(date, "3일 전"){
		td = n.AddDate(0,0,-3).Format("20060102")
	}else if strings.Contains(date, "4일 전"){
		td = n.AddDate(0,0,-4).Format("20060102")
	}else if strings.Contains(date, "5일 전"){
		td = n.AddDate(0,0,-5).Format("20060102")
	}else if strings.Contains(date, "6일 전"){
		td = n.AddDate(0,0,-6).Format("20060102")
	}else if strings.Contains(date, "이번 주"){
		td = n.AddDate(0,0,-3).Format("20060102")
	}else if strings.Contains(date, "지난 주"){
		td = n.AddDate(0,0,-7).Format("20060102")
	}else if strings.Contains(date, "이번 달"){
		td = n.AddDate(0,0,-8).Format("20060102")
	}else if strings.Contains(date, "지난 달"){
		td = n.AddDate(0,-1,0).Format("20060102")
	}else if strings.Contains(date, "2개월 전"){
		td = n.AddDate(0,-2,0).Format("20060102")
	}else if strings.Contains(date, "3개월 전"){
		td = n.AddDate(0,-3,0).Format("20060102")
	}else if strings.Contains(date, "4개월 전"){
		td = n.AddDate(0,-4,0).Format("20060102")
	}else if strings.Contains(date, "-") && len(date) == 10{
		date = strings.ReplaceAll(date, "-", "")
		if len(date) == 8{
			return date
		}else{
			td = n.AddDate(-1,0,0).Format("20060102")
		}
	}else{
		td = n.AddDate(-1,0,0).Format("20060102")
	}

	return td

}

func makeWordCloud(addr, port, id, pwd string){

	rst, comps := GetStoreInfo("")
	if rst < 0{
		lprintf(1, "[ERROR] get store err\n")
		return
	}

	maskName := "mask"
	dirPath := "src"
	now := time.Now()
	startDt := fmt.Sprintf("%s01",now.AddDate(0, -1, 0).Format("200601"))
	endDt := now.AddDate(0, 0, -now.Day()).Format("20060102")
	oldDt := fmt.Sprintf("%s01",now.AddDate(0, -3, 0).Format("200601"))
	var fPaths []string

	for _, comp := range comps{
		b := comp["baemin_id"]
		y := comp["yogiyo_id"]
		n := comp["naver_id"]
		restId := comp["rest_id"]

		fPath := fmt.Sprintf("%s/%s",dirPath,restId)

		if _, err := os.Stat(fPath); !os.IsExist(err){
			os.MkdirAll(fPath, os.ModePerm)
		}

		rst, wcPath := wordCloudReview(restId, b, y, n, dirPath, maskName, "", oldDt, OLD_MONTH)
		if rst > 0{
			fPaths = append(fPaths, wcPath)
		}

		time.Sleep(10*time.Second)
		runtime.GC()
		time.Sleep(20*time.Second)

		rst, wcPath = wordCloudReview(restId, b, y, n, dirPath, maskName, startDt, endDt, LAST_MONTH)
		if rst > 0{
			fPaths = append(fPaths, wcPath)
		}

		time.Sleep(10*time.Second)
		runtime.GC()
		time.Sleep(20*time.Second)
	}

	pushSftp(addr,port,id,pwd,fPaths)

	lprintf(4, "[INFO] word finish")
}

func pushSftp(addr, port, id, pwd string, fPath []string){

	// addr 106.10.42.130: 15022

	if len(fPath) == 0{
		lprintf(4, "[ERROR] fPath is size 0\n")
		return
	}

	//hostKey := getHostKey(addr)
	//if hostKey == nil{
	//	lprintf(1, "[ERROR] hostKey is nil\n")
	//	return
	//}

	config := &ssh.ClientConfig{
		User: id,
		Auth: []ssh.AuthMethod{
			ssh.Password(pwd),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		//HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", addr, port), config)
	if err != nil{
		lprintf(1, "[ERROR] connect err(%s)\n", err.Error())
		return
	}
	defer conn.Close()

	client, err := sftp.NewClient(conn)
	if err != nil{
		lprintf(1, "[ERROR] sftp client err(%s)\n", err.Error())
		return
	}
	defer client.Close()

	for _,path := range fPath{

		fInfo := strings.Split(path,"/")
		if len(fInfo) != 3{
			continue
		}

		mPath := fmt.Sprintf("/app/SharedStorage/wordCloud/%s", fInfo[1])
		cPath := fmt.Sprintf("/app/SharedStorage/wordCloud/%s/%s", fInfo[1], fInfo[2])

		err := client.MkdirAll(mPath)
		if err != nil{
			lprintf(1, "[ERROR] sftp client mkdir err(%s)\n", err.Error())
			continue
		}

		dstFile, err := client.Create(cPath)
		if err != nil{
			lprintf(1, "[ERROR] sftp client create err(%s)\n", err.Error())
			continue
		}

		srcFile, err := os.Open(path)
		if err != nil{
			lprintf(1, "[ERROR] os open err(%s)\n", err.Error())
			dstFile.Close()
			continue
		}

		_,err = io.Copy(dstFile, srcFile)
		if err != nil{
			lprintf(1, "[ERROR] io copy err(%s)\n", err.Error())
		}
		dstFile.Close()
		srcFile.Close()

	}
}

func getHostKey(host string) ssh.PublicKey {
	// parse OpenSSH known_hosts file
	// ssh or use ssh-keyscan to get initial key
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		lprintf(1, "[ERROR] op open err(%s)\n", err.Error())
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				lprintf(1, "[ERROR] ssh err(%s)\n", err.Error())
				return nil
			}
			break
		}
	}

	if hostKey == nil {
		return nil
	}

	return hostKey
}

/*
	가맹점 리뷰를 이용해서 워드 클라우드 이미지 생성

	restId - 가맹점 id
	baeminId - 가맹점 baemin id
	yogiyoId - 가맹점 yogiyo id
	naverId - 가맹점 naver id
	dstPath - 워드 클라우드 생성 path
	mask - 워드 클라우드 형태 image mask 이름
	startDt - 워드 클라우드 생성 시 리뷰 기간
	endDt - 워드 클라우드 생성 시 리뷰 기간
	reviewType - 워드 클라우드 생성 리뷰 기간 타입
 */
func wordCloudReview(restId, baeminId, yogiyoId, naverId, dirPath, mask, startDt, endDt string, reviewType int) (int, string){

	lprintf(4, "[INFO] restId(%s), baeminId(%s), yogiyoId(%s), naverId(%s) \n", restId, baeminId, yogiyoId, naverId)

	if len(baeminId) == 0 && len(yogiyoId) == 0 && len(naverId) == 0{
		//copyWordCloud(bizNum)
		lprintf(4, "[INFO] restId(%s) no delivery company\n", restId)
		return -1, ""
	}

	if len(baeminId) == 0{
		baeminId = "baeminId"
	}
	if len(yogiyoId) == 0{
		yogiyoId = "yogiyoId"
	}
	if len(naverId) == 0{
		naverId = "naverId"
	}

	var params map[string]string
	params = make(map[string]string)

	params["baeminId"] = baeminId
	params["yogiyoId"] = yogiyoId
	params["naverId"] = naverId
	params["startDt"] = startDt
	params["endDt"] = endDt

	var reviewQuery, fileName string

	if reviewType == RECENT{
		reviewQuery = SelectWordCloudReview
		fileName = fmt.Sprintf("recent_%s.png", restId)
	}else if reviewType == LAST_MONTH{
		reviewQuery = SelectWordCloudLastReview
		fileName = fmt.Sprintf("%s_%s.png", startDt[:6],restId)
	}else if reviewType == OLD_MONTH{
		reviewQuery = SelectWordCloudLastReview2
		fileName = fmt.Sprintf("old_%s.png", restId)
	}else{
		reviewQuery = SelectWordCloudLastReview2
		fileName = fmt.Sprintf("old_%s.png", restId)
	}

	reviews, err := cls.SelectData(reviewQuery, params)
	if err != nil{
		lprintf(1, "[ERROR] get SelectWordCloudReview err(%s)\n", err.Error())
		//copyWordCloud(bizNum)
		return -1, ""
	}

	if len(reviews) == 0{
		return -1, ""
	}

	// make world cloud
	// set color
	colorsRGBA := []color.RGBA{
		// water
		{0x00,	0x00,	0x80,0xff},
		{0x57,	0xa0,	0xd2,0xff},
		{0x0f,	0x52,	0xba,0xff},
		{0x00,	0x80,	0xfe,0xff},
		{0x11,	0x34,	0xa6,0xff},
	}

	colors := make([]color.Color, 0)
	for _, c := range colorsRGBA {
		colors = append(colors, c)
	}

	// review text ranking
	tr := textrank.NewTextRank()
	rule := textrank.NewDefaultRule()
	language := textrank.NewDefaultLanguage()
	algorithmDef := textrank.NewDefaultAlgorithm()
	//algorithmDef :=textrank.NewChainAlgorithm()

	for _, v := range reviews{
		tr.Populate(v["content"], language, rule)
	}

	tr.Ranking(algorithmDef)

	//rankedPhrases := textrank.FindPhrases(tr)
	words := textrank.FindSingleWords(tr)
	if len(words) == 0{
		lprintf(1, "[ERROR] find single words cnt 0 \n")
		//copyWordCloud(bizNum)
		return -1, ""
	}

	var inputWords map[string]int
	inputWords = make(map[string]int)

	// 단어 rank에 따른 word cloud에 표현되는 단어 size
	for _,v := range words{
		if v.Qty > 4{
			inputWords[strings.TrimSpace(v.Word)] = rand.Intn(30)+60
		}else if v.Qty > 2{
			inputWords[strings.TrimSpace(v.Word)] = rand.Intn(25)+35
		}else if v.Qty > 1{
			inputWords[strings.TrimSpace(v.Word)] = rand.Intn(20)+15
		}else{
			inputWords[strings.TrimSpace(v.Word)] = rand.Intn(15)+1
		}

		// word cloud text count
		if len(inputWords) > 180{
			break
		}
	}

	//renderNow(restId, inputWords)
	//return 1,0

	// size - 256, 512, 2048
	var size int
	if len(inputWords) >= 180{ // 워드 클라우드 잘 나옴
		size = 2048
	}else if len(inputWords) <= 50{ // 워드 클라우드 원형으로 작게
		size = 256
	}else{
		size = 512 // 워드 클라우드 원형으로 작게
	}

	boxes := wordclouds.Mask(
		fmt.Sprintf("%s/img/%s.png", dirPath,mask),
		size,
		size,
		color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 0,
		})

	w := wordclouds.NewWordcloud(inputWords,
		wordclouds.FontFile(fmt.Sprintf("%s/fonts/gulim.ttc",dirPath)),
		wordclouds.FontMaxSize(300),
		wordclouds.FontMinSize(30),
		wordclouds.Colors(colors),
		wordclouds.MaskBoxes(boxes),
		wordclouds.Height(2048),
		wordclouds.Width(2048),
	)

	fPath := fmt.Sprintf("%s/%s/%s",dirPath,restId,fileName)

	img := w.Draw()
	outputFile, err := os.Create(fPath)
	if err != nil{
		lprintf(1, "[ERROR] png file create err(%s) \n", err.Error())
		//copyWordCloud(bizNum)
		return -1, ""
	}

	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA

	err = png.Encode(outputFile, img)
	if err != nil{
		lprintf(1, "[ERROR] png encode(%s)\n", err.Error())
	}

	// Don't forget to close files
	err = outputFile.Close()
	if err != nil{
		lprintf(1, "[ERROR] outputFile Close(%s)\n", err.Error())
	}

	//w.Clear()
	//w = nil

	/*
	fInfo, err := os.Stat(fmt.Sprintf("%s/%s",dstPath,fileName))
	if err != nil{
		lprintf(1, "[ERROR] file stat err(%s) \n", err.Error())
		//copyWordCloud(bizNum)
		return -1
	}

	if fInfo.Size() < 250000 {
		//copyWordCloud(bizNum)
		lprintf(4, "[IFNO] lack of words.. file(%s) size(%d)\n", fileName, fInfo.Size())
		return -1
	}
	 */


	for k := range inputWords{
		delete(inputWords, k)
	}

	for _,v := range reviews{
		for k := range v{
			delete(v, k)
		}
	}

	inputWords = nil
	reviews = nil

	return 1, fPath

}

func copyWordCloud(dstPath, fileName string){

	srcFile := "src/img/sample.png"
	dstFile := fmt.Sprintf("%s/%s.png",dstPath, fileName)

	fInfo, err := os.Stat(dstFile)
	if err != nil{
		lprintf(1, "[ERROR] file stat err(%s) \n", err.Error())
		copyFile(srcFile, dstFile)
		return
	}

	if fInfo.Size() < 250000 {
		copyFile(srcFile, dstFile)
	}

}

func copyFile(srcFile, dstFile string) int{
	src, err := os.Open(srcFile)
	if err != nil{
		lprintf(1, "[ERROR] file open err(%s) \n",err.Error())
		return -1
	}
	defer src.Close()

	dst, err := os.Create(dstFile)
	if err != nil{
		lprintf(1, "[ERROR] file create err(%s) \n", err.Error())
		return -1
	}

	_, err = io.Copy(dst, src)
	if err != nil{
		lprintf(1, "[ERROR] file copy err(%s) \n", err.Error())
		return -1
	}

	return 1
}

func sendChannel(title, msg, roomNumber string) {
	body := `{ "conversation_id": ${ROOM}, "text": "아프니까 사장이다",	"blocks": [ { "type": "header",	"text": "${TITLE}", "style": "blue" },  { "type": "text", "text": "${MSG}", "markdown": true } ] }`
	body = strings.Replace(body, "${TITLE}", title, -1)
	body = strings.Replace(body, "${MSG}", msg, -1)
	body = strings.Replace(body, "${ROOM}", roomNumber, -1)

	urlStr := "https://api.kakaowork.com/v1/messages.send?Content-Type=application/json"
	lprintf(4, "[INFO][go] url str(%s) \n", urlStr)
	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer([]byte(body)))
	if err != nil {
		lprintf(1, "[ERROR] http NewRequest (%s) \n", err.Error())
		return
	}

	req.Header.Set("Authorization", "Bearer 177f6c7f.dfa16ed40fd1493782f308ac9d15ce25")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		lprintf(1, "[ERROR] do error: http (%s) \n", err)
		return
	}
	defer resp.Body.Close()

	return
}