CREATE TABLE executors (
    executor_id BIGINT PRIMARY KEY
);

CREATE TABLE customers (
    customer_id BIGINT PRIMARY KEY
);

CREATE TABLE zones (
    zone_id BIGINT PRIMARY KEY,
    zone_name VARCHAR(255)
);

CREATE TYPE order_execution_status AS ENUM ('assigned', 'acquired', 'cancelled', 'completed');

CREATE TABLE orders (
    order_id BIGINT NOT NULL PRIMARY KEY,
    executor_id BIGINT NOT NULL REFERENCES executors(executor_id),
    customer_id BIGINT NOT NULL REFERENCES customers(customer_id),
    execution_status order_execution_status NOT NULL,
    coin_coefficient DOUBLE PRECISION NOT NULL,
    coin_bonus_amount DOUBLE PRECISION NOT NULL,
    final_coin_amount DOUBLE PRECISION NOT NULL,
    zone_id BIGINT NOT NULL REFERENCES zones(zone_id),
    assign_time TIMESTAMP WITH TIME ZONE NOT NULL,
    first_acquire_time TIMESTAMP WITH TIME ZONE,
    executor_rating FLOAT NOT NULL,
    completed_time TIMESTAMP WITH TIME ZONE,

    UNIQUE (order_id)
);
