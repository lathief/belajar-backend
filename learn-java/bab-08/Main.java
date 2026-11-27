public class Main {
    public static void main(String[] args) {
        for (String teks : new String[]{"2", "0", "x"}) {
            try {
                int n = Integer.parseInt(teks);
                System.out.println(n > 0 ? "jumlah=" + n : "di luar batas");
            } catch (NumberFormatException e) {
                // Kesalahan pada satu input ditangani dalam iterasi itu.
                System.out.println("bukan angka");
            }
        }
    }
}
