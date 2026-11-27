public class Main {
    public static void main(String[] args) {
        int[] daftar = {0, 1, 4, 8};
        int valid = 0;
        for (int n : daftar) {
            // Pesanan sah berisi satu sampai lima unit.
            if (n >= 1 && n <= 5) {
                valid++;
            }
        }
        System.out.println("valid=" + valid);
    }
}
