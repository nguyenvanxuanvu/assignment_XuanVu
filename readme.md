# Assignment build APIs based on MicroServices architecture:
# Results:
### go-booking: Service create, push job to system
### go-price: Service calculate price of job
### go-send-job: Service send job (notification) to user

# Database design
### ./documents/db.md

# API document:
https://documenter.getpostman.com/view/19892884/2s9YeK5Ai9



# Step to run source code
I built docker compose file and dockerfile, it can started services but cannot connect to mongodb, I still find solution to fix this.
So for mannual run the source code, we can do below steps:
- MongoDB:
    - Create collections: bookings, bookings_detail, device_tokens, notifications, prices, unit prices
    - Insert data for 3 collections: device_tokens, prices, unit_prices with data in folder 'data'
- Go services (from root folder):
  - go-booking:
    - cd go-booking
    - go mod tidy
    - go run main.go
  - go-price:
    - cd go-price
    - go mod tidy
    - go run main.go
  - go-send-job:
    - cd go-send-job
    - go mod tidy
    - go run main.go
- Note: With each service, some of my solutions or my stucks was described in readme file of each sub folder.

- With this assignment, it was an opportunity for me to self-reflect and reevaluate my own knowledge. I only had less than 2 days to make them, so the testing part as well as some parts were not completed yet. And I think if I had more time, I could complete what I set out to do better and learn more from them. 
- Thanks you!!