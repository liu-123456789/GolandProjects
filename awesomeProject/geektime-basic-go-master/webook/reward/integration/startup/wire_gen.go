// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package startup

import (
	"gitee.com/geekbang/basic-go/webook/api/proto/gen/payment/v1"
	"gitee.com/geekbang/basic-go/webook/reward/repository"
	"gitee.com/geekbang/basic-go/webook/reward/repository/cache"
	"gitee.com/geekbang/basic-go/webook/reward/repository/dao"
	"gitee.com/geekbang/basic-go/webook/reward/service"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitWechatNativeSvc(client pmtv1.WechatPaymentServiceClient) *service.WechatNativeRewardService {
	gormDB := InitTestDB()
	rewardDAO := dao.NewRewardGORMDAO(gormDB)
	cmdable := InitRedis()
	rewardCache := cache.NewRewardRedisCache(cmdable)
	rewardRepository := repository.NewRewardRepository(rewardDAO, rewardCache)
	loggerV1 := InitLogger()
	wechatNativeRewardService := service.NewWechatNativeRewardService(client, rewardRepository, loggerV1)
	return wechatNativeRewardService
}

// wire.go:

var thirdPartySet = wire.NewSet(InitTestDB, InitLogger, InitRedis)
