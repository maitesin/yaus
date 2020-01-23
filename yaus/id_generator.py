from string import ascii_letters, digits
from random import choices


def _get_id_generator():
    values = ascii_letters + digits + " "

    while True:
        yield _trim("".join(choices(values, k=8)))


def _trim(value):
    return value.strip().replace(" ", "")


id_generator = _get_id_generator()
