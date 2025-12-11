// BankAccount Sstruct with Deposit() and Withdraw() Methods
package main
import "fmt"

type BankAccount struct{
	balance float64
}
func new() BankAccount{
	return BankAccount{balance: 0}
}
func (b *BankAccount) Deposit(amount float64){
	b.balance += amount
}
func (b *BankAccount) Withdraw(amount float64) string{
	if amount > b.balance{
		return "insufficient funds"
	}
	b.balance -= amount
	return "withdrawal successful"
}
func main(){
	account :=new()
	var deposit float64
	fmt.Println("Enter amount to deposit:")
	fmt.Scanf("%f", &deposit)
	account.Deposit(deposit)
	fmt.Println("Current Balance:", account.balance)

	var withdraw float64
	fmt.Println("Enter amount to withdraw:")
	fmt.Scanf("%f", &withdraw)
	message := account.Withdraw(withdraw)
	fmt.Println(message)
	fmt.Println("Current Balance:", account.balance)
}