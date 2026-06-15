package logic

import (
	"prophet/internal/dao"
)

type MeditationLogic struct {
	meditationDAO *dao.MeditationDAO
}

func NewMeditationLogic() *MeditationLogic {
	return &MeditationLogic{meditationDAO: &dao.MeditationDAO{}}
}

func (l *MeditationLogic) Catalog() (tracks interface{}, guided []map[string]interface{}, quote map[string]string) {
	tracks, _ = l.meditationDAO.FindActiveTracks()

	guided = []map[string]interface{}{
		{"id": "10min", "title": "十分钟入门", "subtitle": "适合初学者", "duration": 600,
			"steps": []string{"盘腿端坐，背挺直", "深呼吸三次，吸气数4秒，呼气数6秒", "把注意力放在鼻尖呼吸的进出", "杂念升起时不评判，温柔回到呼吸", "结束时双手合掌，回向众生"}},
		{"id": "20min", "title": "二十分钟正念", "subtitle": "进阶练习", "duration": 1200,
			"steps": []string{"三下吐纳调息", "观呼吸：注意力锁定鼻尖出入气", "扫描身体：从头顶到脚趾，依次放松每一处", "观念头来去：见妄念升起即知见，不跟随", "回向：愿一切众生离苦得乐"}},
	}

	quote = map[string]string{
		"text":   "不忘初心，方得始终",
		"source": "《华严经》",
	}
	return
}
