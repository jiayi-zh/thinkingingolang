package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

const (
	addr     = "47.100.114.192:6379"
	password = "8421"
	DB       = 0
)

func Test_redisApi(t *testing.T) {
	rc := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       DB,
	})

	ctx := context.Background()

	key := "task"

	idArr := []int{5, 2, 1, 4}

	for _, id := range idArr {
		user := User{
			Id:   id,
			Name: fmt.Sprintf("user%d", id),
		}
		bytes, _ := json.Marshal(user)

		result := rc.ZAdd(ctx, key, &redis.Z{
			Score:  0,
			Member: string(bytes),
		})
		log.Infof("key[%s] zset add [%+v]", key, result)
	}

	countRes := rc.ZCount(ctx, key, "0", "5")
	log.Infof("count: %d", countRes.Val())

	userSlice := rc.ZRange(ctx, key, 0, 5)
	log.Infof("%v", userSlice.Val())

	setCmd := rc.Set(ctx, "k", 0, time.Minute)
	log.Infof("%v", setCmd)

	incrCmd := rc.IncrBy(ctx, "incrBy1", 0)
	log.Infof("%v", incrCmd)

	getCmd := rc.Get(ctx, "incrBy1")
	log.Infof("%v", getCmd)

	del := rc.Del(ctx, "111")
	if del.Err() != nil {

	}

	nxCmd := rc.SetNX(context.Background(), "knx", time.Now().Unix(), time.Minute*30)
	log.Infof("%v", nxCmd)
	nxCmd1 := rc.Get(context.Background(), "knx")
	log.Infof("%v", nxCmd1)
}

type User struct {
	Id   int
	Name string
}

type TaskProgress struct {
	Chunk  int   `json:"chunk"`
	Chunks int   `json:"chunks"`
	Fid    int64 `json:"fid"`
}

func (tp *TaskProgress) CalCurrentProgress() int {
	if tp.Chunks == 0 {
		return 0
	}
	return tp.Chunk * 100 / tp.Chunks
}

func Test_TaskProgress(t *testing.T) {
	tp := TaskProgress{
		Chunk:  31,
		Chunks: 33,
	}
	fmt.Println(tp.CalCurrentProgress())
}
