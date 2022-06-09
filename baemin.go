package main

import (
	"delivery/src/controller/cls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"regexp"
	"strconv"
)

type BaeminMenus struct {
	Status         string `json:"status"`
	Message        string `json:"message"`
	Serverdatetime string `json:"serverDatetime"`
	Data           struct {
		Viewtype string `json:"viewType"`
		ShopInfo struct {
			Themecodes          []int `json:"themeCodes"`
			Deliveryinoperation bool  `json:"deliveryInOperation"`
			Cesco               struct {
				Blue struct {
					Tooltip       string `json:"tooltip"`
					Lastcheckdate string `json:"lastCheckDate"`
				} `json:"blue"`
				White interface{} `json:"white"`
			} `json:"cesco"`
			Usereservedorder  bool        `json:"useReservedOrder"`
			Usedelivery       bool        `json:"useDelivery"`
			Ridersdeliverytip interface{} `json:"ridersDeliveryTip"`
			Actualaddress     struct {
				Address   string  `json:"address"`
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"actualAddress"`
			Ordercounttext    string `json:"orderCountText"`
			Shopstatus        string `json:"shopStatus"`
			Shopstopreason    string `json:"shopStopReason"`
			Indeliveryarea    bool   `json:"inDeliveryArea"`
			Soloshop          bool   `json:"soloShop"`
			Fastdeliverybadge bool   `json:"fastDeliveryBadge"`
			Trackinglog       struct {
				Fastdelivery               bool `json:"fastDelivery"`
				Onepickdelivery            bool `json:"onePickDelivery"`
				Baeminoneactivated         bool `json:"baeminOneActivated"`
				Baeminexpecteddeliverytime struct {
					Type             string  `json:"type"`
					Model            string  `json:"model"`
					Deliverytime     []int   `json:"deliveryTime"`
					Calcdeliverytime float64 `json:"calcDeliveryTime"`
					Puredeliverytime int     `json:"pureDeliveryTime"`
				} `json:"baeminExpectedDeliveryTime"`
			} `json:"trackingLog"`
			Distancetextphrase                       string        `json:"distanceTextPhrase"`
			Servicebadges                            []interface{} `json:"serviceBadges"`
			Baeminorderinoperation                   bool          `json:"baeminOrderInOperation"`
			Usetakeout                               bool          `json:"useTakeout"`
			Usetableorder                            bool          `json:"useTableOrder"`
			Baeminorderdeliverytypetextphrase        string        `json:"baeminOrderDeliveryTypeTextPhrase"`
			Takeoutdiscountrate                      int           `json:"takeoutDiscountRate"`
			Takeoutdiscountprice                     int           `json:"takeoutDiscountPrice"`
			Baeminorderdiscountcosthtml              string        `json:"baeminOrderDiscountCostHtml"`
			Baeminorderdiscounthtmlphrase            string        `json:"baeminOrderDiscountHtmlPhrase"`
			Takeoutdiscountwhenoverminimumorderprice bool          `json:"takeoutDiscountWhenOverMinimumOrderPrice"`
			Expectedcooktime                         string        `json:"expectedCookTime"`
			Expectedcooktooltiptext                  string        `json:"expectedCookTooltipText"`
			Coupons                                  []struct {
				Coupongroupseq   int         `json:"couponGroupSeq"`
				Badges           interface{} `json:"badges"`
				Pricetext        string      `json:"priceText"`
				Expirytext       string      `json:"expiryText"`
				Useconditiontext string      `json:"useConditionText"`
				Issuable         bool        `json:"issuable"`
				Stockout         bool        `json:"stockOut"`
				Types            []string    `json:"types"`
				Minorderprice    interface{} `json:"minOrderPrice"`
			} `json:"coupons"`
			Contents               []interface{} `json:"contents"`
			Convenience            string        `json:"convenience"`
			Statisticstooltiptext  string        `json:"statisticsTooltipText"`
			Reviewcounttext        string        `json:"reviewCountText"`
			Reviewtooltiptext      string        `json:"reviewTooltipText"`
			ShopNo                 string        `json:"Shop_No"`
			ShopNm                 string        `json:"Shop_Nm"`
			TelNo                  string        `json:"Tel_No"`
			VelNo                  string        `json:"Vel_No"`
			FrNo                   string        `json:"Fr_No"`
			FrTelNo                string        `json:"Fr_Tel_No"`
			Addr                   string        `json:"Addr"`
			LocPntLat              float64       `json:"Loc_Pnt_Lat"`
			LocPntLng              float64       `json:"Loc_Pnt_Lng"`
			ReviewCnt              string        `json:"Review_Cnt"`
			StarPntAvg             string        `json:"Star_Pnt_Avg"`
			DlvryTmB               string        `json:"Dlvry_Tm_B"`
			DlvryMiB               string        `json:"Dlvry_Mi_B"`
			DlvryTmE               string        `json:"Dlvry_Tm_E"`
			DlvryMiE               string        `json:"Dlvry_Mi_E"`
			DlvryDate1B            string        `json:"Dlvry_Date_1_B"`
			DlvryDate1E            string        `json:"Dlvry_Date_1_E"`
			DlvryDate2B            string        `json:"Dlvry_Date_2_B"`
			DlvryDate2E            string        `json:"Dlvry_Date_2_E"`
			DlvryDate3B            string        `json:"Dlvry_Date_3_B"`
			DlvryDate3E            string        `json:"Dlvry_Date_3_E"`
			BlockDateB             string        `json:"Block_Date_B"`
			BlockDateE             string        `json:"Block_Date_E"`
			CloseDateB             string        `json:"Close_Date_B"`
			CloseDateE             string        `json:"Close_Date_E"`
			LogoHost               string        `json:"Logo_Host"`
			LogoPath               string        `json:"Logo_Path"`
			LogoFile               string        `json:"Logo_File"`
			DlvryInfo              string        `json:"Dlvry_Info"`
			CloseDay               string        `json:"Close_Day"`
			ShopIntro              string        `json:"Shop_Intro"`
			FavoriteCnt            string        `json:"Favorite_Cnt"`
			ViewCnt                int           `json:"View_Cnt"`
			CallCnt                string        `json:"Call_Cnt"`
			OrdCnt                 string        `json:"Ord_Cnt"`
			CtCd                   string        `json:"Ct_Cd"`
			CtCdNm                 string        `json:"Ct_Cd_Nm"`
			CtCdNmEn               string        `json:"Ct_Cd_Nm_En"`
			CtTyCd                 string        `json:"Ct_Ty_Cd"`
			UseYnOrd               string        `json:"Use_Yn_Ord"`
			UseYnOrdMenu           string        `json:"Use_Yn_Ord_Menu"`
			BizNo                  string        `json:"Biz_No"`
			ShopOwnerNm            string        `json:"Shop_Owner_Nm"`
			OrdAvailYn             string        `json:"Ord_Avail_Yn"`
			SvcShopAdList          []interface{} `json:"Svc_Shop_Ad_List"`
			ShopIconCd             []string      `json:"Shop_Icon_Cd"`
			EvtLandTyVal           string        `json:"Evt_Land_Ty_Val"`
			DhImgHost              string        `json:"Dh_Img_Host"`
			DhImgPath              string        `json:"Dh_Img_Path"`
			DhImgFile              string        `json:"Dh_Img_File"`
			ReviewCntLatest        int           `json:"Review_Cnt_Latest"`
			ReviewCntCeoLatest     int           `json:"Review_Cnt_Ceo_Latest"`
			ReviewCntCeoSayLatest  int           `json:"Review_Cnt_Ceo_Say_Latest"`
			ReviewCntImg           int           `json:"Review_Cnt_Img"`
			ReviewCntCeo           int           `json:"Review_Cnt_Ceo"`
			ReviewCntCeoSay        int           `json:"Review_Cnt_Ceo_Say"`
			CompNo                 string        `json:"Comp_No"`
			CompNm                 string        `json:"Comp_Nm"`
			DhRgnTyCd              string        `json:"Dh_Rgn_Ty_Cd"`
			MovURL                 string        `json:"Mov_Url"`
			ContractStandardFee    string        `json:"Contract_Standard_Fee"`
			ContractSaleFee        string        `json:"Contract_Sale_Fee"`
			ContractSaleFeeYn      string        `json:"Contract_Sale_Fee_Yn"`
			NoncontractStandardFee string        `json:"Noncontract_Standard_Fee"`
			NoncontractSaleFee     string        `json:"Noncontract_Sale_Fee"`
			NoncontractSaleFeeYn   string        `json:"Noncontract_Sale_Fee_Yn"`
			ContractShopYn         string        `json:"Contract_Shop_Yn"`
			BaeminKitchenYn        string        `json:"Baemin_Kitchen_Yn"`
			ShopProm               struct {
				ShopPromCd   string `json:"Shop_Prom_Cd"`
				ShopPromCont string `json:"Shop_Prom_Cont"`
			} `json:"Shop_Prom"`
			CeoNotice struct {
				ReviewCont interface{} `json:"Review_Cont"`
				RegDt      interface{} `json:"Reg_Dt"`
			} `json:"Ceo_Notice"`
			AdYn        string        `json:"Ad_Yn"`
			MeetCash    string        `json:"Meet_Cash"`
			MeetCard    string        `json:"Meet_Card"`
			DlvryTm     string        `json:"Dlvry_Tm"`
			CloseDayTmp string        `json:"Close_Day_Tmp"`
			AwardType   []interface{} `json:"Award_Type"`
			AwardInfo   []interface{} `json:"Award_Info"`
			Cache       string        `json:"Cache"`
			LiveYnShop  string        `json:"Live_Yn_Shop"`
			ShopCpnInfo struct {
			} `json:"Shop_Cpn_Info"`
			ShopCpnYn   string  `json:"Shop_Cpn_Yn"`
			LiveYnOrd   string  `json:"Live_Yn_Ord"`
			ShopBreakYn string  `json:"Shop_Break_Yn"`
			BreakTmInfo string  `json:"Break_Tm_Info"`
			FavoriteYn  string  `json:"Favorite_Yn"`
			Distance    float64 `json:"Distance"`
			DistanceTxt string  `json:"Distance_Txt"`
			Badge       struct {
				Free     string `json:"Free"`
				Discount string `json:"Discount"`
			} `json:"badge"`
			Sanitation struct {
				IsExist bool `json:"IS_EXIST"`
			} `json:"sanitation"`
			CeoNm            string `json:"Ceo_Nm"`
			BusinessLocation string `json:"Business_Location"`
			Deliverytip      struct {
				Version                            string `json:"version"`
				Deliverytipinfophrase              string `json:"deliveryTipInfoPhrase"`
				Deliverytiprangephrasewithdiscount string `json:"deliveryTipRangePhraseWithDiscount"`
				Deliverytiprangephrase             string `json:"deliveryTipRangePhrase"`
				Deliverytipchargephrase            string `json:"deliveryTipChargePhrase"`
				Deliverytipcharges                 []struct {
					Groupname   string `json:"groupName"`
					Groupphrase string `json:"groupPhrase"`
				} `json:"deliveryTipCharges"`
				Deliverytipdetails []struct {
					Index                         int    `json:"index"`
					Orderpricerangephrase         string `json:"orderPriceRangePhrase"`
					Deliverytipphrasewithdiscount string `json:"deliveryTipPhraseWithDiscount"`
					Deliverytipphrase             string `json:"deliveryTipPhrase"`
				} `json:"deliveryTipDetails"`
				Orderpricerangedeliverytips []struct {
					Index                         int    `json:"index"`
					Orderpricerangephrase         string `json:"orderPriceRangePhrase"`
					Deliverytipphrasewithdiscount string `json:"deliveryTipPhraseWithDiscount"`
					Deliverytipphrase             string `json:"deliveryTipPhrase"`
				} `json:"orderPriceRangeDeliveryTips"`
			} `json:"deliveryTip"`
			DlvryExactlyTime                string `json:"Dlvry_Exactly_Time"`
			ExpectedDeliveryTime            string `json:"Expected_Delivery_Time"`
			ExpectedDeliveryTimeTooltipText string `json:"Expected_Delivery_Time_Tooltip_Text"`
			DhFee                           string `json:"Dh_Fee"`
		} `json:"shop_info"`
		ShopMenu struct {
			MenuInfo struct {
				Liquororder struct {
					Menupopup            interface{} `json:"menuPopup"`
					Basketinfophrase     string      `json:"basketInfoPhrase"`
					Basketinfohtmlphrase string      `json:"basketInfoHtmlPhrase"`
				} `json:"liquorOrder"`
				AttCont           string        `json:"Att_Cont"`
				MinOrdPrice       string        `json:"Min_Ord_Price"`
				ShopOrdAtt        string        `json:"Shop_Ord_Att"`
				ShopHeaderImgHost string        `json:"Shop_Header_Img_Host"`
				ShopHeaderImgPath string        `json:"Shop_Header_Img_Path"`
				ShopHeaderImgFile string        `json:"Shop_Header_Img_File"`
				ShopHeaderImg     []interface{} `json:"Shop_Header_Img"`
				BanggaMsg         string        `json:"Bangga_Msg"`
				DsmTelNo          string        `json:"Dsm_Tel_No"`
				MinOrdPriceTxt    string        `json:"Min_Ord_Price_Txt"`
				CtTyCd            string        `json:"Ct_Ty_Cd"`
				OrdTakeTyCd       string        `json:"Ord_Take_Ty_Cd"`
				IsMenuall         string        `json:"IS_MENUALL"`
				BaedalNotice      []string      `json:"Baedal_Notice"`
				FoodOrg           string        `json:"Food_Org"`
				MenuIcon          []interface{} `json:"Menu_Icon"`
				MenuImgURL        string        `json:"Menu_Img_Url"`
				Disposition       struct {
					IsExist  bool          `json:"IS_EXIST"`
					Contents []interface{} `json:"CONTENTS"`
				} `json:"disposition"`
			} `json:"menu_info"`
			MenuOrd struct { // 메뉴 가격 및 이름 정보
				Version string `json:"version"`
				Rec     []struct {
					Menupromotionid interface{} `json:"menuPromotionId"`
					Menustock       interface{} `json:"menuStock"`
					Menupromotion   bool        `json:"menuPromotion"`
					ShopFoodGrpSeq  string      `json:"Shop_Food_Grp_Seq"`
					ShopFoodSeq     string      `json:"Shop_Food_Seq"`
					FoodNm          string      `json:"Food_Nm"`
					ImgURL          string      `json:"Img_Url"`
					Images          []struct {
						Order       int `json:"order"`
						ImageDetail struct {
							Square struct {
								URL    string `json:"url"`
								Width  int    `json:"width"`
								Height int    `json:"height"`
							} `json:"square"`
							Thumbnail struct {
								URL    string `json:"url"`
								Width  int    `json:"width"`
								Height int    `json:"height"`
							} `json:"thumbnail"`
							Rectangle struct {
								URL    string `json:"url"`
								Width  int    `json:"width"`
								Height int    `json:"height"`
							} `json:"rectangle"`
						} `json:"Image_Detail"`
					} `json:"Images"`
					Remark           string `json:"Remark"`
					FoodCont         string `json:"Food_Cont"`
					CookType         string `json:"Cook_Type"`
					UseYnOrd         string `json:"Use_Yn_Ord"`
					MenuTyCd         string `json:"Menu_Ty_Cd"`
					FoodNutrition    string `json:"Food_Nutrition"`
					FoodAllergy      string `json:"Food_Allergy"`
					SoldOut          bool   `json:"Sold_Out"`
					FoodNutritionURL string `json:"Food_Nutrition_Url"`
					ImgUrls          []struct {
						Baseurl string `json:"baseUrl"`
					} `json:"Img_Urls"`
					ListShopFoodPriceGrp []struct {
						ShopFoodSeq         string `json:"Shop_Food_Seq"`
						ShopFoodGrpSeq      string `json:"Shop_Food_Grp_Seq"`
						ShopFoodPriceGrpSeq int    `json:"Shop_Food_Price_Grp_Seq"`
						ShopFoodPriceGrpNm  string `json:"Shop_Food_Price_Grp_Nm"`
						MinSel              string `json:"Min_Sel"`
						MaxSel              string `json:"Max_Sel"`
						DefPriceYn          string `json:"Def_Price_Yn"`
						Discount            bool   `json:"Discount"`
						ListShopFoodPrice   []struct {
							Paymentprice        interface{} `json:"paymentPrice"`
							ShopFoodSeq         string      `json:"Shop_Food_Seq"`
							ShopFoodGrpSeq      string      `json:"Shop_Food_Grp_Seq"`
							ShopFoodPriceGrpSeq int         `json:"Shop_Food_Price_Grp_Seq"`
							ShopFoodPriceSeq    string      `json:"Shop_Food_Price_Seq"`
							FoodPriceNm         string      `json:"Food_Price_Nm"`
							FoodPrice           string      `json:"Food_Price"`
							NormalFoodPrice     string      `json:"Normal_Food_Price"`
							UseYnOrd            string      `json:"Use_Yn_Ord"`
							SoldOut             bool        `json:"Sold_Out"`
						} `json:"List_Shop_Food_Price"`
					} `json:"List_Shop_Food_Price_Grp"`
					Solo           bool `json:"Solo"`
					Representative bool `json:"representative"`
				} `json:"rec"` // 대표 메뉴
				Normal []struct {
					ShopFoodGrpSeq string `json:"Shop_Food_Grp_Seq"`
					ShopFoodGrpNm  string `json:"Shop_Food_Grp_Nm"`
					ImgURL         string `json:"Img_Url"`
					Remark         string `json:"Remark"`
					ListShopFood   []struct {
						Menupromotionid interface{} `json:"menuPromotionId"`
						Menustock       interface{} `json:"menuStock"`
						Menupromotion   bool        `json:"menuPromotion"`
						ShopFoodGrpSeq  string      `json:"Shop_Food_Grp_Seq"`
						ShopFoodSeq     string      `json:"Shop_Food_Seq"`
						FoodNm          string      `json:"Food_Nm"`
						ImgURL          string      `json:"Img_Url"`
						Images          []struct {
							Order       int `json:"order"`
							ImageDetail struct {
								Square struct {
									URL    string `json:"url"`
									Width  int    `json:"width"`
									Height int    `json:"height"`
								} `json:"square"`
								Thumbnail struct {
									URL    string `json:"url"`
									Width  int    `json:"width"`
									Height int    `json:"height"`
								} `json:"thumbnail"`
								Rectangle struct {
									URL    string `json:"url"`
									Width  int    `json:"width"`
									Height int    `json:"height"`
								} `json:"rectangle"`
							} `json:"Image_Detail"`
						} `json:"Images"`
						Remark           string `json:"Remark"`
						FoodCont         string `json:"Food_Cont"`
						CookType         string `json:"Cook_Type"`
						UseYnOrd         string `json:"Use_Yn_Ord"`
						MenuTyCd         string `json:"Menu_Ty_Cd"`
						FoodNutrition    string `json:"Food_Nutrition"`
						FoodAllergy      string `json:"Food_Allergy"`
						SoldOut          bool   `json:"Sold_Out"`
						FoodNutritionURL string `json:"Food_Nutrition_Url"`
						ImgUrls          []struct {
							Baseurl string `json:"baseUrl"`
						} `json:"Img_Urls"`
						ListShopFoodPriceGrp []struct {
							ShopFoodSeq         string `json:"Shop_Food_Seq"`
							ShopFoodGrpSeq      string `json:"Shop_Food_Grp_Seq"`
							ShopFoodPriceGrpSeq int    `json:"Shop_Food_Price_Grp_Seq"`
							ShopFoodPriceGrpNm  string `json:"Shop_Food_Price_Grp_Nm"`
							MinSel              string `json:"Min_Sel"`
							MaxSel              string `json:"Max_Sel"`
							DefPriceYn          string `json:"Def_Price_Yn"`
							Discount            bool   `json:"Discount"`
							ListShopFoodPrice   []struct {
								Paymentprice        interface{} `json:"paymentPrice"`
								ShopFoodSeq         string      `json:"Shop_Food_Seq"`
								ShopFoodGrpSeq      string      `json:"Shop_Food_Grp_Seq"`
								ShopFoodPriceGrpSeq int         `json:"Shop_Food_Price_Grp_Seq"`
								ShopFoodPriceSeq    string      `json:"Shop_Food_Price_Seq"`
								FoodPriceNm         string      `json:"Food_Price_Nm"`
								FoodPrice           string      `json:"Food_Price"`
								NormalFoodPrice     string      `json:"Normal_Food_Price"`
								UseYnOrd            string      `json:"Use_Yn_Ord"`
								SoldOut             bool        `json:"Sold_Out"`
							} `json:"List_Shop_Food_Price"`
						} `json:"List_Shop_Food_Price_Grp"`
						Solo           bool `json:"Solo"`
						Representative bool `json:"representative"`
					} `json:"List_Shop_Food"`
				} `json:"normal"` // 메뉴별 카테고리
				Set  []interface{} `json:"set"`
				Solo []interface{} `json:"solo"`
			} `json:"menu_ord"`
			MenuImg []interface{} `json:"menu_img"`
		} `json:"shop_menu"`
	} `json:"data"`
}

// 가맹점 리뷰 받기
type BaeminReviews struct {
	Status         string `json:"status"`
	Message        string `json:"message"`
	Serverdatetime string `json:"serverDatetime"`
	Data           struct {
		Reviews []struct {
			ID     int64 `json:"id"`
			Member struct {
				Memberno    int64  `json:"memberNo"`
				Nickname    string `json:"nickname"`
				Imageurl    string `json:"imageUrl"`
				Showreviews bool   `json:"showReviews"`
			} `json:"member"`
			Rating                  float64 `json:"rating"`
			Ceoonlymessage          string  `json:"ceoOnlyMessage"`
			Abusingsuspectedmessage string  `json:"abusingSuspectedMessage"`
			Blockmessage            string  `json:"blockMessage"`
			Contents                string  `json:"contents"`
			Modifiable              bool    `json:"modifiable"`
			Deletable               bool    `json:"deletable"`
			Displaystatus           string  `json:"displayStatus"`
			Displaytype             string  `json:"displayType"`
			Menus                   []struct {
				Menuid         int    `json:"menuId"`
				Reviewmenuid   int64  `json:"reviewMenuId"`
				Name           string `json:"name"`
				Recommendation string `json:"recommendation"`
				Contents       string `json:"contents"`
			} `json:"menus"`
			Comments []interface{} `json:"comments"`
			Images   []interface{} `json:"images"`
			Datetext string        `json:"dateText"`
		} `json:"reviews"`
		Shop struct {
			No          int    `json:"no"`
			Name        string `json:"name"`
			Servicetype string `json:"serviceType"`
		} `json:"shop"`
	} `json:"data"`
}

// 가맹점 리스트 얻기
type BaeminComps struct {
	Status         string `json:"status"`
	Message        string `json:"message"`
	Serverdatetime string `json:"serverDatetime"`
	Data           struct {
		Totalcount  int `json:"totalCount"`
		Serviceinfo struct {
			Existresult          bool   `json:"existResult"`
			Requestid            string `json:"requestId"`
			Resultkind           string `json:"resultKind"`
			Noresultimageurl     string `json:"noResultImageUrl"`
			Noresulttext         string `json:"noResultText"`
			Correctedkeywordtext string `json:"correctedKeywordText"`
			Correctedkeyword     string `json:"correctedKeyword"`
			Blocks               []struct {
				Blocktype   string `json:"blockType"`
				Description string `json:"description"`
				Tooltiptext string `json:"tooltipText"`
				Existbar    bool   `json:"existBar"`
				Ad          bool   `json:"ad"`
			} `json:"blocks"`
			Recommend struct {
				Text     string `json:"text"`
				Subtext  string `json:"subText"`
				Category string `json:"category"`
			} `json:"recommend"`
		} `json:"serviceInfo"`
		Shops []struct {
			Blocktype   string `json:"blockType"`
			Landingtype string `json:"landingType"`
			Shopinfo    struct {
				Shopnumber         int    `json:"shopNumber"`
				Shopname           string `json:"shopName"`
				Servicetype        string `json:"serviceType"`
				Categorycode       string `json:"categoryCode"`
				Categorynamekor    string `json:"categoryNameKor"`
				Categorynameeng    string `json:"categoryNameEng"`
				Logourl            string `json:"logoUrl"`
				Introtext          string `json:"introText"`
				Closedaytext       string `json:"closeDayText"`
				Address            string `json:"address"`
				Telnumber          string `json:"telNumber"`
				Virtualtelnumber   string `json:"virtualTelNumber"`
				Franchisenumber    int    `json:"franchiseNumber"`
				Franchisetelnumber string `json:"franchiseTelNumber"`
				Representationmenu string `json:"representationMenu"`
			} `json:"shopInfo"`
			Shopstatus struct {
				Inoperation bool `json:"inOperation"`
			} `json:"shopStatus"`
			Deliveryinfo struct {
				Deliveryareatext           string `json:"deliveryAreaText"`
				Minimumorderprice          int    `json:"minimumOrderPrice"`
				Deliverytipphrase          string `json:"deliveryTipPhrase"`
				Expecteddeliverytimephrase string `json:"expectedDeliveryTimePhrase"`
				Distancephrase             string `json:"distancePhrase"`
				Deliverytipdiscount        bool   `json:"deliveryTipDiscount"`
				Deliverytipzero            bool   `json:"deliveryTipZero"`
				Fastdelivery               bool   `json:"fastDelivery"`
			} `json:"deliveryInfo"`
			Adinfo struct {
				Campaignid string `json:"campaignId"`
			} `json:"adInfo"`
			Shopstatistics struct {
				Averagestarscore      float64 `json:"averageStarScore"`
				Favoritecount         int     `json:"favoriteCount"`
				Latestreviewcount     int     `json:"latestReviewCount"`
				Latestceocommentcount int     `json:"latestCeoCommentCount"`
				Latestordercount      int     `json:"latestOrderCount"`
			} `json:"shopStatistics"`
			Decoinfo struct {
				Thumbnail      bool `json:"thumbnail"`
				Backgrounddeco bool `json:"backgroundDeco"`
				Addonbadges    []struct {
					Type       string `json:"type"`
					Text       string `json:"text"`
					Background struct {
						Color string `json:"color"`
						Alpha string `json:"alpha"`
					} `json:"background"`
					Border struct {
						Color string `json:"color"`
						Alpha string `json:"alpha"`
					} `json:"border"`
					Font struct {
						Color string `json:"color"`
						Alpha string `json:"alpha"`
					} `json:"font"`
				} `json:"addonBadges"`
				Servicebadges []interface{} `json:"serviceBadges"`
				Shopbadges    []interface{} `json:"shopBadges"`
			} `json:"decoInfo"`
			Loginfo struct {
				Displaymenus          []string `json:"displayMenus"`
				Deliverytips          []int    `json:"deliveryTips"`
				Expecteddeliverytimes []int    `json:"expectedDeliveryTimes"`
				Trackinglog           struct {
					Fastdelivery                 bool `json:"fastDelivery"`
					Onepickdelivery              bool `json:"onePickDelivery"`
					Fastdeliverydeliverytiplimit int  `json:"fastDeliveryDeliveryTipLimit"`
				} `json:"trackingLog"`
			} `json:"logInfo"`
		} `json:"shops"`
		Resulttype string `json:"resultType"`
		Extension  string `json:"extension"`
	} `json:"data"`
}

// 가맹점 정보,메뉴 얻기
type BaeminCompInfoMenu struct {
	Status         string `json:"status"`
	Message        string `json:"message"`
	Serverdatetime string `json:"serverDatetime"`
	Data           struct {
		ShopInfo struct {
			Themecodes          []int `json:"themeCodes"`
			Deliveryinoperation bool  `json:"deliveryInOperation"`
			Cesco               struct {
				Blue  interface{} `json:"blue"`
				White interface{} `json:"white"`
			} `json:"cesco"`
			Usereservedorder  bool        `json:"useReservedOrder"`
			Usedelivery       bool        `json:"useDelivery"`
			Ridersdeliverytip interface{} `json:"ridersDeliveryTip"`
			Actualaddress     struct {
				Address   string  `json:"address"`
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"actualAddress"`
			Ordercounttext    string `json:"orderCountText"`
			Shopstatus        string `json:"shopStatus"`
			Shopstopreason    string `json:"shopStopReason"`
			Indeliveryarea    bool   `json:"inDeliveryArea"`
			Soloshop          bool   `json:"soloShop"`
			Fastdeliverybadge bool   `json:"fastDeliveryBadge"`
			Trackinglog       struct {
				Fastdelivery                 bool `json:"fastDelivery"`
				Onepickdelivery              bool `json:"onePickDelivery"`
				Fastdeliverydeliverytiplimit int  `json:"fastDeliveryDeliveryTipLimit"`
				Baeminexpecteddeliverytime   struct {
					Type             string  `json:"type"`
					Model            string  `json:"model"`
					Deliverytime     []int   `json:"deliveryTime"`
					Calcdeliverytime float64 `json:"calcDeliveryTime"`
					Puredeliverytime int     `json:"pureDeliveryTime"`
				} `json:"baeminExpectedDeliveryTime"`
			} `json:"trackingLog"`
			Distancetextphrase                       string        `json:"distanceTextPhrase"`
			Servicebadges                            []interface{} `json:"serviceBadges"`
			Baeminorderinoperation                   bool          `json:"baeminOrderInOperation"`
			Usetakeout                               bool          `json:"useTakeout"`
			Usetableorder                            bool          `json:"useTableOrder"`
			Baeminorderdeliverytypetextphrase        string        `json:"baeminOrderDeliveryTypeTextPhrase"`
			Takeoutdiscountrate                      int           `json:"takeoutDiscountRate"`
			Takeoutdiscountprice                     int           `json:"takeoutDiscountPrice"`
			Baeminorderdiscountcosthtml              string        `json:"baeminOrderDiscountCostHtml"`
			Baeminorderdiscounthtmlphrase            string        `json:"baeminOrderDiscountHtmlPhrase"`
			Takeoutdiscountwhenoverminimumorderprice bool          `json:"takeoutDiscountWhenOverMinimumOrderPrice"`
			Expectedcooktime                         string        `json:"expectedCookTime"`
			Expectedcooktooltiptext                  string        `json:"expectedCookTooltipText"`
			Coupons                                  []interface{} `json:"coupons"`
			Contents                                 []interface{} `json:"contents"`
			Convenience                              string        `json:"convenience"`
			Statisticstooltiptext                    string        `json:"statisticsTooltipText"`
			Reviewcounttext                          string        `json:"reviewCountText"`
			Reviewtooltiptext                        string        `json:"reviewTooltipText"`
			ShopNo                                   string        `json:"Shop_No"`
			ShopNm                                   string        `json:"Shop_Nm"`
			TelNo                                    string        `json:"Tel_No"`
			VelNo                                    string        `json:"Vel_No"`
			FrNo                                     string        `json:"Fr_No"`
			FrTelNo                                  string        `json:"Fr_Tel_No"`
			Addr                                     string        `json:"Addr"`
			LocPntLat                                float64       `json:"Loc_Pnt_Lat"`
			LocPntLng                                float64       `json:"Loc_Pnt_Lng"`
			ReviewCnt                                string        `json:"Review_Cnt"`
			StarPntAvg                               string        `json:"Star_Pnt_Avg"`
			DlvryTmB                                 string        `json:"Dlvry_Tm_B"`
			DlvryMiB                                 string        `json:"Dlvry_Mi_B"`
			DlvryTmE                                 string        `json:"Dlvry_Tm_E"`
			DlvryMiE                                 string        `json:"Dlvry_Mi_E"`
			DlvryDate1B                              string        `json:"Dlvry_Date_1_B"`
			DlvryDate1E                              string        `json:"Dlvry_Date_1_E"`
			DlvryDate2B                              string        `json:"Dlvry_Date_2_B"`
			DlvryDate2E                              string        `json:"Dlvry_Date_2_E"`
			DlvryDate3B                              string        `json:"Dlvry_Date_3_B"`
			DlvryDate3E                              string        `json:"Dlvry_Date_3_E"`
			BlockDateB                               string        `json:"Block_Date_B"`
			BlockDateE                               string        `json:"Block_Date_E"`
			CloseDateB                               string        `json:"Close_Date_B"`
			CloseDateE                               string        `json:"Close_Date_E"`
			LogoHost                                 string        `json:"Logo_Host"`
			LogoPath                                 string        `json:"Logo_Path"`
			LogoFile                                 string        `json:"Logo_File"`
			DlvryInfo                                string        `json:"Dlvry_Info"`
			CloseDay                                 string        `json:"Close_Day"`
			ShopIntro                                string        `json:"Shop_Intro"`
			FavoriteCnt                              string        `json:"Favorite_Cnt"`
			ViewCnt                                  int           `json:"View_Cnt"`
			CallCnt                                  string        `json:"Call_Cnt"`
			OrdCnt                                   string        `json:"Ord_Cnt"`
			CtCd                                     string        `json:"Ct_Cd"`
			CtCdNm                                   string        `json:"Ct_Cd_Nm"`
			CtCdNmEn                                 string        `json:"Ct_Cd_Nm_En"`
			CtTyCd                                   string        `json:"Ct_Ty_Cd"`
			UseYnOrd                                 string        `json:"Use_Yn_Ord"`
			UseYnOrdMenu                             string        `json:"Use_Yn_Ord_Menu"`
			BizNo                                    string        `json:"Biz_No"`
			ShopOwnerNm                              string        `json:"Shop_Owner_Nm"`
			OrdAvailYn                               string        `json:"Ord_Avail_Yn"`
			SvcShopAdList                            []interface{} `json:"Svc_Shop_Ad_List"`
			ShopIconCd                               []string      `json:"Shop_Icon_Cd"`
			EvtLandTyVal                             string        `json:"Evt_Land_Ty_Val"`
			DhImgHost                                string        `json:"Dh_Img_Host"`
			DhImgPath                                string        `json:"Dh_Img_Path"`
			DhImgFile                                string        `json:"Dh_Img_File"`
			ReviewCntLatest                          int           `json:"Review_Cnt_Latest"`
			ReviewCntCeoLatest                       int           `json:"Review_Cnt_Ceo_Latest"`
			ReviewCntCeoSayLatest                    int           `json:"Review_Cnt_Ceo_Say_Latest"`
			ReviewCntImg                             int           `json:"Review_Cnt_Img"`
			ReviewCntCeo                             int           `json:"Review_Cnt_Ceo"`
			ReviewCntCeoSay                          int           `json:"Review_Cnt_Ceo_Say"`
			CompNo                                   string        `json:"Comp_No"`
			CompNm                                   string        `json:"Comp_Nm"`
			DhRgnTyCd                                string        `json:"Dh_Rgn_Ty_Cd"`
			MovURL                                   string        `json:"Mov_Url"`
			ContractStandardFee                      string        `json:"Contract_Standard_Fee"`
			ContractSaleFee                          string        `json:"Contract_Sale_Fee"`
			ContractSaleFeeYn                        string        `json:"Contract_Sale_Fee_Yn"`
			NoncontractStandardFee                   string        `json:"Noncontract_Standard_Fee"`
			NoncontractSaleFee                       string        `json:"Noncontract_Sale_Fee"`
			NoncontractSaleFeeYn                     string        `json:"Noncontract_Sale_Fee_Yn"`
			ContractShopYn                           string        `json:"Contract_Shop_Yn"`
			BaeminKitchenYn                          string        `json:"Baemin_Kitchen_Yn"`
			ShopProm                                 struct {
				ShopPromCd   string `json:"Shop_Prom_Cd"`
				ShopPromCont string `json:"Shop_Prom_Cont"`
			} `json:"Shop_Prom"`
			CeoNotice struct {
				ReviewCont interface{} `json:"Review_Cont"`
				RegDt      interface{} `json:"Reg_Dt"`
			} `json:"Ceo_Notice"`
			AdYn        string        `json:"Ad_Yn"`
			MeetCash    string        `json:"Meet_Cash"`
			MeetCard    string        `json:"Meet_Card"`
			DlvryTm     string        `json:"Dlvry_Tm"`
			CloseDayTmp string        `json:"Close_Day_Tmp"`
			AwardType   []interface{} `json:"Award_Type"`
			AwardInfo   []interface{} `json:"Award_Info"`
			Cache       string        `json:"Cache"`
			LiveYnShop  string        `json:"Live_Yn_Shop"`
			ShopCpnInfo struct {
			} `json:"Shop_Cpn_Info"`
			ShopCpnYn   string  `json:"Shop_Cpn_Yn"`
			LiveYnOrd   string  `json:"Live_Yn_Ord"`
			ShopBreakYn string  `json:"Shop_Break_Yn"`
			BreakTmInfo string  `json:"Break_Tm_Info"`
			FavoriteYn  string  `json:"Favorite_Yn"`
			Distance    float64 `json:"Distance"`
			DistanceTxt string  `json:"Distance_Txt"`
			Badge       struct {
				Free     string `json:"Free"`
				Discount string `json:"Discount"`
			} `json:"badge"`
			Sanitation struct {
				IsExist bool `json:"IS_EXIST"`
			} `json:"sanitation"`
			CeoNm            string `json:"Ceo_Nm"`
			BusinessLocation string `json:"Business_Location"`
			Deliverytip      struct {
				Version                            string        `json:"version"`
				Deliverytipinfophrase              string        `json:"deliveryTipInfoPhrase"`
				Deliverytiprangephrasewithdiscount string        `json:"deliveryTipRangePhraseWithDiscount"`
				Deliverytiprangephrase             string        `json:"deliveryTipRangePhrase"`
				Deliverytipchargephrase            string        `json:"deliveryTipChargePhrase"`
				Deliverytipcharges                 []interface{} `json:"deliveryTipCharges"`
				Deliverytipdetails                 []struct {
					Index                         int    `json:"index"`
					Orderpricerangephrase         string `json:"orderPriceRangePhrase"`
					Deliverytipphrasewithdiscount string `json:"deliveryTipPhraseWithDiscount"`
					Deliverytipphrase             string `json:"deliveryTipPhrase"`
				} `json:"deliveryTipDetails"`
				Orderpricerangedeliverytips []struct {
					Index                         int    `json:"index"`
					Orderpricerangephrase         string `json:"orderPriceRangePhrase"`
					Deliverytipphrasewithdiscount string `json:"deliveryTipPhraseWithDiscount"`
					Deliverytipphrase             string `json:"deliveryTipPhrase"`
				} `json:"orderPriceRangeDeliveryTips"`
			} `json:"deliveryTip"`
			DlvryExactlyTime                string `json:"Dlvry_Exactly_Time"`
			ExpectedDeliveryTime            string `json:"Expected_Delivery_Time"`
			ExpectedDeliveryTimeTooltipText string `json:"Expected_Delivery_Time_Tooltip_Text"`
			DhFee                           string `json:"Dh_Fee"`
		} `json:"shop_info"`
		ShopMenu struct {
			MenuInfo struct {
				Liquororder struct {
					Menupopup            interface{} `json:"menuPopup"`
					Basketinfophrase     string      `json:"basketInfoPhrase"`
					Basketinfohtmlphrase string      `json:"basketInfoHtmlPhrase"`
				} `json:"liquorOrder"`
				AttCont           string        `json:"Att_Cont"`
				MinOrdPrice       string        `json:"Min_Ord_Price"`
				ShopOrdAtt        string        `json:"Shop_Ord_Att"`
				ShopHeaderImgHost string        `json:"Shop_Header_Img_Host"`
				ShopHeaderImgPath string        `json:"Shop_Header_Img_Path"`
				ShopHeaderImgFile string        `json:"Shop_Header_Img_File"`
				ShopHeaderImg     []interface{} `json:"Shop_Header_Img"`
				BanggaMsg         string        `json:"Bangga_Msg"`
				DsmTelNo          string        `json:"Dsm_Tel_No"`
				MinOrdPriceTxt    string        `json:"Min_Ord_Price_Txt"`
				CtTyCd            string        `json:"Ct_Ty_Cd"`
				OrdTakeTyCd       string        `json:"Ord_Take_Ty_Cd"`
				IsMenuall         string        `json:"IS_MENUALL"`
				BaedalNotice      []string      `json:"Baedal_Notice"`
				FoodOrg           string        `json:"Food_Org"`
				MenuIcon          []interface{} `json:"Menu_Icon"`
				MenuImgURL        string        `json:"Menu_Img_Url"`
				Disposition       struct {
					IsExist  bool          `json:"IS_EXIST"`
					Contents []interface{} `json:"CONTENTS"`
				} `json:"disposition"`
			} `json:"menu_info"`
			MenuOrd struct {
				Version string `json:"version"`
				Rec     []struct {
					Menupromotionid interface{} `json:"menuPromotionId"`
					Menustock       interface{} `json:"menuStock"`
					Menupromotion   bool        `json:"menuPromotion"`
					ShopFoodGrpSeq  string      `json:"Shop_Food_Grp_Seq"`
					ShopFoodSeq     string      `json:"Shop_Food_Seq"`
					FoodNm          string      `json:"Food_Nm"`
					ImgURL          string      `json:"Img_Url"`
					Images          []struct {
						Order       int `json:"order"`
						ImageDetail struct {
							Normal struct {
								URL    string      `json:"url"`
								Width  interface{} `json:"width"`
								Height interface{} `json:"height"`
							} `json:"normal"`
						} `json:"Image_Detail"`
					} `json:"Images"`
					Remark           string `json:"Remark"`
					FoodCont         string `json:"Food_Cont"`
					CookType         string `json:"Cook_Type"`
					UseYnOrd         string `json:"Use_Yn_Ord"`
					MenuTyCd         string `json:"Menu_Ty_Cd"`
					FoodNutrition    string `json:"Food_Nutrition"`
					FoodAllergy      string `json:"Food_Allergy"`
					SoldOut          bool   `json:"Sold_Out"`
					FoodNutritionURL string `json:"Food_Nutrition_Url"`
					ImgUrls          []struct {
						Baseurl string `json:"baseUrl"`
					} `json:"Img_Urls"`
					ListShopFoodPriceGrp []struct {
						ShopFoodSeq         string `json:"Shop_Food_Seq"`
						ShopFoodGrpSeq      string `json:"Shop_Food_Grp_Seq"`
						ShopFoodPriceGrpSeq int    `json:"Shop_Food_Price_Grp_Seq"`
						ShopFoodPriceGrpNm  string `json:"Shop_Food_Price_Grp_Nm"`
						MinSel              string `json:"Min_Sel"`
						MaxSel              string `json:"Max_Sel"`
						DefPriceYn          string `json:"Def_Price_Yn"`
						Discount            bool   `json:"Discount"`
						ListShopFoodPrice   []struct {
							Paymentprice        interface{} `json:"paymentPrice"`
							ShopFoodSeq         string      `json:"Shop_Food_Seq"`
							ShopFoodGrpSeq      string      `json:"Shop_Food_Grp_Seq"`
							ShopFoodPriceGrpSeq int         `json:"Shop_Food_Price_Grp_Seq"`
							ShopFoodPriceSeq    string      `json:"Shop_Food_Price_Seq"`
							FoodPriceNm         string      `json:"Food_Price_Nm"`
							FoodPrice           string      `json:"Food_Price"`
							NormalFoodPrice     string      `json:"Normal_Food_Price"`
							UseYnOrd            string      `json:"Use_Yn_Ord"`
							SoldOut             bool        `json:"Sold_Out"`
						} `json:"List_Shop_Food_Price"`
					} `json:"List_Shop_Food_Price_Grp"`
					Solo           bool `json:"Solo"`
					Representative bool `json:"representative"`
				} `json:"rec"`
				Normal []struct {
					ShopFoodGrpSeq string `json:"Shop_Food_Grp_Seq"`
					ShopFoodGrpNm  string `json:"Shop_Food_Grp_Nm"`
					ImgURL         string `json:"Img_Url"`
					Remark         string `json:"Remark"`
					ListShopFood   []struct {
						Menupromotionid interface{} `json:"menuPromotionId"`
						Menustock       interface{} `json:"menuStock"`
						Menupromotion   bool        `json:"menuPromotion"`
						ShopFoodGrpSeq  string      `json:"Shop_Food_Grp_Seq"`
						ShopFoodSeq     string      `json:"Shop_Food_Seq"`
						FoodNm          string      `json:"Food_Nm"`
						ImgURL          string      `json:"Img_Url"`
						Images          []struct {
							Order       int `json:"order"`
							ImageDetail struct {
								Normal struct {
									URL    string      `json:"url"`
									Width  interface{} `json:"width"`
									Height interface{} `json:"height"`
								} `json:"normal"`
							} `json:"Image_Detail"`
						} `json:"Images"`
						Remark           string `json:"Remark"`
						FoodCont         string `json:"Food_Cont"`
						CookType         string `json:"Cook_Type"`
						UseYnOrd         string `json:"Use_Yn_Ord"`
						MenuTyCd         string `json:"Menu_Ty_Cd"`
						FoodNutrition    string `json:"Food_Nutrition"`
						FoodAllergy      string `json:"Food_Allergy"`
						SoldOut          bool   `json:"Sold_Out"`
						FoodNutritionURL string `json:"Food_Nutrition_Url"`
						ImgUrls          []struct {
							Baseurl string `json:"baseUrl"`
						} `json:"Img_Urls"`
						ListShopFoodPriceGrp []struct {
							ShopFoodSeq         string `json:"Shop_Food_Seq"`
							ShopFoodGrpSeq      string `json:"Shop_Food_Grp_Seq"`
							ShopFoodPriceGrpSeq int    `json:"Shop_Food_Price_Grp_Seq"`
							ShopFoodPriceGrpNm  string `json:"Shop_Food_Price_Grp_Nm"`
							MinSel              string `json:"Min_Sel"`
							MaxSel              string `json:"Max_Sel"`
							DefPriceYn          string `json:"Def_Price_Yn"`
							Discount            bool   `json:"Discount"`
							ListShopFoodPrice   []struct {
								Paymentprice        interface{} `json:"paymentPrice"`
								ShopFoodSeq         string      `json:"Shop_Food_Seq"`
								ShopFoodGrpSeq      string      `json:"Shop_Food_Grp_Seq"`
								ShopFoodPriceGrpSeq int         `json:"Shop_Food_Price_Grp_Seq"`
								ShopFoodPriceSeq    string      `json:"Shop_Food_Price_Seq"`
								FoodPriceNm         string      `json:"Food_Price_Nm"`
								FoodPrice           string      `json:"Food_Price"`
								NormalFoodPrice     string      `json:"Normal_Food_Price"`
								UseYnOrd            string      `json:"Use_Yn_Ord"`
								SoldOut             bool        `json:"Sold_Out"`
							} `json:"List_Shop_Food_Price"`
						} `json:"List_Shop_Food_Price_Grp"`
						Solo           bool `json:"Solo"`
						Representative bool `json:"representative"`
					} `json:"List_Shop_Food"`
				} `json:"normal"`
				Set  []interface{} `json:"set"`
				Solo []interface{} `json:"solo"`
			} `json:"menu_ord"`
			MenuImg []interface{} `json:"menu_img"`
		} `json:"shop_menu"`
	} `json:"data"`
}

type BaeminCustomerInfo struct {
	Status         string `json:"status"`
	Message        string `json:"message"`
	ServerDatetime string `json:"serverDatetime"`
	Data           struct {
		Member struct {
			Nickname string `json:"nickname"`
			Grade    string `json:"grade"`
			ImageURL string `json:"imageUrl"`
		} `json:"member"`
		ReviewCount int `json:"reviewCount"`
		Reviews     []struct {
			ID                      int64   `json:"id"`
			Rating                  float64 `json:"rating"`
			CeoOnlyMessage          string  `json:"ceoOnlyMessage"`
			AbusingSuspectedMessage string  `json:"abusingSuspectedMessage"`
			BlockMessage            string  `json:"blockMessage"`
			Contents                string  `json:"contents"`
			Modifiable              bool    `json:"modifiable"`
			Deletable               bool    `json:"deletable"`
			DisplayType             string  `json:"displayType"`
			DisplayStatus           string  `json:"displayStatus"`
			MenuDisplayType         string  `json:"menuDisplayType"`
			Menus                   []struct {
				MenuID         int    `json:"menuId"`
				ReviewMenuID   int64  `json:"reviewMenuId"`
				Name           string `json:"name"`
				Recommendation string `json:"recommendation"`
				Contents       string `json:"contents"`
			} `json:"menus"`
			Comments []struct {
				ID             int64  `json:"id"`
				Nickname       string `json:"nickname"`
				ImageURL       string `json:"imageUrl"`
				Contents       string `json:"contents"`
				DisplayStatus  string `json:"displayStatus"`
				BlockMessage   string `json:"blockMessage"`
				CeoOnlyMessage string `json:"ceoOnlyMessage"`
				DateText       string `json:"dateText"`
			} `json:"comments"`
			Images []struct {
				ID  int64  `json:"id"`
				URL string `json:"url"`
			} `json:"images"`
			Shop struct {
				No          int    `json:"no"`
				Name        string `json:"name"`
				ServiceType string `json:"serviceType"`
			} `json:"shop"`
			DateText string `json:"dateText"`
		} `json:"reviews"`
	} `json:"data"`
}

func BaeminCompList(lat, lng, compNm, bizNum string) (int, string){

	rst, b := GetBaeminComp(lat, lng, compNm)
	if rst > 0{
		SetBaeminComp(b)

		for _, vShop := range b.Data.Shops{
			v := vShop.Shopinfo
			rst, baeminBizNum := GetBaeminCompInfo(v.Shopnumber, lat, lng)
			if rst > 0{
				if bizNum == baeminBizNum{
					lprintf(4, "[INFO] bizNum(%s), baeminId(%d)\n", bizNum, v.Shopnumber)
					return 1, strconv.Itoa(v.Shopnumber)
				}
			}
		}
	}

	return -1, ""
}

func GetBaeminComp(lat, lng, compNm string) (int, BaeminComps){
	var b BaeminComps

	url := fmt.Sprintf("v2/SEARCH/shops?keyword=%s&filter=&sort=SORT__DISTANCE&referral=Search&kind=DEFAULT&offset=0&limit=30&latitude=%s&longitude=%s&extension=&appver=10.27.1&carrier=45008&site=7jWXRELC2e&deviceModel=SM-G906S&dvcid=OPUD5687e4685d245b9c&adid=NONE&sessionId=91f4e62a5d1d615dfb1865&osver=23&oscd=2", url.QueryEscape(compNm), lat, lng)

	resp, err := cls.HttpRequestDetail("HTTPS", "GET", "shopdp-api.baemin.com", "443", url, nil, nil, "", false)
	if err != nil{
		lprintf(1, "[ERROR] baemin get comp addr err(%s)\n", err.Error())
		return -1, b
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		lprintf(1, "[ERROR] baemin get comp addr body read err(%s)\n", err.Error())
		return -1, b
	}

	err = json.Unmarshal(data, &b)
	if err != nil{
		lprintf(1, "[ERROR] baemin get comp addr parsing json err(%s)\n", err.Error())
		return -1, b
	}

	return 1, b
}

func SetBaeminComp(b BaeminComps) int {

	var query string

	for _, vShop := range b.Data.Shops{
		query = "REPLACE INTO a_baemin(BAEMIN_ID, COMP_NM, BAEMIN_TYPE, CATEGORY_NAME_KR, CATEGORY_NAME_ENG, LOGO_URL, INTRO_TEXT, CLOSE_DAY_TEXT, ADDRESS, TEL, TEL_VIRTUAL, " +
			"FRANCHISE_NUMBER, FRANCHISE_TEL_NUMBER, REPRESENTATION_MENU, DELIVERY_AREA_TEXT, MINIMUM_ORDER_PRICE, DELIVERY_TIP_PHRASE, EXPECTED_DELIVERY_TIME_PHRASE, DISTANCE_PHRASE, " +
			"DELIVERY_TIP_DISCOUNT, DELIVERY_TIP_ZERO, FASE_DELIVERY) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"

		v := vShop.Shopinfo

		var params []interface{}
		params = append(params, strconv.Itoa(v.Shopnumber))
		params = append(params, v.Shopname)
		params = append(params, v.Servicetype)
		params = append(params, v.Categorynamekor)
		params = append(params, v.Categorynameeng)
		params = append(params, v.Logourl)
		params = append(params, v.Introtext)
		params = append(params, v.Closedaytext)
		params = append(params, v.Address)
		params = append(params, v.Telnumber)
		params = append(params, v.Virtualtelnumber)
		params = append(params, v.Franchisenumber)
		params = append(params, v.Franchisetelnumber)
		params = append(params, v.Representationmenu)
		params = append(params, vShop.Deliveryinfo.Deliveryareatext)
		params = append(params, vShop.Deliveryinfo.Minimumorderprice)
		params = append(params, vShop.Deliveryinfo.Deliverytipphrase)
		params = append(params, vShop.Deliveryinfo.Expecteddeliverytimephrase)
		params = append(params, vShop.Deliveryinfo.Distancephrase)
		params = append(params, vShop.Deliveryinfo.Deliverytipdiscount)
		params = append(params, vShop.Deliveryinfo.Deliverytipzero)
		params = append(params, vShop.Deliveryinfo.Fastdelivery)

		_, err := cls.ExecDBbyParam(query, params)
		if err != nil {
			lprintf(1, "[ERROR] cls.ExecDBbyParam error(%s) \n", err.Error())
			continue
		}
	}

	return 1
}

func GetBaeminCompInfo(baeminId int, lat, lng string) (int, string){
	var b BaeminCompInfoMenu

	url := fmt.Sprintf("v8/shop/%d/detail?lat=%s&lng=%s&limit=25&mem=&memid=&defaultreview=N&campaignId=-1&displayGroup=SEARCH&lat4Distance=37.52165288&lng4Distance=126.92416613&appver=10.27.1&carrier=45008&site=7jWXRELC2e&deviceModel=SM-G906S&dvcid=OPUD5687e4685d245b9c&adid=NONE&sessionId=91f4e62a5d1d615dfb1865&osver=23&oscd=2", baeminId, lat, lng)

	resp, err := cls.HttpRequestDetail("HTTPS", "GET", "shopdp-api.baemin.com", "443", url, nil, nil, "", false)
	if err != nil{
		lprintf(1, "[ERROR] baemin get comp addr err(%s)\n", err.Error())
		return -1, ""
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		lprintf(1, "[ERROR] baemin get comp addr body read err(%s)\n", err.Error())
		return -1, ""
	}

	err = json.Unmarshal(data, &b)
	if err != nil{
		lprintf(1, "[ERROR] baemin get comp addr parsing json err(%s)\n", err.Error())
		return -1, ""
	}

	return 1, b.Data.ShopInfo.BizNo
}

func GetBaeminReviews(baeminId string) int{

	page := 0

	for {
		rst, b := httpBaeminReviews(baeminId, page)
		if rst > 0{

			if len(b.Data.Reviews) == 0{
				if page > 0{
					return page
				}
				lprintf(4, "[INFO] review cnt 0\n")
				break
			}

			page += len(b.Data.Reviews)

			setBaeminReviews(b, baeminId)
			if IsReviewFinished(b.Data.Reviews[len(b.Data.Reviews)-1].Datetext){
				lprintf(4, "[INFO] review finish\n")
				return page
			}
		}else{

			if page > 0{
				return page
			}
			break
		}
	}

	return -1
}

func httpBaeminReviews(baeminId string, page int) (int, BaeminReviews){
	var b BaeminReviews

	url := fmt.Sprintf("v1/shops/%s/reviews?sort=MOST_RECENT&filter=ALL&offset=%d&limit=20&appver=10.27.1&carrier=45008&site=7jWXRELC2e&deviceModel=SM-G906S&dvcid=OPUD5687e4685d245b9c&adid=NONE&sessionId=91f4e62a5d1d615dfb1865&osver=23&oscd=2", baeminId, page)

	var httpHeader map[string]string
	httpHeader = make(map[string]string)
	httpHeader["authorization"] = "bearer guest"

	resp, err := cls.HttpRequestDetail("HTTPS", "GET", "review-api.baemin.com", "443", url, nil, httpHeader, "", false)
	if err != nil{
		lprintf(1, "[ERROR] baemin get comp addr err(%s)\n", err.Error())
		return -1, b
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		lprintf(1, "[ERROR] baemin get comp addr body read err(%s)\n", err.Error())
		return -1, b
	}

	err = json.Unmarshal(data, &b)
	if err != nil{
		lprintf(1, "[ERROR] baemin get comp addr parsing json err(%s)\n", err.Error())
		return -1, b
	}

	return 1, b
}

func httpBaeminCustomerInfo(customerId, limit string) (int, BaeminCustomerInfo){
	var b BaeminCustomerInfo

	url := fmt.Sprintf("v1/members/%s/reviews?offset=0&limit=%s&appver=10.27.1&carrier=45008&site=7jWXRELC2e&deviceModel=SM-G906S&dvcid=OPUD5687e4685d245b9c&adid=NONE&sessionId=f0a64c6f0b15c3f27f6e38&osver=23&oscd=2", customerId, limit)

	resp, err := cls.HttpRequestDetail("HTTPS", "GET", "review-api.baemin.com", "443", url, nil, nil, "", false)
	if err != nil{
		lprintf(1, "[ERROR] baemin get comp addr err(%s)\n", err.Error())
		return -1, b
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		lprintf(1, "[ERROR] baemin get comp addr body read err(%s)\n", err.Error())
		return -1, b
	}

	err = json.Unmarshal(data, &b)
	if err != nil{
		lprintf(1, "[ERROR] baemin get comp addr parsing json err(%s)\n", err.Error())
		return -1, b
	}

	return 1, b
}

func httpBaeminMenus(baeminId string) (int, BaeminMenus){
	var b BaeminMenus

	url := fmt.Sprintf("v8/shop/%s/detail?lat=37.52165288&lng=126.92416613&limit=30&mem=&memid=&defaultreview=N&campaignId=-1&displayGroup=SEARCH&lat4Distance=37.52165288&lng4Distance=126.92416613&appver=10.27.1&carrier=45008&site=7jWXRELC2e&deviceModel=SM-G906S&dvcid=OPUD5687e4685d245b9c&adid=NONE&sessionId=91f4e62a5d1d615dfb1865&osver=23&oscd=2", baeminId)

	resp, err := cls.HttpRequestDetail("HTTPS", "GET", "shopdp-api.baemin.com", "443", url, nil, nil, "", false)
	if err != nil{
		lprintf(1, "[ERROR] baemin get comp addr err(%s)\n", err.Error())
		return -1, b
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		lprintf(1, "[ERROR] baemin get comp addr body read err(%s)\n", err.Error())
		return -1, b
	}

	err = json.Unmarshal(data, &b)
	if err != nil{
		lprintf(1, "[ERROR] baemin get comp addr parsing json err(%s)\n", err.Error())
		return -1, b
	}

	return 1, b
}

func GetBaeminCustomerInfo(memberNo string) int{
	rst, b := httpBaeminCustomerInfo(memberNo, "50")
	if rst < 0{
		return -1
	}

	SetBaeminCustomerInfo(b, memberNo)

	return 1
}

func GetBaeminMenu(baeminId, bizNum, restId string) int{
	rst, b := httpBaeminMenus(baeminId)
	if rst < 0{
		return -1
	}

	setBaeminMenus(b, baeminId, bizNum, restId)

	return 1
}

func setBaeminMenus(b BaeminMenus, baeminId, bizNum, restId string) {

	// 대표 메뉴 다음 카테고리
	for _,v := range b.Data.ShopMenu.MenuOrd.Normal{
		for _, menu := range v.ListShopFood{
			menuNm := menu.FoodNm

			for _, menuList := range menu.ListShopFoodPriceGrp{
				for _, rMenuList := range menuList.ListShopFoodPrice{

					if menuList.ShopFoodPriceGrpNm == "기본"{
						if len(rMenuList.FoodPriceNm) >0 {
							menuNm += " " + rMenuList.FoodPriceNm
						}

						SetMenu(restId, bizNum, "baemin", baeminId, rMenuList.ShopFoodSeq, menuNm, rMenuList.FoodPrice, menu.ImgURL, "", v.ShopFoodGrpNm, "", 0)
						menuNm = menu.FoodNm
					}else{
						SetMenu(restId, bizNum, "baemin", baeminId, rMenuList.ShopFoodPriceSeq, menuList.ShopFoodPriceGrpNm + " " + rMenuList.FoodPriceNm, rMenuList.FoodPrice, "", "", "추가메뉴", rMenuList.ShopFoodSeq, 0)
					}


				}
			}
		}
	}

	// 대표 메뉴
	for _,v := range b.Data.ShopMenu.MenuOrd.Rec{

		menuNm := v.FoodNm

		for _,menu := range v.ListShopFoodPriceGrp{
			for _,menuList := range menu.ListShopFoodPrice{
				// 대표 메뉴의 메인 메뉴
				if menu.ShopFoodPriceGrpNm == "기본"{
					if len(menuList.FoodPriceNm) > 0 {
						menuNm += " " + menuList.FoodPriceNm
					}

					SetMenu(restId, bizNum, "baemin", baeminId, menuList.ShopFoodSeq, menuNm, menuList.FoodPrice, v.ImgURL, "", "대표메뉴", "", 0)
					menuNm = v.FoodNm
				}else{ // 대표 메뉴의 메인 메뉴에 추가되는 서브 메뉴
					SetMenu(restId, bizNum, "baemin", baeminId, menuList.ShopFoodPriceSeq, menuList.FoodPriceNm, menuList.FoodPrice, "", "", "추가메뉴", menuList.ShopFoodSeq, 0)
				}

			}
		}
	}

}

func setBaeminReviews(b BaeminReviews, baeminId string){

	var query, tmp, tmp1, tmp2, tmp3, tmp4 string

	// 이모티콘 제거
	// mysql varchar2 -> uft8 기준(3비트)
	// 이모티콘 utf16? -> 여튼 4비트라서 mysql에 안들어감
	reg,regErr := regexp.Compile("[^\u0000-\uFFFF]")
	if regErr != nil{
		lprintf(1, "[ERROR] regexp compile err(%s)\n")
	}

	for _, v := range b.Data.Reviews{
		query = "REPLACE INTO a_baemin_review(REVIEW_ID, BAEMIN_ID, MEMBER_NO, MEMBER_NAME, MEMBER_URL, MEMBER_SHOW_REVIEWS, RATING, CEO_ONLY_MESSAGE, ABUSING_SUSPECTED_MESSAGE, BLOCK_MESSAGE, CONTENTS, MODIFIABLE, DELETABLE, " +
			"DISPLAY_STATUS, DISPLAY_TYPE, MENU_ID, MENU_REVIEW_ID, MENU_NAME, MENU_RECOMMENDATION, MENU_CONTENTS, COMMENTS, IMAGES, DATETEXT, DATE) " +
			"VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"

		var params []interface{}
		params = append(params, strconv.FormatInt(v.ID, 10))
		params = append(params, baeminId)
		params = append(params, strconv.FormatInt(v.Member.Memberno, 10))
		params = append(params, v.Member.Nickname)
		params = append(params, v.Member.Imageurl)
		params = append(params, v.Member.Showreviews)
		params = append(params, v.Rating)

		if regErr == nil{
			params = append(params, reg.ReplaceAllString(v.Ceoonlymessage, ""))
			params = append(params, reg.ReplaceAllString(v.Abusingsuspectedmessage, ""))
			params = append(params, reg.ReplaceAllString(v.Blockmessage, ""))
			params = append(params, reg.ReplaceAllString(v.Contents, ""))
		}else{
			params = append(params, v.Ceoonlymessage)
			params = append(params, v.Abusingsuspectedmessage)
			params = append(params, v.Blockmessage)
			params = append(params, v.Contents)
		}
		params = append(params, v.Modifiable)
		params = append(params, v.Deletable)
		params = append(params, v.Displaystatus)
		params = append(params, v.Displaytype)

		for idx, v := range v.Menus {
			if idx == 0 {
				tmp = fmt.Sprintf("%d", v.Menuid)
				tmp1 = fmt.Sprintf("%s", strconv.FormatInt(v.Reviewmenuid,10))
				tmp2 = fmt.Sprintf("%s", v.Name)
				tmp3 = fmt.Sprintf("%s", v.Recommendation)

				if regErr == nil{
					tmp4 = fmt.Sprintf("%s", reg.ReplaceAllString(v.Contents, ""))
				}else{
					tmp4 = fmt.Sprintf("%s", v.Contents)
				}

			} else {
				tmp += "," + fmt.Sprintf("%d", v.Menuid)
				tmp1 += "," + fmt.Sprintf("%s", strconv.FormatInt(v.Reviewmenuid,10))
				tmp2 += "," + fmt.Sprintf("%s", v.Name)
				tmp3 += "," + fmt.Sprintf("%s", v.Recommendation)

				if regErr == nil{
					tmp4 += "," + fmt.Sprintf("%s", reg.ReplaceAllString(v.Contents, ""))
				}else{
					tmp4 += "," + fmt.Sprintf("%s", v.Contents)
				}
			}
		}
		params = append(params, tmp)
		params = append(params, tmp1)
		params = append(params, tmp2)
		params = append(params, tmp3)
		params = append(params, tmp4)

		tmp = ""
		for idx, v := range v.Comments {
			if idx == 0 {
				if regErr == nil{
					tmp = fmt.Sprintf("%s", reg.ReplaceAllString(fmt.Sprintf("%v",v), ""))
				}else{
					tmp = fmt.Sprintf("%v", v)
				}
			} else {
				if regErr == nil{
					tmp += "," + fmt.Sprintf("%s", reg.ReplaceAllString(fmt.Sprintf("%v",v), ""))
				}else{
					tmp += "," + fmt.Sprintf("%v", v)
				}
			}
		}
		params = append(params, tmp)

		tmp = ""
		for idx, v := range v.Images {
			if idx == 0 {
				tmp = fmt.Sprintf("%v", v)
			} else {
				tmp += "," + fmt.Sprintf("%v", v)
			}
		}
		params = append(params, tmp)

		params = append(params, v.Datetext)
		params = append(params, krToDate(v.Datetext))

		_, err := cls.ExecDBbyParam(query, params)
		if err != nil {
			lprintf(1, "[ERROR] cls.ExecDBbyParam error(%s) \n", err.Error())
			continue
		}
	}
}

func SetBaeminCustomerInfo(b BaeminCustomerInfo, memberNo string){

	query := "REPLACE INTO a_baemin_customer(REVIEW_ID, MEMBER_NO, MEMBER_NAME, RATING, CONTENTS, MENUS, BAEMIN_ID, BAEMIN_NAME, " +
		"BOSS_NICNAME, BOSS_CONTENTS, DATETEXT, DATE) " +
		"VALUES(?,?,?,?,?,?,?,?,?,?,?,?);"

	// 이모티콘 제거
	// mysql varchar2 -> uft8 기준(3비트)
	// 이모티콘 utf16? -> 여튼 4비트라서 mysql에 안들어감
	reg,regErr := regexp.Compile("[^\u0000-\uFFFF]")
	if regErr != nil{
		lprintf(1, "[ERROR] regexp compile err(%s)\n")
	}

	for _, review := range b.Data.Reviews{

		var params []interface{}

		params = append(params, strconv.FormatInt(review.ID, 10))
		params = append(params, memberNo)
		params = append(params, b.Data.Member.Nickname)
		params = append(params, review.Rating)

		if regErr == nil{
			params = append(params, reg.ReplaceAllString(review.Contents, ""))
		}else{
			params = append(params, review.Contents)
		}

		// 메뉴
		var tmp string
		for _,menu := range review.Menus{
			tmp += fmt.Sprintf("%s,",menu.Name)
		}

		if len(tmp) > 0{
			params = append(params, tmp)
		}else{
			params = append(params, "")
		}

		params = append(params, fmt.Sprintf("%d",review.Shop.No))
		params = append(params, review.Shop.Name)

		if len(review.Comments) > 0{
			params = append(params, review.Comments[0].Nickname)

			if regErr == nil{
				params = append(params, reg.ReplaceAllString(review.Comments[0].Contents, ""))
			}else{
				params = append(params, review.Comments[0].Contents)
			}

		}else{
			params = append(params, "사장님")
			params = append(params, "")
		}

		params = append(params, review.DateText)
		params = append(params, krToDate(review.DateText))

		_, err := cls.ExecDBbyParam(query, params)
		if err != nil {
			lprintf(1, "[ERROR] cls.ExecDBbyParam error(%s) \n", err.Error())
			continue
		}
	}
}