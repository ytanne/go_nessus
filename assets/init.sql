CREATE TABLE IF NOT EXISTS arp_targets (
    id INTEGER PRIMARY KEY,
    target TEXT NOT NULL,
    num_of_ips INTEGER NOT NULL DEFAULT 0,
    ips BLOB DEFAULT NULL,
    scan_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    error_status INTEGER NOT NULL DEFAULT 0,
    error_msg TEXT NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS nmap_targets (
    id INTEGER PRIMARY KEY,
    arpscan_id INTEGER NOT NULL DEFAULT -1,
    ip TEXT NOT NULL,
    result TEXT NOT NULL DEFAULT '',
    scan_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    error_status INTEGER NOT NULL DEFAULT 0,
    error_msg TEXT NOT NULL DEFAULT '',
    FOREIGN KEY(arpscan_id) REFERENCES arp_targets(id)
);

CREATE TABLE IF NOT EXISTS web_targets (
    id INTEGER PRIMARY KEY,
    arpscan_id INTEGER NOT NULL DEFAULT -1,
    ip TEXT NOT NULL,
    result TEXT NOT NULL DEFAULT '',
    scan_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    error_status INTEGER NOT NULL DEFAULT 0,
    error_msg TEXT NOT NULL DEFAULT '',
    FOREIGN KEY(arpscan_id) REFERENCES arp_targets(id)
);