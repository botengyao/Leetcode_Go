class Solution:
    def isAlienSorted(self, words: List[str], order: str) -> bool:
        #words
        #word
        #row
        orderdict = {}
        for i in range(len(order)):
            orderdict[order[i]] = i
        for i in range(1, len(words)):
            #words[i]
            #words[i - 1]
            count = 0
            length = min(len(words[i]), len(words[i - 1]))
            for j in range(length):
                if orderdict[words[i][j]] < orderdict[words[i - 1][j]]:
                    return False
                elif orderdict[words[i][j]] == orderdict[words[i - 1][j]]:
                    count += 1  #use count to handle this situation: words > word
                else:
                    break
                
            if count == length and len(words[i]) < len(words[i - 1]):
                return False
        return True