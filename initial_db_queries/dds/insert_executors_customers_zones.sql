INSERT INTO executors (executor_id)
SELECT executor_id FROM dblink('host=rc1b-xy6y7apt7j7jdpc7.mdb.yandexcloud.net,rc1d-opy4a78yulgzu7z2.mdb.yandexcloud.net dbname=db1 user=user1 password=NgdXRLUNn67d8tR port=6432 sslmode=verify-full target_session_attrs=read-write',
                    'SELECT DISTINCT executor_id FROM assigned_orders WHERE executor_id IS NOT NULL') AS t(executor_id integer);

INSERT INTO customers (customer_id)
SELECT customer_id FROM dblink('host=rc1b-xy6y7apt7j7jdpc7.mdb.yandexcloud.net,rc1d-opy4a78yulgzu7z2.mdb.yandexcloud.net dbname=db1 user=user1 password=NgdXRLUNn67d8tR port=6432 sslmode=verify-full target_session_attrs=read-write',
                    'SELECT DISTINCT customer_id FROM assigned_orders WHERE customer_id IS NOT NULL') AS t(customer_id integer);

INSERT INTO zones (zone_id)
SELECT zone_id FROM dblink('host=rc1b-xy6y7apt7j7jdpc7.mdb.yandexcloud.net,rc1d-opy4a78yulgzu7z2.mdb.yandexcloud.net dbname=db1 user=user1 password=NgdXRLUNn67d8tR port=6432 sslmode=verify-full target_session_attrs=read-write',
                    'SELECT DISTINCT zone_id FROM assigned_orders WHERE zone_id IS NOT NULL') AS t(zone_id integer);
