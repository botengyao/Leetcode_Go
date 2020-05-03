class Solution:   
    def numberToWords(self, num: int) -> str:
        #1 billion
        #1 million
        #1 thousand
        #1 handred
        #0 -> Zero
        #100_000_001 -> One Handred Million One
        #100 -> One Handred !no space at the end
        #111 -> One Handred Eleven
        #19 -> Nineteen
        def one(num):
            switcher = {
                1: 'One',
                2: 'Two',
                3: 'Three',
                4: 'Four',
                5: 'Five',
                6: 'Six',
                7: 'Seven',
                8: 'Eight',
                9: 'Nine'
            }
            return switcher.get(num)

        def two_less_20(num):
            switcher = {
                10: 'Ten',
                11: 'Eleven',
                12: 'Twelve',
                13: 'Thirteen',
                14: 'Fourteen',
                15: 'Fifteen',
                16: 'Sixteen',
                17: 'Seventeen',
                18: 'Eighteen',
                19: 'Nineteen'
            }
            return switcher.get(num)

        def ten(num):
            switcher = {
                2: 'Twenty',
                3: 'Thirty',
                4: 'Forty',
                5: 'Fifty',
                6: 'Sixty',
                7: 'Seventy',
                8: 'Eighty',
                9: 'Ninety'
            }
            return switcher.get(num)
        
        def dfsnumber(num):
            string = ""
            if num == 0:
                return ""
            elif num < 100:                               
                res = num // 10
                remain = num % 10
                #10 - 19
                if res == 1:
                    return two_less_20(num)
                #1 - 9          
                elif res == 0:
                    return one(remain)
                else:
                    # 20, 30, 40, 50 ... 90
                    if remain == 0:
                        return ten(res)
                    #others
                    else:
                        return ten(res) + " " + one(remain)
            else:    
                unit = ["Billion", "Million", "Thousand", "Hundred"]
                devide = []
                for dec in [1000000000, 1000000, 1000, 100]:
                    res = num // dec
                    num = num % dec
                    devide.append(res)
                #The length of devide is 5, the last one is the number that is smaller than 100
                devide.append(num)

                for i in range(len(unit)):
                    temp = dfsnumber(devide[i]) #Devide and Conquer
                    if temp != "":
                        if temp[-1] == " ": # Handle "One Handred Million" case
                            temp = temp.strip(" ")
                        string += temp + " " + unit[i] + " "
                        
                string += dfsnumber(devide[-1])
                return string.strip(" ") #Handle "One Handred" case
        
        #"Zero" is an edge case
        if num == 0:
            return "Zero"

        return dfsnumber(num)