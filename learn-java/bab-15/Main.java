public class Main {
    static class Ember {
        final int kapasitas;
        int token;
        int terakhir;

        Ember(int kapasitas) {
            this.kapasitas = kapasitas;
            this.token = kapasitas;
        }

        String ambil(int detik) {
            // Waktu menganggur mengisi ember hingga kapasitas maksimum.
            token = Math.min(kapasitas, token + detik - terakhir);
            terakhir = detik;
            if (token == 0) return "tunda";
            token--;
            return "lolos";
        }
    }

    public static void main(String[] args) {
        var e = new Ember(2);
        for (int t : new int[]{0, 0, 0, 1, 6, 6, 6}) {
            String hasil = e.ambil(t);
            System.out.printf("t=%d %s token=%d%n", t, hasil, e.token);
        }
    }
}
