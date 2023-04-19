n = int(input())

errOne = [int(x) for x in input().split(" ")]
errTwo = [int(x) for x in input().split(" ")]
errThree = [int(x) for x in input().split(" ")]

sumOne = sum(errOne)
sumTwo = sum(errTwo)
sumThree = sum(errThree)
    
f = sumOne - sumTwo
s = sumTwo - sumThree

print(f)
print(s)