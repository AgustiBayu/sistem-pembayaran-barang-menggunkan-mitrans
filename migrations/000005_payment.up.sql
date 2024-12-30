CREATE TYPE payment_method_enum AS ENUM ('mitrans');
CREATE TYPE payment_status_enum AS ENUM ('pending', 'completed', 'failed');

CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    pesanan_id int NOT NULL,
    payment_method payment_method_enum NOT NULL,
    payment_status payment_status_enum DEFAULT 'pending',
    payment_amount int NOT NULL,
    payment_date Date,    
    FOREIGN KEY (pesanan_id) REFERENCES pesanans(id)
);