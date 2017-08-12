package game

import (
	"image/color"

	"github.com/oakmound/oak/render"
)

func DemoBody() *Body {
	b := new(Body)
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

	b.overlay = render.NewColorBox(64, 64, color.RGBA{0, 255, 100, 255})
	b.veinColor = color.RGBA{255, 0, 0, 255}
	b.veinColor2 = color.RGBA{0, 0, 255, 255}
	b.AddNodes(NewLung(24, 25),
		NewLiver(27, 55),
		NewHeart(35, 8),
		NewRLung(34, 25))
	b.AddNodes(NewVeinNode(10, 40, b.veinColor),
		NewVeinNode(23, 5, b.veinColor),
		NewVeinNode(30, 38, b.veinColor),
		NewVeinNode(50, 22, b.veinColor))

	b.Connect(0, 5)
	b.Connect(0, 4)
	b.Connect(0, 6)
	b.Connect(2, 5)
	b.Connect(2, 7)
	b.Connect(3, 5)
	b.Connect(3, 7)
	b.Connect(1, 4)
	b.Connect(1, 7)
	b.Connect(6, 7)
	b.Connect(4, 5)
	//

	b.InfectionPattern([][]int{{0}, {1}})
	for _, oNum := range b.infectionPattern[b.infectionSet] {
		b.Infect(oNum)
	}

	b.InitVeins()
	return b
}

//Body2 infection of the (Liver Stomach) Heart Lung
func Body2() *Body {
	//(Liver Stomach) Heart Lung
	b := new(Body)

	b.overlay = render.NewColorBox(64, 64, color.RGBA{0, 255, 100, 255})
	b.veinColor = color.RGBA{255, 0, 0, 255}
	b.veinColor2 = color.RGBA{0, 0, 255, 255}
	b.AddNodes(NewLiver(24, 25),
		NewStomach(27, 55),
		NewHeart(35, 8),
		NewRLung(34, 25),
		NewLung(2, 2))
	b.AddNodes(NewVeinNode(10, 40, b.veinColor))

	b.InfectionPattern([][]int{{0, 1}, {2}, {3}})
	for _, oNum := range b.infectionPattern[b.infectionSet] {
		b.Infect(oNum)
	}

	b.InitVeins()
	return b
}

//Body3 infection of the (Stomach Heart Liver) Stomach
func Body3() *Body {
	b := new(Body)
	b.overlay = render.NewColorBox(64, 64, color.RGBA{0, 255, 100, 255})
	b.veinColor = color.RGBA{255, 0, 0, 255}
	b.veinColor2 = color.RGBA{0, 0, 255, 255}
	b.AddNodes(NewStomach(24, 25),
		NewHeart(27, 55),
		NewLiver(35, 8),
		NewStomach(34, 25))
	b.AddNodes(NewVeinNode(10, 40, b.veinColor))

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

	b.overlay = render.NewColorBox(64, 64, color.RGBA{0, 255, 100, 255})
	b.veinColor = color.RGBA{255, 0, 0, 255}
	b.veinColor2 = color.RGBA{0, 0, 255, 255}
	b.AddNodes(NewLung(24, 25),
		NewStomach(27, 55),
		NewBrain(35, 8),
		NewLiver(34, 25))
	b.AddNodes(NewVeinNode(10, 40, b.veinColor))

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

	b.overlay = render.NewColorBox(64, 64, color.RGBA{0, 255, 100, 255})
	b.veinColor = color.RGBA{255, 0, 0, 255}
	b.veinColor2 = color.RGBA{0, 0, 255, 255}
	b.AddNodes(NewLiver(24, 25),
		NewHeart(27, 55),
		NewStomach(35, 8),

		NewBrain(34, 25),
		NewLung(34, 25),
		NewRLung(34, 25),
		NewHeart(34, 25),
	)
	b.AddNodes(NewVeinNode(10, 40, b.veinColor))

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
