public class Main {
    public static void main(String[] args) {
        final int harga = 5000;
        int jumlah = 2;
        // Jumlah masih dapat berubah ketika barang ditambahkan.
        jumlah++;
        System.out.println("subtotal=" + harga * jumlah);
    }
}
