CREATE TABLE groups
(
    id serial not null unique,
    name varchar(255) not null
   
);

CREATE TABLE songs
(
    id serial not null unique,
    group_id int references groups (id) on delete cascade not null,
    title varchar(255) not null,
    text text not null,
    release_date varchar(255) not null,
    link varchar(255) not null
);

INSERT INTO groups (name) VALUES 
('The Beatles'),
('Metallica'),
('Nirvana'),
('Muse');

INSERT INTO songs (group_id, title, text, release_date, link) VALUES
(1, 'Hey Jude', 'Hey Jude, don\'t make it bad...', '26.08.1968', 'https://example.com/heyjude'),
(1, 'Let It Be', 'When I find myself in times of trouble...', '06.03.1970', 'https://example.com/letitbe'),
(2, 'Enter Sandman', 'Exit light, enter night...', '29.07.1991', 'https://example.com/entersandman'),
(2, 'Nothing Else Matters', 'So close, no matter how far...', '20.03.1991', 'https://example.com/nothingelsematters'),
(3, 'Smells Like Teen Spirit', 'Load up on guns, bring your friends...', '10.09.1991', 'https://example.com/teenspirit'),
(4, 'Supermassive Black Hole', 'Ooh baby, don\'t you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight'
, '16.07.2006', 'https://www.youtube.com/watch?v=Xsp3_a-PMTw');