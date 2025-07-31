/*
 * @Author: WinstonRD
 * @Date: 2025-07-16 12:57:16
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-24 22:29:21
 * @FilePath: /34-go-redis/main.go
 * @Description: reids-go
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"34-redis-go/internal/config"
	"34-redis-go/internal/db"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type cacheClientKey string

func SetMethod(ctx context.Context) {
	cache := ctx.Value(cacheClientKey("cache")).(*redis.Client)
	err := cache.Set(ctx, "testkey", "go-redis", 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cache.Get(ctx, "testkey").Result())
}

// 更新一个值并同时取出旧值
func GetSetMethod(ctx context.Context) {
	cache := ctx.Value(cacheClientKey("cache")).(*redis.Client)
	str, err := cache.GetSet(ctx, "testkey", "redis-go").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}

// 为一个不存在的key设置值，如果存在则不更新
func SetNxMethod(ctx context.Context) {
	cache := ctx.Value(cacheClientKey("cache")).(*redis.Client)
	res, err := cache.SetNX(ctx, "testkey", "gogogo", 0).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

// 批量获取
func MGetMethod(ctx context.Context) {
	cache := ctx.Value(cacheClientKey("cache")).(*redis.Client)
	vals, err := cache.MGet(ctx, "testkey", "testkey1").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vals)
}

// 批量写入
func MSetMethod(ctx context.Context) {
	cache := ctx.Value(cacheClientKey("cache")).(*redis.Client)
	err := cache.MSet(ctx, map[string]interface{}{"testkey": "go-redis", "testkey1": "go-redis1"}).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func IncrMethod(ctx context.Context) {
	cache := ctx.Value(cacheClientKey("cache")).(*redis.Client)
	res, err := cache.IncrBy(ctx, "testintkey1", 2).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func DecrMethod(ctx context.Context) {
	cache := ctx.Value(cacheClientKey("cache")).(*redis.Client)
	res, err := cache.DecrBy(ctx, "testintkey1", 3).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func DelMethod(ctx context.Context) {
	cache := ctx.Value(cacheClientKey("cache")).(*redis.Client)
	res, err := cache.Del(ctx, "testkey", "testkey1").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func ExpireMethod(ctx context.Context) {
	cache := ctx.Value(cacheClientKey("cache")).(*redis.Client)
	res, err := cache.Expire(ctx, "testintkey", 5*time.Second).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func HSetMethod(r *redis.Client, ctx context.Context) {
	err := r.HSet(ctx, "User_1", "username", "recker").Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func HGetMethod(r *redis.Client, ctx context.Context) {
	res, err := r.HGetAll(ctx, "User_1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func HMSetMethod(r *redis.Client, ctx context.Context) {
	data := make(map[string]any)
	data["username"] = "winston"
	data["id"] = 1
	res, err := r.HMSet(ctx, "User_1", data).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func HExistsMethod(r *redis.Client, ctx context.Context) {
	res, err := r.HExists(ctx, "User_1", "idx").Result() // return bool, err
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

// RPush右侧写入
func RPush(r *redis.Client, ctx context.Context) {
	r.RPush(ctx, "recent_user", 1, 2, 3, 4, 5)

	len, _ := r.LLen(ctx, "recent_user").Result()
	fmt.Println(len)
}

func LRange(r *redis.Client, ctx context.Context) {
	// start 0 stop -1 取全部
	res, _ := r.LRange(ctx, "recent_user", 0, 2).Result()
	fmt.Println(res)
}

func LRem(r *redis.Client, ctx context.Context) {
	// 从列表左侧开始删除2个为1的元素
	r.LRem(ctx, "recent_user", 2, 1)
}

func LIndex(r *redis.Client, ctx context.Context) {
	res, _ := r.LIndex(ctx, "recent_user", 2).Result()
	fmt.Println(res)
}

func LInsert(r *redis.Client, ctx context.Context) {
	// 在元素为5的前面插入1
	r.LInsertBefore(ctx, "recent_user", 5, 1)
}

func SAdd(r *redis.Client, ctx context.Context) {
	// 写入集合，重复的不会再写入
	r.SAdd(ctx, "user_orders", 100, 100, 200, 300)

	// 获取元素个数
	len, _ := r.SCard(ctx, "user_orders").Result()
	fmt.Println(len)

	// 判断元素是否在集合中
	res, _ := r.SIsMember(ctx, "user_orders", 100).Result()
	fmt.Println(res)

	res1, _ := r.SMembers(ctx, "user_orders").Result()
	fmt.Println(res1)

	res2, _ := r.SMembersMap(ctx, "user_orders").Result()
	fmt.Println(res2)
}

func ZAdd(r *redis.Client, ctx context.Context) {
	data := make([]redis.Z, 0)
	data = append(
		data,
		redis.Z{Score: 14, Member: "user_1"},
		redis.Z{Score: 25, Member: "user_2"},
		redis.Z{Score: 44, Member: "user_3"},
		redis.Z{Score: 63, Member: "user_4"})
	r.ZAdd(ctx, "top_kill", data...)
}

func ZCount(r *redis.Client, ctx context.Context) {
	res, _ := r.ZCount(ctx, "top_kill", "10", "40").Result()
	fmt.Println(res)
}

func ZIncrBy(r *redis.Client, ctx context.Context) {
	r.ZIncrBy(ctx, "top_kill", 5, "user_1")
}

func ZRange(r *redis.Client, ctx context.Context) {
	// 查询
	var opt redis.ZRangeBy
	opt.Count = 3
	opt.Min = "10"
	opt.Max = "40"
	opt.Offset = 0
	res, _ := r.ZRangeByScoreWithScores(ctx, "top_kill", &opt).Result()
	fmt.Println(res)

	// 排序
	res1, _ := r.ZRevRange(ctx, "top_kill", 0, -1).Result()
	fmt.Println(res1)
}

func ZRank(r *redis.Client, ctx context.Context) {
	// 获取元素所在排位
	res, _ := r.ZRankWithScore(ctx, "top_kill", "user_3").Result()
	fmt.Println(res)
}

func Subscribe(r *redis.Client, ctx context.Context) {
	sub := r.Subscribe(ctx, "channel_1")
	for msg := range sub.Channel() {
		fmt.Println(msg.Channel, msg.Payload)
	}
}

func Publish(r *redis.Client, ctx context.Context) {
	res, _ := r.Publish(ctx, "channel_1", "message_1").Result()
	fmt.Println(res)
}

func Watch(r *redis.Client, ctx context.Context, wg *sync.WaitGroup) {
	fn := func(tx *redis.Tx) error {
		v, _ := tx.Get(ctx, "testintkey1").Int()
		v--
		_, err := tx.Pipelined(ctx, func(pipe redis.Pipeliner) error {
			// 在这里给key设置最新值
			fmt.Println("change to", v)
			pipe.Set(ctx, "testintkey1", v, 0)
			fmt.Println("set success!")
			return nil
		})
		return err
	}
	err := r.Watch(ctx, fn, "testintkey1")
	switch err {
	case nil:
		fmt.Println("success to set")
	case redis.TxFailedErr:
		fmt.Println("key has been changed")
	default:
		panic(err)
	}
	wg.Done()
}

type redLock interface {
	lock(context.Context)
	unlock(context.Context)
}

type redLockImpl struct {
	r   *redis.Client
	ttl time.Duration
	key string
}

func NewRedLock(r *redis.Client, ttl time.Duration, key string) (rl *redLockImpl) {
	rl = &redLockImpl{r: r, ttl: ttl, key: key}
	return
}

func (rl *redLockImpl) Lock(ctx context.Context) {
	for {
		ok, err := rl.r.SetNX(ctx, rl.key, 1, rl.ttl).Result()
		if err == nil && ok {
			fmt.Println("lock!")
			break
		}
		fmt.Println("wait lock!")
		time.Sleep(100 * time.Millisecond)
		continue
	}
}

func (rl *redLockImpl) Unlock(ctx context.Context) {
	rl.r.Del(ctx, rl.key)
}

func main() {
	conf := config.NewConfig("internal/etc/config.yaml")
	ctx := context.WithValue(context.Background(), config.CtxKey("conf"), conf)
	r := db.NewCache(ctx)
	l := NewRedLock(r, 3*time.Second, "redlock1")
	var wg sync.WaitGroup
	wg.Add(5)
	for range 5 {
		go func() {
			l.Lock(ctx)
			defer l.Unlock(ctx)
			time.Sleep(2 * time.Second)
			wg.Done()
		}()
	}
	wg.Wait()
	// hash表操作
	// HGetMethod(r, ctx)
	// HMSetMethod(r, ctx)
	// HExistsMethod(r, ctx)

	// 列表操作
	// RPush(r, ctx)
	// LRange(r, ctx)
	// LRem(r, ctx)
	// LIndex(r, ctx)
	// LInsert(r, ctx)

	// 集合操作
	// SAdd(r, ctx)

	// 有序集合
	// ZAdd(r, ctx)
	// ZCount(r, ctx)
	// ZIncrBy(r, ctx)
	// ZRange(r, ctx)
	// ZRank(r, ctx)
	// Subscribe(r, ctx)
	// var wg sync.WaitGroup
	// wg.Add(10)
	// for range 10 {
	// 	go Watch(r, ctx, &wg)
	// 	// time.Sleep(time.Second * 1)
	// }
	// wg.Wait()

	// hyperloglog 不重复的元素集合
	// res, _ := r.PFAdd(ctx, "UV-user_info", "user_1", "user_2", "user_3").Result()
	// fmt.Println(res)

	// res, _ = r.PFCount(ctx, "UV-user_info").Result()
	// fmt.Println(res)

	// bitmap
	// r.SetBit(ctx, "checkin", 0, 1)
	// r.SetBit(ctx, "checkin", 1, 1)
	// r.SetBit(ctx, "checkin", 2, 1)
	// r.SetBit(ctx, "checkin", 3, 1)
	// r.SetBit(ctx, "checkin", 4, 1)
	// r.SetBit(ctx, "checkin", 5, 0)
	// r.SetBit(ctx, "checkin", 6, 0)

	// res, _ := r.GetBit(ctx, "checkin", 2).Result()
	// fmt.Println(res)

	// var bitCount *redis.BitCount
	// res, _ = r.BitCount(ctx, "checkin", bitCount).Result()
	// fmt.Println(res)

	// redis分布式锁

}
