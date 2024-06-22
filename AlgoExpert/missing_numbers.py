def missingNumbers(nums):
    total = sum(range(1, len(nums) + 3))
    for num in  nums:
        total -= num

    average = total // 2
    fristHalf, secondHalf = 0, 0
    for num in nums:
        if num <= average:
            fristHalf += num
        else:
            secondHalf += num

    expectedFristHalf = sum(range(1, average + 1))
    expectedSecondHalf = sum(range(average + 1, len(nums) + 3))

    return [expectedFristHalf - fristHalf , expectedSecondHalf - secondHalf]
    