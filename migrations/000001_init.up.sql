CREATE TABLE IF NOT EXISTS persons (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255),
	surname VARCHAR(255),
	patronymic VARCHAR(255),
	age INTEGER,
	gender VARCHAR(10),
	country VARCHAR(255)
);
CREATE INDEX idx_name ON persons (name);
CREATE INDEX idx_surname ON persons (surname);
CREATE INDEX idx_age ON persons (age);
CREATE INDEX idx_country ON persons (country);
