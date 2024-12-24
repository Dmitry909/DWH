CREATE TYPE order_execution_status AS ENUM ('assigned', 'acquired', 'cancelled', 'completed');

CREATE TABLE assigned_orders (
    assigned_order_id BIGINT NOT NULL,
    order_id BIGINT NOT NULL,
    executor_id BIGINT NOT NULL,
    execution_status order_execution_status NOT NULL,
    coin_coefficient DOUBLE PRECISION NOT NULL,
    coin_bonus_amount DOUBLE PRECISION NOT NULL,
    final_coin_amount DOUBLE PRECISION NOT NULL,
    zone_name VARCHAR(255),
    has_executor_fallback_been_used BOOLEAN NOT NULL,
    assign_time TIMESTAMP WITH TIME ZONE NOT NULL,
    last_acquire_time TIMESTAMP WITH TIME ZONE,

    PRIMARY KEY (assigned_order_id),
    UNIQUE (order_id)
);