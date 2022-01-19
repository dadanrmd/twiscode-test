CREATE TABLE IF NOT EXISTS transaction_details(
    id BIGINT AUTO_INCREMENT,
    id_transaksi BIGINT,
    harga BIGINT,
    jumlah BIGINT,
    subtotal BIGINT,
    constraint transaction_details_pk
        primary key (id)
);

INSERT INTO transaction_details (`id`, `id_transaksi`, `harga`, `jumlah`, `subtotal`) VALUES (NULL, '1', '2000', '2', '4000');
INSERT INTO transaction_details (`id`, `id_transaksi`, `harga`, `jumlah`, `subtotal`) VALUES (NULL, '1', '25000', '4', '100000');