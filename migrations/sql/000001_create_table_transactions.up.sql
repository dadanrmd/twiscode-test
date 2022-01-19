CREATE TABLE IF NOT EXISTS transactions(
    id BIGINT AUTO_INCREMENT,
    tanggal_order timestamp NOT NULL DEFAULT current_timestamp(),
    status VARCHAR(50) NULL,
    tanggal_pembayaran timestamp NOT NULL DEFAULT current_timestamp(),
    constraint transactions_pk
        primary key (id)
);
INSERT INTO transactions (`id`, `tanggal_order`, `status`, `tanggal_pembayaran`) VALUES (NULL, current_timestamp(), 'lunas', current_timestamp());
INSERT INTO transactions (`id`, `tanggal_order`, `status`, `tanggal_pembayaran`) VALUES (NULL, current_timestamp(), 'pending', current_timestamp());