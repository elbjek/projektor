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


CREATE TABLE checklist (
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    included_checks UUID[] NOT NULL
);

CREATE TABLE checks (
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    description text NOT NULL,
    assist text NOT NULL
);

INSERT INTO checks (id, description, assist)  VALUES
('977bec85-3c91-467e-bc5b-00f2513ae198', 'Preduzeće je registrovano u APR', ''),
('ea425c27-775e-4e56-9144-27a3d5792ddb', 'Malo ili srednje preduzeće', ''),
('f37318d4-7b76-42ce-a065-e7fa050f0207', 'Preduzeće je registrovano u opštini izovan Vojvodine i Beograda', ''),
('088fc12b-2a59-4d92-befa-c9cf1db97c02', 'Preduzeće obavlja delatnost koja ne spada u primarnu poljoprivrednu proizvodnju, proizvodnju kontrolisanih supstanci (duvan, alkohol, oružje i vojna oprema, organizovanje igara na sreću, promet bilo kojih supstanci regulisanih zakonom)', ''),
('f5042094-6b91-4d45-ac39-1a664e3e0422', 'Preduzeće je pretežno u privatnom vlasništvu', ''),
('1b76c01d-819c-45b5-837f-e22e24bb6c58', 'Preduzeće je pretežno u domaćem vlasništvu', ''),
('9da5a482-acf2-44aa-904e-93ba86f152f3', 'Preduzeće ima više od jednog zaposlenog', ''),
('ae998429-e1b7-4b28-96a2-84dbdec912cd', 'Preduzeće je ostvarilo profit u prethodnoj kalendarskoj godini', ''),
('f01d943c-df86-40d4-a291-1e233e712a05', 'Preduzeće nije u stečaju', ''),
('a95c37b4-659d-4756-9156-9f0495f08d7d', 'Preduzeće nema neisplaćena dugovanja za porez', 'Tvoje preduzeće ne ispunjava kriterijume za ovaj poziv. Kontaktiraj lokalnu filijalu poreske uprave putem linka http://www.purs.gov.rs/kontakt/adresar.html, ili pristupi portalu ePorez http://www.poreskauprava.gov.rs/e-porezi  i izmiri svoja porezna dugovanja, nakon čega možeš aplicirati'),
('072d3a65-f47b-44e2-8d97-9071cead8464', 'Preduzeće nije korisnik sredstava od drugih donatoru za isti projekat', 'Tvoje preduzeće ne ispunjava kriterijume za ovaj poziv. Pogledajt druge pozive na linku.'),
('1de98e78-0303-43ad-a27d-77ee0146fc6d', 'Preduzeće poseduje prostor ili ima ugovor o zakupu na duži rok', 'Tvoje preduzeće ne ispunjava kriterijume za ovaj poziv. Pogledajt povoljne prilike za kreditiranje uz garancije Evropskog investicionog fonda u okviru COSME programa koje nudi Societe Generale Banka https://www.societegenerale.rs/cosme-program/ (sponzored)'),
('5585e54f-e83a-4557-a111-cb29a3172747', 'Preduzeće će opremu pribavljenu tokom ovog projekta koristiti u svrhu obavljanja sadašnje delatnosti', ''),
('afa844d6-7d9b-4f05-8eea-757293a18bfa', 'Zaposleni poseduju iskostvo i ekspertizu za realizaciju projekta', 'oje preduzeće može imati teškoće u ispunjavanju uslova, kao i realizaciju projekta. Preporučujemo da iskoristiš prilike za obuke Ministarstva privrede https://www.preduzetnistvo.gov.rs/otvoreni-programi/ ili pronađeš relevantne informacije u našoj Bazi znanja (link)'),
('2a822628-8ac4-4446-b15d-76043e068009', 'Preduzeće poseduje dovoljno sredstava da obezbedi nedostajući deo od 30% vrednosti projekta', 'Iskoristi prilike povoljnih komercijalnih kredita Societe Generale banke https://www.societegenerale.rs/privreda/ (sponzored)'),
('a54e4e84-6253-4b31-9c15-f21508cc521e', 'Preduzeće ima dovoljno iskustva u razvoju projekata i apliciranju za EU fondove', 'Možeš iskoristiti prilike za unapređenje svojih veština kroz pristup našoj Bazi znanja (link), prisustvovanje treningu u okviru Događaja (link) ili angažovanjem konsultanta iz naše baze Konsultanti (link), ili apliciranjem za savetodavnu podršku Evropske banke za obnovu i razvoj https://europa.rs/otvoren-poziv-ebrd-za-savetodavnu-podrsku/');

INSERT INTO checklist (id, included_checks) VALUES 
('cc2f5eec-9dae-44a4-8656-310a43cbe239', '{"977bec85-3c91-467e-bc5b-00f2513ae198", "ea425c27-775e-4e56-9144-27a3d5792ddb", "f37318d4-7b76-42ce-a065-e7fa050f0207", "088fc12b-2a59-4d92-befa-c9cf1db97c02", "f5042094-6b91-4d45-ac39-1a664e3e0422", "1b76c01d-819c-45b5-837f-e22e24bb6c58", "9da5a482-acf2-44aa-904e-93ba86f152f3", "ae998429-e1b7-4b28-96a2-84dbdec912cd", "f01d943c-df86-40d4-a291-1e233e712a05", "a95c37b4-659d-4756-9156-9f0495f08d7d", "072d3a65-f47b-44e2-8d97-9071cead8464", "1de98e78-0303-43ad-a27d-77ee0146fc6d", "5585e54f-e83a-4557-a111-cb29a3172747", "afa844d6-7d9b-4f05-8eea-757293a18bfa", "2a822628-8ac4-4446-b15d-76043e068009", "a54e4e84-6253-4b31-9c15-f21508cc521e"}');

