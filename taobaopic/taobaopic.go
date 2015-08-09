package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/axgle/mahonia"
)

//type TaobaoPingjia struct {
//Watershed      int       `json:"watershed"`
//MaxPage        int       `json:"maxPage"`
//CurrentPageNum int       `json:"currentPageNum"`
//Comments       []Comment `json:"comments"`
//QnaDisabled    bool      `json:"qnaDisabled"`
//}
//type Comment struct {
//Auction struct {
//Title      string `json:"title"`
//Thumbnail  string `json:"thumbnail"`
//AucNumId   string `json:"aucNumId"`
//Link       string `json:"link"`
//AuctionPic string `json:"auctionPic"`
//Sku        string `json:"sku"`
//} `json:"auction"`
//PromotionType        string        `json:"promotionType"`
//EnableSNS            bool          `json:"enableSNS"`
//AppendCanExplainable bool          `json:"appendCanExplainable"`
//Tag                  string        `json:"tag"`
//ShowCuIcon           bool          `json:"showCuIcon"`
//Award                string        `json:"award"`
//Validscore           int           `json:"validscore"`
//NoQna                bool          `json:"noQna"`
//AppendList           []interface{} `json:"appendList"`
//From                 string        `json:"from"`
//Date                 string        `json:"date"`
//PayTime              struct {
//Time           int64 `json:"time"`
//Minutes        int   `json:"minutes"`
//Seconds        int   `json:"seconds"`
//Hours          int   `json:"hours"`
//Month          int   `json:"month"`
//TimezoneOffset int   `json:"timezoneOffset"`
//Year           int   `json:"year"`
//Day            int   `json:"day"`
//Date           int   `json:"date"`
//} `json:"payTime"`
//DayAfterConfirm int `json:"dayAfterConfirm"`
//BidPriceMoney   struct {
//Amount       int    `json:"amount"`
//Cent         int    `json:"cent"`
//CurrencyCode string `json:"currencyCode"`
//CentFactor   int    `json:"centFactor"`
//DisplayUnit  string `json:"displayUnit"`
//Currency     struct {
//DefaultFractionDigits int    `json:"defaultFractionDigits"`
//CurrencyCode          string `json:"currencyCode"`
//Symbol                string `json:"symbol"`
//} `json:"currency"`
//} `json:"bidPriceMoney"`
//Rate            string        `json:"rate"`
//O2oRate         interface{}   `json:"o2oRate"`
//PropertiesAvg   string        `json:"propertiesAvg"`
//ShowDepositIcon bool          `json:"showDepositIcon"`
//CreditFraudRule int           `json:"creditFraudRule"`
//RateId          int64         `json:"rateId"`
//Useful          int           `json:"useful"`
//Reply           interface{}   `json:"reply"`
//Append          interface{}   `json:"append"`
//SpuRatting      []interface{} `json:"spuRatting"`
//Status          int           `json:"status"`
//RaterType       int           `json:"raterType"`
//Photos          []struct {
//FileId    int64  `json:"fileId"`
//ReceiveId int64  `json:"receiveId"`
//Thumbnail string `json:"thumbnail"`
//Status    int    `json:"status"`
//Url       string `json:"url"`
//} `json:"photos"`
//Content   string `json:"content"`
//Vicious   string `json:"vicious"`
//ShareInfo struct {
//Share           bool   `json:"share"`
//UserNumIdBase64 string `json:"userNumIdBase64"`
//Reply           int    `json:"reply"`
//Pic             int    `json:"pic"`
//LastReplyTime   string `json:"lastReplyTime"`
//} `json:"shareInfo"`
//LastModifyFrom int `json:"lastModifyFrom"`
//User           struct {
//Vip            string `json:"vip"`
//Rank           int    `json:"rank"`
//Nick           string `json:"nick"`
//UserId         string `json:"userId"`
//DisplayRatePic string `json:"displayRatePic"`
//NickUrl        string `json:"nickUrl"`
//VipLevel       int    `json:"vipLevel"`
//Avatar         string `json:"avatar"`
//Anony          bool   `json:"anony"`
//RankUrl        string `json:"rankUrl"`
//} `json:"user"`
//BuyAmount int `json:"buyAmount"`
//}

var Tb *TbClient

type TbClient struct {
	*http.Client
}

func init() {
	Tb = &TbClient{
		http.DefaultClient,
	}
}
func (tb *TbClient) GetPingjiaWithPicUrl(pid string, cp int) string {
	return fmt.Sprintf("http://rate.taobao.com/feedRateList.htm?_ksTS=1433519502403_1786&callback=jsonp_reviews_list&userNumId=50852803&auctionNumId=%s&siteID=1&currentPageNum=%d&rateType3=&orderType=sort_weight&showContent=1&attribute=&ua=", pid, cp)
}

func (tb *TbClient) GetAllTbComments(pid string) (*[]Comment, error) {
	all := make([]Comment, 200)
	pj, err := tb.GetComments(pid, 1)
	if err != nil {
		log.Fatal(err)
	}

	all = append(all, pj.Comments...)
	if pj.Watershed >= 1 {
		for i := 1; i <= pj.Watershed; i++ {
			pj, err := tb.GetComments(pid, pj.CurrentPageNum)
			if err != nil {
				log.Fatal(err)
			}
			all = append(all, pj.Comments...)
		}
	}
	return &all, nil
}
func (tb *TbClient) GetComments(pid string, cp int) (*TaobaoPingjia, error) {
	rs, err := tb.Get(tb.GetPingjiaWithPicUrl(pid, cp))
	defer rs.Body.Close()
	if err != nil {
		log.Printf("err = %+v\n", err)
		log.Fatal(err)
	}
	dc := mahonia.NewDecoder("gb18030")
	r, err := ioutil.ReadAll(dc.NewReader(rs.Body))
	str := string(r)
	rj := strings.Replace(str, "jsonp_reviews_list(", "", 1)
	rj = strings.TrimSpace(rj)
	rj2 := strings.TrimSuffix(rj, ")")

	//log.Println("result:", rj2)
	if err != nil {
		log.Fatal(err)
	}
	var dist TaobaoPingjia
	err = json.Unmarshal([]byte(rj2), &dist)
	if err != nil {
		log.Fatal(err)
	}
	return &dist, nil
}
func main() {
	all, err := Tb.GetAllTbComments("44382142175")
	if err != nil {
		log.Fatal(err)
	}
	l := len(*all)
	_ = "breakpoint"
	log.Println(l, all)
}
