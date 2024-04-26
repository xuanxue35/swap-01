package main

import (
    "fmt"
)

type Transaction struct {
    From     string
    To       string
    Amount   uint
    Fee      uint
    Status   string
}

func NewTransaction(from, to string, amount uint) *Transaction {
    return &Transaction{From: from, To: to, Amount: amount}
}

func (t *Transaction) Confirm() {
    fmt.Printf("确认交易：从 %s 向 %s 转账 %d TEA\n", t.From, t.To, t.Amount)
    // 在此执行交易确认逻辑，需要调用 Binance 链的 API
}

func (t *Transaction) CalculateFee() {
    // 在此计算交易手续费
    // 这里只是一个示例，实际应根据 Binance 链的规则计算手续费
    t.Fee = t.Amount / 100
    fmt.Printf("交易手续费：%d TEA\n", t.Fee)
}
