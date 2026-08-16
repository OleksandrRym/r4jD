
CREATE TABLE foods (
                       uuid UUID PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       descriptions VARCHAR(255),
                       protein DOUBLE PRECISION NOT NULL,
                       fat DOUBLE PRECISION NOT NULL,
                       carbohydrates DOUBLE PRECISION NOT NULL,
                       calories DOUBLE PRECISION NOT NULL
);