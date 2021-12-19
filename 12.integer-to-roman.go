/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 */

func intToRoman(num int) string {
    nums := [13]int{1,4,5,9,10,40,50,90,100,400,500,900,1000}
    chars := [13]string{"I","IV","V","IX","X","XL","L","XC","C","CD","D","CM","M"}

    roman := ""

    for i := 0; i < 13; i++{
        j := 13 - (i + 1)
        if num >= 0 {
            q := num / nums[j]
            num = num % nums[j]
            if q > 0 {
                for k := 0; k < q; k++{
                    roman += chars[j]
                }
            }
        }
    }

    return roman
}
// @lc code=end

