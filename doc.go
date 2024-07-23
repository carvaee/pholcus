// Pholcus（幽灵蛛）是一款纯Go语言编写的支持分布式的高并发、重量级爬虫软件，定位于互联网数据采集，为具备一定Go或JS编程基础的人提供一个只需关注规则定制的功能强大的爬虫工具。
// 它支持单机、服务端、客户端三种运行模式，拥有Web、GUI、命令行三种操作界面；规则简单灵活、批量任务并发、输出方式丰富（mysql/mongodb/kafka/csv/excel等）、有大量Demo共享；另外它还支持横纵向两种抓取模式，支持模拟登录和任务暂停、取消等一系列高级功能。
// （官方QQ群：Go大数据 42731170）。
package main

// 基础包
import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/henrylee2cn/pholcus/app/downloader/request" //必需
	// "github.com/henrylee2cn/pholcus/logs"               //信息输出
	. "github.com/henrylee2cn/pholcus/app/spider" //必需
	// . "github.com/henrylee2cn/pholcus/app/spider/common"          //选用
	// net包
	// "net/http" //设置http.Header
	// "net/url"
	// 编码包
	// "encoding/xml"
	// "encoding/json"
	// 其他包
	// "fmt"
	// "math"
	// "time"
)

func init() {
	XueqiuComment.Register()
}

type StockListResp struct {
	Data *StockList `json:"data"`
}
type StockList struct {
	Count int64          `json:"count"`
	List  []*StockDetail `json:"list"`
}

var (
	allStocks []*StockDetail
	lock      sync.Mutex
)

type StockDetail struct {
	Symbol                   string      `json:"symbol"`
	NetProfitCagr            float64     `json:"net_profit_cagr"`
	NorthNetInflow           float64     `json:"north_net_inflow"`
	Ps                       float64     `json:"ps"`
	Type                     int         `json:"type"`
	Percent                  float64     `json:"percent"`
	HasFollow                bool        `json:"has_follow"`
	TickSize                 float64     `json:"tick_size"`
	PbTtm                    float64     `json:"pb_ttm"`
	FloatShares              int         `json:"float_shares"`
	Current                  float64     `json:"current"`
	Amplitude                float64     `json:"amplitude"`
	Pcf                      interface{} `json:"pcf"`
	CurrentYearPercent       float64     `json:"current_year_percent"`
	FloatMarketCapital       float64     `json:"float_market_capital"`
	NorthNetInflowTime       int64       `json:"north_net_inflow_time"`
	MarketCapital            float64     `json:"market_capital"`
	DividendYield            interface{} `json:"dividend_yield"`
	LotSize                  int         `json:"lot_size"`
	RoeTtm                   float64     `json:"roe_ttm"`
	TotalPercent             float64     `json:"total_percent"`
	Percent5M                float64     `json:"percent5m"`
	IncomeCagr               float64     `json:"income_cagr"`
	Amount                   float64     `json:"amount"`
	Chg                      float64     `json:"chg"`
	IssueDateTs              int64       `json:"issue_date_ts"`
	Eps                      float64     `json:"eps"`
	MainNetInflows           float64     `json:"main_net_inflows"`
	Volume                   int         `json:"volume"`
	VolumeRatio              float64     `json:"volume_ratio"`
	Pb                       float64     `json:"pb"`
	Followers                int         `json:"followers"`
	TurnoverRate             float64     `json:"turnover_rate"`
	MappingQuoteCurrent      interface{} `json:"mapping_quote_current"`
	FirstPercent             float64     `json:"first_percent"`
	Name                     string      `json:"name"`
	PeTtm                    interface{} `json:"pe_ttm"`
	DualCounterMappingSymbol interface{} `json:"dual_counter_mapping_symbol"`
	TotalShares              int         `json:"total_shares"`
	LimitupDays              int         `json:"limitup_days"`
}

type CommentResponse struct {
	About string `json:"about"`
	Count int    `json:"count"`
	Key   string `json:"key"`
	List  []struct {
		Blocked        bool   `json:"blocked"`
		Blocking       bool   `json:"blocking"`
		CanEdit        bool   `json:"canEdit"`
		CommentId      int    `json:"commentId"`
		Controversial  bool   `json:"controversial"`
		CreatedAt      int64  `json:"created_at"`
		Description    string `json:"description"`
		DonateCount    int    `json:"donate_count"`
		DonateSnowcoin int    `json:"donate_snowcoin"`
		Editable       bool   `json:"editable"`
		Expend         bool   `json:"expend"`
		FavCount       int    `json:"fav_count"`
		Favorited      bool   `json:"favorited"`
		Flags          int64  `json:"flags"`
		FlagsObj       struct {
			Flags int64 `json:"flags"`
		} `json:"flagsObj"`
		Hot              bool   `json:"hot"`
		Id               int    `json:"id"`
		IsAnswer         bool   `json:"is_answer"`
		IsBonus          bool   `json:"is_bonus"`
		IsRefused        bool   `json:"is_refused"`
		IsReward         bool   `json:"is_reward"`
		IsSsMultiPic     bool   `json:"is_ss_multi_pic"`
		LegalUserVisible bool   `json:"legal_user_visible"`
		LikeCount        int    `json:"like_count"`
		Liked            bool   `json:"liked"`
		Mark             int    `json:"mark"`
		Pic              string `json:"pic"`
		PromotionId      int    `json:"promotion_id"`
		ReplyCount       int    `json:"reply_count"`
		RetweetCount     int    `json:"retweet_count"`
		RetweetStatusId  int    `json:"retweet_status_id"`
		RewardCount      int    `json:"reward_count"`
		RewardUserCount  int    `json:"reward_user_count"`
		Rqid             int    `json:"rqid"`
		Source           string `json:"source"`
		SourceFeed       bool   `json:"source_feed"`
		Target           string `json:"target"`
		Text             string `json:"text"`
		TimeBefore       string `json:"timeBefore"`
		Title            string `json:"title"`
		TrackJson        string `json:"trackJson"`
		Truncated        bool   `json:"truncated"`
		TruncatedBy      int    `json:"truncated_by"`
		Type             string `json:"type"`
		User             struct {
			AllowAllStock   bool   `json:"allow_all_stock"`
			BlockStatus     int    `json:"block_status"`
			Blocking        bool   `json:"blocking"`
			City            string `json:"city"`
			CommonCount     int    `json:"common_count"`
			Description     string `json:"description"`
			DonateCount     int    `json:"donate_count"`
			FollowMe        bool   `json:"follow_me"`
			FollowersCount  int    `json:"followers_count"`
			Following       bool   `json:"following"`
			FortuneUser     bool   `json:"fortuneUser"`
			FriendsCount    int    `json:"friends_count"`
			Gender          string `json:"gender"`
			Id              int64  `json:"id"`
			LastCommentId   int    `json:"last_comment_id"`
			LastStatusId    int    `json:"last_status_id"`
			PhotoDomain     string `json:"photo_domain"`
			Profile         string `json:"profile"`
			ProfileImageUrl string `json:"profile_image_url"`
			Province        string `json:"province"`
			ScreenName      string `json:"screen_name"`
			StColor         string `json:"st_color"`
			Status          int    `json:"status"`
			StatusCount     int    `json:"status_count"`
			Step            string `json:"step"`
			Subscribeable   bool   `json:"subscribeable"`
			Type            string `json:"type"`
			User            struct {
				AllowAllStock       bool   `json:"allow_all_stock"`
				Anonymous           bool   `json:"anonymous"`
				AreaCode            string `json:"areaCode"`
				Blocking            bool   `json:"blocking"`
				City                string `json:"city"`
				CreatedAt           int64  `json:"created_at"`
				Description         string `json:"description"`
				DonateCount         int    `json:"donate_count"`
				DonateSnowcoin      int    `json:"donate_snowcoin"`
				FollowMe            bool   `json:"follow_me"`
				FollowersCount      int    `json:"followers_count"`
				Following           bool   `json:"following"`
				FriendsCount        int    `json:"friends_count"`
				Gender              string `json:"gender"`
				Id                  int64  `json:"id"`
				LastCommentId       int    `json:"last_comment_id"`
				LastStatusId        int    `json:"last_status_id"`
				MaskedEmail         string `json:"maskedEmail"`
				PhotoDomain         string `json:"photo_domain"`
				Profile             string `json:"profile"`
				ProfileImageUrl     string `json:"profile_image_url"`
				Province            string `json:"province"`
				ScreenName          string `json:"screen_name"`
				StColor             string `json:"st_color"`
				Status              int    `json:"status"`
				StatusCount         int    `json:"status_count"`
				Step                string `json:"step"`
				Truncated           bool   `json:"truncated"`
				Type                string `json:"type"`
				Verified            bool   `json:"verified"`
				VerifiedDescription string `json:"verified_description"`
				VerifiedInfos       []struct {
					VerifiedDesc string `json:"verified_desc"`
					VerifiedType string `json:"verified_type"`
				} `json:"verified_infos"`
				VerifiedType int `json:"verified_type"`
			} `json:"user"`
			Verified            bool   `json:"verified"`
			VerifiedDescription string `json:"verified_description"`
			VerifiedInfos       []struct {
				VerifiedDesc string `json:"verified_desc"`
				VerifiedType string `json:"verified_type"`
			} `json:"verified_infos"`
			VerifiedRealname bool `json:"verified_realname"`
			VerifiedType     int  `json:"verified_type"`
		} `json:"user"`
		UserId    int64 `json:"user_id"`
		ViewCount int   `json:"view_count"`
	} `json:"list"`
	MaxPage        int           `json:"maxPage"`
	Page           int           `json:"page"`
	Q              string        `json:"q"`
	QueryId        int64         `json:"query_id"`
	RecommendCards []interface{} `json:"recommend_cards"`
}

var XueqiuComment = &Spider{
	Name:        "xueqiu",
	Description: "xueqiu帖子 [https://xueqiu.com/hq#hot]",
	// Pausetime: 300,
	// Keyin:   KEYIN,
	// Limit:        LIMIT,
	EnableCookie: true,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			ctx.AddQueue(&request.Request{
				Url:  "https://www.xueqiu.com/hq#hot",
				Rule: "行情中心",
				Temp: map[string]interface{}{"p": 1},
			})
		},

		Trunk: map[string]*Rule{
			"行情中心": {
				ParseFunc: func(ctx *Context) {
					curr := ctx.GetTemp("p", 0).(int)
					// 非第一次进来
					if curr != 1 {
						v := ctx.GetText()
						resp := &StockListResp{}
						if err := json.Unmarshal([]byte(v), &resp); err != nil {
							panic(fmt.Sprintf("unmashal json failed,err=%s", err.Error()))
						}
						// 获取评论
						if resp.Data.Count == 0 {
							for _, stock := range allStocks {
								pg := ctx.GetTemp("page", 1).(int)
								commentUrl := fmt.Sprintf("https://xueqiu.com/query/v1/symbol/search/status.json?count=20&comment=0&symbol=%s&hl=0&source=user&sort=time&page=%d&q=&type=11", stock.Symbol, pg)
								ctx.AddQueue(&request.Request{
									Url:  commentUrl,
									Rule: "评论列表",
									Temp: map[string]interface{}{"page": pg + 1, "symbol": stock.Symbol},
								})
							}
							return
						}
						lock.Lock()
						allStocks = append(allStocks, resp.Data.List...)
						lock.Unlock()
					}
					listUrl := fmt.Sprintf("https://stock.xueqiu.com/v5/stock/screener/quote/list.json?page=%d&size=90&order=desc&order_by=percent&market=CN&type=sh_sz", curr)
					ctx.AddQueue(&request.Request{
						Url:  listUrl,
						Rule: "行情中心",
						Temp: map[string]interface{}{"p": curr + 1},
					})
				},
			},

			"评论列表": {
				ParseFunc: func(ctx *Context) {
					pg := ctx.GetTemp("page", 1).(int)
					symbol := ctx.GetTemp("symbol", "x").(string)
					commentUrl := fmt.Sprintf("https://xueqiu.com/query/v1/symbol/search/status.json?count=20&comment=0&symbol=%s&hl=0&source=user&sort=time&page=%d&q=&type=11", symbol, pg)
					ctx.AddQueue(&request.Request{
						Url:  commentUrl,
						Rule: "评论列表",
						Temp: map[string]interface{}{"page": pg + 1, "symbol": symbol},
					})
					v := ctx.GetText()
					resp := &CommentResponse{}
					if err := json.Unmarshal([]byte(v), &resp); err != nil {
						panic(fmt.Sprintf("unmashal json failed,err=%s", err.Error()))
					}
					fmt.Println(resp)

				},
			},

			"输出结果": {
				//注意：有无字段语义和是否输出数据必须保持一致
				ItemFields: []string{

					"当前积分",
					"帖子数",
					"关注的车",
					"注册时间",
					"作者",
				},
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()

					var 当前积分, 帖子数, 关注的车, 注册时间, 作者 string

					积分 := strings.Split(query.Find(".lv-curr").First().Text(), "当前积分：")
					if len(积分) > 1 {
						当前积分 = 积分[1]
					}

					info := query.Find(".conleft").Eq(0).Find(".leftlist li")

					if len(info.Eq(3).Nodes) > 0 {
						帖子数 = strings.Split(info.Eq(3).Find("a").Text(), "帖")[0]
					}

					for i := 6; !info.Eq(i).HasClass("leftimgs") &&
						len(info.Eq(i).Nodes) > 0 &&
						len(info.Eq(i).Find("a").Nodes) > 0; i++ {
						if strings.Contains(info.Eq(i).Text(), "所属：") {
							continue
						}

						fs := info.Eq(i).Find("a")
						var f string
						if len(fs.Nodes) > 1 {
							f, _ = info.Eq(i).Find("a").Eq(1).Attr("title")
						} else {
							f, _ = info.Eq(i).Find("a").First().Attr("title")
						}
						if f == "" {
							continue
						}
						关注的车 += f + "|"
					}

					关注的车 = strings.Trim(关注的车, "|")

					if len(info.Eq(4).Nodes) > 0 {
						注册 := strings.Split(info.Eq(4).Text(), "注册：")
						if len(注册) > 1 {
							注册时间 = 注册[1]
						}
					}
					作者 = query.Find(".conleft").Eq(0).Find("a").Text()
					// 结果存入Response中转
					ctx.Output(map[int]interface{}{
						0: 当前积分,
						1: 帖子数,
						2: 关注的车,
						3: 注册时间,
						4: 作者,
					})
				},
			},

			// "联系方式": {
			// 	ParseFunc: func(ctx *Context) {
			// 		ctx.AddFile(ctx.GetTemp("n").(string))
			// 	},
			// },
		},
	},
}
