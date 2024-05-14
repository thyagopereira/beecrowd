def spiralTraverse(array):
    result = []

    start_row, end_row = 0, len(array) - 1
    start_col, end_col = 0, len(array[0]) - 1

    if len(array) == 0:
        return []
    if len(array) == 1:
        return array[0]

    while start_row <= end_row and start_col <= end_col:

        # -->
        for col in range(start_col, end_col + 1):
            result.append(array[start_row][col])
        # |
        # v
        for row in range(start_row + 1, end_row + 1):
            result.append(array[row][end_col])

        # <--
        if end_row != start_row:
            for fcol in reversed(range(start_col, end_col)):
                result.append(array[end_row][fcol])

        # ^
        # |
        if end_col != start_col:
            for frow in reversed(range(start_row + 1, end_row)):
                result.append(array[frow][start_col])

        start_row += 1
        end_row -= 1
        start_col += 1
        end_col -= 1

    return result