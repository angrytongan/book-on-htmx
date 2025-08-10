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
  ('Retro', 'retro'),
  ('Cyberpunk', 'cyberpunk'),
  ('Valentine', 'valentine'),
  ('Aqua', 'aqua')
;
