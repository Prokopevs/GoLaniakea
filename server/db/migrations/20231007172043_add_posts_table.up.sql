CREATE TABLE "posts" (
    id serial primary key,
    image_url varchar NOT NULL,
    name varchar NOT NULL,
    description varchar NOT NULL,
    date varchar NOT NULL,
    category int NOT NULL,
    category_name varchar NOT NULL,
    like_count int NOT NULL,
    text varchar NOT NULL
);


