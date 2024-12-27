CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    pesanan_id INT,
    produk_id INT,
    quantity INT NOT NULL,
    total int NOT NULL,
    FOREIGN KEY (pesanan_id) REFERENCES pesanans(id),
    FOREIGN KEY (produk_id) REFERENCES produks(id)
);