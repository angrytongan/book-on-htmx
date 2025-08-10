DROP TABLE IF EXISTS settings;
CREATE TABLE settings (
  id INTEGER PRIMARY KEY,
  theme TEXT NOT NULL DEFAULT 'light'
);

INSERT INTO settings (id, theme) VALUES (1, 'light');

DROP TABLE IF EXISTS themes;
CREATE TABLE themes (
  id INTEGER PRIMARY KEY,
  label TEXT NOT NULL,
  value TEXT NOT NULL
);

INSERT INTO themes (label, value) VALUES
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
