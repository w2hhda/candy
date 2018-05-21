package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"errors"
	"github.com/w2hhda/candy/utils"
	"math/big"
	"strconv"
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

func InsertNoExistUser(o orm.Ormer, players []GamePlayer) error {
	for _, player := range players {
		user := User{
			Addr:     player.Addr,
			CreateAt: utils.GetTimestampString(),
			Name:     player.Name,
			Status:   0,
		}
		b, i, err := o.ReadOrCreate(&user, "addr")
		beego.Info("InsertNoExistUser", b, i, err)
		if err != nil {
			return err
		}
	}
	return nil

}

func DispatchCandy(o orm.Ormer, game *Game, candy *Candy, request *GameRunningData, len int) (string, error) {
	remainingCount, _ := new(big.Int).SetString(candy.RemainingCount, 10)
	dispatchCount := big.NewInt(int64(candy.Average * len))
	remainingCount = new(big.Int).Sub(remainingCount, dispatchCount)
	_, err := o.QueryTable(CandyTableName()).Filter("id", candy.Id).Update(orm.Params{
		"remaining_count": remainingCount.String(),
	})
	if err != nil {
		beego.Warn(err)
		return "", err
	}
	err = InsertNoExistGameCandy(o, game, candy, request.GameFieldId, dispatchCount.Int64())
	if err != nil {
		beego.Warn(err)
		return "", err
	}
	return dispatchCount.String(), nil
}

func InsertNoExistGameCandy(o orm.Ormer, game *Game, candy *Candy, gameFieldId string, candyPool int64) error {
	gameCandy := GameCandy{
		GameFieldId:  gameFieldId,
		CandyPool:    candyPool,
		CandyConsume: 0,
		Candy:        candy,
		Game:         game,
	}

	//是否记录已经创建过了
	err := o.Read(&gameCandy, "game_field_id", "game_id", "candy_id")
	if err != nil && err != orm.ErrNoRows {
		beego.Warn(err)
		return err
	}

	if gameCandy.Id > 0 {
		beego.Warn("游戏已经创健了")
		return errors.New("游戏已经创健了")
	}

	_, err = o.Insert(&gameCandy)
	if err != nil {
		beego.Warn(err)
		return err
	}

	return nil
}

//1 . 分配糖果, 修改candy表和gamecandy表
//2 . 记录到record和gamecandy表中
func ReadOrInsert(normal, diamond Candy, request GameRunningData) (string, string, error) {
	o := orm.NewOrm()
	o.Begin()

	//1. 判断有没有这个用户
	err := InsertNoExistUser(o, request.GamePlayer)
	if err != nil {
		beego.Warn(err)
		o.Rollback()
		return "", "", err
	}

	//3. 插入gamecandy
	var game Game
	err = o.QueryTable(GameTableName()).Filter("id", request.GameId).One(&game)
	if err != nil {
		beego.Warn(err)
		o.Rollback()
		return "", "", err
	}

	// 2. 分配糖果, 修改candy表和 gamecandy表
	// 普通糖果
	normalDispatchCount, err := DispatchCandy(o, &game, &normal, &request, len(request.GamePlayer))
	if err != nil {
		beego.Warn(err)
		o.Rollback()
		return "", "", err
	}
	// 钻石糖果
	var diamondDispatchCount string
	if diamond.Id > 0 {
		drc, err := DispatchCandy(o, &game, &diamond, &request, DIAMOND_LEN)
		diamondDispatchCount = drc
		if err != nil {
			beego.Warn(err)
			o.Rollback()
			return "", "", err
		}
	}

	//4. 插入record表
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
	return normalDispatchCount, diamondDispatchCount, nil
}

func UpdateGameData(request GameRunningData) error {
	o := orm.NewOrm()
	o.Begin()
	//1. 查询game_candy表
	gameCandy := GameCandy{
		GameFieldId: request.GameFieldId,
		Game: &Game{
			Id: request.GameId,
		},
	}
	err := orm.NewOrm().Read(&gameCandy, "game_field_id", "game_id")
	if err != nil && err != orm.ErrNoRows {
		beego.Warn(err)
		o.Rollback()
		return err
	}

	// 存在游戏记录
	if gameCandy.Id < 0 {
		beego.Warn("没有分配记录")
		o.Rollback()
		return errors.New("没有找到游戏记录")
	}

	// 得分不能大于分配的数量
	var allScore int64
	for _, player := range request.GamePlayer {

		score, _ := strconv.ParseInt(player.Score, 10, 64)
		allScore += score

		if score > gameCandy.CandyPool {
			beego.Warn("得到的糖果不对")
			o.Rollback()
			return errors.New("得到的糖果不对")
		}

		// 查询record表
		record := Record{
			Addr:        player.Addr,
			GameFieldId: request.GameFieldId,
			Game:        gameCandy.Game,
		}
		err := o.Read(&record, "addr", "game_field_id", "game_id")
		if err != nil {
			beego.Warn("账单中没有记录", err)
			o.Rollback()
			return errors.New("账单中没有记录")
		}

		// 判断是否已经加过了，防止重复提交
		count, _ := strconv.ParseInt(record.Count, 10, 64)
		if count > 0 {
			beego.Warn("糖果已经加过了", err)
			o.Rollback()
			return errors.New("糖果已经加过了")
		}

		record.Count = strconv.FormatInt(score+count, 10)
		o.Update(&record, "count")

		//更新token表
		token := Token{
			Addr:  player.Addr,
			Candy: record.Candy,
		}

		err = o.Read(&token, "addr", "candy_id")
		if err != nil && err != orm.ErrNoRows {
			beego.Warn("查询个人糖果失败", err)
			o.Rollback()
			return errors.New("查询个人糖果失败")
		}

		token.UpdateAt = utils.GetTimestampString()
		if token.Id > 0 {
			tCount, _ := new(big.Int).SetString(token.Count, 10)
			rCount := new(big.Int).And(tCount, big.NewInt(score))
			token.Count = rCount.String()
			_, err = o.Update(&token, "count", "update_at")
			if err != nil {
				beego.Warn("更新个人糖果失败", err)
				o.Rollback()
				return errors.New("更新个人糖果失败")
			}
		} else {
			token.Count = big.NewInt(score).String()
			_, err = o.Insert(&token)
			if err != nil {
				beego.Warn("插入个人糖果失败", err)
				o.Rollback()
				return errors.New("插入个人糖果失败")
			}
		}

	}
	//更新game_candy
	if gameCandy.CandyPool < allScore {
		beego.Warn("插入个人糖果大于分配的数量", err)
		o.Rollback()
		return errors.New("插入个人糖果大于分配的数量")
	}

	gameCandy.CandyConsume = allScore
	_, err = o.Update(&gameCandy, "candy_consume")
	if err != nil {
		beego.Warn("更新游戏糖果表失败", err)
		o.Rollback()
		return errors.New("更新游戏糖果表失败")
	}

	// 更新candy表
	err = o.Read(gameCandy.Candy)
	if err != nil && err != orm.ErrNoRows {
		beego.Warn("查询糖果表失败", err)
		o.Rollback()
		return errors.New("查询糖果表失败")
	}
	reCount, _ := new(big.Int).SetString(gameCandy.Candy.RemainingCount, 10)
	reCount = reCount.And(reCount, big.NewInt(gameCandy.CandyPool-allScore))
	gameCandy.Candy.RemainingCount = reCount.String()

	_, err = o.Update(gameCandy.Candy, "remaining_count")
	if err != nil {
		beego.Warn("更新糖果表失败", err)
		o.Rollback()
		return errors.New("更新糖果表失败")
	}

	o.Commit()
	return nil
}
