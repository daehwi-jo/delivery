package main

var SelectCompInfo string = `
							SELECT a.rest_id, a.biz_num, a.comp_nm, b.lat, b.lng, b.addr 
							FROM cc_comp_inf a left JOIN priv_rest_info b ON a.rest_id = b.rest_id  
							WHERE a.comp_sts_cd=1 AND length(b.lat) > 0 AND b.lng IS NOT null
							`

var SelectStoreInfo string = `
							SELECT rest_id, biz_num, comp_nm, IFNULL(baemin_id,"") as baemin_id, IFNULL(yogiyo_id,"") as yogiyo_id, IFNULL(naver_id,"") as naver_id, IFNULL(coupang_id,"") as coupang_id 
							FROM b_store 
							WHERE 
								rest_id = '#{restId}'
               				 `

var SelectBaeminReview string = `
							SELECT member_no, contents, rating 
							FROM a_baemin_review 
							WHERE 
								baemin_id = '#{baeminId}' 
							AND DATE 
								BETWEEN '#{stadtDt}' 
								AND '#{endDt}'
`

var SelectStoreInfoAll string = `
							SELECT rest_id, biz_num, comp_nm, IFNULL(baemin_id,"") as baemin_id, IFNULL(yogiyo_id,"") as yogiyo_id, IFNULL(naver_id,"") as naver_id, IFNULL(coupang_id,"") as coupang_id
							FROM b_store
               				 `

var SelectWordCloudReview string = `
							(SELECT contents AS content 
							FROM a_baemin_review 
							WHERE 
								baemin_id ='#{baeminId}' 
							AND LENGTH(contents) > 0 
							AND rating>=5 order by date desc LIMIT 30)
							UNION ALL
							(SELECT body AS content 
							FROM a_naver_review 
							WHERE 
								naver_id = '#{naverId}' 
							AND LENGTH(body) > 0 
							AND rating>=5 order by created desc LIMIT 30)
							UNION ALL
							(SELECT COMMENT AS content 
							FROM a_yogiyo_review 
							WHERE 
								yogiyo_id = '#{yogiyoId}' 
							AND LENGTH(comment) > 0 
							AND rating>=5 order by time desc LIMIT 30)
                             `

var SelectWordCloudLastReview string = `
							(SELECT contents AS content 
							FROM a_baemin_review 
							WHERE 
								baemin_id ='#{baeminId}' 
							AND LENGTH(contents) > 0 
							AND date 
								BETWEEN '#{startDt}' 
								AND '#{endDt}'
							AND rating>=5 order by date desc)
							UNION ALL
							(SELECT body AS content 
							FROM a_naver_review 
							WHERE 
								naver_id = '#{naverId}' 
							AND LENGTH(body) > 0 
							AND date_format(created,'%Y%m%d')
								BETWEEN '#{startDt}' 
								AND '#{endDt}'
							AND rating>=5 order by created desc)
							UNION ALL
							(SELECT COMMENT AS content 
							FROM a_yogiyo_review 
							WHERE 
								yogiyo_id = '#{yogiyoId}' 
							AND LENGTH(comment) > 0 
							AND date_format(time,'%Y%m%d')
								BETWEEN '#{startDt}' 
								AND '#{endDt}'
							AND rating>=5 order by time desc)
                             `

var SelectWordCloudLastReview2 string = `
							(SELECT contents AS content 
							FROM a_baemin_review 
							WHERE 
								baemin_id ='#{baeminId}' 
							AND LENGTH(contents) > 0 
							AND date 
								< '#{endDt}'
							AND rating>=5 order by date desc LIMIT 30)
							UNION ALL
							(SELECT body AS content 
							FROM a_naver_review 
							WHERE 
								naver_id = '#{naverId}' 
							AND LENGTH(body) > 0 
							AND date_format(created,'%Y%m%d')
								< '#{endDt}'
							AND rating>=5 order by created desc LIMIT 30)
							UNION ALL
							(SELECT COMMENT AS content 
							FROM a_yogiyo_review 
							WHERE 
								yogiyo_id = '#{yogiyoId}' 
							AND LENGTH(comment) > 0 
							AND date_format(time,'%Y%m%d') 
								< '#{endDt}'
							AND rating>=5 order by time desc LIMIT 30)
                             `