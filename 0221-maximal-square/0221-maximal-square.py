
"""

max int 

if cell == 1


get the current square of the current position


from x,y 

find the width and the height

[1] => 1

[0] => 0



all the members of a square have at least 2 neighbors 




max height 
max width



"""

class Solution:
    def maximalSquare(self, matrix: List[List[str]]) -> int:
        coln, rown = len(matrix), len(matrix[0])
        dp = [[0] * (rown+1) for _ in range(coln+1)]
        m = 0 
        print(dp)
        for col in range(1,coln+1):
            for row in range(1, rown+1):
                if matrix[col-1][row-1] == "1":
                    dp[col][row] = min(dp[col][row-1], dp[col-1][row], dp[col-1][row-1]) +1
                    m = max(dp[col][row], m)       

        return m * m 