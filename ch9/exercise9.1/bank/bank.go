package bank

var deposits = make(chan int) // 发送存款额
var balances = make(chan int) // 接受余额

// Deposit 存入 amount 金额入银行
func Deposit(amount int) { deposits <- amount }

// Balance 返回银行余额
func Balance() int { return <-balances }

type draw struct {
	amount  int
	success chan bool
}

var withDraws = make(chan draw)

// Withdraw 返回布尔值，代表交易是否成功
func Withdraw(amount int) bool {
	success := make(chan bool)
	withDraws <- draw{amount, success}
	return <-success
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case draw := <-withDraws:
			if draw.amount <= balance {
				balance -= draw.amount
				draw.success <- true
			} else {
				draw.success <- false
			}
		}
	}
}

// 初始化，启动监控余额的goroutine
func init() {
	go teller()
}
