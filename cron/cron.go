package cron

import (
	"CarLibrary/score/model"
	"fmt"
	"github.com/robfig/cron/v3"
)

func RefreshScore()  {
	model.GetAVRScore()
}

func FreshAvg (){
	c := cron.New()
	spec := "@every 5m"
	c.AddFunc(spec, func() {
		fmt.Println("hello")
		//todo 将平均分 注入 车系表中


	})
	c.Run()

}
