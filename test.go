package main

import (
	"fmt"
	"sync"
	"time"
)

func imprimirMenu() {
	fmt.Println("0. Salir")
	fmt.Println("1. Pausar cronometro")
	fmt.Println("2. Reanudar cronometro")
	fmt.Println("3. Detener cronometro")
	fmt.Println("4. Iniciar cronometro")
	fmt.Println("5. Detener por 5 segundos")
}

var cronometro1, cronometro2, cronometro3 Cronometro

func test() {
	// tick := cronometro1.Init(1 * time.Second)
	cronometro1.Init(1 * time.Second)
	cronometro2.Init(3 * time.Second)
	cronometro3.Init(6 * time.Second)

	var wg sync.WaitGroup
	wg.Add(3)

	// go func() {
	// 	for {
	// 		for segundo := range tick {
	// 			fmt.Println(segundo)
	// 		}
	// 	}
	// }()

	imprimirMenu()
	opcion := -1
	var cron_select *Cronometro
	for opcion != 0 {
		fmt.Print("> ")
		fmt.Scanf("%d", &opcion)

		switch opcion {
		case 1:
			cron_select = seleccionarCronometro()
			cron_select.Pausar()
		case 2:
			cron_select = seleccionarCronometro()
			cron_select.Reanudar()
		case 3:
			cron_select = seleccionarCronometro()
			cron_select.Detener()
		case 4:
			cron_select = seleccionarCronometro()
			cron_select.Iniciar()
		case 5:
			cronometro1.Pausar()
			cronometro2.Pausar()
			cronometro3.Pausar()
			time.Sleep(5 * time.Second)
			cronometro1.Reanudar()
			cronometro2.Reanudar()
			cronometro3.Reanudar()
		default:
			if opcion != 0 {
				imprimirMenu()
			} else {
				cronometro1.Detener()
				cronometro2.Detener()
				cronometro3.Detener()
				wg.Done()
				wg.Done()
				wg.Done()
			}
		}
	}
	wg.Wait()
}

func seleccionarCronometro() *Cronometro {
	var eleccion int
	var cron_select *Cronometro
	fmt.Print(">> ")
	fmt.Scanf("%d", &eleccion)
	switch eleccion {
	case 1:
		cron_select = &cronometro1
	case 2:
		cron_select = &cronometro2
	case 3:
		cron_select = &cronometro3
	}

	return cron_select
}
