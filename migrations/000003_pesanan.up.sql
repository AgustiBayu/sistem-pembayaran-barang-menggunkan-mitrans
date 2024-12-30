CREATE TYPE status_pesanan_enum AS ENUM ('pending', 'completed', 'cancelled');

CREATE TABLE pesanans (
    id SERIAL PRIMARY KEY,
    pelanggan_id INT,
    total_amount int NOT NULL,
    status status_pesanan_enum DEFAULT 'pending',
    created_at Date,
    FOREIGN KEY (pelanggan_id) REFERENCES pelanggans(id)
);