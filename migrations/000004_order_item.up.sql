CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    pesanan_id INT,
    produk_id INT,
    quantity INT NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (pesanan_id) REFERENCES pesanans(id),
    FOREIGN KEY (produk_id) REFERENCES produks(id)
);