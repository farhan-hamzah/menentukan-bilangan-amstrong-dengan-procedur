// Bilangan Amstrong Dengan Procedur
package main
import "fmt"
func amstrong(n int, hasil *int){
    var i, j, k, l, pangkat int
    l = n
    i = n
    *hasil = 0
    for i > 0{
        j++
        i = i/10
    }
    i = j
    for n > 0{
        pangkat = 1
        k = n%10
        for j > 0{
            pangkat = pangkat * k
            j--
        }
        *hasil += pangkat
        j = i
        n = n/10
    }
    if *hasil == l{
        fmt.Print(true)
    }else{
        fmt.Print(false)
    }
}

func main(){
    var masukan, keluaran int
    fmt.Scan(&masukan)
    amstrong(masukan, &keluaran)
}