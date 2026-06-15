package services

import (
	"fmt"
	"math"
	"time"
)

type AlmanacData struct {
	Solar      SolarDate      `json:"solar"`
	Lunar      LunarDate      `json:"lunar"`
	Ganzhi     GanzhiDate     `json:"ganzhi"`
	Overall    OverallLevel   `json:"overall_level"`
	Jieqi      JieqiInfo      `json:"jieqi"`
	Yi         []string       `json:"yi"`
	Ji         []string       `json:"ji"`
	Shen       ShenSha        `json:"shen"`
	Chong      string         `json:"chong"`
	TaiPos     string         `json:"tai_position"`
	Xiu        string         `json:"xiu"`
	XiuLuck    string         `json:"xiu_luck"`
	Zhixing    string         `json:"zhixing"`
	Shichen    []ShichenInfo  `json:"shichen"`
}

type SolarDate struct {
	Year        int    `json:"year"`
	Month       int    `json:"month"`
	Day         int    `json:"day"`
	Weekday     int    `json:"weekday"`
	WeekdayFull string `json:"weekday_full"`
}

type LunarDate struct {
	YearInChinese   string `json:"year_in_chinese"`
	MonthInChinese  string `json:"month_in_chinese"`
	DayInChinese    string `json:"day_in_chinese"`
	YearZodiac      string `json:"year_zodiac"`
}

type GanzhiDate struct {
	Year  string `json:"year"`
	Month string `json:"month"`
	Day   string `json:"day"`
	Nayin string `json:"nayin"`
}

type OverallLevel struct {
	Level   string `json:"level"`
	Summary string `json:"summary"`
}

type JieqiInfo struct {
	Today string `json:"today"`
}

type ShenSha struct {
	Lucky   []string `json:"lucky"`
	Unlucky []string `json:"unlucky"`
}

type ShichenInfo struct {
	Name   string `json:"name"`
	Lucky  string `json:"lucky"`
	Ganzhi string `json:"ganzhi"`
	Chong  string `json:"chong"`
}

var (
	tianGan   = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	diZhi     = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	shengXiao = []string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}
	nayinList = []string{"海中金", "炉中火", "大林木", "路旁土", "剑锋金", "山头火", "涧下水", "城头土", "白蜡金", "杨柳木", "泉中水", "屋上土", "霹雳火", "松柏木", "长流水", "沙中金", "山下火", "平地木", "壁上土", "金箔金", "覆灯火", "天河水", "大驿土", "钗钏金", "桑柘木", "大溪水", "沙中土", "天上火", "石榴木", "大海水"}
	weekdays  = []string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"}

	lunarMonthName = []string{"正", "二", "三", "四", "五", "六", "七", "八", "九", "十", "冬", "腊"}
	lunarDayName   = []string{"初一", "初二", "初三", "初四", "初五", "初六", "初七", "初八", "初九", "初十", "十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九", "二十", "廿一", "廿二", "廿三", "廿四", "廿五", "廿六", "廿七", "廿八", "廿九", "三十"}
	lunarYearName  = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十", "十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九", "二十", "廿一", "廿二", "廿三", "廿四", "廿五", "廿六", "廿七", "廿八", "廿九", "三十"}
)

// 简化的农历计算（基于查表法）
func ComputeAlmanac(dateStr string) *AlmanacData {
	t, _ := time.Parse("2006-01-02", dateStr)
	data := &AlmanacData{}

	// 公历
	data.Solar = SolarDate{
		Year:        t.Year(),
		Month:       int(t.Month()),
		Day:         t.Day(),
		Weekday:     int(t.Weekday()),
		WeekdayFull: weekdays[t.Weekday()],
	}

	// 干支年
	yearGanIdx := (t.Year() - 4) % 10
	yearZhiIdx := (t.Year() - 4) % 12
	data.Ganzhi.Year = tianGan[yearGanIdx] + diZhi[yearZhiIdx]
	data.Lunar.YearZodiac = shengXiao[yearZhiIdx]

	// 干支月
	monthGanIdx := (yearGanIdx*2 + int(t.Month())) % 10
	monthZhiIdx := (int(t.Month()) + 1) % 12
	data.Ganzhi.Month = tianGan[monthGanIdx] + diZhi[monthZhiIdx]

	// 干支日（简化算法）
	dayOffset := t.Unix()/86400
	dayGanIdx := int((dayOffset + 49) % 10)
	dayZhiIdx := int((dayOffset + 49) % 12)
	data.Ganzhi.Day = tianGan[dayGanIdx] + diZhi[dayZhiIdx]

	// 纳音
	nayinIdx := (yearGanIdx*12 + yearZhiIdx) % 30
	if nayinIdx < 0 {
		nayinIdx += 30
	}
	data.Ganzhi.Nayin = nayinList[nayinIdx]

	// 农历（简化）
	lunarMonth := (int(t.Month()) + 10) % 12
	data.Lunar.MonthInChinese = lunarMonthName[lunarMonth] + "月"
	lunarDay := t.Day()
	if lunarDay <= 30 {
		data.Lunar.DayInChinese = lunarDayName[lunarDay-1]
	} else {
		data.Lunar.DayInChinese = "三十"
	}
	yearStr := fmt.Sprintf("%d", t.Year())
	data.Lunar.YearInChinese = ""
	for _, ch := range yearStr {
		idx := int(ch - '0')
		data.Lunar.YearInChinese += lunarYearName[idx]
	}

	// 节气
	data.Jieqi = computeJieqi(t)

	// 宜忌
	data.Yi, data.Ji = computeYiJi(dayGanIdx, dayZhiIdx, t)

	// 综合评级
	data.Overall = computeOverall(data.Yi, data.Ji)

	// 神煞
	data.Shen = computeShenSha(dayGanIdx, dayZhiIdx)

	// 冲煞
	data.Chong = fmt.Sprintf("%s日冲%s", data.Ganzhi.Day, diZhi[(dayZhiIdx+6)%12])

	// 胎神方位
	data.TaiPos = computeTaiPosition(dayZhiIdx)

	// 28宿
	data.Xiu, data.XiuLuck = computeXiu(t)

	// 12建除
	data.Zhixing = computeZhixing(dayZhiIdx)

	// 时辰
	data.Shichen = computeShichen(dayGanIdx)

	return data
}

func computeJieqi(t time.Time) JieqiInfo {
	jieqi := JieqiInfo{}
	month := int(t.Month())
	day := t.Day()

	jieqiMap := map[[2]int]string{
		{1, 5}: "小寒", {1, 20}: "大寒",
		{2, 3}: "立春", {2, 18}: "雨水",
		{3, 5}: "惊蛰", {3, 20}: "春分",
		{4, 4}: "清明", {4, 19}: "谷雨",
		{5, 5}: "立夏", {5, 20}: "小满",
		{6, 5}: "芒种", {6, 21}: "夏至",
		{7, 6}: "小暑", {7, 22}: "大暑",
		{8, 7}: "立秋", {8, 22}: "处暑",
		{9, 7}: "白露", {9, 22}: "秋分",
		{10, 8}: "寒露", {10, 23}: "霜降",
		{11, 7}: "立冬", {11, 22}: "小雪",
		{12, 6}: "大雪", {12, 21}: "冬至",
	}

	for k, v := range jieqiMap {
		if k[0] == month && math.Abs(float64(k[1]-day)) <= 1 {
			jieqi.Today = v
			break
		}
	}
	return jieqi
}

func computeYiJi(ganIdx, zhiIdx int, t time.Time) ([]string, []string) {
	yiPool := []string{
		"嫁娶", "出行", "搬家", "装修", "开业", "求职", "签约",
		"祈福", "求嗣", "纳采", "裁衣", "竖柱", "上梁", "栽种",
		"牧养", "纳畜", "开光", "塑绘", "斋醮", "出行", "会友",
	}
	jiPool := []string{
		"动土", "安葬", "破土", "开仓", "置产", "造船", "词讼",
		"栽种", "掘井", "置产", "行丧", "针灸", "远行", "搬迁",
	}

	daySeed := int(t.Unix()/86400) % 10
	yiCount := 4 + daySeed%3
	jiCount := 3 + daySeed%2

	yi := make([]string, 0, yiCount)
	ji := make([]string, 0, jiCount)
	usedYi := make(map[int]bool)
	usedJi := make(map[int]bool)

	for i := 0; i < yiCount; i++ {
		idx := (daySeed*3 + i*7) % len(yiPool)
		for usedYi[idx] {
			idx = (idx + 1) % len(yiPool)
		}
		usedYi[idx] = true
		yi = append(yi, yiPool[idx])
	}

	for i := 0; i < jiCount; i++ {
		idx := (daySeed*5 + i*3) % len(jiPool)
		for usedJi[idx] {
			idx = (idx + 1) % len(jiPool)
		}
		usedJi[idx] = true
		ji = append(ji, jiPool[idx])
	}

	return yi, ji
}

func computeOverall(yi, ji []string) OverallLevel {
	score := len(yi)*10 - len(ji)*8
	switch {
	case score >= 35:
		return OverallLevel{Level: "上上", Summary: "万事皆宜，大吉大利"}
	case score >= 25:
		return OverallLevel{Level: "上吉", Summary: "诸事顺遂，贵人相助"}
	case score >= 15:
		return OverallLevel{Level: "中吉", Summary: "平稳顺遂，小有收获"}
	case score >= 5:
		return OverallLevel{Level: "中平", Summary: "平常之日，谨慎行事"}
	default:
		return OverallLevel{Level: "下下", Summary: "诸事不宜，静待时机"}
	}
}

func computeShenSha(ganIdx, zhiIdx int) ShenSha {
	lucky := []string{"天德", "月德", "天恩", "天赦"}
	unlucky := []string{"天火", "月厌", "四废"}
	s := int(time.Now().Unix() / 86400)
	if s%3 == 0 {
		lucky = append(lucky, "文昌")
	}
	if s%5 == 0 {
		unlucky = append(unlucky, "五虚")
	}
	return ShenSha{Lucky: lucky, Unlucky: unlucky}
}

func computeTaiPosition(zhiIdx int) string {
	positions := []string{
		"占门碓 外东南", "占厨灶 外正南", "占房床 外正东",
		"占门炉 外正东", "占碓磨 外正南", "占厨灶 外正西",
		"占房床 外正西", "占碓磨 外正北", "占门厕 外正北",
		"占门鸡栖 外东南", "占碓磨 外正东", "占房床 外正南",
	}
	return positions[zhiIdx]
}

func computeXiu(t time.Time) (string, string) {
	xiuList := []string{"角", "亢", "氐", "房", "心", "尾", "箕", "斗", "牛", "女", "虚", "危", "室", "壁", "奎", "娄", "胃", "昴", "毕", "觜", "参", "井", "鬼", "柳", "星", "张", "翼", "轸"}
	xiuLuckList := []string{"吉", "凶", "平", "吉", "凶", "平", "吉", "凶", "平", "吉", "凶", "平", "吉", "凶", "平", "吉", "凶", "平", "吉", "凶", "平", "吉", "凶", "平", "吉", "凶", "平", "吉"}
	idx := int(t.Unix()/86400) % len(xiuList)
	return xiuList[idx], xiuLuckList[idx]
}

func computeZhixing(zhiIdx int) string {
	list := []string{"建", "除", "满", "平", "定", "执", "破", "危", "成", "收", "开", "闭"}
	return list[zhiIdx]
}

func computeShichen(dayGanIdx int) []ShichenInfo {
	shichenNames := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	luckPattern := []string{"吉", "凶", "吉", "凶", "吉", "凶", "吉", "凶", "吉", "凶", "吉", "凶"}

	result := make([]ShichenInfo, 12)
	for i := 0; i < 12; i++ {
		ganIdx := (dayGanIdx*12 + i) % 10
		zhiIdx := (dayGanIdx*12 + i) % 12
		result[i] = ShichenInfo{
			Name:   shichenNames[i],
			Lucky:  luckPattern[i],
			Ganzhi: tianGan[ganIdx] + diZhi[zhiIdx],
			Chong:  fmt.Sprintf("冲%s", diZhi[(zhiIdx+6)%12]),
		}
	}
	return result
}

func ComputeWeekAlmanac() []map[string]interface{} {
	now := time.Now()
	result := make([]map[string]interface{}, 7)
	weekdaysShort := []string{"一", "二", "三", "四", "五", "六", "日"}

	for i := 0; i < 7; i++ {
		d := now.AddDate(0, 0, i)
		data := ComputeAlmanac(d.Format("2006-01-02"))
		result[i] = map[string]interface{}{
			"date":      d.Format("2006-01-02"),
			"weekday":   weekdaysShort[(d.Weekday()+6)%7],
			"lunar_day": data.Lunar.DayInChinese,
			"level":     data.Overall.Level,
		}
	}
	return result
}
