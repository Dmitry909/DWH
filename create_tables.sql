-- ODS: assigned_orders

-- DDS

CREATE TABLE executors (
    executor_id TEXT PRIMARY KEYS,
);

CREATE TABLE zones (
    zone_id   TEXT PRIMARY KEYS,
    zone_name TEXT,
);

-- orders = assigned_orders

CREATE TABLE marts (
    ts                                            TIMESTAMP NOT NULL,
    sum_cost_finished                             INT NOT NULL
    median_cost_finished_and_cancelled            FLOAT NOT NULL
    median_acquire_time                           FLOAT NOT NULL
    number_different_executors_of_finished_orders INT NOT NULL,
    number_different_executors_of_all_orders      INT NOT NULL
    correlation_rating_status                     FLOAT NOT NULL
    correlation_final_coin_amount_status          FLOAT NOT NULL
);

