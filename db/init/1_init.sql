DROP TABLE IF EXISTS docs;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS bundleddocs;
DROP TABLE IF EXISTS bundles;

CREATE TABLE IF NOT EXISTS bundles
(
   id SERIAL NOT NULL,
   title TEXT NOT NULL,
   user_id INTEGER NOT NULL,
   created_at TIMESTAMP NULL,
   updated_at TIMESTAMP NULL,
   deleted_at TIMESTAMP NULL,
   PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS bundleddocs
(
   id SERIAL NOT NULL,
   bundle_id INTEGER NOT NULL,
   doc_id INTEGER NOT NULL,
   created_at TIMESTAMP NULL,
   updated_at TIMESTAMP NULL,
   deleted_at TIMESTAMP NULL,
   PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS docs
(
   id SERIAL NOT NULL,
   title TEXT NOT NULL,
   text TEXT NOT NULL,
   user_id INTEGER NOT NULL,
   created_at TIMESTAMP NULL,
   updated_at TIMESTAMP NULL,
   deleted_at TIMESTAMP NULL,
   PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users
(
   id SERIAL NOT NULL,
   name TEXT NOT NULL,
   email TEXT UNIQUE NOT NULL,
   password TEXT NOT NULL,
   created_at TIMESTAMP NULL,
   updated_at TIMESTAMP NULL,
   deleted_at TIMESTAMP NULL,
   PRIMARY KEY (id)
);

INSERT INTO users(name, email, password) VALUES ('taro', 'a@a', 'aaaaaaaaaaa');
INSERT INTO docs(title, text, user_id) VALUES ('main.go', 'package\tmain\nfunc main(){\nprintln("hello")\n}', 1);
INSERT INTO docs(title, text, user_id) VALUES ('main1.go', 'package\tmain\nfunc main(){\nprintln("hello1")\n}', 1);
INSERT INTO docs(title, text, user_id) VALUES ('main2.go', 'package\tmain\nfunc main(){\nprintln("hello2")\n}', 1);

INSERT INTO bundles(title, user_id) VALUES ('goMainPackage', 1);

INSERT INTO bundleddocs(bundle_id, doc_id) VALUES (1, 1);
INSERT INTO bundleddocs(bundle_id, doc_id) VALUES (1, 2);
INSERT INTO bundleddocs(bundle_id, doc_id) VALUES (1, 3);