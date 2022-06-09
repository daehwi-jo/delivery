package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"delivery/src/controller/cls"
)

type VentureComp struct {
	Response struct {
		Header struct {
			Status int `json:"status"`
		} `json:"header"`
		Body struct {
			DocCnt int `json:"docCnt"`
			Docs   []struct {
				Page         int         `json:"page"`
				ViewSize     int         `json:"viewSize"`
				TotCnt       int         `json:"totCnt"`
				Title        interface{} `json:"title"`
				ChkComNm     interface{} `json:"chkComNm"`
				ChkAreaNm    interface{} `json:"chkAreaNm"`
				ChkBizKindNm interface{} `json:"chkBizKindNm"`
				ComNm        string      `json:"comNm"`
				CeoNm        string      `json:"ceoNm"`
				BizKind      string      `json:"bizKind"`
				ComFoundYmd  string      `json:"comFoundYmd"`
				ComHomepage  interface{} `json:"comHomepage"`
				ComIntro     string      `json:"comIntro"`
				ComSido      string      `json:"comSido"`
			} `json:"docs"`
		} `json:"body"`
	} `json:"response"`
}

func collectVenture(){

	rst, v := httpVenture("서울", 1)
	if rst < 0{
		return
	}

	setVenture(v)

}

func setVenture(v VentureComp){

	var query string

	for _,comp := range v.Response.Body.Docs{

		query = "REPLACE INTO a_venture(COM_NM, CEO_NM, BIZ_KIND, COM_FOUND_YMD, COM_HOMEPAGE, COM_INTRO, COM_SIDO) VALUES(?,?,?,?,?,?,?);"

		var params []interface{}
		params = append(params, comp.ComNm)
		params = append(params, comp.CeoNm)
		params = append(params, comp.BizKind)
		params = append(params, comp.ComFoundYmd)
		params = append(params, comp.ComHomepage)
		params = append(params, comp.ComIntro)
		params = append(params, comp.ComSido)

		_, err := cls.ExecDBbyParam(query, params)
		if err != nil {
			lprintf(1, "[ERROR] cls.ExecDBbyParam error(%s) \n", err.Error())
			continue
		}
	}

}

func httpVenture(sido string, page int)(int, VentureComp){

	var v VentureComp

	url := fmt.Sprintf("api/openapi/getMemberIntroList.json?chkAreaNm=1&page=%d&title=%s&viewSize=100",page, sido)

	resp, err := cls.HttpRequestDetail("HTTPS", "GET", "www.venture.or.kr", "443", url, nil, nil, "", true)
	if err != nil{
		lprintf(1, "[ERROR] venture get comp addr err(%s)\n", err.Error())
		return -1, v
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		lprintf(1, "[ERROR] venture get comp addr body read err(%s)\n", err.Error())
		return -1, v
	}

	err = json.Unmarshal(data, &v)
	if err != nil{
		lprintf(1, "[ERROR] venture get comp addr parsing json err(%s)\n", err.Error())
		return -1, v
	}

	return 1, v

}