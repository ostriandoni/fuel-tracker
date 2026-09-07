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

CREATE TABLE locations (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME NULL
);

CREATE TABLE petrol_types (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    location_id BIGINT UNSIGNED NOT NULL,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(12,2) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME NULL,
    FOREIGN KEY (location_id) REFERENCES locations(id)
);

ALTER TABLE fuel_logs
    DROP COLUMN location,
    ADD COLUMN location_id BIGINT UNSIGNED,
    ADD COLUMN petrol_type_id BIGINT UNSIGNED,
    ADD FOREIGN KEY (location_id) REFERENCES locations(id),
    ADD FOREIGN KEY (petrol_type_id) REFERENCES petrol_types(id);
