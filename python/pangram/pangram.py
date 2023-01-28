import string


def is_pangram(sentence):
    return len(set(string.ascii_lowercase).difference(sentence.lower())) == 0
