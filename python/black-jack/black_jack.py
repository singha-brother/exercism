"""Functions to help play and score a game of blackjack.

How to play blackjack:    https://bicyclecards.com/how-to-play/blackjack/
"Standard" playing cards: https://en.wikipedia.org/wiki/Standard_52-card_deck
"""


def value_of_card(card):
    """Determine the scoring value of a card.

    :param card: str - given card.
    :return: int - value of a given card.  See below for values.

    1.  'J', 'Q', or 'K' (otherwise known as "face cards") = 10
    2.  'A' (ace card) = 1
    3.  '2' - '10' = numerical value.
    """
    if card in ['J', 'Q', 'K']:
        return 10
    if card == 'A':
        return 1
    try:
        val = int(card)
        return val
    except ValueError:
        pass


def higher_card(card_one, card_two):
    val_one = value_of_card(card_one)
    val_two = value_of_card(card_two)
    if val_one == val_two:
        return card_one, card_two
    if val_one > val_two:
        return card_one
    return card_two


def value_of_ace(card_one, card_two):
    has_A = (card_one == 'A') or (card_two == 'A')
    sum_of_two = value_of_card(card_one) + value_of_card(card_two)
    if has_A:
        return 1
    if sum_of_two + 11 > 21:
        return 1
    return 11


def is_blackjack(card_one, card_two):

    contain_A = (card_one == 'A') or (card_two == 'A')
    is_ten = (value_of_card(card_one) == 10) or (value_of_card(card_two) == 10)
    return contain_A and is_ten


def can_split_pairs(card_one, card_two):
    return value_of_card(card_one) == value_of_card(card_two)


def can_double_down(card_one, card_two):
    if card_one == 'A' and card_two == 'A':
        return False
    return value_of_card(card_one) + value_of_card(card_two) < 12
