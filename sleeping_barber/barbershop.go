package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients.", barber)

		for {
			if len(shop.ClientsChan) == 0 {
				color.Yellow("There is nothing to do, so %s takes a nap.", barber)
				isSleeping = true
			}

			client, shopOpen := <-shop.ClientsChan

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up.", client, barber)
					isSleeping = false
				}

				// cut hair
				shop.cutHair(barber, client)
			} else {
				// shop is closed, so send the barber home and close this goroutine
				shop.sendBarberHome(barber)
				return
			}
		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Blue("%s is cutting %s' hair.", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Blue("%s is finished cutting %s' hair.", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Red("%s is going home.", barber)
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) closeShopForDay() {
	color.Cyan("Closing shop for the day.")
	close(shop.ClientsChan)
	shop.Open = false

	// waiting until all the barber are done
	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<-shop.BarbersDoneChan
	}

	close(shop.BarbersDoneChan)

	color.Magenta("-------------------------------------------------------------")
	color.Red("The barbershop is now closed for the day, and everyone has gone home.")
}

func (shop *BarberShop) addClient(client string) {
	color.Cyan("*** %s arrives!!", client)

	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Yellow("%s takes a seat in the waiting room.", client)
		default:
			color.Red("The waiting room is full, so %s leaves.", client)
		}

	} else {
		color.Red("The shop is already closed, so %s leaves!", client)
	}
}
