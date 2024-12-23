CREATE TABLE executors (
    executor_id VARCHAR(255) PRIMARY KEY
);

CREATE TABLE zones (
    zone_id VARCHAR(255) PRIMARY KEY
    zone_name VARCHAR(255)
);

CREATE TYPE order_execution_status AS ENUM ('assigned', 'acquired', 'cancelled', 'completed');

CREATE TABLE orders (
    order_id VARCHAR(255) NOT NULL,
    executor_id VARCHAR(255) NOT NULL,
    execution_status order_execution_status NOT NULL,
    coin_coefficient DOUBLE PRECISION NOT NULL,
    coin_bonus_amount DOUBLE PRECISION NOT NULL,
    final_coin_amount DOUBLE PRECISION NOT NULL,
    zone_name VARCHAR(255),
    assign_time TIMESTAMP WITH TIME ZONE NOT NULL,
    first_acquire_time TIMESTAMP WITH TIME ZONE,
    cancel_time TIMESTAMP WITH TIME ZONE,
    executor_rating FLOAT NOT NULL,

    UNIQUE (order_id)
);
