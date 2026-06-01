package hb

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"wxbot/engine/control"
	"wxbot/engine/robot"
)

func init() {
	engine := control.Register("hb", &control.Options{
		Alias: "模拟抢红包",
		Help: "指令:\n" +
			"* 抢红包[金额]元[个数]个\n" +
			"例:\n" +
			"* 抢红包11.21元7个\n",
	})
	engine.OnPrefix("抢红包").SetBlock(true).Handle(handle)
}

func handle(ctx *robot.Ctx) {
	rawOptions := strings.Split(ctx.State["args"].(string), "元")
	if len(rawOptions) == 0 {
		return
	}

	fmt.Println(rawOptions)
	float, _ := strconv.ParseFloat(rawOptions[0], 10)

	a := int64(float * 100)
	numStrA := rawOptions[1]
	i := len(numStrA)
	fmt.Println(i)
	fmt.Println(numStrA)
	numStr := numStrA[:strings.Index(numStrA, "个")]
	fmt.Println(numStr)
	parseInt, _ := strconv.ParseInt(numStr, 10, 10)

	hb := computeHb(a, int(parseInt))

	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	ord := rd.Intn(int(parseInt))
	money := hb[ord]
	//toString, _ := jsoniter.MarshalToString(hb)v
	var toString = ""
	for _, v := range hb {
		toString += *MoneyStrConvert(v) + "元 "
	}

	console := "模拟抢红包\n"
	console += fmt.Sprintf("总金额:  %s 元\n", *MoneyStrConvert(a))
	console += fmt.Sprintf("红包个数:  %d 个\n", parseInt)
	console += fmt.Sprintf("你是第%d个抢到的\n", ord+1)
	console += fmt.Sprintf("你抢到了:  %s 元\n", *MoneyStrConvert(money))
	console += fmt.Sprintf("抢红包详情:\n%s\n", toString)

	err := ctx.ReplyTextAndAt(console)
	// 将结果放到匹配队列，触发其它插件
	if err == nil {
		//ctx.PushEvent(console)
	}
}

func computeHb(totalAmount int64, totalNum int) []int64 {
	rs := make([]int64, 0)
	rsAmount := totalAmount
	rsTotalNum := totalNum
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < totalNum-1; i++ {
		rda := (rsAmount/int64(rsTotalNum))*2 - 1
		rsOne := rd.Int63n(rda) + 1
		rsAmount -= rsOne
		rsTotalNum--
		rs = append(rs, rsOne)
	}
	rs = append(rs, rsAmount)
	return rs
}

func MoneyStrConvert(amount int64) *string {
	var rs string

	switch len(strconv.FormatInt(amount, 10)) {
	case 1: // 分
		geWei := amount % 10
		rs = fmt.Sprintf("0.0%d", geWei)
	case 2: // 毛
		geWei := amount % 10
		shiWei := amount / 10 % 10
		rs = fmt.Sprintf("0.%d%d", shiWei, geWei)
	case 3: // 元
		geWei := amount % 10
		shiWei := amount / 10 % 10
		baiWei := amount / 100 % 10
		rs = fmt.Sprintf("%d.%d%d", baiWei, shiWei, geWei)
	case 4: // 十元
		geWei := amount % 10
		shiWei := amount / 10 % 10
		baiWei := amount / 100 % 10
		qianWei := amount / 1000 % 10
		rs = fmt.Sprintf("%d%d.%d%d", qianWei, baiWei, shiWei, geWei)
	case 5: // 百元
		geWei := amount % 10
		shiWei := amount / 10 % 10
		baiWei := amount / 100 % 10
		qianWei := amount / 1000 % 10
		wanWei := amount / 10000 % 10
		rs = fmt.Sprintf("%d%d%d.%d%d", wanWei, qianWei, baiWei, shiWei, geWei)
	case 6: // 千元
		geWei := amount % 10
		shiWei := amount / 10 % 10
		baiWei := amount / 100 % 10
		qianWei := amount / 1000 % 10
		wanWei := amount / 10000 % 10
		shiWanWei := amount / 100000 % 10
		rs = fmt.Sprintf("%d%d%d%d.%d%d", shiWanWei, wanWei, qianWei, baiWei, shiWei, geWei)
	case 7: // 万元
		geWei := amount % 10
		shiWei := amount / 10 % 10
		baiWei := amount / 100 % 10
		qianWei := amount / 1000 % 10
		wanWei := amount / 10000 % 10
		shiWanWei := amount / 100000 % 10
		baiWanWanWei := amount / 1000000 % 10
		rs = fmt.Sprintf("%d%d%d%d%d.%d%d", baiWanWanWei, shiWanWei, wanWei, qianWei, baiWei, shiWei, geWei)
	default:
		rs = "0"
	}
	return &rs
}
