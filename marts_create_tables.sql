CREATE TYPE order_execution_status AS ENUM ('assigned', 'acquired', 'cancelled', 'completed');

CREATE TABLE general_analytics_01 (
    assign_time TIMESTAMP NOT NULL,
    final_coin_amount DOUBLE PRECISION NOT NULL,
    execution_status order_execution_status NOT NULL,
    acquire_seconds INT,
    executor_rating FLOAT NOT NULL,
    order_id BIGINT NOT NULL,
    executor_id BIGINT NOT NULL,
    zone_id BIGINT NOT NULL,
    price_for_toll_road DOUBLE PRECISION NOT NULL
);
