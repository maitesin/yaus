from string import ascii_letters, digits
from random import choices


def get_id_generator():
    values = ascii_letters + digits + ' '

    while True:
        yield ''.join(choices(values, k=8)).strip().replace(' ', '')


id_generator = get_id_generator()
