public class Main {
    static int subtotal(int harga, int jumlah) {
        // Method perhitungan bisa dipanggil tanpa mencetak apa pun.
        return harga * jumlah;
    }

    public static void main(String[] args) {
        System.out.println("total=" + subtotal(5000, 3));
    }
}
