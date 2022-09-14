package models

import (
	"fmt"
	"log"
	"strconv"
)

func SelectByGroup(groupType string) *[]SportInfo{
	var sportInfos []SportInfo
	// 预编译
	db.Where("group_type=?",groupType).Find(&sportInfos)
	for i := range  sportInfos{

		var minuteNum float64
		var err error
		if len(sportInfos[i].Meter800) != 0 {
			minuteNum,err = strconv.ParseFloat(sportInfos[i].Meter800,64)
			if err != nil {
				log.Fatal("parse fail!!")
			}
			minute := int(minuteNum) / 60
			second := int(minuteNum) % 60
			sportInfos[i].Meter800 = strconv.Itoa(minute) +"'"+ strconv.Itoa(second)
			fmt.Println(sportInfos[i].Meter800)

		}

		if len(sportInfos[i].Meter1000) != 0 {
			minuteNum,err = strconv.ParseFloat(sportInfos[i].Meter1000,64)
			if err != nil {
				log.Fatal("parse fail!!")
			}
			minute := int(minuteNum) / 60
			second := int(minuteNum) % 60
			sportInfos[i].Meter1000 = strconv.Itoa(minute) +"'"+ strconv.Itoa(second)
			fmt.Println(sportInfos[i].Meter1000)

		}

	}
	return &sportInfos

}
