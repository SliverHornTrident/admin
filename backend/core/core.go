package core

import (
	"context"
	"fmt"
	"github.com/SliverHornTrident/shadow/global"
	"github.com/SliverHornTrident/shadow/interfaces"
)

func Register(cores ...interfaces.Corer) {
	ctx := context.Background()
	length := len(cores)
	{
		var has bool
		for i := 0; i < length; i++ {
			var ok bool
			_, ok = cores[i].(*_viper)
			if ok {
				has = ok
				err := Viper.Viper(global.Viper)
				if err != nil {
					Println(Viper, err)
					return
				}
				length--
				cores = append(cores[:i], cores[i+1:]...) // 移除viper核心实例
				Viper.cores = make(map[string]interfaces.Corer, length)
				for j := 0; j < length; j++ {
					if cores[j].ConfigName() != "" {
						Viper.cores[cores[j].ConfigName()] = cores[j]
					}
				}
				break
			}
		}
		if !has {
			panic("viper core is not registered")
		}
	} // 保证初始化第一个为viper
	{
		err := Viper.Initialization(ctx)
		if err != nil {
			Println(Viper, err)
		}
	} // viper初始化
	for i := 0; i < length; i++ {
		err := cores[i].Viper(global.Viper)
		if err != nil {
			Println(cores[i], err)
		}
		err = cores[i].Initialization(ctx)
		if err != nil {
			Println(cores[i], err)
		}
		fmt.Println(cores[i].Name(), "Initialization success.")
	}
}

func Println(core interfaces.Corer, err error) {
	message := fmt.Sprintf("%s=>%+v", core.Name(), err)
	if core.IsPanic() {
		panic(message)
	}
	fmt.Println(message)
}
