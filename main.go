package main

import "fmt"
import "os"
import "strconv"
import "time"
import "math/rand"
import "atomicgo.dev/cursor"


//Here are a lot of ANSI codes for the coloured output
var o=[]string{
//"black",
"red",
"green",
"yellow",
"blue",
"magenta",
"cyan",
"white",
//"reset",
}
var f=map[string]string{
"black": "\x1b[30m",
"red": "\x1b[31m",
"green": "\x1b[32m",
"yellow": "\x1b[33m",
"blue": "\x1b[34m",
"magenta": "\x1b[35m",
"cyan": "\x1b[36m",
"white": "\x1b[37m",
"reset":"\033[0m",
}
var b=map[string]string{
"black": "\x1b[40m",
"red": "\x1b[41m",
"green": "\x1b[42m",
"yellow": "\x1b[43m",
"blue": "\x1b[44m",
"magenta": "\x1b[45m",
"cyan": "\x1b[46m",
"white": "\x1b[47m",
"reset":"\033[0m",
}

func main(){
	//How many rows and columns shall be rendered?
	rows, columns:=getconfig()
	// We print a block repeatedly
	repeat(5, func(){block(rows, columns)})
}
// creates a random escape code for foreground and background and they are concatenated
func ranCol()string{
	return f[o[rand.Intn(len(o))]]+b[o[rand.Intn(len(o))]]
}
// repeat the function 'a' 'r' many times, wait 1 second in between
func repeat(r int, a func()){
	for i:=0;i<r;i++{
		a()
		time.Sleep(time.Second)
		cursor.Up(10)
	}
	a()
}
// We read number of rows and columns from commandline arguments, the default are 20x20
func getconfig()(rows, columns int){
		rows=20
		columns=20
		if len(os.Args)==3{
			s,err:=strconv.Atoi(os.Args[1])
			rows=s
			columns,err=strconv.Atoi(os.Args[2])
			if err!=nil{
				fmt.Fprint(os.Stderr,"wrong input")
				return
			}
		}
		return
}
// make and print the block with given number of rows and columns
func block(rows, columns int){
	// because the number of rows must be even, it is increased by one if necessary
	rows=rows+(rows%2)
	//'a' contains the info about the block
	a :=make([][]int,rows)
	//populate every row with column entries
	for i:=range a{
		a[i]=make([]int,columns)
	}
	//fill the block with random values
	populate(&a)
	//render the block
	draw(&a)
}
// fill the block with either zero or one
func populate(a *[][]int){
	for i:=0;i<len(*a);i++{
		for j:=0;j<len((*a)[i]);j++{
			(*a)[i][j]=rand.Intn(2)
		}
	}
}
// render the block
func draw(a *[][]int){
	// every real row contains two rows from the block, that's why we increment by 2
	for i:=0;i<len(*a);i+=2{
		for j:=0;j<len((*a)[i]);j++{
			// We print every line with a prepented random colour escape code
			// But the character we print needs to be translated from the values of the block
			fmt.Print(ranCol()+translate((*a)[i][j], (*a)[i+1][j]))
		}
		fmt.Println(f["reset"])
	}
}
// return either a full block, low block or top block
func translate(a,b int)string{
	switch {
		case a==1 && b==1:
			return "█"
		case a==1 && b==0:
			return "▀"
		case a==0 && b==1:
			return "▄"
		default:
			return " "
	}
}
