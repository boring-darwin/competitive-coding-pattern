package main

func main() {

}

func letterCombination(num int) {
	output := make([]string, 0)
	words := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var str string
	var backtrack func(n int)
	nums := make([]int, 0)

	for num > 0 {
		nums = append(nums, num%10)
		num = num / 10
	}

	backtrack = func(n int) {
		if len(str) > 0 {
			output = append(output, str)
			return
		}

		for i:=len(nums)-1; i<=n; i++ {
			word := words[i-2]
			
			for j:=0; j<len(word); j++ {
				str = str+string(word[j])
			}
		}
	}
}
