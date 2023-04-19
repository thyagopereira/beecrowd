books = [int(x) for x in input().split(" ")]

min_cost = 0

if books[0] % 4 == 1:
    min_cost = min(books[1] *3, books[1] + books[2], books[3])
elif books[0] % 4 == 2:
    min_cost = min(books[1] * 2, books[2], books[3] * 2)
elif books[0] % 4 == 3:
    min_cost = min(books[1], books[2] + books[3], books[3]*3)

print(min_cost)