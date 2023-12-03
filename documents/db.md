## Database 


### bookings: Information about job
| Field           | Type         | Settings |                 Description                      | Example | Default & Notes |
|-----------------|----------|----------|----------------------------------------|------|---|
| _id (PK)       | objectId      |  PK, not null | Id of job                           | 656cb1f18cad4a1a6f233e7a       | |
| customer_id          | string |  not null | Id of job creator  |  123   | |
| booking_detail | object      |   | Detail information of job      |     | |
| status | string   | not null  | Status of job     | Waiting    | |
| booking_detail_id | objectId   | not null  | Id of detail information of job     | 656cb1f08cad4a1a6f233e79    |  |
| domestic_worker_id | Array[string]   |   | List of chosen domestic worker      | ["userId1", "userId2"]    |  |
| created_by     | string  |   | Create record creator     | userId | 
| created_at     | time  |   | Record time     | "2023-12-02T17:00:00.000+00:00" | 
| updated_by     | string  |   | Record updator     | userId | 
| updated_at     | time  |  | Update record time     | "2023-12-02T17:00:00.000+00:00" | 
### booking_details: Detail information of job
| Field           | Type    | Settings | Description               | Example        |  Default & Notes |
|-----------------|--------|--------|---------------------------|------|----------|
| _id (PK)      | objectId | PK, not null | Id of detail information | 656cb1f08cad4a1a6f233e79              | |
| booking_type     | string  | not null  | Type of job     | DD | DD - Dọn dẹp |
| booking_service_time     | object  | not null  | Time of job     | {"start_time": "2023-12-02T17:00:00.000+00:00" , "duration": 2} | 
| detail_service     | object  | null  | Bonus information     | {"a": "b"} | 
| created_by     | string  |   | Record creator     | userId | 
| created_at     | time  |   | Record time     | "2023-12-02T17:00:00.000+00:00" | 
| updated_by     | string  |   | Record updator     | userId | 
| updated_at     | time  |   | Update record time     | "2023-12-02T17:00:00.000+00:00" | 


### prices: Price information
| Field           | Type    | Settings | Description               | Example        |  Default & Notes |
|-----------------|--------|--------|---------------------------|------|----------|
| _id (PK)      | objectId | PK, not null | Id of price | 656cb1f08cad4a1a6f233e79              | |
| date     | string  | not null  | Date do job     | 2023-12-11 | "0000-00-00" (default) |
| booking_type     | string  | not null  | Type of job     | DD | 
| unit_price_id     | string  | not null  | Id of unit price    | 656cb1f08cad4a1a6f233e78  | 

### unit_prices: Unit price information
| Field           | Type    | Settings | Description               | Example        |  Default & Notes |
|-----------------|--------|--------|---------------------------|------|----------|
| _id (PK)      | objectId | PK, not null | Id of unit price | 656cb1f08cad4a1a6f233e79              | |
| duration     | int  | not null  | Duration unit     | 3 |  |
| price     | int  | not null   | Price    | 50000 | 

### notifications: Notifications record
| Field           | Type    | Settings | Description               | Example        |  Default & Notes |
|-----------------|--------|--------|---------------------------|------|----------|
| _id (PK)      | objectId | PK, not null | Id of notification| 656cb1f08cad4a1a6f233e79              | |
| job_id     | objectId  | not null  | Id of job     | 656cb1f08cad4a1a6f233e78 |  |
| job_detail     | object  | not null   | Detail information of job    |  | 
| user_id     | array[string]  | not null   | List of user receive job noti    | ["userId1", "userId2"] | 
| device_token     | array[string]  | not null   | List of device token receive job noti    | ["token1", "token2"] | 
| created_by     | string  |   | Record creator     | userId | 
| created_at     | time  |   | Create record time     | "2023-12-02T17:00:00.000+00:00" | 
| updated_by     | string  |   | Record updator     | userId | 
| updated_at     | time  |   | Update record time     | "2023-12-02T17:00:00.000+00:00" | 

### device_tokens: Information of device token and user
| Field           | Type    | Settings | Description               | Example        |  Default & Notes |
|-----------------|--------|--------|---------------------------|------|----------|
| _id (PK)      | objectId | PK, not null | Id of device token | 656cb1f08cad4a1a6f233e79              | |
| user_id     | string  | not null  | Id of user     | "userId" |  |
| token     | string  | not null   | device token of user    | "token" | 