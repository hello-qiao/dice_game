package dice_game

import (
	"math/rand"
	"time"
)

var game GameData

type GameData struct {
	Values       []GameLocation // 游戏的路径地图
	GameId       int            // 游戏id
	UserPathList map[int][]int  // 用户的路径
	Status       bool           // true 表示游戏结束
}

type GameLocation struct {
	Val        int           // 当前的位置
	NextAction *GameLocation // 遇到的动作，梯子或者蛇

}

// 这里封装一个新建游戏的接口，传入用户id之类信息，返回游戏id
func NewGameMap(uidList []int) int {
	values := []GameLocation{}
	for i := 1; i < 101; i++ {
		value := GameLocation{
			Val:        i,
			NextAction: nil,
		}
		values = append(values, value)
	}
	// 随机增加梯子或蛇
	for i := 0; i < 10; i++ {
		// 找到随机位置
		randAction := rand.Intn(100)
		// 此位置增加下一跳，可能是梯子或蛇
		values[randAction].NextAction = &values[rand.Intn(100)]
	}
	gameId := time.Now().Second()
	userPathList := make(map[int][]int)
	for _, u := range uidList {
		userPathList[u] = []int{0}
	}
	game = GameData{Values: values, GameId: gameId, UserPathList: userPathList}
	return gameId
}

// 用户轮流调用下一步的方法，返回路径图到前端
func Next(gameId int, uid int) GameData {
	if game.GameId != gameId {
		return GameData{}
	}
	if game.Status {
		return game
	}
	step := rand.Intn(6) + 1
	uidPath := game.UserPathList[uid]
	next := uidPath[len(uidPath)-1] + step
	// 超过上限
	if next > len(game.Values) {
		next = len(game.Values)*2 - next
	}
	if game.Values[next].NextAction != nil {
		next = game.Values[next].NextAction.Val - 1
	}
	game.UserPathList[uid] = append(uidPath, next)
	if next == len(game.Values) {
		game.Status = true
		// 做个写数据库操作，等待记录回放
	}
	return game
}

func GetLog(gameId int) GameData {
	// 实际上从数据库获取并返回
	if game.GameId == gameId {
		return game
	}
	return GameData{}
}

func main() {

}
