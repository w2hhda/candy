package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"errors"
	"github.com/w2hhda/candy/utils"
	"math/big"
)

type GameCandy struct {
	Id           int    `json:"id" orm:"column(id)"`
	GameFieldId  string `json:"game_field_id" orm:"column(game_field_id)"`
	CandyPool    int64  `json:"candy_pool" orm:"column(candy_pool)"`
	CandyConsume int64  `json:"candy_consume" orm:"column(candy_consume)"`
	Candy        *Candy `json:"candy" orm:"rel(one)"`
	Game         *Game  `json:"game" orm:"rel(one)"`
}

type Game struct {
	Id     int    `json:"id" orm:"column(id)"`
	Link   string `json:"url" orm:"column(link)"`
	Sort   int    `json:"sort" orm:"column(sort)"`
	Status int    `json:"status" orm:"column(status)"`
	Icon   string `json:"icon" orm:"column(icon)"`
	Name   string `json:"name" orm:"column(name)"`
}

type GameData struct {
	Request
	Name   string `json:"name"`
	Addr   string `json:"addr"`
	Type   int    `json:"candy_type"`
	Count  string `json:"count"`
	GameId int    `json:"game_id"`
}

type GamePlayer struct {
	Name  string `json:"name"`
	Addr  string `json:"addr"`
	Score string `json:"score"`
}

type GameRunningData struct {
	Request
	GameFieldId string       `json:"game_room_id"`
	GameId      int          `json:"game_id"`
	GamePlayer  []GamePlayer `json:"players"`
}

type GameStartData struct {
	CandyType  int    `json:"candy_type"`
	CandyCount string `json:"candy_count"`
	CandyLabel string `json:"candy_label"`
}

func GameCandyTableName() string {
	return "gamecandy"
}

func GameTableName() string {
	return "game"
}

func (this *Game) TableName() string {
	return GameTableName()
}

func (this *Game) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this.TableName())
}

func (this *Game) ListAllGame() ([]Game, error) {
	var list []Game
	_, err := this.Query().All(&list)
	return list, err
}

//1 . 分配糖果, 修改candy表和gamecandy表
//2 . 记录到record和gamecandy表中
func ReadOrInsert(normal, diamond Candy, request GameRunningData) (string, string, error) {
	o := orm.NewOrm()
	o.Begin()

	// 3. 分配糖果, 修改candy表和 gamecandy表
	// 普通糖果
	remainingCount, _ := new(big.Int).SetString(normal.RemainingCount, 10)
	normalCut := big.NewInt(int64(normal.Average * len(request.GamePlayer)))
	beego.Error("normalCut", normalCut)
	remainingCount = new(big.Int).Sub(remainingCount, normalCut)
	_, err := o.QueryTable(CandyTableName()).Filter("id", normal.Id).Update(orm.Params{
		"remaining_count": remainingCount.String(),
	})
	if err != nil {
		beego.Warn(err)
		o.Rollback()
		return "", "", err
	}
	// 钻石糖果
	var diamondCut *big.Int
	if diamond.Id > 0 {
		diamondCount, _ := new(big.Int).SetString(diamond.RemainingCount, 10)
		diamondCut = big.NewInt(int64(diamond.Average * 3))
		diamondCount = new(big.Int).Sub(diamondCount, diamondCut)
		_, err = o.QueryTable(CandyTableName()).Filter("id", diamond.Id).Update(orm.Params{
			"remaining_count": diamondCount.String(),
		})
		if err != nil {
			beego.Warn(err)
			o.Rollback()
			return "", "", err
		}
	}

	//4. 插入gamecandy
	var game Game
	err = o.QueryTable(GameTableName()).Filter("id", request.GameId).One(&game)
	if err != nil {
		beego.Warn(err)
		o.Rollback()
		return "", "", err
	}

	gameCandy := &GameCandy{
		GameFieldId:  request.GameFieldId,
		CandyPool:    remainingCount.Int64(),
		CandyConsume: 0,
		Candy:        &normal,
		Game:         &game,
	}

	_, err = o.Insert(gameCandy)
	if err != nil {
		beego.Warn(err)
		o.Rollback()
		return "", "", err
	}

	//插入record表
	for _, player := range request.GamePlayer {
		normalRecord := Record{
			Addr:     player.Addr,
			CreateAt: utils.GetTimestampString(),
			Count:    big.NewInt(0).String(),
			Candy:    &normal,
			Game: &Game{
				Id: request.GameId,
			},
			GameFieldId: request.GameFieldId,
		}

		if diamond.Id > 0 {
			diamondRecord := Record{
				Addr:     player.Addr,
				CreateAt: utils.GetTimestampString(),
				Count:    big.NewInt(0).String(),
				Candy:    &diamond,
				Game: &Game{
					Id: request.GameId,
				},
				GameFieldId: request.GameFieldId,
			}
			_, err = o.InsertMulti(2, []Record{normalRecord, diamondRecord})
			if err != nil {
				beego.Warn(err)
				o.Rollback()
				return "", "", err
			}
		} else {
			_, err = o.Insert(&normalRecord)
			if err != nil {
				beego.Warn(err)
				o.Rollback()
				return "", "", err
			}
		}

		if err != nil {
			beego.Warn(err)
			o.Rollback()
			return "", "", err
		}
	}

	o.Commit()
	return normalCut.String(), diamondCut.String(), nil
}

func RecordGameData(request GameData) error {
	o := orm.NewOrm()
	o.Begin()
	//1. 判断有没有这个用户
	exist := o.QueryTable(UserTableName()).Filter("addr", request.Addr).Exist()
	if !exist {
		// 2. 插入用户表 user 表
		user := User{
			Addr:     request.Addr,
			Status:   0,
			Name:     request.Name,
			CreateAt: utils.GetTimestampString(),
		}
		_, err := o.Insert(&user)
		if err != nil {
			beego.Warn(err)
			o.Rollback()
			return err
		}
	}

	// 3. 查询这个糖果的类型是不是正确的
	exist = o.QueryTable(CandyTableName()).Filter("candy_type", request.Type).Exist()
	if !exist {
		beego.Warn("输入糖果类型不正确")
		o.Rollback()
		return errors.New("输入糖果类型不正确")
	}

	// 4. 更新糖果表 token 表
	token := Token{}
	err := o.QueryTable(TokenTableName()).Filter("addr", request.Addr).
		Filter("candy_id", request.Type).One(&token)
	if err != nil {
		beego.Warn("查询token错误: ", err)
	}

	reqCount, _ := new(big.Int).SetString(request.Count, 10)
	tCount, _ := new(big.Int).SetString(token.Count, 10)
	beego.Info("update token", token)
	if token.Id > 0 {
		beego.Warn("token id", token.Id)
		count := new(big.Int).Add(tCount, reqCount)
		_, err := o.QueryTable(TokenTableName()).Filter("addr", request.Addr).
			Filter("candy_id", request.Type).Update(orm.Params{
			"count":     count.String(),
			"update_at": utils.GetTimestampString(),
		})
		if err != nil {
			beego.Warn(err)
			o.Rollback()
			return err
		}
	} else {
		token.Count = reqCount.String()
		token.Addr = request.Addr
		token.UpdateAt = utils.GetTimestampString()
		token.Candy = &Candy{
			CandyType: request.Type,
		}
		_, err := o.Insert(&token)
		if err != nil {
			beego.Warn(err)
			o.Rollback()
			return err
		}
	}

	//5. 插入记录表record
	exist = o.QueryTable(GameTableName()).Filter("id", request.GameId).Exist()
	if !exist {
		beego.Warn("非官方游戏")
		o.Rollback()
		return errors.New("非官方游戏")
	}

	record := Record{
		Addr:     request.Addr,
		CreateAt: utils.GetTimestampString(),
		Count:    request.Count,
		Candy: &Candy{
			CandyType: request.Type,
		},
		Game: &Game{
			Id: request.GameId,
		},
	}

	_, err = o.Insert(&record)
	if err != nil {
		beego.Warn(err)
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}
