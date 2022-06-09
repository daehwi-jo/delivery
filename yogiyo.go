package main

import (
	"delivery/src/controller/cls"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"regexp"
	"strings"
	"strconv"
)

// 가맹점 리스트 받기
type YogiyoComps struct {
	Pagination struct {
		PerPage      int `json:"per_page"`
		TotalObjects int `json:"total_objects"`
		CurrentPage  int `json:"current_page"`
		TotalPages   int `json:"total_pages"`
	} `json:"pagination"`
	Restaurants []struct {
		Rating                            float64     `json:"rating"`
		Subtitle                          string      `json:"subtitle"`
		AdditionalDiscountCurrentlyActive bool        `json:"additional_discount_currently_active"`
		FreeDeliveryThreshold             int         `json:"free_delivery_threshold"`
		DeliveryFeeExplanation            string      `json:"delivery_fee_explanation"`
		IsAvailableDelivery               bool        `json:"is_available_delivery"`
		NextEnd                           interface{} `json:"next_end"`
		DiscountedDeliveryFee             int         `json:"discounted_delivery_fee"`
		Threshold                         int         `json:"threshold"`
		AdvDistance                       interface{} `json:"adv_distance"`
		Open                              bool        `json:"open"`
		CentralBilling                    bool        `json:"central_billing"`
		City                              string      `json:"city"`
		ReviewReplyCount                  int         `json:"review_reply_count"`
		SmsBonus                          bool        `json:"sms_bonus"`
		FranchiseID                       interface{} `json:"franchise_id"`
		PhoneDownlisted                   bool        `json:"phone_downlisted"`
		NewMarkDate                       int         `json:"new_mark_date"`
		FranchiseName                     interface{} `json:"franchise_name"`
		ReviewImageCount                  int         `json:"review_image_count"`
		Top100Restaurant                  bool        `json:"top100_restaurant"`
		HasLoyaltySupport                 bool        `json:"has_loyalty_support"`
		IsAvailablePickup                 bool        `json:"is_available_pickup"`
		RepresentativeMenus               string      `json:"representative_menus"`
		AdditionalDiscountPerMenu         int         `json:"additional_discount_per_menu"`
		New                               bool        `json:"new"`
		DeliveryFee                       int         `json:"delivery_fee"`
		AdditionalDiscountPickup          int         `json:"additional_discount_pickup"`
		Phone                             string      `json:"phone"`
		AdditionalDiscountOnlyForUser     bool        `json:"additional_discount_only_for_user"`
		LogoCuration                      string      `json:"logo_curation"`
		HasTerminal                       bool        `json:"has_terminal"`
		DiscountUntil                     interface{} `json:"discount_until"`
		EstimatedDeliveryTimeKey          interface{} `json:"estimated_delivery_time_key"`
		Categories                        []string    `json:"categories"`
		NextBegin                         interface{} `json:"next_begin"`
		MinOrderAmount                    int         `json:"min_order_amount"`
		Distance                          float64     `json:"distance"`
		LogoURL                           string      `json:"logo_url"`
		ExceptCash                        bool        `json:"except_cash"`
		Name                              string      `json:"name"`
		MenuItemImages                    string      `json:"menu_item_images"`
		SectionPos                        int         `json:"section_pos"`
		RestaurantType                    string      `json:"restaurant_type"`
		RelayMethods                      []string    `json:"relay_methods"`
		OneDish                           interface{}         `json:"one_dish"`
		DiscountPercent                   int         `json:"discount_percent"`
		DeliveryFeeToDisplay              struct {
			Basic string `json:"basic"`
		} `json:"delivery_fee_to_display"`
		AdvertisementType               string      `json:"advertisement_type"`
		PhoneOrder                      bool        `json:"phone_order"`
		AdvertisementRank               int         `json:"advertisement_rank"`
		AdditionalDiscountPickupPerMenu int         `json:"additional_discount_pickup_per_menu"`
		OpenTimeDescription             string      `json:"open_time_description"`
		Zones                           interface{} `json:"zones"`
		AdjustedDeliveryFee             int         `json:"adjusted_delivery_fee"`
		DeliveryMethod                  string      `json:"delivery_method"`
		NewRating                       float64     `json:"new_rating"`
		AppOrder                        bool        `json:"app_order"`
		Premium                         struct {
		} `json:"premium"`
		Lng                                 float64       `json:"lng"`
		ListPos                             int           `json:"list_pos"`
		ID                                  int           `json:"id"`
		CanReview                           int           `json:"can_review"`
		EstimatedDeliveryTime               string        `json:"estimated_delivery_time"`
		AdditionalDiscountPickupOnlyForUser bool          `json:"additional_discount_pickup_only_for_user"`
		ReviewCount                         int           `json:"review_count"`
		OwnerReplyCount                     int           `json:"owner_reply_count"`
		IsDeliverable                       bool          `json:"is_deliverable"`
		Section                             string        `json:"section"`
		LoyalScore                          float64       `json:"loyal_score"`
		ThumbnailURL                        string        `json:"thumbnail_url"`
		Reachable                           bool          `json:"reachable"`
		AdditionalDiscount                  int           `json:"additional_discount"`
		FranchiseTitle                      interface{}   `json:"franchise_title"`
		MinimumPickupMinutes                int           `json:"minimum_pickup_minutes"`
		ThumbnailMessage                    string        `json:"thumbnail_message"`
		Begin                               string        `json:"begin"`
		Description                         string        `json:"description"`
		Tags                                []interface{} `json:"tags"`
		IsAdditionalDiscountEnabled         bool          `json:"is_additional_discount_enabled"`
		Address                             string        `json:"address"`
		Lat                                 float64       `json:"lat"`
		End                                 string        `json:"end"`
		Slug                                string        `json:"slug"`
		ReviewAvg                           float64       `json:"review_avg"`
		DiscountFrom                        interface{}   `json:"discount_from"`
		Top28                               bool          `json:"top28"`
		URL                                 string        `json:"url"`
		FoodflyRestaurantID                 interface{}   `json:"foodfly_restaurant_id"`
		PaymentMethods                      []string      `json:"payment_methods"`
		Keywords                            string        `json:"keywords"`
	} `json:"restaurants"`
}

// 가맹점 정보 받기
type YogiyoCompInfo struct {
	MinOrderAmount          int           `json:"min_order_amount"`
	NutritionFactDesktopURL string        `json:"nutrition_fact_desktop_url"`
	ExceptCash              bool          `json:"except_cash"`
	Description             string        `json:"description"`
	Tags                    []interface{} `json:"tags"`
	NutritionFactMobileURL  string        `json:"nutrition_fact_mobile_url"`
	Address                 string        `json:"address"`
	SuspensionText          interface{}   `json:"suspension_text"`
	AllergyInfoMobileURL    string        `json:"allergy_info_mobile_url"`
	Phone                   string        `json:"phone"`
	Crmdata                 struct {
		CompanyName   string `json:"company_name"`
		CompanyNumber string `json:"company_number"`
	} `json:"crmdata"`
	AllergyInfoDesktopURL  string  `json:"allergy_info_desktop_url"`
	OpeningTimeDescription string  `json:"opening_time_description"`
	Lat                    float64 `json:"lat"`
	IntroductionByOwner    struct {
		Images []struct {
			ThumbnailURL string `json:"thumbnail_url"`
			ImageURL     string `json:"image_url"`
		} `json:"images"`
		IntroductionText string `json:"introduction_text"`
	} `json:"introduction_by_owner"`
	Lng           float64     `json:"lng"`
	Violations    interface{} `json:"violations"`
	CountryOrigin string      `json:"country_origin"`
}

type YogiyoMenus []struct {
	Items []struct {
		OriginalImage  string `json:"original_image"`
		Slug           string `json:"slug"`
		ReviewCount    int    `json:"review_count"`
		Subtitle       string `json:"subtitle"`
		Description    string `json:"description"`
		Price          string `json:"price"`
		IndustrialInfo struct {
		} `json:"industrial_info"`
		OneDish    bool   `json:"one_dish"`
		Image      string `json:"image"`
		Section    string `json:"section"`
		Subchoices []struct {
			Multiple            bool   `json:"multiple"`
			Name                string `json:"name"`
			MultipleCount       int    `json:"multiple_count"`
			IsAvailableQuantity bool   `json:"is_available_quantity"`
			Subchoices          []struct {
				Slug        string `json:"slug"`
				Description string `json:"description"`
				Price       string `json:"price"`
				ID          int    `json:"id"`
				Soldout     bool   `json:"soldout"`
				Name        string `json:"name"`
			} `json:"subchoices"`
			Mandatory bool   `json:"mandatory"`
			Slug      string `json:"slug"`
		} `json:"subchoices"`
		TopDisplayedItemOrder int    `json:"top_displayed_item_order"`
		Soldout               bool   `json:"soldout"`
		ID                    int    `json:"id"`
		Name                  string `json:"name"`
	} `json:"items"`
	Slug        string      `json:"slug"`
	Name        string      `json:"name"`
	Description interface{} `json:"description,omitempty"`
	Image       interface{} `json:"image,omitempty"`
	MsType      interface{} `json:"ms_type,omitempty"`
}

type YogiyoReviews []struct {
	Comment   string `json:"comment"`
	Rating    float64    `json:"rating"`
	Blind     bool   `json:"blind"`
	IsDeleted bool   `json:"is_deleted"`
	MenuItems []struct {
		Ingredients []string `json:"ingredients"`
		ID          int      `json:"id"`
		Name        string   `json:"name"`
	} `json:"menu_items"`
	Level          int           `json:"level"`
	MenuSummary    string        `json:"menu_summary"`
	LikeCount      int           `json:"like_count"`
	IsMineReview   bool          `json:"is_mine_review"`
	Phone          string        `json:"phone"`
	RatingQuantity float64           `json:"rating_quantity"`
	RatingTaste    float64           `json:"rating_taste"`
	IsMineLike     bool          `json:"is_mine_like"`
	RatingDelivery float64           `json:"rating_delivery"`
	Clean          bool          `json:"clean"`
	Time           string        `json:"time"`
	ReviewImages   []struct {
		Full  string `json:"full"`
		Thumb string `json:"thumb"`
	} `json:"review_images"`
	Nickname       string        `json:"nickname"`
	ID             int           `json:"id"`
	IsMenuVisible  bool          `json:"is_menu_visible"`
	OwnerReply     struct {
		Comment   string `json:"comment"`
		CreatedAt string `json:"created_at"`
		ID        int    `json:"id"`
	} `json:"owner_reply"`
}

func YogiyoCompList(lat, lng, compNm, bizNum string) (int, int){

	rst, y := GetYogiyoComp(lat, lng, compNm)
	if rst > 0{
		SetYogiyoComp(y)

		for _, v := range y.Restaurants{
			rst, yogiyoBizNum := GetYogiyoCompInfo(v.ID)
			if rst > 0{
				if bizNum == yogiyoBizNum{
					lprintf(4, "[INFO] yogiyo id(%d), bizNum(%s)", v.ID, bizNum)
					return 1, v.ID
				}
			}
		}
	}

	return -1, 0
}

func GetYogiyoMenu(yogiyoId, bizNum, restId string) int{
	rst, y := httpYogiyoMenus(yogiyoId)
	if rst < 0{
		return -1
	}

	setYogiyoMenus(y, yogiyoId, bizNum, restId)

	return 1
}

func GetYogiyoReviews(yogiyoId int) int{

	page := 1
	totalCnt := 0

	for {
		rst, y := httpYogiyoReviews(yogiyoId, page)
		if rst > 0{

			if len(y) == 0{
				if page > 1{
					return totalCnt
				}
				lprintf(4, "[INFO] review cnt 0\n")
				break
			}

			totalCnt += setYogiyoReviews(y, yogiyoId)
			if IsReviewFinished(y[len(y)-1].Time){
				lprintf(4, "[INFO] review finish\n")
				return totalCnt
			}
		}else{
			if page > 1{
				return totalCnt
			}
			break
		}
		page++
	}

	return -1
}

func SetMenu(restId, bizNum, deliveryNm, deliveryStoreId, menuId, menuNm, menuPrice, imageUrl1, imageUrl2, category, mainMenuId string, menuReviewCnt int){

	query := "REPLACE INTO a_menu(MENU_ID, REST_ID, BIZ_NUM, DELIVERY_NM, DELIVERY_STORE_ID, MENU_NM, MENU_PRICE, " +
		"IMAGE_URL_1, IMAGE_URL_2, MENU_REVIEW_CNT, CATEGORY, MAIN_MENU_ID) " +
		"VALUES(?,?,?,?,?,?,?,?,?,?,?,?);"

	var params []interface{}

	params = append(params, menuId)
	params = append(params, restId)
	params = append(params, bizNum)
	params = append(params, deliveryNm)
	params = append(params, deliveryStoreId)
	params = append(params, menuNm)
	params = append(params, menuPrice)
	params = append(params, imageUrl1)
	params = append(params, imageUrl2)
	params = append(params, menuReviewCnt)
	params = append(params, category)
	params = append(params, mainMenuId)

	_, err := cls.ExecDBbyParam(query, params)
	if err != nil {
		lprintf(1, "[ERROR] cls.ExecDBbyParam error(%s) \n", err.Error())
		return
	}
}

func setYogiyoMenus(y YogiyoMenus, yogiyoId, bizNum, restId string) {
	for _, v := range y{
		for _, menu := range v.Items{
			mainMenuId := strconv.Itoa(menu.ID)
			SetMenu(restId, bizNum, "yogiyo", yogiyoId, mainMenuId, menu.Name, menu.Price, menu.OriginalImage, menu.Image, v.Name, "", menu.ReviewCount)

			for _, sMenu := range menu.Subchoices{
				for _, subMenu := range sMenu.Subchoices{
					SetMenu(restId, bizNum, "yogiyo", yogiyoId, strconv.Itoa(subMenu.ID), subMenu.Name, subMenu.Price, "", "", sMenu.Name, mainMenuId, 0)
				}
			}
		}
	}
}

func setYogiyoReviews(y YogiyoReviews, yogiyoId int) int {

	var query, tmp, tmp1, tmp2 string
	var totalCnt int

	// 이모티콘 제거
	// mysql varchar2 -> uft8 기준(3비트)
	// 이모티콘 utf16? -> 여튼 4비트라서 mysql에 안들어감
	reg,regErr := regexp.Compile("[^\u0000-\uFFFF]")
	if regErr != nil{
		lprintf(1, "[ERROR] regexp compile err(%s)\n")
	}

	for _, v := range y{
		query = "REPLACE INTO a_yogiyo_review(REVIEW_ID, YOGIYO_ID, COMMENT, RATING, BLIND, IS_DELETED, MENU_ID, MENU_NAME, MENU_INGREDIENTS, LEVEL, MENU_SUMMARY, LIKE_COUNT, " +
			"IS_MINE_REVIEW, PHONE, RATING_QUANTITY, RATING_TASTE, IS_MINE_LIKE, RATING_DELEVERY, CLEAN, TIME, REVIEW_IMAGES, NICKNAME, IS_MENU_VISIBLE, OWNER_REPLY) " +
			"VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"

		var params []interface{}
		params = append(params, strconv.Itoa(v.ID))
		params = append(params, strconv.Itoa(yogiyoId))

		if regErr == nil{
			params = append(params, reg.ReplaceAllString(v.Comment, ""))
		}else{
			params = append(params, v.Comment)
		}

		params = append(params, v.Rating)
		params = append(params, v.Blind)
		params = append(params, v.IsDeleted)

		for idx, v := range v.MenuItems{
			if idx == 0{
				tmp = fmt.Sprintf("%d", v.ID)
				tmp1 = fmt.Sprintf("%s", v.Name)
			}else{
				tmp += "," + fmt.Sprintf("%d", v.ID)
				tmp1 += "," + fmt.Sprintf("%s", v.Name)
			}

			for i, vv := range v.Ingredients{
				if i == 0{
					tmp2 = fmt.Sprintf("%s", vv)
				}else{
					tmp2 += "," + fmt.Sprintf("%s", vv)
				}
			}
		}
		params = append(params, tmp)
		params = append(params, tmp1)
		params = append(params, tmp2)

		params = append(params, v.Level)
		params = append(params, v.MenuSummary)
		params = append(params, v.LikeCount)
		params = append(params, v.IsMineReview)
		params = append(params, v.Phone)
		params = append(params, v.RatingQuantity)
		params = append(params, v.RatingTaste)
		params = append(params, v.IsMineLike)
		params = append(params, v.RatingDelivery)
		params = append(params, v.Clean)
		params = append(params, v.Time)

		tmp = ""
		for idx, v := range v.ReviewImages{
			if idx == 0{
				tmp = fmt.Sprintf("%s", v.Full)
			}else{
				tmp += "," + fmt.Sprintf("%v", v.Full)
			}
		}
		params = append(params, tmp)

		params = append(params, v.Nickname)
		params = append(params, v.IsMenuVisible)

		if regErr == nil{
			params = append(params, reg.ReplaceAllString(v.OwnerReply.Comment, ""))
		}else{
			params = append(params, v.OwnerReply.Comment)
		}

		_, err := cls.ExecDBbyParam(query, params)
		if err != nil {
			lprintf(1, "[ERROR] cls.ExecDBbyParam error(%s) \n", err.Error())
			continue
		}

		totalCnt++
	}

	return totalCnt
}

func httpYogiyoMenus(yogiyoId string) (int, YogiyoMenus){
	var y YogiyoMenus

	url := fmt.Sprintf("api/v1/restaurants/%s/menu/?add_photo_menu=android&add_one_dish_menu=true&order_serving_type=delivery", yogiyoId)

	var httpHeader map[string]string
	httpHeader = make(map[string]string)
	httpHeader["x-apisecret"] = "fe5183cc3dea12bd0ce299cf110a75a2"
	httpHeader["x-apikey"] = "iphoneap"

	resp, err := cls.HttpRequestDetail("HTTPS", "GET", "www.yogiyo.co.kr", "443", url, nil, httpHeader, "", false)
	if err != nil{
		lprintf(1, "[ERROR] yogiyo get comp addr err(%s)\n", err.Error())
		return -1, y
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		lprintf(1, "[ERROR] yogiyo get comp addr body read err(%s)\n", err.Error())
		return -1, y
	}

	err = json.Unmarshal(data, &y)
	if err != nil{
		lprintf(1, "[ERROR] yogiyo get comp addr parsing json err(%s)\n", err.Error())
		return -1, y
	}

	return 1, y
}

func httpYogiyoReviews(yogiyoId, page int) (int, YogiyoReviews){
	var y YogiyoReviews

	url := fmt.Sprintf("api/v1/reviews/%d/?count=30&only_photo_review=false&page=%d&sort=time", yogiyoId, page)

	var httpHeader map[string]string
	httpHeader = make(map[string]string)
	httpHeader["x-apisecret"] = "fe5183cc3dea12bd0ce299cf110a75a2"
	httpHeader["x-apikey"] = "iphoneap"

	resp, err := cls.HttpRequestDetail("HTTPS", "GET", "www.yogiyo.co.kr", "443", url, nil, httpHeader, "", false)
	if err != nil{
		lprintf(1, "[ERROR] yogiyo get comp addr err(%s)\n", err.Error())
		return -1, y
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		lprintf(1, "[ERROR] yogiyo get comp addr body read err(%s)\n", err.Error())
		return -1, y
	}

	err = json.Unmarshal(data, &y)
	if err != nil{
		lprintf(1, "[ERROR] yogiyo get comp addr parsing json err(%s)\n", err.Error())
		return -1, y
	}

	return 1, y
}

func GetYogiyoComp(lat, lng, compNm string) (int, YogiyoComps){
	var y YogiyoComps

	url := fmt.Sprintf("api/v1/restaurants-geo/?items=50&lat=%s&lng=%s&order=distance&page=0&category=%s&search=", lat, lng, getCategory(compNm))

	var httpHeader map[string]string
	httpHeader = make(map[string]string)
	httpHeader["x-apisecret"] = "fe5183cc3dea12bd0ce299cf110a75a2"
	httpHeader["x-apikey"] = "iphoneap"

	resp, err := cls.HttpRequestDetail("HTTPS", "GET", "www.yogiyo.co.kr", "443", url, nil, httpHeader, "", false)
	if err != nil{
		lprintf(1, "[ERROR] yogiyo get comp addr err(%s)\n", err.Error())
		return -1, y
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		lprintf(1, "[ERROR] yogiyo get comp addr body read err(%s)\n", err.Error())
		return -1, y
	}

	err = json.Unmarshal(data, &y)
	if err != nil{
		lprintf(1, "[ERROR] yogiyo get comp addr parsing json err(%s)\n", err.Error())
		return -1, y
	}

	return 1, y
}

func GetYogiyoCompInfo(yogiyoId int) (int, string){
	var y YogiyoCompInfo

	url := fmt.Sprintf("api/v1/restaurants/%d/info/", yogiyoId)

	var httpHeader map[string]string
	httpHeader = make(map[string]string)
	httpHeader["x-apisecret"] = "fe5183cc3dea12bd0ce299cf110a75a2"
	httpHeader["x-apikey"] = "iphoneap"

	resp, err := cls.HttpRequestDetail("HTTPS", "GET", "www.yogiyo.co.kr", "443", url, nil, httpHeader, "", false)
	if err != nil{
		lprintf(1, "[ERROR] yogiyo get comp addr err(%s)\n", err.Error())
		return -1, ""
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		lprintf(1, "[ERROR] yogiyo get comp addr body read err(%s)\n", err.Error())
		return -1, ""
	}

	err = json.Unmarshal(data, &y)
	if err != nil{
		lprintf(1, "[ERROR] yogiyo get comp addr parsing json err(%s)\n", err.Error())
		return -1, ""
	}

	return 1, strings.ReplaceAll(y.Crmdata.CompanyNumber, "-", "")
}

func getCategory(compNm string) string{
	if strings.Contains(compNm, "치킨"){
		return "치킨"
	}else if strings.Contains(compNm, "족발") || strings.Contains(compNm, "보쌈"){
		return "족발보쌈"
	}else if strings.Contains(compNm, "피자") || strings.Contains(compNm, "양식") || strings.Contains(compNm, "스파게티"){
		return "피자양식"
	}else if strings.Contains(compNm, "중국집") || strings.Contains(compNm, "반점"){
		return "중식"
	}else if strings.Contains(compNm, "일식") || strings.Contains(compNm, "돈까스"){
		return "일식돈까스"
	}else if strings.Contains(compNm, "야식"){
		return "야식"
	}else if strings.Contains(compNm, "분식") || strings.Contains(compNm, "떡볶이"){
		return "분식"
	}else if strings.Contains(compNm, "카페") || strings.Contains(compNm, "cafe") || strings.Contains(compNm, "커피") || strings.Contains(compNm, "coffee") || strings.Contains(compNm, "디저트"){
		return "카페디저트"
	}

	return "한식"
}

func SetYogiyoComp(y YogiyoComps) int {

	var query, tmp string

	for _, v := range y.Restaurants{
		query = "REPLACE INTO a_yogiyo(YOGIYO_ID, COMP_NM, CATEGORIES, THUMBNAIL_URL, BEGIN, END, ADDRESS, YOGIYO_TYPE, FRANCHISE_TITLE, REVIEW_AVG, CITY, " +
			"DELIVERY_FEE_EXPLANATION, REVIEW_REPLY_COUNT, FRANCHISE_ID, FRANCHISE_NAME, REVIEW_IMAGE_COUNT, OWNER_REPLY_COUNT, LAT, LNG, KEYWORDS, MIN_ORDER_AMOUNT, " +
			"ADDITIONAL_DISCOUNT) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"

		var params []interface{}
		params = append(params, strconv.Itoa(v.ID))
		params = append(params, v.Name)

		for idx, ctg := range v.Categories{
			if idx == 0 {
				tmp = ctg
			}else{
				tmp += ","+ctg
			}
		}
		params = append(params, tmp)

		params = append(params, v.ThumbnailURL)
		params = append(params, v.Begin)
		params = append(params, v.End)
		params = append(params, v.Address)
		params = append(params, v.RestaurantType)
		params = append(params, v.FranchiseTitle)
		params = append(params, v.ReviewAvg)
		params = append(params, v.City)
		params = append(params, v.DeliveryFeeExplanation)
		params = append(params, v.ReviewReplyCount)
		params = append(params, v.FranchiseID)
		params = append(params, v.FranchiseName)
		params = append(params, v.ReviewImageCount)
		params = append(params, v.OwnerReplyCount)
		params = append(params, v.Lat)
		params = append(params, v.Lng)
		params = append(params, v.Keywords)
		params = append(params, v.MinOrderAmount)
		params = append(params, v.AdditionalDiscount)

		_, err := cls.ExecDBbyParam(query, params)
		if err != nil {
			lprintf(1, "[ERROR] cls.ExecDBbyParam error(%s) \n", err.Error())
			continue
		}
	}

	return 1
}