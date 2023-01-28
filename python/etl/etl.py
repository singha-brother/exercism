def transform(legacy_data):
    result = {}
    for point, chars in legacy_data.items():
        for char in chars:
            result[char.lower()] = point

    return result
