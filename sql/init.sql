BEGIN;

DROP TABLE IF EXISTS setting;
CREATE TABLE setting (
  user_id INTEGER GENERATED ALWAYS AS IDENTITY,
  theme TEXT NOT NULL DEFAULT 'light',
  
  PRIMARY KEY (user_id)
);

INSERT INTO setting (theme) VALUES ('light');

DROP TABLE IF EXISTS dog;
CREATE TABLE dog (
  dog_id INTEGER GENERATED ALWAYS AS IDENTITY,
  colour TEXT NOT NULL,
  breed TEXT NOT NULL,
  name TEXT NOT NULL,
  PRIMARY KEY (dog_id)
);

INSERT INTO dog (colour, breed, name) VALUES
  ('black', 'cocker spaniel', 'Banjo'),
  ('black', 'border collie', 'Sebastian'),
  ('tan', 'border collie', 'Marcus'),
  ('gold', 'labrador', 'Fido')
;

DROP TABLE IF EXISTS theme;
CREATE TABLE theme (
  theme_id INTEGER GENERATED ALWAYS AS IDENTITY,
  label TEXT NOT NULL,
  value TEXT NOT NULL,
  PRIMARY KEY (theme_id)
);

INSERT INTO theme (label, value) VALUES
  ('Light', 'light'),
  ('Dark', 'dark'),
  ('Cupcake', 'cupcake'),
  ('Bumblebee', 'bumblebee'),
  ('Emerald', 'emerald'),
  ('Corporate', 'corporate'),
  ('Synthwave', 'synthwave'),
  ('Retro', 'retro'),
  ('Cyberpunk', 'cyberpunk'),
  ('Valentine', 'valentine'),
  ('Halloween', 'halloween'),
  ('Garden', 'garden'),
  ('Forest', 'forest'),
  ('Aqua', 'aqua'),
  ('Lofi', 'lofi'),
  ('Pastel', 'pastel'),
  ('Fantasy', 'fantasy'),
  ('Wireframe', 'wireframe'),
  ('Black', 'black'),
  ('Luxury', 'luxury'),
  ('Dracula', 'dracula'),
  ('Cmyk', 'cmyk'),
  ('Autumn', 'autumn'),
  ('Business', 'business'),
  ('Acid', 'acid'),
  ('Lemonade', 'lemonade'),
  ('Night', 'night'),
  ('Coffee', 'coffee'),
  ('Winter', 'winter'),
  ('Dim', 'dim'),
  ('Nord', 'nord'),
  ('Sunset', 'sunset'),
  ('Caramellatte', 'caramellatte'),
  ('Abyss', 'abyss'),
  ('Silk', 'silk')
;

DROP TABLE IF EXISTS word;
CREATE TABLE word (
  word_id INTEGER GENERATED ALWAYS AS IDENTITY,
  word TEXT NOT NULL,
  PRIMARY KEY (word_id)
);

INSERT INTO word (word) VALUES
  ('domineering'),
  ('rag'),
  ('burrowstown'),
  ('unown'),
  ('uncalcined'),
  ('narwhal'),
  ('dragonfly'),
  ('hold'),
  ('pyrrhotite'),
  ('buttocked'),
  ('octogenarianism'),
  ('periwinkled'),
  ('sjambok'),
  ('decadal'),
  ('blooddrop'),
  ('consulship'),
  ('glycosine'),
  ('Teutonize'),
  ('Sybil'),
  ('eupolyzoan'),
  ('sibilatory'),
  ('Carica'),
  ('matronly'),
  ('interstellary'),
  ('siphonoplax'),
  ('tyrannicly'),
  ('Fourierite'),
  ('Pop'),
  ('slagging'),
  ('kokan'),
  ('shelteringly'),
  ('endocystitis'),
  ('chordoid'),
  ('angiostegnosis'),
  ('tuberculid'),
  ('bouillabaisse'),
  ('batsman'),
  ('contracture'),
  ('oversmoothly'),
  ('trachelomastoid'),
  ('orbitofrontal'),
  ('retinerved'),
  ('fadable'),
  ('chromoplasm'),
  ('pericardotomy'),
  ('complex'),
  ('thoracic'),
  ('allude'),
  ('preinitiate'),
  ('Eogaea')
;

COMMIT;
