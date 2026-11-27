public class Main {
    public static void main(String[] args) {
        Persediaan stok = new Persediaan(2);
        stok.ambilSatu();
        System.out.println("sisa=" + stok.sisa());
    }
}

class Persediaan {
    private int sisa;

    Persediaan(int sisa) {
        this.sisa = sisa;
    }

    void ambilSatu() {
        // Pelanggan class tidak perlu mengubah field secara langsung.
        if (sisa > 0) sisa--;
    }

    int sisa() {
        return sisa;
    }
}
