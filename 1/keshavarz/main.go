package main

import "fmt"

func main() {
	var w, sh int
	var p float64
	for {
		fmt.Scanf("%d %d", &w, &sh)
		if w >= 1 && w <= 100 && sh >= 1 && sh <= 100 {
			break
		} else {
			fmt.Println("your input must be between 1 and 100")
		}
	}
	for {
		fmt.Scanf("%f", &p)
		if p > 0 && p < 1 {
			break
		} else {
			fmt.Println(("your input must be between 0 and 1"))
		}
	}
	//fmt.Println(w, sh)
	//fmt.Println(p)

	attackDay := p * 365.0
	//fmt.Println(attackDay)
	annualSalary := w * 365
	//fmt.Println(annualSalary)
	annualDamage := float64(sh) * attackDay
	//fmt.Println(annualDamage)
	Pure := float64(annualSalary) - annualDamage
	fmt.Printf("%.2f\n", Pure)
}
