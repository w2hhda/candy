package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"github.com/astaxie/beego"
	"strconv"
	"errors"
	"github.com/w2hhda/candy/utils"
)

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
	Count  int    `json:"count"`
	GameId int    `json:"game_id"`
}

func GameTableName() string {
	return "game"
}

func (this *Game) TableName() string {
	return "game"
}

func (this *Game) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this.TableName())
}

func (this *Game) ListAllGame() ([]Game, error) {
	var list []Game
	_, err := this.Query().All(&list)
	return list, err
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
	err := o.QueryTable(UserCandyTableName()).Filter("addr", request.Addr).
		Filter("candy_id", request.Type).One(&token)
	if err != nil {
		beego.Warn("查询token错误: ", err)
	}

	beego.Info("update token", token)
	countStr := strconv.Itoa(request.Count)
	if token.Id > 0 { //数据转换
		count := parseFloat(token.Count) + float64(request.Count)
		token.Count = parseString(count)
		token.UpdateAt = time.Now().String()
		_, err := o.Update(&token)
		if err != nil {
			beego.Warn(err)
			o.Rollback()
			return err
		}
	} else {
		token.Count = countStr
		token.Addr = request.Addr
		token.UpdateAt = time.Now().String()
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
		return err
	}

	record := Record{
		Addr:     request.Addr,
		CreateAt: utils.GetTimestampString(),
		Count:    countStr,
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
