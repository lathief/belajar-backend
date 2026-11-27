public class Main {
    enum Status { BARU, BAYAR }
    record Pesanan(String kode, Status status) {}

    public static void main(String[] args) {
        var pesanan = new Pesanan("PS-1", Status.BAYAR);
        // Setiap status memerlukan label tampilan.
        String label = switch (pesanan.status()) {
            case BARU -> "menunggu";
            case BAYAR -> "lunas";
        };
        System.out.println(pesanan.kode() + " " + label);
    }
}
