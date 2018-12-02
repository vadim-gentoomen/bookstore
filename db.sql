CREATE TABLE books
(
  isbn   char(14)     NOT NULL,
  title  varchar(255) NOT NULL,
  author varchar(255) NOT NULL,
  price  decimal(5,2) NOT NULL
);

INSERT INTO books (isbn, title, author, price)
VALUES
  ('978-1-78712-349-6', 'Go: Building Web Applications', 'Nathan Kozyra, Nathan Kozyra', 0.1),
  ('978-5-97060-477-9', 'GO на практике', 'Мэтт Батчер, Мэтт Фарина', 0.1),
  ('978-5-94074-854-0', 'Программирование на Go', 'Марк Саммерфильд', 0.1);

ALTER TABLE books
  ADD PRIMARY KEY (isbn);