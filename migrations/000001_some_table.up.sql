CREATE TABLE if not exists "categories"(
    "id" serial PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE default current_timestamp
);

CREATE TABLE if not exists "users"(
    "id" serial PRIMARY KEY,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255),
    "phone_number" VARCHAR(255) UNIQUE,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "created_at" TIMESTAMP WITH TIME ZONE default current_timestamp,
    "gender" VARCHAR(255) CHECK("gender" IN('male','female')) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "username" VARCHAR(255) NOT NULL UNIQUE,
    "profile_image_url" VARCHAR(255),
    "type" VARCHAR(255)CHECK("type" IN('admin','user')) NOT NULL
);

CREATE TABLE if not exists "posts"(
    "id" serial PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL,
    "description" TEXT NOT NULL,
    "image_url" VARCHAR(255) NOT NULL,
    "user_id" INTEGER NOT NULL REFERENCES users(id)ON DELETE CASCADE,
    "category_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "views_count" INTEGER not NULL default 0
);

CREATE TABLE "comments"(
    "id" serial PRIMARY KEY,
    "post_id" INTEGER NOT NULL REFERENCES posts(id)ON DELETE CASCADE,
    "user_id" INTEGER NOT NULL REFERENCES users(id)ON DELETE CASCADE,
    "description" TEXT,
    "created_at" TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITHOUT TIME ZONE
);


CREATE TABLE "likes"(
    "id" serial PRIMARY KEY,
    "post_id" INTEGER NOT NULL REFERENCES posts(id)ON DELETE CASCADE,
    "user_id" INTEGER NOT NULL REFERENCES users(id)ON DELETE CASCADE,
    "status" VARCHAR(255) CHECK("status" IN('like','dislike')) NOT NULL
);



