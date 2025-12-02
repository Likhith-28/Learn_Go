// Write a function that increases Product by 10%
package main
import "fmt"
type Product struct{
	name string
	price float64
}
func new(name string,price float64)*Product{
	return &Product{
		name:name,
		price:price,
	}
}
func (p *Product) increaseprice(price float64){
	p.price=price+(p.price*0.10)
}
func main(){
	prod:=new("Laptop",800.00)
	fmt.Println("Before increase, price:", prod.price)
	prod.increaseprice(prod.price)
	fmt.Println("After 10% increase, price:", prod.price)
}