CREATE TABLE IF NOT EXISTS stock (
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deletedAt TIMESTAMP NULL,
  id VARCHAR(36) PRIMARY KEY NOT NULL,
  name VARCHAR(255) NOT NULL,
  cusip VARCHAR(20) NOT NULL,
  ticker VARCHAR(20) NOT NULL,
  exchange VARCHAR(50) NOT NULL,
  cik VARCHAR(20) NOT NULL,
  sic VARCHAR(20) NOT NULL,
  industry VARCHAR(255) NOT NULL,
  sector VARCHAR(255) NOT NULL,
  isDelisted BOOLEAN NOT NULL,
  location VARCHAR(255) NOT NULL,
  secId VARCHAR(100) NOT NULL,
  sicSector VARCHAR(255) NOT NULL,
  sicIndustry VARCHAR(255) NOT NULL,
  famaIndustry VARCHAR(255) NOT NULL,
  currency VARCHAR(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;