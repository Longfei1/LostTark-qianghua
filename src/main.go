package main

import (
	"math/rand"
	"os"
	"qianghua/src/ui"
	"time"

	"github.com/flopp/go-findfont"
	"github.com/goki/freetype/truetype"
)

func init() {
	rand.Seed(time.Now().Unix())

	path, er := findfont.Find("simhei.ttf")
	if er != nil {
		return
	}

	fontData, er := os.ReadFile(path)
	if er != nil {
		return
	}
	_, er = truetype.Parse(fontData)
	if er != nil {
		return
	}
	os.Setenv("FYNE_FONT", path)
}

func main() {
	myApp := ui.NewApp()
	myApp.Init()

	myApp.Run()

	//cfg, er := ReadConfig()
	//if er != nil {
	//	fmt.Printf("read config er:%v", er.Error())
	//	return
	//}
	//
	//fmt.Println("---读取配置成功---")
	//
	////裸强
	//ex1 := CalcExpectation(cfg.BaseProbability, cfg.FailAddProbability, cfg.BaseFee, cfg.MaxCount)
	//
	////保护材料
	//ex2 := CalcExpectation(cfg.BaseProbability+cfg.ExtraProbability, cfg.FailAddProbability, cfg.BaseFee+cfg.ExtraFee, cfg.ExtraMaxCount)
	//
	//fmt.Printf("裸强期望：%v\n", ex1)
	//fmt.Printf("保护期望：%v\n", ex2)
	//
	//fmt.Println("\n输入指令 1：q （退出） 2.j 次数 （模拟强化n次装备）")
	//for {
	//	var s string
	//	var n int
	//	c, er := fmt.Scan(&s, &n)
	//	if er != nil || c == 0 {
	//		fmt.Println("指令识别错误")
	//		continue
	//	}
	//
	//	switch s {
	//	case "q":
	//		return
	//	case "j":
	//		r, f := SimulateN(cfg.BaseProbability, cfg.FailAddProbability, cfg.BaseFee, cfg.MaxCount, n)
	//		fmt.Println("       平均强化次数  平均费用")
	//		fmt.Printf("裸强    %1.2f        %f\n", r, f)
	//
	//		r, f = SimulateN(cfg.BaseProbability+cfg.ExtraProbability, cfg.FailAddProbability, cfg.BaseFee+cfg.ExtraFee, cfg.ExtraMaxCount, n)
	//		fmt.Printf("保护    %1.2f        %f\n", r, f)
	//
	//		fmt.Println("")
	//	}
	//}
}
