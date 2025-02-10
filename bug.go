func main() {
    var i int
    fmt.Println(i) // Output: 0
    fmt.Println(&i) // Output: 0xc0000140a8

    j := 10
    fmt.Println(j)
    fmt.Println(&j) // Output: 0xc0000140b0
}