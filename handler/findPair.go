package handler

import (
	"backend/types"

	"github.com/gin-gonic/gin"
)

func FindPairHandler(c *gin.Context) {
	reqBody := types.FindPairreq{}
	if err := c.ShouldBindBodyWithJSON(&reqBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := findPair(reqBody.Number, reqBody.Target)
	res := types.FindPairRes{Solutions: result}
	c.JSON(200, res)
}

func findPair(nums []int, target int) [][]int {
	result := [][]int{}
	tempResult := make(map[int]int)
	for idx, value := range nums {
		compliment := target - value

		if num, ok := tempResult[compliment]; ok {
			result = append(result, []int{num, idx})

		}
		tempResult[value] = idx
	}
	return result

}
