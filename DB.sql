CREATE TABLE "users" (
  "id" serial not null unique,
  "name" varchar(255) not null
);

CREATE TABLE "categories" (
  "id" serial not null unique,
  "label" varchar(255) not null
);

CREATE TABLE "operations" (
  "id" serial not null unique,
  "user_id" serial not null,
  "month_date" date not null,
  "category_id" serial not null,
  "total" int not null
);

ALTER TABLE "operations" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "operations" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

INSERT INTO "categories" (label) values ('Аренда квартиры');
INSERT INTO "categories" (label) values ('Одежда и обувь');
INSERT INTO "categories" (label) values ('Супермаркеты');
INSERT INTO "categories" (label) values ('Транспорт');
INSERT INTO "categories" (label) values ('Остальное');
INSERT INTO "categories" (label) values ('Подписки, мобильная связь, интернет');
INSERT INTO "categories" (label) values ('Дом, ремонт');
INSERT INTO "categories" (label) values ('Подарки');
INSERT INTO "categories" (label) values ('Красота');
INSERT INTO "categories" (label) values ('Здоровье, аптека');
INSERT INTO "categories" (label) values ('Развлечения');
INSERT INTO "categories" (label) values ('Путешествие');
INSERT INTO "categories" (label) values ('Крупные покупки');
INSERT INTO "categories" (label) values ('Помощь');

INSERT INTO "users" (name) values ('Константин');
INSERT INTO "users" (name) values ('Дарья');

INSERT INTO "operations" (user_id, month_date, category_id, total) values (1, '2023-01-21', 1, 45000);
