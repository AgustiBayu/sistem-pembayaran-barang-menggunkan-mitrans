CREATE TABLE pesanans (
    id SERIAL PRIMARY KEY,
    pelanggan_id INT,
    total_amount DECIMAL(10, 2) NOT NULL,
    status TEXT CHECK (status IN ('pending', 'completed', 'cancelled')) DEFAULT 'pending',
    created_at Date,
    FOREIGN KEY (pelanggan_id) REFERENCES pelanggans(id)
);