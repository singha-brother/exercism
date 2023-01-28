import string


def is_isogram(word):
    test = []
    for char in word:
        char = char.lower()
        if char in test:
            return False
        if char in string.ascii_lowercase:
            test.append(char)
    return True
