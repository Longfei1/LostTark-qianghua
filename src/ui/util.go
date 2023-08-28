package ui

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func NewEntryWithLabel(name string, value interface{}, onSubmitted func(string)) fyne.CanvasObject {
	label := widget.NewLabel(name)

	var data binding.String
	switch v := value.(type) {
	case *float64:
		data = binding.FloatToString(binding.BindFloat(v))
	case *string:
		data = binding.BindString(v)
	case *int:
		data = binding.IntToString(binding.BindInt(v))
	case binding.String:
		data = v
	case binding.Int:
		data = binding.IntToString(v)
	case binding.Float:
		data = binding.FloatToString(v)
	default:

	}

	var input *widget.Entry
	if data != nil {
		input = widget.NewEntryWithData(data)
	} else {
		input = widget.NewEntry()
		if v, ok := value.(string); ok {
			input.SetText(v)
		}
	}
	input.Wrapping = fyne.TextWrapOff

	return container.NewGridWithColumns(2, label, input)
}

func CalcExpectation(p, fp, fee float64, maxCount int) float64 {
	var ex float64
	lastP := float64(1)

	for i := 1; i <= maxCount; i++ {
		feeNow := float64(i) * fee
		pNow := p + float64(i-1)*fp

		if pNow >= float64(1) {
			ex += lastP * feeNow
			break
		}

		if i == maxCount {
			ex += lastP * feeNow
		} else {
			ex += lastP * pNow * feeNow
			lastP = lastP * (float64(1) - pNow)
		}
	}

	return ex
}

func Simulate(p, fp, fee float64, maxCount int) ([]bool, float64) {
	n := 1
	ret := []bool{}
	for {
		pNow := p + float64(n-1)*fp
		if n >= maxCount || rand.Float64() < pNow {
			ret = append(ret, true)
			break
		}

		ret = append(ret, false)
		n++
	}

	return ret, float64(n) * fee
}

func SimulateN(p, fp, fee float64, maxCount, count int) (float64, float64) {
	tFee := float64(0)
	tCount := 0

	for i := 0; i < count; i++ {
		r, f := Simulate(p, fp, fee, maxCount)
		tCount += len(r)
		tFee += f

		//PrintSimulate(r, f)
	}

	return float64(tCount) / float64(count), tFee / float64(count)
}

func PrintSimulate(success []bool, fee float64) {
	rm := map[bool]string{
		false: "失败",
		true:  "成功",
	}
	s := ""
	for i, v := range success {
		s += fmt.Sprintf("%v：%v ", i+1, rm[v])
	}

	s += fmt.Sprintf("总费用：%v", fee)
	fmt.Println(s)
}
