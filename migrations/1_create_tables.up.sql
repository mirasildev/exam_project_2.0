CREATE TABLE "users"(
    "id" SERIAL PRIMARY KEY,
    "first_name" VARCHAR(30) NOT NULL,
    "last_name" VARCHAR(30) NOT NULL,
    "email" VARCHAR(50) NOT NULL UNIQUE,
    "password" VARCHAR NOT NULL,
    "phone_number" VARCHAR(20) UNIQUE,
    "type" VARCHAR CHECK ("type" IN('user', 'partner', 'superadmin')) NOT NULL DEFAULT 'user',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

CREATE TABLE "hotels"(
    "id" SERIAL PRIMARY KEY,
    "hotel_name" VARCHAR(40) NOT NULL,
    "description" VARCHAR(100) ,
    "address" VARCHAR(255) NOT NULL,
    "image_url" VARCHAR(255) NOT NULL,
    "num_of_rooms" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL
);

CREATE TABLE "hotel_images"(
    "id" SERIAL PRIMARY KEY,
    "hotel_id" INTEGER NOT NULL,
    "image_url" VARCHAR(255),
    "sequence_number" INTEGER NOT NULL
);

CREATE TABLE "rooms"(
    "id" SERIAL PRIMARY KEY,
    "room_number" INTEGER NOT NULL,
    "type" VARCHAR(255) CHECK ("type" IN('single', 'double', 'family')) NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "hotel_id" INTEGER NOT NULL,
    "price_per_night" NUMERIC(18, 2) NOT NULL,
    "status" BOOLEAN NOT NULL
);

CREATE TABLE "bookings"(
    "id" SERIAL PRIMARY KEY,
    "arrival" DATE NOT NULL,
    "checkout" DATE NOT NULL,
    "room_id" INTEGER NOT NULL,
    "room_number" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "booked_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE
    "rooms" ADD CONSTRAINT "rooms_hotel_id_foreign" FOREIGN KEY("hotel_id") REFERENCES "hotels"("id");
ALTER TABLE
    "hotel_images" ADD CONSTRAINT "hotel_images_hotel_id_foreign" FOREIGN KEY("hotel_id") REFERENCES "hotels"("id");
ALTER TABLE
    "hotels" ADD CONSTRAINT "hotels_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "bookings" ADD CONSTRAINT "bookings_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "bookings" ADD CONSTRAINT "bookings_room_id_foreign" FOREIGN KEY("room_id") REFERENCES "rooms"("id");