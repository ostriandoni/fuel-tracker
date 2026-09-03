CREATE TABLE fuel_logs (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    date DATE NOT NULL,
    price_per_liter DECIMAL(12,2) NOT NULL,
    total_paid DECIMAL(12,2) NOT NULL,
    liters_filled DECIMAL(10,2) NOT NULL,
    km_start INT NOT NULL,
    km_end INT NOT NULL,
    location VARCHAR(255),
    notes TEXT,
    created_at DATETIME,
    updated_at DATETIME
);

ALTER TABLE fuel_logs ADD COLUMN deleted_at DATETIME NULL;
CREATE INDEX idx_fuel_logs_deleted_at ON fuel_logs(deleted_at);
