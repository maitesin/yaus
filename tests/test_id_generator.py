from yaus.id_generator import _trim


def test_trim():
    assert "12345678" == _trim("12345678")
    assert "2345678" == _trim(" 2345678")
    assert "1234567" == _trim("1234567 ")
    assert "1234678" == _trim("1234 678")
    assert "2367" == _trim(" 23  67 ")
