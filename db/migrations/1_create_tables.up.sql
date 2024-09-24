CREATE TABLE cars (
    id SERIAL PRIMARY KEY,
    engine_volume DECIMAL(1,1) NOT NULL,
    color VARCHAR(50) NOT NULL,
    brand VARCHAR(50) CHECK (brand IN ('BMW', 'Mercedes', 'Lada')),
    wheel_position VARCHAR(10) CHECK (wheel_position IN ('Left', 'Right')),
    price DECIMAL NOT NULL
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    balance NUMERIC DEFAULT 1000000
);

CREATE TABLE user_cars (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    car_id INT REFERENCES cars(id),
    UNIQUE(user_id, car_id),
    FOREIGN KEY (model_id) REFERENCES car_models(id)
);