package ui

import (
	"fmt"
	"qianghua/src/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type MyApp struct {
	uiApp  fyne.App
	frameW fyne.Window
	frameC fyne.CanvasObject

	cfgC fyne.CanvasObject

	opeC fyne.CanvasObject

	outC    *container.Scroll
	outList *fyne.Container

	cfgData *data.Config
	init    bool
}

func NewApp() *MyApp {
	return &MyApp{
		uiApp: app.NewWithID("longfei.qianghua"),
	}
}

func (a *MyApp) Init() {
	cfgData, er := data.ReadConfig()
	if er != nil {
		panic(er.Error())
	}
	a.cfgData = cfgData

	a.createFrame()

	a.init = true
}

func (a *MyApp) createFrame() {
	a.frameW = a.uiApp.NewWindow("强化模拟工具")
	a.frameW.SetMaster()

	//图标
	logo, _ := fyne.LoadResourceFromPath("../resource/icon.png")
	a.frameW.SetIcon(logo)

	a.createCfgBlock()
	a.createOpeBlock()
	a.createOutBlock()

	//主布局
	a.frameC = container.NewVSplit(a.outC, container.NewVBox(a.cfgC, widget.NewSeparator(), a.opeC))
	a.frameW.SetContent(a.frameC)

	a.frameW.SetOnClosed(func() {
		a.saveData()
	})

	a.frameW.Resize(fyne.NewSize(1024, 768))
}

func (a *MyApp) createCfgBlock() {
	a.cfgC = container.NewGridWithColumns(3,
		NewEntryWithLabel("基础强化成功概率：", &a.cfgData.BaseProbability, nil),
		NewEntryWithLabel("基础费用：", &a.cfgData.BaseFee, nil),
		NewEntryWithLabel("最大强化次数（保底）：", &a.cfgData.MaxCount, nil),
		NewEntryWithLabel("额外增加强化概率：", &a.cfgData.ExtraProbability, nil),
		NewEntryWithLabel("额外所需费用：", &a.cfgData.ExtraFee, nil),
		NewEntryWithLabel("额外最大强化次数（保底）：", &a.cfgData.ExtraMaxCount, nil),
		NewEntryWithLabel("失败增加概率：", &a.cfgData.FailAddProbability, nil),
	)
}

func (a *MyApp) saveData() {
	fmt.Println("save cfg data", a.cfgData)

	er := data.SaveConfig(a.cfgData)
	if er != nil {
		fmt.Println(er.Error())
	}
}

func (a *MyApp) createOpeBlock() {
	a.opeC = container.NewHBox(
		NewEntryWithLabel("模拟次数：", binding.BindPreferenceInt("simulateCount", a.uiApp.Preferences()), nil),
		widget.NewButton("计算期望", a.onBtnCalcExpectation),
		widget.NewButton("模拟强化", a.onBtnSimulate),
		layout.NewSpacer(),
		widget.NewButton("清空输出", a.onBtnClear),
	)
}

func (a *MyApp) onBtnCalcExpectation() {
	//裸强
	ex1 := CalcExpectation(a.cfgData.BaseProbability, a.cfgData.FailAddProbability, a.cfgData.BaseFee, a.cfgData.MaxCount)

	//保护材料
	ex2 := CalcExpectation(a.cfgData.BaseProbability+a.cfgData.ExtraProbability, a.cfgData.FailAddProbability,
		a.cfgData.BaseFee+a.cfgData.ExtraFee, a.cfgData.ExtraMaxCount)

	a.addOutString("裸强期望：%v\n保护期望：%v", ex1, ex2)
}

func (a *MyApp) onBtnSimulate() {
	count := a.uiApp.Preferences().Int("simulateCount")
	if count <= 0 {
		return
	}

	//裸强
	c1, f1 := SimulateN(a.cfgData.BaseProbability, a.cfgData.FailAddProbability, a.cfgData.BaseFee, a.cfgData.MaxCount, count)

	//保护材料
	c2, f2 := SimulateN(a.cfgData.BaseProbability+a.cfgData.ExtraProbability, a.cfgData.FailAddProbability,
		a.cfgData.BaseFee+a.cfgData.ExtraFee, a.cfgData.ExtraMaxCount, count)

	a.addOutString("        平均次数    平均费用\n裸强：  %1.2f        %.2f\n保护：  %1.2f        %.2f",
		c1, f1, c2, f2)
}

func (a *MyApp) onBtnClear() {
	a.outList.RemoveAll()
}

func (a *MyApp) createOutBlock() {
	a.outList = container.NewVBox()

	sl := container.NewVScroll(a.outList)
	//sl.SetMinSize(fyne.NewSize(1024, 590))
	a.outC = sl
}

func (a *MyApp) Run() {
	a.frameW.ShowAndRun()
}

func (a *MyApp) addOutString(s string, args ...interface{}) {
	item := widget.NewLabel(fmt.Sprintf(s, args...))

	a.outList.Add(item)

	a.outC.ScrollToBottom()
	fmt.Println(a.outC.Size())
}
