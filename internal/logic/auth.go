package logic

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"prophet/internal/dao"
	"prophet/internal/model"
)

type AuthLogic struct {
	userDAO *dao.UserDAO
}

func NewAuthLogic() *AuthLogic {
	return &AuthLogic{userDAO: &dao.UserDAO{}}
}

func (l *AuthLogic) AnonymousInit(deviceID string) (*model.User, string, error) {
	user, err := l.userDAO.FindByDeviceID(deviceID)
	if err != nil {
		luckyCode := fmt.Sprintf("佛缘%04d", time.Now().UnixNano()%10000)
		user = &model.User{
			DeviceID:  deviceID,
			Nickname:  "善信" + luckyCode,
			LuckyCode: luckyCode,
		}
		if err := l.userDAO.Create(user); err != nil {
			return nil, "", err
		}
	}

	token := fmt.Sprintf("token_%d_%s", user.ID, generateID()[:16])
	return user, token, nil
}

func (l *AuthLogic) GetMe(userID interface{}) (*model.User, error) {
	return l.userDAO.FindByID(userID)
}

func generateID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func GenerateOrderNo() string {
	return fmt.Sprintf("ord_%s_%s", time.Now().Format("20060102150405"), generateID()[:8])
}
