package main

import (
	"delivery/src/controller/cls"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type InnoForestComp struct {
	SuccessMsg       string `json:"successMsg"`
	TotalDataCount   int    `json:"totalDataCount"`
	TotalPages       int    `json:"totalPages"`
	CurrentPages     int    `json:"currentPages"`
	Offset           int    `json:"offset"`
	CurrentDataCount int    `json:"currentDataCount"`
	Data             []struct {
		CorpID          string `json:"corpId"`
		CorpLogoImg     string `json:"corpLogoImg"`
		InvstSumVal     int64  `json:"invstSumVal"`
		InvstCnt        int    `json:"invstCnt"`
		LastInvstVal    int64  `json:"lastInvstVal"`
		LastInvstAt     string `json:"lastInvstAt"`
		EmpWholeVal     int    `json:"empWholeVal"`
		CapStockVal     int64    `json:"capStockVal"`
		FinacRevenueVal int64    `json:"finacRevenueVal"`
		FinacRevenueAt  string `json:"finacRevenueAt"`
		ShowFlag        bool   `json:"showFlag"`
		CorpViewCnt     int    `json:"corpViewCnt"`
		CorpNameKr      string `json:"corpNameKr"`
		CorpNameEn      string `json:"corpNameEn"`
		CorpIntroKr     string `json:"corpIntroKr"`
		CorpIntroEn     string `json:"corpIntroEn"`
		CorpStockCdKr   string `json:"corpStockCdKr"`
		CorpStockCdEn   string `json:"corpStockCdEn"`
		InvstCdKr       string `json:"invstCdKr"`
		InvstCdEn       string `json:"invstCdEn"`
	} `json:"data"`
	Success bool `json:"success"`
}

func collectInnoForest(){
	rst, v := httpInnoForest(1)
	if rst < 0{
		return
	}

	setInnoForest(v)
}

func setInnoForest(v InnoForestComp){

	var query string

	for _,comp := range v.Data{

		query = "REPLACE INTO a_innoforest(CORP_ID, CORP_LOGO_IMG, INVST_SUM_VAL, INVST_CNT, LAST_INVST_VAL, LAST_INVST_AT, EMP_WHOLE_VAL, CAP_STOCK_VAL, FINAC_REVENUE_VAL" +
			", FINAC_REVENUE_AT, CORP_VIEW_CNT, CORP_NAME_KR, CORP_INTRO_KR, CORP_STOCK_CD_KR, INVST_CD_KR) " +
			"VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"

		var params []interface{}
		params = append(params, comp.CorpID)
		params = append(params, comp.CorpLogoImg)
		params = append(params, comp.InvstSumVal)
		params = append(params, comp.InvstCnt)
		params = append(params, comp.LastInvstVal)
		params = append(params, comp.LastInvstAt)
		params = append(params, comp.EmpWholeVal)

		params = append(params, comp.CapStockVal)
		params = append(params, comp.FinacRevenueVal)
		params = append(params, comp.FinacRevenueAt)
		params = append(params, comp.CorpViewCnt)
		params = append(params, comp.CorpNameKr)
		params = append(params, comp.CorpIntroKr)
		params = append(params, comp.CorpStockCdKr)
		params = append(params, comp.InvstCdKr)

		_, err := cls.ExecDBbyParam(query, params)
		if err != nil {
			lprintf(1, "[ERROR] cls.ExecDBbyParam error(%s) \n", err.Error())
			continue
		}
	}

}

func httpInnoForest(page int)(int, InnoForestComp){

	var v InnoForestComp

	url := fmt.Sprintf("seed/corp/v1/findseedcorpsummary?page=%d&limit=20",page)

	header := make(map[string]string)
	header["User-Agent"] = "PostmanRuntime/7.28.4"
	header["Accept"] = "*/*"
	header["Accept-Encoding"] = "gzip, deflate, br"
	header["Connection"] = "keep-alive"

	resp, err := cls.HttpRequestDetail("HTTP", "GET", "49.50.172.227", "7979", url, nil, header, "application/json", true)
	//resp, err := cls.HttpRequestDetail("HTTPS", "GET", "liveapi.innoforest.co.kr", "443", url, nil, header, "application/json", true)
	if err != nil{
		lprintf(1, "[ERROR] innoForest get comp addr err(%s)\n", err.Error())
		return -1, v
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		lprintf(1, "[ERROR] innoForest get comp addr body read err(%s)\n", err.Error())
		return -1, v
	}

	err = json.Unmarshal(data, &v)
	if err != nil{
		lprintf(1, "[ERROR] innoForest get comp addr parsing json err(%s)\n", err.Error())
		return -1, v
	}

	return 1, v

}