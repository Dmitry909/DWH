CREATE TYPE order_execution_status AS ENUM ('assigned', 'acquired', 'cancelled', 'completed');

CREATE TABLE marts (
    order_id BIGINT PRIMARY KEY,
    assign_time TIMESTAMP NOT NULL,
    final_coin_amount DOUBLE PRECISION NOT NULL,
    execution_status order_execution_status NOT NULL,
    acquire_seconds INT,
    executor_rating FLOAT NOT NULL
);
