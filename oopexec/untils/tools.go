package tools

import "fmt"

type account struct {
	accountNo string
	pwd       string
	balance   float64
}

func Newacc(accountNo string, pwd string, balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 20 {
		fmt.Println("账号长度不对。。。。")
		return nil
	}

	if len(pwd) < 6 {
		fmt.Println("密码长度不对。。。。")
		return nil
	}

	if balance < 20 {
		fmt.Println("余额太小。。。")
		return nil
	}

	return &account{
		accountNo: accountNo,
		pwd:       pwd,
		balance:   balance,
	}

}

func (a *account) Setpwd(pwd string) {
	if len(pwd) < 6 {
		fmt.Println("请输入一个大于6位数的密码")
	} else {
		a.pwd = pwd
	}
}

func (a *account) Query(pwd string) {
	if pwd == a.pwd {
		fmt.Printf("当前账户：%V, 余额为： %v", a.accountNo, a.balance)
	} else {
		fmt.Println("输入密码有误")
	}
}

func (a *account) Getmeny(pwd string, meny float64) {
	if pwd == a.pwd {
		if meny <= a.balance {
			a.balance -= meny
		} else {
			fmt.Println("当前余额数太少")
		}
	} else {
		fmt.Println("密码不正确")
	}

}
