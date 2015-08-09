package main

type TaobaoPingjia struct {
	Watershed      int       `json:"watershed"`
	MaxPage        int       `json:"maxPage"`
	CurrentPageNum int       `json:"currentPageNum"`
	Comments       []Comment `json:"comments"`
	QnaDisabled    bool      `json:"qnaDisabled"`
}
type Comment struct {
	Auction struct {
		Title      string `json:"title"`
		Thumbnail  string `json:"thumbnail"`
		AucNumId   string `json:"aucNumId"`
		Link       string `json:"link"`
		AuctionPic string `json:"auctionPic"`
		Sku        string `json:"sku"`
	} `json:"auction"`
	PromotionType        string        `json:"promotionType"`
	EnableSNS            bool          `json:"enableSNS"`
	AppendCanExplainable bool          `json:"appendCanExplainable"`
	Tag                  string        `json:"tag"`
	ShowCuIcon           bool          `json:"showCuIcon"`
	Award                string        `json:"award"`
	Validscore           int           `json:"validscore"`
	NoQna                bool          `json:"noQna"`
	AppendList           []interface{} `json:"appendList"`
	From                 string        `json:"from"`
	Date                 string        `json:"date"`
	PayTime              struct {
		Time           int64 `json:"time"`
		Minutes        int   `json:"minutes"`
		Seconds        int   `json:"seconds"`
		Hours          int   `json:"hours"`
		Month          int   `json:"month"`
		TimezoneOffset int   `json:"timezoneOffset"`
		Year           int   `json:"year"`
		Day            int   `json:"day"`
		Date           int   `json:"date"`
	} `json:"payTime"`
	DayAfterConfirm int `json:"dayAfterConfirm"`
	BidPriceMoney   struct {
		Amount       int    `json:"amount"`
		Cent         int    `json:"cent"`
		CurrencyCode string `json:"currencyCode"`
		CentFactor   int    `json:"centFactor"`
		DisplayUnit  string `json:"displayUnit"`
		Currency     struct {
			DefaultFractionDigits int    `json:"defaultFractionDigits"`
			CurrencyCode          string `json:"currencyCode"`
			Symbol                string `json:"symbol"`
		} `json:"currency"`
	} `json:"bidPriceMoney"`
	Rate            string        `json:"rate"`
	O2oRate         interface{}   `json:"o2oRate"`
	PropertiesAvg   string        `json:"propertiesAvg"`
	ShowDepositIcon bool          `json:"showDepositIcon"`
	CreditFraudRule int           `json:"creditFraudRule"`
	RateId          int64         `json:"rateId"`
	Useful          int           `json:"useful"`
	Reply           interface{}   `json:"reply"`
	Append          interface{}   `json:"append"`
	SpuRatting      []interface{} `json:"spuRatting"`
	Status          int           `json:"status"`
	RaterType       int           `json:"raterType"`
	Photos          []struct {
		FileId    int64  `json:"fileId"`
		ReceiveId int64  `json:"receiveId"`
		Thumbnail string `json:"thumbnail"`
		Status    int    `json:"status"`
		Url       string `json:"url"`
	} `json:"photos"`
	Content   string `json:"content"`
	Vicious   string `json:"vicious"`
	ShareInfo struct {
		Share           bool   `json:"share"`
		UserNumIdBase64 string `json:"userNumIdBase64"`
		Reply           int    `json:"reply"`
		Pic             int    `json:"pic"`
		LastReplyTime   string `json:"lastReplyTime"`
	} `json:"shareInfo"`
	LastModifyFrom int `json:"lastModifyFrom"`
	User           struct {
		Vip            string `json:"vip"`
		Rank           int    `json:"rank"`
		Nick           string `json:"nick"`
		UserId         string `json:"userId"`
		DisplayRatePic string `json:"displayRatePic"`
		NickUrl        string `json:"nickUrl"`
		VipLevel       int    `json:"vipLevel"`
		Avatar         string `json:"avatar"`
		Anony          bool   `json:"anony"`
		RankUrl        string `json:"rankUrl"`
	} `json:"user"`
	BuyAmount int `json:"buyAmount"`
}
