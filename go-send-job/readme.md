Currently I just register a firebase app
And in golang code, I can send a notification via this firebase app, but cannot testing with real token of user



- With send service, I have 2 solution for implement:
  1. When new job is created, we call api directly to send service to send notification
  2. Use 'oplogs', I think this way is better than the the first way, We use this to get notified when new job was created, and after that, we can implement a message queue to produce and consume notification

I will learn about this and implement in the future.