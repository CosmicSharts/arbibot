package main

import (
	"fmt"
	"net/http"

	"github.com/pablonlr/arbibot/exchange/cex/binance/pclient"

	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pablonlr/arbibot/coin"
	"github.com/pablonlr/arbibot/exchange/dex/bscdex"
)

const address = "0x325E343f1dE602396E256B67eFd1F61C3A6B38Bd"

func main() {
	cl := pclient.Client{&http.Client{}}
	p1, p2, err := cl.UpdatedPrice("XLM", "USDT")
	if err != nil {
		panic(err)
	}
	fmt.Println(p1)
	fmt.Println(p2)
	client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		log.Fatal(err)
	}
	routeToUsd := []common.Address{
		common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"),
	}
	dex, err := bscdex.NewBSCDEX("babyswap", address, client, routeToUsd)
	if err != nil {
		panic(err)
	}
	xlm := coin.NewBSCCoin("XLM", "Stellar", "0x43C934A845205F0b514417d757d7235B8f53f1B9")
	usdt := coin.NewBSCCoin("USDT", "USDT", "0x55d398326f99059fF775485246999027B3197955")
	i, err := dex.GetExchangeAmount(100000, xlm, usdt)
	if err != nil {
		panic(err)
	}
	f1 := float64(i) / 100000
	i2, err := dex.GetExchangeAmount(100000, usdt, xlm)
	if err != nil {
		panic(err)
	}
	f2 := (1.0 / float64(i2)) * 100000
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println((1 - (f1 / p1)) * 100)

}
