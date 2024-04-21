package main

import (
	"runtime"
	"sync"
	"time"
)

type Cronometro struct {
	sync.Mutex
	transcurrido time.Duration
	incremento   time.Duration
	tick         chan time.Duration
	estado       int
	cronometro   time.Ticker
}

const (
	Detenido      = 0
	Pausado       = 1
	Cronometrando = 2
)

func (c *Cronometro) Init(inc time.Duration) chan time.Duration {
	c.transcurrido = 0
	c.incremento = inc
	c.estado = Detenido
	c.tick = make(chan time.Duration)
	return c.tick
}

func cronometrar(c *Cronometro) {
	for {
		c.Lock()
		switch c.estado {
		case Detenido:
			runtime.Gosched()
		case Pausado:
			runtime.Gosched()
		default:
			select {
			case <-c.cronometro.C:
				c.transcurrido += c.incremento
				c.tick <- c.transcurrido
			default:
				runtime.Gosched()
			}
		}
		c.Unlock()
	}
}

func (c *Cronometro) Iniciar() {
	c.transcurrido = 0
	if c.estado == Detenido {
		c.cronometro = *time.NewTicker(c.incremento)
		go cronometrar(c)
	}
	c.estado = Cronometrando
}

func (c *Cronometro) Reanudar() {
	if c.estado == Pausado {
		c.cronometro = *time.NewTicker(c.incremento)
		c.estado = Cronometrando
	}
}

func (c *Cronometro) Detener() {
	c.cronometro.Stop()
	c.estado = Detenido
}

func (c *Cronometro) Pausar() {
	if c.estado != Detenido {
		c.cronometro.Stop()
		c.estado = Pausado
	}
}
