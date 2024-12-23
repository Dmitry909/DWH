CREATE TABLE executors (
    executor_id VARCHAR(255) PRIMARY KEY
);

CREATE TABLE zones (
    zone_name VARCHAR(255) PRIMARY KEY
);

CREATE TYPE completed_or_cancelled_status AS ENUM ('cancelled', 'completed');

CREATE TABLE completed_and_cancelled_orders (
    order_id VARCHAR(255) NOT NULL,
    executor_id VARCHAR(255) NOT NULL,
    execution_status completed_or_cancelled_status NOT NULL,
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

CREATE TYPE status_in_progress AS ENUM ('assigned', 'acquired', 'cancelled', 'completed');

CREATE TABLE orders_in_progress (
    order_id VARCHAR(255) NOT NULL,
    executor_id VARCHAR(255) NOT NULL,
    execution_status status_in_progress NOT NULL,
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
