package main

func main() {
	var y = []string{"A", "B", "C", "D"} //y的长度为4 cap 为4
	var x = y[:3]                        //x的长度为3 cap 为4
	//var d = x[0:3]                       //d的长度为3 cap 为4
	//x = append(x, "Z")                   //此时x的长度为4 cap为4 因为y的长度同样为4 且底层指向的是同一个数组，所以y也改变了
	//fmt.Println(x)                       //而因为d的长度没有改变所以还是 [A B c]
	//fmt.Println(y)
	//fmt.Println(d)
	for i, s := range x {
		print(i, s, ",")

		x = append(x, "Z")
		x[i+1] = "Z"
	}
}
