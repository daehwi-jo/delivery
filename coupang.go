package main

import (
	"delivery/src/controller/cls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// 가맹점 리스트 받기
type CoupangComps struct {
	Data struct {
		MappedKeyword interface{} `json:"mappedKeyword"`
		NextToken     interface{} `json:"nextToken"`
		EntityList    []struct {
			Entity struct {
				Data struct {
					OpenStatus            string      `json:"openStatus"`
					OpenStatusText        string      `json:"openStatusText"`
					NextOpenAt            string      `json:"nextOpenAt"`
					RemainingTime         interface{} `json:"remainingTime"`
					Distance              string      `json:"distance"`
					EstimatedDeliveryTime string      `json:"estimatedDeliveryTime"`
					Shareable             bool        `json:"shareable"`
					Benefit               interface{} `json:"benefit"`
					Favorite              bool        `json:"favorite"`
					ReviewRating          float64     `json:"reviewRating"`
					ReviewCount           int         `json:"reviewCount"`
					ReviewCountTexts      []struct {
						Text          string `json:"text"`
						Color         string `json:"color"`
						Size          int    `json:"size"`
						Bold          bool   `json:"bold"`
						StrikeThrough bool   `json:"strikeThrough"`
						Underline     bool   `json:"underline"`
					} `json:"reviewCountTexts"`
					OrderCountText  interface{} `json:"orderCountText"`
					ShowOrderCount  bool        `json:"showOrderCount"`
					DeliveryFeeInfo string      `json:"deliveryFeeInfo"`
					DeliveryFeeText struct {
						Text          string `json:"text"`
						Color         string `json:"color"`
						Size          int    `json:"size"`
						Bold          bool   `json:"bold"`
						StrikeThrough bool   `json:"strikeThrough"`
						Underline     bool   `json:"underline"`
					} `json:"deliveryFeeText"`
					ServiceFeeInfo []struct {
						Text          string `json:"text"`
						Color         string `json:"color"`
						Size          int    `json:"size"`
						Bold          bool   `json:"bold"`
						StrikeThrough bool   `json:"strikeThrough"`
						Underline     bool   `json:"underline"`
					} `json:"serviceFeeInfo"`
					ExtraDocuments interface{} `json:"extraDocuments"`
					Badges         struct {
					} `json:"badges"`
					SearchID             string `json:"searchId"`
					ExpressDeliveryBadge struct {
						TierName     string `json:"tierName"`
						ImagePath    string `json:"imagePath"`
						BadgeToolTip string `json:"badgeToolTip"`
						ExpressBadge bool   `json:"expressBadge"`
					} `json:"expressDeliveryBadge"`
					BenefitCouponInfo      interface{}   `json:"benefitCouponInfo"`
					PreviouslyOrderedInfo  interface{}   `json:"previouslyOrderedInfo"`
					CuratedType            interface{}   `json:"curatedType"`
					NewStoreBadge          interface{}   `json:"newStoreBadge"`
					NameTexts              interface{}   `json:"nameTexts"`
					MatchedDishes          interface{}   `json:"matchedDishes"`
					ReviewCarousels        interface{}   `json:"reviewCarousels"`
					ReviewShortCuts        interface{}   `json:"reviewShortCuts"`
					Emphasized             bool          `json:"emphasized"`
					MinimumOrderThresholds interface{}   `json:"minimumOrderThresholds"`
					AdProperty             interface{}   `json:"adProperty"`
					ID                     int           `json:"id"`
					MerchantID             int           `json:"merchantId"`
					Categories             []string      `json:"categories"`
					PaymentStoreID         string        `json:"paymentStoreId"`
					Name                   string        `json:"name"`
					Description            interface{}   `json:"description"`
					TelNo                  string        `json:"telNo"`
					BizNo                  string        `json:"bizNo"`
					ApprovalStatus         string        `json:"approvalStatus"`
					ZipNo                  string        `json:"zipNo"`
					Address                string        `json:"address"`
					AddressDetail          string        `json:"addressDetail"`
					Latitude               float64       `json:"latitude"`
					Longitude              float64       `json:"longitude"`
					ServiceFeeRatio        float64       `json:"serviceFeeRatio"`
					Menus                  []string      `json:"menus"`
					MenuSource             interface{}   `json:"menuSource"`
					ImagePaths             []string      `json:"imagePaths"`
					TopDishImagePaths      []string      `json:"topDishImagePaths"`
					ImageHeightRatio       float64       `json:"imageHeightRatio"`
					TaxBaseType            string        `json:"taxBaseType"`
					StoreLevelInfoID       int           `json:"storeLevelInfoId"`
					ManuallyShutdown       bool          `json:"manuallyShutdown"`
					Deleted                bool          `json:"deleted"`
					BrandLogoPath          string        `json:"brandLogoPath"`
				} `json:"data"`
			} `json:"entity"`
			ViewType string `json:"viewType"`
		} `json:"entityList"`
	} `json:"data"`
	Error interface{} `json:"error"`
}

type CoupangReviews struct {
	Data struct {
		MappedKeyword interface{} `json:"mappedKeyword"`
		NextToken     interface{} `json:"nextToken"`
		EntityList    []struct {
			Entity struct {
				Data struct {
					ReviewRating float64       `json:"reviewRating"`
					ImagePaths   []interface{} `json:"imagePaths"`
					Writer       string        `json:"writer"`
					WrittenDay   string        `json:"writtenDay"`
					IsOwner      bool          `json:"isOwner"`
					ReviewText   string        `json:"reviewText"`
					ReviewID     int           `json:"reviewId"`
					OrderedMenu  []struct {
						Text          string `json:"text"`
						Color         string `json:"color"`
						Size          int    `json:"size"`
						Bold          bool   `json:"bold"`
						StrikeThrough bool   `json:"strikeThrough"`
						Underline     bool   `json:"underline"`
					} `json:"orderedMenu"`
					IsReportAvailable bool `json:"isReportAvailable"`
					IsEditAvailable   bool `json:"isEditAvailable"`
					ThumbUpInfo       struct {
						ThumbUpCount     int    `json:"thumbUpCount"`
						ThumbAction      string `json:"thumbAction"`
						ThumbUpAvailable bool   `json:"thumbUpAvailable"`
					} `json:"thumbUpInfo"`
					MerchantReply interface{} `json:"merchantReply"`
				} `json:"data"`
			} `json:"entity"`
			ViewType string `json:"viewType"`
		} `json:"entityList"`
		StoreTitle  string `json:"storeTitle"`
		ReviewTitle string `json:"reviewTitle"`
	} `json:"data"`
	Error interface{} `json:"error"`
}

func CoupangCompList(lat, lng, compNm, bizNum string) (int, int){

	rst, c := GetCoupangComp(lat, lng, compNm)
	if rst > 0{
		rst, cId := SetCoupangComp(c, bizNum)
		if rst > 0{
			return 1, cId
		}
	}

	return -1, 0
}

func GetCoupangReviews(coupandId int) int{

	rst, c := httpCoupangReviews(coupandId)

	if rst > 0 && len(c.Data.EntityList) > 0{
		if setCoupangReviews(c, coupandId) > 0{
			return 1
		}
	}

	return -1
}

func setCoupangReviews(c CoupangReviews, coupangId int) int {

	var query, tmp string

	// 이모티콘 제거
	// mysql varchar2 -> uft8 기준(3비트)
	// 이모티콘 utf16? -> 여튼 4비트라서 mysql에 안들어감
	reg,regErr := regexp.Compile("[^\u0000-\uFFFF]")
	if regErr != nil{
		lprintf(1, "[ERROR] regexp compile err(%s)\n")
	}

	for _, v := range c.Data.EntityList{

		c := v.Entity.Data

		if c.ReviewID == 0{
			continue
		}

		query = "REPLACE INTO a_coupang_review(REVIEW_ID, COUPANG_ID, RATING, IMAGEPATH, WRITER, WRITTEN_DAY, DATE, REVIEW_TEXT, ORDER_MENUS) " +
			"VALUES(?,?,?,?,?,?,?,?,?);"

		var params []interface{}
		params = append(params, strconv.Itoa(c.ReviewID))
		params = append(params, strconv.Itoa(coupangId))
		params = append(params, c.ReviewRating)

		if len(c.ImagePaths) > 0{
			params = append(params, c.ImagePaths[0])
		}else{
			params = append(params, "")
		}

		params = append(params, c.Writer)
		params = append(params, c.WrittenDay)
		params = append(params, krToDate(c.WrittenDay))

		if regErr == nil{
			params = append(params, reg.ReplaceAllString(c.ReviewText, ""))
		}else{
			params = append(params, c.ReviewText)
		}

		for idx, v := range c.OrderedMenu{
			if idx == 0{
				tmp = v.Text
			}else{

				if strings.Contains(v.Text, "·"){
					continue
				}

				tmp += "," + v.Text
			}
		}

		params = append(params, tmp)


		_, err := cls.ExecDBbyParam(query, params)
		if err != nil {
			lprintf(1, "[ERROR] cls.ExecDBbyParam error(%s) \n", err.Error())
			continue
		}

	}

	return 1
}

func httpCoupangReviews(coupandId int) (int, CoupangReviews){
	var c CoupangReviews

	url := fmt.Sprintf("endpoint/store.get_reviews?storeId=%d&sort=LATEST_DESC", coupandId)

	var httpHeader map[string]string
	httpHeader = make(map[string]string)
	httpHeader["X-EATS-LOCALE"] = "ko-KR"
	httpHeader["X-EATS-APP-VERSION"] = "1.3.8"
	httpHeader["X-EATS-PCID"] = "03e2556d-8f3b-3f74-982e-5264228f0f78"
	httpHeader["X-EATS-DEVICE-ID"] = "03e2556d-8f3b-3f74-982e-5264228f0f78"
	httpHeader["X-EATS-OS-TYPE"] = "ANDROID"

	resp, err := cls.HttpRequestDetail("HTTPS", "GET", "api.coupangeats.com", "443", url, nil, httpHeader, "", false)
	if err != nil{
		lprintf(1, "[ERROR] coupang get comp addr err(%s)\n", err.Error())
		return -1, c
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		lprintf(1, "[ERROR] coupang get comp addr body read err(%s)\n", err.Error())
		return -1, c
	}

	err = json.Unmarshal(data, &c)
	if err != nil{
		lprintf(1, "[ERROR] coupang get comp addr parsing json err(%s)\n", err.Error())
		return -1, c
	}

	return 1, c
}

func GetCoupangComp(lat, lng, compNm string) (int, CoupangComps){
	var c CoupangComps

	url := fmt.Sprintf("endpoint/store.get_search?keyWord=%s&sort=nearby", url.QueryEscape(compNm))

	var httpHeader map[string]string
	httpHeader = make(map[string]string)
	httpHeader["X-EATS-LOCALE"] = "ko-KR"
	httpHeader["X-EATS-LOCATION"] = fmt.Sprintf("{\"addressId\":0,\"latitude\":%s,\"longitude\":%s,\"regionId\":23}", lat, lng)
	httpHeader["X-EATS-APP-VERSION"] = "1.3.8"
	httpHeader["X-EATS-PCID"] = "03e2556d-8f3b-3f74-982e-5264228f0f78"
	httpHeader["X-EATS-DEVICE-ID"] = "03e2556d-8f3b-3f74-982e-5264228f0f78"
	httpHeader["X-EATS-OS-TYPE"] = "ANDROID"

	resp, err := cls.HttpRequestDetail("HTTPS", "GET", "api.coupangeats.com", "443", url, nil, httpHeader, "", false)
	if err != nil{
		lprintf(1, "[ERROR] coupang get comp addr err(%s)\n", err.Error())
		return -1, c
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		lprintf(1, "[ERROR] coupang get comp addr body read err(%s)\n", err.Error())
		return -1, c
	}

	err = json.Unmarshal(data, &c)
	if err != nil{
		lprintf(1, "[ERROR] coupang get comp addr parsing json err(%s)\n", err.Error())
		return -1, c
	}

	return 1, c
}


func SetCoupangComp(c CoupangComps, bizNum string) (int,int) {

	var query string
	var coupangId int

	for _, v := range c.Data.EntityList{

		c := v.Entity.Data

		if c.ID == 0{
			continue
		}

		var tmp string
		query = "REPLACE INTO a_coupang(COUPANG_ID, CATEGORIES, NAME, PAYMENT_STORE_ID, MERCHANT_ID, DESCRIPTION, TEL_NO, BIZNUM, ZIP_NO, ADDRESS, ADDRESS_DETAIL, " +
			"LAT, LNG, SERVICE_FEE_RATIO, MENUS, IMAGEPATH, TOPDISH_IMAGEPATH, BRANDLOGO_PATH) " +
			"VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"

		var params []interface{}
		params = append(params, strconv.Itoa(c.ID))

		for idx, ctg := range c.Categories{
			if idx == 0 {
				tmp = ctg
			}else{
				tmp += ","+ctg
			}
		}
		params = append(params, tmp)

		params = append(params, c.Name)
		params = append(params, c.PaymentStoreID)
		params = append(params, c.MerchantID)
		params = append(params, c.Description)
		params = append(params, c.TelNo)

		bizNo := strings.ReplaceAll(c.BizNo, "-", "")
		if bizNum == bizNo{
			coupangId = c.ID
		}

		params = append(params, bizNo)
		params = append(params, c.ZipNo)
		params = append(params, c.Address)
		params = append(params, c.AddressDetail)
		params = append(params, c.Latitude)
		params = append(params, c.Longitude)
		params = append(params, c.ServiceFeeRatio)

		tmp = ""
		for idx, menu := range c.Menus{
			if idx == 0 {
				tmp = menu
			}else{
				tmp += ","+menu
			}
		}
		params = append(params, tmp)

		if len(c.ImagePaths) > 0{
			params = append(params, c.ImagePaths[0])
		}else{
			params = append(params, "")
		}

		if len(c.TopDishImagePaths) > 0{
			params = append(params, c.TopDishImagePaths[0])
		}else{
			params = append(params, "")
		}

		params = append(params, c.BrandLogoPath)

		_, err := cls.ExecDBbyParam(query, params)
		if err != nil {
			lprintf(1, "[ERROR] cls.ExecDBbyParam error(%s) \n", err.Error())
			continue
		}
	}

	if coupangId > 0{
		lprintf(4, "[INFO] bizNum(%s) coupang id(%d) \n", bizNum, coupangId)
		return 1, coupangId
	}

	return -1, coupangId
}