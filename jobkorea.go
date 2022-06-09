package main

import (
	"bufio"
	"delivery/src/controller/cls"
	"fmt"
	"net/url"
	"strings"
	"time"
)

/*
회사 리스트 페이지 호출 후 응답 html
a 태그 /Recruit/Co_Read/C/companyCode
companyCode 값을 이용해서 상세 페이지 호출

* 기업 정보

키워드 -> 기업정보</h2>
	<th class="field-label"> -> key
	<div class="value"> -> value
키워드 -> 재무분석</h2>
=> 키워드 사이 key, value 값

* 복리

키워드 -> 복리후생</h3> 밑으로
-> p 태그 안에 글자
키워드 -> 기업위치</h3>
*/

func GetJobkoreaCompay() {
	rst := httpJobkoreaList()
	if rst < 0{
		return
	}
}

// 기업 리스트
func httpJobkoreaList() int{

	var httpHeader map[string]string
	httpHeader = make(map[string]string)
	httpHeader["X-Requested-With"] = "XMLHttpRequest"

	netUrl := url.Values{}
	netUrl.Add("page","1")
	netUrl.Add("pagesize","50")
	netUrl.Add("condition[local]","I070")    // 지역 코드 - 구로구
	netUrl.Add("condition[subway]","")  // 지하철 코드 - 구로디지털단지역
	netUrl.Add("condition[menucode]","")
	netUrl.Add("direct","0")

	/*
		order
		2 - 등록일순
		3 - 최신업데이트순
		4 - 마감일순
		5 - 경력순
		7 - 학력순
		20 - 추천순
	 */

	netUrl.Add("order","4")
	netUrl.Add("tabindex","0")
	netUrl.Add("fulltime","0")
	netUrl.Add("confirm","0")

	resp, err := cls.HttpRequestDetail("HTTPS", "POST", "www.jobkorea.co.kr", "443", "Recruit/Home/_GI_List/", []byte(netUrl.Encode()), httpHeader, "application/x-www-form-urlencoded", false)
	if err != nil{
		lprintf(1, "[ERROR] naver get review addr err(%s)\n", err.Error())
		return -1
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	keyword := "/Recruit/Co_Read/C"

	for scanner.Scan(){

		data := scanner.Text()

		if strings.Contains(data, keyword){
			sindex := strings.Index(data, keyword)
			index := len(keyword)+sindex
			sindex2 := strings.Index(data[index:],"\"")

			//fmt.Println("line : ", data, " code : ", data[index+1:index+sindex2])

			httpJobkoreaInfo(data[index+1:index+sindex2])

			// 10분에 한번씩
			time.Sleep(10*time.Minute)
		}
	}
/*
	keyword := "/Recruit/Co_Read/C"
	data := "<a href=\"/Recruit/Co_Read/C/mktkid\" class=\"link normalLog\" data-clickctgrcode=\"B01\" target=\"_blank\">㈜마케팅키드</a>"
	sindex := strings.Index(data, keyword)
	index := len(keyword)+sindex
	fmt.Println(data[sindex:index])

	sindex2 := strings.Index(data[index:],"\"")
	fmt.Println(data[index+1:index+sindex2])
	return
 */

	return 1
}

// 기업 상세
func httpJobkoreaInfo(companyCode string) int{

	url := fmt.Sprintf("Recruit/Co_Read/C/%s?Oem_Code=C1", companyCode)

	resp, err := cls.HttpRequestDetail("HTTPS", "GET", "www.jobkorea.co.kr", "443", url, nil, nil, "", false)
	if err != nil{
		lprintf(1, "[ERROR] naver get review addr err(%s)\n", err.Error())
		return -1
	}
	defer resp.Body.Close()

	var companyMap map[string][]string // 기업 정보
	companyMap = make(map[string][]string)

	var companyKey string
	var companySlice []string // 기업 복리

	// 기업 정보
	var start, end bool
	startKeyword := "기업정보</h2>"
	endKeyword := "재무분석</h2>"

	keyword1 := "<th class=\"field-label\">"
	keyword2 := "<div class=\"value\">"

	// 기업 복리
	var start2, end2 bool
	startKeyword2 := "복리후생</h3>"
	endKeyword2 := "기업위치</h3>"

	keyword3 := "<p>"

	var companyName string

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan(){

		// 기업 이름
		if len(companyName) == 0{
			if strings.Contains(scanner.Text(), "<title>"){
				companyName = searchText(scanner.Text(), "<title>")
			}
		}

		// 기업 정보
		if strings.Contains(scanner.Text(), startKeyword){
			start = true
		}

		if strings.Contains(scanner.Text(), endKeyword){
			end = true
		}

		if start && !end{
			if strings.Contains(scanner.Text(), keyword1){
				m,k := findText(scanner.Text(),keyword1,companyKey, 1, companyMap)
				if m != nil{
					companyMap = m
				}else{
					companyKey = k
				}
			}else if strings.Contains(scanner.Text(), keyword2){
				m,k := findText(scanner.Text(),keyword2,companyKey, 2, companyMap)
				if m != nil{
					companyMap = m
				}else{
					companyKey = k
				}
			}
		}

		// 기업 복리
		if strings.Contains(scanner.Text(), startKeyword2){
			start2 = true
		}

		if strings.Contains(scanner.Text(), endKeyword2){
			end2 = true
		}

		if start2 && !end2{
			if strings.Contains(scanner.Text(), keyword3){
				companySlice = findText2(scanner.Text(),keyword3,companySlice)
			}
		}
	}

	/*
	// 결과 테스트
	fmt.Println("기업정보")
	for k,v := range companyMap{
		fmt.Println("key : ", k)
		for _,v2 := range v{
			fmt.Println("val : ", v2)
		}
	}

	fmt.Println("기업복지")
	for _,v := range companySlice{
		fmt.Println("val : ", v)
	}
	 */

	setCompany(companyCode, companyName, "I070", companyMap, companySlice)

	return 1
}

// 기업 정보 추출
func findText(data, keyword, companyKey string, flag int, companyMap map[string][]string) (map[string][]string, string) {
	sindex := strings.Index(data,keyword)
	index := len(keyword)+sindex
	sindex2 := strings.Index(data[index:],"<")

	f := data[index:index+sindex2]
	//fmt.Println(data[index:index+sindex2])

	//if flag == 1{ // key

	//fmt.Printf("%v %s \n",companyMap,companyKey)

	if flag == 2{ //value
		val, exists := companyMap[companyKey]
		if !exists {
			companyMap[companyKey] = []string{f}
		}else{
			val = append(val, f)
			companyMap[companyKey] = val
		}

		return companyMap, ""
	}

	return nil, f
}

// 기업 복지 추출
func findText2(data, keyword string, companySlice []string) []string{

	sindex := strings.Index(data,keyword)
	index := len(keyword)+sindex
	sindex2 := strings.Index(data[index:],"<")

	f := data[index:index+sindex2]
	companySlice = append(companySlice, strings.TrimSpace(f))
	//fmt.Println(strings.TrimSpace(f))

	return companySlice
}

func searchText(data, keyword string) string{
	sindex := strings.Index(data,keyword)
	index := len(keyword)+sindex
	sindex2 := strings.Index(data[index:],"</title>")

	return data[index:index+sindex2]
}

func setCompany(companyCode, companyName, companyPlace string, companyMap map[string][]string, companySlice []string){

	query := "REPLACE INTO a_jobkorea(COMPANY_CODE, COMPANY_NAME, COMPANY_PLACE_CODE, AMT, CEO, INSURANCE, HOMEPAGE, INDUSTRY, TYPE, DATE, ADDR, PEOPLE, CAPITAL, " +
		"BUSINESS, WELFARE) " +
		"VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"

	var params []interface{}
	params = append(params, companyCode)
	params = append(params, companyName)
	params = append(params, companyPlace)

	tmp := ""
	for idx, v := range companyMap["매출액"] {
		if idx == 0 {
			tmp = fmt.Sprintf("%v", v)
		} else {
			tmp += "," + fmt.Sprintf("%v", v)
		}
	}
	params = append(params, tmp)

	tmp = ""
	for idx, v := range companyMap["대표자"] {
		if idx == 0 {
			tmp = fmt.Sprintf("%v", v)
		} else {
			tmp += "," + fmt.Sprintf("%v", v)
		}
	}
	params = append(params, tmp)

	tmp = ""
	for idx, v := range companyMap["4대보험"] {
		if idx == 0 {
			tmp = fmt.Sprintf("%v", v)
		} else {
			tmp += "," + fmt.Sprintf("%v", v)
		}
	}
	params = append(params, tmp)

	tmp = ""
	for idx, v := range companyMap["홈페이지"] {
		if idx == 0 {
			tmp = fmt.Sprintf("%v", v)
		} else {
			tmp += "," + fmt.Sprintf("%v", v)
		}
	}
	params = append(params, tmp)

	tmp = ""
	for idx, v := range companyMap["산업"] {
		if idx == 0 {
			tmp = fmt.Sprintf("%v", v)
		} else {
			tmp += "," + fmt.Sprintf("%v", v)
		}
	}
	params = append(params, tmp)

	tmp = ""
	for idx, v := range companyMap["기업구분"] {
		if idx == 0 {
			tmp = fmt.Sprintf("%v", v)
		} else {
			tmp += "," + fmt.Sprintf("%v", v)
		}
	}
	params = append(params, tmp)

	tmp = ""
	for idx, v := range companyMap["설립일"] {
		if idx == 0 {
			tmp = fmt.Sprintf("%v", v)
		} else {
			tmp += "," + fmt.Sprintf("%v", v)
		}
	}
	params = append(params, tmp)

	tmp = ""
	for idx, v := range companyMap["주소"] {
		if idx == 0 {
			tmp = fmt.Sprintf("%v", v)
		} else {
			tmp += "," + fmt.Sprintf("%v", v)
		}
	}
	params = append(params, tmp)

	tmp = ""
	for idx, v := range companyMap["사원수"] {
		if idx == 0 {
			tmp = fmt.Sprintf("%v", v)
		} else {
			tmp += "," + fmt.Sprintf("%v", v)
		}
	}
	params = append(params, tmp)

	tmp = ""
	for idx, v := range companyMap["자본금"] {
		if idx == 0 {
			tmp = fmt.Sprintf("%v", v)
		} else {
			tmp += "," + fmt.Sprintf("%v", v)
		}
	}
	params = append(params, tmp)

	tmp = ""
	for idx, v := range companyMap["주요사업"] {
		if idx == 0 {
			tmp = fmt.Sprintf("%v", v)
		} else {
			tmp += "," + fmt.Sprintf("%v", v)
		}
	}
	params = append(params, tmp)

	tmp = ""
	for idx, v := range companySlice {
		if idx == 0 {
			tmp = fmt.Sprintf("%v", v)
		} else {
			tmp += "," + fmt.Sprintf("%v", v)
		}
	}
	params = append(params, tmp)

	_, err := cls.ExecDBbyParam(query, params)
	if err != nil {
		lprintf(1, "[ERROR] cls.ExecDBbyParam error(%s) \n", err.Error())
	}

	return
}