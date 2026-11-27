<?php

interface AturanOngkir
{
    public function hitung(int $subtotal): int;
}

final class OngkirTetap implements AturanOngkir
{
    public function hitung(int $subtotal): int
    {
        return 7000;
    }
}

final class OngkirGratisMinimum implements AturanOngkir
{
    public function hitung(int $subtotal): int
    {
        return $subtotal >= 50000 ? 0 : 7000;
    }
}

function hitungTagihan(int $subtotal, AturanOngkir $aturan): int
{
    return $subtotal + $aturan->hitung($subtotal);
}

foreach ([new OngkirTetap(), new OngkirGratisMinimum()] as $aturan) {
    echo "total=" . hitungTagihan(52000, $aturan) . "\n";
}
