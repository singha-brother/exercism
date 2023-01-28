def square(number):
    if number >= 1 and number <= 64:
        return 1 << (number - 1)
    raise ValueError("square must be between 1 and 64")


def total():
    # return sum([1 << n for n in range(64)])
    return (1 << 64) - 1
