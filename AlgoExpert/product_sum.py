# Tip: You can use the type(element) function to check whether an item
# is a list or an integer.
def productSum(array):
    return productWrap(1, array)
    
def productWrap(factor, array):
    sum = 0
    for elem in array:
        if type(elem) == list:
            sum += productWrap(factor + 1, elem)
        else:
            sum += elem

    return sum * factor


def main(): 
    input = [5, 2, [7, -1], 3, [6, [-13, 8], 4]]
    expected  = 12
    output = productSum(input)


    print(output)
    assert(output == expected)


main()