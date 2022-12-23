package main

import (
	"fmt"
)

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName string
	biddedAmout int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours int
	hourlyRate int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmout
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func calculateNetIncome(ic []Income) {
	var netIncome int = 0
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netIncome += income.calculate()
	}
	fmt.Printf("Net income of organisation = $%d", netIncome)
}

func main() {
	project1 := FixedBilling{"project1", 5000}	
	project2 := FixedBilling{"project2", 10000}	
	project3 := TimeAndMaterial{"project3", 160, 25}
	project4 := Advertisement{"Banner", 2, 500}
	incomeStreams := []Income{project1, project2, project3, project4}
	calculateNetIncome(incomeStreams)
}

type Advertisement struct {
	adName string
	CPC int
	noOfClicks int
}

func (a Advertisement) calculate() int {
	return a.CPC * a.noOfClicks
}

func (a Advertisement) source() string {
	return a.adName
}