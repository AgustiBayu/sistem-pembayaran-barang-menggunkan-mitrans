CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    pesanan_id int NOT NULL,
    payment_method TEXT CHECK (payment_method IN ('mitrans')) NOT NULL,
    payment_status TEXT CHECK (payment_status IN ('pending', 'completed', 'failed')) DEFAULT 'pending',
    payment_amount int NOT NULL,
    payment_date Date,
    FOREIGN KEY (pesanan_id) REFERENCES pesanans(id)
);