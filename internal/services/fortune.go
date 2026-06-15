package services

import (
	"fmt"
	"math/rand"
	"time"
)

type LotteryResult struct {
	Code          string `json:"code"`
	Title         string `json:"title"`
	Level         string `json:"level"`
	Poem          string `json:"poem"`
	Analysis      string `json:"analysis"`
	MasterReading string `json:"master_reading"`
	Advice        string `json:"advice"`
}

var signData = []struct {
	Level string
	Title string
	Poem  string
	Advice string
}{
	{"上上", "鲤鱼化龙", "一跃龙门身价高，春风得意马蹄忙。前程万里从此始，功名富贵自然来。\n\n此签为鲤鱼化龙之象，大吉大利。凡事谋望，皆能如愿。", "把握时机，大胆行动，必有所成。"},
	{"上吉", "喜鹊登枝", "喜鹊枝头叫喳喳，佳音传来福满家。所求之事皆如意，贵人相助事事佳。\n\n此签为喜鹊报喜之象，有好消息将至。", "耐心等待，好消息即将到来。"},
	{"上吉", "明月照江", "一轮明月照江明，万里波涛尽澄清。心地光明人自在，何愁前途不光明。\n\n此签为光明磊落之象，前途光明。", "保持正直，光明正大行事。"},
	{"中吉", "春风化雨", "春风化雨润无声，万物生长各有时。守得云开见月明，静待时机莫心急。\n\n此签为蓄势待发之象，需要耐心等待。", "目前时机未到，耐心等待为上策。"},
	{"中吉", "行路遇桥", "行路途中遇小桥，桥下流水向东流。过桥便是康庄道，一步一稳莫心焦。\n\n此签为逢凶化吉之象，虽有小阻，终能顺利。", "遇到困难不要急躁，稳扎稳打。"},
	{"中平", "水中捞月", "水中月影映天光，看得分明摸不着。不如回头寻正路，脚踏实地最稳当。\n\n此签为虚花之象，不可强求。", "不要执着于虚幻之事，务实行事。"},
	{"中平", "守株待兔", "株旁静坐待兔来，日暮西山兔未回。不如起身寻他处，天涯何处无芳草。\n\n此签为固守之象，需要灵活变通。", "换个思路，不要死守一个方向。"},
	{"下下", "逆水行舟", "逆水行舟用力撑，一篙松劲退千寻。如今识得风波险，稳舵前行莫暂停。\n\n此签为艰难之象，前进困难。", "目前形势不利，宜守不宜进。"},
	{"下下", "落花流水", "落花有意随流水，流水无心恋落花。缘来缘去终有时，莫为此事乱心怀。\n\n此签为失落之象，不必强求。", "顺其自然，该放手时就放手。"},
}

var masterReadings = map[string]func(string, string) string{
	"huiming": func(level, title string) string {
		return fmt.Sprintf("【慧明长老开示】\n\n此签为%s之象，签题「%s」。\n\n依《渊海子平》所载，此签主事须循序渐进，不可急功近利。施主当下之境，正如古人所言：欲速则不达。凡事稳中求进，自有福报。\n\n老衲劝施主：放下执念，随缘而行。心诚则灵，福报自来。", level, title)
	},
	"mingxin": func(level, title string) string {
		return fmt.Sprintf("【明心师父开示】\n\n施主求得此签，签题「%s」，签级%s。\n\n师父慈悲为怀，劝施主：世间万事皆有因缘，不必过于执着。当下所遇，皆是修行。保持一颗平常心，善待身边人，自然逢凶化吉。\n\n愿施主福慧双增，六时吉祥。🙏", title, level)
	},
	"xuanzhen": func(level, title string) string {
		return fmt.Sprintf("【玄真道长解签】\n\n嘿，这签有意思。「%s」，签级%s。\n\n直说了吧——这签的意思就是：%s。别想太多，该干嘛干嘛。命里有时终须有，命里无时莫强求。做好眼前事，其余的交给天意。\n\n道长我就说这么多，自己悟去吧。☯️", title, level, title)
	},
}

func DrawLottery(master string) *LotteryResult {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(signData))
	sign := signData[idx]

	result := &LotteryResult{
		Code:     fmt.Sprintf("第%d签", idx+1),
		Title:    sign.Title,
		Level:    sign.Level,
		Poem:     sign.Poem,
		Advice:   sign.Advice,
		Analysis: sign.Advice,
	}

	if reader, ok := masterReadings[master]; ok {
		result.MasterReading = reader(sign.Level, sign.Title)
	} else {
		result.MasterReading = fmt.Sprintf("此签为%s之象，签题「%s」。\n\n%s", sign.Level, sign.Title, sign.Advice)
	}

	return result
}

func BaziAnalyze(year, month, day int, shichen, gender, master string) map[string]interface{} {
	ganIdx := (year - 4) % 10
	zhiIdx := (year - 4) % 12
	shengXiao := []string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}
	tianGan := []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	diZhi := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

	yearGZ := tianGan[ganIdx] + diZhi[zhiIdx]
	monthGZ := tianGan[(ganIdx*2+month)%10] + diZhi[(month+1)%12]
	dayGZ := tianGan[(int(day)%10)] + diZhi[(int(day)%12)]

	wuxing := []string{"金", "木", "水", "火", "土"}
	wuxingMap := map[int]string{0: "金", 1: "木", 2: "水", 3: "火", 4: "土"}
	benMing := wuxingMap[ganIdx%5]

	personality := "您性格沉稳，做事踏实，有责任心。善于思考，注重细节，是一个可靠的人。"
	career := "事业方面，适合从事稳定型工作。中年后运势渐佳，贵人运旺。"
	wealth := "财运平稳，不宜冒险投资。积少成多，晚年丰足。"
	relationship := "感情方面较为专一，家庭观念强。宜寻属相相合之人。"
	health := "身体素质一般，注意劳逸结合。中年后注意保养。"

	if gender == "male" {
		personality = "您性格刚毅果断，有领导才能。做事有魄力，但有时过于急躁。"
	} else {
		personality = "您性格温婉细腻，善解人意。有艺术天赋，但有时过于敏感。"
	}

	result := map[string]interface{}{
		"bazi": map[string]string{
			"year":  yearGZ,
			"month": monthGZ,
			"day":   dayGZ,
			"shichen": shichen,
		},
		"sheng_xiao": shengXiao[zhiIdx],
		"ben_ming":   benMing,
		"wuxing":     []string{wuxing[ganIdx%5], wuxing[(ganIdx+1)%5], wuxing[(ganIdx+2)%5], wuxing[(ganIdx+3)%5], wuxing[(ganIdx+4)%5]},
		"personality": personality,
		"career":      career,
		"wealth":      wealth,
		"relationship": relationship,
		"health":      health,
	}
	return result
}

func InterpretDream(description string) map[string]interface{} {
	luckyLevels := []string{"吉", "凶", "平", "吉", "平"}
	rand.Seed(time.Now().UnixNano())
	level := luckyLevels[rand.Intn(len(luckyLevels))]

	interpretations := map[string]string{
		"蛇": "梦见蛇，主财运将至。蛇为小龙，象征蜕变与新生。",
		"水": "梦见水，主财源广进。水为财，清澈之水更好。",
		"火": "梦见火，主事业兴旺。火为阳，象征热情与动力。",
		"山": "梦见山，主有靠山。山为稳，象征踏实与安定。",
		"鱼": "梦见鱼，主有喜事。鱼为余，象征年年有余。",
		"花": "梦见花，主有桃花运。花开富贵，好事将至。",
		"树": "梦见树，主有贵人。大树底下好乘凉。",
		"路": "梦见路，主前程。路在脚下，方向由己。",
	}

	advice := "日有所思，夜有所梦。梦境反映内心所想，不必过于执着。"
	interpretation := "此梦反映您近期的心境状态。"

	for keyword, interp := range interpretations {
		if contains(description, keyword) {
			interpretation = interp
			break
		}
	}

	return map[string]interface{}{
		"interpretation": interpretation,
		"lucky_level":    level,
		"advice":         advice,
	}
}

func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && (s == substr || len(s) > len(substr) && (s[0:len(substr)] == substr || contains(s[1:], substr)))
}

func InterpretDivination(question string) map[string]interface{} {
	hexagrams := []string{"乾", "坤", "屯", "蒙", "需", "讼", "师", "比", "小畜", "履", "泰", "否", "同人", "大有", "谦", "豫", "随", "蛊", "临", "观", "噬嗑", "贲", "剥", "复", "无妄", "大畜", "颐", "大过", "坎", "离"}
	descriptions := []string{"元亨利贞", "元亨利牝马之贞", "元亨利贞勿用有攸往", "亨匪我求童蒙", "有孚光亨贞吉", "有孚窒惕中吉", "贞丈人吉无咎", "原筮元永贞无咎", "密云不雨自我西郊", "履虎尾不咥人亨", "小往大来吉亨", "否之匪人不利君子贞", "同人于野亨", "元亨", "亨君子有终", "利建侯行师", "元亨利贞无咎", "元亨利涉大川", "元亨利贞至于八月有凶", "盥而不荐有孚颙若", "亨利用狱", "亨小利有攸往", "不利有攸往", "亨出入无疾朋来无咎", "元亨利贞其匪正有眚", "利贞不家食吉", "贞吉观颐自求口实", "栋桡利有攸往亨", "习坎有孚维心亨行有尚", "利贞亨畜牝牛吉"}

	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(hexagrams))

	return map[string]interface{}{
		"hexagram":      hexagrams[idx],
		"description":   descriptions[idx],
		"interpretation": fmt.Sprintf("您所问之事，得「%s」卦，卦辞：%s。\n\n此卦象暗示当前形势需要审慎判断。建议您多方考量，不可贸然行事。若能保持冷静，终能化险为夷。", hexagrams[idx], descriptions[idx]),
		"advice":        "三思而后行，谨慎决策。",
	}
}
