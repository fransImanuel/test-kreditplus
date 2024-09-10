
# Test kreditplus

test kreditplus


## Explaination

The Project I did only gave general outline of the solution ( since there is no specific instruction ) 

Regarding The Minimum Requirement 

- Point 3
I use pessimistic Locking for handling concurrent transaction

- Point 4
Adopt min 3 from 10 OWASP

-- SQL Injection Prevention: GORM Framework already guard our code from sql injection.

-- Input Validation: Input Validation already managed by Gin Struct tag validation.

-- Authentication: I add JWT Middleware for every endpoint