package main

import (
    "fmt"
)

type Token struct {
    Name   string
    Symbol string
    Amount uint
}

func NewToken(name, symbol string, amount uint) *Token {
    return &Token{Name: name, Symbol: symbol, Amount: amount}
}

func (t *Token) Transfer(to string, amount uint) {
    fmt.Printf("向 %s 转账 %d %s\n", to, amount, t.Symbol)
    // 在此执行代币转账逻辑，需要调用 Binance 链的 API
}
