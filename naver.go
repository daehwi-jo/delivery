package main

import (
	"delivery/src/controller/cls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"regexp"
	"strings"
	"time"

	textrank "github.com/DavidBelicza/TextRank"
)

const(
	ASSA int = iota
	INSA
)

// naver cafe 아싸 인기글 리스트
type NaverASsaList struct {
	Articles struct {
		Items []struct {
			ID           int    `json:"id"`
			Subject      string `json:"subject"`
			Writerid     string `json:"writerId"`
			Writernick   string `json:"writerNick"`
			Writedate    int64  `json:"writeDate"`
			Commentcount int    `json:"commentCount"`
			Upcount      int    `json:"upCount"`
			Image        struct {
				URL        string `json:"url"`
				Service    string `json:"service"`
				Type       string `json:"type"`
				Isanimated bool   `json:"isAnimated"`
			} `json:"image,omitempty"`
			Salestatus          string `json:"saleStatus"`
			Isattachedmap       bool   `json:"isAttachedMap"`
			Isattachedmovie     bool   `json:"isAttachedMovie"`
			Isattachedlink      bool   `json:"isAttachedLink"`
			Isattachedmusic     bool   `json:"isAttachedMusic"`
			Isattachedcalendar  bool   `json:"isAttachedCalendar"`
			Isattachedpoll      bool   `json:"isAttachedPoll"`
			Isattachedfile      bool   `json:"isAttachedFile"`
			Isattachedimage     bool   `json:"isAttachedImage"`
			Isnewarticle        bool   `json:"isNewArticle"`
			Iscafebook          bool   `json:"isCafeBook"`
			Isbadmenubyrestrict bool   `json:"isBadMenuByRestrict"`
		} `json:"items"`
		Totalpages int `json:"totalPages"`
	} `json:"articles"`
	Isshow bool `json:"isShow"`
}

// 가맹점 이름으로 아이디 얻기
type NaverCompsWithName struct {
	Meta struct {
		Model     string `json:"model"`
		Query     string `json:"query"`
		Requestid string `json:"requestId"`
	} `json:"meta"`
	Ac    []string `json:"ac"`
	Place []struct {
		Type         string  `json:"type"`
		ID           string  `json:"id"`
		Title        string  `json:"title"`
		X            string  `json:"x"`
		Y            string  `json:"y"`
		Dist         float64 `json:"dist"`
		Totalscore   float64 `json:"totalScore"`
		Sid          string  `json:"sid"`
		Ctg          string  `json:"ctg"`
		Cid          string  `json:"cid"`
		Jibunaddress string  `json:"jibunAddress"`
		Roadaddress  string  `json:"roadAddress"`
		Review       struct {
			Count string `json:"count"`
		} `json:"review"`
	} `json:"place"`
	Address []interface{} `json:"address"`
	Bus     []interface{} `json:"bus"`
	Menu    []interface{} `json:"menu"`
	All     []struct {
		Place struct {
			Type         string  `json:"type"`
			ID           string  `json:"id"`
			Title        string  `json:"title"`
			X            string  `json:"x"`
			Y            string  `json:"y"`
			Dist         float64 `json:"dist"`
			Totalscore   float64 `json:"totalScore"`
			Sid          string  `json:"sid"`
			Ctg          string  `json:"ctg"`
			Cid          string  `json:"cid"`
			Jibunaddress string  `json:"jibunAddress"`
			Roadaddress  string  `json:"roadAddress"`
			Review       struct {
				Count string `json:"count"`
			} `json:"review"`
		} `json:"place"`
		Address interface{} `json:"address"`
		Bus     interface{} `json:"bus"`
	} `json:"all"`
}

// 가맹점 주소로 아이디 얻기
type NaverCompsWithAddr struct {
	Place struct {
		Count      int `json:"count"`
		Totalcount int `json:"totalCount"`
		List       []struct {
			Index           string        `json:"index"`
			Rank            string        `json:"rank"`
			ID              string        `json:"id"`
			Name            string        `json:"name"`
			Tel             string        `json:"tel"`
			Iscalllink      bool          `json:"isCallLink"`
			Virtualtel      string        `json:"virtualTel"`
			Ppc             string        `json:"ppc"`
			Category        []string      `json:"category"`
			Categorypath    [][]string    `json:"categoryPath"`
			Address         string        `json:"address"`
			Roadaddress     string        `json:"roadAddress"`
			Abbraddress     string        `json:"abbrAddress"`
			Shortaddress    []string      `json:"shortAddress"`
			Display         string        `json:"display"`
			Teldisplay      string        `json:"telDisplay"`
			Context         []interface{} `json:"context"`
			Reviewcount     int           `json:"reviewCount"`
			Ktcallmd        string        `json:"ktCallMd"`
			Coupon          string        `json:"coupon"`
			Thumurl         string        `json:"thumUrl"`
			Type            string        `json:"type"`
			Issite          string        `json:"isSite"`
			Posexact        string        `json:"posExact"`
			X               string        `json:"x"`
			Y               string        `json:"y"`
			Itemlevel       string        `json:"itemLevel"`
			Isadultbusiness bool          `json:"isAdultBusiness"`
			Streetpanorama  struct {
				ID   string `json:"id"`
				Pan  string `json:"pan"`
				Tilt string `json:"tilt"`
				Lng  string `json:"lng"`
				Lat  string `json:"lat"`
				Fov  string `json:"fov"`
			} `json:"streetPanorama"`
			Skypanorama struct {
				ID   string `json:"id"`
				Pan  string `json:"pan"`
				Tilt string `json:"tilt"`
				Lng  string `json:"lng"`
				Lat  string `json:"lat"`
				Fov  string `json:"fov"`
			} `json:"skyPanorama"`
			Insidepanorama   interface{} `json:"insidePanorama"`
			Interiorpanorama interface{} `json:"interiorPanorama"`
			Indoorpanorama   interface{} `json:"indoorPanorama"`
			Theme            interface{} `json:"theme"`
			Poiinfo          struct {
				Relation    interface{} `json:"relation"`
				Hasrelation bool        `json:"hasRelation"`
				Road        struct {
					Poishapetype string      `json:"poiShapeType"`
					Shapekey     interface{} `json:"shapeKey"`
					Boundary     interface{} `json:"boundary"`
					Detail       interface{} `json:"detail"`
				} `json:"road"`
				Hasroad    bool        `json:"hasRoad"`
				Land       interface{} `json:"land"`
				Hasland    bool        `json:"hasLand"`
				Polygon    interface{} `json:"polygon"`
				Haspolygon bool        `json:"hasPolygon"`
			} `json:"poiInfo"`
			Homepage       string `json:"homePage"`
			Description    string `json:"description"`
			Entrancecoords struct {
				Car []struct {
					Rep bool    `json:"rep"`
					X   float64 `json:"x"`
					Y   float64 `json:"y"`
				} `json:"car"`
				Walk []struct {
					Rep bool    `json:"rep"`
					X   float64 `json:"x"`
					Y   float64 `json:"y"`
				} `json:"walk"`
			} `json:"entranceCoords"`
			Ispollingplace     bool        `json:"isPollingPlace"`
			Bizhourinfo        string      `json:"bizhourInfo"`
			Menuinfo           string      `json:"menuInfo"`
			Petrolinfo         interface{} `json:"petrolInfo"`
			Couponurl          interface{} `json:"couponUrl"`
			Couponurlmobile    interface{} `json:"couponUrlMobile"`
			Hascardbenefit     bool        `json:"hasCardBenefit"`
			Menuexist          string      `json:"menuExist"`
			Hasnaverbooking    bool        `json:"hasNaverBooking"`
			Naverbookingurl    string      `json:"naverBookingUrl"`
			Navereasyorderurl  interface{} `json:"naverEasyOrderUrl"`
			Hasnaversmartorder bool        `json:"hasNaverSmartOrder"`
			Reservationlabel   struct {
				Standard bool `json:"standard"`
				Preorder bool `json:"preOrder"`
				Table    bool `json:"table"`
				Takeout  bool `json:"takeout"`
			} `json:"reservationLabel"`
			Reservation struct {
				Benefit string `json:"benefit"`
			} `json:"reservation"`
			Hasbroadcastinfo bool          `json:"hasBroadcastInfo"`
			Broadcastinfo    interface{}   `json:"broadcastInfo"`
			Shopwindowinfo   interface{}   `json:"shopWindowInfo"`
			Hasnpay          bool          `json:"hasNPay"`
			Carwash          string        `json:"carWash"`
			Parkingprice     string        `json:"parkingPrice"`
			Card             string        `json:"card"`
			Distance         string        `json:"distance"`
			Marker           string        `json:"marker"`
			Markerselected   string        `json:"markerSelected"`
			Microreview      []interface{} `json:"microReview"`
			Michelinguide    interface{}   `json:"michelinGuide"`
			Indoor           interface{}   `json:"indoor"`
			Markerlabel      interface{}   `json:"markerLabel"`
			Subway           interface{}   `json:"subway"`
			Evchargerinfo    struct {
				Summary struct {
					Fastevcharger struct {
						Type       string        `json:"type"`
						Typecode   string        `json:"typeCode"`
						Status     string        `json:"status"`
						Statuscode string        `json:"statusCode"`
						Names      []interface{} `json:"names"`
						Namecodes  []interface{} `json:"nameCodes"`
					} `json:"fastEvCharger"`
					Standardevcharger struct {
						Type       string        `json:"type"`
						Typecode   string        `json:"typeCode"`
						Status     string        `json:"status"`
						Statuscode string        `json:"statusCode"`
						Names      []interface{} `json:"names"`
						Namecodes  []interface{} `json:"nameCodes"`
					} `json:"standardEvCharger"`
				} `json:"summary"`
				Evchargerlist []interface{} `json:"evChargerList"`
			} `json:"evChargerInfo"`
		} `json:"list"`
	} `json:"place"`
}

// 가맹점 리뷰 받기
type NaverReviews struct {
	Data struct {
		Visitorreviews struct {
			Items []struct {
				ID     string `json:"id"`
				Rating float64    `json:"rating"`
				Author struct {
					ID       string      `json:"id"`
					Nickname string      `json:"nickname"`
					From     string      `json:"from"`
					Imageurl interface{} `json:"imageUrl"`
					Objectid string      `json:"objectId"`
					URL      string      `json:"url"`
					Review   struct {
						Totalcount int    `json:"totalCount"`
						Imagecount int    `json:"imageCount"`
						Avgrating  float64    `json:"avgRating"`
						Typename   string `json:"__typename"`
					} `json:"review"`
					Typename string `json:"__typename"`
				} `json:"author"`
				Body       string        `json:"body"`
				Thumbnail  interface{}   `json:"thumbnail"`
				Media      []interface{} `json:"media"`
				Tags       []string      `json:"tags"`
				Status     string        `json:"status"`
				Visitcount int           `json:"visitCount"`
				Viewcount  int           `json:"viewCount"`
				Visited    string        `json:"visited"`
				Created    string        `json:"created"`
				Reply      struct {
					Editurl    string      `json:"editUrl"`
					Body       string `json:"body"`
					Editedby   interface{} `json:"editedBy"`
					Created    string      `json:"created"`
					Replytitle string      `json:"replyTitle"`
					Typename   string      `json:"__typename"`
				} `json:"reply"`
				Origintype             string      `json:"originType"`
				Item                   interface{} `json:"item"`
				Language               string      `json:"language"`
				Highlightoffsets       interface{} `json:"highlightOffsets"`
				Translatedtext         interface{} `json:"translatedText"`
				Businessname           string      `json:"businessName"`
				Showbookingitemname    bool        `json:"showBookingItemName"`
				Showbookingitemoptions interface{} `json:"showBookingItemOptions"`
				Bookingitemname        interface{} `json:"bookingItemName"`
				Bookingitemoptions     string      `json:"bookingItemOptions"`
				Typename               string      `json:"__typename"`
			} `json:"items"`
			Stardistribution     []interface{} `json:"starDistribution"`
			Hideproductselectbox bool          `json:"hideProductSelectBox"`
			Total                int           `json:"total"`
			Typename             string        `json:"__typename"`
		} `json:"visitorReviews"`
	} `json:"data"`
}

type NaverReviewReq struct {
	Operationname string `json:"operationName"`
	Query         string `json:"query"`
	Variables     struct {
		Input struct {
			Bookingbusinessid interface{} `json:"bookingBusinessId"`
			Businessid        string      `json:"businessId"`
			Businesstype      string      `json:"businessType"`
			Display           int         `json:"display"`
			Getauthorinfo     bool        `json:"getAuthorInfo"`
			Includecontent    bool        `json:"includeContent"`
			Isphotoused       bool        `json:"isPhotoUsed"`
			Item              string      `json:"item"`
			Page              int         `json:"page"`
			Theme             string      `json:"theme"`
		} `json:"input"`
	} `json:"variables"`
}

func GetNaverReviews(naverId string) int{

	page := 1
	totalReviewCnt := 0

	for {
		rst, n := httpNaverReviews(naverId, page)
		if rst > 0{

			if len(n.Data.Visitorreviews.Items) == 0{
				if page > 1{
					return totalReviewCnt
				}
				lprintf(4, "[INFO] review cnt 0\n")
				break
			}

			setNaverReviews(n, naverId)
			getReviewCnt := len(n.Data.Visitorreviews.Items)
			totalReviewCnt += getReviewCnt
			if n.Data.Visitorreviews.Total == totalReviewCnt || IsReviewFinished(n.Data.Visitorreviews.Items[getReviewCnt-1].Created){
				lprintf(4, "[INFO] review finish\n")
				return totalReviewCnt
			}
		}else{
			if page > 1{
				return totalReviewCnt
			}
			break
		}
		page++
	}

	return -1
}

func setNaverCafe(n NaverASsaList, cafeType int){

	var query string

	if cafeType == ASSA{
		query = "INSERT INTO a_naver_assa(SUBJECT, COMMENT_COUNT, UP_COUNT, IMAGE_URL, WRITE_DATE) VALUES(?,?,?,?,?);"
	}else if cafeType == INSA{
		query = "INSERT INTO a_naver_insa(SUBJECT, COMMENT_COUNT, UP_COUNT, IMAGE_URL, WRITE_DATE) VALUES(?,?,?,?,?);"
	}else{
		return
	}

	for _,v := range n.Articles.Items{
		var params []interface{}
		params = append(params, v.Subject)
		params = append(params, v.Commentcount)
		params = append(params, v.Upcount)
		params = append(params, v.Image.URL)

		tm := time.Unix(0,v.Writedate*1000000)
		params = append(params, tm.Format("20060102"))

		_, err := cls.ExecDBbyParam(query, params)
		if err != nil {
			lprintf(1, "[ERROR] cls.ExecDBbyParam error(%s) \n", err.Error())
			continue
		}
	}
}

func setNaverReviews(n NaverReviews, naverId string){

	var query, tmp string

	// 이모티콘 제거
	// mysql varchar2 -> uft8 기준(3비트)
	// 이모티콘 utf16? -> 여튼 4비트라서 mysql에 안들어감
	reg,regErr := regexp.Compile("[^\u0000-\uFFFF]")
	if regErr != nil{
		lprintf(1, "[ERROR] regexp compile err(%s)\n")
	}

	for _, v := range n.Data.Visitorreviews.Items{
		query = "REPLACE INTO a_naver_review(REVIEW_ID, NAVER_ID, RATING, AUTHOR_ID, AUTHOR_NICKNAME, AUTHOR_FROM, AUTHOR_IMAGEURL, AUTHOR_OBJECT_ID, AUTHOR_URL, REVIEW_TOTAL_COUNT, REVIEW_IMAGE_COUNT, REVIEW_AVG_RATING, " +
			"REVIEW_TYPE_NAME, AUTHOR_TYPE_NAME, BODY, THUMBNAIL, MEDIA, TAGS, STATUS, VISIT_COUNT, VIEW_COUNT, VISITED, CREATED, REPLY_EDIT_URL, REPLY_BODY, REPLY_EDITED_BY, REPLY_CREATED, " +
			"REPLY_TITLE, REPLY_TYPE_NAME, ORIGIN_TYPE, LANGUAGE, HIGHLIGHT_OFFSETS, TRANSLATED_TEXT, BUSINESS_NAME, SHOW_BOOKING_ITEM_NAME, SHOW_BOOKING_ITEM_OPTIONS, BOOKING_ITEM_NAME, " +
			"BOOKING_ITEM_OPTIONS, TYPE_NAME) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"

		var params []interface{}
		params = append(params, v.ID)
		params = append(params, naverId)
		params = append(params, v.Rating)
		params = append(params, v.Author.ID)
		params = append(params, v.Author.Nickname)
		params = append(params, v.Author.From)
		params = append(params, v.Author.Imageurl)
		params = append(params, v.Author.Objectid)
		params = append(params, v.Author.URL)
		params = append(params, v.Author.Review.Totalcount)
		params = append(params, v.Author.Review.Imagecount)
		params = append(params, v.Author.Review.Avgrating)
		params = append(params, v.Author.Review.Typename)
		params = append(params, v.Author.Typename)

		if regErr == nil{
			params = append(params, reg.ReplaceAllString(v.Body, ""))
		}else{
			params = append(params, v.Body)
		}
		params = append(params, v.Thumbnail)

		for idx, v := range v.Media{
			if idx == 0{
				tmp = fmt.Sprintf("%v", v)
			}else{
				tmp += "," + fmt.Sprintf("%v", v)
			}
		}

		params = append(params, tmp)

		tmp = ""
		for idx, v := range v.Tags{
			if idx == 0{
				tmp = v
			}else{
				tmp += "," + v
			}
		}

		params = append(params, tmp)

		params = append(params, v.Status)
		params = append(params, v.Visitcount)
		params = append(params, v.Viewcount)
		params = append(params, v.Visited)
		params = append(params, v.Created)
		params = append(params, v.Reply.Editurl)

		if regErr == nil{
			params = append(params, reg.ReplaceAllString(v.Reply.Body, ""))
		}else{
			params = append(params, v.Reply.Body)
		}

		params = append(params, v.Reply.Editedby)
		params = append(params, v.Reply.Created)
		params = append(params, v.Reply.Replytitle)
		params = append(params, v.Reply.Typename)
		params = append(params, v.Origintype)
		//params = append(params, v.Item)
		params = append(params, v.Language)
		params = append(params, v.Highlightoffsets)
		params = append(params, v.Translatedtext)
		params = append(params, v.Businessname)
		params = append(params, v.Showbookingitemname)
		params = append(params, v.Showbookingitemoptions)
		params = append(params, v.Bookingitemname)
		params = append(params, v.Bookingitemoptions)
		params = append(params, v.Typename)

		_, err := cls.ExecDBbyParam(query, params)
		if err != nil {
			lprintf(1, "[ERROR] cls.ExecDBbyParam error(%s) \n", err.Error())
			continue
		}
	}
}

func CollectINsa(){

	rst, n := httpNaverCafe("10733571")
	if rst < 0{
		return
	}

	setNaverCafe(n, INSA)

	// review text ranking
	tr := textrank.NewTextRank()
	rule := textrank.NewDefaultRule()
	language := textrank.NewDefaultLanguage()
	algorithmDef := textrank.NewDefaultAlgorithm()
	//algorithmDef :=textrank.NewChainAlgorithm()

	for _, pList := range n.Articles.Items{
		tr.Populate(pList.Subject, language, rule)
	}

	tr.Ranking(algorithmDef)

	words := textrank.FindSingleWords(tr)
	if len(words) == 0{
		return
	}

	var cnt int
	var msg string
	for _,v := range words{

		if checkSendChannel(v.Word){
			msg += fmt.Sprintf("%s,", v.Word)
			cnt ++
		}

		if cnt >= 10{
			break
		}
	}

	sendChannel("인사쟁이 핫 키워드 TOP(10)", msg[:len(msg)-1], "1411731")
}

func CollectASsa(){

	rst, n := httpNaverCafe("23611966")
	if rst < 0{
		return
	}

	setNaverCafe(n, ASSA)

	// review text ranking
	tr := textrank.NewTextRank()
	rule := textrank.NewDefaultRule()
	language := textrank.NewDefaultLanguage()
	algorithmDef := textrank.NewDefaultAlgorithm()
	//algorithmDef :=textrank.NewChainAlgorithm()

	for _, pList := range n.Articles.Items{
		tr.Populate(pList.Subject, language, rule)
	}

	tr.Ranking(algorithmDef)

	words := textrank.FindSingleWords(tr)
	if len(words) == 0{
		return
	}

	var cnt int
	var msg string
	for _,v := range words{

		if checkSendChannel(v.Word){
			msg += fmt.Sprintf("%s,", v.Word)
			cnt ++
		}

		if cnt >= 10{
			break
		}
	}

	sendChannel("아프니까 핫 키워드 TOP(10)", msg[:len(msg)-1], "1411731")
}

func checkSendChannel(msg string) bool{

	delMsg1 := []string{"게", "다", "난", "는", "요", "록", "로", "어", "의", "면", "원", "할", "온", "여", "서", "죠", "에", "있", "도", "가", "한", "만", "린", "은", "려", "지", "을", "들", "운"}
	delMsg2 := []string{"님들", "러운", "부터", "않고", "ㅋㅋ", "ㄷㄷ"}
	delMsg3 := []string{"당하지", "하소연", "ㅋㅋㅋ", "ㄷㄷㄷ"}
	delMsg4 := []string{"ㅋㅋㅋㅋ", "ㄷㄷㄷㄷ"}

	for _,v := range delMsg1{
		if strings.Contains(msg, v){
			return false
		}
	}

	for _,v := range delMsg2{
		if strings.Contains(msg, v){
			return false
		}
	}

	for _,v := range delMsg3{
		if strings.Contains(msg, v){
			return false
		}
	}

	for _,v := range delMsg4{
		if strings.Contains(msg, v){
			return false
		}
	}

	return true
}

// 네이버 카페 인기글 스크랩핑
func httpNaverCafe(cafeId string)(int, NaverASsaList){

	var n NaverASsaList

	url := fmt.Sprintf("cafe-web/cafe-articleapi/cafes/%s/popular/articles?limit=45&fromAllArticleList=true&page=1", cafeId)

	resp, err := cls.HttpRequestDetail("HTTPS", "GET", "apis.naver.com", "443", url, nil, nil, "", true)
	if err != nil{
		lprintf(1, "[ERROR] naver get comp addr err(%s)\n", err.Error())
		return -1, n
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		lprintf(1, "[ERROR] naver get comp addr body read err(%s)\n", err.Error())
		return -1, n
	}

	err = json.Unmarshal(data, &n)
	if err != nil{
		lprintf(1, "[ERROR] naver get comp addr parsing json err(%s)\n", err.Error())
		return -1, n
	}

	return 1, n

}

func httpNaverReviews(naverId string, page int)(int, NaverReviews){

	var n NaverReviews
	var nq NaverReviewReq

	nq.Operationname = "getVisitorReviews"
	nq.Query = "query getVisitorReviews($input: VisitorReviewsInput) {\n  visitorReviews(input: $input) {\n    items {\n      id\n      rating\n      author {\n        id\n        nickname\n        from\n        imageUrl\n        objectId\n        url\n        review {\n          totalCount\n          imageCount\n          avgRating\n          __typename\n        }\n        __typename\n      }\n      body\n      thumbnail\n      media {\n        type\n        thumbnail\n        __typename\n      }\n      tags\n      status\n      visitCount\n      viewCount\n      visited\n      created\n      reply {\n        editUrl\n        body\n        editedBy\n        created\n        replyTitle\n        __typename\n      }\n      originType\n      item {\n        name\n        code\n        options\n        __typename\n      }\n      language\n      highlightOffsets\n      translatedText\n      businessName\n      showBookingItemName\n      showBookingItemOptions\n      bookingItemName\n      bookingItemOptions\n      __typename\n    }\n    starDistribution {\n      score\n      count\n      __typename\n    }\n    hideProductSelectBox\n    total\n    __typename\n  }\n}\n"
	nq.Variables.Input.Businessid= naverId
	nq.Variables.Input.Businesstype="restaurant"
	nq.Variables.Input.Display = 20
	nq.Variables.Input.Getauthorinfo = true
	nq.Variables.Input.Includecontent = true
	nq.Variables.Input.Isphotoused = false
	nq.Variables.Input.Item = "0"
	nq.Variables.Input.Page = page
	nq.Variables.Input.Theme = "allTypes"

	req, err := json.Marshal(nq)
	if err != nil{
		lprintf(1, "[ERROR] naver review req json err(%s)\n", err.Error())
		return -1, n
	}

	url := fmt.Sprintf("graphql")

	resp, err := cls.HttpRequestDetail("HTTPS", "POST", "pcmap-api.place.naver.com", "443", url, req, nil, "application/json", false)
	if err != nil{
		lprintf(1, "[ERROR] naver get review addr err(%s)\n", err.Error())
		return -1, n
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		lprintf(1, "[ERROR] naver get review body read err(%s)\n", err.Error())
		return -1, n
	}

	err = json.Unmarshal(data, &n)
	if err != nil{
		lprintf(1, "[ERROR] naver get review json parsing err(%s)\n", err.Error())
		return -1, n
	}

	return 1, n
}



func NaverCompList(addr, lat, lng, compNm string) (int, string){

	lprintf(4, "[INFO] get naver id comp info(%s, %s)\n", compNm, addr)

	var naverIdWithAddr, naverIdWithName string

	if len(addr) > 0 {
		rst, nAddr := GetNaverCompWithAddr(addr)
		if rst > 0{
			SetNaverCompWithAddr(nAddr)
			naverIdWithAddr = naverMappingWithAddr(nAddr, addr)
		}
	}

	if len(lat) > 0 && len(lng) > 0{
		rst, nPlace := GetNaverCompWithName(lat, lng, compNm)
		if rst > 0{
			SetNaverCompWithName(nPlace)
			naverIdWithName = naverMappingWithName(nPlace, compNm)
		}
	}

	if len(naverIdWithAddr) > 0{
		return 1, naverIdWithAddr
	}

	if len(naverIdWithName) > 0{
		return 1, naverIdWithName
	}

	return -1, ""
}

func naverMappingWithAddr(n NaverCompsWithAddr, addr string) string{
	for _, v := range n.Place.List {
		if v.Roadaddress == addr || v.Address == addr{
			return v.ID
		}
	}

	return ""
}

func naverMappingWithName(n NaverCompsWithName, name string) string{
	for _, v := range n.Place {
		if v.Title == name {
			return v.ID
		}
	}

	return ""
}

func SetNaverCompWithAddr(n NaverCompsWithAddr) int {

	var query, tmp string

	for _, v := range n.Place.List{
		query = "REPLACE INTO a_naver(NAVER_ID, NAVER_TYPE, COMP_NM, LAT, LNG, CTG, JIBUN_ADDRESS, ROAD_ADDRESS) VALUES(?,?,?,?,?,?,?,?);"

		var params []interface{}
		params = append(params, v.ID)
		params = append(params, v.Type)
		params = append(params, v.Name)
		params = append(params, v.Y)
		params = append(params, v.X)

		for idx, ctg := range v.Category{
			if idx == 0 {
				tmp = ctg
			}else{
				tmp += ","+ctg
			}
		}
		params = append(params, tmp)

		params = append(params, v.Address)
		params = append(params, v.Roadaddress)

		_, err := cls.ExecDBbyParam(query, params)
		if err != nil {
			lprintf(1, "[ERROR] cls.ExecDBbyParam error(%s) \n", err.Error())
			continue
		}
	}

	return 1
}

func SetNaverCompWithName(n NaverCompsWithName) int {

	var query string

	for _, v := range n.Place{
		query = "REPLACE INTO a_naver(NAVER_ID, NAVER_TYPE, COMP_NM, LAT, LNG, DIST, TOTAL_SCORE, CTG, CID, JIBUN_ADDRESS, ROAD_ADDRESS) VALUES(?,?,?,?,?,?,?,?,?,?,?);"

		var params []interface{}
		params = append(params, v.ID)
		params = append(params, v.Type)
		params = append(params, v.Title)
		params = append(params, v.Y)
		params = append(params, v.X)
		params = append(params, v.Dist)
		params = append(params, v.Totalscore)
		params = append(params, v.Ctg)
		params = append(params, v.Cid)
		params = append(params, v.Jibunaddress)
		params = append(params, v.Roadaddress)

		_, err := cls.ExecDBbyParam(query, params)
		if err != nil {
			lprintf(1, "[ERROR] cls.ExecDBbyParam error(%s) \n", err.Error())
			continue
		}
	}

	return 1
}

func GetNaverCompWithAddr(addr string) (int, NaverCompsWithAddr){

	var n NaverCompsWithAddr

	addr = ReplaceAddr(addr)
	url := fmt.Sprintf("v5/api/addresses/place?address=%s&lang=ko&page=1", url.QueryEscape(addr))

	resp, err := cls.HttpRequestDetail("HTTPS", "GET", "map.naver.com", "443", url, nil, nil, "", false)
	if err != nil{
		lprintf(1, "[ERROR] naver get comp addr err(%s)\n", err.Error())
		return -1, n
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		lprintf(1, "[ERROR] naver get comp addr body read err(%s)\n", err.Error())
		return -1, n
	}

	err = json.Unmarshal(data, &n)
	if err != nil{
		lprintf(1, "[ERROR] naver get comp addr parsing json err(%s)\n", err.Error())
		return -1, n
	}

	return 1, n
}

func GetNaverCompWithName(lat, lng, compNm string) (int, NaverCompsWithName){

	var n NaverCompsWithName

	url := fmt.Sprintf("v5/api/instantSearch?lang=ko&caller=pcweb&types=place,address,bus&coords=%s,%s&query=%s", lat, lng, url.QueryEscape(compNm))

	resp, err := cls.HttpRequestDetail("HTTPS", "GET", "map.naver.com", "443", url, nil, nil, "", false)
	if err != nil{
		lprintf(1, "[ERROR] naver get comp addr err(%s)\n", err.Error())
		return -1, n
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		lprintf(1, "[ERROR] naver get comp addr body read err(%s)\n", err.Error())
		return -1, n
	}

	err = json.Unmarshal(data, &n)
	if err != nil{
		lprintf(1, "[ERROR] naver get comp addr parsing json err(%s)\n", err.Error())
		return -1, n
	}

	return 1, n
}