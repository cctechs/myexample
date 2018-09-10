package main

import "fmt"

type Base interface {
	Draw() string
}

type Pen struct {
}

func (this *Pen) Draw() string {
	return "pen"
}

type Brush struct {
}

func (this *Brush) Draw() string {
	return "brush"
}

func Tesss(base Base) {
	fmt.Println(base.Draw())
}

func main() {
	//Tesss(&Pen{})
	//return
	fmt.Println("start")
	draws := make([]Base, 0)

	draws = append(draws, &Pen{})
	draws = append(draws, &Brush{})

	for _, v := range draws {
		fmt.Println(v.Draw())
	}
	fmt.Println("\nshit\n")

	fmt.Println(len(draws))
	for _, base := range draws {
		switch base.(type) {
		case *Pen:
			{
				fmt.Println(base.(*Pen).Draw())

			}
		case *Brush:
			{
				fmt.Println(base.(*Brush).Draw())
			}
		}
	}
}
