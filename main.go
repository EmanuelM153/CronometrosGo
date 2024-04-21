package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"runtime"
	"time"
)

func ponerTiempo(label *widget.Label, segundos time.Duration) {
	label.SetText(fmt.Sprintf("Tiempo: %v segundos", segundos))
}

func updateLabel(label *widget.Label, tick chan time.Duration) {
	for {
		select {
		case segundos := <-tick:
			ponerTiempo(label, segundos)
		default:
			runtime.Gosched()
		}
	}
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Cronómetros")

	cron1TiempoLbl := widget.NewLabel("Tiempo: 0s segundos")
	cron2TiempoLbl := widget.NewLabel("Tiempo: 0s segundos")
	cron3TiempoLbl := widget.NewLabel("Tiempo: 0s segundos")

	var cronometro1, cronometro2, cronometro3 Cronometro
	tick1 := cronometro1.Init(1 * time.Second)
	tick2 := cronometro2.Init(3 * time.Second)
	tick3 := cronometro3.Init(6 * time.Second)

	go updateLabel(cron1TiempoLbl, tick1)
	go updateLabel(cron2TiempoLbl, tick2)
	go updateLabel(cron3TiempoLbl, tick3)

	botonPausaGeneral := widget.NewButton("Pausar durante 5 segundos", func() {
		cronometro1.Pausar()
		cronometro2.Pausar()
		cronometro3.Pausar()
		time.Sleep(5 * time.Second)
		cronometro1.Reanudar()
		cronometro2.Reanudar()
		cronometro3.Reanudar()
	})

	cron1Iniciar := widget.NewButton("Iniciar", func() {
		ponerTiempo(cron1TiempoLbl, 0)
		cronometro1.Iniciar()
	})
	cron1Pausar := widget.NewButton("Pausar", func() {
		cronometro1.Pausar()
	})
	cron1Reanudar := widget.NewButton("Reanudar", func() {
		cronometro1.Reanudar()
	})
	cron1Detener := widget.NewButton("Detener", func() {
		cronometro1.Detener()
	})

	cron2Iniciar := widget.NewButton("Iniciar", func() {
		ponerTiempo(cron2TiempoLbl, 0)
		cronometro2.Iniciar()
	})
	cron2Pausar := widget.NewButton("Pausar", func() {
		cronometro2.Pausar()
	})
	cron2Reanudar := widget.NewButton("Reanudar", func() {
		cronometro2.Reanudar()
	})
	cron2Detener := widget.NewButton("Detener", func() {
		cronometro2.Detener()
	})

	cron3Iniciar := widget.NewButton("Iniciar", func() {
		ponerTiempo(cron3TiempoLbl, 0)
		cronometro3.Iniciar()
	})
	cron3Pausar := widget.NewButton("Pausar", func() {
		cronometro3.Pausar()
	})
	cron3Reanudar := widget.NewButton("Reanudar", func() {
		cronometro3.Reanudar()
	})
	cron3Detener := widget.NewButton("Detener", func() {
		cronometro3.Detener()
	})

	cron1Container := container.NewVBox(
		widget.NewLabel("Cronómetro 1"),
		cron1Iniciar,
		cron1Pausar,
		cron1Reanudar,
		cron1Detener,
		cron1TiempoLbl,
	)
	cron2Container := container.NewVBox(
		widget.NewLabel("Cronómetro 2"),
		cron2Iniciar,
		cron2Pausar,
		cron2Reanudar,
		cron2Detener,
		cron2TiempoLbl,
	)
	cron3Container := container.NewVBox(
		widget.NewLabel("Cronómetro 3"),
		cron3Iniciar,
		cron3Pausar,
		cron3Reanudar,
		cron3Detener,
		cron3TiempoLbl,
	)

	contenido := container.NewVBox(
		botonPausaGeneral,
		cron1Container,
		cron2Container,
		cron3Container,
	)

	myWindow.SetContent(contenido)
	myWindow.ShowAndRun()
}
