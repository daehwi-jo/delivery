package main

import (
	"delivery/src/controller/cls"
	"github.com/jasonlvhit/gocron"
	"os"
	"strconv"
	"strings"
)

var lprintf func(int, string, ...interface{}) = cls.Lprintf

func main() {

	fname := cls.Cls_conf(os.Args)
	lprintf(3, "++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n")
	lprintf(3, "** start delivery scrapping : fname(%s)\n", fname)
	lprintf(3, "++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n")

	// DB1 connect
	// 수집한 결과 저장하는 db
	ret := cls.Db_conf(fname)
	if ret < 0 {
		lprintf(1, "[ERROR] DB connection error\n")
		return
	}
	defer cls.DBc.Close()

	// DB2 connect
	// 수집하기 위해 comp select 용 db
	ret = cls.Db_conf2(fname)
	if ret < 0 {
		lprintf(1, "[ERROR] DB connection 2 error\n")
		return
	}
	defer cls.DBc2.Close()

	word, r := cls.GetTokenValue("WORD", fname)
	if r == cls.CONF_ERR {
		word = "05:00"
	}

	assa, r := cls.GetTokenValue("ASSA", fname)
	if r == cls.CONF_ERR {
		assa = "11:00"
	}

	insa, r := cls.GetTokenValue("INSA", fname)
	if r == cls.CONF_ERR {
		insa = "11:10"
	}

	collect, r := cls.GetTokenValue("COLLECT", fname)
	if r == cls.CONF_ERR {
		collect = "09:00"
	}

	menu, r := cls.GetTokenValue("MENU", fname)
	if r == cls.CONF_ERR {
		menu = "12:00"
	}

	review, r := cls.GetTokenValue("REVIEW", fname)
	if r == cls.CONF_ERR {
		review = "22:00"
	}

	sftpAddr, r := cls.GetTokenValue("SFTPADDR", fname)
	if r == cls.CONF_ERR {
		sftpAddr = "106.10.42.130"
	}

	sftpPort, r := cls.GetTokenValue("SFTPPORT", fname)
	if r == cls.CONF_ERR {
		sftpPort = "15022"
	}

	sftpId, r := cls.GetTokenValue("SFTPID", fname)
	if r == cls.CONF_ERR {
		sftpId = "root"
	}

	sftpPwd, r := cls.GetTokenValue("SFTPPWD", fname)
	if r == cls.CONF_ERR {
		lprintf(1, "[ERROR] SFTPPWD not found \n")
		return
	}

	g := gocron.NewScheduler()

	// 벤처기업회사
	// https://www.venture.or.kr/#/home/member/h0403
	g.Every(1).Friday().At("09:35").Do(collectVenture)

	// 데이터룸 혁신의 숲
	// https://www.innoforest.co.kr/dataroom?page=1&limit=15
	g.Every(1).Friday().At("09:25").Do(collectInnoForest)

	// 아프니까 사장이다 - 네이버 가맹점 카페
	g.Every(1).Friday().At(assa).Do(CollectASsa)

	// 인사쟁이가 보는 실무까페 - 네이버 인사 담당자 카페
	g.Every(1).Friday().At(insa).Do(CollectINsa)

	// 배달업체 매칭, 리뷰 수집
	g.Every(1).Day().At(collect).Do(collectFull)

	// 배민 메뉴 수집
	g.Every(1).Day().At(menu).Do(collectMenu)

	// 배민 평점, 키워드 작성자 수집
	g.Every(1).Sunday().At("02:00").Do(collectCustomer)

	// 배달업체 리뷰 수집
	g.Every(1).Day().At(review).Do(collectReview)

	// 배달업체 리뷰로 워드클라우드 생성
	g.Every(1).Monday().At(word).Do(makeWordCloud, sftpAddr, sftpPort, sftpId, sftpPwd)

	// 잡코리아 지역별 기업정보 수집
	g.Every(1).Day().At("08:00").Do(GetJobkoreaCompay)

	<- g.Start()
	defer g.Clear()
}

func collectFull(){

	lprintf(4, ">> collectFull START .... << \n")

	rst, comps := GetCompInfo()
	if rst < 0{
		lprintf(1, "[ERROR] get comp err\n")
		return
	}

	for _, comp := range comps{

		var n, y, b, c string

		addr := comp["addr"]
		lat := comp["lat"]
		lng := comp["lng"]
		compNm := comp["comp_nm"]
		restId := comp["rest_id"]
		bizNum := comp["biz_num"]

		// naver
		rst, naverId := NaverCompList(addr, lat, lng, compNm)
		if rst > 0{
			lprintf(4, "[INFO] get comp naverId(%s)\n", naverId)
			if GetNaverReviews(naverId) > 0 {
				n = naverId
			}
		}

		// yogiyo
		rst, yogiyoId := YogiyoCompList(lat, lng, compNm, bizNum)
		if rst > 0{
			lprintf(4, "[INFO] get comp yogiyoId(%d)\n", yogiyoId)
			if GetYogiyoReviews(yogiyoId) > 0{
				y = strconv.Itoa(yogiyoId)
			}
		}

		// baemin
		rst, baeminId := BaeminCompList(lat, lng, compNm, bizNum)
		if rst > 0{
			lprintf(4, "[INFO] get comp baeminId(%s)\n", baeminId)
			if GetBaeminReviews(baeminId) > 0{
				b = baeminId
			}
		}

		// coupang
		rst, coupangId := CoupangCompList(lat, lng, compNm, bizNum)
		if rst > 0{
			lprintf(4, "[INFO] get comp coupangId(%s)\n", coupangId)
			if GetCoupangReviews(coupangId) > 0{
				c = strconv.Itoa(coupangId)
			}
		}

		SetStore(restId, bizNum, compNm, b, y, n, c)
	}

	lprintf(4, ">> collectFull END .... << \n")
}

func collectReview(){

	lprintf(4, ">> collectReview START .... << \n")

	rst, comps := GetStoreInfo("")
	if rst < 0{
		lprintf(1, "[ERROR] get store err\n")
		return
	}

	var reviewCnt int

	for _, comp := range comps{
		b := comp["baemin_id"]
		y := comp["yogiyo_id"]
		n := comp["naver_id"]
		c := comp["coupang_id"]
		bizNum := comp["biz_num"]

		if len(n) > 0{
			lprintf(4, "[INFO] naver id(%s) len(%d)\n", n, len(n))
			reviewCnt = GetNaverReviews(n)
			lprintf(4, "[INFO] biz_num(%s) naver review cnt(%d) \n", bizNum, reviewCnt)
		}

		if len(y) > 0{
			lprintf(4, "[INFO] yogiyo id(%s) len(%d)\n", y, len(y))

			iy,err := strconv.Atoi(y)
			if err != nil{
				lprintf(1, "[ERROR] (%s) atoi err(%s)\n", y, err.Error())
			}else{
				reviewCnt = GetYogiyoReviews(iy)
				lprintf(4, "[INFO] biz_num(%s) yogiyo review cnt(%d) \n", bizNum, reviewCnt)
			}
		}

		if len(b) > 0{
			lprintf(4, "[INFO] baemin id(%s) len(%d)\n", b, len(b))

			reviewCnt = GetBaeminReviews(b)
			lprintf(4, "[INFO] biz_num(%s) baemin review cnt(%d) \n", bizNum, reviewCnt)
		}

		if len(c) > 0{
			ic,_ := strconv.Atoi(c)
			GetCoupangReviews(ic)
		}


		//wordCloudReview(bizNum, b, y, n)
	}

	lprintf(4, ">> collectReview END .... << \n")
}

func collectMenu(){

	lprintf(4, ">> collectMenu START .... << \n")

	rst, comps := GetStoreInfo("")
	if rst < 0{
		lprintf(1, "[ERROR] get store err\n")
		return
	}

	for _, comp := range comps{
		b := comp["baemin_id"]
		y := comp["yogiyo_id"]
		bizNum := comp["biz_num"]
		restId := comp["rest_id"]

		if len(b) > 0{
			lprintf(4, "[INFO] baemin id(%s) len(%d)\n", b, len(b))
			GetBaeminMenu(b, bizNum, restId)
		}else if len(y) > 0{
			lprintf(4, "[INFO] yogiyo id(%s) len(%d)\n", y, len(y))
			GetYogiyoMenu(y, bizNum, restId)
		}
	}

	lprintf(4, ">> collectMenu END .... << \n")
}

func collectCustomer(){

	lprintf(4, ">> collectCustomer START .... << \n")

	rst, comps := GetStoreInfo("")
	if rst < 0{
		lprintf(1, "[ERROR] get store err\n")
		return
	}

	for _, comp := range comps{
		b := comp["baemin_id"]

		if len(b) > 0{

			rst, reviews := GetBaeminReivew(b)
			if rst < 0{
				continue
			}

			for _,review := range reviews{
				// 키워드 매칭
				if review["rating"] == "1" || strings.Contains(review["contents"], "맛있어요"){
					// 작성자 수집
					GetBaeminCustomerInfo(review["member_no"])
				}
			}
		}
	}

	lprintf(4, ">> collectCustomer END .... << \n")
}