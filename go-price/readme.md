With service to calculate price based on date and time.
My solution is querying the price and unit price in database based on 'bookingType' and 'date' param.
If that date exist in database, we will load this price and unit price to calculate, and if not exist, we will choose the default value for price.

And in the future, we can design some api and algorithm to create records for price information via date of job and type of job.