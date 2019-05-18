INSERT INTO company_fields (id, name) VALUES
('9e0ac61a-c0fa-46be-891d-c62b87a406cb', 'preduzetnistvo'),
('b03d5057-8369-4a4a-911b-9278f6c6dd01', 'poljoprivreda'),
('efaee74d-cf29-426d-b009-1b05d08e2d6e', 'vocarstvo'),
('70d69a58-554a-4aaa-8a96-74308c111338', 'fizika'),
('518f87fd-81ae-4df8-9a89-4151d6732a5c', 'blejarstvo'),
('1a933eb5-98ae-474e-a05d-e97afea593f7', 'gradjevina'),
('23863540-63dc-4ffb-b998-2a68ae637ba1', 'stocarstvo');

INSERT INTO tags VALUES
('b20c8070-c25e-4042-86a4-0d2cfc4aa505', 'beograd'),
('1d7292b6-814f-4b3d-9d8d-6e17123b56f6', 'novi sad'),
('f6649db4-6f18-406e-9fd0-2ea7e3a7e17d', 'centralna srbija'),
('93d612af-d256-4b08-a670-ddd162975ee3', 'mali preduzetnik'),
('697afef2-cb93-48db-8f05-772eb4deeda3', 'veliki preduzetnik'),
('1c9aaedf-ace3-49f0-b764-0911f52ab2b2', 'srednji preduzetnik'),
('d3fbce50-c7ea-4bf7-bfb2-fe8aef0d43c6', 'ngo'),
('e0fb765a-0aee-4ed8-b574-af5de8853a53', 'fizika');

INSERT INTO companies VALUES 
('67f8f522-c1f5-4e08-b9d4-8884a21935e2', 'ISWiB', '9e0ac61a-c0fa-46be-891d-c62b87a406cb', '{"b20c8070-c25e-4042-86a4-0d2cfc4aa505", "d3fbce50-c7ea-4bf7-bfb2-fe8aef0d43c6"}'),
('81adc65e-8455-4752-8b87-0c6c8fcb604a', 'Frutina d.o.o.', 'efaee74d-cf29-426d-b009-1b05d08e2d6e', '{"b20c8070-c25e-4042-86a4-0d2cfc4aa505", "93d612af-d256-4b08-a670-ddd162975ee3"}'),
('583d994e-d012-4b0c-956a-eebb1b80b034', 'Delta Holding', '70d69a58-554a-4aaa-8a96-74308c111338', '{"1d7292b6-814f-4b3d-9d8d-6e17123b56f6", "e0fb765a-0aee-4ed8-b574-af5de8853a53", "697afef2-cb93-48db-8f05-772eb4deeda3"}'),
('18a32fa4-eee1-48ed-bc2c-0475f24f7a9d', 'Stepin Vajat', '23863540-63dc-4ffb-b998-2a68ae637ba1', '{"b20c8070-c25e-4042-86a4-0d2cfc4aa505", "1c9aaedf-ace3-49f0-b764-0911f52ab2b2"}'); 

INSERT INTO user_roles VALUES
('a98959b1-8e8a-4821-9b8b-6dd7083b3ad6', 'admin'),
('03860102-7b05-4803-abb9-74622e829327', 'primary_user'),
('b04c1ac5-b7d5-494f-845b-e7c0d33389be', 'secondary_user'),
('82576de3-3d4f-4c82-a101-47ac49c39068', 'consultant');

INSERT INTO users VALUES
('c9b207f9-1b01-4653-95f0-2df358fc0e3e', 'gliga', '7six@protonmail.com', '$2a$10$a.NLtTMCJ8pQcGe73Iui3.iQNbAfT8HKVIYCssGE.kvMksrdpqH4q', '18a32fa4-eee1-48ed-bc2c-0475f24f7a9d', '03860102-7b05-4803-abb9-74622e829327', 'Dusan', 'Gligoric'),
('8d597f8e-8f2b-4cb1-88ac-ff50d9e00657', 'suza', 'dusan76shi@gmail.com', '$2a$10$a.NLtTMCJ8pQcGe73Iui3.iQNbAfT8HKVIYCssGE.kvMksrdpqH4q', '67f8f522-c1f5-4e08-b9d4-8884a21935e2', '03860102-7b05-4803-abb9-74622e829327', 'Suzana', 'Mancic'),
('8616e6fb-b46c-4de4-8986-e94cbbc2098b', 'jdoe', 'jdoe@gmail.com', '$2a$10$a.NLtTMCJ8pQcGe73Iui3.iQNbAfT8HKVIYCssGE.kvMksrdpqH4q', NULL, '82576de3-3d4f-4c82-a101-47ac49c39068', 'John', 'Doe');

INSERT INTO applications VALUES 
('a862b3bd-480b-4e05-ac08-9341fa262df8', 'XYZ-1001 Poziv za privrednike u oblasti stocarstva', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque pellentesque lacus non efficitur rutrum. Donec id diam velit. Mauris volutpat ipsum leo, vitae rutrum ipsum accumsan eu. Vivamus vitae aliquam lectus, sit amet dictum ex. In lacinia interdum posuere. Vestibulum dignissim vestibulum ligula, id ullamcorper erat tincidunt at. Etiam vitae hendrerit magna, vitae dapibus est. Sed scelerisque augue sapien. Nulla interdum elit et risus commodo, vitae ullamcorper libero scelerisque. Nam eu lorem ultricies, pellentesque lectus quis, tincidunt massa.','{"b20c8070-c25e-4042-86a4-0d2cfc4aa505", "93d612af-d256-4b08-a670-ddd162975ee3", "1c9aaedf-ace3-49f0-b764-0911f52ab2b2"}','2019-01-15 00:00:00', '2019-06-15 00:00:00', true)




'3b71be97-4a76-4369-8b7e-5fc2d1a1bdf4'
'b106206f-e705-4331-8783-f22c124a490b'
'2c9bd13a-135b-42e1-a206-ff9416a70625'
'1b8432a8-5d3f-4811-9380-14a6e9174084'
'0e362b51-ad91-4fd1-b73c-b4d8595358d2'
'c53a8530-3ad8-4e2c-aee9-dd4f40b9b246'
'1edd6fdc-e3f3-4d7f-9cea-351c150c5aa5'
'c42c3a8b-cd78-4991-a933-53b7329e19fd'
'9df139a3-fea2-4d01-b58b-21d1aa5497f6'
'3a9286b6-c6c2-4623-ad44-5730285ad493'
'b08ff18e-80fc-4c3a-a60b-2f7fb574b6c6'
'f9a775c2-3a43-4bd0-b14f-c81c6e7033b9'
'8306d23f-f2c8-4705-8990-a72eac2308c7'

