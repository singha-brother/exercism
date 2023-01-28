def response(hey_bob):
    hey_bob = hey_bob.strip()
    is_question = hey_bob.endswith("?")
    all_upper = hey_bob.isupper()
    is_empty = hey_bob == ""
    if is_question and all_upper:
        return "Calm down, I know what I'm doing!"
    if all_upper:
        return "Whoa, chill out!"
    if is_question:
        return "Sure."
    if is_empty:
        return "Fine. Be that way!"
    return "Whatever."
