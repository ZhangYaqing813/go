package tools

import "strconv"

func reverse (slice []string ) []string {

	for i,j :=0, len(slice)-1; i <j; i++{
		slice[i], slice[j] = slice[j],slice[i]
		j--
	}

	return slice
}

func ToBy(num int) []string {
	tmp := ""
	slice := []string{}
	for ;num>0; {
		tmp = strconv.Itoa(num % 2)
		num /= 2
		slice = append(slice, tmp)
	}
	//翻转字符串


	return reverse(slice)
}

func ToOX(num int) []string  {
	tmp := ""
	slice := []string{}
	for ;num>0; {
		tmp = strconv.Itoa(num % 8)
		num /= 8
		slice = append(slice, tmp)
	}
	//翻转字符串
	return reverse(slice)

}

func powerf(BasNum int ,power int) int {
	/*
		BasNum 是二进制每一个位的数字，只有两种情况，0或者是1
		power 是代表了BasNum 是在第几位，当BasNum 不为0的情况下，能计算出该位的十进制数字
		例如 二进制1010= 十进制的 10
		计算过程（从右到左）
		二进制位数 第一位0，第二位1，第三位0，第四位1，计算机是从0开始，那么最高位就是4-1 = 3
		1*2^3+0*2^2+1*2^1+0*2^0 = 10

	*/
	var sum int = 1  //定义一个初始值sum = 1 ,1乘以任何数都不变
	if BasNum == 0{
		return 0
	}  // 如果该位是0，不用计算，直接返回0,
	// 计算每一位不为零所代表的十进制数字，并返回
	for i := power; i >0; i--{
		sum *= 2
	}
	return sum
}



func BaniyToDec(num int) int {
	var res int = 0
	slice := []string{}
	var tmp string = strconv.Itoa(num)
	// 遍历字符串存入中间变量slice
	for _, v := range tmp{
		slice = append(slice, strconv.Itoa(int(v)))
	}

	for i,v := range reverse(slice){
		v, _ := strconv.Atoi(v)
		//调用 powerf 函数获取二进制单位字符的所代表的十进制数字，使用res 进行累加
		res = res + powerf(v-48, int(i))
	}
	return res
}