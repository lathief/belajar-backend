import java.util.OptionalInt;

public class Main {
    static class Tugas {
        int generasi;
        int kedaluwarsa;
        boolean selesai;

        OptionalInt klaim(int kini) {
            if (selesai || kini < kedaluwarsa) return OptionalInt.empty();
            kedaluwarsa = kini + 5;
            return OptionalInt.of(++generasi);
        }

        String simpan(int token, int kini) {
            // Token lama dan token kedaluwarsa sama-sama kehilangan hak tulis.
            if (token == 0 || selesai || token != generasi || kini >= kedaluwarsa) return "ditolak";
            selesai = true;
            return "tersimpan";
        }
    }

    public static void main(String[] args) {
        var t = new Tugas();
        int a = t.klaim(0).orElseThrow();
        System.out.println("klaim_dini=" + t.klaim(4).isPresent());
        System.out.println("batas=" + t.simpan(a, 5));
        int b = t.klaim(5).orElseThrow();
        System.out.println("lama=" + t.simpan(a, 6));
        System.out.println("baru=" + t.simpan(b, 6));
        System.out.println("ulang=" + t.simpan(b, 7));
    }
}
