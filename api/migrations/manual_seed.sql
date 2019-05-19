INSERT INTO company_fields (id, name) VALUES
('9e0ac61a-c0fa-46be-891d-c62b87a406cb', 'proizvodnja'),
('b03d5057-8369-4a4a-911b-9278f6c6dd01', 'usluga'),
('ffafd74c-9799-450f-ba6a-3affc8df5669', 'gradjevina'),
('d6ca5614-efe2-43e5-865a-a1f729b467f9', 'turizam'),
('c3b3bb5e-1705-43f4-88d2-998619d2009a', 'transport');

INSERT INTO tag_location VALUES
('a6dc93c5-6962-44fc-856d-9a2cabae4dae', 'Beograd'),
('38cd1591-bf73-4081-b7af-6695f48f82cb', 'Centralna Srbija'),
('49d64df1-9e40-4aa5-8650-9c2c4574195c', 'Nis'),
('796fb166-a23d-4e23-b88f-911a25322635', 'Novi Sad');

INSERT INTO tag_size (id, name) VALUES
('8f6f7f10-05f7-4391-8cf6-9c3cddb400f9', 'mikro'),
('02978911-ae6b-4fbc-a89b-5777056f13eb', 'mala'),
('91b28a8e-890c-494b-905d-9eb136ce796a', 'srednja'),
('1352a6ec-c878-445b-bae9-1745dd992bad', 'start-up');

INSERT INTO tag_investment (id, name) VALUES
('eced4e07-afc8-45d3-95d1-b19fe4002d3b', 'oprema'),
('5495c208-ff15-4800-911a-1df909eabc34', 'obuka'),
('811a2478-30a9-4fe8-8db5-0fc672bdbfd9', 'prostor'),
('e6f4dbe8-7759-498e-acb5-6bc8146a9ad8', 'izvoz');

INSERT INTO user_roles VALUES
('a98959b1-8e8a-4821-9b8b-6dd7083b3ad6', 'admin'),
('03860102-7b05-4803-abb9-74622e829327', 'primary_user'),
('b04c1ac5-b7d5-494f-845b-e7c0d33389be', 'secondary_user'),
('82576de3-3d4f-4c82-a101-47ac49c39068', 'consultant');

INSERT INTO users (id, username, email, password, company_id, role_id, firstname, surname, e_eu_applied_before, e_eu_funded_before, e_eu_subsidy_before, e_wrote_projects, e_enlish_fluent, e_edu_writing_projects, e_edu_project_mngmt, e_edu_financial_reporting, e_edu_eu_standards, e_edu_market_emerge) VALUES
(
    'c9b207f9-1b01-4653-95f0-2df358fc0e3e',
    'gliga',
    '7six@protonmail.com',
    '$2a$10$a.NLtTMCJ8pQcGe73Iui3.iQNbAfT8HKVIYCssGE.kvMksrdpqH4q',
    NULL,
    '03860102-7b05-4803-abb9-74622e829327',
    'Dusan',
    'Gligoric',
    false,
    false,
    false,
    false,
    true,
    false,
    true,
    true,
    false,
    false
),
(
    '8d597f8e-8f2b-4cb1-88ac-ff50d9e00657',
    'suza',
    'dusan76shi@gmail.com',
    '$2a$10$a.NLtTMCJ8pQcGe73Iui3.iQNbAfT8HKVIYCssGE.kvMksrdpqH4q',
    NULL,
    '03860102-7b05-4803-abb9-74622e829327',
    'Suzana',
    'Mancic',
    false,
    false,
    false,
    false,
    true,
    false,
    true,
    true,
    false,
    false
),
(
    '8616e6fb-b46c-4de4-8986-e94cbbc2098b',
    'jdoe',
    'jdoe@gmail.com',
    '$2a$10$a.NLtTMCJ8pQcGe73Iui3.iQNbAfT8HKVIYCssGE.kvMksrdpqH4q',
    NULL,
    '03860102-7b05-4803-abb9-74622e829327',
    'John',
    'Doe',
    false,
    false,
    false,
    false,
    true,
    false,
    true,
    true,
    false,
    false
);



INSERT INTO applications VALUES 
('a862b3bd-480b-4e05-ac08-9341fa262df8', true, 'XYZ-1001 Poziv za privrednike u oblasti stocarstva', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque pellentesque lacus non efficitur rutrum. Donec id diam velit. Mauris volutpat ipsum leo, vitae rutrum ipsum accumsan eu. Vivamus vitae aliquam lectus, sit amet dictum ex. In lacinia interdum posuere. Vestibulum dignissim vestibulum ligula, id ullamcorper erat tincidunt at. Etiam vitae hendrerit magna, vitae dapibus est. Sed scelerisque augue sapien. Nulla interdum elit et risus commodo, vitae ullamcorper libero scelerisque. Nam eu lorem ultricies, pellentesque lectus quis, tincidunt massa.','{"b20c8070-c25e-4042-86a4-0d2cfc4aa505", "93d612af-d256-4b08-a670-ddd162975ee3", "1c9aaedf-ace3-49f0-b764-0911f52ab2b2"}','2019-01-15 00:00:00', '2019-06-15 00:00:00');


INSERT INTO companies (id, owner_id, name, created_at, employees, field_id, t_size, t_location, t_investment, t_pib, t_mib) VALUES  
(
    '67f8f522-c1f5-4e08-b9d4-8884a21935e2',
    'c9b207f9-1b01-4653-95f0-2df358fc0e3e',
    'ISWiB',
    '2006-06-15 00:00:00',
    100,
    'd6ca5614-efe2-43e5-865a-a1f729b467f9',
    '8f6f7f10-05f7-4391-8cf6-9c3cddb400f9',
    'a6dc93c5-6962-44fc-856d-9a2cabae4dae',
    'eced4e07-afc8-45d3-95d1-b19fe4002d3b',
    100000,
    10000000
),
(
    '81adc65e-8455-4752-8b87-0c6c8fcb604a',
    '8d597f8e-8f2b-4cb1-88ac-ff50d9e00657',
    'Frutina d.o.o.',
    '2011-05-15 00:00:00',
    15,
    '9e0ac61a-c0fa-46be-891d-c62b87a406cb',
    '02978911-ae6b-4fbc-a89b-5777056f13eb',
    '49d64df1-9e40-4aa5-8650-9c2c4574195c',
    '5495c208-ff15-4800-911a-1df909eabc34',
    200000,
    20000000
),
(
    '583d994e-d012-4b0c-956a-eebb1b80b034',
    '8616e6fb-b46c-4de4-8986-e94cbbc2098b',
    'Delta Holding',
    '1995-12-30 00:00:00',
    20,
    'b03d5057-8369-4a4a-911b-9278f6c6dd01',
    '02978911-ae6b-4fbc-a89b-5777056f13eb',
    '49d64df1-9e40-4aa5-8650-9c2c4574195c',
    '811a2478-30a9-4fe8-8db5-0fc672bdbfd9',
    300000,
    30000000
),
(
    '18a32fa4-eee1-48ed-bc2c-0475f24f7a9d',
    '8616e6fb-b46c-4de4-8986-e94cbbc2098b',
    'Stepin Vajat',
    '1992-06-24 00:00:00',
    40,
    'b03d5057-8369-4a4a-911b-9278f6c6dd01',
    '02978911-ae6b-4fbc-a89b-5777056f13eb',
    '49d64df1-9e40-4aa5-8650-9c2c4574195c',
    'e6f4dbe8-7759-498e-acb5-6bc8146a9ad8',
    400000,
    40000000
); 

UPDATE users SET company_id='67f8f522-c1f5-4e08-b9d4-8884a21935e2' WHERE id='c9b207f9-1b01-4653-95f0-2df358fc0e3e';
UPDATE users SET company_id='81adc65e-8455-4752-8b87-0c6c8fcb604a' WHERE id='8d597f8e-8f2b-4cb1-88ac-ff50d9e00657';
UPDATE users SET company_id='18a32fa4-eee1-48ed-bc2c-0475f24f7a9d' WHERE id='8616e6fb-b46c-4de4-8986-e94cbbc2098b';