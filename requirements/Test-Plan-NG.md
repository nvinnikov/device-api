# Тест план

## Расширенный тест-план

## ActDeviceApiService

### Тест 1 Регистрация устройства GRPC

Предварительные условия:

1) Система запущена.
2) Проходит Readiness probe

| Шаг | Действие                                                                                                            | Ожидаемый результат                                                                                                                                                                                                                                                                                                      |
|-----|---------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 1.  | Отправьте запрос CreateDeviceV1Request со следующими параметрами:<br/> `user_id = {user_id}, platform = {platform}` | Пришел ответ CreateDeviceV1Response с deviceId > 0 (`{device_ID_1}`)<br/> В БД появилась 1 запись: `user_id = {user_id}, platform = {platform}, entered_at >= Текущее время (**timestamp_1**), removed =false, created_at >= entered_at, updated_at >= created_at`. Отправлено уведомление о создании устройства в Kafka |
| 2.  | Выполните запрос: DescribeDeviceV1Request c device_id =**device_ID_1**                                              | Получен ответ DescribeDeviceV1Response с Device =`id = {device_ID_1}, user_id = {user_id}, platform = {platform}, etered_at = ± **timestamp_1** `}                                                                                                                                                                       |

#### Тестовые данные:

| platform      | user_id | Ожидаемый результат |
|---------------|---------|---------------------|
| ios           | 2       | &check;             |
| android       | 3       | &check;             |
| windows_phone | -1      | &cross;             |
| ""            | Null    | &cross;             |


### Тест 2 Редактирование устройства GRPC

Предварительные условия:

1) Система запущена.
2) Проходит Readiness probe

| Шаг | Действие                                                                                                                                         | Ожидаемый результат                                                                                                                                                                                                                                                |
|-----|--------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 1.  | Отправьте запрос CreateDeviceV1Request со следующими параметрами:<br/> `user_id = {user_id}, platform = {platform}`                              | Пришел ответ CreateDeviceV1Response с deviceId > 0 (`{device_ID_1}`)<br/> В БД появилась 1 запись: `user_id = {user_id}, platform = {platform}, entered_at >= Текущее время (**timestamp_1**), removed =false, created_at >= entered_at, updated_at >= created_at` |
| 2.  | Выполните запрос: DescribeDeviceV1Request c device_id =**device_ID_1**                                                                           | Получен ответ DescribeDeviceV1Response с Device =`id = **device_ID_1**, user_id = {user_id}, platform = {platform}, etered_at = ± **timestamp_1** `}                                                                                                               |
| 3.  | Отправьте запрос: UpdateDeviceV1Request со следующими параметрами:<br/> `device_id ={device_id}, platform = {new_platform}, user_id = {user_id}` | * Пришел ответ UpdateDeviceV1Response с `success ={True}`"<br/> * В БД обновилась запись: `user_id = {user_id}, platform = {new_platform}, entered_at >= (**timestamp_1**), removed =false, created_at >= entered_at, updated_at >= created_at`                    |
| 4.  | Выполните запрос: DescribeDeviceV1Request c device_id =**device_ID_1**                                                                           | Получен ответ DescribeDeviceV1Response с Device =`id = **device_ID_1**, user_id = {user_id}, platform = {new_platform}, etered_at = ± **timestamp_1** `}                                                                                                           |

#### Тестовые данные:

| platform      | user_id | new_platform                      | Ожидаемый результат |
|---------------|---------|-----------------------------------|---------------------|
| ios           | 4       | android                           | &check;             |
| android       | 5       | ios                               | &check;             |
| windows_phone | -1      | ios                               | &cross;             |
| ""            | Null    | macos                             | &cross;             |
| android       | 5       |                                   | &cross;             |
| ios           | 4       | androoidandrooidandrooidandrooidz | &cross;             |

### Тест 3 Удаление устройства GRPC

Предварительные условия:

1) Система запущена.
2) Проходит Readiness probe

| Шаг | Действие                                                                                                            | Ожидаемый результат                                                                                                                                                                                                                                                |
|-----|---------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 1.  | Отправьте запрос CreateDeviceV1Request со следующими параметрами:<br/> `user_id = {user_id}, platform = {platform}` | Пришел ответ CreateDeviceV1Response с deviceId > 0 (`{device_ID_1}`)<br/> В БД появилась 1 запись: `user_id = {user_id}, platform = {platform}, entered_at >= Текущее время (**timestamp_1**), removed =false, created_at >= entered_at, updated_at >= created_at` |
| 2.  | Выполните запрос: DescribeDeviceV1Request c device_id =**device_ID_1**                                              | Получен ответ DescribeDeviceV1Response с Device =`id = **device_ID_1**, user_id = {user_id}, platform = {platform}, etered_at = ± **timestamp_1** `}                                                                                                               |
| 3.  | Отправьте запрос: RemoveDeviceV1Request со следующими параметрами:<br/> `device_id ={device_id}`                    | Пришел ответ RemoveDeviceV1Response с `found ={True}`"<br/> * В БД обновилась запись, добавилось поле `removed ={True}`                                                                                                                                            |
| 4.  | Выполните запрос: DescribeDeviceV1Request c device_id =**device_ID_1**                                              | Получен ответ DescribeDeviceV1Response с кодом `404` и `code= {5}, message = {device not found}`                                                                                                                                                                   |

#### Тестовые данные:

| platform | user_id | new_platform | Ожидаемый результат |
|----------|---------|--------------|---------------------|
| ios      | 6       | android      | &check;             |
| android  | 7       | ios          | &check;             |
| ""       | Null    | Null         | &cross;             |

### Тест 4 Список устройств GRPC

Предварительные условия:

1) Система запущена.
2) Проходит Readiness probe

| Шаг | Действие                                                                                                     | Ожидаемый результат                                                                                                                                                 |
|-----|--------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 1.  | Отправьте запрос ListDevicesV1Request со следующими параметрами:<br/> `page = {page}, per_page = {per_page}` | Пришел ответ ListDevicesV1Response с items соответсвенно с коллекцией из БД `id = {id}, user_id = {user_id}, platform = {platform}, etered_at = ± **timestamp_1** ` |

#### Тестовые данные:

| page | per_page | Ожидаемый результат |
|------|----------|---------------------|
| 0    | 1        | &check;             |
| 1    | 5        | &check;             |
| -1   | 0        | &cross;             |
| qwe  | qwe      | &cross;             |


## ActNotificationApiService

### Тест 5 Отправка уведомлений по устройству GRPC

Предварительные условия:

1) Система запущена.
2) Проходит Readiness probe

| Шаг | Действие                                                                                                                                                                                                                                 | Ожидаемый результат                                                            |
|-----|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------|
| 1.  | Зарегистрировать устрово, выполнить тест кейс `Тест 1 Регистрация устройства GRPC`, получение `{device_ID_1}`                                                                                                                            | Устройство зарегистрировано, Тест кейс прошел успешно, получен `{device_ID_1}` |
| 2.  | Отправить запрос SendNotificationV1Request со следующими параметрами:<br/> `Notification.{notificationId={notificationId}, deviceId={deviceId}, message={message}, lang={Language.lang}, notificationStatus{Status.notificationStatus}}` | Получен ответ SendNotificationV1Response `SendNotificationV1Response`          |

#### Тестовые данные:

| Language     | Status             | Time  |
|--------------|--------------------|-------|
| LANG_ENGLISH | STATUS_CREATED     | 6:01  |
| LANG_RUSSIAN | STATUS_IN_PROGRESS | 11:01 |
| LANG_ESPANOL | STATUS_DELIVERED   | 15:01 |
| LANG_ITALIAN |                    | 21:01 |


#### Ожидаемый результат

Language

| Поле         | Значение в БД | Описание    |
|--------------|---------------|-------------|
| LANG_ENGLISH | 1             | Английский  |
| LANG_RUSSIAN | 2             | Русский     |
| LANG_ESPANOL | 3             | Испанский   |
| LANG_ITALIAN | 4             | Итальянский |

Status

| Поле               | Значение в БД | Описание                        |
|--------------------|---------------|---------------------------------|
| STATUS_CREATED     | 1             | Уведомление создано             |
| STATUS_IN_PROGRESS | 2             | Уведомление в процессе отправки |
| STATUS_DELIVERED   | 3             | Уведомление доставлено          |


Приветствия - Утро

| Язык         | Текст        |
|--------------|--------------|
| LANG_ENGLISH | Good morning |
| LANG_RUSSIAN | Доброе утро  |
| LANG_ESPANOL | Buenos dias  |
| LANG_ITALIAN | Buon giorno  |

Приветствия - День

| Язык         | Текст           |
|--------------|-----------------|
| LANG_ENGLISH | Good afternoon  |
| LANG_RUSSIAN | Добрый день     |
| LANG_ESPANOL | Buenas tardes   |
| LANG_ITALIAN | Buon pomeriggio |


Приветствия - Вечер

| Язык         | Текст         |
|--------------|---------------|
| LANG_ENGLISH | Good evening  |
| LANG_RUSSIAN | Добрый вечер  |
| LANG_ESPANOL | Buenas noches |
| LANG_ITALIAN | Buona serata  |

Приветствия - Ночь

| Язык         | Текст         |
|--------------|---------------|
| LANG_ENGLISH | Good night    |
| LANG_RUSSIAN | Доброй ночи   |
| LANG_ESPANOL | Buenas noches |
| LANG_ITALIAN | Buona notte   |

### Тест 6 Получение и подтверждение уведомлений по устройству GRPC

Предварительные условия:

1) Система запущена.
2) Проходит Readiness probe

| Шаг | Действие                                                                                                                                                                                                                                 | Ожидаемый результат                                                                                         |
|-----|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------|
| 1.  | Зарегистрировать устрово, выполнить тест кейс `Тест 1 Регистрация устройства GRPC`, получение `{device_ID_1}`                                                                                                                            | Устройство зарегистрировано, Тест кейс прошел успешно, получен `{device_ID_1}`                              |
| 2.  | Отправить запрос SendNotificationV1Request со следующими параметрами:<br/> `Notification.{notificationId={notificationId}, deviceId={deviceId}, message={message}, lang={Language.lang}, notificationStatus{Status.notificationStatus}}` | Получен ответ SendNotificationV1Response `SendNotificationV1Response`                                       |
| 3.  | Отправить запрос GetNotificationV1Request со следующими параметрами:<br/> `deviceId={deviceId}`                                                                                                                                          | Получен ответ GetNotificationV1Response `notification.{notificationId={notificationId}, message={message}}` |
| 4.  | Отправить запрос AckNotificationV1Request со следующими параметрами:<br/> `notificationId={notificationId}`                                                                                                                              | Получен ответ AckNotificationV1Response `success={success}`                                                 |

#### Тестовые данные:

| Language     | Status             | Time  | message |
|--------------|--------------------|-------|---------|
| LANG_ENGLISH | STATUS_CREATED     | 6:01  | message |
| LANG_RUSSIAN | STATUS_IN_PROGRESS | 11:01 | None    |
| LANG_ESPANOL | STATUS_DELIVERED   | 15:01 |         |
| LANG_ITALIAN |                    | 21:01 |         |

#### Ожидаемый результат

Language

| Поле         | Значение в БД | Описание    |
|--------------|---------------|-------------|
| LANG_ENGLISH | 1             | Английский  |
| LANG_RUSSIAN | 2             | Русский     |
| LANG_ESPANOL | 3             | Испанский   |
| LANG_ITALIAN | 4             | Итальянский |

Status

| Поле               | Значение в БД | Описание                        |
|--------------------|---------------|---------------------------------|
| STATUS_CREATED     | 1             | Уведомление создано             |
| STATUS_IN_PROGRESS | 2             | Уведомление в процессе отправки |
| STATUS_DELIVERED   | 3             | Уведомление доставлено          |


Приветствия - Утро

| Язык         | Текст        |
|--------------|--------------|
| LANG_ENGLISH | Good morning |
| LANG_RUSSIAN | Доброе утро  |
| LANG_ESPANOL | Buenos dias  |
| LANG_ITALIAN | Buon giorno  |

Приветствия - День

| Язык         | Текст           |
|--------------|-----------------|
| LANG_ENGLISH | Good afternoon  |
| LANG_RUSSIAN | Добрый день     |
| LANG_ESPANOL | Buenas tardes   |
| LANG_ITALIAN | Buon pomeriggio |


Приветствия - Вечер

| Язык         | Текст         |
|--------------|---------------|
| LANG_ENGLISH | Good evening  |
| LANG_RUSSIAN | Добрый вечер  |
| LANG_ESPANOL | Buenas noches |
| LANG_ITALIAN | Buona serata  |

Приветствия - Ночь

| Язык         | Текст         |
|--------------|---------------|
| LANG_ENGLISH | Good night    |
| LANG_RUSSIAN | Доброй ночи   |
| LANG_ESPANOL | Buenas noches |
| LANG_ITALIAN | Buona notte   |

### Тест 7 Подписка на получение уведомлений

Предварительные условия:

1) Система запущена.
2) Проходит Readiness probe

| Шаг | Действие                                                                                                                                                                                                                                 | Ожидаемый результат                                                            |
|-----|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------|
| 1.  | Зарегистрировать устрово, выполнить тест кейс `Тест 1 Регистрация устройства GRPC`, получение `{device_ID_1}`                                                                                                                            | Устройство зарегистрировано, Тест кейс прошел успешно, получен `{device_ID_1}` |
| 2.  | Отправить запрос SubscribeNotificationRequest со следующими параметрами:<br/> `deviceId={device_ID_1}`                                                                                                                                   |                                                                                |
| 3.  | Отправить запрос SendNotificationV1Request со следующими параметрами:<br/> `Notification.{notificationId={notificationId}, deviceId={deviceId}, message={message}, lang={Language.lang}, notificationStatus{Status.notificationStatus}}` | Получен ответ SendNotificationV1Response `SendNotificationV1Response`          |
| 4.  | Отправить запрос SendNotificationV1Request со следующими параметрами:<br/> `Notification.{notificationId={notificationId}, deviceId={deviceId}, message={message}, lang={Language.lang}, notificationStatus{Status.notificationStatus}}` | Получен ответ SendNotificationV1Response `SendNotificationV1Response`          |
| 5.  | Отправить запрос SendNotificationV1Request со следующими параметрами:<br/> `Notification.{notificationId={notificationId}, deviceId={deviceId}, message={message}, lang={Language.lang}, notificationStatus{Status.notificationStatus}}` | Получен ответ SendNotificationV1Response `SendNotificationV1Response`          |
| 6.  | Отправить запрос AckNotificationV1Request со следующими параметрами:<br/> `notificationId={notificationId}`                                                                                                                              | Получен ответ AckNotificationV1Response `success={success}`                    |

#### Тестовые данные:

| Language     | Status             | Time  | message |
|--------------|--------------------|-------|---------|
| LANG_ENGLISH | STATUS_CREATED     | 6:01  | message |
| LANG_RUSSIAN | STATUS_IN_PROGRESS | 11:01 | None    |
| LANG_ESPANOL | STATUS_DELIVERED   | 15:01 |         |
| LANG_ITALIAN |                    | 21:01 |         |

#### Ожидаемый результат:
STATUS_CREATED переходит в STATUS_IN_PROGRESS в DB
Новые уведомления приходят сразу после отправки, записываются в DB STATUS_IN_PROGRESS
