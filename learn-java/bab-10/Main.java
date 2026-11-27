import java.util.HashMap;
import java.util.Map;

public class Main {
    record Permintaan(String id, int jumlah) {}
    record Bukti(int jumlah, String hasil) {}

    static class Gudang {
        private int stok;
        private final Map<String, Bukti> riwayat = new HashMap<>();

        Gudang(int stok) {
            this.stok = stok;
        }

        String pesan(Permintaan r) {
            if (r.jumlah() < 1 || r.jumlah() > 5) return "jumlah_invalid";

            // Record menyimpan jumlah asli untuk mendeteksi pemakaian ulang ID.
            Bukti awal = riwayat.get(r.id());
            if (awal != null) {
                if (awal.jumlah() != r.jumlah()) return "konflik";
                return awal.hasil();
            }

            String hasil = "stok_kurang";
            if (r.jumlah() <= stok) {
                stok -= r.jumlah();
                hasil = "diterima";
            }
            riwayat.put(r.id(), new Bukti(r.jumlah(), hasil));
            return hasil;
        }
    }

    public static void main(String[] args) {
        var gudang = new Gudang(5);
        var urutan = new Permintaan[] {
            new Permintaan("A", 3), new Permintaan("B", 3),
            new Permintaan("C", 0), new Permintaan("C", 2),
            new Permintaan("A", 3), new Permintaan("A", 1),
            new Permintaan("B", 3)
        };
        for (var r : urutan) {
            String hasil = gudang.pesan(r);
            System.out.printf("%s %d %s stok=%d%n", r.id(), r.jumlah(), hasil, gudang.stok);
        }
    }
}
