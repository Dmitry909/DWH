### Описание схем баз данных

#### ODS

Содержит одну таблицу `assigned_orders`. Её колонки:
- `assigned_order_id` -- уникальный id в пределах этой таблицы.
- `order_id` -- id заказа. Приходит в сервис заказов извне.
- `executor_id` -- id исполнителя. Приходит в сервис заказов извне.
- `customer_id` -- id заказчика. Приходит в сервис заказов извне.
- `execution_status` -- статус исполненности заказа. Приходит в сервис заказов извне.
- `coin_coefficient` -- коэффициент стоимости, зависящий от зоны.
- `coin_bonus_amount` -- стоимость платных дорог.
- `final_coin_amount` -- финальная стоимость заказа.
- `zone_id` -- числовой id зоны заказа.
- `has_executor_fallback_been_used` -- использовался ли fallback-источник executor.
- `assign_time` -- таймстемп назначения заказа.
- `first_acquire_time` -- таймстемп назначения заказа исполнителю.
- `executor_rating` -- рейтинг исполнителя в момент назначения ему заказа.
- `completed_time` -- таймстемп завершения заказа.

#### DDS

Таблица `executors`. Колонки:
- `executor_id`
То есть список всех исполнителей.

Таблица `customers`:
- `customer_id`
То есть список всех заказчиков.

Таблица `zones`. Колонки:
- `zone_id`
- `zone_name`
То есть список всех зон.

Таблица `orders`:
- `order_id`
- `executor_id`
- `customer_id`
- `execution_status`
- `coin_coefficient`
- `coin_bonus_amount`
- `final_coin_amount`
- `zone_id`
- `assign_time`
- `first_acquire_time`
- `executor_rating`
- `completed_time`
Отличия от `assigned_orders` из слоя ODS: убраны `assigned_order_id`, `has_executor_fallback_been_used`.

#### marts

Таблица `general_analytics_01`
- `assign_time`
- `final_coin_amount`
- `is_cancelled` -- отменен ли заказ.
- `acquire_seconds` -- время в секундах от создания заказа до назначения его какому-то исполнителю
- `executor_rating`
- `order_id`
- `executor_id`
- `zone_id`
- `price_for_toll_road` -- финальная стоимость заказа.
Отличия от `orders` -- убраны ненужные на дашборде поля и посчитано `acquire_seconds`.
