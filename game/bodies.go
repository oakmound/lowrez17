package game

import (
	"image/color"
	"time"

	"github.com/oakmound/oak/render"
)

func DemoBody() *Body {
	b := new(Body)
	b.level = -1
	b.startTime = time.Now()
	b.overlay = render.NewColorBox(64, 64, color.RGBA{0, 255, 100, 255})
	b.veinColor = color.RGBA{255, 0, 0, 255}
	b.veinColor2 = color.RGBA{0, 0, 255, 255}
	b.AddNodes(NewVeinNode(10, 10, b.veinColor),
		NewVeinNode(15, 20, b.veinColor),
		NewLiver(40, 5),
		NewHeart(50, 40),
		NewBrain(50, 10),
		NewLung(30, 30),
		NewStomach(10, 40))
	b.Connect(0, 1)
	b.Connect(0, 2)
	b.Connect(1, 2)
	b.Connect(1, 3)
	b.Connect(1, 4)
	b.Connect(1, 5)
	b.Connect(1, 6)

	//b.Infect(0)
	//b.Infect(1)
	//b.Infect(2)
	b.Infect(3)
	//b.Infect(4)
	//b.Infect(5)
	//b.Infect(6)
	b.InitVeins()
	return b
}

//Body1 infection of the Lung Liver
func Body1() *Body {
	b := new(Body)
	b.level = 0
	b.startTime = time.Now()

	b.overlay = render.NewColorBox(64, 64, color.RGBA{0, 255, 100, 255})
	b.veinColor = color.RGBA{255, 0, 0, 255}
	b.veinColor2 = color.RGBA{0, 0, 255, 255}
	b.AddNodes(NewLung(20, 27),
		NewLiver(20, 55),
		NewHeart(27, 40),
		NewRLung(38, 27),
		NewBrain(27, 10))
	b.AddNodes(NewVeinNode(27, 30, b.veinColor),
		NewVeinNode(30, 48, b.veinColor),
		NewVeinNode(16, 45, b.veinColor))

	b.Connect(0, 4)
	b.Connect(3, 4)

	b.Connect(5, 0)
	b.Connect(5, 3)

	b.Connect(5, 2)
	b.Connect(2, 1)

	b.Connect(0, 7)
	b.Connect(3, 6)

	b.Connect(1, 7)
	b.Connect(1, 6)

	b.Connect(2, 7)
	b.Connect(2, 6)

	b.InfectionPattern([][]int{{0}, {1}})
	for _, oNum := range b.infectionPattern[b.infectionSet] {
		b.Infect(oNum)
	}

	b.InitVeins()
	return b
}

//Body2 infection of the (Liver Stomach) Heart RLung
func Body2() *Body {
	//(Liver Stomach) Heart Lung
	b := new(Body)
	b.level = 1
	b.startTime = time.Now()

	b.overlay = render.NewColorBox(64, 64, color.RGBA{0, 255, 100, 255})
	b.veinColor = color.RGBA{255, 0, 0, 255}
	b.veinColor2 = color.RGBA{0, 0, 255, 255}
	b.AddNodes(NewLiver(19, 55),
		NewStomach(31, 40),
		NewHeart(24, 14),
		NewRLung(38, 27),
		NewLung(20, 27),
		NewBrain(28, 4))
	b.AddNodes(NewVeinNode(16, 41, b.veinColor),
		NewVeinNode(28, 23, b.veinColor),
		NewVeinNode(42, 53, b.veinColor),
		NewVeinNode(36, 14, b.veinColor))

	b.InfectionPattern([][]int{{0, 1}, {2}, {3}})
	for _, oNum := range b.infectionPattern[b.infectionSet] {
		b.Infect(oNum)
	}

	b.Connect(6, 0)
	b.Connect(6, 1)

	b.Connect(7, 3)
	b.Connect(7, 4)

	b.Connect(0, 1)
	b.Connect(1, 7)

	b.Connect(2, 3)
	b.Connect(2, 4)
	b.Connect(2, 7)
	b.Connect(2, 3)

	b.Connect(8, 6)
	b.Connect(8, 3)
	b.Connect(8, 1)

	b.Connect(5, 2)

	b.Connect(5, 9)
	b.Connect(3, 9)

	b.InitVeins()
	return b
}

//Body3 infection of the (Stomach Heart Liver) Stomach
func Body3() *Body {
	b := new(Body)
	b.level = 2
	b.startTime = time.Now()
	b.overlay = render.NewColorBox(64, 64, color.RGBA{0, 255, 100, 255})
	b.veinColor = color.RGBA{255, 0, 0, 255}
	b.veinColor2 = color.RGBA{0, 0, 255, 255}
	b.AddNodes(NewStomach(25, 45),
		NewHeart(26, 18),
		NewLiver(15, 28),
		NewStomach(39, 45),
		NewLiver(48, 28))
	b.AddNodes(NewVeinNode(36, 10, b.veinColor),
		NewVeinNode(33, 21, b.veinColor))

	b.Connect(0, 2)
	b.Connect(3, 4)
	b.Connect(1, 2)

	//b.Connect(2, 4)
	b.Connect(0, 3)

	//Vein Connection
	b.Connect(5, 4)
	b.Connect(5, 1)
	b.Connect(6, 2)
	b.Connect(6, 4)

	b.Connect(6, 5)

	b.InfectionPattern([][]int{{0, 1, 2}, {3}})
	for _, oNum := range b.infectionPattern[b.infectionSet] {
		b.Infect(oNum)
	}

	b.InitVeins()
	return b
}

//----------------------------
//Body4 infection of the  (Lung Stomach) (Brain Liver) Heart
func Body4() *Body {
	b := new(Body)
	b.level = 3
	b.startTime = time.Now()

	b.overlay = render.NewColorBox(64, 64, color.RGBA{0, 255, 100, 255})
	b.veinColor = color.RGBA{255, 0, 0, 255}
	b.veinColor2 = color.RGBA{0, 0, 255, 255}
	b.AddNodes(NewLung(8, 55),
		NewStomach(27, 30),
		NewBrain(20, 8),
		NewLiver(34, 55),
		NewBrain(35, 8),
		NewRLung(52, 55))
	b.AddNodes(NewVeinNode(28, 43, b.veinColor),
		NewVeinNode(14, 33, b.veinColor),
		NewVeinNode(46, 33, b.veinColor))

	b.Connect(2, 4)
	b.Connect(2, 1)
	b.Connect(4, 1)
	b.Connect(0, 3)
	b.Connect(3, 5)
	//VeinConnections
	b.Connect(7, 2)
	b.Connect(8, 4)
	b.Connect(6, 7)
	b.Connect(6, 8)
	b.Connect(6, 5)
	b.Connect(6, 0)
	b.Connect(0, 7)
	b.Connect(5, 8)

	b.InfectionPattern([][]int{{0, 1}, {2, 3}, {4}})
	for _, oNum := range b.infectionPattern[b.infectionSet] {
		b.Infect(oNum)
	}

	b.InitVeins()
	return b
}

//Body5 infection of the (Liver Heart Stomach) (Brain Lung Lung Heart) (Brain)
func Body5() *Body {
	b := new(Body)
	b.level = 4
	b.startTime = time.Now()

	b.overlay = render.NewColorBox(64, 64, color.RGBA{0, 255, 100, 255})
	b.veinColor = color.RGBA{255, 0, 0, 255}
	b.veinColor2 = color.RGBA{0, 0, 255, 255}
	b.AddNodes(NewLiver(18, 42),
		NewHeart(27, 25),
		NewStomach(37, 47),
		NewBrain(31, 6),
		NewLung(48, 28),
		NewRLung(55, 28),
		NewHeart(34, 25),
		NewLung(8, 28),
		NewRLung(15, 28),
	)
	b.AddNodes(NewVeinNode(46, 12, b.veinColor),
		NewVeinNode(12, 12, b.veinColor),
		NewVeinNode(25, 55, b.veinColor),
	)

	b.Connect(4, 5)
	b.Connect(1, 6)
	b.Connect(2, 0)
	b.Connect(0, 1)
	b.Connect(2, 6)

	//Vein connect
	b.Connect(9, 4)
	b.Connect(9, 5)
	b.Connect(9, 6)
	b.Connect(9, 3)

	b.Connect(10, 8)
	b.Connect(10, 7)
	b.Connect(10, 1)
	b.Connect(10, 3)

	b.Connect(11, 0)
	b.Connect(11, 2)

	b.InfectionPattern([][]int{{0, 1, 2}, {3, 4, 5, 6}, {3}})
	for _, oNum := range b.infectionPattern[b.infectionSet] {
		b.Infect(oNum)
	}
	b.InitVeins()
	return b
}

func GetBody(level string) *Body {
	if level == "endurance" {
		// randomize
		return nil
	}
	return levelBodies[level]
}
