INSERT INTO groups (name) VALUES
    ('Muse'),
    ('The Beatles'),
    ('Queen'),
    ('Pink Floyd');

INSERT INTO songs (group_id, title, text, release_date, link) VALUES
  
    (1, 'Supermassive Black Hole', ' Ooh baby, don''t you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight', 
                      '16.07.2006', 'https://www.youtube.com/watch?v=Xsp3_a-PMTw'),
    (1, 'Uprising', 'Paranoia is in bloom\\nThe PR transmissions will resume\\nThey''ll try to push drugs that keep us all dumbed down\\nAnd hope that we will never see the truth around\\n\\nAnother promise, another seed\\nAnother packaged lie to keep us trapped in greed\\nAnd all the green belts wrapped around our minds\\nAnd endless red tape to keep the truth confined\\n', 
                      '14.09.2009', 'https://www.youtube.com/watch?v=w8KQmps-Sog'),
    
    
    (2, 'Hey Jude', 'Hey Jude, don\''t make it bad.\n\nRemember to let her into your heart', 
                      '1968-08-26', 'http://example.com/hey_jude'),
    (2, 'Let It Be', 'When I find myself in times of trouble.\n\nWhisper words of wisdom, let it be', 
                      '1970-05-08', 'http://example.com/let_it_be'),
    
 
    (3, 'Bohemian Rhapsody', 'Is this the real life? Is this just fantasy?\n\nCaught in a landslide, no escape from reality', 
                               '1975-10-31', 'http://example.com/bohemian_rhapsody'),
    (3, 'We Will Rock You', 'Buddy, you''re a boy, make a big noise.\n\nYou\''re gonna be a big man some day', 
                               '1977-10-07', 'http://example.com/we_will_rock_you'),
    
  
    (4, 'Comfortably Numb', 'Hello, is there anybody in there?\n\nJust nod if you can hear me', 
                               '1979-11-30', 'http://example.com/comfortably_numb'),
    (4, 'Another Brick in the Wall', 'We don''t need no education.\n\nWe don''t need no thought control', 
                                      '1979-11-30', 'http://example.com/another_brick_in_the_wall');
