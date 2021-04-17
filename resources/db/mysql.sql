CREATE TABLE provinsi (
    id   VARCHAR(36),
    nama VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (nama)
);

CREATE TABLE kabupaten_kota (
    id          VARCHAR(36),
    id_provinsi VARCHAR(36) NOT NULL,
    nama        VARCHAR (255) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (nama),
    FOREIGN KEY (id_provinsi) REFERENCES provinsi (id)
);

CREATE TABLE kecamatan (
    id                VARCHAR(36),
    id_kabupaten_kota VARCHAR(36)  NOT NULL,
    nama              VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (nama),
    FOREIGN KEY (id_kabupaten_kota) REFERENCES kabupaten_kota (id)
);

CREATE TABLE kelurahan (
    id           VARCHAR(36),
    id_kecamatan VARCHAR(36)  NOT NULL,
    nama         VARCHAR(255) NOT NULL,
    kodepos      VARCHAR(10)  NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (nama),
    FOREIGN KEY (id_kecamatan) REFERENCES kecamatan (id)
);

CREATE TABLE registrasi (
    id        VARCHAR(36),
    email     VARCHAR(255) NOT NULL,
    password  VARCHAR(255) NOT NULL,
    nama      VARCHAR(255) NOT NULL,
    alamat    VARCHAR(255) NOT NULL,
    id_kelurahan VARCHAR(36) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (email),
    FOREIGN KEY (id_kelurahan) REFERENCES kelurahan(id)
);

INSERT INTO provinsi (id, nama) VALUES ('dki', 'DKI Jakarta');
INSERT INTO provinsi (id, nama) VALUES ('jatim', 'Jawa Timur');
INSERT INTO provinsi (id, nama) VALUES ('jabar', 'Jawa Barat');

INSERT INTO kabupaten_kota (id, id_provinsi, nama) VALUES ('kabbogor', 'jabar', 'Kabupaten Bogor');
INSERT INTO kabupaten_kota (id, id_provinsi, nama) VALUES ('kotabogor', 'jabar', 'Kota Bogor');
INSERT INTO kabupaten_kota (id, id_provinsi, nama) VALUES ('surabaya', 'jatim', 'Surabaya');
INSERT INTO kabupaten_kota (id, id_provinsi, nama) VALUES ('mojokerto', 'jatim', 'Mojokerto');

INSERT INTO kecamatan (id, id_kabupaten_kota, nama) VALUES ('cibinong', 'kabbogor', 'Cibinong');
INSERT INTO kecamatan (id, id_kabupaten_kota, nama) VALUES ('gnputri', 'kabbogor', 'Gunung Putri');
INSERT INTO kecamatan (id, id_kabupaten_kota, nama) VALUES ('bootimur', 'kotabogor', 'Bogor Timur');
INSERT INTO kecamatan (id, id_kabupaten_kota, nama) VALUES ('boobarat', 'kotabogor', 'Bogor Barat');
INSERT INTO kecamatan (id, id_kabupaten_kota, nama) VALUES ('rungkut', 'surabaya', 'Rungkut');
INSERT INTO kecamatan (id, id_kabupaten_kota, nama) VALUES ('wonokromo', 'surabaya', 'Wonokromo');
INSERT INTO kecamatan (id, id_kabupaten_kota, nama) VALUES ('mojosari', 'mojokerto', 'Mojosari');
INSERT INTO kecamatan (id, id_kabupaten_kota, nama) VALUES ('trowulan', 'mojokerto', 'Trowulan');

INSERT INTO kelurahan (id, id_kecamatan, nama, kodepos) VALUES ('tengah', 'cibinong', 'Tengah', '16914');
INSERT INTO kelurahan (id, id_kecamatan, nama, kodepos) VALUES ('pakansari', 'cibinong', 'Pakansari', '16915');
INSERT INTO kelurahan (id, id_kecamatan, nama, kodepos) VALUES ('ciangsana', 'gnputri', 'Ciangsana', '16968');
INSERT INTO kelurahan (id, id_kecamatan, nama, kodepos) VALUES ('cikeas', 'gnputri', 'Cikeas', '16966');

INSERT INTO registrasi (id,email, password, nama, alamat, id_kelurahan) VALUES ('testing', 'sammidev@gm.com','myPasswords','sammidev', 'lapau durian', 'tengah');